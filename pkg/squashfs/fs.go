// Copyright 2025 Chainguard, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package squashfs

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/diskfs/go-diskfs"
	"github.com/diskfs/go-diskfs/filesystem"
	"github.com/diskfs/go-diskfs/filesystem/squashfs"

	apkfs "chainguard.dev/apko/pkg/apk/fs"
)

// memFS is an in-memory implementation that will be converted to SquashFS
type memFS struct {
	tree *node
	mu   sync.RWMutex
}

type node struct {
	name     string
	mode     fs.FileMode
	uid, gid int
	dir      bool
	data     []byte
	modTime  time.Time
	children map[string]*node
	mu       sync.RWMutex
	xattrs   map[string][]byte
}

// New creates a new in-memory filesystem that can be converted to SquashFS
func New() *memFS {
	return &memFS{
		tree: &node{
			dir:      true,
			children: map[string]*node{},
			xattrs:   map[string][]byte{},
			name:     "/",
			mode:     fs.ModeDir | 0o755,
			modTime:  time.Now(),
		},
	}
}

// Mkdir creates a directory with the given name and permissions
func (m *memFS) Mkdir(path string, perm fs.FileMode) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	parent := filepath.Dir(path)
	anode, err := m.getNode(parent)
	if err != nil {
		return err
	}
	if !anode.dir {
		return fmt.Errorf("parent is not a directory")
	}
	
	anode.mu.Lock()
	defer anode.mu.Unlock()
	
	base := filepath.Base(path)
	if _, ok := anode.children[base]; ok {
		return fs.ErrExist
	}
	
	anode.children[base] = &node{
		name:     base,
		mode:     fs.ModeDir | perm,
		dir:      true,
		children: map[string]*node{},
		xattrs:   map[string][]byte{},
		modTime:  time.Now(),
	}
	return nil
}

// MkdirAll creates all directories in the path
func (m *memFS) MkdirAll(path string, perm fs.FileMode) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	parts := strings.Split(strings.Trim(path, "/"), "/")
	current := m.tree
	
	for _, part := range parts {
		if part == "" {
			continue
		}
		
		current.mu.Lock()
		if current.children == nil {
			current.children = map[string]*node{}
		}
		
		child, exists := current.children[part]
		if !exists {
			child = &node{
				name:     part,
				mode:     fs.ModeDir | perm,
				dir:      true,
				children: map[string]*node{},
				xattrs:   map[string][]byte{},
				modTime:  time.Now(),
			}
			current.children[part] = child
		}
		current.mu.Unlock()
		
		if !child.dir {
			return fmt.Errorf("path component is not a directory")
		}
		current = child
	}
	return nil
}

// getNode retrieves a node by path
func (m *memFS) getNode(path string) (*node, error) {
	if path == "/" || path == "." {
		return m.tree, nil
	}
	
	parts := strings.Split(strings.Trim(path, "/"), "/")
	current := m.tree
	
	for _, part := range parts {
		if part == "" {
			continue
		}
		
		current.mu.RLock()
		child, exists := current.children[part]
		current.mu.RUnlock()
		
		if !exists {
			return nil, fs.ErrNotExist
		}
		current = child
	}
	return current, nil
}

// Open opens a file for reading
func (m *memFS) Open(name string) (fs.File, error) {
	return m.OpenFile(name, os.O_RDONLY, 0)
}

// OpenReaderAt opens a file that supports ReaderAt
func (m *memFS) OpenReaderAt(name string) (apkfs.File, error) {
	return m.OpenFile(name, os.O_RDONLY, 0)
}

// OpenFile opens a file with specified flags and permissions
func (m *memFS) OpenFile(name string, flag int, perm fs.FileMode) (apkfs.File, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(name)
	if err != nil {
		if flag&os.O_CREATE == 0 {
			return nil, err
		}
		// Create the file
		return m.createFile(name, flag, perm)
	}
	
	if node.dir && flag&os.O_CREATE == 0 {
		return &memFile{node: node, name: name, fs: m}, nil
	}
	
	return &memFile{node: node, name: name, fs: m, flag: flag}, nil
}

