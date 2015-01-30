// Code generated by protoc-gen-gogo.
// source: controlmessage.proto
// DO NOT EDIT!

/*
Package control is a generated protocol buffer package.

It is generated from these files:
	controlmessage.proto
	heartbeatrequest.proto
	uuid.proto

It has these top-level messages:
	ControlMessage
*/
package control

import proto "github.com/gogo/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// / Type of the wrapped control.
type ControlMessage_ControlType int32

const (
	ControlMessage_HeartbeatRequest ControlMessage_ControlType = 1
)

var ControlMessage_ControlType_name = map[int32]string{
	1: "HeartbeatRequest",
}
var ControlMessage_ControlType_value = map[string]int32{
	"HeartbeatRequest": 1,
}

func (x ControlMessage_ControlType) Enum() *ControlMessage_ControlType {
	p := new(ControlMessage_ControlType)
	*p = x
	return p
}
func (x ControlMessage_ControlType) String() string {
	return proto.EnumName(ControlMessage_ControlType_name, int32(x))
}
func (x *ControlMessage_ControlType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ControlMessage_ControlType_value, data, "ControlMessage_ControlType")
	if err != nil {
		return err
	}
	*x = ControlMessage_ControlType(value)
	return nil
}

// / ControlMessage wraps a control command and adds metadata.
type ControlMessage struct {
	Origin           *string                     `protobuf:"bytes,1,req,name=origin" json:"origin,omitempty"`
	Identifier       *UUID                       `protobuf:"bytes,2,req,name=identifier" json:"identifier,omitempty"`
	Timestamp        *int64                      `protobuf:"varint,3,req,name=timestamp" json:"timestamp,omitempty"`
	ControlType      *ControlMessage_ControlType `protobuf:"varint,4,req,name=controlType,enum=control.ControlMessage_ControlType" json:"controlType,omitempty"`
	HeartbeatRequest *HeartbeatRequest           `protobuf:"bytes,5,opt,name=heartbeatRequest" json:"heartbeatRequest,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *ControlMessage) Reset()         { *m = ControlMessage{} }
func (m *ControlMessage) String() string { return proto.CompactTextString(m) }
func (*ControlMessage) ProtoMessage()    {}

func (m *ControlMessage) GetOrigin() string {
	if m != nil && m.Origin != nil {
		return *m.Origin
	}
	return ""
}

func (m *ControlMessage) GetIdentifier() *UUID {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *ControlMessage) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *ControlMessage) GetControlType() ControlMessage_ControlType {
	if m != nil && m.ControlType != nil {
		return *m.ControlType
	}
	return ControlMessage_HeartbeatRequest
}

func (m *ControlMessage) GetHeartbeatRequest() *HeartbeatRequest {
	if m != nil {
		return m.HeartbeatRequest
	}
	return nil
}

func init() {
	proto.RegisterEnum("control.ControlMessage_ControlType", ControlMessage_ControlType_name, ControlMessage_ControlType_value)
}
