// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: internal/proto/grpc.proto

package grpc

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

type CustomerId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CustomerId) Reset() {
	*x = CustomerId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerId) ProtoMessage() {}

func (x *CustomerId) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerId.ProtoReflect.Descriptor instead.
func (*CustomerId) Descriptor() ([]byte, []int) {
	return file_internal_proto_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	Firstname  string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname   string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Address1   string `protobuf:"bytes,4,opt,name=address1,proto3" json:"address1,omitempty"`
	Address2   string `protobuf:"bytes,5,opt,name=address2,proto3" json:"address2,omitempty"`
	City       string `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	Postcode   string `protobuf:"bytes,7,opt,name=postcode,proto3" json:"postcode,omitempty"`
	CountryId  string `protobuf:"bytes,8,opt,name=countryId,proto3" json:"countryId,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_internal_proto_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Address) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Address) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Address) GetAddress1() string {
	if x != nil {
		return x.Address1
	}
	return ""
}

func (x *Address) GetAddress2() string {
	if x != nil {
		return x.Address2
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetPostcode() string {
	if x != nil {
		return x.Postcode
	}
	return ""
}

func (x *Address) GetCountryId() string {
	if x != nil {
		return x.CountryId
	}
	return ""
}

type AddressutsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AddressutsRequest) Reset() {
	*x = AddressutsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressutsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressutsRequest) ProtoMessage() {}

func (x *AddressutsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressutsRequest.ProtoReflect.Descriptor instead.
func (*AddressutsRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *AddressutsRequest) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type AddressutsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *AddressutsResponse) Reset() {
	*x = AddressutsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressutsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressutsResponse) ProtoMessage() {}

func (x *AddressutsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressutsResponse.ProtoReflect.Descriptor instead.
func (*AddressutsResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *AddressutsResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

type AddressctuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customerid *CustomerId `protobuf:"bytes,1,opt,name=customerid,proto3" json:"customerid,omitempty"`
}

func (x *AddressctuRequest) Reset() {
	*x = AddressctuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressctuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressctuRequest) ProtoMessage() {}

func (x *AddressctuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressctuRequest.ProtoReflect.Descriptor instead.
func (*AddressctuRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *AddressctuRequest) GetCustomerid() *CustomerId {
	if x != nil {
		return x.Customerid
	}
	return nil
}

type AddressctuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AddressctuResponse) Reset() {
	*x = AddressctuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressctuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressctuResponse) ProtoMessage() {}

func (x *AddressctuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressctuResponse.ProtoReflect.Descriptor instead.
func (*AddressctuResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *AddressctuResponse) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Credential) Reset() {
	*x = Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_internal_proto_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *Credential) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Credential) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credential *Credential `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (x *CredentialRequest) Reset() {
	*x = CredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialRequest) ProtoMessage() {}

func (x *CredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialRequest.ProtoReflect.Descriptor instead.
func (*CredentialRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *CredentialRequest) GetCredential() *Credential {
	if x != nil {
		return x.Credential
	}
	return nil
}

type CredentialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customerid string `protobuf:"bytes,1,opt,name=customerid,proto3" json:"customerid,omitempty"`
	Ispresent  bool   `protobuf:"varint,2,opt,name=ispresent,proto3" json:"ispresent,omitempty"`
}

func (x *CredentialResponse) Reset() {
	*x = CredentialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredentialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialResponse) ProtoMessage() {}

func (x *CredentialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialResponse.ProtoReflect.Descriptor instead.
func (*CredentialResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *CredentialResponse) GetCustomerid() string {
	if x != nil {
		return x.Customerid
	}
	return ""
}

func (x *CredentialResponse) GetIspresent() bool {
	if x != nil {
		return x.Ispresent
	}
	return false
}

var File_internal_proto_grpc_proto protoreflect.FileDescriptor

var file_internal_proto_grpc_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xe9, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x11,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x75, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x30, 0x0a, 0x12, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x75, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x0a,
	0x11, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x63, 0x74, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x63, 0x74, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x44, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x46, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x31, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x22, 0x52, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x32, 0xa1, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x75, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x75, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x75, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_grpc_proto_rawDescOnce sync.Once
	file_internal_proto_grpc_proto_rawDescData = file_internal_proto_grpc_proto_rawDesc
)

func file_internal_proto_grpc_proto_rawDescGZIP() []byte {
	file_internal_proto_grpc_proto_rawDescOnce.Do(func() {
		file_internal_proto_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_grpc_proto_rawDescData)
	})
	return file_internal_proto_grpc_proto_rawDescData
}

var file_internal_proto_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internal_proto_grpc_proto_goTypes = []interface{}{
	(*CustomerId)(nil),         // 0: proto.CustomerId
	(*Address)(nil),            // 1: proto.Address
	(*AddressutsRequest)(nil),  // 2: proto.AddressutsRequest
	(*AddressutsResponse)(nil), // 3: proto.AddressutsResponse
	(*AddressctuRequest)(nil),  // 4: proto.AddressctuRequest
	(*AddressctuResponse)(nil), // 5: proto.AddressctuResponse
	(*Credential)(nil),         // 6: proto.Credential
	(*CredentialRequest)(nil),  // 7: proto.CredentialRequest
	(*CredentialResponse)(nil), // 8: proto.CredentialResponse
}
var file_internal_proto_grpc_proto_depIdxs = []int32{
	1, // 0: proto.AddressutsRequest.address:type_name -> proto.Address
	0, // 1: proto.AddressctuRequest.customerid:type_name -> proto.CustomerId
	1, // 2: proto.AddressctuResponse.address:type_name -> proto.Address
	6, // 3: proto.CredentialRequest.credential:type_name -> proto.Credential
	2, // 4: proto.Service.AddressutsService:input_type -> proto.AddressutsRequest
	7, // 5: proto.Service.CredentialService:input_type -> proto.CredentialRequest
	3, // 6: proto.Service.AddressutsService:output_type -> proto.AddressutsResponse
	8, // 7: proto.Service.CredentialService:output_type -> proto.CredentialResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_proto_grpc_proto_init() }
func file_internal_proto_grpc_proto_init() {
	if File_internal_proto_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerId); i {
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
		file_internal_proto_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_internal_proto_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressutsRequest); i {
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
		file_internal_proto_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressutsResponse); i {
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
		file_internal_proto_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressctuRequest); i {
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
		file_internal_proto_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressctuResponse); i {
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
		file_internal_proto_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credential); i {
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
		file_internal_proto_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredentialRequest); i {
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
		file_internal_proto_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredentialResponse); i {
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
			RawDescriptor: file_internal_proto_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_grpc_proto_goTypes,
		DependencyIndexes: file_internal_proto_grpc_proto_depIdxs,
		MessageInfos:      file_internal_proto_grpc_proto_msgTypes,
	}.Build()
	File_internal_proto_grpc_proto = out.File
	file_internal_proto_grpc_proto_rawDesc = nil
	file_internal_proto_grpc_proto_goTypes = nil
	file_internal_proto_grpc_proto_depIdxs = nil
}
