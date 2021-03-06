// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: internal/portdomain/proto/port.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type GetPortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortId int64 `protobuf:"varint,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
}

func (x *GetPortRequest) Reset() {
	*x = GetPortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_portdomain_proto_port_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPortRequest) ProtoMessage() {}

func (x *GetPortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_portdomain_proto_port_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPortRequest.ProtoReflect.Descriptor instead.
func (*GetPortRequest) Descriptor() ([]byte, []int) {
	return file_internal_portdomain_proto_port_proto_rawDescGZIP(), []int{0}
}

func (x *GetPortRequest) GetPortId() int64 {
	if x != nil {
		return x.PortId
	}
	return 0
}

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IdStr       string    `protobuf:"bytes,2,opt,name=id_str,json=idStr,proto3" json:"id_str,omitempty"`
	Name        string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	City        string    `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Country     string    `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Coordinates []float64 `protobuf:"fixed64,6,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"` // long, lat
	Provice     string    `protobuf:"bytes,7,opt,name=provice,proto3" json:"provice,omitempty"`
	Timezone    string    `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Code        string    `protobuf:"bytes,9,opt,name=code,proto3" json:"code,omitempty"`
	Regions     []string  `protobuf:"bytes,10,rep,name=regions,proto3" json:"regions,omitempty"`
	Unlocs      []string  `protobuf:"bytes,11,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Alias       []string  `protobuf:"bytes,12,rep,name=alias,proto3" json:"alias,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_portdomain_proto_port_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_internal_portdomain_proto_port_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_internal_portdomain_proto_port_proto_rawDescGZIP(), []int{1}
}

func (x *Port) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Port) GetIdStr() string {
	if x != nil {
		return x.IdStr
	}
	return ""
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Port) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Port) GetCoordinates() []float64 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Port) GetProvice() string {
	if x != nil {
		return x.Provice
	}
	return ""
}

func (x *Port) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Port) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Port) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Port) GetUnlocs() []string {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

func (x *Port) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

var File_internal_portdomain_proto_port_proto protoreflect.FileDescriptor

var file_internal_portdomain_proto_port_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0xa3, 0x02, 0x0a, 0x04, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x64, 0x53, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x32, 0x7b, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x34,
	0x0a, 0x06, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x28, 0x01, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x1a, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x1b, 0x5a,
	0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_portdomain_proto_port_proto_rawDescOnce sync.Once
	file_internal_portdomain_proto_port_proto_rawDescData = file_internal_portdomain_proto_port_proto_rawDesc
)

func file_internal_portdomain_proto_port_proto_rawDescGZIP() []byte {
	file_internal_portdomain_proto_port_proto_rawDescOnce.Do(func() {
		file_internal_portdomain_proto_port_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_portdomain_proto_port_proto_rawDescData)
	})
	return file_internal_portdomain_proto_port_proto_rawDescData
}

var file_internal_portdomain_proto_port_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_portdomain_proto_port_proto_goTypes = []interface{}{
	(*GetPortRequest)(nil), // 0: portdomain.GetPortRequest
	(*Port)(nil),           // 1: portdomain.Port
	(*empty.Empty)(nil),    // 2: google.protobuf.Empty
}
var file_internal_portdomain_proto_port_proto_depIdxs = []int32{
	1, // 0: portdomain.PortDomain.Upsert:input_type -> portdomain.Port
	0, // 1: portdomain.PortDomain.GetPort:input_type -> portdomain.GetPortRequest
	2, // 2: portdomain.PortDomain.Upsert:output_type -> google.protobuf.Empty
	1, // 3: portdomain.PortDomain.GetPort:output_type -> portdomain.Port
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_portdomain_proto_port_proto_init() }
func file_internal_portdomain_proto_port_proto_init() {
	if File_internal_portdomain_proto_port_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_portdomain_proto_port_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPortRequest); i {
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
		file_internal_portdomain_proto_port_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
			RawDescriptor: file_internal_portdomain_proto_port_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_portdomain_proto_port_proto_goTypes,
		DependencyIndexes: file_internal_portdomain_proto_port_proto_depIdxs,
		MessageInfos:      file_internal_portdomain_proto_port_proto_msgTypes,
	}.Build()
	File_internal_portdomain_proto_port_proto = out.File
	file_internal_portdomain_proto_port_proto_rawDesc = nil
	file_internal_portdomain_proto_port_proto_goTypes = nil
	file_internal_portdomain_proto_port_proto_depIdxs = nil
}
