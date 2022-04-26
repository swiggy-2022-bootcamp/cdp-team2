// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: protos/shipping.proto

package shipping

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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId   string `protobuf:"bytes,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	Firstname    string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname     string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	AddressLine1 string `protobuf:"bytes,4,opt,name=addressLine1,proto3" json:"addressLine1,omitempty"`
	AddressLine2 string `protobuf:"bytes,5,opt,name=addressLine2,proto3" json:"addressLine2,omitempty"`
	City         string `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	PostCode     int32  `protobuf:"varint,7,opt,name=postCode,proto3" json:"postCode,omitempty"`
	CountryCode  string `protobuf:"bytes,8,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shipping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shipping_proto_msgTypes[0]
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
	return file_protos_shipping_proto_rawDescGZIP(), []int{0}
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

func (x *Address) GetAddressLine1() string {
	if x != nil {
		return x.AddressLine1
	}
	return ""
}

func (x *Address) GetAddressLine2() string {
	if x != nil {
		return x.AddressLine2
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetPostCode() int32 {
	if x != nil {
		return x.PostCode
	}
	return 0
}

func (x *Address) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type AddressstuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
}

func (x *AddressstuRequest) Reset() {
	*x = AddressstuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shipping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressstuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressstuRequest) ProtoMessage() {}

func (x *AddressstuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shipping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressstuRequest.ProtoReflect.Descriptor instead.
func (*AddressstuRequest) Descriptor() ([]byte, []int) {
	return file_protos_shipping_proto_rawDescGZIP(), []int{1}
}

func (x *AddressstuRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type AddressstuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []*Address `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
}

func (x *AddressstuResponse) Reset() {
	*x = AddressstuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shipping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressstuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressstuResponse) ProtoMessage() {}

func (x *AddressstuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shipping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressstuResponse.ProtoReflect.Descriptor instead.
func (*AddressstuResponse) Descriptor() ([]byte, []int) {
	return file_protos_shipping_proto_rawDescGZIP(), []int{2}
}

func (x *AddressstuResponse) GetAddress() []*Address {
	if x != nil {
		return x.Address
	}
	return nil
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
		mi := &file_protos_shipping_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressutsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressutsRequest) ProtoMessage() {}

func (x *AddressutsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shipping_proto_msgTypes[3]
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
	return file_protos_shipping_proto_rawDescGZIP(), []int{3}
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
		mi := &file_protos_shipping_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressutsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressutsResponse) ProtoMessage() {}

func (x *AddressutsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shipping_proto_msgTypes[4]
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
	return file_protos_shipping_proto_rawDescGZIP(), []int{4}
}

func (x *AddressutsResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

type OrderAddressUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	AddressId int32  `protobuf:"varint,2,opt,name=addressId,proto3" json:"addressId,omitempty"`
}

func (x *OrderAddressUpdateRequest) Reset() {
	*x = OrderAddressUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shipping_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAddressUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAddressUpdateRequest) ProtoMessage() {}

func (x *OrderAddressUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shipping_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAddressUpdateRequest.ProtoReflect.Descriptor instead.
func (*OrderAddressUpdateRequest) Descriptor() ([]byte, []int) {
	return file_protos_shipping_proto_rawDescGZIP(), []int{5}
}

func (x *OrderAddressUpdateRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderAddressUpdateRequest) GetAddressId() int32 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

type OrderAddressUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *OrderAddressUpdateResponse) Reset() {
	*x = OrderAddressUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_shipping_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAddressUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAddressUpdateResponse) ProtoMessage() {}

func (x *OrderAddressUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_shipping_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAddressUpdateResponse.ProtoReflect.Descriptor instead.
func (*OrderAddressUpdateResponse) Descriptor() ([]byte, []int) {
	return file_protos_shipping_proto_rawDescGZIP(), []int{6}
}

func (x *OrderAddressUpdateResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

var File_protos_shipping_proto protoreflect.FileDescriptor

var file_protos_shipping_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd,
	0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c,
	0x69, 0x6e, 0x65, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x31, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x32, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x33,
	0x0a, 0x11, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x73, 0x74, 0x75, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x73, 0x74,
	0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x3d, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x75, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x30, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x75, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a, 0x19, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x1a, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x85, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4a, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x75, 0x74, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x75, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x75, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x11, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x73, 0x74, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x73, 0x74, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x73, 0x74, 0x75, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x19, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2f,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_shipping_proto_rawDescOnce sync.Once
	file_protos_shipping_proto_rawDescData = file_protos_shipping_proto_rawDesc
)

func file_protos_shipping_proto_rawDescGZIP() []byte {
	file_protos_shipping_proto_rawDescOnce.Do(func() {
		file_protos_shipping_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_shipping_proto_rawDescData)
	})
	return file_protos_shipping_proto_rawDescData
}

var file_protos_shipping_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_shipping_proto_goTypes = []interface{}{
	(*Address)(nil),                    // 0: proto.Address
	(*AddressstuRequest)(nil),          // 1: proto.AddressstuRequest
	(*AddressstuResponse)(nil),         // 2: proto.AddressstuResponse
	(*AddressutsRequest)(nil),          // 3: proto.AddressutsRequest
	(*AddressutsResponse)(nil),         // 4: proto.AddressutsResponse
	(*OrderAddressUpdateRequest)(nil),  // 5: proto.OrderAddressUpdateRequest
	(*OrderAddressUpdateResponse)(nil), // 6: proto.OrderAddressUpdateResponse
}
var file_protos_shipping_proto_depIdxs = []int32{
	0, // 0: proto.AddressstuResponse.address:type_name -> proto.Address
	0, // 1: proto.AddressutsRequest.address:type_name -> proto.Address
	3, // 2: proto.Service.AddressutsService:input_type -> proto.AddressutsRequest
	1, // 3: proto.Service.AddressstuService:input_type -> proto.AddressstuRequest
	5, // 4: proto.Service.OrderAddressUpdateService:input_type -> proto.OrderAddressUpdateRequest
	4, // 5: proto.Service.AddressutsService:output_type -> proto.AddressutsResponse
	2, // 6: proto.Service.AddressstuService:output_type -> proto.AddressstuResponse
	6, // 7: proto.Service.OrderAddressUpdateService:output_type -> proto.OrderAddressUpdateResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_shipping_proto_init() }
func file_protos_shipping_proto_init() {
	if File_protos_shipping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_shipping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_shipping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressstuRequest); i {
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
		file_protos_shipping_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressstuResponse); i {
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
		file_protos_shipping_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_shipping_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_shipping_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAddressUpdateRequest); i {
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
		file_protos_shipping_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAddressUpdateResponse); i {
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
			RawDescriptor: file_protos_shipping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_shipping_proto_goTypes,
		DependencyIndexes: file_protos_shipping_proto_depIdxs,
		MessageInfos:      file_protos_shipping_proto_msgTypes,
	}.Build()
	File_protos_shipping_proto = out.File
	file_protos_shipping_proto_rawDesc = nil
	file_protos_shipping_proto_goTypes = nil
	file_protos_shipping_proto_depIdxs = nil
}