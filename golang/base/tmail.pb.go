// Code generated by protoc-gen-go.
// source: tmail.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	tmail.proto

It has these top-level messages:
	SmtpdResponse
	SmtpdNewClientMsg
*/
package main

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SmtpdResponse struct {
	SmtpCode         *int32   `protobuf:"varint,1,req,name=smtp_code" json:"smtp_code,omitempty"`
	SmtpMsg          *string  `protobuf:"bytes,2,req,name=smtp_msg" json:"smtp_msg,omitempty"`
	CloseConnection  *bool    `protobuf:"varint,16,opt,name=close_connection" json:"close_connection,omitempty"`
	DataLink         *string  `protobuf:"bytes,17,opt,name=data_link" json:"data_link,omitempty"`
	ExtraHeaders     []string `protobuf:"bytes,18,rep,name=extra_headers" json:"extra_headers,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SmtpdResponse) Reset()         { *m = SmtpdResponse{} }
func (m *SmtpdResponse) String() string { return proto.CompactTextString(m) }
func (*SmtpdResponse) ProtoMessage()    {}

func (m *SmtpdResponse) GetSmtpCode() int32 {
	if m != nil && m.SmtpCode != nil {
		return *m.SmtpCode
	}
	return 0
}

func (m *SmtpdResponse) GetSmtpMsg() string {
	if m != nil && m.SmtpMsg != nil {
		return *m.SmtpMsg
	}
	return ""
}

func (m *SmtpdResponse) GetCloseConnection() bool {
	if m != nil && m.CloseConnection != nil {
		return *m.CloseConnection
	}
	return false
}

func (m *SmtpdResponse) GetDataLink() string {
	if m != nil && m.DataLink != nil {
		return *m.DataLink
	}
	return ""
}

func (m *SmtpdResponse) GetExtraHeaders() []string {
	if m != nil {
		return m.ExtraHeaders
	}
	return nil
}

type SmtpdNewClientMsg struct {
	SessionId        *string `protobuf:"bytes,1,req,name=session_id" json:"session_id,omitempty"`
	RemoteIp         *string `protobuf:"bytes,2,req,name=remote_ip" json:"remote_ip,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SmtpdNewClientMsg) Reset()         { *m = SmtpdNewClientMsg{} }
func (m *SmtpdNewClientMsg) String() string { return proto.CompactTextString(m) }
func (*SmtpdNewClientMsg) ProtoMessage()    {}

func (m *SmtpdNewClientMsg) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *SmtpdNewClientMsg) GetRemoteIp() string {
	if m != nil && m.RemoteIp != nil {
		return *m.RemoteIp
	}
	return ""
}

func init() {
}
