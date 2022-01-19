// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: apps/host/pb/host.proto

package host

import (
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

type UpdateMode int32

const (
	UpdateMode_PATCH UpdateMode = 0
	UpdateMode_PUT   UpdateMode = 1
)

// Enum value maps for UpdateMode.
var (
	UpdateMode_name = map[int32]string{
		0: "PATCH",
		1: "PUT",
	}
	UpdateMode_value = map[string]int32{
		"PATCH": 0,
		"PUT":   1,
	}
)

func (x UpdateMode) Enum() *UpdateMode {
	p := new(UpdateMode)
	*p = x
	return p
}

func (x UpdateMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateMode) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_host_pb_host_proto_enumTypes[0].Descriptor()
}

func (UpdateMode) Type() protoreflect.EnumType {
	return &file_apps_host_pb_host_proto_enumTypes[0]
}

func (x UpdateMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateMode.Descriptor instead.
func (UpdateMode) EnumDescriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{0}
}

type Vendor int32

const (
	Vendor_ALI_CLOUD Vendor = 0
	Vendor_TX_CLOUD  Vendor = 1
	Vendor_HW_CLOUD  Vendor = 2
)

// Enum value maps for Vendor.
var (
	Vendor_name = map[int32]string{
		0: "ALI_CLOUD",
		1: "TX_CLOUD",
		2: "HW_CLOUD",
	}
	Vendor_value = map[string]int32{
		"ALI_CLOUD": 0,
		"TX_CLOUD":  1,
		"HW_CLOUD":  2,
	}
)

func (x Vendor) Enum() *Vendor {
	p := new(Vendor)
	*p = x
	return p
}

func (x Vendor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Vendor) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_host_pb_host_proto_enumTypes[1].Descriptor()
}

func (Vendor) Type() protoreflect.EnumType {
	return &file_apps_host_pb_host_proto_enumTypes[1]
}

func (x Vendor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Vendor.Descriptor instead.
func (Vendor) EnumDescriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{1}
}

type QueryHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize   int64  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNumber int64  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	Keywords   string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *QueryHostRequest) Reset() {
	*x = QueryHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryHostRequest) ProtoMessage() {}

func (x *QueryHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryHostRequest.ProtoReflect.Descriptor instead.
func (*QueryHostRequest) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{0}
}

func (x *QueryHostRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryHostRequest) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *QueryHostRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

type DescribeHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeHostRequest) Reset() {
	*x = DescribeHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeHostRequest) ProtoMessage() {}

func (x *DescribeHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeHostRequest.ProtoReflect.Descriptor instead.
func (*DescribeHostRequest) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeHostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateMode UpdateMode `protobuf:"varint,1,opt,name=update_mode,json=updateMode,proto3,enum=pb.UpdateMode" json:"update_mode,omitempty"`
	Host       *Host      `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *UpdateHostRequest) Reset() {
	*x = UpdateHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHostRequest) ProtoMessage() {}

func (x *UpdateHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHostRequest.ProtoReflect.Descriptor instead.
func (*UpdateHostRequest) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHostRequest) GetUpdateMode() UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return UpdateMode_PATCH
}

func (x *UpdateHostRequest) GetHost() *Host {
	if x != nil {
		return x.Host
	}
	return nil
}

type DeleteHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteHostRequest) Reset() {
	*x = DeleteHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHostRequest) ProtoMessage() {}

func (x *DeleteHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHostRequest.ProtoReflect.Descriptor instead.
func (*DeleteHostRequest) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteHostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"resource_hash"
	ResourceHash string `protobuf:"bytes,1,opt,name=resource_hash,json=resourceHash,proto3" json:"resource_hash,omitempty"`
	// @gotags: json:"describe_hash"
	DescribeHash string `protobuf:"bytes,2,opt,name=describe_hash,json=describeHash,proto3" json:"describe_hash,omitempty"`
	// @gotags: json:"resource"
	Resource *Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	// @gotags: json:"describe"
	Describe *Describe `protobuf:"bytes,4,opt,name=describe,proto3" json:"describe,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{4}
}

func (x *Host) GetResourceHash() string {
	if x != nil {
		return x.ResourceHash
	}
	return ""
}

func (x *Host) GetDescribeHash() string {
	if x != nil {
		return x.DescribeHash
	}
	return ""
}