// createFile creates a new file
func (m *memFS) createFile(name string, flag int, perm fs.FileMode) (apkfs.File, error) {
	parent := filepath.Dir(name)
	base := filepath.Base(name)
	
	parentNode, err := m.getNode(parent)
	if err != nil {
		return nil, err
	}
	
	if !parentNode.dir {
		return nil, fmt.Errorf("parent is not a directory")
	}
	
	parentNode.mu.Lock()
	defer parentNode.mu.Unlock()
	
	if parentNode.children == nil {
		parentNode.children = map[string]*node{}
	}
	
	node := &node{
		name:    base,
		mode:    perm,
		dir:     false,
		data:    []byte{},
		xattrs:  map[string][]byte{},
		modTime: time.Now(),
	}
	
	parentNode.children[base] = node
	return &memFile{node: node, name: name, fs: m, flag: flag}, nil
}

// ReadFile reads the entire file
func (m *memFS) ReadFile(name string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(name)
	if err != nil {
		return nil, err
	}
	
	if node.dir {
		return nil, fmt.Errorf("is a directory")
	}
	
	node.mu.RLock()
	defer node.mu.RUnlock()
	return append([]byte{}, node.data...), nil
}

// WriteFile writes data to a file
func (m *memFS) WriteFile(name string, data []byte, mode fs.FileMode) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	node, err := m.getNode(name)
	if err != nil {
		// Create the file
		if err := m.MkdirAll(filepath.Dir(name), 0o755); err != nil {
			return err
		}
		node, err = m.createFileNode(name, mode)
		if err != nil {
			return err
		}
	}
	
	if node.dir {
		return fmt.Errorf("is a directory")
	}
	
	node.mu.Lock()
	defer node.mu.Unlock()
	node.data = append([]byte{}, data...)
	node.modTime = time.Now()
	return nil
}

// createFileNode creates a new file node
func (m *memFS) createFileNode(name string, mode fs.FileMode) (*node, error) {
	parent := filepath.Dir(name)
	base := filepath.Base(name)
	
	parentNode, err := m.getNode(parent)
	if err != nil {
		return nil, err
	}
	
	parentNode.mu.Lock()
	defer parentNode.mu.Unlock()
	
	if parentNode.children == nil {
		parentNode.children = map[string]*node{}
	}
	
	node := &node{
		name:    base,
		mode:    mode,
		dir:     false,
		data:    []byte{},
		xattrs:  map[string][]byte{},
		modTime: time.Now(),
	}
	
	parentNode.children[base] = node
	return node, nil
}

// ReadDir reads directory entries
func (m *memFS) ReadDir(name string) ([]fs.DirEntry, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(name)
	if err != nil {
		return nil, err
	}
	
	if !node.dir {
		return nil, fmt.Errorf("not a directory")
	}
	
	node.mu.RLock()
	defer node.mu.RUnlock()
	
	entries := make([]fs.DirEntry, 0, len(node.children))
	for _, child := range node.children {
		entries = append(entries, &dirEntry{node: child})
	}
	
	return entries, nil
}

// Stat returns file information
func (m *memFS) Stat(name string) (fs.FileInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(name)
	if err != nil {
		return nil, err
	}
	
	return &fileInfo{node: node, name: filepath.Base(name)}, nil
}

// Lstat returns file information (same as Stat for this implementation)
func (m *memFS) Lstat(name string) (fs.FileInfo, error) {
	return m.Stat(name)
}

// Create creates a new file
func (m *memFS) Create(name string) (apkfs.File, error) {
	return m.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0o666)
}

// Remove removes a file or directory
func (m *memFS) Remove(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	parent := filepath.Dir(name)
	base := filepath.Base(name)
	
	parentNode, err := m.getNode(parent)
	if err != nil {
		return err
	}
	
	parentNode.mu.Lock()
	defer parentNode.mu.Unlock()
	
	if _, exists := parentNode.children[base]; !exists {
		return fs.ErrNotExist
	}
	
	delete(parentNode.children, base)
	return nil
}

