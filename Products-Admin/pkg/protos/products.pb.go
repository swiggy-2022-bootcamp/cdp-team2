// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: products.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ProductDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LanguageID     int64  `protobuf:"varint,1,opt,name=LanguageID,json=languageID,proto3" json:"LanguageID,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Description    string `protobuf:"bytes,3,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	MetaTitle      string `protobuf:"bytes,4,opt,name=MetaTitle,json=metaTitle,proto3" json:"MetaTitle,omitempty"`
	MetaDescriptio string `protobuf:"bytes,5,opt,name=MetaDescriptio,json=metaDescriptio,proto3" json:"MetaDescriptio,omitempty"`
	MetaKeyword    string `protobuf:"bytes,6,opt,name=MetaKeyword,json=metaKeyword,proto3" json:"MetaKeyword,omitempty"`
	Tag            string `protobuf:"bytes,7,opt,name=Tag,json=tag,proto3" json:"Tag,omitempty"`
}

func (x *ProductDescription) Reset() {
	*x = ProductDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDescription) ProtoMessage() {}

func (x *ProductDescription) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDescription.ProtoReflect.Descriptor instead.
func (*ProductDescription) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{0}
}

func (x *ProductDescription) GetLanguageID() int64 {
	if x != nil {
		return x.LanguageID
	}
	return 0
}

func (x *ProductDescription) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductDescription) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductDescription) GetMetaTitle() string {
	if x != nil {
		return x.MetaTitle
	}
	return ""
}

func (x *ProductDescription) GetMetaDescriptio() string {
	if x != nil {
		return x.MetaDescriptio
	}
	return ""
}

func (x *ProductDescription) GetMetaKeyword() string {
	if x != nil {
		return x.MetaKeyword
	}
	return ""
}

func (x *ProductDescription) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type ProductSeoUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword    string `protobuf:"bytes,1,opt,name=Keyword,json=keyword,proto3" json:"Keyword,omitempty"`
	LanguageID int64  `protobuf:"varint,2,opt,name=LanguageID,json=languageID,proto3" json:"LanguageID,omitempty"`
	StoreID    int64  `protobuf:"varint,3,opt,name=StoreID,json=storeID,proto3" json:"StoreID,omitempty"`
}

func (x *ProductSeoUrl) Reset() {
	*x = ProductSeoUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSeoUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSeoUrl) ProtoMessage() {}

func (x *ProductSeoUrl) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSeoUrl.ProtoReflect.Descriptor instead.
func (*ProductSeoUrl) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{1}
}

func (x *ProductSeoUrl) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ProductSeoUrl) GetLanguageID() int64 {
	if x != nil {
		return x.LanguageID
	}
	return 0
}

func (x *ProductSeoUrl) GetStoreID() int64 {
	if x != nil {
		return x.StoreID
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 int64                 `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Model              string                `protobuf:"bytes,2,opt,name=Model,json=model,proto3" json:"Model,omitempty"`
	Quantity           int64                 `protobuf:"varint,3,opt,name=Quantity,json=quantity,proto3" json:"Quantity,omitempty"`
	Price              string                `protobuf:"bytes,4,opt,name=Price,json=price,proto3" json:"Price,omitempty"`
	ManufacturerID     int64                 `protobuf:"varint,5,opt,name=ManufacturerID,json=manufacturerID,proto3" json:"ManufacturerID,omitempty"`
	Sku                string                `protobuf:"bytes,6,opt,name=Sku,json=sku,proto3" json:"Sku,omitempty"`
	ProductSeoUrl      *ProductSeoUrl        `protobuf:"bytes,7,opt,name=ProductSeoUrl,json=productSeoUrl,proto3" json:"ProductSeoUrl,omitempty"`
	Points             int64                 `protobuf:"varint,8,opt,name=Points,json=points,proto3" json:"Points,omitempty"`
	Rewards            int64                 `protobuf:"varint,9,opt,name=Rewards,json=rewards,proto3" json:"Rewards,omitempty"`
	Image              string                `protobuf:"bytes,10,opt,name=Image,json=image,proto3" json:"Image,omitempty"`
	ShippingID         int64                 `protobuf:"varint,11,opt,name=ShippingID,json=shippingID,proto3" json:"ShippingID,omitempty"`
	Weight             int64                 `protobuf:"varint,12,opt,name=Weight,json=weight,proto3" json:"Weight,omitempty"`
	Langth             int64                 `protobuf:"varint,13,opt,name=Langth,json=langth,proto3" json:"Langth,omitempty"`
	Width              int64                 `protobuf:"varint,14,opt,name=Width,json=width,proto3" json:"Width,omitempty"`
	Height             int64                 `protobuf:"varint,15,opt,name=Height,json=height,proto3" json:"Height,omitempty"`
	Minimum            int64                 `protobuf:"varint,16,opt,name=Minimum,json=minimum,proto3" json:"Minimum,omitempty"`
	ProductRelated     []int64               `protobuf:"varint,17,rep,packed,name=ProductRelated,json=productRelated,proto3" json:"ProductRelated,omitempty"`
	ProductDescription []*ProductDescription `protobuf:"bytes,18,rep,name=ProductDescription,json=productDescription,proto3" json:"ProductDescription,omitempty"`
	ProductCategory    []int64               `protobuf:"varint,19,rep,packed,name=ProductCategory,json=productCategory,proto3" json:"ProductCategory,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Product) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Product) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Product) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Product) GetManufacturerID() int64 {
	if x != nil {
		return x.ManufacturerID
	}
	return 0
}

func (x *Product) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *Product) GetProductSeoUrl() *ProductSeoUrl {
	if x != nil {
		return x.ProductSeoUrl
	}
	return nil
}

