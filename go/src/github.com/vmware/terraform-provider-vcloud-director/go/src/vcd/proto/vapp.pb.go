// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vapp.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateVAppInfo struct {
	Name             string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	CatalogName      string `protobuf:"bytes,2,opt,name=catalog_name,json=catalogName" json:"catalog_name,omitempty"`
	TemplateName     string `protobuf:"bytes,3,opt,name=template_name,json=templateName" json:"template_name,omitempty"`
	Vdc              string `protobuf:"bytes,4,opt,name=vdc" json:"vdc,omitempty"`
	Network          string `protobuf:"bytes,5,opt,name=network" json:"network,omitempty"`
	IpAllocationMode string `protobuf:"bytes,6,opt,name=ip_allocation_mode,json=ipAllocationMode" json:"ip_allocation_mode,omitempty"`
	Memory           string `protobuf:"bytes,7,opt,name=memory" json:"memory,omitempty"`
	Cpu              int32  `protobuf:"varint,8,opt,name=cpu" json:"cpu,omitempty"`
	PowerOn          bool   `protobuf:"varint,9,opt,name=power_on,json=powerOn" json:"power_on,omitempty"`
	StorageProfile   string `protobuf:"bytes,10,opt,name=storage_profile,json=storageProfile" json:"storage_profile,omitempty"`
	AcceptAllEulas   bool   `protobuf:"varint,11,opt,name=accept_all_eulas,json=acceptAllEulas" json:"accept_all_eulas,omitempty"`
}

func (m *CreateVAppInfo) Reset()                    { *m = CreateVAppInfo{} }
func (m *CreateVAppInfo) String() string            { return proto1.CompactTextString(m) }
func (*CreateVAppInfo) ProtoMessage()               {}
func (*CreateVAppInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateVAppInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateVAppInfo) GetCatalogName() string {
	if m != nil {
		return m.CatalogName
	}
	return ""
}

func (m *CreateVAppInfo) GetTemplateName() string {
	if m != nil {
		return m.TemplateName
	}
	return ""
}

func (m *CreateVAppInfo) GetVdc() string {
	if m != nil {
		return m.Vdc
	}
	return ""
}

func (m *CreateVAppInfo) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *CreateVAppInfo) GetIpAllocationMode() string {
	if m != nil {
		return m.IpAllocationMode
	}
	return ""
}

func (m *CreateVAppInfo) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func (m *CreateVAppInfo) GetCpu() int32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *CreateVAppInfo) GetPowerOn() bool {
	if m != nil {
		return m.PowerOn
	}
	return false
}

func (m *CreateVAppInfo) GetStorageProfile() string {
	if m != nil {
		return m.StorageProfile
	}
	return ""
}

func (m *CreateVAppInfo) GetAcceptAllEulas() bool {
	if m != nil {
		return m.AcceptAllEulas
	}
	return false
}

type CreateVAppResult struct {
	Created    bool            `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	InVappInfo *CreateVAppInfo `protobuf:"bytes,2,opt,name=in_vapp_info,json=inVappInfo" json:"in_vapp_info,omitempty"`
}

func (m *CreateVAppResult) Reset()                    { *m = CreateVAppResult{} }
func (m *CreateVAppResult) String() string            { return proto1.CompactTextString(m) }
func (*CreateVAppResult) ProtoMessage()               {}
func (*CreateVAppResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CreateVAppResult) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *CreateVAppResult) GetInVappInfo() *CreateVAppInfo {
	if m != nil {
		return m.InVappInfo
	}
	return nil
}

type DeleteVAppInfo struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Vdc  string `protobuf:"bytes,2,opt,name=vdc" json:"vdc,omitempty"`
}

func (m *DeleteVAppInfo) Reset()                    { *m = DeleteVAppInfo{} }
func (m *DeleteVAppInfo) String() string            { return proto1.CompactTextString(m) }
func (*DeleteVAppInfo) ProtoMessage()               {}
func (*DeleteVAppInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *DeleteVAppInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteVAppInfo) GetVdc() string {
	if m != nil {
		return m.Vdc
	}
	return ""
}

type DeleteVAppResult struct {
	Deleted bool `protobuf:"varint,1,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *DeleteVAppResult) Reset()                    { *m = DeleteVAppResult{} }
func (m *DeleteVAppResult) String() string            { return proto1.CompactTextString(m) }
func (*DeleteVAppResult) ProtoMessage()               {}
func (*DeleteVAppResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *DeleteVAppResult) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

type ReadVAppInfo struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Vdc  string `protobuf:"bytes,2,opt,name=vdc" json:"vdc,omitempty"`
}

