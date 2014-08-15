// Code generated by protoc-gen-gogo.
// source: http.proto
// DO NOT EDIT!

package events

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type PeerType int32

const (
	PeerType_Client PeerType = 1
	PeerType_Server PeerType = 2
)

var PeerType_name = map[int32]string{
	1: "Client",
	2: "Server",
}
var PeerType_value = map[string]int32{
	"Client": 1,
	"Server": 2,
}

func (x PeerType) Enum() *PeerType {
	p := new(PeerType)
	*p = x
	return p
}
func (x PeerType) String() string {
	return proto.EnumName(PeerType_name, int32(x))
}
func (x *PeerType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PeerType_value, data, "PeerType")
	if err != nil {
		return err
	}
	*x = PeerType(value)
	return nil
}

type Method int32

const (
	Method_GET    Method = 1
	Method_POST   Method = 2
	Method_PUT    Method = 3
	Method_DELETE Method = 4
	Method_HEAD   Method = 5
)

var Method_name = map[int32]string{
	1: "GET",
	2: "POST",
	3: "PUT",
	4: "DELETE",
	5: "HEAD",
}
var Method_value = map[string]int32{
	"GET":    1,
	"POST":   2,
	"PUT":    3,
	"DELETE": 4,
	"HEAD":   5,
}

func (x Method) Enum() *Method {
	p := new(Method)
	*p = x
	return p
}
func (x Method) String() string {
	return proto.EnumName(Method_name, int32(x))
}
func (x *Method) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Method_value, data, "Method")
	if err != nil {
		return err
	}
	*x = Method(value)
	return nil
}

type UUID struct {
	Low              *uint64 `protobuf:"varint,1,req,name=low" json:"low,omitempty"`
	High             *uint64 `protobuf:"varint,2,req,name=high" json:"high,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UUID) Reset()         { *m = UUID{} }
func (m *UUID) String() string { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()    {}

func (m *UUID) GetLow() uint64 {
	if m != nil && m.Low != nil {
		return *m.Low
	}
	return 0
}

func (m *UUID) GetHigh() uint64 {
	if m != nil && m.High != nil {
		return *m.High
	}
	return 0
}

type HttpStart struct {
	Timestamp        *int64    `protobuf:"varint,1,req,name=timestamp" json:"timestamp,omitempty"`
	RequestId        *UUID     `protobuf:"bytes,2,req,name=requestId" json:"requestId,omitempty"`
	PeerType         *PeerType `protobuf:"varint,3,req,name=peerType,enum=events.PeerType" json:"peerType,omitempty"`
	Method           *Method   `protobuf:"varint,4,req,name=method,enum=events.Method" json:"method,omitempty"`
	Uri              *string   `protobuf:"bytes,5,req,name=uri" json:"uri,omitempty"`
	RemoteAddress    *string   `protobuf:"bytes,6,req,name=remoteAddress" json:"remoteAddress,omitempty"`
	UserAgent        *string   `protobuf:"bytes,7,req,name=userAgent" json:"userAgent,omitempty"`
	ParentRequestId  *UUID     `protobuf:"bytes,8,opt,name=parentRequestId" json:"parentRequestId,omitempty"`
	ApplicationId    *UUID     `protobuf:"bytes,9,opt,name=applicationId" json:"applicationId,omitempty"`
	InstanceIndex    *int32    `protobuf:"varint,10,opt,name=instanceIndex" json:"instanceIndex,omitempty"`
	InstanceId       *string   `protobuf:"bytes,11,opt,name=instanceId" json:"instanceId,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *HttpStart) Reset()         { *m = HttpStart{} }
func (m *HttpStart) String() string { return proto.CompactTextString(m) }
func (*HttpStart) ProtoMessage()    {}

func (m *HttpStart) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *HttpStart) GetRequestId() *UUID {
	if m != nil {
		return m.RequestId
	}
	return nil
}

func (m *HttpStart) GetPeerType() PeerType {
	if m != nil && m.PeerType != nil {
		return *m.PeerType
	}
	return PeerType_Client
}

func (m *HttpStart) GetMethod() Method {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return Method_GET
}

func (m *HttpStart) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *HttpStart) GetRemoteAddress() string {
	if m != nil && m.RemoteAddress != nil {
		return *m.RemoteAddress
	}
	return ""
}

func (m *HttpStart) GetUserAgent() string {
	if m != nil && m.UserAgent != nil {
		return *m.UserAgent
	}
	return ""
}

func (m *HttpStart) GetParentRequestId() *UUID {
	if m != nil {
		return m.ParentRequestId
	}
	return nil
}

func (m *HttpStart) GetApplicationId() *UUID {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *HttpStart) GetInstanceIndex() int32 {
	if m != nil && m.InstanceIndex != nil {
		return *m.InstanceIndex
	}
	return 0
}

func (m *HttpStart) GetInstanceId() string {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return ""
}

