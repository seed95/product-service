// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: microService.proto

package micro

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

type RequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	// Headers
	//
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	//
	OpCode      int32  `protobuf:"varint,10,opt,name=opCode,proto3" json:"opCode,omitempty"`
	Username    string `protobuf:"bytes,11,opt,name=username,proto3" json:"username,omitempty"`
	CompanyName string `protobuf:"bytes,12,opt,name=companyName,proto3" json:"companyName,omitempty"`
	//
	// Payload
	//
	Payload string `protobuf:"bytes,51,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *RequestMessage) Reset() {
	*x = RequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMessage) ProtoMessage() {}

func (x *RequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_microService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMessage.ProtoReflect.Descriptor instead.
func (*RequestMessage) Descriptor() ([]byte, []int) {
	return file_microService_proto_rawDescGZIP(), []int{0}
}

func (x *RequestMessage) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *RequestMessage) GetOpCode() int32 {
	if x != nil {
		return x.OpCode
	}
	return 0
}

func (x *RequestMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RequestMessage) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *RequestMessage) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type ResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	// Payload & status
	//
	StatusCode    int32  `protobuf:"varint,51,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage string `protobuf:"bytes,52,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
	Payload       string `protobuf:"bytes,53,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ResponseMessage) Reset() {
	*x = ResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMessage) ProtoMessage() {}

func (x *ResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_microService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMessage.ProtoReflect.Descriptor instead.
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return file_microService_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseMessage) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ResponseMessage) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *ResponseMessage) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_microService_proto protoreflect.FileDescriptor

var file_microService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x0e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x70, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x33, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x71, 0x0a, 0x0f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x33, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x34,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x35,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x4e, 0x0a,
	0x0c, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x0b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a,
	0x06, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_microService_proto_rawDescOnce sync.Once
	file_microService_proto_rawDescData = file_microService_proto_rawDesc
)

func file_microService_proto_rawDescGZIP() []byte {
	file_microService_proto_rawDescOnce.Do(func() {
		file_microService_proto_rawDescData = protoimpl.X.CompressGZIP(file_microService_proto_rawDescData)
	})
	return file_microService_proto_rawDescData
}

var file_microService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_microService_proto_goTypes = []interface{}{
	(*RequestMessage)(nil),  // 0: micro.RequestMessage
	(*ResponseMessage)(nil), // 1: micro.ResponseMessage
}
var file_microService_proto_depIdxs = []int32{
	0, // 0: micro.MicroService.generalCall:input_type -> micro.RequestMessage
	1, // 1: micro.MicroService.generalCall:output_type -> micro.ResponseMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_microService_proto_init() }
func file_microService_proto_init() {
	if File_microService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_microService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMessage); i {
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
		file_microService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMessage); i {
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
			RawDescriptor: file_microService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_microService_proto_goTypes,
		DependencyIndexes: file_microService_proto_depIdxs,
		MessageInfos:      file_microService_proto_msgTypes,
	}.Build()
	File_microService_proto = out.File
	file_microService_proto_rawDesc = nil
	file_microService_proto_goTypes = nil
	file_microService_proto_depIdxs = nil
}