// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: core.proto

package proto

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

type Challenge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type              string              `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Status            string              `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Uri               string              `protobuf:"bytes,9,opt,name=uri,proto3" json:"uri,omitempty"`
	Token             string              `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	KeyAuthorization  string              `protobuf:"bytes,5,opt,name=keyAuthorization,proto3" json:"keyAuthorization,omitempty"`
	Validationrecords []*ValidationRecord `protobuf:"bytes,10,rep,name=validationrecords,proto3" json:"validationrecords,omitempty"`
	Error             *ProblemDetails     `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
	Validated         int64               `protobuf:"varint,11,opt,name=validated,proto3" json:"validated,omitempty"`
}

func (x *Challenge) Reset() {
	*x = Challenge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Challenge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Challenge) ProtoMessage() {}

func (x *Challenge) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Challenge.ProtoReflect.Descriptor instead.
func (*Challenge) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

func (x *Challenge) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Challenge) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Challenge) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Challenge) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Challenge) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Challenge) GetKeyAuthorization() string {
	if x != nil {
		return x.KeyAuthorization
	}
	return ""
}

func (x *Challenge) GetValidationrecords() []*ValidationRecord {
	if x != nil {
		return x.Validationrecords
	}
	return nil
}

func (x *Challenge) GetError() *ProblemDetails {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Challenge) GetValidated() int64 {
	if x != nil {
		return x.Validated
	}
	return 0
}

type ValidationRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname          string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Port              string   `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	AddressesResolved [][]byte `protobuf:"bytes,3,rep,name=addressesResolved,proto3" json:"addressesResolved,omitempty"` // net.IP.MarshalText()
	AddressUsed       []byte   `protobuf:"bytes,4,opt,name=addressUsed,proto3" json:"addressUsed,omitempty"`             // net.IP.MarshalText()
	Authorities       []string `protobuf:"bytes,5,rep,name=authorities,proto3" json:"authorities,omitempty"`
	Url               string   `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	// A list of addresses tried before the address used (see
	// core/objects.go and the comment on the ValidationRecord structure
	// definition for more information.
	AddressesTried [][]byte `protobuf:"bytes,7,rep,name=addressesTried,proto3" json:"addressesTried,omitempty"` // net.IP.MarshalText()
}

func (x *ValidationRecord) Reset() {
	*x = ValidationRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationRecord) ProtoMessage() {}

func (x *ValidationRecord) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationRecord.ProtoReflect.Descriptor instead.
func (*ValidationRecord) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1}
}

func (x *ValidationRecord) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ValidationRecord) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ValidationRecord) GetAddressesResolved() [][]byte {
	if x != nil {
		return x.AddressesResolved
	}
	return nil
}

func (x *ValidationRecord) GetAddressUsed() []byte {
	if x != nil {
		return x.AddressUsed
	}
	return nil
}

func (x *ValidationRecord) GetAuthorities() []string {
	if x != nil {
		return x.Authorities
	}
	return nil
}

func (x *ValidationRecord) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ValidationRecord) GetAddressesTried() [][]byte {
	if x != nil {
		return x.AddressesTried
	}
	return nil
}

type ProblemDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemType string `protobuf:"bytes,1,opt,name=problemType,proto3" json:"problemType,omitempty"`
	Detail      string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	HttpStatus  int32  `protobuf:"varint,3,opt,name=httpStatus,proto3" json:"httpStatus,omitempty"`
}

func (x *ProblemDetails) Reset() {
	*x = ProblemDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProblemDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemDetails) ProtoMessage() {}

func (x *ProblemDetails) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemDetails.ProtoReflect.Descriptor instead.
func (*ProblemDetails) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{2}
}

func (x *ProblemDetails) GetProblemType() string {
	if x != nil {
		return x.ProblemType
	}
	return ""
}

func (x *ProblemDetails) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *ProblemDetails) GetHttpStatus() int32 {
	if x != nil {
		return x.HttpStatus
	}
	return 0
}

type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegistrationID int64  `protobuf:"varint,1,opt,name=registrationID,proto3" json:"registrationID,omitempty"`
	Serial         string `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial,omitempty"`
	Digest         string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	Der            []byte `protobuf:"bytes,4,opt,name=der,proto3" json:"der,omitempty"`
	Issued         int64  `protobuf:"varint,5,opt,name=issued,proto3" json:"issued,omitempty"`   // Unix timestamp (nanoseconds)
	Expires        int64  `protobuf:"varint,6,opt,name=expires,proto3" json:"expires,omitempty"` // Unix timestamp (nanoseconds)
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{3}
}

func (x *Certificate) GetRegistrationID() int64 {
	if x != nil {
		return x.RegistrationID
	}
	return 0
}

func (x *Certificate) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *Certificate) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Certificate) GetDer() []byte {
	if x != nil {
		return x.Der
	}
	return nil
}

func (x *Certificate) GetIssued() int64 {
	if x != nil {
		return x.Issued
	}
	return 0
}

