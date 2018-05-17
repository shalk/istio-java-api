// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/home/graph/v1/homegraph.proto

package graph // import "google.golang.org/genproto/googleapis/home/graph/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _struct "github.com/golang/protobuf/ptypes/struct"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request type for RequestSyncDevices call.
type RequestSyncDevicesRequest struct {
	// Required. Third-party user id issued by agent's third-party identity
	// provider.
	AgentUserId string `protobuf:"bytes,1,opt,name=agent_user_id,json=agentUserId" json:"agent_user_id,omitempty"`
	// Optional. If set, the request will be added to a queue and a response will
	// be returned immediately. The queue allows for de-duplication of
	// simultaneous requests.
	Async                bool     `protobuf:"varint,2,opt,name=async" json:"async,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestSyncDevicesRequest) Reset()         { *m = RequestSyncDevicesRequest{} }
func (m *RequestSyncDevicesRequest) String() string { return proto.CompactTextString(m) }
func (*RequestSyncDevicesRequest) ProtoMessage()    {}
func (*RequestSyncDevicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_homegraph_5bfcd56e20a4e9c3, []int{0}
}
func (m *RequestSyncDevicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestSyncDevicesRequest.Unmarshal(m, b)
}
func (m *RequestSyncDevicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestSyncDevicesRequest.Marshal(b, m, deterministic)
}
func (dst *RequestSyncDevicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestSyncDevicesRequest.Merge(dst, src)
}
func (m *RequestSyncDevicesRequest) XXX_Size() int {
	return xxx_messageInfo_RequestSyncDevicesRequest.Size(m)
}
func (m *RequestSyncDevicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestSyncDevicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RequestSyncDevicesRequest proto.InternalMessageInfo

func (m *RequestSyncDevicesRequest) GetAgentUserId() string {
	if m != nil {
		return m.AgentUserId
	}
	return ""
}

func (m *RequestSyncDevicesRequest) GetAsync() bool {
	if m != nil {
		return m.Async
	}
	return false
}

// Response type for RequestSyncDevices call. Intentionally empty upon success.
// An HTTP response code is returned with more details upon failure.
type RequestSyncDevicesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestSyncDevicesResponse) Reset()         { *m = RequestSyncDevicesResponse{} }
func (m *RequestSyncDevicesResponse) String() string { return proto.CompactTextString(m) }
func (*RequestSyncDevicesResponse) ProtoMessage()    {}
func (*RequestSyncDevicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_homegraph_5bfcd56e20a4e9c3, []int{1}
}
func (m *RequestSyncDevicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestSyncDevicesResponse.Unmarshal(m, b)
}
func (m *RequestSyncDevicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestSyncDevicesResponse.Marshal(b, m, deterministic)
}
func (dst *RequestSyncDevicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestSyncDevicesResponse.Merge(dst, src)
}
func (m *RequestSyncDevicesResponse) XXX_Size() int {
	return xxx_messageInfo_RequestSyncDevicesResponse.Size(m)
}
func (m *RequestSyncDevicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestSyncDevicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestSyncDevicesResponse proto.InternalMessageInfo

// Sample ReportStateAndNotificationRequest, with states and notifications
// defined per device_id (eg: "123" and "456" in the following example):
// {
//   "requestId": "ff36a3cc-ec34-11e6-b1a0-64510650abcf",
//   "agent_user_id": "1234",
//   "payload": {
//     "devices": {
//       "states": {
//         "123": {
//           "on": true
//         },
//         "456": {
//           "on": true,
//           "brightness": 10
//         }
//       },
//       "notifications": {
//         "123": {
//           "ObjectDetected": {
//             "priority": 0,
//             "objects": {
//               "NAMED": ["Alice", "Bob", "Carol", "Eve"]
//             }
//           },
//           "DoorUnlocked": {
//             "priority": 0,
//             "keyUsed": {
//               "keyName": "Wife's key"
//             }
//           }
//         },
//         "456": {
//           "SprinklersOn": {
//             "priority": 0,
//             "timeStarted": "1513792702"
//           }
//         }
//       }
//     }
//   }
// }
// Request type for ReportStateAndNotification call. It may include States,
// Notifications, or both. This request uses globally unique flattened state
// names instead of namespaces based on traits to align with the existing QUERY
// and EXECUTE APIs implemented by 90+ Smart Home partners.
// Next tag: 5.
type ReportStateAndNotificationRequest struct {
	// Request id used for debugging.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	// Unique identifier per event (eg: doorbell press).
	EventId string `protobuf:"bytes,4,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	// Required. Third-party user id.
	AgentUserId string `protobuf:"bytes,2,opt,name=agent_user_id,json=agentUserId" json:"agent_user_id,omitempty"`
	// State of devices to update and notification metadata for devices. For
	// example, if a user turns a light on manually, a State update should be
	// sent so that the information is always the current status of the device.
	// Notifications are independent from the state and its piece of the payload
	// should contain everything necessary to notify the user. Although it may be
	// related to a state change, it does not need to be. For example, if a
	// device can turn on/off and change temperature, the states reported would
	// include both "on" and "70 degrees" but the 3p may choose not to send any
	// notification for that, or to only say that the "the room is heating up",
	// keeping state and notification independent.
	Payload              *StateAndNotificationPayload `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ReportStateAndNotificationRequest) Reset()         { *m = ReportStateAndNotificationRequest{} }
func (m *ReportStateAndNotificationRequest) String() string { return proto.CompactTextString(m) }
func (*ReportStateAndNotificationRequest) ProtoMessage()    {}
func (*ReportStateAndNotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_homegraph_5bfcd56e20a4e9c3, []int{2}
}
func (m *ReportStateAndNotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportStateAndNotificationRequest.Unmarshal(m, b)
}
func (m *ReportStateAndNotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportStateAndNotificationRequest.Marshal(b, m, deterministic)
}
func (dst *ReportStateAndNotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportStateAndNotificationRequest.Merge(dst, src)
}
func (m *ReportStateAndNotificationRequest) XXX_Size() int {
	return xxx_messageInfo_ReportStateAndNotificationRequest.Size(m)
}
func (m *ReportStateAndNotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportStateAndNotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportStateAndNotificationRequest proto.InternalMessageInfo

func (m *ReportStateAndNotificationRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *ReportStateAndNotificationRequest) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *ReportStateAndNotificationRequest) GetAgentUserId() string {
	if m != nil {
		return m.AgentUserId
	}
	return ""
}

func (m *ReportStateAndNotificationRequest) GetPayload() *StateAndNotificationPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Response type for ReportStateAndNotification call.
type ReportStateAndNotificationResponse struct {
	// Request id copied from ReportStateAndNotificationRequest.
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportStateAndNotificationResponse) Reset()         { *m = ReportStateAndNotificationResponse{} }
func (m *ReportStateAndNotificationResponse) String() string { return proto.CompactTextString(m) }
func (*ReportStateAndNotificationResponse) ProtoMessage()    {}
func (*ReportStateAndNotificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_homegraph_5bfcd56e20a4e9c3, []int{3}
}
func (m *ReportStateAndNotificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportStateAndNotificationResponse.Unmarshal(m, b)
}
func (m *ReportStateAndNotificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportStateAndNotificationResponse.Marshal(b, m, deterministic)
}
func (dst *ReportStateAndNotificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportStateAndNotificationResponse.Merge(dst, src)
}
func (m *ReportStateAndNotificationResponse) XXX_Size() int {
	return xxx_messageInfo_ReportStateAndNotificationResponse.Size(m)
}
func (m *ReportStateAndNotificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportStateAndNotificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportStateAndNotificationResponse proto.InternalMessageInfo

func (m *ReportStateAndNotificationResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

// Payload containing the State and Notification information for devices.
type StateAndNotificationPayload struct {
	// The devices for updating State and sending Notifications.
	Devices              *ReportStateAndNotificationDevice `protobuf:"bytes,1,opt,name=devices" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *StateAndNotificationPayload) Reset()         { *m = StateAndNotificationPayload{} }