func (m *ReadVAppInfo) Reset()                    { *m = ReadVAppInfo{} }
func (m *ReadVAppInfo) String() string            { return proto1.CompactTextString(m) }
func (*ReadVAppInfo) ProtoMessage()               {}
func (*ReadVAppInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ReadVAppInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadVAppInfo) GetVdc() string {
	if m != nil {
		return m.Vdc
	}
	return ""
}

type ReadVAppResult struct {
	Present bool     `protobuf:"varint,1,opt,name=present" json:"present,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Network []string `protobuf:"bytes,3,rep,name=network" json:"network,omitempty"`
}

func (m *ReadVAppResult) Reset()                    { *m = ReadVAppResult{} }
func (m *ReadVAppResult) String() string            { return proto1.CompactTextString(m) }
func (*ReadVAppResult) ProtoMessage()               {}
func (*ReadVAppResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ReadVAppResult) GetPresent() bool {
	if m != nil {
		return m.Present
	}
	return false
}

func (m *ReadVAppResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadVAppResult) GetNetwork() []string {
	if m != nil {
		return m.Network
	}
	return nil
}

type UpdateVAppInfo struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Vdc     string `protobuf:"bytes,2,opt,name=vdc" json:"vdc,omitempty"`
	PowerOn bool   `protobuf:"varint,3,opt,name=power_on,json=powerOn" json:"power_on,omitempty"`
}

func (m *UpdateVAppInfo) Reset()                    { *m = UpdateVAppInfo{} }
func (m *UpdateVAppInfo) String() string            { return proto1.CompactTextString(m) }
func (*UpdateVAppInfo) ProtoMessage()               {}
func (*UpdateVAppInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *UpdateVAppInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateVAppInfo) GetVdc() string {
	if m != nil {
		return m.Vdc
	}
	return ""
}

func (m *UpdateVAppInfo) GetPowerOn() bool {
	if m != nil {
		return m.PowerOn
	}
	return false
}

type UpdateVAppResult struct {
	Updated bool `protobuf:"varint,3,opt,name=updated" json:"updated,omitempty"`
}

func (m *UpdateVAppResult) Reset()                    { *m = UpdateVAppResult{} }
func (m *UpdateVAppResult) String() string            { return proto1.CompactTextString(m) }
func (*UpdateVAppResult) ProtoMessage()               {}
func (*UpdateVAppResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *UpdateVAppResult) GetUpdated() bool {
	if m != nil {
		return m.Updated
	}
	return false
}

func init() {
	proto1.RegisterType((*CreateVAppInfo)(nil), "proto.CreateVAppInfo")
	proto1.RegisterType((*CreateVAppResult)(nil), "proto.CreateVAppResult")
	proto1.RegisterType((*DeleteVAppInfo)(nil), "proto.DeleteVAppInfo")
	proto1.RegisterType((*DeleteVAppResult)(nil), "proto.DeleteVAppResult")
	proto1.RegisterType((*ReadVAppInfo)(nil), "proto.ReadVAppInfo")
	proto1.RegisterType((*ReadVAppResult)(nil), "proto.ReadVAppResult")
	proto1.RegisterType((*UpdateVAppInfo)(nil), "proto.UpdateVAppInfo")
	proto1.RegisterType((*UpdateVAppResult)(nil), "proto.UpdateVAppResult")
}

func init() { proto1.RegisterFile("proto/vapp.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0x9b, 0xed, 0xc7, 0x4e, 0x4b, 0x88, 0x2c, 0x81, 0xcc, 0xad, 0x84, 0x03, 0x3d, 0xac,
	0x16, 0x09, 0x10, 0x9c, 0x2b, 0xe0, 0xc0, 0x81, 0xaf, 0x48, 0xac, 0xb8, 0x59, 0x26, 0x9e, 0xae,
	0x22, 0x1c, 0xdb, 0x72, 0x9d, 0x5d, 0xf1, 0x67, 0xf9, 0x2d, 0xc8, 0x93, 0xaf, 0xf6, 0xb4, 0xea,
	0x29, 0xf3, 0xde, 0xbc, 0xcc, 0x8c, 0xdf, 0x0c, 0x64, 0xce, 0xdb, 0x60, 0x5f, 0xdd, 0x49, 0xe7,
	0xae, 0x29, 0x64, 0x33, 0xfa, 0xe4, 0xff, 0xa6, 0x90, 0x7e, 0xf0, 0x28, 0x03, 0xde, 0xec, 0x9c,
	0xfb, 0x6c, 0xf6, 0x96, 0x31, 0xb8, 0x30, 0xb2, 0x46, 0x3e, 0xd9, 0x4c, 0xb6, 0x97, 0x05, 0xc5,
	0xec, 0x39, 0xac, 0x4b, 0x19, 0xa4, 0xb6, 0xb7, 0x82, 0x72, 0x53, 0xca, 0xad, 0x3a, 0xee, 0x6b,
	0x94, 0xbc, 0x80, 0x47, 0x01, 0x6b, 0xa7, 0x65, 0xc0, 0x56, 0x93, 0x90, 0x66, 0xdd, 0x93, 0x24,
	0xca, 0x20, 0xb9, 0x53, 0x25, 0xbf, 0xa0, 0x54, 0x0c, 0x19, 0x87, 0x85, 0xc1, 0x70, 0x6f, 0xfd,
	0x1f, 0x3e, 0x23, 0xb6, 0x87, 0xec, 0x0a, 0x58, 0xe5, 0x84, 0xd4, 0xda, 0x96, 0x32, 0x54, 0xd6,
	0x88, 0xda, 0x2a, 0xe4, 0x73, 0x12, 0x65, 0x95, 0xdb, 0x0d, 0x89, 0x2f, 0x56, 0x21, 0x7b, 0x0a,
	0xf3, 0x1a, 0x6b, 0xeb, 0xff, 0xf2, 0x05, 0x29, 0x3a, 0x14, 0x3b, 0x96, 0xae, 0xe1, 0xcb, 0xcd,
	0x64, 0x3b, 0x2b, 0x62, 0xc8, 0x9e, 0xc1, 0xd2, 0xd9, 0x7b, 0xf4, 0xc2, 0x1a, 0x7e, 0xb9, 0x99,
	0x6c, 0x97, 0xc5, 0x82, 0xf0, 0x37, 0xc3, 0x5e, 0xc2, 0xe3, 0x43, 0xb0, 0x5e, 0xde, 0xa2, 0x70,
	0xde, 0xee, 0x2b, 0x8d, 0x1c, 0xa8, 0x5a, 0xda, 0xd1, 0xdf, 0x5b, 0x96, 0x6d, 0x21, 0x93, 0x65,
	0x89, 0x2e, 0xc4, 0xf9, 0x04, 0x36, 0x5a, 0x1e, 0xf8, 0x8a, 0x6a, 0xa5, 0x2d, 0xbf, 0xd3, 0xfa,
	0x53, 0x64, 0x73, 0x84, 0x6c, 0xf4, 0xb7, 0xc0, 0x43, 0xa3, 0x43, 0x7c, 0x73, 0x49, 0x9c, 0x22,
	0x93, 0x97, 0x45, 0x0f, 0xd9, 0x7b, 0x58, 0x57, 0x46, 0xc4, 0x35, 0x89, 0xca, 0xec, 0x2d, 0xf9,
	0xbc, 0x7a, 0xfd, 0xa4, 0xdd, 0xd9, 0xf5, 0xe9, 0xa2, 0x0a, 0xa8, 0xcc, 0x8d, 0x6c, 0xe3, 0xfc,
	0x1d, 0xa4, 0x1f, 0x51, 0xe3, 0x03, 0x6b, 0xec, 0xec, 0x9f, 0x0e, 0xf6, 0xe7, 0x57, 0x90, 0x8d,
	0xff, 0x8d, 0xe3, 0x29, 0xe2, 0x86, 0xf1, 0x3a, 0x98, 0xbf, 0x85, 0x75, 0x81, 0x52, 0x9d, 0xd9,
	0xe3, 0x17, 0xa4, 0xfd, 0x5f, 0x63, 0x07, 0xe7, 0xf1, 0x80, 0x26, 0xf4, 0x1d, 0x3a, 0x38, 0x54,
	0x9c, 0x1e, 0x55, 0x3c, 0x3a, 0x91, 0x64, 0x93, 0x1c, 0x9d, 0x48, 0xfe, 0x03, 0xd2, 0x9f, 0x4e,
	0xc9, 0x73, 0x5f, 0x7d, 0x72, 0x02, 0xc9, 0xc9, 0x09, 0x44, 0x43, 0xc6, 0x92, 0xe3, 0xb8, 0x0d,
	0x71, 0xaa, 0x57, 0x77, 0xf0, 0xf7, 0x9c, 0x16, 0xf3, 0xe6, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xaf, 0xdf, 0xcf, 0x02, 0x60, 0x03, 0x00, 0x00,
}
