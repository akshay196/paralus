// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: proto/types/systempb/v3/partner.proto

package systemv3

import (
	v3 "github.com/RafayLabs/rcloud-base/proto/types/commonpb/v3"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PartnerSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host              string           `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Domain            string           `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	TosLink           string           `protobuf:"bytes,3,opt,name=tosLink,proto3" json:"tosLink,omitempty"`
	LogoLink          string           `protobuf:"bytes,4,opt,name=logoLink,proto3" json:"logoLink,omitempty"`
	NotificationEmail string           `protobuf:"bytes,5,opt,name=notificationEmail,proto3" json:"notificationEmail,omitempty"`
	HelpdeskEmail     string           `protobuf:"bytes,6,opt,name=helpdeskEmail,proto3" json:"helpdeskEmail,omitempty"`
	ProductName       string           `protobuf:"bytes,7,opt,name=productName,proto3" json:"productName,omitempty"`
	SupportTeamName   string           `protobuf:"bytes,8,opt,name=supportTeamName,proto3" json:"supportTeamName,omitempty"`
	OpsHost           string           `protobuf:"bytes,9,opt,name=opsHost,proto3" json:"opsHost,omitempty"`
	FavIconLink       string           `protobuf:"bytes,10,opt,name=favIconLink,proto3" json:"favIconLink,omitempty"`
	IsTOTPEnabled     bool             `protobuf:"varint,11,opt,name=isTOTPEnabled,proto3" json:"isTOTPEnabled,omitempty"`
	Settings          *structpb.Struct `protobuf:"bytes,12,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *PartnerSpec) Reset() {
	*x = PartnerSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_systempb_v3_partner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartnerSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartnerSpec) ProtoMessage() {}

func (x *PartnerSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_systempb_v3_partner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartnerSpec.ProtoReflect.Descriptor instead.
func (*PartnerSpec) Descriptor() ([]byte, []int) {
	return file_proto_types_systempb_v3_partner_proto_rawDescGZIP(), []int{0}
}

func (x *PartnerSpec) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *PartnerSpec) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *PartnerSpec) GetTosLink() string {
	if x != nil {
		return x.TosLink
	}
	return ""
}

func (x *PartnerSpec) GetLogoLink() string {
	if x != nil {
		return x.LogoLink
	}
	return ""
}

func (x *PartnerSpec) GetNotificationEmail() string {
	if x != nil {
		return x.NotificationEmail
	}
	return ""
}

func (x *PartnerSpec) GetHelpdeskEmail() string {
	if x != nil {
		return x.HelpdeskEmail
	}
	return ""
}

func (x *PartnerSpec) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *PartnerSpec) GetSupportTeamName() string {
	if x != nil {
		return x.SupportTeamName
	}
	return ""
}

func (x *PartnerSpec) GetOpsHost() string {
	if x != nil {
		return x.OpsHost
	}
	return ""
}

func (x *PartnerSpec) GetFavIconLink() string {
	if x != nil {
		return x.FavIconLink
	}
	return ""
}

func (x *PartnerSpec) GetIsTOTPEnabled() bool {
	if x != nil {
		return x.IsTOTPEnabled
	}
	return false
}

func (x *PartnerSpec) GetSettings() *structpb.Struct {
	if x != nil {
		return x.Settings
	}
	return nil
}

type Partner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string       `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind       string       `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata   *v3.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec       *PartnerSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	Status     *v3.Status   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Partner) Reset() {
	*x = Partner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_systempb_v3_partner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partner) ProtoMessage() {}

func (x *Partner) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_systempb_v3_partner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partner.ProtoReflect.Descriptor instead.
func (*Partner) Descriptor() ([]byte, []int) {
	return file_proto_types_systempb_v3_partner_proto_rawDescGZIP(), []int{1}
}

func (x *Partner) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Partner) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Partner) GetMetadata() *v3.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Partner) GetSpec() *PartnerSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Partner) GetStatus() *v3.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_proto_types_systempb_v3_partner_proto protoreflect.FileDescriptor

var file_proto_types_systempb_v3_partner_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x07, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x32, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0x92, 0x41, 0x1b, 0x2a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x32,
	0x13, 0x48, 0x6f, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0x92, 0x41, 0x1f, 0x2a,
	0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x32, 0x15, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x20,
	0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x47, 0x0a, 0x07, 0x74, 0x6f, 0x73, 0x4c, 0x69, 0x6e,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0x92, 0x41, 0x2a, 0x2a, 0x03, 0x54, 0x4f,
	0x53, 0x32, 0x23, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x20, 0x75, 0x72, 0x6c, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x07, 0x74, 0x6f, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x3e, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x22, 0x92, 0x41, 0x1f, 0x2a, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x32, 0x17, 0x4c, 0x6f,
	0x67, 0x6f, 0x20, 0x75, 0x72, 0x6c, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x68, 0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0x92, 0x41, 0x37, 0x2a,
	0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x32, 0x21, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x58, 0x0a, 0x0d, 0x68, 0x65, 0x6c,
	0x70, 0x64, 0x65, 0x73, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x32, 0x92, 0x41, 0x2f, 0x2a, 0x0e, 0x48, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x20,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0x1d, 0x48, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x20,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x52, 0x0d, 0x68, 0x65, 0x6c, 0x70, 0x64, 0x65, 0x73, 0x6b, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x41, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0x92, 0x41, 0x1c, 0x2a, 0x0c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x0c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x53, 0x0a, 0x0f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x29, 0x92, 0x41, 0x26, 0x2a, 0x11, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x54, 0x65,
	0x61, 0x6d, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x11, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x20, 0x54, 0x65, 0x61, 0x6d, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0f, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x6f,
	0x70, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0x92, 0x41,
	0x2a, 0x2a, 0x08, 0x4f, 0x50, 0x53, 0x20, 0x48, 0x6f, 0x73, 0x74, 0x32, 0x1e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x68, 0x6f, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x07, 0x6f, 0x70, 0x73,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0b, 0x66, 0x61, 0x76, 0x49, 0x63, 0x6f, 0x6e, 0x4c,
	0x69, 0x6e, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0x92, 0x41, 0x1e, 0x2a, 0x0d,
	0x46, 0x61, 0x76, 0x20, 0x49, 0x63, 0x6f, 0x6e, 0x20, 0x4c, 0x69, 0x6e, 0x6b, 0x32, 0x0d, 0x46,
	0x61, 0x76, 0x20, 0x49, 0x63, 0x6f, 0x6e, 0x20, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0b, 0x66, 0x61,
	0x76, 0x49, 0x63, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x5f, 0x0a, 0x0d, 0x69, 0x73, 0x54,
	0x4f, 0x54, 0x50, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x39, 0x92, 0x41, 0x36, 0x2a, 0x0c, 0x54, 0x4f, 0x54, 0x50, 0x20, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x32, 0x26, 0x54, 0x4f, 0x54, 0x50, 0x20, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x20, 0x66, 0x6c, 0x61, 0x67, 0x20, 0x61, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x20, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0d, 0x69, 0x73, 0x54,
	0x4f, 0x54, 0x50, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x5b, 0x0a, 0x08, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x26, 0x92, 0x41, 0x23, 0x2a, 0x08, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x32, 0x17, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x20, 0x6f,
	0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x08, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x3a, 0x33, 0x92, 0x41, 0x30, 0x0a, 0x2e, 0x2a, 0x15,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x20, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x15, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x20, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x98, 0x04, 0x0a,
	0x07, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x65, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x45, 0x92, 0x41,
	0x42, 0x2a, 0x0b, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x1b,
	0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x14, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x6b, 0x38, 0x73, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x76,
	0x33, 0x40, 0x01, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x3e, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0x92,
	0x41, 0x27, 0x2a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x32, 0x14, 0x4b, 0x69, 0x6e, 0x64, 0x20, 0x6f,
	0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x07,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x40, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12,
	0x68, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x27, 0x92, 0x41, 0x24, 0x2a, 0x08, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0x18, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x20,
	0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5b, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x33, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x42,
	0x1f, 0x92, 0x41, 0x1c, 0x2a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x32, 0x14, 0x53, 0x70, 0x65, 0x63,
	0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x60, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x25, 0x92, 0x41, 0x22, 0x2a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x16, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x6f,
	0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x40, 0x01,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x3d, 0x92, 0x41, 0x3a, 0x0a, 0x38, 0x2a,
	0x07, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x32, 0x07, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0xd2, 0x01, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0xd2, 0x01,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0xd2, 0x01, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xd2, 0x01, 0x04, 0x73, 0x70, 0x65, 0x63, 0x42, 0xfc, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e,
	0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x33, 0x42, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x66, 0x61, 0x79, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x73, 0x2f, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x3b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x76, 0x33, 0xa2,
	0x02, 0x04, 0x52, 0x44, 0x54, 0x53, 0xaa, 0x02, 0x19, 0x52, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x44,
	0x65, 0x76, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x56, 0x33, 0xca, 0x02, 0x19, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x56, 0x33, 0xe2, 0x02,
	0x25, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x5c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x52, 0x61, 0x66, 0x61, 0x79, 0x3a, 0x3a,
	0x44, 0x65, 0x76, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_systempb_v3_partner_proto_rawDescOnce sync.Once
	file_proto_types_systempb_v3_partner_proto_rawDescData = file_proto_types_systempb_v3_partner_proto_rawDesc
)

func file_proto_types_systempb_v3_partner_proto_rawDescGZIP() []byte {
	file_proto_types_systempb_v3_partner_proto_rawDescOnce.Do(func() {
		file_proto_types_systempb_v3_partner_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_systempb_v3_partner_proto_rawDescData)
	})
	return file_proto_types_systempb_v3_partner_proto_rawDescData
}

var file_proto_types_systempb_v3_partner_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_types_systempb_v3_partner_proto_goTypes = []interface{}{
	(*PartnerSpec)(nil),     // 0: rafay.dev.types.system.v3.PartnerSpec
	(*Partner)(nil),         // 1: rafay.dev.types.system.v3.Partner
	(*structpb.Struct)(nil), // 2: google.protobuf.Struct
	(*v3.Metadata)(nil),     // 3: rafay.dev.types.common.v3.Metadata
	(*v3.Status)(nil),       // 4: rafay.dev.types.common.v3.Status
}
var file_proto_types_systempb_v3_partner_proto_depIdxs = []int32{
	2, // 0: rafay.dev.types.system.v3.PartnerSpec.settings:type_name -> google.protobuf.Struct
	3, // 1: rafay.dev.types.system.v3.Partner.metadata:type_name -> rafay.dev.types.common.v3.Metadata
	0, // 2: rafay.dev.types.system.v3.Partner.spec:type_name -> rafay.dev.types.system.v3.PartnerSpec
	4, // 3: rafay.dev.types.system.v3.Partner.status:type_name -> rafay.dev.types.common.v3.Status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_types_systempb_v3_partner_proto_init() }
func file_proto_types_systempb_v3_partner_proto_init() {
	if File_proto_types_systempb_v3_partner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_systempb_v3_partner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartnerSpec); i {
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
		file_proto_types_systempb_v3_partner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Partner); i {
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
			RawDescriptor: file_proto_types_systempb_v3_partner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_systempb_v3_partner_proto_goTypes,
		DependencyIndexes: file_proto_types_systempb_v3_partner_proto_depIdxs,
		MessageInfos:      file_proto_types_systempb_v3_partner_proto_msgTypes,
	}.Build()
	File_proto_types_systempb_v3_partner_proto = out.File
	file_proto_types_systempb_v3_partner_proto_rawDesc = nil
	file_proto_types_systempb_v3_partner_proto_goTypes = nil
	file_proto_types_systempb_v3_partner_proto_depIdxs = nil
}