// Chmod changes file permissions
func (m *memFS) Chmod(path string, perm fs.FileMode) error {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(path)
	if err != nil {
		return err
	}
	
	node.mu.Lock()
	defer node.mu.Unlock()
	node.mode = perm | (node.mode & fs.ModeType)
	return nil
}

// Chown changes file ownership
func (m *memFS) Chown(path string, uid, gid int) error {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(path)
	if err != nil {
		return err
	}
	
	node.mu.Lock()
	defer node.mu.Unlock()
	node.uid = uid
	node.gid = gid
	return nil
}

// Chtimes changes file times
func (m *memFS) Chtimes(path string, atime, mtime time.Time) error {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(path)
	if err != nil {
		return err
	}
	
	node.mu.Lock()
	defer node.mu.Unlock()
	node.modTime = mtime
	return nil
}

// Extended attributes methods (simplified implementations)
func (m *memFS) SetXattr(path string, attr string, data []byte) error {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(path)
	if err != nil {
		return err
	}
	
	node.mu.Lock()
	defer node.mu.Unlock()
	node.xattrs[attr] = append([]byte{}, data...)
	return nil
}

func (m *memFS) GetXattr(path string, attr string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(path)
	if err != nil {
		return nil, err
	}
	
	node.mu.RLock()
	defer node.mu.RUnlock()
	data, exists := node.xattrs[attr]
	if !exists {
		return nil, os.ErrNotExist
	}
	return append([]byte{}, data...), nil
}

func (m *memFS) RemoveXattr(path string, attr string) error {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(path)
	if err != nil {
		return err
	}
	
	node.mu.Lock()
	defer node.mu.Unlock()
	delete(node.xattrs, attr)
	return nil
}

func (m *memFS) ListXattrs(path string) (map[string][]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(path)
	if err != nil {
		return nil, err
	}
	
	node.mu.RLock()
	defer node.mu.RUnlock()
	
	result := make(map[string][]byte)
	for k, v := range node.xattrs {
		result[k] = append([]byte{}, v...)
	}
	return result, nil
}

// Symlink creates a symbolic link
func (m *memFS) Symlink(oldname, newname string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	parent := filepath.Dir(newname)
	base := filepath.Base(newname)
	
	parentNode, err := m.getNode(parent)
	if err != nil {
		return err
	}
	
	parentNode.mu.Lock()
	defer parentNode.mu.Unlock()
	
	if parentNode.children == nil {
		parentNode.children = map[string]*node{}
	}
	
	if _, exists := parentNode.children[base]; exists {
		return fs.ErrExist
	}
	
	parentNode.children[base] = &node{
		name:    base,
		mode:    0o777 | fs.ModeSymlink,
		dir:     false,
		data:    []byte(oldname),
		xattrs:  map[string][]byte{},
		modTime: time.Now(),
	}
	
	return nil
}

// Link creates a hard link
func (m *memFS) Link(oldname, newname string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	// Get the target node
	targetNode, err := m.getNode(oldname)
	if err != nil {
		return err
	}
	
	parent := filepath.Dir(newname)
	base := filepath.Base(newname)
	
	parentNode, err := m.getNode(parent)
	if err != nil {
		return err
	}
	
	parentNode.mu.Lock()
	defer parentNode.mu.Unlock()
	
	if _, exists := parentNode.children[base]; exists {
		return fs.ErrExist
	}
	
	parentNode.children[base] = targetNode
	return nil
}

// Readlink reads the target of a symbolic link
func (m *memFS) Readlink(name string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(name)
	if err != nil {
		return "", err
	}
	
	if node.mode&fs.ModeSymlink == 0 {
		return "", fmt.Errorf("file is not a symbolic link")
	}
	
	node.mu.RLock()
	defer node.mu.RUnlock()
	return string(node.data), nil
}

// Device file operations (simplified)
func (m *memFS) Mknod(path string, mode uint32, dev int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	parent := filepath.Dir(path)
	base := filepath.Base(path)
	
	parentNode, err := m.getNode(parent)
	if err != nil {
		return err
	}
	
	parentNode.mu.Lock()
	defer parentNode.mu.Unlock()
	
	if _, exists := parentNode.children[base]; exists {
		return fs.ErrExist
	}
	
	parentNode.children[base] = &node{
		name:    base,
		mode:    fs.FileMode(mode) | fs.ModeDevice,
		dir:     false,
		xattrs:  map[string][]byte{},
		modTime: time.Now(),
	}
	
	return nil
}

