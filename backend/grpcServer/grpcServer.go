// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: miniZomato.proto

package __

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

type Restaurant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rating      string `protobuf:"bytes,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Cuisine     string `protobuf:"bytes,3,opt,name=cuisine,proto3" json:"cuisine,omitempty"`
	Address     string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	OpeningTime string `protobuf:"bytes,5,opt,name=openingTime,proto3" json:"openingTime,omitempty"`
	ClosingTime string `protobuf:"bytes,6,opt,name=closingTime,proto3" json:"closingTime,omitempty"`
	CostForTwo  string `protobuf:"bytes,7,opt,name=costForTwo,proto3" json:"costForTwo,omitempty"`
}

func (x *Restaurant) Reset() {
	*x = Restaurant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_miniZomato_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restaurant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restaurant) ProtoMessage() {}

func (x *Restaurant) ProtoReflect() protoreflect.Message {
	mi := &file_miniZomato_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restaurant.ProtoReflect.Descriptor instead.
func (*Restaurant) Descriptor() ([]byte, []int) {
	return file_miniZomato_proto_rawDescGZIP(), []int{0}
}

func (x *Restaurant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Restaurant) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

func (x *Restaurant) GetCuisine() string {
	if x != nil {
		return x.Cuisine
	}
	return ""
}

func (x *Restaurant) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Restaurant) GetOpeningTime() string {
	if x != nil {
		return x.OpeningTime
	}
	return ""
}

func (x *Restaurant) GetClosingTime() string {
	if x != nil {
		return x.ClosingTime
	}
	return ""
}

func (x *Restaurant) GetCostForTwo() string {
	if x != nil {
		return x.CostForTwo
	}
	return ""
}

//containing result for inserting new data
type AddedConfirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddedConfirmation string `protobuf:"bytes,1,opt,name=addedConfirmation,proto3" json:"addedConfirmation,omitempty"`
}

func (x *AddedConfirmation) Reset() {
	*x = AddedConfirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_miniZomato_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddedConfirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddedConfirmation) ProtoMessage() {}

func (x *AddedConfirmation) ProtoReflect() protoreflect.Message {
	mi := &file_miniZomato_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddedConfirmation.ProtoReflect.Descriptor instead.
func (*AddedConfirmation) Descriptor() ([]byte, []int) {
	return file_miniZomato_proto_rawDescGZIP(), []int{1}
}

func (x *AddedConfirmation) GetAddedConfirmation() string {
	if x != nil {
		return x.AddedConfirmation
	}
	return ""
}

//containing request for list
type GetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetListRequest string `protobuf:"bytes,1,opt,name=getListRequest,proto3" json:"getListRequest,omitempty"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_miniZomato_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_miniZomato_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_miniZomato_proto_rawDescGZIP(), []int{2}
}

func (x *GetListRequest) GetGetListRequest() string {
	if x != nil {
		return x.GetListRequest
	}
	return ""
}

//containing marshalled list
type RestaurantList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantList string `protobuf:"bytes,1,opt,name=restaurantList,proto3" json:"restaurantList,omitempty"`
}

func (x *RestaurantList) Reset() {
	*x = RestaurantList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_miniZomato_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantList) ProtoMessage() {}

func (x *RestaurantList) ProtoReflect() protoreflect.Message {
	mi := &file_miniZomato_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantList.ProtoReflect.Descriptor instead.
func (*RestaurantList) Descriptor() ([]byte, []int) {
	return file_miniZomato_proto_rawDescGZIP(), []int{3}
}

func (x *RestaurantList) GetRestaurantList() string {
	if x != nil {
		return x.RestaurantList
	}
	return ""
}

var File_miniZomato_proto protoreflect.FileDescriptor

var file_miniZomato_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x69, 0x5a, 0x6f, 0x6d, 0x61, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x69, 0x6e, 0x69, 0x5a, 0x6f, 0x6d, 0x61, 0x74, 0x6f, 0x22, 0xd0,
	0x01, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x69,
	0x73, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x75, 0x69, 0x73,
	0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x54, 0x77, 0x6f, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x54, 0x77,
	0x6f, 0x22, 0x41, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x65, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x61, 0x64, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xab, 0x01, 0x0a, 0x12, 0x41, 0x64, 0x64,
	0x5f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12,
	0x4b, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x73, 0x12, 0x1a, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5a, 0x6f, 0x6d, 0x61, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5a, 0x6f, 0x6d, 0x61, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x2e,
	0x6d, 0x69, 0x6e, 0x69, 0x5a, 0x6f, 0x6d, 0x61, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5a, 0x6f, 0x6d, 0x61,
	0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_miniZomato_proto_rawDescOnce sync.Once
	file_miniZomato_proto_rawDescData = file_miniZomato_proto_rawDesc
)

func file_miniZomato_proto_rawDescGZIP() []byte {
	file_miniZomato_proto_rawDescOnce.Do(func() {
		file_miniZomato_proto_rawDescData = protoimpl.X.CompressGZIP(file_miniZomato_proto_rawDescData)
	})
	return file_miniZomato_proto_rawDescData
}

var file_miniZomato_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_miniZomato_proto_goTypes = []interface{}{
	(*Restaurant)(nil),        // 0: miniZomato.Restaurant
	(*AddedConfirmation)(nil), // 1: miniZomato.AddedConfirmation
	(*GetListRequest)(nil),    // 2: miniZomato.GetListRequest
	(*RestaurantList)(nil),    // 3: miniZomato.RestaurantList
}
var file_miniZomato_proto_depIdxs = []int32{
	2, // 0: miniZomato.Add_get_restaurant.ListRestaurants:input_type -> miniZomato.GetListRequest
	0, // 1: miniZomato.Add_get_restaurant.AddRestaurant:input_type -> miniZomato.Restaurant
	3, // 2: miniZomato.Add_get_restaurant.ListRestaurants:output_type -> miniZomato.RestaurantList
	1, // 3: miniZomato.Add_get_restaurant.AddRestaurant:output_type -> miniZomato.AddedConfirmation
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_miniZomato_proto_init() }
func file_miniZomato_proto_init() {
	if File_miniZomato_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_miniZomato_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restaurant); i {
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
		file_miniZomato_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddedConfirmation); i {
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
		file_miniZomato_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRequest); i {
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
		file_miniZomato_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantList); i {
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
			RawDescriptor: file_miniZomato_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_miniZomato_proto_goTypes,
		DependencyIndexes: file_miniZomato_proto_depIdxs,
		MessageInfos:      file_miniZomato_proto_msgTypes,
	}.Build()
	File_miniZomato_proto = out.File
	file_miniZomato_proto_rawDesc = nil
	file_miniZomato_proto_goTypes = nil
	file_miniZomato_proto_depIdxs = nil
}
