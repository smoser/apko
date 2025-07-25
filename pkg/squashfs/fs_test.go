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
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestNewMemFS(t *testing.T) {
	memfs := New()
	if memfs == nil {
		t.Fatal("New() returned nil")
	}
	
	// Test root directory exists
	info, err := memfs.Stat("/")
	if err != nil {
		t.Fatal("Root directory should exist:", err)
	}
	if !info.IsDir() {
		t.Fatal("Root should be a directory")
	}
}

func TestMkdirAll(t *testing.T) {
	memfs := New()
	
	// Create nested directories
	err := memfs.MkdirAll("/usr/local/bin", 0o755)
	if err != nil {
		t.Fatal("MkdirAll failed:", err)
	}
	
	// Verify directories exist
	info, err := memfs.Stat("/usr")
	if err != nil {
		t.Fatal("Failed to stat /usr:", err)
	}
	if !info.IsDir() {
		t.Fatal("/usr should be a directory")
	}
	
	info, err = memfs.Stat("/usr/local/bin")
	if err != nil {
		t.Fatal("Failed to stat /usr/local/bin:", err)
	}
	if !info.IsDir() {
		t.Fatal("/usr/local/bin should be a directory")
	}
}

func TestWriteReadFile(t *testing.T) {
	memfs := New()
	
	// Create a directory first
	err := memfs.MkdirAll("/etc", 0o755)
	if err != nil {
		t.Fatal("MkdirAll failed:", err)
	}
	
	// Write a file
	content := []byte("Hello, SquashFS!")
	err = memfs.WriteFile("/etc/test.txt", content, 0o644)
	if err != nil {
		t.Fatal("WriteFile failed:", err)
	}
	
	// Read the file back
	readContent, err := memfs.ReadFile("/etc/test.txt")
	if err != nil {
		t.Fatal("ReadFile failed:", err)
	}
	
	if string(readContent) != string(content) {
		t.Fatalf("Content mismatch: expected %q, got %q", content, readContent)
	}
}

func TestSymlink(t *testing.T) {
	memfs := New()
	
	// Create directories
	err := memfs.MkdirAll("/usr/bin", 0o755)
	if err != nil {
		t.Fatal("MkdirAll failed:", err)
	}
	
	// Create a file
	err = memfs.WriteFile("/usr/bin/app", []byte("#!/bin/sh\necho hello"), 0o755)
	if err != nil {
		t.Fatal("WriteFile failed:", err)
	}
	
	// Create a symlink
	err = memfs.Symlink("/usr/bin/app", "/usr/bin/myapp")
	if err != nil {
		t.Fatal("Symlink failed:", err)
	}
	
	// Read the symlink
	target, err := memfs.Readlink("/usr/bin/myapp")
	if err != nil {
		t.Fatal("Readlink failed:", err)
	}
	
	if target != "/usr/bin/app" {
		t.Fatalf("Symlink target mismatch: expected %q, got %q", "/usr/bin/app", target)
	}
}

func TestWriteToSquashFS(t *testing.T) {
	memfs := New()
	
	// Create some test data
	err := memfs.MkdirAll("/usr/local/bin", 0o755)
	if err != nil {
		t.Fatal("MkdirAll failed:", err)
	}
	
	err = memfs.WriteFile("/usr/local/bin/hello", []byte("#!/bin/sh\necho 'Hello, World!'"), 0o755)
	if err != nil {
		t.Fatal("WriteFile failed:", err)
	}
	
	err = memfs.Symlink("/usr/local/bin/hello", "/usr/local/bin/hi")
	if err != nil {
		t.Fatal("Symlink failed:", err)
	}
	
	// Create temporary file for SquashFS output
	tempDir := t.TempDir()
	squashfsPath := filepath.Join(tempDir, "test.squashfs")
	
	// Write to SquashFS
	err = memfs.WriteToSquashFS(squashfsPath, 0)
	if err != nil {
		t.Fatal("WriteToSquashFS failed:", err)
	}
	
	// Verify the file was created
	info, err := os.Stat(squashfsPath)
	if err != nil {
		t.Fatal("SquashFS file was not created:", err)
	}
	
	if info.Size() == 0 {
		t.Fatal("SquashFS file is empty")
	}
	
	t.Logf("SquashFS file created successfully: %s (size: %d bytes)", squashfsPath, info.Size())
}

func TestChmod(t *testing.T) {
	memfs := New()
	
	// Create a file
	err := memfs.WriteFile("/test.txt", []byte("test content"), 0o644)
	if err != nil {
		t.Fatal("WriteFile failed:", err)
	}
	
	// Change permissions
	err = memfs.Chmod("/test.txt", 0o755)
	if err != nil {
		t.Fatal("Chmod failed:", err)
	}
	
	// Verify permissions changed
	info, err := memfs.Stat("/test.txt")
	if err != nil {
		t.Fatal("Stat failed:", err)
	}
	
	if info.Mode().Perm() != 0o755 {
		t.Fatalf("Permission mismatch: expected %o, got %o", 0o755, info.Mode().Perm())
	}
}