func (m *memFS) Readnod(name string) (dev int, err error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	node, err := m.getNode(name)
	if err != nil {
		return 0, err
	}
	
	if node.mode&fs.ModeDevice == 0 {
		return 0, fmt.Errorf("not a device file")
	}
	
	return 0, nil // Simplified - return 0 device number
}

// Sub returns a sub-filesystem
func (m *memFS) Sub(path string) (apkfs.FullFS, error) {
	cleanPath := filepath.Clean(path)
	if cleanPath == "." {
		return m, nil
	}
	
	info, err := m.Stat(cleanPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, errors.New("not a directory")
	}
	
	return &apkfs.SubFS{
		FS:   m,
		Root: cleanPath,
	}, nil
}

// WriteToSquashFS writes the filesystem to a SquashFS image
func (m *memFS) WriteToSquashFS(imagePath string, sizeHint int64) error {
	// Calculate filesystem size if not provided
	if sizeHint <= 0 {
		sizeHint = m.calculateSize() * 2 // Add some padding
		if sizeHint < 1024*1024 { // Minimum 1MB
			sizeHint = 1024 * 1024
		}
	}
	
	// Create SquashFS disk image
	disk, err := diskfs.Create(imagePath, sizeHint, diskfs.Raw, diskfs.SectorSizeDefault)
	if err != nil {
		return fmt.Errorf("creating disk: %w", err)
	}
	defer disk.Close()
	
	fs, err := disk.CreateFilesystem(diskfs.FilesystemSpec{
		Partition:   0,
		FSType:      filesystem.TypeSquashfs,
		VolumeLabel: "apko-squashfs",
	})
	if err != nil {
		return fmt.Errorf("creating filesystem: %w", err)
	}
	
	// Write all files to the SquashFS
	if err := m.writeNodeToSquashFS(fs, m.tree, "/"); err != nil {
		return fmt.Errorf("writing files to squashfs: %w", err)
	}
	
	// Finalize the SquashFS
	squashFS, ok := fs.(*squashfs.FileSystem)
	if !ok {
		return fmt.Errorf("filesystem is not squashfs")
	}
	
	if err := squashFS.Finalize(squashfs.FinalizeOptions{}); err != nil {
		return fmt.Errorf("finalizing squashfs: %w", err)
	}
	
	return nil
}

// calculateSize estimates the total size needed for the filesystem
func (m *memFS) calculateSize() int64 {
	return m.calculateNodeSize(m.tree)
}

func (m *memFS) calculateNodeSize(node *node) int64 {
	node.mu.RLock()
	defer node.mu.RUnlock()
	
	size := int64(len(node.data))
	for _, child := range node.children {
		size += m.calculateNodeSize(child)
	}
	return size
}

// writeNodeToSquashFS recursively writes nodes to the SquashFS filesystem
func (m *memFS) writeNodeToSquashFS(fs filesystem.FileSystem, node *node, path string) error {
	node.mu.RLock()
	defer node.mu.RUnlock()
	
	if node.dir {
		// Create directory
		if path != "/" {
			if err := fs.Mkdir(path); err != nil && !os.IsExist(err) {
				return fmt.Errorf("creating directory %s: %w", path, err)
			}
		}
		
		// Write children
		for name, child := range node.children {
			childPath := filepath.Join(path, name)
			if path == "/" {
				childPath = "/" + name
			}
			if err := m.writeNodeToSquashFS(fs, child, childPath); err != nil {
				return err
			}
		}
	} else {
		// Handle files
		if node.mode&fs.ModeSymlink != 0 {
			// Create symbolic link
			target := string(node.data)
			if err := fs.Symlink(target, path); err != nil {
				return fmt.Errorf("creating symlink %s -> %s: %w", path, target, err)
			}
		} else {
			// Create regular file
			file, err := fs.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
			if err != nil {
				return fmt.Errorf("creating file %s: %w", path, err)
			}
			
			if len(node.data) > 0 {
				if _, err := file.Write(node.data); err != nil {
					file.Close()
					return fmt.Errorf("writing file %s: %w", path, err)
				}
			}
			
			if err := file.Close(); err != nil {
				return fmt.Errorf("closing file %s: %w", path, err)
			}
		}
	}
	
	return nil
}