func (x *Certificate) GetExpires() int64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

type CertificateStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Serial                string `protobuf:"bytes,1,opt,name=serial,proto3" json:"serial,omitempty"`
	Status                string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	OcspLastUpdated       int64  `protobuf:"varint,4,opt,name=ocspLastUpdated,proto3" json:"ocspLastUpdated,omitempty"`
	RevokedDate           int64  `protobuf:"varint,5,opt,name=revokedDate,proto3" json:"revokedDate,omitempty"`
	RevokedReason         int64  `protobuf:"varint,6,opt,name=revokedReason,proto3" json:"revokedReason,omitempty"`
	LastExpirationNagSent int64  `protobuf:"varint,7,opt,name=lastExpirationNagSent,proto3" json:"lastExpirationNagSent,omitempty"`
	OcspResponse          []byte `protobuf:"bytes,8,opt,name=ocspResponse,proto3" json:"ocspResponse,omitempty"`
	NotAfter              int64  `protobuf:"varint,9,opt,name=notAfter,proto3" json:"notAfter,omitempty"`
	IsExpired             bool   `protobuf:"varint,10,opt,name=isExpired,proto3" json:"isExpired,omitempty"`
	IssuerID              int64  `protobuf:"varint,11,opt,name=issuerID,proto3" json:"issuerID,omitempty"`
}

func (x *CertificateStatus) Reset() {
	*x = CertificateStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateStatus) ProtoMessage() {}

func (x *CertificateStatus) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateStatus.ProtoReflect.Descriptor instead.
func (*CertificateStatus) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{4}
}

func (x *CertificateStatus) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *CertificateStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CertificateStatus) GetOcspLastUpdated() int64 {
	if x != nil {
		return x.OcspLastUpdated
	}
	return 0
}

func (x *CertificateStatus) GetRevokedDate() int64 {
	if x != nil {
		return x.RevokedDate
	}
	return 0
}

func (x *CertificateStatus) GetRevokedReason() int64 {
	if x != nil {
		return x.RevokedReason
	}
	return 0
}

func (x *CertificateStatus) GetLastExpirationNagSent() int64 {
	if x != nil {
		return x.LastExpirationNagSent
	}
	return 0
}

func (x *CertificateStatus) GetOcspResponse() []byte {
	if x != nil {
		return x.OcspResponse
	}
	return nil
}

func (x *CertificateStatus) GetNotAfter() int64 {
	if x != nil {
		return x.NotAfter
	}
	return 0
}

func (x *CertificateStatus) GetIsExpired() bool {
	if x != nil {
		return x.IsExpired
	}
	return false
}

func (x *CertificateStatus) GetIssuerID() int64 {
	if x != nil {
		return x.IssuerID
	}
	return 0
}

type Registration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Key             []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Contact         []string `protobuf:"bytes,3,rep,name=contact,proto3" json:"contact,omitempty"`
	ContactsPresent bool     `protobuf:"varint,4,opt,name=contactsPresent,proto3" json:"contactsPresent,omitempty"`
	Agreement       string   `protobuf:"bytes,5,opt,name=agreement,proto3" json:"agreement,omitempty"`
	InitialIP       []byte   `protobuf:"bytes,6,opt,name=initialIP,proto3" json:"initialIP,omitempty"`
	CreatedAt       int64    `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"` // Unix timestamp (nanoseconds)
	Status          string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Registration) Reset() {
	*x = Registration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registration) ProtoMessage() {}

func (x *Registration) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registration.ProtoReflect.Descriptor instead.
func (*Registration) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{5}
}

func (x *Registration) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Registration) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Registration) GetContact() []string {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *Registration) GetContactsPresent() bool {
	if x != nil {
		return x.ContactsPresent
	}
	return false
}

func (x *Registration) GetAgreement() string {
	if x != nil {
		return x.Agreement
	}
	return ""
}

func (x *Registration) GetInitialIP() []byte {
	if x != nil {
		return x.InitialIP
	}
	return nil
}

func (x *Registration) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Registration) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Authorization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Identifier     string       `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	RegistrationID int64        `protobuf:"varint,3,opt,name=registrationID,proto3" json:"registrationID,omitempty"`
	Status         string       `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Expires        int64        `protobuf:"varint,5,opt,name=expires,proto3" json:"expires,omitempty"` // Unix timestamp (nanoseconds)
	Challenges     []*Challenge `protobuf:"bytes,6,rep,name=challenges,proto3" json:"challenges,omitempty"`
}

func (x *Authorization) Reset() {
	*x = Authorization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authorization) ProtoMessage() {}

func (x *Authorization) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authorization.ProtoReflect.Descriptor instead.
func (*Authorization) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{6}
}

func (x *Authorization) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Authorization) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Authorization) GetRegistrationID() int64 {
	if x != nil {
		return x.RegistrationID
	}
	return 0
}

