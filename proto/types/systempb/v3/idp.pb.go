// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: proto/types/systempb/v3/idp.proto

package systemv3

import (
	v3 "github.com/RafayLabs/rcloud-base/proto/types/commonpb/v3"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type Idp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string       `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind       string       `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata   *v3.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec       *IdpSpec     `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	Status     *v3.Status   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Idp) Reset() {
	*x = Idp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_systempb_v3_idp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Idp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Idp) ProtoMessage() {}

func (x *Idp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_systempb_v3_idp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Idp.ProtoReflect.Descriptor instead.
func (*Idp) Descriptor() ([]byte, []int) {
	return file_proto_types_systempb_v3_idp_proto_rawDescGZIP(), []int{0}
}

func (x *Idp) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Idp) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Idp) GetMetadata() *v3.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Idp) GetSpec() *IdpSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Idp) GetStatus() *v3.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type IdpSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdpName            string `protobuf:"bytes,1,opt,name=idpName,proto3" json:"idpName,omitempty"`
	Domain             string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	AcsUrl             string `protobuf:"bytes,3,opt,name=acsUrl,proto3" json:"acsUrl,omitempty"`
	SsoUrl             string `protobuf:"bytes,4,opt,name=ssoUrl,proto3" json:"ssoUrl,omitempty"`
	IdpCert            string `protobuf:"bytes,5,opt,name=idpCert,proto3" json:"idpCert,omitempty"`
	SpCert             string `protobuf:"bytes,6,opt,name=spCert,proto3" json:"spCert,omitempty"`
	MetadataUrl        string `protobuf:"bytes,7,opt,name=metadataUrl,proto3" json:"metadataUrl,omitempty"`
	MetadataFilename   string `protobuf:"bytes,8,opt,name=metadataFilename,proto3" json:"metadataFilename,omitempty"`
	SaeEnabled         bool   `protobuf:"varint,9,opt,name=saeEnabled,proto3" json:"saeEnabled,omitempty"`
	GroupAttributeName string `protobuf:"bytes,10,opt,name=groupAttributeName,proto3" json:"groupAttributeName,omitempty"`
	NameIdFormat       string `protobuf:"bytes,11,opt,name=nameIdFormat,proto3" json:"nameIdFormat,omitempty"`
	ConsumerBinding    string `protobuf:"bytes,12,opt,name=consumerBinding,proto3" json:"consumerBinding,omitempty"`
	SpEntityId         string `protobuf:"bytes,13,opt,name=spEntityId,proto3" json:"spEntityId,omitempty"`
}

func (x *IdpSpec) Reset() {
	*x = IdpSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_systempb_v3_idp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdpSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdpSpec) ProtoMessage() {}

func (x *IdpSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_systempb_v3_idp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdpSpec.ProtoReflect.Descriptor instead.
func (*IdpSpec) Descriptor() ([]byte, []int) {
	return file_proto_types_systempb_v3_idp_proto_rawDescGZIP(), []int{1}
}

func (x *IdpSpec) GetIdpName() string {
	if x != nil {
		return x.IdpName
	}
	return ""
}

func (x *IdpSpec) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *IdpSpec) GetAcsUrl() string {
	if x != nil {
		return x.AcsUrl
	}
	return ""
}

func (x *IdpSpec) GetSsoUrl() string {
	if x != nil {
		return x.SsoUrl
	}
	return ""
}

func (x *IdpSpec) GetIdpCert() string {
	if x != nil {
		return x.IdpCert
	}
	return ""
}

func (x *IdpSpec) GetSpCert() string {
	if x != nil {
		return x.SpCert
	}
	return ""
}

func (x *IdpSpec) GetMetadataUrl() string {
	if x != nil {
		return x.MetadataUrl
	}
	return ""
}

func (x *IdpSpec) GetMetadataFilename() string {
	if x != nil {
		return x.MetadataFilename
	}
	return ""
}

func (x *IdpSpec) GetSaeEnabled() bool {
	if x != nil {
		return x.SaeEnabled
	}
	return false
}

func (x *IdpSpec) GetGroupAttributeName() string {
	if x != nil {
		return x.GroupAttributeName
	}
	return ""
}

func (x *IdpSpec) GetNameIdFormat() string {
	if x != nil {
		return x.NameIdFormat
	}
	return ""
}

func (x *IdpSpec) GetConsumerBinding() string {
	if x != nil {
		return x.ConsumerBinding
	}
	return ""
}

func (x *IdpSpec) GetSpEntityId() string {
	if x != nil {
		return x.SpEntityId
	}
	return ""
}

type IdpList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string           `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind       string           `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata   *v3.ListMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Items      []*Idp           `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *IdpList) Reset() {
	*x = IdpList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_systempb_v3_idp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdpList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdpList) ProtoMessage() {}

