// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.27.3
// source: pkg/proto/configuration/cloud/aws/aws.proto

package aws

import (
	http "github.com/buildbarn/bb-storage/pkg/proto/configuration/http"
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

type StaticCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKeyId     string `protobuf:"bytes,1,opt,name=access_key_id,json=accessKeyId,proto3" json:"access_key_id,omitempty"`
	SecretAccessKey string `protobuf:"bytes,2,opt,name=secret_access_key,json=secretAccessKey,proto3" json:"secret_access_key,omitempty"`
}

func (x *StaticCredentials) Reset() {
	*x = StaticCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticCredentials) ProtoMessage() {}

func (x *StaticCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticCredentials.ProtoReflect.Descriptor instead.
func (*StaticCredentials) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescGZIP(), []int{0}
}

func (x *StaticCredentials) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *StaticCredentials) GetSecretAccessKey() string {
	if x != nil {
		return x.SecretAccessKey
	}
	return ""
}

type WebIdentityRoleCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleArn   string `protobuf:"bytes,1,opt,name=role_arn,json=roleArn,proto3" json:"role_arn,omitempty"`
	TokenFile string `protobuf:"bytes,2,opt,name=token_file,json=tokenFile,proto3" json:"token_file,omitempty"`
}

func (x *WebIdentityRoleCredentials) Reset() {
	*x = WebIdentityRoleCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebIdentityRoleCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebIdentityRoleCredentials) ProtoMessage() {}

func (x *WebIdentityRoleCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebIdentityRoleCredentials.ProtoReflect.Descriptor instead.
func (*WebIdentityRoleCredentials) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescGZIP(), []int{1}
}

func (x *WebIdentityRoleCredentials) GetRoleArn() string {
	if x != nil {
		return x.RoleArn
	}
	return ""
}

func (x *WebIdentityRoleCredentials) GetTokenFile() string {
	if x != nil {
		return x.TokenFile
	}
	return ""
}

type SessionConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// Types that are assignable to Credentials:
	//
	//	*SessionConfiguration_StaticCredentials
	//	*SessionConfiguration_WebIdentityRoleCredentials
	Credentials isSessionConfiguration_Credentials `protobuf_oneof:"credentials"`
	HttpClient  *http.ClientConfiguration          `protobuf:"bytes,6,opt,name=http_client,json=httpClient,proto3" json:"http_client,omitempty"`
}

func (x *SessionConfiguration) Reset() {
	*x = SessionConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionConfiguration) ProtoMessage() {}

func (x *SessionConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionConfiguration.ProtoReflect.Descriptor instead.
func (*SessionConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescGZIP(), []int{2}
}

func (x *SessionConfiguration) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (m *SessionConfiguration) GetCredentials() isSessionConfiguration_Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (x *SessionConfiguration) GetStaticCredentials() *StaticCredentials {
	if x, ok := x.GetCredentials().(*SessionConfiguration_StaticCredentials); ok {
		return x.StaticCredentials
	}
	return nil
}

func (x *SessionConfiguration) GetWebIdentityRoleCredentials() *WebIdentityRoleCredentials {
	if x, ok := x.GetCredentials().(*SessionConfiguration_WebIdentityRoleCredentials); ok {
		return x.WebIdentityRoleCredentials
	}
	return nil
}

func (x *SessionConfiguration) GetHttpClient() *http.ClientConfiguration {
	if x != nil {
		return x.HttpClient
	}
	return nil
}

type isSessionConfiguration_Credentials interface {
	isSessionConfiguration_Credentials()
}

type SessionConfiguration_StaticCredentials struct {
	StaticCredentials *StaticCredentials `protobuf:"bytes,5,opt,name=static_credentials,json=staticCredentials,proto3,oneof"`
}

type SessionConfiguration_WebIdentityRoleCredentials struct {
	WebIdentityRoleCredentials *WebIdentityRoleCredentials `protobuf:"bytes,7,opt,name=web_identity_role_credentials,json=webIdentityRoleCredentials,proto3,oneof"`
}

func (*SessionConfiguration_StaticCredentials) isSessionConfiguration_Credentials() {}

func (*SessionConfiguration_WebIdentityRoleCredentials) isSessionConfiguration_Credentials() {}

var File_pkg_proto_configuration_cloud_aws_aws_proto protoreflect.FileDescriptor

var file_pkg_proto_configuration_cloud_aws_aws_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x77, 0x73, 0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x77, 0x73,
	0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x11, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x22,
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x22, 0x56,
	0x0a, 0x1a, 0x57, 0x65, 0x62, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x6f, 0x6c,
	0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x8f, 0x03, 0x0a, 0x14, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x65, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x11, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x82,
	0x01, 0x0a, 0x1d, 0x77, 0x65, 0x62, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61,
	0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x57, 0x65, 0x62, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x1a, 0x77, 0x65, 0x62, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x12, 0x52, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x68, 0x74, 0x74,
	0x70, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x03,
	0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e,
	0x2f, 0x62, 0x62, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x77, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescData = file_pkg_proto_configuration_cloud_aws_aws_proto_rawDesc
)

func file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescData)
	})
	return file_pkg_proto_configuration_cloud_aws_aws_proto_rawDescData
}

var file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_proto_configuration_cloud_aws_aws_proto_goTypes = []interface{}{
	(*StaticCredentials)(nil),          // 0: buildbarn.configuration.cloud.aws.StaticCredentials
	(*WebIdentityRoleCredentials)(nil), // 1: buildbarn.configuration.cloud.aws.WebIdentityRoleCredentials
	(*SessionConfiguration)(nil),       // 2: buildbarn.configuration.cloud.aws.SessionConfiguration
	(*http.ClientConfiguration)(nil),   // 3: buildbarn.configuration.http.ClientConfiguration
}
var file_pkg_proto_configuration_cloud_aws_aws_proto_depIdxs = []int32{
	0, // 0: buildbarn.configuration.cloud.aws.SessionConfiguration.static_credentials:type_name -> buildbarn.configuration.cloud.aws.StaticCredentials
	1, // 1: buildbarn.configuration.cloud.aws.SessionConfiguration.web_identity_role_credentials:type_name -> buildbarn.configuration.cloud.aws.WebIdentityRoleCredentials
	3, // 2: buildbarn.configuration.cloud.aws.SessionConfiguration.http_client:type_name -> buildbarn.configuration.http.ClientConfiguration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_cloud_aws_aws_proto_init() }
func file_pkg_proto_configuration_cloud_aws_aws_proto_init() {
	if File_pkg_proto_configuration_cloud_aws_aws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticCredentials); i {
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
		file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebIdentityRoleCredentials); i {
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
		file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionConfiguration); i {
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
	file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SessionConfiguration_StaticCredentials)(nil),
		(*SessionConfiguration_WebIdentityRoleCredentials)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_configuration_cloud_aws_aws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_cloud_aws_aws_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_cloud_aws_aws_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_cloud_aws_aws_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_cloud_aws_aws_proto = out.File
	file_pkg_proto_configuration_cloud_aws_aws_proto_rawDesc = nil
	file_pkg_proto_configuration_cloud_aws_aws_proto_goTypes = nil
	file_pkg_proto_configuration_cloud_aws_aws_proto_depIdxs = nil
}
