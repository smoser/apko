# SquashFS Support in apko

This document describes the SquashFS layer format support that has been added to apko, allowing users to create container image layers in SquashFS format instead of the traditional tar.gz format.

## Overview

SquashFS is a compressed read-only filesystem that can provide better compression ratios and faster access times compared to tar.gz archives. This implementation allows apko to create OCI-compatible container image layers using SquashFS format.

## Implementation Details

### 1. SquashFS Filesystem Implementation (`pkg/squashfs/fs.go`)

The SquashFS support is built around a complete filesystem implementation that provides:

- **Full Interface Compatibility**: Implements the `apkfs.FullFS` interface, ensuring compatibility with existing apko build processes
- **In-Memory Storage**: Uses an in-memory tree structure that can be exported to SquashFS format
- **Thread Safety**: All operations are protected with appropriate mutex usage for concurrent access
- **Complete Feature Set**: Supports all required filesystem operations including:
  - Directory creation and management
  - File reading and writing
  - Symbolic links (in-memory, with export limitations)
  - File permissions and ownership
  - Extended attributes (xattrs)
  - Device files

### 2. Build System Integration

The SquashFS support is seamlessly integrated into the existing apko build system:

- **Options System**: Added `UseSquashFS` flag to the build options
- **Layer Creation**: Modified layer creation process to choose between tar.gz and SquashFS based on configuration
- **OCI Compliance**: SquashFS layers use the media type `application/vnd.oci.image.layer.squashfs`
- **Backward Compatibility**: Existing tar.gz functionality remains unchanged when SquashFS is not enabled

### 3. Command Line Interface

Users can enable SquashFS layer creation using the `--squashfs` flag:

```bash
apko build --squashfs config.yaml my-image:latest output.tar
```

### 4. Dependencies

The implementation uses the `github.com/diskfs/go-diskfs` library for SquashFS image creation, which provides:
- Cross-platform SquashFS support
- Efficient compression
- Standard SquashFS format compliance

## Usage Examples

### Basic Usage

Create a container image with SquashFS layers:

```bash
apko build --squashfs alpine-config.yaml my-alpine:latest alpine-squashfs.tar
```

### With Additional Options

Combine SquashFS with other apko features:

```bash
apko build --squashfs --arch amd64,arm64 --sbom config.yaml multi-arch:latest output/
```

## Technical Specifications

### File Format
- **Container Format**: Standard OCI image format
- **Layer Format**: SquashFS compressed filesystem
- **Media Type**: `application/vnd.oci.image.layer.squashfs`
- **Compression**: Built-in SquashFS compression (typically better than gzip)

### Performance Characteristics
- **Space Efficiency**: SquashFS typically provides better compression ratios than tar.gz
- **Access Speed**: Faster random access to files within layers
- **Memory Usage**: In-memory filesystem during build, then exported to SquashFS

### Compatibility
- **OCI Compliance**: Layers are OCI-compatible with custom media type
- **Runtime Support**: Requires container runtime support for SquashFS layers
- **Backward Compatibility**: Can be used alongside traditional tar.gz layers

## Limitations and Known Issues

### Current Limitations

1. **Symlink Export**: Due to limitations in the go-diskfs library, symbolic links are currently skipped during SquashFS export. The symlinks exist and function correctly in the in-memory filesystem during the build process, but are not included in the final SquashFS image.

2. **Sector Size**: SquashFS creation requires a minimum 4K sector size, which is automatically configured.

3. **Runtime Support**: Container runtimes must support SquashFS layer format to use the resulting images.

### Future Improvements

- **Symlink Support**: Will be added when the underlying go-diskfs library implements symlink creation for SquashFS
- **Optimization**: Additional compression and performance optimizations
- **Extended Features**: Support for additional SquashFS-specific features

## Testing

The implementation includes comprehensive tests covering:

- Basic filesystem operations (create, read, write, delete)
- Directory management (mkdir, readdir)
- File permissions and ownership
- Extended attributes
- Symlink operations (in-memory)
- SquashFS image creation and validation

All tests can be run with:

```bash
go test ./pkg/squashfs -v
```

## Architecture

### Component Structure

```
pkg/squashfs/
├── fs.go         # Main SquashFS filesystem implementation
└── fs_test.go    # Comprehensive test suite

pkg/build/
├── build_implementation.go  # SquashFS layer creation logic
└── options.go              # Build options for SquashFS

internal/cli/
└── build.go      # Command line interface integration
```

### Integration Points

1. **Filesystem Selection**: Build system chooses between `tarfs.New()` and `squashfs.New()` based on options
2. **Layer Creation**: `createSquashFSLayer()` function handles SquashFS-specific layer creation
3. **File Output**: SquashFS layers are written directly to disk as `.squashfs` files

## Contributing

When contributing to SquashFS support:

1. **Testing**: Ensure all tests pass and add tests for new functionality
2. **Compatibility**: Maintain backward compatibility with existing tar.gz format
3. **Documentation**: Update this document for any significant changes
4. **Performance**: Consider memory usage and build time impacts

## Security Considerations

- **File Permissions**: All file permissions and ownership are preserved in SquashFS layers
- **Extended Attributes**: Security-related extended attributes are maintained
- **Read-Only**: SquashFS layers are read-only by design, providing immutability
- **Compression**: Built-in compression reduces storage requirements

## Related Documentation

- [OCI Image Format Specification](https://github.com/opencontainers/image-spec)
- [SquashFS Format Documentation](https://www.kernel.org/doc/html/latest/filesystems/squashfs.html)
- [go-diskfs Library](https://github.com/diskfs/go-diskfs)
- [apko Build System](https://edu.chainguard.dev/open-source/apko/overview/)