func (x *IdpList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_systempb_v3_idp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdpList.ProtoReflect.Descriptor instead.
func (*IdpList) Descriptor() ([]byte, []int) {
	return file_proto_types_systempb_v3_idp_proto_rawDescGZIP(), []int{2}
}

func (x *IdpList) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *IdpList) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *IdpList) GetMetadata() *v3.ListMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *IdpList) GetItems() []*Idp {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_proto_types_systempb_v3_idp_proto protoreflect.FileDescriptor

var file_proto_types_systempb_v3_idp_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x69, 0x64, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x19, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x1a, 0x24,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x04, 0x0a, 0x03, 0x49, 0x64, 0x70, 0x12, 0x67, 0x0a, 0x0a,
	0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x47, 0x92, 0x41, 0x44, 0x2a, 0x0b, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x32, 0x1f, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20,
	0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x70, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x3a, 0x14, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x6b, 0x38, 0x73, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x33, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x28, 0x92, 0x41, 0x25, 0x2a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x32, 0x18,
	0x4b, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x70, 0x20,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x03, 0x49, 0x64, 0x70, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x6c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2b, 0x92, 0x41, 0x28, 0x2a,
	0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0x1c, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x70, 0x20, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x5b, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x64, 0x70, 0x53,
	0x70, 0x65, 0x63, 0x42, 0x23, 0x92, 0x41, 0x20, 0x2a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x32, 0x18,
	0x53, 0x70, 0x65, 0x63, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x70, 0x20,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x60,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x25, 0x92, 0x41, 0x22, 0x2a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x16,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x40, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x3a, 0x35, 0x92, 0x41, 0x32, 0x0a, 0x30, 0x2a, 0x03, 0x49, 0x64, 0x70, 0x32, 0x03, 0x49, 0x64,
	0x70, 0xd2, 0x01, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0xd2, 0x01,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0xd2, 0x01, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xd2, 0x01, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0xa9, 0x03, 0x0a, 0x07, 0x49, 0x64, 0x70, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x73, 0x55, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x73, 0x6f, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x73, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x70, 0x43, 0x65, 0x72, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x70, 0x43, 0x65, 0x72, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x70, 0x43, 0x65, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x70, 0x43, 0x65, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x72, 0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x61, 0x65, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x61, 0x65, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x61, 0x6d,
	0x65, 0x49, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x22, 0xb6, 0x03, 0x0a, 0x07, 0x49, 0x64, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x6e, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x4e, 0x92, 0x41, 0x4b, 0x2a, 0x0b, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x24, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x70, 0x20, 0x6c, 0x69,
	0x73, 0x74, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x14, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x6b, 0x38, 0x73, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x76,
	0x33, 0x40, 0x01, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x47, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0x92,
	0x41, 0x30, 0x2a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x32, 0x1d, 0x4b, 0x69, 0x6e, 0x64, 0x20, 0x6f,
	0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x70, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x07, 0x49, 0x64, 0x70, 0x4c, 0x69, 0x73, 0x74,
	0x40, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x77, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x61, 0x66,
	0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x32, 0x92, 0x41, 0x2f, 0x2a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x32, 0x21, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6f, 0x66, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x70, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x40, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x5d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x64, 0x70,
	0x42, 0x27, 0x92, 0x41, 0x24, 0x2a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x19, 0x4c, 0x69,
	0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x70, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x40, 0x01, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x3a, 0x1a, 0x92, 0x41, 0x17, 0x0a, 0x15, 0x2a, 0x07, 0x49, 0x64, 0x70, 0x4c, 0x69, 0x73, 0x74,
	0x32, 0x08, 0x69, 0x64, 0x70, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x40, 0x01, 0x42, 0xf8, 0x01, 0x0a,
	0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x42, 0x08,
	0x49, 0x64, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x66, 0x61, 0x79, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x73, 0x2f, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x3b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x76, 0x33,
	0xa2, 0x02, 0x04, 0x52, 0x44, 0x54, 0x53, 0xaa, 0x02, 0x19, 0x52, 0x61, 0x66, 0x61, 0x79, 0x2e,
	0x44, 0x65, 0x76, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x56, 0x33, 0xca, 0x02, 0x19, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x56, 0x33, 0xe2,
	0x02, 0x25, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x52, 0x61, 0x66, 0x61, 0x79, 0x3a,
	0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_systempb_v3_idp_proto_rawDescOnce sync.Once
	file_proto_types_systempb_v3_idp_proto_rawDescData = file_proto_types_systempb_v3_idp_proto_rawDesc
)

