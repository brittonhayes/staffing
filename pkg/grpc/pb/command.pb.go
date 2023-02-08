// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/command.proto

package pb

import (
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

type MoveEmployeeToBenchCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId   string `protobuf:"bytes,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	EmployeeName string `protobuf:"bytes,2,opt,name=employee_name,json=employeeName,proto3" json:"employee_name,omitempty"`
}

func (x *MoveEmployeeToBenchCommand) Reset() {
	*x = MoveEmployeeToBenchCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveEmployeeToBenchCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveEmployeeToBenchCommand) ProtoMessage() {}

func (x *MoveEmployeeToBenchCommand) ProtoReflect() protoreflect.Message {
	mi := &file_api_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveEmployeeToBenchCommand.ProtoReflect.Descriptor instead.
func (*MoveEmployeeToBenchCommand) Descriptor() ([]byte, []int) {
	return file_api_command_proto_rawDescGZIP(), []int{0}
}

func (x *MoveEmployeeToBenchCommand) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

func (x *MoveEmployeeToBenchCommand) GetEmployeeName() string {
	if x != nil {
		return x.EmployeeName
	}
	return ""
}

type AlignEmployeeToProjectCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId   string                 `protobuf:"bytes,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	EmployeeName string                 `protobuf:"bytes,2,opt,name=employee_name,json=employeeName,proto3" json:"employee_name,omitempty"`
	ProjectId    string                 `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ProjectName  string                 `protobuf:"bytes,4,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	StartDate    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *AlignEmployeeToProjectCommand) Reset() {
	*x = AlignEmployeeToProjectCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlignEmployeeToProjectCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlignEmployeeToProjectCommand) ProtoMessage() {}

func (x *AlignEmployeeToProjectCommand) ProtoReflect() protoreflect.Message {
	mi := &file_api_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlignEmployeeToProjectCommand.ProtoReflect.Descriptor instead.
func (*AlignEmployeeToProjectCommand) Descriptor() ([]byte, []int) {
	return file_api_command_proto_rawDescGZIP(), []int{1}
}

func (x *AlignEmployeeToProjectCommand) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

func (x *AlignEmployeeToProjectCommand) GetEmployeeName() string {
	if x != nil {
		return x.EmployeeName
	}
	return ""
}

func (x *AlignEmployeeToProjectCommand) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *AlignEmployeeToProjectCommand) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *AlignEmployeeToProjectCommand) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *AlignEmployeeToProjectCommand) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

type EmployeeMovedToBenchEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId   string `protobuf:"bytes,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	EmployeeName string `protobuf:"bytes,2,opt,name=employee_name,json=employeeName,proto3" json:"employee_name,omitempty"`
}

func (x *EmployeeMovedToBenchEvent) Reset() {
	*x = EmployeeMovedToBenchEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeMovedToBenchEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeMovedToBenchEvent) ProtoMessage() {}

func (x *EmployeeMovedToBenchEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeMovedToBenchEvent.ProtoReflect.Descriptor instead.
func (*EmployeeMovedToBenchEvent) Descriptor() ([]byte, []int) {
	return file_api_command_proto_rawDescGZIP(), []int{2}
}

func (x *EmployeeMovedToBenchEvent) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

func (x *EmployeeMovedToBenchEvent) GetEmployeeName() string {
	if x != nil {
		return x.EmployeeName
	}
	return ""
}

type EmployeeAlignedToProjectEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId string                 `protobuf:"bytes,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	ProjectId  string                 `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	StartDate  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *EmployeeAlignedToProjectEvent) Reset() {
	*x = EmployeeAlignedToProjectEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeAlignedToProjectEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeAlignedToProjectEvent) ProtoMessage() {}

func (x *EmployeeAlignedToProjectEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeAlignedToProjectEvent.ProtoReflect.Descriptor instead.
func (*EmployeeAlignedToProjectEvent) Descriptor() ([]byte, []int) {
	return file_api_command_proto_rawDescGZIP(), []int{3}
}

func (x *EmployeeAlignedToProjectEvent) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

func (x *EmployeeAlignedToProjectEvent) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *EmployeeAlignedToProjectEvent) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *EmployeeAlignedToProjectEvent) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

var File_api_command_proto protoreflect.FileDescriptor

var file_api_command_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x62, 0x0a, 0x1a, 0x4d, 0x6f, 0x76, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x54, 0x6f, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x99, 0x02, 0x0a, 0x1d, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x22, 0x61, 0x0a, 0x19, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x4d, 0x6f, 0x76,
	0x65, 0x64, 0x54, 0x6f, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd1, 0x01, 0x0a, 0x1d, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x32, 0xa8, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x4a, 0x0a, 0x13, 0x4d, 0x6f, 0x76, 0x65, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x6f, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x4d,
	0x6f, 0x76, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x6f, 0x42, 0x65, 0x6e,
	0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x50, 0x0a, 0x16, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x2e, 0x41, 0x6c,
	0x69, 0x67, 0x6e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x6f, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_command_proto_rawDescOnce sync.Once
	file_api_command_proto_rawDescData = file_api_command_proto_rawDesc
)

func file_api_command_proto_rawDescGZIP() []byte {
	file_api_command_proto_rawDescOnce.Do(func() {
		file_api_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_command_proto_rawDescData)
	})
	return file_api_command_proto_rawDescData
}

var file_api_command_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_command_proto_goTypes = []interface{}{
	(*MoveEmployeeToBenchCommand)(nil),    // 0: MoveEmployeeToBenchCommand
	(*AlignEmployeeToProjectCommand)(nil), // 1: AlignEmployeeToProjectCommand
	(*EmployeeMovedToBenchEvent)(nil),     // 2: EmployeeMovedToBenchEvent
	(*EmployeeAlignedToProjectEvent)(nil), // 3: EmployeeAlignedToProjectEvent
	(*timestamppb.Timestamp)(nil),         // 4: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                 // 5: google.protobuf.Empty
}
var file_api_command_proto_depIdxs = []int32{
	4, // 0: AlignEmployeeToProjectCommand.start_date:type_name -> google.protobuf.Timestamp
	4, // 1: AlignEmployeeToProjectCommand.end_date:type_name -> google.protobuf.Timestamp
	4, // 2: EmployeeAlignedToProjectEvent.start_date:type_name -> google.protobuf.Timestamp
	4, // 3: EmployeeAlignedToProjectEvent.end_date:type_name -> google.protobuf.Timestamp
	0, // 4: Commands.MoveEmployeeToBench:input_type -> MoveEmployeeToBenchCommand
	1, // 5: Commands.AlignEmployeeToProject:input_type -> AlignEmployeeToProjectCommand
	5, // 6: Commands.MoveEmployeeToBench:output_type -> google.protobuf.Empty
	5, // 7: Commands.AlignEmployeeToProject:output_type -> google.protobuf.Empty
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_command_proto_init() }
func file_api_command_proto_init() {
	if File_api_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveEmployeeToBenchCommand); i {
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
		file_api_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlignEmployeeToProjectCommand); i {
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
		file_api_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeMovedToBenchEvent); i {
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
		file_api_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeAlignedToProjectEvent); i {
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
			RawDescriptor: file_api_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_command_proto_goTypes,
		DependencyIndexes: file_api_command_proto_depIdxs,
		MessageInfos:      file_api_command_proto_msgTypes,
	}.Build()
	File_api_command_proto = out.File
	file_api_command_proto_rawDesc = nil
	file_api_command_proto_goTypes = nil
	file_api_command_proto_depIdxs = nil
}
