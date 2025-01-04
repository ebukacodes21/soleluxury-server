// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.2
// source: rpc_product.proto

package pb

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

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId     int64    `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	CategoryId  int64    `protobuf:"varint,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	SizeId      int64    `protobuf:"varint,3,opt,name=size_id,json=sizeId,proto3" json:"size_id,omitempty"`
	ColorId     int64    `protobuf:"varint,4,opt,name=color_id,json=colorId,proto3" json:"color_id,omitempty"`
	Name        string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Price       float64  `protobuf:"fixed64,7,opt,name=price,proto3" json:"price,omitempty"`
	IsFeatured  bool     `protobuf:"varint,8,opt,name=is_featured,json=isFeatured,proto3" json:"is_featured,omitempty"`
	IsArchived  bool     `protobuf:"varint,9,opt,name=is_archived,json=isArchived,proto3" json:"is_archived,omitempty"`
	Images      []string `protobuf:"bytes,10,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_rpc_product_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *CreateProductRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *CreateProductRequest) GetSizeId() int64 {
	if x != nil {
		return x.SizeId
	}
	return 0
}

func (x *CreateProductRequest) GetColorId() int64 {
	if x != nil {
		return x.ColorId
	}
	return 0
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetIsFeatured() bool {
	if x != nil {
		return x.IsFeatured
	}
	return false
}

func (x *CreateProductRequest) GetIsArchived() bool {
	if x != nil {
		return x.IsArchived
	}
	return false
}

func (x *CreateProductRequest) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

type CreateProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *CreateProductResponse) Reset() {
	*x = CreateProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductResponse) ProtoMessage() {}

func (x *CreateProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductResponse.ProtoReflect.Descriptor instead.
func (*CreateProductResponse) Descriptor() ([]byte, []int) {
	return file_rpc_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_rpc_product_proto protoreflect.FileDescriptor

var file_rpc_product_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x69,
	0x7a, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x73, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x69, 0x73, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x73, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x62, 0x75, 0x6b, 0x61, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x32, 0x31,
	0x2f, 0x73, 0x6f, 0x6c, 0x65, 0x6c, 0x75, 0x78, 0x75, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_product_proto_rawDescOnce sync.Once
	file_rpc_product_proto_rawDescData = file_rpc_product_proto_rawDesc
)

func file_rpc_product_proto_rawDescGZIP() []byte {
	file_rpc_product_proto_rawDescOnce.Do(func() {
		file_rpc_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_product_proto_rawDescData)
	})
	return file_rpc_product_proto_rawDescData
}

var file_rpc_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_product_proto_goTypes = []interface{}{
	(*CreateProductRequest)(nil),  // 0: pb.CreateProductRequest
	(*CreateProductResponse)(nil), // 1: pb.CreateProductResponse
	(*Product)(nil),               // 2: pb.Product
}
var file_rpc_product_proto_depIdxs = []int32{
	2, // 0: pb.CreateProductResponse.product:type_name -> pb.Product
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_product_proto_init() }
func file_rpc_product_proto_init() {
	if File_rpc_product_proto != nil {
		return
	}
	file_general_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
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
		file_rpc_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductResponse); i {
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
			RawDescriptor: file_rpc_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_product_proto_goTypes,
		DependencyIndexes: file_rpc_product_proto_depIdxs,
		MessageInfos:      file_rpc_product_proto_msgTypes,
	}.Build()
	File_rpc_product_proto = out.File
	file_rpc_product_proto_rawDesc = nil
	file_rpc_product_proto_goTypes = nil
	file_rpc_product_proto_depIdxs = nil
}
