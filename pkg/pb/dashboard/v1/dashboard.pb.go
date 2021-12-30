// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: dashboard/v1/dashboard.proto

package dashboard

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListPullRequestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPullRequestsRequest) Reset() {
	*x = ListPullRequestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_v1_dashboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPullRequestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPullRequestsRequest) ProtoMessage() {}

func (x *ListPullRequestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_v1_dashboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPullRequestsRequest.ProtoReflect.Descriptor instead.
func (*ListPullRequestsRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_v1_dashboard_proto_rawDescGZIP(), []int{0}
}

type StreamPullRequestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamPullRequestsRequest) Reset() {
	*x = StreamPullRequestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_v1_dashboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamPullRequestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamPullRequestsRequest) ProtoMessage() {}

func (x *StreamPullRequestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_v1_dashboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamPullRequestsRequest.ProtoReflect.Descriptor instead.
func (*StreamPullRequestsRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_v1_dashboard_proto_rawDescGZIP(), []int{1}
}

type ListPullRequestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Repository `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListPullRequestsResponse) Reset() {
	*x = ListPullRequestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_v1_dashboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPullRequestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPullRequestsResponse) ProtoMessage() {}

func (x *ListPullRequestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_v1_dashboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPullRequestsResponse.ProtoReflect.Descriptor instead.
func (*ListPullRequestsResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_v1_dashboard_proto_rawDescGZIP(), []int{2}
}

func (x *ListPullRequestsResponse) GetItems() []*Repository {
	if x != nil {
		return x.Items
	}
	return nil
}

type Repository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pullrequests []*PullRequest `protobuf:"bytes,2,rep,name=pullrequests,proto3" json:"pullrequests,omitempty"`
}

func (x *Repository) Reset() {
	*x = Repository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_v1_dashboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repository) ProtoMessage() {}

func (x *Repository) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_v1_dashboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repository.ProtoReflect.Descriptor instead.
func (*Repository) Descriptor() ([]byte, []int) {
	return file_dashboard_v1_dashboard_proto_rawDescGZIP(), []int{3}
}

func (x *Repository) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Repository) GetPullrequests() []*PullRequest {
	if x != nil {
		return x.Pullrequests
	}
	return nil
}

type PullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url   string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	User  string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *PullRequest) Reset() {
	*x = PullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_v1_dashboard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullRequest) ProtoMessage() {}

func (x *PullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_v1_dashboard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullRequest.ProtoReflect.Descriptor instead.
func (*PullRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_v1_dashboard_proto_rawDescGZIP(), []int{4}
}

func (x *PullRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PullRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PullRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

var File_dashboard_v1_dashboard_proto protoreflect.FileDescriptor

var file_dashboard_v1_dashboard_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x19, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1b, 0x0a,
	0x19, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x18, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x5f, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x75, 0x6c, 0x6c,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x70, 0x75, 0x6c, 0x6c, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x49, 0x0a, 0x0b, 0x50, 0x75, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x32, 0xd5, 0x01, 0x0a, 0x10, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x25, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x12,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x27, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x75, 0x6d, 0x68, 0x61, 0x72,
	0x64, 0x2f, 0x70, 0x72, 0x2d, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dashboard_v1_dashboard_proto_rawDescOnce sync.Once
	file_dashboard_v1_dashboard_proto_rawDescData = file_dashboard_v1_dashboard_proto_rawDesc
)

func file_dashboard_v1_dashboard_proto_rawDescGZIP() []byte {
	file_dashboard_v1_dashboard_proto_rawDescOnce.Do(func() {
		file_dashboard_v1_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_dashboard_v1_dashboard_proto_rawDescData)
	})
	return file_dashboard_v1_dashboard_proto_rawDescData
}

var file_dashboard_v1_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dashboard_v1_dashboard_proto_goTypes = []interface{}{
	(*ListPullRequestsRequest)(nil),   // 0: dashboard.v1.ListPullRequestsRequest
	(*StreamPullRequestsRequest)(nil), // 1: dashboard.v1.StreamPullRequestsRequest
	(*ListPullRequestsResponse)(nil),  // 2: dashboard.v1.ListPullRequestsResponse
	(*Repository)(nil),                // 3: dashboard.v1.Repository
	(*PullRequest)(nil),               // 4: dashboard.v1.PullRequest
}
var file_dashboard_v1_dashboard_proto_depIdxs = []int32{
	3, // 0: dashboard.v1.ListPullRequestsResponse.items:type_name -> dashboard.v1.Repository
	4, // 1: dashboard.v1.Repository.pullrequests:type_name -> dashboard.v1.PullRequest
	0, // 2: dashboard.v1.DashboardService.ListPullRequests:input_type -> dashboard.v1.ListPullRequestsRequest
	1, // 3: dashboard.v1.DashboardService.StreamPullRequests:input_type -> dashboard.v1.StreamPullRequestsRequest
	2, // 4: dashboard.v1.DashboardService.ListPullRequests:output_type -> dashboard.v1.ListPullRequestsResponse
	4, // 5: dashboard.v1.DashboardService.StreamPullRequests:output_type -> dashboard.v1.PullRequest
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dashboard_v1_dashboard_proto_init() }
func file_dashboard_v1_dashboard_proto_init() {
	if File_dashboard_v1_dashboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dashboard_v1_dashboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPullRequestsRequest); i {
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
		file_dashboard_v1_dashboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamPullRequestsRequest); i {
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
		file_dashboard_v1_dashboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPullRequestsResponse); i {
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
		file_dashboard_v1_dashboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repository); i {
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
		file_dashboard_v1_dashboard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullRequest); i {
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
			RawDescriptor: file_dashboard_v1_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dashboard_v1_dashboard_proto_goTypes,
		DependencyIndexes: file_dashboard_v1_dashboard_proto_depIdxs,
		MessageInfos:      file_dashboard_v1_dashboard_proto_msgTypes,
	}.Build()
	File_dashboard_v1_dashboard_proto = out.File
	file_dashboard_v1_dashboard_proto_rawDesc = nil
	file_dashboard_v1_dashboard_proto_goTypes = nil
	file_dashboard_v1_dashboard_proto_depIdxs = nil
}
