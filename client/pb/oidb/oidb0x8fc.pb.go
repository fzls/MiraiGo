// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/oidb/oidb0x8fc.proto

package oidb

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type D8FCReqBody struct {
	GroupCode      proto.Option[int64] `protobuf:"varint,1,opt"`
	ShowFlag       proto.Option[int32] `protobuf:"varint,2,opt"`
	MemLevelInfo   []*D8FCMemberInfo   `protobuf:"bytes,3,rep"`
	LevelName      []*D8FCLevelName    `protobuf:"bytes,4,rep"`
	UpdateTime     proto.Option[int32] `protobuf:"varint,5,opt"`
	OfficeMode     proto.Option[int32] `protobuf:"varint,6,opt"`
	GroupOpenAppid proto.Option[int32] `protobuf:"varint,7,opt"`
	MsgClientInfo  *D8FCClientInfo     `protobuf:"bytes,8,opt"`
	AuthKey        []byte              `protobuf:"bytes,9,opt"`
}

type D8FCMemberInfo struct {
	Uin                    proto.Option[int64] `protobuf:"varint,1,opt"`
	Point                  proto.Option[int32] `protobuf:"varint,2,opt"`
	ActiveDay              proto.Option[int32] `protobuf:"varint,3,opt"`
	Level                  proto.Option[int32] `protobuf:"varint,4,opt"`
	SpecialTitle           []byte              `protobuf:"bytes,5,opt"`
	SpecialTitleExpireTime proto.Option[int32] `protobuf:"varint,6,opt"`
	UinName                []byte              `protobuf:"bytes,7,opt"`
	MemberCardName         []byte              `protobuf:"bytes,8,opt"`
	Phone                  []byte              `protobuf:"bytes,9,opt"`
	Email                  []byte              `protobuf:"bytes,10,opt"`
	Remark                 []byte              `protobuf:"bytes,11,opt"`
	Gender                 proto.Option[int32] `protobuf:"varint,12,opt"`
	Job                    []byte              `protobuf:"bytes,13,opt"`
	TribeLevel             proto.Option[int32] `protobuf:"varint,14,opt"`
	TribePoint             proto.Option[int32] `protobuf:"varint,15,opt"`
	RichCardName           []*D8FCCardNameElem `protobuf:"bytes,16,rep"`
	CommRichCardName       []byte              `protobuf:"bytes,17,opt"`
}

type D8FCCardNameElem struct {
	EnumCardType proto.Option[int32] `protobuf:"varint,1,opt"`
	Value        []byte              `protobuf:"bytes,2,opt"`
}

type D8FCLevelName struct {
	Level proto.Option[int32]  `protobuf:"varint,1,opt"`
	Name  proto.Option[string] `protobuf:"bytes,2,opt"`
}

type D8FCClientInfo struct {
	Implat       proto.Option[int32]  `protobuf:"varint,1,opt"`
	IngClientver proto.Option[string] `protobuf:"bytes,2,opt"`
}

type D8FCCommCardNameBuf struct {
	RichCardName []*D8FCRichCardNameElem `protobuf:"bytes,1,rep"`
}

type D8FCRichCardNameElem struct {
	Ctrl []byte `protobuf:"bytes,1,opt"`
	Text []byte `protobuf:"bytes,2,opt"`
}
