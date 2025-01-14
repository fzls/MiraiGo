// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/web/WebSsoBody.proto

package web

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type STServiceMonitItem struct {
	Cmd     proto.Option[string] `protobuf:"bytes,1,opt"`
	Url     proto.Option[string] `protobuf:"bytes,2,opt"`
	Errcode proto.Option[int32]  `protobuf:"varint,3,opt"`
	Cost    proto.Option[uint32] `protobuf:"varint,4,opt"`
	Src     proto.Option[uint32] `protobuf:"varint,5,opt"`
}

type STServiceMonitReq struct {
	List []*STServiceMonitItem `protobuf:"bytes,1,rep"`
}

type WebSsoControlData struct {
	Frequency   proto.Option[uint32] `protobuf:"varint,1,opt"`
	PackageSize proto.Option[uint32] `protobuf:"varint,2,opt"`
}

type WebSsoRequestBody struct {
	Version proto.Option[uint32] `protobuf:"varint,1,opt"`
	Type    proto.Option[uint32] `protobuf:"varint,2,opt"`
	Data    proto.Option[string] `protobuf:"bytes,3,opt"`
	WebData proto.Option[string] `protobuf:"bytes,4,opt"`
}

type WebSsoResponseBody struct {
	Version     proto.Option[uint32] `protobuf:"varint,1,opt"`
	Type        proto.Option[uint32] `protobuf:"varint,2,opt"`
	Ret         proto.Option[uint32] `protobuf:"varint,3,opt"`
	Data        proto.Option[string] `protobuf:"bytes,4,opt"`
	ControlData *WebSsoControlData   `protobuf:"bytes,5,opt"`
}
