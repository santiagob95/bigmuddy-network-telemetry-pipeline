// Code generated by protoc-gen-go.
// source: snmp_snmpHostXml.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_snmp_agent_oper_snmp_information_hosts_host_host_information is a generated protocol buffer package.

It is generated from these files:
	snmp_snmpHostXml.proto

It has these top-level messages:
	SnmpSnmpHostXml_KEYS
	SnmpSnmpHostXml
*/
package cisco_ios_xr_snmp_agent_oper_snmp_information_hosts_host_host_information

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SNMP host xml information
type SnmpSnmpHostXml_KEYS struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *SnmpSnmpHostXml_KEYS) Reset()                    { *m = SnmpSnmpHostXml_KEYS{} }
func (m *SnmpSnmpHostXml_KEYS) String() string            { return proto.CompactTextString(m) }
func (*SnmpSnmpHostXml_KEYS) ProtoMessage()               {}
func (*SnmpSnmpHostXml_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SnmpSnmpHostXml_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SnmpSnmpHostXml_KEYS) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type SnmpSnmpHostXml struct {
	// Transport type of address
	SnmpTargetAddressTHost string `protobuf:"bytes,50,opt,name=snmp_target_address_t_host,json=snmpTargetAddressTHost" json:"snmp_target_address_t_host,omitempty"`
	// Target UDP port
	SnmpTargetAddressPort string `protobuf:"bytes,51,opt,name=snmp_target_address_port,json=snmpTargetAddressPort" json:"snmp_target_address_port,omitempty"`
	// Target host type (Inform or Trap)
	SnmpTargetAddresstype string `protobuf:"bytes,52,opt,name=snmp_target_addresstype,json=snmpTargetAddresstype" json:"snmp_target_addresstype,omitempty"`
	// Security model
	SnmpTargetParamsSecurityModel string `protobuf:"bytes,53,opt,name=snmp_target_params_security_model,json=snmpTargetParamsSecurityModel" json:"snmp_target_params_security_model,omitempty"`
	// Security name
	SnmpTargetParamsSecurityName string `protobuf:"bytes,54,opt,name=snmp_target_params_security_name,json=snmpTargetParamsSecurityName" json:"snmp_target_params_security_name,omitempty"`
	// Security level
	SnmpTargetParamsSecurityLevel string `protobuf:"bytes,55,opt,name=snmp_target_params_security_level,json=snmpTargetParamsSecurityLevel" json:"snmp_target_params_security_level,omitempty"`
}

func (m *SnmpSnmpHostXml) Reset()                    { *m = SnmpSnmpHostXml{} }
func (m *SnmpSnmpHostXml) String() string            { return proto.CompactTextString(m) }
func (*SnmpSnmpHostXml) ProtoMessage()               {}
func (*SnmpSnmpHostXml) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SnmpSnmpHostXml) GetSnmpTargetAddressTHost() string {
	if m != nil {
		return m.SnmpTargetAddressTHost
	}
	return ""
}

func (m *SnmpSnmpHostXml) GetSnmpTargetAddressPort() string {
	if m != nil {
		return m.SnmpTargetAddressPort
	}
	return ""
}

func (m *SnmpSnmpHostXml) GetSnmpTargetAddresstype() string {
	if m != nil {
		return m.SnmpTargetAddresstype
	}
	return ""
}

func (m *SnmpSnmpHostXml) GetSnmpTargetParamsSecurityModel() string {
	if m != nil {
		return m.SnmpTargetParamsSecurityModel
	}
	return ""
}

func (m *SnmpSnmpHostXml) GetSnmpTargetParamsSecurityName() string {
	if m != nil {
		return m.SnmpTargetParamsSecurityName
	}
	return ""
}

func (m *SnmpSnmpHostXml) GetSnmpTargetParamsSecurityLevel() string {
	if m != nil {
		return m.SnmpTargetParamsSecurityLevel
	}
	return ""
}