func (x *Host) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Host) GetDescribe() *Describe {
	if x != nil {
		return x.Describe
	}
	return nil
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id" validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // 全局唯一Id
	// @gotags: json:"vendor" validate:"required"
	Vendor Vendor `protobuf:"varint,2,opt,name=vendor,proto3,enum=pb.Vendor" json:"vendor,omitempty"` // 厂商
	// @gotags: json:"region" validate:"required"`
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"` // 地域
	// @gotags: json:"zone"
	Zone string `protobuf:"bytes,4,opt,name=zone,proto3" json:"zone,omitempty"` // 区域
	// @gotags: json:"create_at" validate:"required"
	CreateAt int64 `protobuf:"varint,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"` // 创建时间
	// @gotags: json:"expire_at"
	ExpireAt int64 `protobuf:"varint,6,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"` // 过期时间
	// @gotags: json:"category"
	Category string `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty"` // 种类
	// @gotags: json:"type"
	Type string `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"` // 规格
	// @gotags: json:"instance_id"`
	InstanceId string `protobuf:"bytes,9,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"` // 实例ID
	// @gotags: json:"name" validate:"required"
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"` // 名称
	// @gotags: json:"description"
	Description string `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"` // 描述
	// @gotags: json:"status" validate:"required"
	Status string `protobuf:"bytes,12,opt,name=status,proto3" json:"status,omitempty"` // 服务商中的状态
	// @gotags: json:"tags"
	Tags map[string]string `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 标签
	// @gotags: json:"update_at"
	UpdateAt int64 `protobuf:"varint,14,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"` // 更新时间
	// @gotags: json:"sync_at"
	SyncAt int64 `protobuf:"varint,15,opt,name=sync_at,json=syncAt,proto3" json:"sync_at,omitempty"` // 同步时间
	// @gotags: json:"sync_accout"
	SyncAccount string `protobuf:"bytes,16,opt,name=sync_account,json=syncAccount,proto3" json:"sync_account,omitempty"` // 同步的账号
	// @gotags: json:"public_ip"
	PublicIp string `protobuf:"bytes,17,opt,name=public_ip,json=publicIp,proto3" json:"public_ip,omitempty"` // 公网IP
	// @gotags: json:"private_ip" validate:"required"
	PrivateIp string `protobuf:"bytes,18,opt,name=private_ip,json=privateIp,proto3" json:"private_ip,omitempty"` // 内网IP
	// @gotags: json:"pay_type"
	PayType string `protobuf:"bytes,19,opt,name=pay_type,json=payType,proto3" json:"pay_type,omitempty"` // 实例付费方式
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{5}
}

func (x *Resource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Resource) GetVendor() Vendor {
	if x != nil {
		return x.Vendor
	}
	return Vendor_ALI_CLOUD
}

func (x *Resource) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Resource) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Resource) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Resource) GetExpireAt() int64 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

func (x *Resource) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Resource) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Resource) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Resource) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Resource) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Resource) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Resource) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Resource) GetSyncAt() int64 {
	if x != nil {
		return x.SyncAt
	}
	return 0
}

func (x *Resource) GetSyncAccount() string {
	if x != nil {
		return x.SyncAccount
	}
	return ""
}

func (x *Resource) GetPublicIp() string {
	if x != nil {
		return x.PublicIp
	}
	return ""
}

func (x *Resource) GetPrivateIp() string {
	if x != nil {
		return x.PrivateIp
	}
	return ""
}

func (x *Resource) GetPayType() string {
	if x != nil {
		return x.PayType
	}
	return ""
}

type Describe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"cpu" validate:"required"
	Cpu int64 `protobuf:"varint,1,opt,name=cpu,proto3" json:"cpu,omitempty"` // 核数
	// @gotags: json:"memory" validate:"required"
	Memory int64 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"` // 内存
	// @gotags: json:"gpu_amount"
	GpuAmount int64 `protobuf:"varint,3,opt,name=gpu_amount,json=gpuAmount,proto3" json:"gpu_amount,omitempty"` // GPU数量
	// @gotags: json:"gpu_spec"
	GpuSpec string `protobuf:"bytes,4,opt,name=gpu_spec,json=gpuSpec,proto3" json:"gpu_spec,omitempty"` // GPU类型
	// @gotags: json:"os_type"
	OsType string `protobuf:"bytes,5,opt,name=os_type,json=osType,proto3" json:"os_type,omitempty"` // 操作系统类型，分为Windows和Linux
	// @gotags: json:"os_name"
	OsName string `protobuf:"bytes,6,opt,name=os_name,json=osName,proto3" json:"os_name,omitempty"` // 操作系统名称
	// @gotags: json:"serial_number"
	SerialNumber string `protobuf:"bytes,7,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"` // 序列号
	// @gotags: json:"image_id"
	ImageId string `protobuf:"bytes,8,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"` // 镜像ID
	// @gotags: json:"internet_max_bandwidth_out
	InternetMaxBandwidthOut int64 `protobuf:"varint,9,opt,name=internet_max_bandwidth_out,json=internetMaxBandwidthOut,proto3" json:"internet_max_bandwidth_out,omitempty"` // 公网出带宽最大值，单位为 Mbps
	// @gotags: json:"internet_max_bandwidth_in"
	InternetMaxBandwidthIn int64 `protobuf:"varint,10,opt,name=internet_max_bandwidth_in,json=internetMaxBandwidthIn,proto3" json:"internet_max_bandwidth_in,omitempty"` // 公网入带宽最大值，单位为 Mbps
	// @gotags: json:"key_pair_name"
	KeyPairName string `protobuf:"bytes,11,opt,name=key_pair_name,json=keyPairName,proto3" json:"key_pair_name,omitempty"` // 秘钥对名称
	// @gotags: json:"security_groups"
	SecurityGroups string `protobuf:"bytes,12,opt,name=security_groups,json=securityGroups,proto3" json:"security_groups,omitempty"` // 安全组  采用逗号分隔
}

