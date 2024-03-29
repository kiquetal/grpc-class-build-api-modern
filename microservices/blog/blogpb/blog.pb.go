// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: blog/blogpb/blog.proto

package blogpb

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

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_blogpb_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_blog_blogpb_blog_proto_msgTypes[0]
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
	return file_blog_blogpb_blog_proto_rawDescGZIP(), []int{0}
}

func (x *Blog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Blog) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
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

type CreateBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *CreateBlogRequest) Reset() {
	*x = CreateBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_blogpb_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogRequest) ProtoMessage() {}

func (x *CreateBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_blogpb_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogRequest.ProtoReflect.Descriptor instead.
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return file_blog_blogpb_blog_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBlogRequest) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type CreateBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"` // will have blog id
}

func (x *CreateBlogResponse) Reset() {
	*x = CreateBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_blogpb_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogResponse) ProtoMessage() {}

func (x *CreateBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blog_blogpb_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogResponse.ProtoReflect.Descriptor instead.
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return file_blog_blogpb_blog_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type ReadBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId string `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
}

func (x *ReadBlogRequest) Reset() {
	*x = ReadBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_blogpb_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlogRequest) ProtoMessage() {}

func (x *ReadBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_blogpb_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlogRequest.ProtoReflect.Descriptor instead.
func (*ReadBlogRequest) Descriptor() ([]byte, []int) {
	return file_blog_blogpb_blog_proto_rawDescGZIP(), []int{3}
}

func (x *ReadBlogRequest) GetBlogId() string {
	if x != nil {
		return x.BlogId
	}
	return ""
}

type UpdateBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *UpdateBlogRequest) Reset() {
	*x = UpdateBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_blogpb_blog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogRequest) ProtoMessage() {}

func (x *UpdateBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_blogpb_blog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogRequest.ProtoReflect.Descriptor instead.
func (*UpdateBlogRequest) Descriptor() ([]byte, []int) {
	return file_blog_blogpb_blog_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBlogRequest) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type UpdateBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *UpdateBlogResponse) Reset() {
	*x = UpdateBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_blogpb_blog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogResponse) ProtoMessage() {}

func (x *UpdateBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blog_blogpb_blog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogResponse.ProtoReflect.Descriptor instead.
func (*UpdateBlogResponse) Descriptor() ([]byte, []int) {
	return file_blog_blogpb_blog_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type DeleteBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId string `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
}

func (x *DeleteBlogResponse) Reset() {
	*x = DeleteBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_blogpb_blog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogResponse) ProtoMessage() {}

func (x *DeleteBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blog_blogpb_blog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogResponse.ProtoReflect.Descriptor instead.
func (*DeleteBlogResponse) Descriptor() ([]byte, []int) {
	return file_blog_blogpb_blog_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBlogResponse) GetBlogId() string {
	if x != nil {
		return x.BlogId
	}
	return ""
}

type ListBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBlogRequest) Reset() {
	*x = ListBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_blogpb_blog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogRequest) ProtoMessage() {}

func (x *ListBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_blogpb_blog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogRequest.ProtoReflect.Descriptor instead.
func (*ListBlogRequest) Descriptor() ([]byte, []int) {
	return file_blog_blogpb_blog_proto_rawDescGZIP(), []int{7}
}

var File_blog_blogpb_blog_proto protoreflect.FileDescriptor

var file_blog_blogpb_blog_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2f, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62,
	0x22, 0x63, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x62, 0x6c,
	0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70,
	0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x36, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04,
	0x62, 0x6c, 0x6f, 0x67, 0x22, 0x2a, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64,
	0x22, 0x35, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f,
	0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x36, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22,
	0x2d, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0x11,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x32, 0xc2, 0x02, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x43, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12,
	0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c,
	0x6f, 0x67, 0x12, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x43, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x17, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x17, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x2e,
	0x42, 0x6c, 0x6f, 0x67, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blog_blogpb_blog_proto_rawDescOnce sync.Once
	file_blog_blogpb_blog_proto_rawDescData = file_blog_blogpb_blog_proto_rawDesc
)