type HttpStop struct {
	Timestamp        *int64    `protobuf:"varint,1,req,name=timestamp" json:"timestamp,omitempty"`
	Uri              *string   `protobuf:"bytes,2,req,name=uri" json:"uri,omitempty"`
	RequestId        *UUID     `protobuf:"bytes,3,req,name=requestId" json:"requestId,omitempty"`
	PeerType         *PeerType `protobuf:"varint,4,req,name=peerType,enum=events.PeerType" json:"peerType,omitempty"`
	StatusCode       *int32    `protobuf:"varint,5,req,name=statusCode" json:"statusCode,omitempty"`
	ContentLength    *int64    `protobuf:"varint,6,req,name=contentLength" json:"contentLength,omitempty"`
	ApplicationId    *UUID     `protobuf:"bytes,7,opt,name=applicationId" json:"applicationId,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *HttpStop) Reset()         { *m = HttpStop{} }
func (m *HttpStop) String() string { return proto.CompactTextString(m) }
func (*HttpStop) ProtoMessage()    {}

func (m *HttpStop) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *HttpStop) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *HttpStop) GetRequestId() *UUID {
	if m != nil {
		return m.RequestId
	}
	return nil
}

func (m *HttpStop) GetPeerType() PeerType {
	if m != nil && m.PeerType != nil {
		return *m.PeerType
	}
	return PeerType_Client
}

func (m *HttpStop) GetStatusCode() int32 {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return 0
}

func (m *HttpStop) GetContentLength() int64 {
	if m != nil && m.ContentLength != nil {
		return *m.ContentLength
	}
	return 0
}

func (m *HttpStop) GetApplicationId() *UUID {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

type HttpStartStop struct {
	StartTimestamp   *int64    `protobuf:"varint,1,req,name=startTimestamp" json:"startTimestamp,omitempty"`
	StopTimestamp    *int64    `protobuf:"varint,2,req,name=stopTimestamp" json:"stopTimestamp,omitempty"`
	RequestId        *UUID     `protobuf:"bytes,3,req,name=requestId" json:"requestId,omitempty"`
	PeerType         *PeerType `protobuf:"varint,4,req,name=peerType,enum=events.PeerType" json:"peerType,omitempty"`
	Method           *Method   `protobuf:"varint,5,req,name=method,enum=events.Method" json:"method,omitempty"`
	Uri              *string   `protobuf:"bytes,6,req,name=uri" json:"uri,omitempty"`
	RemoteAddress    *string   `protobuf:"bytes,7,req,name=remoteAddress" json:"remoteAddress,omitempty"`
	UserAgent        *string   `protobuf:"bytes,8,req,name=userAgent" json:"userAgent,omitempty"`
	StatusCode       *int32    `protobuf:"varint,9,req,name=statusCode" json:"statusCode,omitempty"`
	ContentLength    *int64    `protobuf:"varint,10,req,name=contentLength" json:"contentLength,omitempty"`
	ParentRequestId  *UUID     `protobuf:"bytes,11,opt,name=parentRequestId" json:"parentRequestId,omitempty"`
	ApplicationId    *UUID     `protobuf:"bytes,12,opt,name=applicationId" json:"applicationId,omitempty"`
	InstanceIndex    *int32    `protobuf:"varint,13,opt,name=instanceIndex" json:"instanceIndex,omitempty"`
	InstanceId       *string   `protobuf:"bytes,14,opt,name=instanceId" json:"instanceId,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *HttpStartStop) Reset()         { *m = HttpStartStop{} }
func (m *HttpStartStop) String() string { return proto.CompactTextString(m) }
func (*HttpStartStop) ProtoMessage()    {}

func (m *HttpStartStop) GetStartTimestamp() int64 {
	if m != nil && m.StartTimestamp != nil {
		return *m.StartTimestamp
	}
	return 0
}

func (m *HttpStartStop) GetStopTimestamp() int64 {
	if m != nil && m.StopTimestamp != nil {
		return *m.StopTimestamp
	}
	return 0
}

func (m *HttpStartStop) GetRequestId() *UUID {
	if m != nil {
		return m.RequestId
	}
	return nil
}

func (m *HttpStartStop) GetPeerType() PeerType {
	if m != nil && m.PeerType != nil {
		return *m.PeerType
	}
	return PeerType_Client
}

func (m *HttpStartStop) GetMethod() Method {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return Method_GET
}

func (m *HttpStartStop) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *HttpStartStop) GetRemoteAddress() string {
	if m != nil && m.RemoteAddress != nil {
		return *m.RemoteAddress
	}
	return ""
}

func (m *HttpStartStop) GetUserAgent() string {
	if m != nil && m.UserAgent != nil {
		return *m.UserAgent
	}
	return ""
}

func (m *HttpStartStop) GetStatusCode() int32 {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return 0
}

func (m *HttpStartStop) GetContentLength() int64 {
	if m != nil && m.ContentLength != nil {
		return *m.ContentLength
	}
	return 0
}

func (m *HttpStartStop) GetParentRequestId() *UUID {
	if m != nil {
		return m.ParentRequestId
	}
	return nil
}

func (m *HttpStartStop) GetApplicationId() *UUID {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *HttpStartStop) GetInstanceIndex() int32 {
	if m != nil && m.InstanceIndex != nil {
		return *m.InstanceIndex
	}
	return 0
}

func (m *HttpStartStop) GetInstanceId() string {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return ""
}

func init() {
	proto.RegisterEnum("events.PeerType", PeerType_name, PeerType_value)
	proto.RegisterEnum("events.Method", Method_name, Method_value)
}