func (x *Describe) Reset() {
	*x = Describe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Describe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Describe) ProtoMessage() {}

func (x *Describe) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Describe.ProtoReflect.Descriptor instead.
func (*Describe) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{6}
}

func (x *Describe) GetCpu() int64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Describe) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Describe) GetGpuAmount() int64 {
	if x != nil {
		return x.GpuAmount
	}
	return 0
}

func (x *Describe) GetGpuSpec() string {
	if x != nil {
		return x.GpuSpec
	}
	return ""
}

func (x *Describe) GetOsType() string {
	if x != nil {
		return x.OsType
	}
	return ""
}

func (x *Describe) GetOsName() string {
	if x != nil {
		return x.OsName
	}
	return ""
}

func (x *Describe) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *Describe) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *Describe) GetInternetMaxBandwidthOut() int64 {
	if x != nil {
		return x.InternetMaxBandwidthOut
	}
	return 0
}

func (x *Describe) GetInternetMaxBandwidthIn() int64 {
	if x != nil {
		return x.InternetMaxBandwidthIn
	}
	return 0
}

func (x *Describe) GetKeyPairName() string {
	if x != nil {
		return x.KeyPairName
	}
	return ""
}

func (x *Describe) GetSecurityGroups() string {
	if x != nil {
		return x.SecurityGroups
	}
	return ""
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// @gotags: json:"items"
	Items []*Host `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_host_pb_host_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_apps_host_pb_host_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_apps_host_pb_host_proto_rawDescGZIP(), []int{7}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*Host {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_host_pb_host_proto protoreflect.FileDescriptor

var file_apps_host_pb_host_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x68,
	0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x6c, 0x0a,
	0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x62, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x04,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x28,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x22, 0xd8, 0x04, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x7a,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x70, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x79,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa5, 0x03,
	0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70,
	0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x70, 0x75, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x67, 0x70, 0x75, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x70, 0x75, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x70, 0x75, 0x53, 0x70, 0x65, 0x63, 0x12, 0x17,
	0x0a, 0x07, 0x6f, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x73, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x3b, 0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x78,
	0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x4d, 0x61,
	0x78, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x4f, 0x75, 0x74, 0x12, 0x39, 0x0a,
	0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x61,
	0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x42, 0x61, 0x6e,
	0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x49, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x5f,
	0x70, 0x61, 0x69, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x3b, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x2a, 0x20, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50,
	0x55, 0x54, 0x10, 0x01, 0x2a, 0x33, 0x0a, 0x06, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x0d,
	0x0a, 0x09, 0x41, 0x4c, 0x49, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x54, 0x58, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x48,
	0x57, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x10, 0x02, 0x32, 0xe8, 0x01, 0x0a, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x1a, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x74, 0x12, 0x31, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x48, 0x6f, 0x73, 0x74, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x48, 0x41, 0x4f, 0x6c, 0x6f, 0x77, 0x6b, 0x65, 0x79, 0x2f, 0x72, 0x65, 0x73,
	0x74, 0x66, 0x75, 0x6c, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_host_pb_host_proto_rawDescOnce sync.Once
	file_apps_host_pb_host_proto_rawDescData = file_apps_host_pb_host_proto_rawDesc
)

func file_apps_host_pb_host_proto_rawDescGZIP() []byte {
	file_apps_host_pb_host_proto_rawDescOnce.Do(func() {
		file_apps_host_pb_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_host_pb_host_proto_rawDescData)
	})
	return file_apps_host_pb_host_proto_rawDescData
}

var file_apps_host_pb_host_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apps_host_pb_host_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_apps_host_pb_host_proto_goTypes = []interface{}{
	(UpdateMode)(0),             // 0: pb.UpdateMode
	(Vendor)(0),                 // 1: pb.Vendor
	(*QueryHostRequest)(nil),    // 2: pb.QueryHostRequest
	(*DescribeHostRequest)(nil), // 3: pb.DescribeHostRequest
	(*UpdateHostRequest)(nil),   // 4: pb.UpdateHostRequest
	(*DeleteHostRequest)(nil),   // 5: pb.DeleteHostRequest
	(*Host)(nil),                // 6: pb.Host
	(*Resource)(nil),            // 7: pb.Resource
	(*Describe)(nil),            // 8: pb.Describe
	(*Set)(nil),                 // 9: pb.Set
	nil,                         // 10: pb.Resource.TagsEntry
}
var file_apps_host_pb_host_proto_depIdxs = []int32{
	0,  // 0: pb.UpdateHostRequest.update_mode:type_name -> pb.UpdateMode
	6,  // 1: pb.UpdateHostRequest.host:type_name -> pb.Host
	7,  // 2: pb.Host.resource:type_name -> pb.Resource
	8,  // 3: pb.Host.describe:type_name -> pb.Describe
	1,  // 4: pb.Resource.vendor:type_name -> pb.Vendor
	10, // 5: pb.Resource.tags:type_name -> pb.Resource.TagsEntry
	6,  // 6: pb.Set.items:type_name -> pb.Host
	6,  // 7: pb.Service.CreateHost:input_type -> pb.Host
	2,  // 8: pb.Service.QueryHost:input_type -> pb.QueryHostRequest
	3,  // 9: pb.Service.DescribeHost:input_type -> pb.DescribeHostRequest
	4,  // 10: pb.Service.UpdateHost:input_type -> pb.UpdateHostRequest
	5,  // 11: pb.Service.DeleteHost:input_type -> pb.DeleteHostRequest
	6,  // 12: pb.Service.CreateHost:output_type -> pb.Host
	9,  // 13: pb.Service.QueryHost:output_type -> pb.Set
	6,  // 14: pb.Service.DescribeHost:output_type -> pb.Host
	6,  // 15: pb.Service.UpdateHost:output_type -> pb.Host
	6,  // 16: pb.Service.DeleteHost:output_type -> pb.Host
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_apps_host_pb_host_proto_init() }
func file_apps_host_pb_host_proto_init() {
	if File_apps_host_pb_host_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_host_pb_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryHostRequest); i {
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
		file_apps_host_pb_host_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeHostRequest); i {
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
		file_apps_host_pb_host_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHostRequest); i {
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
		file_apps_host_pb_host_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHostRequest); i {
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
		file_apps_host_pb_host_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
		file_apps_host_pb_host_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_apps_host_pb_host_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Describe); i {
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
		file_apps_host_pb_host_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
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
			RawDescriptor: file_apps_host_pb_host_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_host_pb_host_proto_goTypes,
		DependencyIndexes: file_apps_host_pb_host_proto_depIdxs,
		EnumInfos:         file_apps_host_pb_host_proto_enumTypes,
		MessageInfos:      file_apps_host_pb_host_proto_msgTypes,
	}.Build()
	File_apps_host_pb_host_proto = out.File
	file_apps_host_pb_host_proto_rawDesc = nil
	file_apps_host_pb_host_proto_goTypes = nil
	file_apps_host_pb_host_proto_depIdxs = nil
}