func file_blog_blogpb_blog_proto_rawDescGZIP() []byte {
	file_blog_blogpb_blog_proto_rawDescOnce.Do(func() {
		file_blog_blogpb_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_blogpb_blog_proto_rawDescData)
	})
	return file_blog_blogpb_blog_proto_rawDescData
}

var file_blog_blogpb_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_blog_blogpb_blog_proto_goTypes = []interface{}{
	(*Blog)(nil),               // 0: blogpb.Blog
	(*CreateBlogRequest)(nil),  // 1: blogpb.CreateBlogRequest
	(*CreateBlogResponse)(nil), // 2: blogpb.CreateBlogResponse
	(*ReadBlogRequest)(nil),    // 3: blogpb.ReadBlogRequest
	(*UpdateBlogRequest)(nil),  // 4: blogpb.UpdateBlogRequest
	(*UpdateBlogResponse)(nil), // 5: blogpb.UpdateBlogResponse
	(*DeleteBlogResponse)(nil), // 6: blogpb.DeleteBlogResponse
	(*ListBlogRequest)(nil),    // 7: blogpb.ListBlogRequest
}
var file_blog_blogpb_blog_proto_depIdxs = []int32{
	0, // 0: blogpb.CreateBlogRequest.blog:type_name -> blogpb.Blog
	0, // 1: blogpb.CreateBlogResponse.blog:type_name -> blogpb.Blog
	0, // 2: blogpb.UpdateBlogRequest.blog:type_name -> blogpb.Blog
	0, // 3: blogpb.UpdateBlogResponse.blog:type_name -> blogpb.Blog
	1, // 4: blogpb.BlogService.CreateBlog:input_type -> blogpb.CreateBlogRequest
	3, // 5: blogpb.BlogService.ReadBlog:input_type -> blogpb.ReadBlogRequest
	4, // 6: blogpb.BlogService.UpdateBlog:input_type -> blogpb.UpdateBlogRequest
	3, // 7: blogpb.BlogService.DeleteBlog:input_type -> blogpb.ReadBlogRequest
	7, // 8: blogpb.BlogService.ListBlog:input_type -> blogpb.ListBlogRequest
	2, // 9: blogpb.BlogService.CreateBlog:output_type -> blogpb.CreateBlogResponse
	0, // 10: blogpb.BlogService.ReadBlog:output_type -> blogpb.Blog
	5, // 11: blogpb.BlogService.UpdateBlog:output_type -> blogpb.UpdateBlogResponse
	6, // 12: blogpb.BlogService.DeleteBlog:output_type -> blogpb.DeleteBlogResponse
	0, // 13: blogpb.BlogService.ListBlog:output_type -> blogpb.Blog
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_blog_blogpb_blog_proto_init() }
func file_blog_blogpb_blog_proto_init() {
	if File_blog_blogpb_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blog_blogpb_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_blog_blogpb_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogRequest); i {
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
		file_blog_blogpb_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogResponse); i {
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
		file_blog_blogpb_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlogRequest); i {
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
		file_blog_blogpb_blog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlogRequest); i {
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
		file_blog_blogpb_blog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlogResponse); i {
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
		file_blog_blogpb_blog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlogResponse); i {
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
		file_blog_blogpb_blog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlogRequest); i {
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
			RawDescriptor: file_blog_blogpb_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_blogpb_blog_proto_goTypes,
		DependencyIndexes: file_blog_blogpb_blog_proto_depIdxs,
		MessageInfos:      file_blog_blogpb_blog_proto_msgTypes,
	}.Build()
	File_blog_blogpb_blog_proto = out.File
	file_blog_blogpb_blog_proto_rawDesc = nil
	file_blog_blogpb_blog_proto_goTypes = nil
	file_blog_blogpb_blog_proto_depIdxs = nil
}
