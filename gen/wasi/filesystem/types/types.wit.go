// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package types represents the imported interface "wasi:filesystem/types@0.2.0".
package types

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	wallclock "go.wasmcloud.dev/component/gen/wasi/clocks/wall-clock"
	ioerror "go.wasmcloud.dev/component/gen/wasi/io/error"
	"go.wasmcloud.dev/component/gen/wasi/io/streams"
)

// FileSize represents the u64 "wasi:filesystem/types@0.2.0#filesize".
//
//	type filesize = u64
type FileSize uint64

// DescriptorType represents the enum "wasi:filesystem/types@0.2.0#descriptor-type".
//
//	enum descriptor-type {
//		unknown,
//		block-device,
//		character-device,
//		directory,
//		fifo,
//		symbolic-link,
//		regular-file,
//		socket
//	}
type DescriptorType uint8

const (
	DescriptorTypeUnknown DescriptorType = iota
	DescriptorTypeBlockDevice
	DescriptorTypeCharacterDevice
	DescriptorTypeDirectory
	DescriptorTypeFIFO
	DescriptorTypeSymbolicLink
	DescriptorTypeRegularFile
	DescriptorTypeSocket
)

var stringsDescriptorType = [8]string{
	"unknown",
	"block-device",
	"character-device",
	"directory",
	"fifo",
	"symbolic-link",
	"regular-file",
	"socket",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e DescriptorType) String() string {
	return stringsDescriptorType[e]
}

// DescriptorFlags represents the flags "wasi:filesystem/types@0.2.0#descriptor-flags".
//
//	flags descriptor-flags {
//		read,
//		write,
//		file-integrity-sync,
//		data-integrity-sync,
//		requested-write-sync,
//		mutate-directory,
//	}
type DescriptorFlags uint8

const (
	DescriptorFlagsRead DescriptorFlags = 1 << iota
	DescriptorFlagsWrite
	DescriptorFlagsFileIntegritySync
	DescriptorFlagsDataIntegritySync
	DescriptorFlagsRequestedWriteSync
	DescriptorFlagsMutateDirectory
)

// PathFlags represents the flags "wasi:filesystem/types@0.2.0#path-flags".
//
//	flags path-flags {
//		symlink-follow,
//	}
type PathFlags uint8

const (
	PathFlagsSymlinkFollow PathFlags = 1 << iota
)

// OpenFlags represents the flags "wasi:filesystem/types@0.2.0#open-flags".
//
//	flags open-flags {
//		create,
//		directory,
//		exclusive,
//		truncate,
//	}
type OpenFlags uint8

const (
	OpenFlagsCreate OpenFlags = 1 << iota
	OpenFlagsDirectory
	OpenFlagsExclusive
	OpenFlagsTruncate
)

// LinkCount represents the u64 "wasi:filesystem/types@0.2.0#link-count".
//
//	type link-count = u64
type LinkCount uint64

// DescriptorStat represents the record "wasi:filesystem/types@0.2.0#descriptor-stat".
//
//	record descriptor-stat {
//		%type: descriptor-type,
//		link-count: link-count,
//		size: filesize,
//		data-access-timestamp: option<datetime>,
//		data-modification-timestamp: option<datetime>,
//		status-change-timestamp: option<datetime>,
//	}
type DescriptorStat struct {
	_                         cm.HostLayout
	Type                      DescriptorType
	LinkCount                 LinkCount
	Size                      FileSize
	DataAccessTimestamp       cm.Option[wallclock.DateTime]
	DataModificationTimestamp cm.Option[wallclock.DateTime]
	StatusChangeTimestamp     cm.Option[wallclock.DateTime]
}

// NewTimestamp represents the variant "wasi:filesystem/types@0.2.0#new-timestamp".
//
//	variant new-timestamp {
//		no-change,
//		now,
//		timestamp(datetime),
//	}
type NewTimestamp cm.Variant[uint8, wallclock.DateTime, wallclock.DateTime]

// NewTimestampNoChange returns a [NewTimestamp] of case "no-change".
func NewTimestampNoChange() NewTimestamp {
	var data struct{}
	return cm.New[NewTimestamp](0, data)
}

// NoChange returns true if [NewTimestamp] represents the variant case "no-change".
func (self *NewTimestamp) NoChange() bool {
	return self.Tag() == 0
}

