// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/recommend.proto

package pb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RecommendationCreateUserCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Comment   string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Labels    []string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	Subscribe []string `protobuf:"bytes,4,rep,name=subscribe,proto3" json:"subscribe,omitempty"`
}

func (x *RecommendationCreateUserCommand) Reset() {
	*x = RecommendationCreateUserCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_recommend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendationCreateUserCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationCreateUserCommand) ProtoMessage() {}

func (x *RecommendationCreateUserCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_recommend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationCreateUserCommand.ProtoReflect.Descriptor instead.
func (*RecommendationCreateUserCommand) Descriptor() ([]byte, []int) {
	return file_proto_recommend_proto_rawDescGZIP(), []int{0}
}

func (x *RecommendationCreateUserCommand) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RecommendationCreateUserCommand) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *RecommendationCreateUserCommand) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *RecommendationCreateUserCommand) GetSubscribe() []string {
	if x != nil {
		return x.Subscribe
	}
	return nil
}

type RecommendationSuggestUserProjectCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RecommendationSuggestUserProjectCommand) Reset() {
	*x = RecommendationSuggestUserProjectCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_recommend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendationSuggestUserProjectCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationSuggestUserProjectCommand) ProtoMessage() {}

func (x *RecommendationSuggestUserProjectCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_recommend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationSuggestUserProjectCommand.ProtoReflect.Descriptor instead.
func (*RecommendationSuggestUserProjectCommand) Descriptor() ([]byte, []int) {
	return file_proto_recommend_proto_rawDescGZIP(), []int{1}
}

func (x *RecommendationSuggestUserProjectCommand) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RecommendationCreateProjectCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RecommendationCreateProjectCommand) Reset() {
	*x = RecommendationCreateProjectCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_recommend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendationCreateProjectCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationCreateProjectCommand) ProtoMessage() {}

func (x *RecommendationCreateProjectCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_recommend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationCreateProjectCommand.ProtoReflect.Descriptor instead.
func (*RecommendationCreateProjectCommand) Descriptor() ([]byte, []int) {
	return file_proto_recommend_proto_rawDescGZIP(), []int{2}
}

func (x *RecommendationCreateProjectCommand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Feedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeedbackType string                 `protobuf:"bytes,1,opt,name=feedback_type,json=feedbackType,proto3" json:"feedback_type,omitempty"`
	EmployeeId   string                 `protobuf:"bytes,2,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	ItemId       string                 `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Timestamp    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Feedback) Reset() {
	*x = Feedback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_recommend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feedback) ProtoMessage() {}

func (x *Feedback) ProtoReflect() protoreflect.Message {
	mi := &file_proto_recommend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feedback.ProtoReflect.Descriptor instead.
func (*Feedback) Descriptor() ([]byte, []int) {
	return file_proto_recommend_proto_rawDescGZIP(), []int{3}
}

func (x *Feedback) GetFeedbackType() string {
	if x != nil {
		return x.FeedbackType
	}
	return ""
}

func (x *Feedback) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

func (x *Feedback) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *Feedback) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type RecommendationInsertFeedbackCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feedback *Feedback `protobuf:"bytes,1,opt,name=feedback,proto3" json:"feedback,omitempty"`
}

func (x *RecommendationInsertFeedbackCommand) Reset() {
	*x = RecommendationInsertFeedbackCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_recommend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendationInsertFeedbackCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationInsertFeedbackCommand) ProtoMessage() {}

func (x *RecommendationInsertFeedbackCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_recommend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationInsertFeedbackCommand.ProtoReflect.Descriptor instead.
func (*RecommendationInsertFeedbackCommand) Descriptor() ([]byte, []int) {
	return file_proto_recommend_proto_rawDescGZIP(), []int{4}
}

func (x *RecommendationInsertFeedbackCommand) GetFeedback() *Feedback {
	if x != nil {
		return x.Feedback
	}
	return nil
}

var File_proto_recommend_proto protoreflect.FileDescriptor

var file_proto_recommend_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x74, 0x61, 0x66, 0x66, 0x69, 0x6e,
	0x67, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x1f, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0x42, 0x0a, 0x27, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x22, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x67, 0x0a, 0x23, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x67, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x08, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x32, 0xba, 0x03, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x61, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x67,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x71, 0x0a, 0x12, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x43, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x67, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3e,
	0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x69, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x3f, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_recommend_proto_rawDescOnce sync.Once
	file_proto_recommend_proto_rawDescData = file_proto_recommend_proto_rawDesc
)

func file_proto_recommend_proto_rawDescGZIP() []byte {
	file_proto_recommend_proto_rawDescOnce.Do(func() {
		file_proto_recommend_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_recommend_proto_rawDescData)
	})
	return file_proto_recommend_proto_rawDescData
}

var file_proto_recommend_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_recommend_proto_goTypes = []interface{}{
	(*RecommendationCreateUserCommand)(nil),         // 0: staffing.recommendation.v1.RecommendationCreateUserCommand
	(*RecommendationSuggestUserProjectCommand)(nil), // 1: staffing.recommendation.v1.RecommendationSuggestUserProjectCommand
	(*RecommendationCreateProjectCommand)(nil),      // 2: staffing.recommendation.v1.RecommendationCreateProjectCommand
	(*Feedback)(nil), // 3: staffing.recommendation.v1.Feedback
	(*RecommendationInsertFeedbackCommand)(nil), // 4: staffing.recommendation.v1.RecommendationInsertFeedbackCommand
	(*timestamppb.Timestamp)(nil),               // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                       // 6: google.protobuf.Empty
}
var file_proto_recommend_proto_depIdxs = []int32{
	5, // 0: staffing.recommendation.v1.Feedback.timestamp:type_name -> google.protobuf.Timestamp
	3, // 1: staffing.recommendation.v1.RecommendationInsertFeedbackCommand.feedback:type_name -> staffing.recommendation.v1.Feedback
	0, // 2: staffing.recommendation.v1.Recommendation.CreateUser:input_type -> staffing.recommendation.v1.RecommendationCreateUserCommand
	1, // 3: staffing.recommendation.v1.Recommendation.SuggestUserProject:input_type -> staffing.recommendation.v1.RecommendationSuggestUserProjectCommand
	2, // 4: staffing.recommendation.v1.Recommendation.CreateProject:input_type -> staffing.recommendation.v1.RecommendationCreateProjectCommand
	4, // 5: staffing.recommendation.v1.Recommendation.InsertFeedback:input_type -> staffing.recommendation.v1.RecommendationInsertFeedbackCommand
	6, // 6: staffing.recommendation.v1.Recommendation.CreateUser:output_type -> google.protobuf.Empty
	6, // 7: staffing.recommendation.v1.Recommendation.SuggestUserProject:output_type -> google.protobuf.Empty
	6, // 8: staffing.recommendation.v1.Recommendation.CreateProject:output_type -> google.protobuf.Empty
	6, // 9: staffing.recommendation.v1.Recommendation.InsertFeedback:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_recommend_proto_init() }
func file_proto_recommend_proto_init() {
	if File_proto_recommend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_recommend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendationCreateUserCommand); i {
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
		file_proto_recommend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendationSuggestUserProjectCommand); i {
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
		file_proto_recommend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendationCreateProjectCommand); i {
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
		file_proto_recommend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feedback); i {
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
		file_proto_recommend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendationInsertFeedbackCommand); i {
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
			RawDescriptor: file_proto_recommend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_recommend_proto_goTypes,
		DependencyIndexes: file_proto_recommend_proto_depIdxs,
		MessageInfos:      file_proto_recommend_proto_msgTypes,
	}.Build()
	File_proto_recommend_proto = out.File
	file_proto_recommend_proto_rawDesc = nil
	file_proto_recommend_proto_goTypes = nil
	file_proto_recommend_proto_depIdxs = nil
}