func file_proto_types_systempb_v3_idp_proto_rawDescGZIP() []byte {
	file_proto_types_systempb_v3_idp_proto_rawDescOnce.Do(func() {
		file_proto_types_systempb_v3_idp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_systempb_v3_idp_proto_rawDescData)
	})
	return file_proto_types_systempb_v3_idp_proto_rawDescData
}

var file_proto_types_systempb_v3_idp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_types_systempb_v3_idp_proto_goTypes = []interface{}{
	(*Idp)(nil),             // 0: rafay.dev.types.system.v3.Idp
	(*IdpSpec)(nil),         // 1: rafay.dev.types.system.v3.IdpSpec
	(*IdpList)(nil),         // 2: rafay.dev.types.system.v3.IdpList
	(*v3.Metadata)(nil),     // 3: rafay.dev.types.common.v3.Metadata
	(*v3.Status)(nil),       // 4: rafay.dev.types.common.v3.Status
	(*v3.ListMetadata)(nil), // 5: rafay.dev.types.common.v3.ListMetadata
}
var file_proto_types_systempb_v3_idp_proto_depIdxs = []int32{
	3, // 0: rafay.dev.types.system.v3.Idp.metadata:type_name -> rafay.dev.types.common.v3.Metadata
	1, // 1: rafay.dev.types.system.v3.Idp.spec:type_name -> rafay.dev.types.system.v3.IdpSpec
	4, // 2: rafay.dev.types.system.v3.Idp.status:type_name -> rafay.dev.types.common.v3.Status
	5, // 3: rafay.dev.types.system.v3.IdpList.metadata:type_name -> rafay.dev.types.common.v3.ListMetadata
	0, // 4: rafay.dev.types.system.v3.IdpList.items:type_name -> rafay.dev.types.system.v3.Idp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_types_systempb_v3_idp_proto_init() }
func file_proto_types_systempb_v3_idp_proto_init() {
	if File_proto_types_systempb_v3_idp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_systempb_v3_idp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Idp); i {
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
		file_proto_types_systempb_v3_idp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdpSpec); i {
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
		file_proto_types_systempb_v3_idp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdpList); i {
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
			RawDescriptor: file_proto_types_systempb_v3_idp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_systempb_v3_idp_proto_goTypes,
		DependencyIndexes: file_proto_types_systempb_v3_idp_proto_depIdxs,
		MessageInfos:      file_proto_types_systempb_v3_idp_proto_msgTypes,
	}.Build()
	File_proto_types_systempb_v3_idp_proto = out.File
	file_proto_types_systempb_v3_idp_proto_rawDesc = nil
	file_proto_types_systempb_v3_idp_proto_goTypes = nil
	file_proto_types_systempb_v3_idp_proto_depIdxs = nil
}