// NewTimestampNow returns a [NewTimestamp] of case "now".
func NewTimestampNow() NewTimestamp {
	var data struct{}
	return cm.New[NewTimestamp](1, data)
}

// Now returns true if [NewTimestamp] represents the variant case "now".
func (self *NewTimestamp) Now() bool {
	return self.Tag() == 1
}

// NewTimestampTimestamp returns a [NewTimestamp] of case "timestamp".
func NewTimestampTimestamp(data wallclock.DateTime) NewTimestamp {
	return cm.New[NewTimestamp](2, data)
}

// Timestamp returns a non-nil *[wallclock.DateTime] if [NewTimestamp] represents the variant case "timestamp".
func (self *NewTimestamp) Timestamp() *wallclock.DateTime {
	return cm.Case[wallclock.DateTime](self, 2)
}

// DirectoryEntry represents the record "wasi:filesystem/types@0.2.0#directory-entry".
//
//	record directory-entry {
//		%type: descriptor-type,
//		name: string,
//	}
type DirectoryEntry struct {
	_    cm.HostLayout
	Type DescriptorType
	Name string
}

// ErrorCode represents the enum "wasi:filesystem/types@0.2.0#error-code".
//
//	enum error-code {
//		access,
//		would-block,
//		already,
//		bad-descriptor,
//		busy,
//		deadlock,
//		quota,
//		exist,
//		file-too-large,
//		illegal-byte-sequence,
//		in-progress,
//		interrupted,
//		invalid,
//		io,
//		is-directory,
//		loop,
//		too-many-links,
//		message-size,
//		name-too-long,
//		no-device,
//		no-entry,
//		no-lock,
//		insufficient-memory,
//		insufficient-space,
//		not-directory,
//		not-empty,
//		not-recoverable,
//		unsupported,
//		no-tty,
//		no-such-device,
//		overflow,
//		not-permitted,
//		pipe,
//		read-only,
//		invalid-seek,
//		text-file-busy,
//		cross-device
//	}
type ErrorCode uint8

const (
	ErrorCodeAccess ErrorCode = iota
	ErrorCodeWouldBlock
	ErrorCodeAlready
	ErrorCodeBadDescriptor
	ErrorCodeBusy
	ErrorCodeDeadlock
	ErrorCodeQuota
	ErrorCodeExist
	ErrorCodeFileTooLarge
	ErrorCodeIllegalByteSequence
	ErrorCodeInProgress
	ErrorCodeInterrupted
	ErrorCodeInvalid
	ErrorCodeIO
	ErrorCodeIsDirectory
	ErrorCodeLoop
	ErrorCodeTooManyLinks
	ErrorCodeMessageSize
	ErrorCodeNameTooLong
	ErrorCodeNoDevice
	ErrorCodeNoEntry
	ErrorCodeNoLock
	ErrorCodeInsufficientMemory
	ErrorCodeInsufficientSpace
	ErrorCodeNotDirectory
	ErrorCodeNotEmpty
	ErrorCodeNotRecoverable
	ErrorCodeUnsupported
	ErrorCodeNoTTY
	ErrorCodeNoSuchDevice
	ErrorCodeOverflow
	ErrorCodeNotPermitted
	ErrorCodePipe
	ErrorCodeReadOnly
	ErrorCodeInvalidSeek
	ErrorCodeTextFileBusy
	ErrorCodeCrossDevice
)

var stringsErrorCode = [37]string{
	"access",
	"would-block",
	"already",
	"bad-descriptor",
	"busy",
	"deadlock",
	"quota",
	"exist",
	"file-too-large",
	"illegal-byte-sequence",
	"in-progress",
	"interrupted",
	"invalid",
	"io",
	"is-directory",
	"loop",
	"too-many-links",
	"message-size",
	"name-too-long",
	"no-device",
	"no-entry",
	"no-lock",
	"insufficient-memory",
	"insufficient-space",
	"not-directory",
	"not-empty",
	"not-recoverable",
	"unsupported",
	"no-tty",
	"no-such-device",
	"overflow",
	"not-permitted",
	"pipe",
	"read-only",
	"invalid-seek",
	"text-file-busy",
	"cross-device",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e ErrorCode) String() string {
	return stringsErrorCode[e]
}

// Advice represents the enum "wasi:filesystem/types@0.2.0#advice".
//
//	enum advice {
//		normal,
//		sequential,
//		random,
//		will-need,
//		dont-need,
//		no-reuse
//	}
type Advice uint8

