// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: proto/devicetoken.proto

package user

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

type DeviceTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DeviceTokenRequest) Reset() {
	*x = DeviceTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_devicetoken_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceTokenRequest) ProtoMessage() {}

func (x *DeviceTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_devicetoken_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceTokenRequest.ProtoReflect.Descriptor instead.
func (*DeviceTokenRequest) Descriptor() ([]byte, []int) {
	return file_proto_devicetoken_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceTokenRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeviceToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId       int32  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	IsProduction int32  `protobuf:"varint,3,opt,name=isProduction,proto3" json:"isProduction,omitempty"`
	Os           string `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	Token        string `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	DeviceId     string `protobuf:"bytes,6,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
}

func (x *DeviceToken) Reset() {
	*x = DeviceToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_devicetoken_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceToken) ProtoMessage() {}

func (x *DeviceToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_devicetoken_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceToken.ProtoReflect.Descriptor instead.
func (*DeviceToken) Descriptor() ([]byte, []int) {
	return file_proto_devicetoken_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceToken) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeviceToken) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeviceToken) GetIsProduction() int32 {
	if x != nil {
		return x.IsProduction
	}
	return 0
}

func (x *DeviceToken) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *DeviceToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeviceToken) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type DeviceTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceTokens []*DeviceToken `protobuf:"bytes,1,rep,name=deviceTokens,proto3" json:"deviceTokens,omitempty"`
}

func (x *DeviceTokenResponse) Reset() {
	*x = DeviceTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_devicetoken_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceTokenResponse) ProtoMessage() {}

func (x *DeviceTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_devicetoken_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceTokenResponse.ProtoReflect.Descriptor instead.
func (*DeviceTokenResponse) Descriptor() ([]byte, []int) {
	return file_proto_devicetoken_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceTokenResponse) GetDeviceTokens() []*DeviceToken {
	if x != nil {
		return x.DeviceTokens
	}
	return nil
}

var File_proto_devicetoken_proto protoreflect.FileDescriptor

var file_proto_devicetoken_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x65, 0x6d, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a,
	0x12, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9b, 0x01, 0x0a, 0x0b,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x73, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x13, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x32, 0x8a, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x64, 0x12, 0x18, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a,
	0x01, 0x2a, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x67, 0x65, 0x74, 0x2d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2d, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2d, 0x69, 0x64, 0x42, 0x07, 0x5a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_devicetoken_proto_rawDescOnce sync.Once
	file_proto_devicetoken_proto_rawDescData = file_proto_devicetoken_proto_rawDesc
)

func file_proto_devicetoken_proto_rawDescGZIP() []byte {
	file_proto_devicetoken_proto_rawDescOnce.Do(func() {
		file_proto_devicetoken_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_devicetoken_proto_rawDescData)
	})
	return file_proto_devicetoken_proto_rawDescData
}

var file_proto_devicetoken_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_devicetoken_proto_goTypes = []interface{}{
	(*DeviceTokenRequest)(nil),  // 0: demo.DeviceTokenRequest
	(*DeviceToken)(nil),         // 1: demo.DeviceToken
	(*DeviceTokenResponse)(nil), // 2: demo.DeviceTokenResponse
}
var file_proto_devicetoken_proto_depIdxs = []int32{
	1, // 0: demo.DeviceTokenResponse.deviceTokens:type_name -> demo.DeviceToken
	0, // 1: demo.DeviceTokenService.GetDeviceTokenId:input_type -> demo.DeviceTokenRequest
	2, // 2: demo.DeviceTokenService.GetDeviceTokenId:output_type -> demo.DeviceTokenResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_devicetoken_proto_init() }
func file_proto_devicetoken_proto_init() {
	if File_proto_devicetoken_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_devicetoken_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceTokenRequest); i {
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
		file_proto_devicetoken_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceToken); i {
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
		file_proto_devicetoken_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceTokenResponse); i {
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
			RawDescriptor: file_proto_devicetoken_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_devicetoken_proto_goTypes,
		DependencyIndexes: file_proto_devicetoken_proto_depIdxs,
		MessageInfos:      file_proto_devicetoken_proto_msgTypes,
	}.Build()
	File_proto_devicetoken_proto = out.File
	file_proto_devicetoken_proto_rawDesc = nil
	file_proto_devicetoken_proto_goTypes = nil
	file_proto_devicetoken_proto_depIdxs = nil
}
