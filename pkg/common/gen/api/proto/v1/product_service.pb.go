// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/proto/v1/product_service.proto

package pbv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type ProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter    *ProductsRequestFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Paginator *Paginator             `protobuf:"bytes,2,opt,name=paginator,proto3" json:"paginator,omitempty"`
}

func (x *ProductsRequest) Reset() {
	*x = ProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_product_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsRequest) ProtoMessage() {}

func (x *ProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_product_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsRequest.ProtoReflect.Descriptor instead.
func (*ProductsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_product_service_proto_rawDescGZIP(), []int{0}
}

func (x *ProductsRequest) GetFilter() *ProductsRequestFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ProductsRequest) GetPaginator() *Paginator {
	if x != nil {
		return x.Paginator
	}
	return nil
}

type ProductsRequestFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UuidIn       []string           `protobuf:"bytes,1,rep,name=uuid_in,json=uuidIn,proto3" json:"uuid_in,omitempty"`
	NameContains string             `protobuf:"bytes,2,opt,name=name_contains,json=nameContains,proto3" json:"name_contains,omitempty"`
	PriceGt      float64            `protobuf:"fixed64,3,opt,name=price_gt,json=priceGt,proto3" json:"price_gt,omitempty"`
	PriceLt      float64            `protobuf:"fixed64,4,opt,name=price_lt,json=priceLt,proto3" json:"price_lt,omitempty"`
	PropertiesEq []*ProductProperty `protobuf:"bytes,5,rep,name=properties_eq,json=propertiesEq,proto3" json:"properties_eq,omitempty"`
}

func (x *ProductsRequestFilter) Reset() {
	*x = ProductsRequestFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_product_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsRequestFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsRequestFilter) ProtoMessage() {}

func (x *ProductsRequestFilter) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_product_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsRequestFilter.ProtoReflect.Descriptor instead.
func (*ProductsRequestFilter) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_product_service_proto_rawDescGZIP(), []int{1}
}

func (x *ProductsRequestFilter) GetUuidIn() []string {
	if x != nil {
		return x.UuidIn
	}
	return nil
}

func (x *ProductsRequestFilter) GetNameContains() string {
	if x != nil {
		return x.NameContains
	}
	return ""
}

func (x *ProductsRequestFilter) GetPriceGt() float64 {
	if x != nil {
		return x.PriceGt
	}
	return 0
}

func (x *ProductsRequestFilter) GetPriceLt() float64 {
	if x != nil {
		return x.PriceLt
	}
	return 0
}