func init() {
	proto.RegisterType((*SnmpSnmpHostXml_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.hosts.host.host_information.snmp_snmpHostXml_KEYS")
	proto.RegisterType((*SnmpSnmpHostXml)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.hosts.host.host_information.snmp_snmpHostXml")
}

func init() { proto.RegisterFile("snmp_snmpHostXml.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xa9, 0x8a, 0x60, 0x4e, 0x12, 0xa8, 0x06, 0x51, 0xa8, 0x3d, 0x79, 0xda, 0x83, 0xd5,
	0x16, 0xbc, 0x88, 0x07, 0xa5, 0xe2, 0x1f, 0x8a, 0xed, 0x41, 0x4f, 0x43, 0xdc, 0x8d, 0x75, 0x61,
	0xb3, 0x13, 0x32, 0x53, 0xb1, 0x9f, 0xc5, 0x2f, 0x2b, 0x99, 0x0a, 0x96, 0xda, 0xd6, 0xcb, 0x90,
	0x9d, 0xf7, 0x7e, 0x2f, 0xb3, 0x43, 0xd4, 0x1e, 0xd5, 0x3e, 0x40, 0x2a, 0x7d, 0x24, 0x7e, 0xf6,
	0x55, 0x16, 0x22, 0x32, 0xea, 0xdb, 0xbc, 0xa4, 0x1c, 0xa1, 0x44, 0x82, 0xcf, 0x28, 0x3a, 0xd8,
	0xb1, 0xab, 0x19, 0x30, 0xb8, 0x98, 0xa5, 0xef, 0xac, 0xac, 0xdf, 0x30, 0x7a, 0xcb, 0x25, 0xd6,
	0xd9, 0x3b, 0x12, 0x93, 0x54, 0x29, 0x30, 0xa7, 0xb5, 0x2f, 0x55, 0x73, 0xf1, 0x12, 0xb8, 0xbb,
	0x7e, 0x19, 0x6a, 0xad, 0xb6, 0x6a, 0xeb, 0x9d, 0x69, 0xb4, 0x1a, 0x27, 0x3b, 0x4f, 0x72, 0x4e,
	0xbd, 0x09, 0xb9, 0x68, 0x36, 0x66, 0xbd, 0x74, 0x6e, 0x7f, 0x6d, 0xaa, 0xdd, 0xc5, 0x04, 0x7d,
	0xa1, 0x0e, 0xa4, 0xc7, 0x36, 0x8e, 0x1d, 0x83, 0x2d, 0x8a, 0xe8, 0x88, 0x80, 0x21, 0xdd, 0x6f,
	0x4e, 0x05, 0x97, 0x9f, 0x1b, 0x89, 0xe1, 0x6a, 0xa6, 0x8f, 0x12, 0xae, 0x7b, 0xca, 0x2c, 0x63,
	0x03, 0x46, 0x36, 0x1d, 0x21, 0x9b, 0x7f, 0xc8, 0x01, 0x46, 0xd6, 0x5d, 0xb5, 0xbf, 0x04, 0xe4,
	0x69, 0x70, 0xe6, 0x6c, 0x05, 0x97, 0x44, 0xdd, 0x57, 0xc7, 0xf3, 0x5c, 0xb0, 0xd1, 0x7a, 0x02,
	0x72, 0xf9, 0x24, 0x96, 0x3c, 0x05, 0x8f, 0x85, 0xab, 0xcc, 0xb9, 0x24, 0x1c, 0xfd, 0x26, 0x0c,
	0xc4, 0x36, 0xfc, 0x71, 0x3d, 0x24, 0x93, 0xbe, 0x51, 0xad, 0x75, 0x49, 0xb2, 0xcf, 0xae, 0x04,
	0x1d, 0xae, 0x0a, 0x7a, 0x4c, 0x7b, 0xfe, 0x67, 0xa2, 0xca, 0x7d, 0xb8, 0xca, 0xf4, 0xd6, 0x4f,
	0x74, 0x9f, 0x4c, 0xaf, 0xdb, 0xf2, 0x60, 0x3a, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x92, 0x69,
	0x28, 0x88, 0x4a, 0x02, 0x00, 0x00,
}