func TestChown(t *testing.T) {
	memfs := New()
	
	// Create a file
	err := memfs.WriteFile("/test.txt", []byte("test content"), 0o644)
	if err != nil {
		t.Fatal("WriteFile failed:", err)
	}
	
	// Change ownership (this is a no-op test since we can't easily verify ownership in the test)
	err = memfs.Chown("/test.txt", 1000, 1000)
	if err != nil {
		t.Fatal("Chown failed:", err)
	}
}

func TestXattrs(t *testing.T) {
	memfs := New()
	
	// Create a file
	err := memfs.WriteFile("/test.txt", []byte("test content"), 0o644)
	if err != nil {
		t.Fatal("WriteFile failed:", err)
	}
	
	// Set extended attribute
	err = memfs.SetXattr("/test.txt", "user.test", []byte("test value"))
	if err != nil {
		t.Fatal("SetXattr failed:", err)
	}
	
	// Get extended attribute
	value, err := memfs.GetXattr("/test.txt", "user.test")
	if err != nil {
		t.Fatal("GetXattr failed:", err)
	}
	
	if string(value) != "test value" {
		t.Fatalf("Xattr value mismatch: expected %q, got %q", "test value", value)
	}
	
	// List extended attributes
	xattrs, err := memfs.ListXattrs("/test.txt")
	if err != nil {
		t.Fatal("ListXattrs failed:", err)
	}
	
	if len(xattrs) != 1 {
		t.Fatalf("Expected 1 xattr, got %d", len(xattrs))
	}
	
	if string(xattrs["user.test"]) != "test value" {
		t.Fatalf("Xattr value mismatch in list: expected %q, got %q", "test value", xattrs["user.test"])
	}
	
	// Remove extended attribute
	err = memfs.RemoveXattr("/test.txt", "user.test")
	if err != nil {
		t.Fatal("RemoveXattr failed:", err)
	}
	
	// Verify it's gone
	_, err = memfs.GetXattr("/test.txt", "user.test")
	if err == nil {
		t.Fatal("Expected error after removing xattr")
	}
}

func TestReadDir(t *testing.T) {
	memfs := New()
	
	// Create directory structure
	err := memfs.MkdirAll("/usr/bin", 0o755)
	if err != nil {
		t.Fatal("MkdirAll failed:", err)
	}
	
	// Create some files
	err = memfs.WriteFile("/usr/bin/app1", []byte("app1 content"), 0o755)
	if err != nil {
		t.Fatal("WriteFile failed:", err)
	}
	
	err = memfs.WriteFile("/usr/bin/app2", []byte("app2 content"), 0o755)
	if err != nil {
		t.Fatal("WriteFile failed:", err)
	}
	
	// Create subdirectory
	err = memfs.Mkdir("/usr/bin/subdir", 0o755)
	if err != nil {
		t.Fatal("Mkdir failed:", err)
	}
	
	// Read directory
	entries, err := memfs.ReadDir("/usr/bin")
	if err != nil {
		t.Fatal("ReadDir failed:", err)
	}
	
	if len(entries) != 3 {
		t.Fatalf("Expected 3 entries, got %d", len(entries))
	}
	
	// Verify entries
	names := make(map[string]bool)
	for _, entry := range entries {
		names[entry.Name()] = entry.IsDir()
	}
	
	if !names["app1"] || names["app1"] {
		t.Fatal("app1 should be a file")
	}
	if !names["app2"] || names["app2"] {
		t.Fatal("app2 should be a file")
	}
	if !names["subdir"] || !names["subdir"] {
		t.Fatal("subdir should be a directory")
	}
}

func TestOpenFile(t *testing.T) {
	memfs := New()
	
	// Create and write to a file using OpenFile
	file, err := memfs.OpenFile("/test.txt", os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		t.Fatal("OpenFile failed:", err)
	}
	
	content := []byte("Hello, World!")
	_, err = file.Write(content)
	if err != nil {
		t.Fatal("Write failed:", err)
	}
	
	err = file.Close()
	if err != nil {
		t.Fatal("Close failed:", err)
	}
	
	// Read the file back
	readContent, err := memfs.ReadFile("/test.txt")
	if err != nil {
		t.Fatal("ReadFile failed:", err)
	}
	
	if string(readContent) != string(content) {
		t.Fatalf("Content mismatch: expected %q, got %q", content, readContent)
	}
}