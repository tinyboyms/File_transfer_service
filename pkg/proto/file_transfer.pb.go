// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: pkg/proto/file_transfer.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileChunk struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Content       []byte                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	TotalSize     int64                  `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	ChunkSize     int64                  `protobuf:"varint,4,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	Offset        int64                  `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	mi := &file_pkg_proto_file_transfer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_file_transfer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_pkg_proto_file_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *FileChunk) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileChunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *FileChunk) GetTotalSize() int64 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *FileChunk) GetChunkSize() int64 {
	if x != nil {
		return x.ChunkSize
	}
	return 0
}

func (x *FileChunk) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	mi := &file_pkg_proto_file_transfer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_file_transfer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_file_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *FileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type FileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Size          int64                  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	mi := &file_pkg_proto_file_transfer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_file_transfer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_file_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *FileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FileResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *FileResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FileInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Size          int64                  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Exists        bool                   `protobuf:"varint,3,opt,name=exists,proto3" json:"exists,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	mi := &file_pkg_proto_file_transfer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_file_transfer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_pkg_proto_file_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *FileInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileInfo) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_pkg_proto_file_transfer_proto protoreflect.FileDescriptor

const file_pkg_proto_file_transfer_proto_rawDesc = "" +
	"\n" +
	"\x1dpkg/proto/file_transfer.proto\x12\ffiletransfer\"\x97\x01\n" +
	"\tFileChunk\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\x12\x18\n" +
	"\acontent\x18\x02 \x01(\fR\acontent\x12\x1d\n" +
	"\n" +
	"total_size\x18\x03 \x01(\x03R\ttotalSize\x12\x1d\n" +
	"\n" +
	"chunk_size\x18\x04 \x01(\x03R\tchunkSize\x12\x16\n" +
	"\x06offset\x18\x05 \x01(\x03R\x06offset\")\n" +
	"\vFileRequest\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\"V\n" +
	"\fFileResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12\x18\n" +
	"\asuccess\x18\x02 \x01(\bR\asuccess\x12\x12\n" +
	"\x04size\x18\x03 \x01(\x03R\x04size\"R\n" +
	"\bFileInfo\x12\x1a\n" +
	"\bfilename\x18\x01 \x01(\tR\bfilename\x12\x12\n" +
	"\x04size\x18\x02 \x01(\x03R\x04size\x12\x16\n" +
	"\x06exists\x18\x03 \x01(\bR\x06exists2\xe2\x01\n" +
	"\x13FileTransferService\x12C\n" +
	"\n" +
	"UploadFile\x12\x17.filetransfer.FileChunk\x1a\x1a.filetransfer.FileResponse(\x01\x12D\n" +
	"\fDownloadFile\x12\x19.filetransfer.FileRequest\x1a\x17.filetransfer.FileChunk0\x01\x12@\n" +
	"\vGetFileInfo\x12\x19.filetransfer.FileRequest\x1a\x16.filetransfer.FileInfoB1Z/github.com/yourusername/file_transfer/pkg/protob\x06proto3"

var (
	file_pkg_proto_file_transfer_proto_rawDescOnce sync.Once
	file_pkg_proto_file_transfer_proto_rawDescData []byte
)

func file_pkg_proto_file_transfer_proto_rawDescGZIP() []byte {
	file_pkg_proto_file_transfer_proto_rawDescOnce.Do(func() {
		file_pkg_proto_file_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_proto_file_transfer_proto_rawDesc), len(file_pkg_proto_file_transfer_proto_rawDesc)))
	})
	return file_pkg_proto_file_transfer_proto_rawDescData
}

var file_pkg_proto_file_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_proto_file_transfer_proto_goTypes = []any{
	(*FileChunk)(nil),    // 0: filetransfer.FileChunk
	(*FileRequest)(nil),  // 1: filetransfer.FileRequest
	(*FileResponse)(nil), // 2: filetransfer.FileResponse
	(*FileInfo)(nil),     // 3: filetransfer.FileInfo
}
var file_pkg_proto_file_transfer_proto_depIdxs = []int32{
	0, // 0: filetransfer.FileTransferService.UploadFile:input_type -> filetransfer.FileChunk
	1, // 1: filetransfer.FileTransferService.DownloadFile:input_type -> filetransfer.FileRequest
	1, // 2: filetransfer.FileTransferService.GetFileInfo:input_type -> filetransfer.FileRequest
	2, // 3: filetransfer.FileTransferService.UploadFile:output_type -> filetransfer.FileResponse
	0, // 4: filetransfer.FileTransferService.DownloadFile:output_type -> filetransfer.FileChunk
	3, // 5: filetransfer.FileTransferService.GetFileInfo:output_type -> filetransfer.FileInfo
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_file_transfer_proto_init() }
func file_pkg_proto_file_transfer_proto_init() {
	if File_pkg_proto_file_transfer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_proto_file_transfer_proto_rawDesc), len(file_pkg_proto_file_transfer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_file_transfer_proto_goTypes,
		DependencyIndexes: file_pkg_proto_file_transfer_proto_depIdxs,
		MessageInfos:      file_pkg_proto_file_transfer_proto_msgTypes,
	}.Build()
	File_pkg_proto_file_transfer_proto = out.File
	file_pkg_proto_file_transfer_proto_goTypes = nil
	file_pkg_proto_file_transfer_proto_depIdxs = nil
}