func (x *ProductsRequestFilter) GetPropertiesEq() []*ProductProperty {
	if x != nil {
		return x.PropertiesEq
	}
	return nil
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Product `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Count uint32     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_product_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_product_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_product_service_proto_rawDescGZIP(), []int{2}
}

func (x *ProductResponse) GetItems() []*Product {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ProductResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetByProductUUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductUuid string `protobuf:"bytes,1,opt,name=product_uuid,json=productUuid,proto3" json:"product_uuid,omitempty"`
}

func (x *GetByProductUUIDRequest) Reset() {
	*x = GetByProductUUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_product_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByProductUUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByProductUUIDRequest) ProtoMessage() {}

func (x *GetByProductUUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_product_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByProductUUIDRequest.ProtoReflect.Descriptor instead.
func (*GetByProductUUIDRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_product_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetByProductUUIDRequest) GetProductUuid() string {
	if x != nil {
		return x.ProductUuid
	}
	return ""
}

var File_api_proto_v1_product_service_proto protoreflect.FileDescriptor

var file_api_proto_v1_product_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x77, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x22, 0xc8, 0x01, 0x0a, 0x15, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x75, 0x69, 0x64, 0x49, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x47, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x65, 0x71, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x45, 0x71, 0x22, 0x4d, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x75, 0x69,
	0x64, 0x32, 0xbc, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x94, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42,
	0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x92, 0x41, 0x3b, 0x12, 0x39, 0xd0, 0x9f,
	0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb5,
	0x20, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xbe, 0xd0, 0xb4, 0xd1, 0x83, 0xd0, 0xba, 0xd1, 0x82, 0xd0,
	0xbe, 0xd0, 0xb2, 0x20, 0xd0, 0xbf, 0xd0, 0xbe, 0x20, 0xd1, 0x84, 0xd0, 0xb8, 0xd0, 0xbb, 0xd1,
	0x8c, 0xd1, 0x82, 0xd1, 0x80, 0xd1, 0x83, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x92, 0x01, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x55,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x55, 0x92, 0x41, 0x2f, 0x12, 0x2d,
	0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8,
	0xd0, 0xb5, 0x20, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xbe, 0xd0, 0xb4, 0xd1, 0x83, 0xd0, 0xba, 0xd1,
	0x82, 0xd0, 0xb0, 0x20, 0xd0, 0xbf, 0xd0, 0xbe, 0x20, 0x55, 0x55, 0x49, 0x44, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x7d,
	0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x79, 0x6f, 0x72, 0x74, 0x79, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_product_service_proto_rawDescOnce sync.Once
	file_api_proto_v1_product_service_proto_rawDescData = file_api_proto_v1_product_service_proto_rawDesc
)

func file_api_proto_v1_product_service_proto_rawDescGZIP() []byte {
	file_api_proto_v1_product_service_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_product_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_product_service_proto_rawDescData)
	})
	return file_api_proto_v1_product_service_proto_rawDescData
}

var file_api_proto_v1_product_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_v1_product_service_proto_goTypes = []interface{}{
	(*ProductsRequest)(nil),         // 0: pb.v1.ProductsRequest
	(*ProductsRequestFilter)(nil),   // 1: pb.v1.ProductsRequestFilter
	(*ProductResponse)(nil),         // 2: pb.v1.ProductResponse
	(*GetByProductUUIDRequest)(nil), // 3: pb.v1.GetByProductUUIDRequest
	(*Paginator)(nil),               // 4: pb.v1.Paginator
	(*ProductProperty)(nil),         // 5: pb.v1.ProductProperty
	(*Product)(nil),                 // 6: pb.v1.Product
}
var file_api_proto_v1_product_service_proto_depIdxs = []int32{
	1, // 0: pb.v1.ProductsRequest.filter:type_name -> pb.v1.ProductsRequestFilter
	4, // 1: pb.v1.ProductsRequest.paginator:type_name -> pb.v1.Paginator
	5, // 2: pb.v1.ProductsRequestFilter.properties_eq:type_name -> pb.v1.ProductProperty
	6, // 3: pb.v1.ProductResponse.items:type_name -> pb.v1.Product
	0, // 4: pb.v1.ProductService.GetAllByFilter:input_type -> pb.v1.ProductsRequest
	3, // 5: pb.v1.ProductService.GetByUUID:input_type -> pb.v1.GetByProductUUIDRequest
	2, // 6: pb.v1.ProductService.GetAllByFilter:output_type -> pb.v1.ProductResponse
	6, // 7: pb.v1.ProductService.GetByUUID:output_type -> pb.v1.Product
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_v1_product_service_proto_init() }
func file_api_proto_v1_product_service_proto_init() {
	if File_api_proto_v1_product_service_proto != nil {
		return
	}
	file_api_proto_v1_common_proto_init()
	file_api_proto_v1_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_product_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsRequest); i {
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
		file_api_proto_v1_product_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsRequestFilter); i {
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
		file_api_proto_v1_product_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
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
		file_api_proto_v1_product_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByProductUUIDRequest); i {
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
			RawDescriptor: file_api_proto_v1_product_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_v1_product_service_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_product_service_proto_depIdxs,
		MessageInfos:      file_api_proto_v1_product_service_proto_msgTypes,
	}.Build()
	File_api_proto_v1_product_service_proto = out.File
	file_api_proto_v1_product_service_proto_rawDesc = nil
	file_api_proto_v1_product_service_proto_goTypes = nil
	file_api_proto_v1_product_service_proto_depIdxs = nil
}