func (m *StateAndNotificationPayload) String() string { return proto.CompactTextString(m) }
func (*StateAndNotificationPayload) ProtoMessage()    {}
func (*StateAndNotificationPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_homegraph_5bfcd56e20a4e9c3, []int{4}
}
func (m *StateAndNotificationPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateAndNotificationPayload.Unmarshal(m, b)
}
func (m *StateAndNotificationPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateAndNotificationPayload.Marshal(b, m, deterministic)
}
func (dst *StateAndNotificationPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateAndNotificationPayload.Merge(dst, src)
}
func (m *StateAndNotificationPayload) XXX_Size() int {
	return xxx_messageInfo_StateAndNotificationPayload.Size(m)
}
func (m *StateAndNotificationPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_StateAndNotificationPayload.DiscardUnknown(m)
}

var xxx_messageInfo_StateAndNotificationPayload proto.InternalMessageInfo

func (m *StateAndNotificationPayload) GetDevices() *ReportStateAndNotificationDevice {
	if m != nil {
		return m.Devices
	}
	return nil
}

// The States and Notifications specific to a device.
type ReportStateAndNotificationDevice struct {
	// States of devices to update.
	States *_struct.Struct `protobuf:"bytes,1,opt,name=states" json:"states,omitempty"`
	// Notifications metadata for devices.
	Notifications        *_struct.Struct `protobuf:"bytes,2,opt,name=notifications" json:"notifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReportStateAndNotificationDevice) Reset()         { *m = ReportStateAndNotificationDevice{} }
func (m *ReportStateAndNotificationDevice) String() string { return proto.CompactTextString(m) }
func (*ReportStateAndNotificationDevice) ProtoMessage()    {}
func (*ReportStateAndNotificationDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_homegraph_5bfcd56e20a4e9c3, []int{5}
}
func (m *ReportStateAndNotificationDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportStateAndNotificationDevice.Unmarshal(m, b)
}
func (m *ReportStateAndNotificationDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportStateAndNotificationDevice.Marshal(b, m, deterministic)
}
func (dst *ReportStateAndNotificationDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportStateAndNotificationDevice.Merge(dst, src)
}
func (m *ReportStateAndNotificationDevice) XXX_Size() int {
	return xxx_messageInfo_ReportStateAndNotificationDevice.Size(m)
}
func (m *ReportStateAndNotificationDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportStateAndNotificationDevice.DiscardUnknown(m)
}

var xxx_messageInfo_ReportStateAndNotificationDevice proto.InternalMessageInfo

func (m *ReportStateAndNotificationDevice) GetStates() *_struct.Struct {
	if m != nil {
		return m.States
	}
	return nil
}

func (m *ReportStateAndNotificationDevice) GetNotifications() *_struct.Struct {
	if m != nil {
		return m.Notifications
	}
	return nil
}

// Request type for DeleteAgentUser call.
type DeleteAgentUserRequest struct {
	// Request id used for debugging.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	// Required. Third-party user id.
	AgentUserId          string   `protobuf:"bytes,2,opt,name=agent_user_id,json=agentUserId" json:"agent_user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAgentUserRequest) Reset()         { *m = DeleteAgentUserRequest{} }
func (m *DeleteAgentUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAgentUserRequest) ProtoMessage()    {}
func (*DeleteAgentUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_homegraph_5bfcd56e20a4e9c3, []int{6}
}
func (m *DeleteAgentUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAgentUserRequest.Unmarshal(m, b)
}
func (m *DeleteAgentUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAgentUserRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteAgentUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAgentUserRequest.Merge(dst, src)
}
func (m *DeleteAgentUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAgentUserRequest.Size(m)
}
func (m *DeleteAgentUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAgentUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAgentUserRequest proto.InternalMessageInfo

func (m *DeleteAgentUserRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *DeleteAgentUserRequest) GetAgentUserId() string {
	if m != nil {
		return m.AgentUserId
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestSyncDevicesRequest)(nil), "google.home.graph.v1.RequestSyncDevicesRequest")
	proto.RegisterType((*RequestSyncDevicesResponse)(nil), "google.home.graph.v1.RequestSyncDevicesResponse")
	proto.RegisterType((*ReportStateAndNotificationRequest)(nil), "google.home.graph.v1.ReportStateAndNotificationRequest")
	proto.RegisterType((*ReportStateAndNotificationResponse)(nil), "google.home.graph.v1.ReportStateAndNotificationResponse")
	proto.RegisterType((*StateAndNotificationPayload)(nil), "google.home.graph.v1.StateAndNotificationPayload")
	proto.RegisterType((*ReportStateAndNotificationDevice)(nil), "google.home.graph.v1.ReportStateAndNotificationDevice")
	proto.RegisterType((*DeleteAgentUserRequest)(nil), "google.home.graph.v1.DeleteAgentUserRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HomeGraphApiServiceClient is the client API for HomeGraphApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HomeGraphApiServiceClient interface {
	// Requests a Sync call from Google to a 3p partner's home control agent for
	// a user.
	//
	//
	// Third-party user's identity is passed in as agent_user_id.
	// (see [RequestSyncDevicesRequest][google.home.graph.v1.RequestSyncDevicesRequest]) and forwarded back to the agent.
	// Agent is identified by the API key or JWT signed by the partner's service
	// account.
	RequestSyncDevices(ctx context.Context, in *RequestSyncDevicesRequest, opts ...grpc.CallOption) (*RequestSyncDevicesResponse, error)
	// Reports device state and optionally sends device notifications. Called by
	// an agent when the device state of a third-party changes or the agent wants
	// to send a notification about the device.
	// This method updates a predefined set of States for a device, which all
	// devices have (for example a light will have OnOff, Color, Brightness).
	// A new State may not be created and an INVALID_ARGUMENT code will be thrown
	// if so. It also optionally takes in a list of Notifications that may be
	// created, which are associated to this State change.
	//
	// Third-party user's identity is passed in as agent_user_id.
	// Agent is identified by the JWT signed by the partner's service account.
	ReportStateAndNotification(ctx context.Context, in *ReportStateAndNotificationRequest, opts ...grpc.CallOption) (*ReportStateAndNotificationResponse, error)
	// Unlink an agent user from Google. As result, all data related to this user
	// will be deleted.
	//
	// Third-party user's identity is passed in as agent_user_id.
	// Agent is identified by the JWT signed by the partner's service account.
	//
	// Note: Special characters (except "/") in agent_user_id must be URL encoded.
	DeleteAgentUser(ctx context.Context, in *DeleteAgentUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type homeGraphApiServiceClient struct {
	cc *grpc.ClientConn
}

func NewHomeGraphApiServiceClient(cc *grpc.ClientConn) HomeGraphApiServiceClient {
	return &homeGraphApiServiceClient{cc}
}

func (c *homeGraphApiServiceClient) RequestSyncDevices(ctx context.Context, in *RequestSyncDevicesRequest, opts ...grpc.CallOption) (*RequestSyncDevicesResponse, error) {
	out := new(RequestSyncDevicesResponse)
	err := c.cc.Invoke(ctx, "/google.home.graph.v1.HomeGraphApiService/RequestSyncDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeGraphApiServiceClient) ReportStateAndNotification(ctx context.Context, in *ReportStateAndNotificationRequest, opts ...grpc.CallOption) (*ReportStateAndNotificationResponse, error) {
	out := new(ReportStateAndNotificationResponse)
	err := c.cc.Invoke(ctx, "/google.home.graph.v1.HomeGraphApiService/ReportStateAndNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeGraphApiServiceClient) DeleteAgentUser(ctx context.Context, in *DeleteAgentUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.home.graph.v1.HomeGraphApiService/DeleteAgentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HomeGraphApiService service

type HomeGraphApiServiceServer interface {
	// Requests a Sync call from Google to a 3p partner's home control agent for
	// a user.
	//
	//
	// Third-party user's identity is passed in as agent_user_id.
	// (see [RequestSyncDevicesRequest][google.home.graph.v1.RequestSyncDevicesRequest]) and forwarded back to the agent.
	// Agent is identified by the API key or JWT signed by the partner's service
	// account.
	RequestSyncDevices(context.Context, *RequestSyncDevicesRequest) (*RequestSyncDevicesResponse, error)
	// Reports device state and optionally sends device notifications. Called by
	// an agent when the device state of a third-party changes or the agent wants
	// to send a notification about the device.
	// This method updates a predefined set of States for a device, which all
	// devices have (for example a light will have OnOff, Color, Brightness).
	// A new State may not be created and an INVALID_ARGUMENT code will be thrown
	// if so. It also optionally takes in a list of Notifications that may be
	// created, which are associated to this State change.
	//
	// Third-party user's identity is passed in as agent_user_id.
	// Agent is identified by the JWT signed by the partner's service account.
	ReportStateAndNotification(context.Context, *ReportStateAndNotificationRequest) (*ReportStateAndNotificationResponse, error)
	// Unlink an agent user from Google. As result, all data related to this user
	// will be deleted.
	//
	// Third-party user's identity is passed in as agent_user_id.
	// Agent is identified by the JWT signed by the partner's service account.
	//
	// Note: Special characters (except "/") in agent_user_id must be URL encoded.
	DeleteAgentUser(context.Context, *DeleteAgentUserRequest) (*empty.Empty, error)
}

func RegisterHomeGraphApiServiceServer(s *grpc.Server, srv HomeGraphApiServiceServer) {
	s.RegisterService(&_HomeGraphApiService_serviceDesc, srv)
}

func _HomeGraphApiService_RequestSyncDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSyncDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeGraphApiServiceServer).RequestSyncDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.home.graph.v1.HomeGraphApiService/RequestSyncDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeGraphApiServiceServer).RequestSyncDevices(ctx, req.(*RequestSyncDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeGraphApiService_ReportStateAndNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportStateAndNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeGraphApiServiceServer).ReportStateAndNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.home.graph.v1.HomeGraphApiService/ReportStateAndNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeGraphApiServiceServer).ReportStateAndNotification(ctx, req.(*ReportStateAndNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeGraphApiService_DeleteAgentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeGraphApiServiceServer).DeleteAgentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.home.graph.v1.HomeGraphApiService/DeleteAgentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeGraphApiServiceServer).DeleteAgentUser(ctx, req.(*DeleteAgentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HomeGraphApiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.home.graph.v1.HomeGraphApiService",
	HandlerType: (*HomeGraphApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestSyncDevices",
			Handler:    _HomeGraphApiService_RequestSyncDevices_Handler,
		},
		{
			MethodName: "ReportStateAndNotification",
			Handler:    _HomeGraphApiService_ReportStateAndNotification_Handler,
		},
		{
			MethodName: "DeleteAgentUser",
			Handler:    _HomeGraphApiService_DeleteAgentUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/home/graph/v1/homegraph.proto",
}

func init() {
	proto.RegisterFile("google/home/graph/v1/homegraph.proto", fileDescriptor_homegraph_5bfcd56e20a4e9c3)
}

var fileDescriptor_homegraph_5bfcd56e20a4e9c3 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0x66, 0x5b, 0x6d, 0xda, 0x13, 0x8a, 0x30, 0x96, 0x76, 0x9b, 0x46, 0x48, 0x47, 0x91, 0x18,
	0x65, 0xd7, 0x44, 0xf0, 0xa7, 0xd2, 0x8b, 0xd4, 0x8a, 0x06, 0x41, 0xc2, 0x86, 0xde, 0xe8, 0x45,
	0x99, 0xee, 0x9e, 0x6e, 0x17, 0x92, 0x99, 0x75, 0x67, 0x12, 0x08, 0xe2, 0x4d, 0x1f, 0x41, 0x2f,
	0x7d, 0x1b, 0x1f, 0xc0, 0x1b, 0x5f, 0xc1, 0xd7, 0x10, 0x64, 0x67, 0x76, 0xad, 0x69, 0x36, 0x69,
	0x7a, 0x37, 0x67, 0xce, 0xdf, 0x77, 0xce, 0xf7, 0xcd, 0xc0, 0xbd, 0x50, 0x88, 0xb0, 0x8f, 0xee,
	0x99, 0x18, 0xa0, 0x1b, 0x26, 0x2c, 0x3e, 0x73, 0x47, 0x4d, 0x6d, 0x69, 0xc3, 0x89, 0x13, 0xa1,
	0x04, 0xd9, 0x30, 0x51, 0x4e, 0x7a, 0xef, 0x18, 0xc7, 0xa8, 0x59, 0xa9, 0x66, 0xb9, 0x2c, 0x8e,
	0x5c, 0xc6, 0xb9, 0x50, 0x4c, 0x45, 0x82, 0x4b, 0x93, 0x53, 0xd9, 0xc9, 0xbc, 0xda, 0x3a, 0x19,
	0x9e, 0xba, 0x38, 0x88, 0xd5, 0x38, 0x73, 0x56, 0x2f, 0x3b, 0xa5, 0x4a, 0x86, 0xbe, 0x32, 0x5e,
	0x7a, 0x04, 0xdb, 0x1e, 0x7e, 0x1a, 0xa2, 0x54, 0xbd, 0x31, 0xf7, 0x0f, 0x71, 0x14, 0xf9, 0x28,
	0xb3, 0x1b, 0x42, 0x61, 0x9d, 0x85, 0xc8, 0xd5, 0xf1, 0x50, 0x62, 0x72, 0x1c, 0x05, 0xb6, 0x55,
	0xb3, 0xea, 0x6b, 0x5e, 0x59, 0x5f, 0x1e, 0x49, 0x4c, 0x3a, 0x01, 0xd9, 0x80, 0x9b, 0x4c, 0x8e,
	0xb9, 0x6f, 0x2f, 0xd5, 0xac, 0xfa, 0xaa, 0x67, 0x0c, 0x5a, 0x85, 0x4a, 0x51, 0x59, 0x19, 0x0b,
	0x2e, 0x91, 0xfe, 0xb4, 0x60, 0xd7, 0xc3, 0x58, 0x24, 0xaa, 0xa7, 0x98, 0xc2, 0x36, 0x0f, 0xde,
	0x0b, 0x15, 0x9d, 0x46, 0xbe, 0x9e, 0x2a, 0xef, 0x7e, 0x07, 0x20, 0x31, 0xc7, 0x8b, 0xd6, 0x6b,
	0xd9, 0x4d, 0x27, 0x20, 0xdb, 0xb0, 0x8a, 0xa3, 0x14, 0x5c, 0x14, 0xd8, 0x37, 0xb4, 0xb3, 0xa4,
	0xed, 0x4e, 0x30, 0x8d, 0x7b, 0x69, 0x1a, 0xf7, 0x3b, 0x28, 0xc5, 0x6c, 0xdc, 0x17, 0x2c, 0xb0,
	0x97, 0x6b, 0x56, 0xbd, 0xdc, 0x6a, 0x3a, 0x45, 0x9b, 0x77, 0x8a, 0x10, 0x76, 0x4d, 0xa2, 0x97,
	0x57, 0xa0, 0xaf, 0x80, 0xce, 0x9b, 0xc7, 0x8c, 0x7d, 0xc5, 0x40, 0x54, 0xc0, 0xce, 0x9c, 0x66,
	0xa4, 0x0b, 0xa5, 0xc0, 0xec, 0x51, 0xa7, 0x96, 0x5b, 0x4f, 0x8b, 0x01, 0xcf, 0x06, 0x62, 0x68,
	0xf0, 0xf2, 0x32, 0xf4, 0xab, 0x05, 0xb5, 0xab, 0xa2, 0x89, 0x0b, 0x2b, 0x32, 0xf5, 0xe6, 0x5d,
	0xb7, 0xf2, 0xae, 0xb9, 0x9e, 0x9c, 0x9e, 0xd6, 0x93, 0x97, 0x85, 0x91, 0x7d, 0x58, 0xe7, 0xff,
	0x95, 0x91, 0x7a, 0xf9, 0x73, 0xf2, 0x26, 0xa3, 0xe9, 0x47, 0xd8, 0x3c, 0xc4, 0x3e, 0x2a, 0x6c,
	0xe7, 0x64, 0x2d, 0xa8, 0x87, 0x05, 0x48, 0x6f, 0xfd, 0x59, 0x86, 0xdb, 0x6f, 0xc5, 0x00, 0xdf,
	0xa4, 0xcb, 0x6a, 0xc7, 0x51, 0x0f, 0x13, 0x3d, 0xe4, 0x77, 0x0b, 0xc8, 0xb4, 0x5e, 0x89, 0x3b,
	0x6b, 0xc3, 0x33, 0x1e, 0x4c, 0xe5, 0xf1, 0xe2, 0x09, 0xd9, 0x53, 0xa0, 0xe7, 0xbf, 0x7e, 0x7f,
	0x5b, 0xaa, 0xd2, 0xad, 0xf4, 0x2f, 0xc8, 0x88, 0xd9, 0x4b, 0x2e, 0xe2, 0xf7, 0xac, 0x06, 0xf9,
	0x61, 0xa5, 0xaf, 0x69, 0x16, 0x4f, 0xe4, 0xd9, 0x75, 0x75, 0x90, 0xa3, 0x7d, 0x7e, 0xfd, 0xc4,
	0x0c, 0x75, 0x53, 0xa3, 0x7e, 0x48, 0xef, 0x4f, 0xa2, 0x9e, 0x95, 0x97, 0x0e, 0x71, 0x6e, 0xc1,
	0xad, 0x4b, 0xc4, 0x92, 0x47, 0xc5, 0x00, 0x8a, 0xf9, 0xaf, 0x6c, 0x4e, 0x29, 0xe8, 0x75, 0xfa,
	0xcd, 0xd1, 0x07, 0x1a, 0xcc, 0xdd, 0xc6, 0x6e, 0x0a, 0xe6, 0xf3, 0x84, 0x06, 0xf6, 0xff, 0x91,
	0x2f, 0xdd, 0x46, 0xe3, 0xcb, 0x81, 0x00, 0xdb, 0x17, 0x83, 0xc2, 0xae, 0x07, 0x76, 0x81, 0x30,
	0xba, 0x69, 0xa7, 0x0f, 0x2f, 0xb2, 0xf8, 0x50, 0xf4, 0x19, 0x0f, 0x1d, 0x91, 0x84, 0x6e, 0x88,
	0x5c, 0xa3, 0x70, 0x8d, 0x8b, 0xc5, 0x91, 0x9c, 0xfc, 0xd7, 0x5f, 0xea, 0xc3, 0xc9, 0x8a, 0x8e,
	0x7a, 0xf2, 0x37, 0x00, 0x00, 0xff, 0xff, 0x92, 0xb2, 0x5b, 0xe2, 0xfc, 0x05, 0x00, 0x00,
}