func (x *Product) GetPoints() int64 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *Product) GetRewards() int64 {
	if x != nil {
		return x.Rewards
	}
	return 0
}

func (x *Product) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Product) GetShippingID() int64 {
	if x != nil {
		return x.ShippingID
	}
	return 0
}

func (x *Product) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Product) GetLangth() int64 {
	if x != nil {
		return x.Langth
	}
	return 0
}

func (x *Product) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Product) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Product) GetMinimum() int64 {
	if x != nil {
		return x.Minimum
	}
	return 0
}

func (x *Product) GetProductRelated() []int64 {
	if x != nil {
		return x.ProductRelated
	}
	return nil
}

func (x *Product) GetProductDescription() []*ProductDescription {
	if x != nil {
		return x.ProductDescription
	}
	return nil
}

func (x *Product) GetProductCategory() []int64 {
	if x != nil {
		return x.ProductCategory
	}
	return nil
}

type ProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=Products,json=products,proto3" json:"Products,omitempty"`
}

func (x *ProductsResponse) Reset() {
	*x = ProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsResponse) ProtoMessage() {}

func (x *ProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsResponse.ProtoReflect.Descriptor instead.
func (*ProductsResponse) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{3}
}

func (x *ProductsResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{4}
}

type CategoryIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryID int64 `protobuf:"varint,1,opt,name=categoryID,proto3" json:"categoryID,omitempty"`
}

func (x *CategoryIDRequest) Reset() {
	*x = CategoryIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryIDRequest) ProtoMessage() {}

func (x *CategoryIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryIDRequest.ProtoReflect.Descriptor instead.
func (*CategoryIDRequest) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{5}
}

func (x *CategoryIDRequest) GetCategoryID() int64 {
	if x != nil {
		return x.CategoryID
	}
	return 0
}

type BoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *BoolResponse) Reset() {
	*x = BoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolResponse) ProtoMessage() {}

func (x *BoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolResponse.ProtoReflect.Descriptor instead.
func (*BoolResponse) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{6}
}

func (x *BoolResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_products_proto protoreflect.FileDescriptor

var file_products_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65, 0x74,
	0x61, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x4d,
	0x65, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x54, 0x61, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22,
	0x63, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x6f, 0x55, 0x72, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x44, 0x22, 0xd6, 0x04, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x75,
	0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x10, 0x0a, 0x03, 0x53, 0x6b, 0x75, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x6b, 0x75, 0x12, 0x3b, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x6f,
	0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x6f, 0x55, 0x72, 0x6c,
	0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x68, 0x69, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6c, 0x61, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x12,
	0x26, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x11, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x12, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x13, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x3f, 0x0a,
	0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2b, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x0e,
	0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33,
	0x0a, 0x11, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x44, 0x22, 0x26, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x32, 0xa8, 0x01, 0x0a, 0x10,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x46, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x77, 0x69, 0x67, 0x67, 0x79, 0x2d, 0x32, 0x30, 0x32, 0x32,
	0x2d, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x2f, 0x63, 0x64, 0x70, 0x2d, 0x74, 0x65,
	0x61, 0x6d, 0x32, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2d, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_products_proto_rawDescOnce sync.Once
	file_products_proto_rawDescData = file_products_proto_rawDesc
)