// memFile implements apkfs.File
type memFile struct {
	node   *node
	name   string
	fs     *memFS
	offset int64
	flag   int
}

func (f *memFile) Read(p []byte) (n int, err error) {
	if f.node.dir {
		return 0, fmt.Errorf("is a directory")
	}
	
	f.node.mu.RLock()
	defer f.node.mu.RUnlock()
	
	if f.offset >= int64(len(f.node.data)) {
		return 0, io.EOF
	}
	
	n = copy(p, f.node.data[f.offset:])
	f.offset += int64(n)
	return n, nil
}

func (f *memFile) ReadAt(p []byte, off int64) (n int, err error) {
	if f.node.dir {
		return 0, fmt.Errorf("is a directory")
	}
	
	f.node.mu.RLock()
	defer f.node.mu.RUnlock()
	
	if off >= int64(len(f.node.data)) {
		return 0, io.EOF
	}
	
	n = copy(p, f.node.data[off:])
	return n, nil
}

func (f *memFile) Write(p []byte) (n int, err error) {
	if f.node.dir {
		return 0, fmt.Errorf("is a directory")
	}
	
	f.node.mu.Lock()
	defer f.node.mu.Unlock()
	
	if f.flag&os.O_APPEND != 0 {
		f.node.data = append(f.node.data, p...)
		f.offset = int64(len(f.node.data))
	} else {
		if f.offset+int64(len(p)) > int64(len(f.node.data)) {
			newData := make([]byte, f.offset+int64(len(p)))
			copy(newData, f.node.data)
			f.node.data = newData
		}
		copy(f.node.data[f.offset:], p)
		f.offset += int64(len(p))
	}
	
	f.node.modTime = time.Now()
	return len(p), nil
}

func (f *memFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.offset = offset
	case io.SeekCurrent:
		f.offset += offset
	case io.SeekEnd:
		f.node.mu.RLock()
		f.offset = int64(len(f.node.data)) + offset
		f.node.mu.RUnlock()
	default:
		return 0, fmt.Errorf("invalid whence")
	}
	return f.offset, nil
}

func (f *memFile) Close() error {
	return nil
}

func (f *memFile) Stat() (fs.FileInfo, error) {
	return &fileInfo{node: f.node, name: f.name}, nil
}

// fileInfo implements fs.FileInfo
type fileInfo struct {
	node *node
	name string
}

func (fi *fileInfo) Name() string {
	return fi.name
}

func (fi *fileInfo) Size() int64 {
	fi.node.mu.RLock()
	defer fi.node.mu.RUnlock()
	return int64(len(fi.node.data))
}

func (fi *fileInfo) Mode() fs.FileMode {
	fi.node.mu.RLock()
	defer fi.node.mu.RUnlock()
	return fi.node.mode
}

func (fi *fileInfo) ModTime() time.Time {
	fi.node.mu.RLock()
	defer fi.node.mu.RUnlock()
	return fi.node.modTime
}

func (fi *fileInfo) IsDir() bool {
	return fi.node.dir
}

func (fi *fileInfo) Sys() interface{} {
	return nil
}

// dirEntry implements fs.DirEntry
type dirEntry struct {
	node *node
}

func (de *dirEntry) Name() string {
	return de.node.name
}

func (de *dirEntry) IsDir() bool {
	return de.node.dir
}

func (de *dirEntry) Type() fs.FileMode {
	de.node.mu.RLock()
	defer de.node.mu.RUnlock()
	return de.node.mode.Type()
}

func (de *dirEntry) Info() (fs.FileInfo, error) {
	return &fileInfo{node: de.node, name: de.node.name}, nil
}