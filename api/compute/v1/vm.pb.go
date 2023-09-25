// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: compute/v1/vm.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateVmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image   string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Port    string   `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Command []string `protobuf:"bytes,3,rep,name=command,proto3" json:"command,omitempty"`
}

func (x *CreateVmRequest) Reset() {
	*x = CreateVmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_v1_vm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVmRequest) ProtoMessage() {}

func (x *CreateVmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_compute_v1_vm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVmRequest.ProtoReflect.Descriptor instead.
func (*CreateVmRequest) Descriptor() ([]byte, []int) {
	return file_compute_v1_vm_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVmRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateVmRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *CreateVmRequest) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

type PortBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip          string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	PrivatePort uint32 `protobuf:"varint,2,opt,name=private_port,json=privatePort,proto3" json:"private_port,omitempty"`
	PublicPort  uint32 `protobuf:"varint,3,opt,name=public_port,json=publicPort,proto3" json:"public_port,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PortBinding) Reset() {
	*x = PortBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_v1_vm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortBinding) ProtoMessage() {}

func (x *PortBinding) ProtoReflect() protoreflect.Message {
	mi := &file_compute_v1_vm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortBinding.ProtoReflect.Descriptor instead.
func (*PortBinding) Descriptor() ([]byte, []int) {
	return file_compute_v1_vm_proto_rawDescGZIP(), []int{1}
}

func (x *PortBinding) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *PortBinding) GetPrivatePort() uint32 {
	if x != nil {
		return x.PrivatePort
	}
	return 0
}

func (x *PortBinding) GetPublicPort() uint32 {
	if x != nil {
		return x.PublicPort
	}
	return 0
}

func (x *PortBinding) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type DeleteVmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteVmRequest) Reset() {
	*x = DeleteVmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_v1_vm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVmRequest) ProtoMessage() {}

func (x *DeleteVmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_compute_v1_vm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVmRequest.ProtoReflect.Descriptor instead.
func (*DeleteVmRequest) Descriptor() ([]byte, []int) {
	return file_compute_v1_vm_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteVmRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteVmReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteVmReply) Reset() {
	*x = DeleteVmReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_v1_vm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVmReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVmReply) ProtoMessage() {}

func (x *DeleteVmReply) ProtoReflect() protoreflect.Message {
	mi := &file_compute_v1_vm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVmReply.ProtoReflect.Descriptor instead.
func (*DeleteVmReply) Descriptor() ([]byte, []int) {
	return file_compute_v1_vm_proto_rawDescGZIP(), []int{3}
}

type GetVmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetVmRequest) Reset() {
	*x = GetVmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_v1_vm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVmRequest) ProtoMessage() {}

func (x *GetVmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_compute_v1_vm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVmRequest.ProtoReflect.Descriptor instead.
func (*GetVmRequest) Descriptor() ([]byte, []int) {
	return file_compute_v1_vm_proto_rawDescGZIP(), []int{4}
}

func (x *GetVmRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetVmReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Image       string         `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Ports       []*PortBinding `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	CpuUsage    uint64         `protobuf:"varint,4,opt,name=cpuUsage,proto3" json:"cpuUsage,omitempty"`
	MemoryUsage uint64         `protobuf:"varint,5,opt,name=memoryUsage,proto3" json:"memoryUsage,omitempty"`
}

func (x *GetVmReply) Reset() {
	*x = GetVmReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_v1_vm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVmReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVmReply) ProtoMessage() {}

func (x *GetVmReply) ProtoReflect() protoreflect.Message {
	mi := &file_compute_v1_vm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVmReply.ProtoReflect.Descriptor instead.
func (*GetVmReply) Descriptor() ([]byte, []int) {
	return file_compute_v1_vm_proto_rawDescGZIP(), []int{5}
}

func (x *GetVmReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetVmReply) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *GetVmReply) GetPorts() []*PortBinding {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *GetVmReply) GetCpuUsage() uint64 {
	if x != nil {
		return x.CpuUsage
	}
	return 0
}

func (x *GetVmReply) GetMemoryUsage() uint64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

type ListVmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListVmRequest) Reset() {
	*x = ListVmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_v1_vm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVmRequest) ProtoMessage() {}

func (x *ListVmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_compute_v1_vm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVmRequest.ProtoReflect.Descriptor instead.
func (*ListVmRequest) Descriptor() ([]byte, []int) {
	return file_compute_v1_vm_proto_rawDescGZIP(), []int{6}
}

type ListVmReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*GetVmReply `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ListVmReply) Reset() {
	*x = ListVmReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compute_v1_vm_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVmReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVmReply) ProtoMessage() {}