func file_products_proto_rawDescGZIP() []byte {
	file_products_proto_rawDescOnce.Do(func() {
		file_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_products_proto_rawDescData)
	})
	return file_products_proto_rawDescData
}

var file_products_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_products_proto_goTypes = []interface{}{
	(*ProductDescription)(nil), // 0: protos.ProductDescription
	(*ProductSeoUrl)(nil),      // 1: protos.ProductSeoUrl
	(*Product)(nil),            // 2: protos.Product
	(*ProductsResponse)(nil),   // 3: protos.ProductsResponse
	(*EmptyRequest)(nil),       // 4: protos.EmptyRequest
	(*CategoryIDRequest)(nil),  // 5: protos.CategoryIDRequest
	(*BoolResponse)(nil),       // 6: protos.BoolResponse
}
var file_products_proto_depIdxs = []int32{
	1, // 0: protos.Product.ProductSeoUrl:type_name -> protos.ProductSeoUrl
	0, // 1: protos.Product.ProductDescription:type_name -> protos.ProductDescription
	2, // 2: protos.ProductsResponse.Products:type_name -> protos.Product
	4, // 3: protos.ProductsServices.GetAvailableProducts:input_type -> protos.EmptyRequest
	5, // 4: protos.ProductsServices.CheckProductsWithCategory:input_type -> protos.CategoryIDRequest
	3, // 5: protos.ProductsServices.GetAvailableProducts:output_type -> protos.ProductsResponse
	6, // 6: protos.ProductsServices.CheckProductsWithCategory:output_type -> protos.BoolResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_products_proto_init() }
func file_products_proto_init() {
	if File_products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_products_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDescription); i {
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
		file_products_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSeoUrl); i {
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
		file_products_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_products_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsResponse); i {
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
		file_products_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_products_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryIDRequest); i {
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
		file_products_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolResponse); i {
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
			RawDescriptor: file_products_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_products_proto_goTypes,
		DependencyIndexes: file_products_proto_depIdxs,
		MessageInfos:      file_products_proto_msgTypes,
	}.Build()
	File_products_proto = out.File
	file_products_proto_rawDesc = nil
	file_products_proto_goTypes = nil
	file_products_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductsServicesClient is the client API for ProductsServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductsServicesClient interface {
	GetAvailableProducts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ProductsResponse, error)
	CheckProductsWithCategory(ctx context.Context, in *CategoryIDRequest, opts ...grpc.CallOption) (*BoolResponse, error)
}

type productsServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsServicesClient(cc grpc.ClientConnInterface) ProductsServicesClient {
	return &productsServicesClient{cc}
}

func (c *productsServicesClient) GetAvailableProducts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ProductsResponse, error) {
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, "/protos.ProductsServices/GetAvailableProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServicesClient) CheckProductsWithCategory(ctx context.Context, in *CategoryIDRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/protos.ProductsServices/CheckProductsWithCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServicesServer is the server API for ProductsServices service.
type ProductsServicesServer interface {
	GetAvailableProducts(context.Context, *EmptyRequest) (*ProductsResponse, error)
	CheckProductsWithCategory(context.Context, *CategoryIDRequest) (*BoolResponse, error)
}

// UnimplementedProductsServicesServer can be embedded to have forward compatible implementations.
type UnimplementedProductsServicesServer struct {
}

func (*UnimplementedProductsServicesServer) GetAvailableProducts(context.Context, *EmptyRequest) (*ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableProducts not implemented")
}
func (*UnimplementedProductsServicesServer) CheckProductsWithCategory(context.Context, *CategoryIDRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckProductsWithCategory not implemented")
}

func RegisterProductsServicesServer(s *grpc.Server, srv ProductsServicesServer) {
	s.RegisterService(&_ProductsServices_serviceDesc, srv)
}

func _ProductsServices_GetAvailableProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServicesServer).GetAvailableProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ProductsServices/GetAvailableProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServicesServer).GetAvailableProducts(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsServices_CheckProductsWithCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServicesServer).CheckProductsWithCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ProductsServices/CheckProductsWithCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServicesServer).CheckProductsWithCategory(ctx, req.(*CategoryIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductsServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ProductsServices",
	HandlerType: (*ProductsServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvailableProducts",
			Handler:    _ProductsServices_GetAvailableProducts_Handler,
		},
		{
			MethodName: "CheckProductsWithCategory",
			Handler:    _ProductsServices_CheckProductsWithCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "products.proto",
}