const (
	AdviceNormal Advice = iota
	AdviceSequential
	AdviceRandom
	AdviceWillNeed
	AdviceDontNeed
	AdviceNoReuse
)

var stringsAdvice = [6]string{
	"normal",
	"sequential",
	"random",
	"will-need",
	"dont-need",
	"no-reuse",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e Advice) String() string {
	return stringsAdvice[e]
}

// MetadataHashValue represents the record "wasi:filesystem/types@0.2.0#metadata-hash-value".
//
//	record metadata-hash-value {
//		lower: u64,
//		upper: u64,
//	}
type MetadataHashValue struct {
	_     cm.HostLayout
	Lower uint64
	Upper uint64
}

// Descriptor represents the imported resource "wasi:filesystem/types@0.2.0#descriptor".
//
//	resource descriptor
type Descriptor cm.Resource

// ResourceDrop represents the imported resource-drop for resource "descriptor".
//
// Drops a resource handle.
//
//go:nosplit
func (self Descriptor) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DescriptorResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [resource-drop]descriptor
//go:noescape
func wasmimport_DescriptorResourceDrop(self0 uint32)

// Advise represents the imported method "advise".
//
//	advise: func(offset: filesize, length: filesize, advice: advice) -> result<_, error-code>
//
//go:nosplit
func (self Descriptor) Advise(offset FileSize, length FileSize, advice Advice) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	offset0 := (uint64)(offset)
	length0 := (uint64)(length)
	advice0 := (uint32)(advice)
	wasmimport_DescriptorAdvise((uint32)(self0), (uint64)(offset0), (uint64)(length0), (uint32)(advice0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.advise
//go:noescape
func wasmimport_DescriptorAdvise(self0 uint32, offset0 uint64, length0 uint64, advice0 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

// AppendViaStream represents the imported method "append-via-stream".
//
//	append-via-stream: func() -> result<output-stream, error-code>
//
//go:nosplit
func (self Descriptor) AppendViaStream() (result cm.Result[streams.OutputStream, streams.OutputStream, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DescriptorAppendViaStream((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.append-via-stream
//go:noescape
func wasmimport_DescriptorAppendViaStream(self0 uint32, result *cm.Result[streams.OutputStream, streams.OutputStream, ErrorCode])

// CreateDirectoryAt represents the imported method "create-directory-at".
//
//	create-directory-at: func(path: string) -> result<_, error-code>
//
//go:nosplit
func (self Descriptor) CreateDirectoryAt(path string) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	path0, path1 := cm.LowerString(path)
	wasmimport_DescriptorCreateDirectoryAt((uint32)(self0), (*uint8)(path0), (uint32)(path1), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.create-directory-at
//go:noescape
func wasmimport_DescriptorCreateDirectoryAt(self0 uint32, path0 *uint8, path1 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

// GetFlags represents the imported method "get-flags".
//
//	get-flags: func() -> result<descriptor-flags, error-code>
//
//go:nosplit
func (self Descriptor) GetFlags() (result cm.Result[DescriptorFlags, DescriptorFlags, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DescriptorGetFlags((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.get-flags
//go:noescape
func wasmimport_DescriptorGetFlags(self0 uint32, result *cm.Result[DescriptorFlags, DescriptorFlags, ErrorCode])

// GetType represents the imported method "get-type".
//
//	get-type: func() -> result<descriptor-type, error-code>
//
//go:nosplit
func (self Descriptor) GetType() (result cm.Result[DescriptorType, DescriptorType, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DescriptorGetType((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.get-type
//go:noescape
func wasmimport_DescriptorGetType(self0 uint32, result *cm.Result[DescriptorType, DescriptorType, ErrorCode])

// IsSameObject represents the imported method "is-same-object".
//
//	is-same-object: func(other: borrow<descriptor>) -> bool
//
//go:nosplit
func (self Descriptor) IsSameObject(other Descriptor) (result bool) {
	self0 := cm.Reinterpret[uint32](self)
	other0 := cm.Reinterpret[uint32](other)
	result0 := wasmimport_DescriptorIsSameObject((uint32)(self0), (uint32)(other0))
	result = cm.U32ToBool((uint32)(result0))
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.is-same-object
//go:noescape
func wasmimport_DescriptorIsSameObject(self0 uint32, other0 uint32) (result0 uint32)

// LinkAt represents the imported method "link-at".
//
//	link-at: func(old-path-flags: path-flags, old-path: string, new-descriptor: borrow<descriptor>,
//	new-path: string) -> result<_, error-code>
//
//go:nosplit
func (self Descriptor) LinkAt(oldPathFlags PathFlags, oldPath string, newDescriptor Descriptor, newPath string) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	oldPathFlags0 := (uint32)(oldPathFlags)
	oldPath0, oldPath1 := cm.LowerString(oldPath)
	newDescriptor0 := cm.Reinterpret[uint32](newDescriptor)
	newPath0, newPath1 := cm.LowerString(newPath)
	wasmimport_DescriptorLinkAt((uint32)(self0), (uint32)(oldPathFlags0), (*uint8)(oldPath0), (uint32)(oldPath1), (uint32)(newDescriptor0), (*uint8)(newPath0), (uint32)(newPath1), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.link-at
//go:noescape
func wasmimport_DescriptorLinkAt(self0 uint32, oldPathFlags0 uint32, oldPath0 *uint8, oldPath1 uint32, newDescriptor0 uint32, newPath0 *uint8, newPath1 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

// MetadataHash represents the imported method "metadata-hash".
//
//	metadata-hash: func() -> result<metadata-hash-value, error-code>
//
//go:nosplit
func (self Descriptor) MetadataHash() (result cm.Result[MetadataHashValueShape, MetadataHashValue, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DescriptorMetadataHash((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.metadata-hash
//go:noescape
func wasmimport_DescriptorMetadataHash(self0 uint32, result *cm.Result[MetadataHashValueShape, MetadataHashValue, ErrorCode])

// MetadataHashAt represents the imported method "metadata-hash-at".
//
//	metadata-hash-at: func(path-flags: path-flags, path: string) -> result<metadata-hash-value,
//	error-code>
//
//go:nosplit
func (self Descriptor) MetadataHashAt(pathFlags PathFlags, path string) (result cm.Result[MetadataHashValueShape, MetadataHashValue, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	pathFlags0 := (uint32)(pathFlags)
	path0, path1 := cm.LowerString(path)
	wasmimport_DescriptorMetadataHashAt((uint32)(self0), (uint32)(pathFlags0), (*uint8)(path0), (uint32)(path1), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.metadata-hash-at
//go:noescape
func wasmimport_DescriptorMetadataHashAt(self0 uint32, pathFlags0 uint32, path0 *uint8, path1 uint32, result *cm.Result[MetadataHashValueShape, MetadataHashValue, ErrorCode])

// OpenAt represents the imported method "open-at".
//
//	open-at: func(path-flags: path-flags, path: string, open-flags: open-flags, %flags:
//	descriptor-flags) -> result<descriptor, error-code>
//
//go:nosplit
func (self Descriptor) OpenAt(pathFlags PathFlags, path string, openFlags OpenFlags, flags DescriptorFlags) (result cm.Result[Descriptor, Descriptor, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	pathFlags0 := (uint32)(pathFlags)
	path0, path1 := cm.LowerString(path)
	openFlags0 := (uint32)(openFlags)
	flags0 := (uint32)(flags)
	wasmimport_DescriptorOpenAt((uint32)(self0), (uint32)(pathFlags0), (*uint8)(path0), (uint32)(path1), (uint32)(openFlags0), (uint32)(flags0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.open-at
//go:noescape
func wasmimport_DescriptorOpenAt(self0 uint32, pathFlags0 uint32, path0 *uint8, path1 uint32, openFlags0 uint32, flags0 uint32, result *cm.Result[Descriptor, Descriptor, ErrorCode])

// Read represents the imported method "read".
//
//	read: func(length: filesize, offset: filesize) -> result<tuple<list<u8>, bool>,
//	error-code>
//
//go:nosplit
func (self Descriptor) Read(length FileSize, offset FileSize) (result cm.Result[TupleListU8BoolShape, cm.Tuple[cm.List[uint8], bool], ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	length0 := (uint64)(length)
	offset0 := (uint64)(offset)
	wasmimport_DescriptorRead((uint32)(self0), (uint64)(length0), (uint64)(offset0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.read
//go:noescape
func wasmimport_DescriptorRead(self0 uint32, length0 uint64, offset0 uint64, result *cm.Result[TupleListU8BoolShape, cm.Tuple[cm.List[uint8], bool], ErrorCode])

// ReadDirectory represents the imported method "read-directory".
//
//	read-directory: func() -> result<directory-entry-stream, error-code>
//
//go:nosplit
func (self Descriptor) ReadDirectory() (result cm.Result[DirectoryEntryStream, DirectoryEntryStream, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DescriptorReadDirectory((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.read-directory
//go:noescape
func wasmimport_DescriptorReadDirectory(self0 uint32, result *cm.Result[DirectoryEntryStream, DirectoryEntryStream, ErrorCode])

// ReadViaStream represents the imported method "read-via-stream".
//
//	read-via-stream: func(offset: filesize) -> result<input-stream, error-code>
//
//go:nosplit
func (self Descriptor) ReadViaStream(offset FileSize) (result cm.Result[streams.InputStream, streams.InputStream, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	offset0 := (uint64)(offset)
	wasmimport_DescriptorReadViaStream((uint32)(self0), (uint64)(offset0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.read-via-stream
//go:noescape
func wasmimport_DescriptorReadViaStream(self0 uint32, offset0 uint64, result *cm.Result[streams.InputStream, streams.InputStream, ErrorCode])

// ReadLinkAt represents the imported method "readlink-at".
//
//	readlink-at: func(path: string) -> result<string, error-code>
//
//go:nosplit
func (self Descriptor) ReadLinkAt(path string) (result cm.Result[string, string, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	path0, path1 := cm.LowerString(path)
	wasmimport_DescriptorReadLinkAt((uint32)(self0), (*uint8)(path0), (uint32)(path1), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.readlink-at
//go:noescape
func wasmimport_DescriptorReadLinkAt(self0 uint32, path0 *uint8, path1 uint32, result *cm.Result[string, string, ErrorCode])

// RemoveDirectoryAt represents the imported method "remove-directory-at".
//
//	remove-directory-at: func(path: string) -> result<_, error-code>
//
//go:nosplit
func (self Descriptor) RemoveDirectoryAt(path string) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	path0, path1 := cm.LowerString(path)
	wasmimport_DescriptorRemoveDirectoryAt((uint32)(self0), (*uint8)(path0), (uint32)(path1), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.remove-directory-at
//go:noescape
func wasmimport_DescriptorRemoveDirectoryAt(self0 uint32, path0 *uint8, path1 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

// RenameAt represents the imported method "rename-at".
//
//	rename-at: func(old-path: string, new-descriptor: borrow<descriptor>, new-path:
//	string) -> result<_, error-code>
//
//go:nosplit
func (self Descriptor) RenameAt(oldPath string, newDescriptor Descriptor, newPath string) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	oldPath0, oldPath1 := cm.LowerString(oldPath)
	newDescriptor0 := cm.Reinterpret[uint32](newDescriptor)
	newPath0, newPath1 := cm.LowerString(newPath)
	wasmimport_DescriptorRenameAt((uint32)(self0), (*uint8)(oldPath0), (uint32)(oldPath1), (uint32)(newDescriptor0), (*uint8)(newPath0), (uint32)(newPath1), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.rename-at
//go:noescape
func wasmimport_DescriptorRenameAt(self0 uint32, oldPath0 *uint8, oldPath1 uint32, newDescriptor0 uint32, newPath0 *uint8, newPath1 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

// SetSize represents the imported method "set-size".
//
//	set-size: func(size: filesize) -> result<_, error-code>
//
//go:nosplit
func (self Descriptor) SetSize(size FileSize) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	size0 := (uint64)(size)
	wasmimport_DescriptorSetSize((uint32)(self0), (uint64)(size0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.set-size
//go:noescape
func wasmimport_DescriptorSetSize(self0 uint32, size0 uint64, result *cm.Result[ErrorCode, struct{}, ErrorCode])

// SetTimes represents the imported method "set-times".
//
//	set-times: func(data-access-timestamp: new-timestamp, data-modification-timestamp:
//	new-timestamp) -> result<_, error-code>
//
//go:nosplit
func (self Descriptor) SetTimes(dataAccessTimestamp NewTimestamp, dataModificationTimestamp NewTimestamp) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	dataAccessTimestamp0, dataAccessTimestamp1, dataAccessTimestamp2 := lower_NewTimestamp(dataAccessTimestamp)
	dataModificationTimestamp0, dataModificationTimestamp1, dataModificationTimestamp2 := lower_NewTimestamp(dataModificationTimestamp)
	wasmimport_DescriptorSetTimes((uint32)(self0), (uint32)(dataAccessTimestamp0), (uint64)(dataAccessTimestamp1), (uint32)(dataAccessTimestamp2), (uint32)(dataModificationTimestamp0), (uint64)(dataModificationTimestamp1), (uint32)(dataModificationTimestamp2), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.set-times
//go:noescape
func wasmimport_DescriptorSetTimes(self0 uint32, dataAccessTimestamp0 uint32, dataAccessTimestamp1 uint64, dataAccessTimestamp2 uint32, dataModificationTimestamp0 uint32, dataModificationTimestamp1 uint64, dataModificationTimestamp2 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

// SetTimesAt represents the imported method "set-times-at".
//
//	set-times-at: func(path-flags: path-flags, path: string, data-access-timestamp:
//	new-timestamp, data-modification-timestamp: new-timestamp) -> result<_, error-code>
//
//go:nosplit
func (self Descriptor) SetTimesAt(pathFlags PathFlags, path string, dataAccessTimestamp NewTimestamp, dataModificationTimestamp NewTimestamp) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	pathFlags0 := (uint32)(pathFlags)
	path0, path1 := cm.LowerString(path)
	dataAccessTimestamp0, dataAccessTimestamp1, dataAccessTimestamp2 := lower_NewTimestamp(dataAccessTimestamp)
	dataModificationTimestamp0, dataModificationTimestamp1, dataModificationTimestamp2 := lower_NewTimestamp(dataModificationTimestamp)
	wasmimport_DescriptorSetTimesAt((uint32)(self0), (uint32)(pathFlags0), (*uint8)(path0), (uint32)(path1), (uint32)(dataAccessTimestamp0), (uint64)(dataAccessTimestamp1), (uint32)(dataAccessTimestamp2), (uint32)(dataModificationTimestamp0), (uint64)(dataModificationTimestamp1), (uint32)(dataModificationTimestamp2), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.set-times-at
//go:noescape
func wasmimport_DescriptorSetTimesAt(self0 uint32, pathFlags0 uint32, path0 *uint8, path1 uint32, dataAccessTimestamp0 uint32, dataAccessTimestamp1 uint64, dataAccessTimestamp2 uint32, dataModificationTimestamp0 uint32, dataModificationTimestamp1 uint64, dataModificationTimestamp2 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

// Stat represents the imported method "stat".
//
//	stat: func() -> result<descriptor-stat, error-code>
//
//go:nosplit
func (self Descriptor) Stat() (result cm.Result[DescriptorStatShape, DescriptorStat, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DescriptorStat((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.stat
//go:noescape
func wasmimport_DescriptorStat(self0 uint32, result *cm.Result[DescriptorStatShape, DescriptorStat, ErrorCode])

// StatAt represents the imported method "stat-at".
//
//	stat-at: func(path-flags: path-flags, path: string) -> result<descriptor-stat,
//	error-code>
//
//go:nosplit
func (self Descriptor) StatAt(pathFlags PathFlags, path string) (result cm.Result[DescriptorStatShape, DescriptorStat, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	pathFlags0 := (uint32)(pathFlags)
	path0, path1 := cm.LowerString(path)
	wasmimport_DescriptorStatAt((uint32)(self0), (uint32)(pathFlags0), (*uint8)(path0), (uint32)(path1), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.stat-at
//go:noescape
func wasmimport_DescriptorStatAt(self0 uint32, pathFlags0 uint32, path0 *uint8, path1 uint32, result *cm.Result[DescriptorStatShape, DescriptorStat, ErrorCode])

// SymlinkAt represents the imported method "symlink-at".
//
//	symlink-at: func(old-path: string, new-path: string) -> result<_, error-code>
//
//go:nosplit
func (self Descriptor) SymlinkAt(oldPath string, newPath string) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	oldPath0, oldPath1 := cm.LowerString(oldPath)
	newPath0, newPath1 := cm.LowerString(newPath)
	wasmimport_DescriptorSymlinkAt((uint32)(self0), (*uint8)(oldPath0), (uint32)(oldPath1), (*uint8)(newPath0), (uint32)(newPath1), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.symlink-at
//go:noescape
func wasmimport_DescriptorSymlinkAt(self0 uint32, oldPath0 *uint8, oldPath1 uint32, newPath0 *uint8, newPath1 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

// Sync represents the imported method "sync".
//
//	sync: func() -> result<_, error-code>
//
//go:nosplit
func (self Descriptor) Sync() (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DescriptorSync((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.sync
//go:noescape
func wasmimport_DescriptorSync(self0 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

// SyncData represents the imported method "sync-data".
//
//	sync-data: func() -> result<_, error-code>
//
//go:nosplit
func (self Descriptor) SyncData() (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DescriptorSyncData((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.sync-data
//go:noescape
func wasmimport_DescriptorSyncData(self0 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

// UnlinkFileAt represents the imported method "unlink-file-at".
//
//	unlink-file-at: func(path: string) -> result<_, error-code>
//
//go:nosplit
func (self Descriptor) UnlinkFileAt(path string) (result cm.Result[ErrorCode, struct{}, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	path0, path1 := cm.LowerString(path)
	wasmimport_DescriptorUnlinkFileAt((uint32)(self0), (*uint8)(path0), (uint32)(path1), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.unlink-file-at
//go:noescape
func wasmimport_DescriptorUnlinkFileAt(self0 uint32, path0 *uint8, path1 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

// Write represents the imported method "write".
//
//	write: func(buffer: list<u8>, offset: filesize) -> result<filesize, error-code>
//
//go:nosplit
func (self Descriptor) Write(buffer cm.List[uint8], offset FileSize) (result cm.Result[uint64, FileSize, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	buffer0, buffer1 := cm.LowerList(buffer)
	offset0 := (uint64)(offset)
	wasmimport_DescriptorWrite((uint32)(self0), (*uint8)(buffer0), (uint32)(buffer1), (uint64)(offset0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.write
//go:noescape
func wasmimport_DescriptorWrite(self0 uint32, buffer0 *uint8, buffer1 uint32, offset0 uint64, result *cm.Result[uint64, FileSize, ErrorCode])

// WriteViaStream represents the imported method "write-via-stream".
//
//	write-via-stream: func(offset: filesize) -> result<output-stream, error-code>
//
//go:nosplit
func (self Descriptor) WriteViaStream(offset FileSize) (result cm.Result[streams.OutputStream, streams.OutputStream, ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	offset0 := (uint64)(offset)
	wasmimport_DescriptorWriteViaStream((uint32)(self0), (uint64)(offset0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]descriptor.write-via-stream
//go:noescape
func wasmimport_DescriptorWriteViaStream(self0 uint32, offset0 uint64, result *cm.Result[streams.OutputStream, streams.OutputStream, ErrorCode])

// DirectoryEntryStream represents the imported resource "wasi:filesystem/types@0.2.0#directory-entry-stream".
//
//	resource directory-entry-stream
type DirectoryEntryStream cm.Resource

// ResourceDrop represents the imported resource-drop for resource "directory-entry-stream".
//
// Drops a resource handle.
//
//go:nosplit
func (self DirectoryEntryStream) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DirectoryEntryStreamResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [resource-drop]directory-entry-stream
//go:noescape
func wasmimport_DirectoryEntryStreamResourceDrop(self0 uint32)

// ReadDirectoryEntry represents the imported method "read-directory-entry".
//
//	read-directory-entry: func() -> result<option<directory-entry>, error-code>
//
//go:nosplit
func (self DirectoryEntryStream) ReadDirectoryEntry() (result cm.Result[OptionDirectoryEntryShape, cm.Option[DirectoryEntry], ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_DirectoryEntryStreamReadDirectoryEntry((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 [method]directory-entry-stream.read-directory-entry
//go:noescape
func wasmimport_DirectoryEntryStreamReadDirectoryEntry(self0 uint32, result *cm.Result[OptionDirectoryEntryShape, cm.Option[DirectoryEntry], ErrorCode])

// FilesystemErrorCode represents the imported function "filesystem-error-code".
//
//	filesystem-error-code: func(err: borrow<error>) -> option<error-code>
//
//go:nosplit
func FilesystemErrorCode(err ioerror.Error) (result cm.Option[ErrorCode]) {
	err0 := cm.Reinterpret[uint32](err)
	wasmimport_FilesystemErrorCode((uint32)(err0), &result)
	return
}

//go:wasmimport wasi:filesystem/types@0.2.0 filesystem-error-code
//go:noescape
func wasmimport_FilesystemErrorCode(err0 uint32, result *cm.Option[ErrorCode])
