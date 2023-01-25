// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: pb/blog/blog.proto

// 定义包名。注意这和上面最后的go文件所属包名不一样，package是proto文件的包名。

package blogGo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求消息
type GetUserAllBlogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *GetUserAllBlogsRequest) Reset() {
	*x = GetUserAllBlogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_blog_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAllBlogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAllBlogsRequest) ProtoMessage() {}

func (x *GetUserAllBlogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_blog_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAllBlogsRequest.ProtoReflect.Descriptor instead.
func (*GetUserAllBlogsRequest) Descriptor() ([]byte, []int) {
	return file_pb_blog_blog_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserAllBlogsRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

// 响应消息
type GetUserAllBlogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog []*Blog `protobuf:"bytes,1,rep,name=blog,proto3" json:"blog,omitempty"`
}

func (x *GetUserAllBlogsResponse) Reset() {
	*x = GetUserAllBlogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_blog_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAllBlogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAllBlogsResponse) ProtoMessage() {}

func (x *GetUserAllBlogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_blog_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAllBlogsResponse.ProtoReflect.Descriptor instead.
func (*GetUserAllBlogsResponse) Descriptor() ([]byte, []int) {
	return file_pb_blog_blog_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserAllBlogsResponse) GetBlog() []*Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author  string `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_blog_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_pb_blog_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_pb_blog_blog_proto_rawDescGZIP(), []int{2}
}

func (x *Blog) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_pb_blog_blog_proto protoreflect.FileDescriptor

var file_pb_blog_blog_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x6c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x35, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x67,
	0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x4e, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x69, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x21, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x42,
	0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x69, 0x63, 0x6b, 0x79, 0x2d, 0x7a, 0x68, 0x66, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x65, 0x62,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x3b,
	0x62, 0x6c, 0x6f, 0x67, 0x47, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_blog_blog_proto_rawDescOnce sync.Once
	file_pb_blog_blog_proto_rawDescData = file_pb_blog_blog_proto_rawDesc
)

func file_pb_blog_blog_proto_rawDescGZIP() []byte {
	file_pb_blog_blog_proto_rawDescOnce.Do(func() {
		file_pb_blog_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_blog_blog_proto_rawDescData)
	})
	return file_pb_blog_blog_proto_rawDescData
}

var file_pb_blog_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_blog_blog_proto_goTypes = []interface{}{
	(*GetUserAllBlogsRequest)(nil),  // 0: blogProto.GetUserAllBlogsRequest
	(*GetUserAllBlogsResponse)(nil), // 1: blogProto.GetUserAllBlogsResponse
	(*Blog)(nil),                    // 2: blogProto.Blog
}
var file_pb_blog_blog_proto_depIdxs = []int32{
	2, // 0: blogProto.GetUserAllBlogsResponse.blog:type_name -> blogProto.Blog
	0, // 1: blogProto.BlogService.GetUserAllBlogs:input_type -> blogProto.GetUserAllBlogsRequest
	1, // 2: blogProto.BlogService.GetUserAllBlogs:output_type -> blogProto.GetUserAllBlogsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_blog_blog_proto_init() }
func file_pb_blog_blog_proto_init() {
	if File_pb_blog_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_blog_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAllBlogsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_blog_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAllBlogsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_blog_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_blog_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_blog_blog_proto_goTypes,
		DependencyIndexes: file_pb_blog_blog_proto_depIdxs,
		MessageInfos:      file_pb_blog_blog_proto_msgTypes,
	}.Build()
	File_pb_blog_blog_proto = out.File
	file_pb_blog_blog_proto_rawDesc = nil
	file_pb_blog_blog_proto_goTypes = nil
	file_pb_blog_blog_proto_depIdxs = nil
}
