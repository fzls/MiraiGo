// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/qweb/protocol.proto

package qweb

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type QWebReq struct {
	Seq         proto.Option[int64]  `protobuf:"varint,1,opt"`
	Qua         proto.Option[string] `protobuf:"bytes,2,opt"`
	DeviceInfo  proto.Option[string] `protobuf:"bytes,3,opt"`
	BusiBuff    []byte               `protobuf:"bytes,4,opt"`
	TraceId     proto.Option[string] `protobuf:"bytes,5,opt"`
	Module      proto.Option[string] `protobuf:"bytes,6,opt"`
	Cmdname     proto.Option[string] `protobuf:"bytes,7,opt"`
	LoginSig    *StAuthInfo          `protobuf:"bytes,8,opt"`
	Crypto      *StEncryption        `protobuf:"bytes,9,opt"`
	Extinfo     []*COMMEntry         `protobuf:"bytes,10,rep"`
	ContentType proto.Option[uint32] `protobuf:"varint,11,opt"`
}

type QWebRsp struct {
	Seq      proto.Option[int64]  `protobuf:"varint,1,opt"`
	RetCode  proto.Option[int64]  `protobuf:"varint,2,opt"`
	ErrMsg   proto.Option[string] `protobuf:"bytes,3,opt"`
	BusiBuff []byte               `protobuf:"bytes,4,opt"`
	Traceid  proto.Option[string] `protobuf:"bytes,5,opt"`
}

type StAuthInfo struct {
	Uin        proto.Option[string] `protobuf:"bytes,1,opt"`
	Sig        []byte               `protobuf:"bytes,2,opt"`
	Platform   proto.Option[string] `protobuf:"bytes,3,opt"`
	Type       proto.Option[uint32] `protobuf:"varint,4,opt"`
	Appid      proto.Option[string] `protobuf:"bytes,5,opt"`
	Openid     proto.Option[string] `protobuf:"bytes,6,opt"`
	Sessionkey []byte               `protobuf:"bytes,7,opt"`
	Extinfo    []*COMMEntry         `protobuf:"bytes,8,rep"`
}

type StEncryption struct {
	Method proto.Option[uint32] `protobuf:"varint,1,opt"`
	Iv     proto.Option[string] `protobuf:"bytes,2,opt"`
}

type COMMEntry struct {
	Key   proto.Option[string] `protobuf:"bytes,1,opt"`
	Value proto.Option[string] `protobuf:"bytes,2,opt"`
}