func (x *ListVmReply) ProtoReflect() protoreflect.Message {
	mi := &file_compute_v1_vm_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVmReply.ProtoReflect.Descriptor instead.
func (*ListVmReply) Descriptor() ([]byte, []int) {
	return file_compute_v1_vm_proto_rawDescGZIP(), []int{7}
}

func (x *ListVmReply) GetResult() []*GetVmReply {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_compute_v1_vm_proto protoreflect.FileDescriptor

var file_compute_v1_vm_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x75, 0x0a, 0x0b, 0x50, 0x6f,
	0x72, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x21, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6d,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x6d, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x56, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x56, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0xb3, 0x04, 0x0a, 0x02, 0x56, 0x6d, 0x12, 0x5a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x6d, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x3a, 0x01, 0x2a, 0x22, 0x06, 0x2f, 0x76, 0x31, 0x2f,
	0x76, 0x6d, 0x12, 0x5f, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6d, 0x12, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6d, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x56, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x56, 0x6d, 0x12, 0x1c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x56, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56,
	0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b,
	0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6d, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x54, 0x0a, 0x06, 0x4c,
	0x69, 0x73, 0x74, 0x56, 0x6d, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6d, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x76, 0x31, 0x2f, 0x76,
	0x6d, 0x12, 0x61, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x56, 0x6d, 0x12, 0x1c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x56, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56,
	0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01,
	0x2a, 0x1a, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6d, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x5f, 0x0a, 0x06, 0x53, 0x74, 0x6f, 0x70, 0x56, 0x6d, 0x12, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x56, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x56, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x3a, 0x01, 0x2a, 0x1a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6d, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x2f, 0x73, 0x74, 0x6f, 0x70, 0x42, 0x4f, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x68, 0x61, 0x69, 0x6a, 0x69, 0x61, 0x6e, 0x67,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_compute_v1_vm_proto_rawDescOnce sync.Once
	file_compute_v1_vm_proto_rawDescData = file_compute_v1_vm_proto_rawDesc
)

func file_compute_v1_vm_proto_rawDescGZIP() []byte {
	file_compute_v1_vm_proto_rawDescOnce.Do(func() {
		file_compute_v1_vm_proto_rawDescData = protoimpl.X.CompressGZIP(file_compute_v1_vm_proto_rawDescData)
	})
	return file_compute_v1_vm_proto_rawDescData
}

var file_compute_v1_vm_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_compute_v1_vm_proto_goTypes = []interface{}{
	(*CreateVmRequest)(nil), // 0: api.compute.v1.CreateVmRequest
	(*PortBinding)(nil),     // 1: api.compute.v1.PortBinding
	(*DeleteVmRequest)(nil), // 2: api.compute.v1.DeleteVmRequest
	(*DeleteVmReply)(nil),   // 3: api.compute.v1.DeleteVmReply
	(*GetVmRequest)(nil),    // 4: api.compute.v1.GetVmRequest
	(*GetVmReply)(nil),      // 5: api.compute.v1.GetVmReply
	(*ListVmRequest)(nil),   // 6: api.compute.v1.ListVmRequest
	(*ListVmReply)(nil),     // 7: api.compute.v1.ListVmReply
}
var file_compute_v1_vm_proto_depIdxs = []int32{
	1, // 0: api.compute.v1.GetVmReply.ports:type_name -> api.compute.v1.PortBinding
	5, // 1: api.compute.v1.ListVmReply.result:type_name -> api.compute.v1.GetVmReply
	0, // 2: api.compute.v1.Vm.CreateVm:input_type -> api.compute.v1.CreateVmRequest
	2, // 3: api.compute.v1.Vm.DeleteVm:input_type -> api.compute.v1.DeleteVmRequest
	4, // 4: api.compute.v1.Vm.GetVm:input_type -> api.compute.v1.GetVmRequest
	6, // 5: api.compute.v1.Vm.ListVm:input_type -> api.compute.v1.ListVmRequest
	4, // 6: api.compute.v1.Vm.StartVm:input_type -> api.compute.v1.GetVmRequest
	4, // 7: api.compute.v1.Vm.StopVm:input_type -> api.compute.v1.GetVmRequest
	5, // 8: api.compute.v1.Vm.CreateVm:output_type -> api.compute.v1.GetVmReply
	3, // 9: api.compute.v1.Vm.DeleteVm:output_type -> api.compute.v1.DeleteVmReply
	5, // 10: api.compute.v1.Vm.GetVm:output_type -> api.compute.v1.GetVmReply
	7, // 11: api.compute.v1.Vm.ListVm:output_type -> api.compute.v1.ListVmReply
	5, // 12: api.compute.v1.Vm.StartVm:output_type -> api.compute.v1.GetVmReply
	5, // 13: api.compute.v1.Vm.StopVm:output_type -> api.compute.v1.GetVmReply
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_compute_v1_vm_proto_init() }
func file_compute_v1_vm_proto_init() {
	if File_compute_v1_vm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_compute_v1_vm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVmRequest); i {
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
		file_compute_v1_vm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortBinding); i {
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
		file_compute_v1_vm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVmRequest); i {
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
		file_compute_v1_vm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVmReply); i {
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
		file_compute_v1_vm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVmRequest); i {
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
		file_compute_v1_vm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVmReply); i {
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
		file_compute_v1_vm_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVmRequest); i {
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
		file_compute_v1_vm_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVmReply); i {
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
			RawDescriptor: file_compute_v1_vm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_compute_v1_vm_proto_goTypes,
		DependencyIndexes: file_compute_v1_vm_proto_depIdxs,
		MessageInfos:      file_compute_v1_vm_proto_msgTypes,
	}.Build()
	File_compute_v1_vm_proto = out.File
	file_compute_v1_vm_proto_rawDesc = nil
	file_compute_v1_vm_proto_goTypes = nil
	file_compute_v1_vm_proto_depIdxs = nil
}
