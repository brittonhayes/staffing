// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/project.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProjectCreateCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProjectCreateCommand) Reset() {
	*x = ProjectCreateCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectCreateCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectCreateCommand) ProtoMessage() {}

func (x *ProjectCreateCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectCreateCommand.ProtoReflect.Descriptor instead.
func (*ProjectCreateCommand) Descriptor() ([]byte, []int) {
	return file_proto_project_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectCreateCommand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProjectAssignEmployeeCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId  string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	EmployeeId string `protobuf:"bytes,2,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
}

func (x *ProjectAssignEmployeeCommand) Reset() {
	*x = ProjectAssignEmployeeCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectAssignEmployeeCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAssignEmployeeCommand) ProtoMessage() {}

func (x *ProjectAssignEmployeeCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAssignEmployeeCommand.ProtoReflect.Descriptor instead.
func (*ProjectAssignEmployeeCommand) Descriptor() ([]byte, []int) {
	return file_proto_project_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectAssignEmployeeCommand) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ProjectAssignEmployeeCommand) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

type ProjectUnassignEmployeeCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId  string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	EmployeeId string `protobuf:"bytes,2,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
}

func (x *ProjectUnassignEmployeeCommand) Reset() {
	*x = ProjectUnassignEmployeeCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectUnassignEmployeeCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectUnassignEmployeeCommand) ProtoMessage() {}

func (x *ProjectUnassignEmployeeCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectUnassignEmployeeCommand.ProtoReflect.Descriptor instead.
func (*ProjectUnassignEmployeeCommand) Descriptor() ([]byte, []int) {
	return file_proto_project_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectUnassignEmployeeCommand) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ProjectUnassignEmployeeCommand) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

var File_proto_project_proto protoreflect.FileDescriptor

var file_proto_project_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x74, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x5e, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x1e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x6e,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x49, 0x64, 0x32, 0xa1, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x29, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x10, 0x55, 0x6e, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x33, 0x2e, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_project_proto_rawDescOnce sync.Once
	file_proto_project_proto_rawDescData = file_proto_project_proto_rawDesc
)

func file_proto_project_proto_rawDescGZIP() []byte {
	file_proto_project_proto_rawDescOnce.Do(func() {
		file_proto_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_project_proto_rawDescData)
	})
	return file_proto_project_proto_rawDescData
}

var file_proto_project_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_project_proto_goTypes = []interface{}{
	(*ProjectCreateCommand)(nil),           // 0: staffing.project.v1.ProjectCreateCommand
	(*ProjectAssignEmployeeCommand)(nil),   // 1: staffing.project.v1.ProjectAssignEmployeeCommand
	(*ProjectUnassignEmployeeCommand)(nil), // 2: staffing.project.v1.ProjectUnassignEmployeeCommand
	(*emptypb.Empty)(nil),                  // 3: google.protobuf.Empty
}
var file_proto_project_proto_depIdxs = []int32{
	0, // 0: staffing.project.v1.Project.CreateProject:input_type -> staffing.project.v1.ProjectCreateCommand
	1, // 1: staffing.project.v1.Project.AssignEmployee:input_type -> staffing.project.v1.ProjectAssignEmployeeCommand
	2, // 2: staffing.project.v1.Project.UnassignEmployee:input_type -> staffing.project.v1.ProjectUnassignEmployeeCommand
	3, // 3: staffing.project.v1.Project.CreateProject:output_type -> google.protobuf.Empty
	3, // 4: staffing.project.v1.Project.AssignEmployee:output_type -> google.protobuf.Empty
	3, // 5: staffing.project.v1.Project.UnassignEmployee:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_project_proto_init() }
func file_proto_project_proto_init() {
	if File_proto_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectCreateCommand); i {
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
		file_proto_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectAssignEmployeeCommand); i {
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
		file_proto_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectUnassignEmployeeCommand); i {
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
			RawDescriptor: file_proto_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_project_proto_goTypes,
		DependencyIndexes: file_proto_project_proto_depIdxs,
		MessageInfos:      file_proto_project_proto_msgTypes,
	}.Build()
	File_proto_project_proto = out.File
	file_proto_project_proto_rawDesc = nil
	file_proto_project_proto_goTypes = nil
	file_proto_project_proto_depIdxs = nil
}
