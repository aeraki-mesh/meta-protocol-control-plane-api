// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: api/meta_protocol_proxy/filters/local_ratelimit/v1alpha/local_ratelimit.proto

package localratelimitv1alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v1alpha "github.com/aeraki-mesh/meta-protocol-control-plane-api/aeraki/meta_protocol_proxy/config/route/v1alpha"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type LocalRateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatPrefix  string                     `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	TokenBucket *v3.TokenBucket            `protobuf:"bytes,2,opt,name=token_bucket,json=tokenBucket,proto3" json:"token_bucket,omitempty"`
	Conditions  []*LocalRateLimitCondition `protobuf:"bytes,3,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *LocalRateLimit) Reset() {
	*x = LocalRateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalRateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalRateLimit) ProtoMessage() {}

func (x *LocalRateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalRateLimit.ProtoReflect.Descriptor instead.
func (*LocalRateLimit) Descriptor() ([]byte, []int) {
	return file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDescGZIP(), []int{0}
}

func (x *LocalRateLimit) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *LocalRateLimit) GetTokenBucket() *v3.TokenBucket {
	if x != nil {
		return x.TokenBucket
	}
	return nil
}

func (x *LocalRateLimit) GetConditions() []*LocalRateLimitCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

type LocalRateLimitCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Match       *v1alpha.RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	TokenBucket *v3.TokenBucket     `protobuf:"bytes,2,opt,name=token_bucket,json=tokenBucket,proto3" json:"token_bucket,omitempty"`
}

func (x *LocalRateLimitCondition) Reset() {
	*x = LocalRateLimitCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalRateLimitCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalRateLimitCondition) ProtoMessage() {}

func (x *LocalRateLimitCondition) ProtoReflect() protoreflect.Message {
	mi := &file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalRateLimitCondition.ProtoReflect.Descriptor instead.
func (*LocalRateLimitCondition) Descriptor() ([]byte, []int) {
	return file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDescGZIP(), []int{1}
}

func (x *LocalRateLimitCondition) GetMatch() *v1alpha.RouteMatch {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *LocalRateLimitCondition) GetTokenBucket() *v3.TokenBucket {
	if x != nil {
		return x.TokenBucket
	}
	return nil
}

var File_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto protoreflect.FileDescriptor

var file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x3a, 0x61, 0x65, 0x72, 0x61, 0x6b, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee,
	0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x3d, 0x0a, 0x0c, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0b, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x73, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x53,
	0x2e, 0x61, 0x65, 0x72, 0x61, 0x6b, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0xbf, 0x01, 0x0a, 0x17, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x05, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x61, 0x65, 0x72,
	0x61, 0x6b, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x47, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x42, 0xe9, 0x01, 0x0a, 0x3e, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x65, 0x72, 0x61, 0x6b, 0x69,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x42, 0x13, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x87, 0x01, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x65, 0x72, 0x61, 0x6b, 0x69, 0x2d,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x65, 0x72, 0x61, 0x6b, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDescOnce sync.Once
	file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDescData = file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDesc
)

func file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDescGZIP() []byte {
	file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDescOnce.Do(func() {
		file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDescData)
	})
	return file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDescData
}

var file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_goTypes = []interface{}{
	(*LocalRateLimit)(nil),          // 0: aeraki.meta_protocol_proxy.filters.local_ratelimit.v1alpha.LocalRateLimit
	(*LocalRateLimitCondition)(nil), // 1: aeraki.meta_protocol_proxy.filters.local_ratelimit.v1alpha.LocalRateLimitCondition
	(*v3.TokenBucket)(nil),          // 2: envoy.type.v3.TokenBucket
	(*v1alpha.RouteMatch)(nil),      // 3: aeraki.meta_protocol_proxy.config.route.v1alpha.RouteMatch
}
var file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_depIdxs = []int32{
	2, // 0: aeraki.meta_protocol_proxy.filters.local_ratelimit.v1alpha.LocalRateLimit.token_bucket:type_name -> envoy.type.v3.TokenBucket
	1, // 1: aeraki.meta_protocol_proxy.filters.local_ratelimit.v1alpha.LocalRateLimit.conditions:type_name -> aeraki.meta_protocol_proxy.filters.local_ratelimit.v1alpha.LocalRateLimitCondition
	3, // 2: aeraki.meta_protocol_proxy.filters.local_ratelimit.v1alpha.LocalRateLimitCondition.match:type_name -> aeraki.meta_protocol_proxy.config.route.v1alpha.RouteMatch
	2, // 3: aeraki.meta_protocol_proxy.filters.local_ratelimit.v1alpha.LocalRateLimitCondition.token_bucket:type_name -> envoy.type.v3.TokenBucket
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_init()
}
func file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_init() {
	if File_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalRateLimit); i {
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
		file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalRateLimitCondition); i {
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
			RawDescriptor: file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_goTypes,
		DependencyIndexes: file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_depIdxs,
		MessageInfos:      file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_msgTypes,
	}.Build()
	File_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto = out.File
	file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_rawDesc = nil
	file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_goTypes = nil
	file_api_meta_protocol_proxy_filters_local_ratelimit_v1alpha_local_ratelimit_proto_depIdxs = nil
}