func (x *Authorization) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Authorization) GetExpires() int64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

func (x *Authorization) GetChallenges() []*Challenge {
	if x != nil {
		return x.Challenges
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RegistrationID    int64           `protobuf:"varint,2,opt,name=registrationID,proto3" json:"registrationID,omitempty"`
	Expires           int64           `protobuf:"varint,3,opt,name=expires,proto3" json:"expires,omitempty"`
	Error             *ProblemDetails `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	CertificateSerial string          `protobuf:"bytes,5,opt,name=certificateSerial,proto3" json:"certificateSerial,omitempty"`
	Status            string          `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Names             []string        `protobuf:"bytes,8,rep,name=names,proto3" json:"names,omitempty"`
	BeganProcessing   bool            `protobuf:"varint,9,opt,name=beganProcessing,proto3" json:"beganProcessing,omitempty"`
	Created           int64           `protobuf:"varint,10,opt,name=created,proto3" json:"created,omitempty"`
	V2Authorizations  []int64         `protobuf:"varint,11,rep,packed,name=v2Authorizations,proto3" json:"v2Authorizations,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{7}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetRegistrationID() int64 {
	if x != nil {
		return x.RegistrationID
	}
	return 0
}

func (x *Order) GetExpires() int64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

func (x *Order) GetError() *ProblemDetails {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Order) GetCertificateSerial() string {
	if x != nil {
		return x.CertificateSerial
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *Order) GetBeganProcessing() bool {
	if x != nil {
		return x.BeganProcessing
	}
	return false
}

func (x *Order) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Order) GetV2Authorizations() []int64 {
	if x != nil {
		return x.V2Authorizations
	}
	return nil
}

var File_core_proto protoreflect.FileDescriptor

var file_core_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0xab, 0x02, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x69, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x6b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x6b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x44, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x22, 0xee, 0x01, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x11, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x73,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x54, 0x72, 0x69, 0x65, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x0e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x54, 0x72, 0x69, 0x65,
	0x64, 0x22, 0x6a, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x68, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa9, 0x01,
	0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x22, 0xeb, 0x02, 0x0a, 0x11, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x28, 0x0a, 0x0f, 0x6f, 0x63, 0x73, 0x70, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6f, 0x63, 0x73, 0x70, 0x4c, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x34, 0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x67, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x67, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x63, 0x73, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f,
	0x63, 0x73, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x6f, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e,
	0x6f, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x49,
	0x44, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0xe6, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x50, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x50, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xd6, 0x01, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x0a,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x73, 0x4a, 0x04, 0x08,
	0x07, 0x10, 0x08, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x09, 0x22, 0xd7, 0x02, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x28, 0x0a,
	0x0f, 0x62, 0x65, 0x67, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x62, 0x65, 0x67, 0x61, 0x6e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x2a, 0x0a, 0x10, 0x76, 0x32, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x03, 0x52, 0x10, 0x76, 0x32, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x04, 0x08,
	0x06, 0x10, 0x07, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x65, 0x74, 0x73, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2f, 0x62, 0x6f,
	0x75, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_proto_rawDescOnce sync.Once
	file_core_proto_rawDescData = file_core_proto_rawDesc
)

func file_core_proto_rawDescGZIP() []byte {
	file_core_proto_rawDescOnce.Do(func() {
		file_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_proto_rawDescData)
	})
	return file_core_proto_rawDescData
}

var file_core_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_core_proto_goTypes = []interface{}{
	(*Challenge)(nil),         // 0: core.Challenge
	(*ValidationRecord)(nil),  // 1: core.ValidationRecord
	(*ProblemDetails)(nil),    // 2: core.ProblemDetails
	(*Certificate)(nil),       // 3: core.Certificate
	(*CertificateStatus)(nil), // 4: core.CertificateStatus
	(*Registration)(nil),      // 5: core.Registration
	(*Authorization)(nil),     // 6: core.Authorization
	(*Order)(nil),             // 7: core.Order
}
var file_core_proto_depIdxs = []int32{
	1, // 0: core.Challenge.validationrecords:type_name -> core.ValidationRecord
	2, // 1: core.Challenge.error:type_name -> core.ProblemDetails
	0, // 2: core.Authorization.challenges:type_name -> core.Challenge
	2, // 3: core.Order.error:type_name -> core.ProblemDetails
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_core_proto_init() }
func file_core_proto_init() {
	if File_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Challenge); i {
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
		file_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationRecord); i {
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
		file_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProblemDetails); i {
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
		file_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
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
		file_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateStatus); i {
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
		file_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registration); i {
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
		file_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authorization); i {
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
		file_core_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_proto_goTypes,
		DependencyIndexes: file_core_proto_depIdxs,
		MessageInfos:      file_core_proto_msgTypes,
	}.Build()
	File_core_proto = out.File
	file_core_proto_rawDesc = nil
	file_core_proto_goTypes = nil
	file_core_proto_depIdxs = nil
}
