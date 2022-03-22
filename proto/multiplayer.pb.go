// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/multiplayer.proto

package proto

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

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	X  int64  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y  int64  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_multiplayer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_proto_multiplayer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_proto_multiplayer_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Player) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type NewPlayerAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewPlayerAction) Reset() {
	*x = NewPlayerAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_multiplayer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPlayerAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPlayerAction) ProtoMessage() {}

func (x *NewPlayerAction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_multiplayer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPlayerAction.ProtoReflect.Descriptor instead.
func (*NewPlayerAction) Descriptor() ([]byte, []int) {
	return file_proto_multiplayer_proto_rawDescGZIP(), []int{1}
}

func (x *NewPlayerAction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MoveAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeftPlayer  *Player `protobuf:"bytes,1,opt,name=leftPlayer,proto3" json:"leftPlayer,omitempty"`
	RightPlayer *Player `protobuf:"bytes,2,opt,name=rightPlayer,proto3" json:"rightPlayer,omitempty"`
}

func (x *MoveAction) Reset() {
	*x = MoveAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_multiplayer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveAction) ProtoMessage() {}

func (x *MoveAction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_multiplayer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveAction.ProtoReflect.Descriptor instead.
func (*MoveAction) Descriptor() ([]byte, []int) {
	return file_proto_multiplayer_proto_rawDescGZIP(), []int{2}
}

func (x *MoveAction) GetLeftPlayer() *Player {
	if x != nil {
		return x.LeftPlayer
	}
	return nil
}

func (x *MoveAction) GetRightPlayer() *Player {
	if x != nil {
		return x.RightPlayer
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//	*Request_NewPlayer
	//	*Request_MoveAction
	Action isRequest_Action `protobuf_oneof:"action"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_multiplayer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_multiplayer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_multiplayer_proto_rawDescGZIP(), []int{3}
}

func (m *Request) GetAction() isRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *Request) GetNewPlayer() *NewPlayerAction {
	if x, ok := x.GetAction().(*Request_NewPlayer); ok {
		return x.NewPlayer
	}
	return nil
}

func (x *Request) GetMoveAction() *MoveAction {
	if x, ok := x.GetAction().(*Request_MoveAction); ok {
		return x.MoveAction
	}
	return nil
}

type isRequest_Action interface {
	isRequest_Action()
}

type Request_NewPlayer struct {
	NewPlayer *NewPlayerAction `protobuf:"bytes,1,opt,name=newPlayer,proto3,oneof"`
}

type Request_MoveAction struct {
	MoveAction *MoveAction `protobuf:"bytes,2,opt,name=moveAction,proto3,oneof"`
}

func (*Request_NewPlayer) isRequest_Action() {}

func (*Request_MoveAction) isRequest_Action() {}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//	*Response_NewPlayer
	//	*Response_MoveAction
	Action isResponse_Action `protobuf_oneof:"action"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_multiplayer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_multiplayer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_multiplayer_proto_rawDescGZIP(), []int{4}
}

func (m *Response) GetAction() isResponse_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *Response) GetNewPlayer() *NewPlayerAction {
	if x, ok := x.GetAction().(*Response_NewPlayer); ok {
		return x.NewPlayer
	}
	return nil
}

func (x *Response) GetMoveAction() *MoveAction {
	if x, ok := x.GetAction().(*Response_MoveAction); ok {
		return x.MoveAction
	}
	return nil
}

type isResponse_Action interface {
	isResponse_Action()
}

type Response_NewPlayer struct {
	NewPlayer *NewPlayerAction `protobuf:"bytes,1,opt,name=newPlayer,proto3,oneof"`
}

type Response_MoveAction struct {
	MoveAction *MoveAction `protobuf:"bytes,2,opt,name=moveAction,proto3,oneof"`
}

func (*Response_NewPlayer) isResponse_Action() {}

func (*Response_MoveAction) isResponse_Action() {}

var File_proto_multiplayer_proto protoreflect.FileDescriptor

var file_proto_multiplayer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x06, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x79, 0x22,
	0x21, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x60, 0x0a, 0x0a, 0x4d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x27, 0x0a, 0x0a, 0x6c, 0x65, 0x66, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x0a, 0x6c,
	0x65, 0x66, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0b, 0x72, 0x69, 0x67,
	0x68, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x0b, 0x72, 0x69, 0x67, 0x68, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x22, 0x74, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x2d, 0x0a, 0x0a, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4e, 0x65, 0x77, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x6e,
	0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x0a, 0x6d, 0x6f, 0x76, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d,
	0x6f, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x6d, 0x6f, 0x76,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0x35, 0x0a, 0x0b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_multiplayer_proto_rawDescOnce sync.Once
	file_proto_multiplayer_proto_rawDescData = file_proto_multiplayer_proto_rawDesc
)

func file_proto_multiplayer_proto_rawDescGZIP() []byte {
	file_proto_multiplayer_proto_rawDescOnce.Do(func() {
		file_proto_multiplayer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_multiplayer_proto_rawDescData)
	})
	return file_proto_multiplayer_proto_rawDescData
}

var file_proto_multiplayer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_multiplayer_proto_goTypes = []interface{}{
	(*Player)(nil),          // 0: Player
	(*NewPlayerAction)(nil), // 1: NewPlayerAction
	(*MoveAction)(nil),      // 2: MoveAction
	(*Request)(nil),         // 3: Request
	(*Response)(nil),        // 4: Response
}
var file_proto_multiplayer_proto_depIdxs = []int32{
	0, // 0: MoveAction.leftPlayer:type_name -> Player
	0, // 1: MoveAction.rightPlayer:type_name -> Player
	1, // 2: Request.newPlayer:type_name -> NewPlayerAction
	2, // 3: Request.moveAction:type_name -> MoveAction
	1, // 4: Response.newPlayer:type_name -> NewPlayerAction
	2, // 5: Response.moveAction:type_name -> MoveAction
	3, // 6: Multiplayer.GameProcess:input_type -> Request
	4, // 7: Multiplayer.GameProcess:output_type -> Response
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_multiplayer_proto_init() }
func file_proto_multiplayer_proto_init() {
	if File_proto_multiplayer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_multiplayer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_proto_multiplayer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPlayerAction); i {
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
		file_proto_multiplayer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveAction); i {
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
		file_proto_multiplayer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_multiplayer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
	file_proto_multiplayer_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Request_NewPlayer)(nil),
		(*Request_MoveAction)(nil),
	}
	file_proto_multiplayer_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Response_NewPlayer)(nil),
		(*Response_MoveAction)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_multiplayer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_multiplayer_proto_goTypes,
		DependencyIndexes: file_proto_multiplayer_proto_depIdxs,
		MessageInfos:      file_proto_multiplayer_proto_msgTypes,
	}.Build()
	File_proto_multiplayer_proto = out.File
	file_proto_multiplayer_proto_rawDesc = nil
	file_proto_multiplayer_proto_goTypes = nil
	file_proto_multiplayer_proto_depIdxs = nil
}