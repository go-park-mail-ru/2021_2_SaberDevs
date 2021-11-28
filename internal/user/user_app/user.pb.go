// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user_server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Login                string   `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Surname              string   `protobuf:"bytes,3,opt,name=Surname,proto3" json:"Surname,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	Score                int32    `protobuf:"varint,6,opt,name=Score,proto3" json:"Score,omitempty"`
	AvatarURL            string   `protobuf:"bytes,7,opt,name=AvatarURL,proto3" json:"AvatarURL,omitempty"`
	Description          string   `protobuf:"bytes,8,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *User) GetAvatarURL() string {
	if m != nil {
		return m.AvatarURL
	}
	return ""
}

func (m *User) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type LoginData struct {
	Login                string   `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Surname              string   `protobuf:"bytes,3,opt,name=Surname,proto3" json:"Surname,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Score                int32    `protobuf:"varint,5,opt,name=Score,proto3" json:"Score,omitempty"`
	AvatarURL            string   `protobuf:"bytes,6,opt,name=AvatarURL,proto3" json:"AvatarURL,omitempty"`
	Description          string   `protobuf:"bytes,7,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginData) Reset()         { *m = LoginData{} }
func (m *LoginData) String() string { return proto.CompactTextString(m) }
func (*LoginData) ProtoMessage()    {}
func (*LoginData) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *LoginData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginData.Unmarshal(m, b)
}
func (m *LoginData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginData.Marshal(b, m, deterministic)
}
func (m *LoginData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginData.Merge(m, src)
}
func (m *LoginData) XXX_Size() int {
	return xxx_messageInfo_LoginData.Size(m)
}
func (m *LoginData) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginData.DiscardUnknown(m)
}

var xxx_messageInfo_LoginData proto.InternalMessageInfo

func (m *LoginData) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *LoginData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginData) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *LoginData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginData) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *LoginData) GetAvatarURL() string {
	if m != nil {
		return m.AvatarURL
	}
	return ""
}

func (m *LoginData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type LoginResponse struct {
	Status               uint32     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 *LoginData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Msg                  string     `protobuf:"bytes,3,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *LoginResponse) GetData() *LoginData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *LoginResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type UpdateInput struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	SesionID             string   `protobuf:"bytes,2,opt,name=sesionID,proto3" json:"sesionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateInput) Reset()         { *m = UpdateInput{} }
func (m *UpdateInput) String() string { return proto.CompactTextString(m) }
func (*UpdateInput) ProtoMessage()    {}
func (*UpdateInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UpdateInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateInput.Unmarshal(m, b)
}
func (m *UpdateInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateInput.Marshal(b, m, deterministic)
}
func (m *UpdateInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInput.Merge(m, src)
}
func (m *UpdateInput) XXX_Size() int {
	return xxx_messageInfo_UpdateInput.Size(m)
}
func (m *UpdateInput) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInput.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInput proto.InternalMessageInfo

func (m *UpdateInput) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateInput) GetSesionID() string {
	if m != nil {
		return m.SesionID
	}
	return ""
}

type LogoutResponse struct {
	Status               uint32   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	GoodbyeMsg           string   `protobuf:"bytes,2,opt,name=GoodbyeMsg,proto3" json:"GoodbyeMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

func (m *LogoutResponse) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *LogoutResponse) GetGoodbyeMsg() string {
	if m != nil {
		return m.GoodbyeMsg
	}
	return ""
}

type GetUserData struct {
	Login                string   `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Surname              string   `protobuf:"bytes,3,opt,name=Surname,proto3" json:"Surname,omitempty"`
	Score                int32    `protobuf:"varint,4,opt,name=Score,proto3" json:"Score,omitempty"`
	AvatarURL            string   `protobuf:"bytes,5,opt,name=AvatarURL,proto3" json:"AvatarURL,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserData) Reset()         { *m = GetUserData{} }
func (m *GetUserData) String() string { return proto.CompactTextString(m) }
func (*GetUserData) ProtoMessage()    {}
func (*GetUserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *GetUserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserData.Unmarshal(m, b)
}
func (m *GetUserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserData.Marshal(b, m, deterministic)
}
func (m *GetUserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserData.Merge(m, src)
}
func (m *GetUserData) XXX_Size() int {
	return xxx_messageInfo_GetUserData.Size(m)
}
func (m *GetUserData) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserData.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserData proto.InternalMessageInfo

func (m *GetUserData) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *GetUserData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetUserData) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *GetUserData) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *GetUserData) GetAvatarURL() string {
	if m != nil {
		return m.AvatarURL
	}
	return ""
}

func (m *GetUserData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type GetUserResponse struct {
	Status               uint32       `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 *GetUserData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Msg                  string       `protobuf:"bytes,3,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetUserResponse) GetData() *GetUserData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetUserResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type SignUpData struct {
	Login                string   `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Surname              string   `protobuf:"bytes,3,opt,name=Surname,proto3" json:"Surname,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Score                int32    `protobuf:"varint,5,opt,name=Score,proto3" json:"Score,omitempty"`
	AvatarURL            string   `protobuf:"bytes,6,opt,name=AvatarURL,proto3" json:"AvatarURL,omitempty"`
	Description          string   `protobuf:"bytes,7,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpData) Reset()         { *m = SignUpData{} }
func (m *SignUpData) String() string { return proto.CompactTextString(m) }
func (*SignUpData) ProtoMessage()    {}
func (*SignUpData) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *SignUpData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpData.Unmarshal(m, b)
}
func (m *SignUpData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpData.Marshal(b, m, deterministic)
}
func (m *SignUpData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpData.Merge(m, src)
}
func (m *SignUpData) XXX_Size() int {
	return xxx_messageInfo_SignUpData.Size(m)
}
func (m *SignUpData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpData.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpData proto.InternalMessageInfo

func (m *SignUpData) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *SignUpData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignUpData) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *SignUpData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpData) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *SignUpData) GetAvatarURL() string {
	if m != nil {
		return m.AvatarURL
	}
	return ""
}

func (m *SignUpData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type SignupResponse struct {
	Status               uint32      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 *SignUpData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Msg                  string      `protobuf:"bytes,3,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SignupResponse) Reset()         { *m = SignupResponse{} }
func (m *SignupResponse) String() string { return proto.CompactTextString(m) }
func (*SignupResponse) ProtoMessage()    {}
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *SignupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupResponse.Unmarshal(m, b)
}
func (m *SignupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupResponse.Marshal(b, m, deterministic)
}
func (m *SignupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupResponse.Merge(m, src)
}
func (m *SignupResponse) XXX_Size() int {
	return xxx_messageInfo_SignupResponse.Size(m)
}
func (m *SignupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignupResponse proto.InternalMessageInfo

func (m *SignupResponse) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SignupResponse) GetData() *SignUpData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SignupResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type SessionID struct {
	SesionID             string   `protobuf:"bytes,1,opt,name=sesionID,proto3" json:"sesionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionID) Reset()         { *m = SessionID{} }
func (m *SessionID) String() string { return proto.CompactTextString(m) }
func (*SessionID) ProtoMessage()    {}
func (*SessionID) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *SessionID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionID.Unmarshal(m, b)
}
func (m *SessionID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionID.Marshal(b, m, deterministic)
}
func (m *SessionID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionID.Merge(m, src)
}
func (m *SessionID) XXX_Size() int {
	return xxx_messageInfo_SessionID.Size(m)
}
func (m *SessionID) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionID.DiscardUnknown(m)
}

var xxx_messageInfo_SessionID proto.InternalMessageInfo

func (m *SessionID) GetSesionID() string {
	if m != nil {
		return m.SesionID
	}
	return ""
}

type CookieValue struct {
	CookieValue          string   `protobuf:"bytes,1,opt,name=cookieValue,proto3" json:"cookieValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CookieValue) Reset()         { *m = CookieValue{} }
func (m *CookieValue) String() string { return proto.CompactTextString(m) }
func (*CookieValue) ProtoMessage()    {}
func (*CookieValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *CookieValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookieValue.Unmarshal(m, b)
}
func (m *CookieValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookieValue.Marshal(b, m, deterministic)
}
func (m *CookieValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookieValue.Merge(m, src)
}
func (m *CookieValue) XXX_Size() int {
	return xxx_messageInfo_CookieValue.Size(m)
}
func (m *CookieValue) XXX_DiscardUnknown() {
	xxx_messageInfo_CookieValue.DiscardUnknown(m)
}

var xxx_messageInfo_CookieValue proto.InternalMessageInfo

func (m *CookieValue) GetCookieValue() string {
	if m != nil {
		return m.CookieValue
	}
	return ""
}

type Author struct {
	Author               string   `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{11}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type Nothing struct {
	Dummy                bool     `protobuf:"varint,1,opt,name=dummy,proto3" json:"dummy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{12}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func (m *Nothing) GetDummy() bool {
	if m != nil {
		return m.Dummy
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "user_server.User")
	proto.RegisterType((*LoginData)(nil), "user_server.LoginData")
	proto.RegisterType((*LoginResponse)(nil), "user_server.LoginResponse")
	proto.RegisterType((*UpdateInput)(nil), "user_server.UpdateInput")
	proto.RegisterType((*LogoutResponse)(nil), "user_server.LogoutResponse")
	proto.RegisterType((*GetUserData)(nil), "user_server.GetUserData")
	proto.RegisterType((*GetUserResponse)(nil), "user_server.GetUserResponse")
	proto.RegisterType((*SignUpData)(nil), "user_server.SignUpData")
	proto.RegisterType((*SignupResponse)(nil), "user_server.SignupResponse")
	proto.RegisterType((*SessionID)(nil), "user_server.SessionID")
	proto.RegisterType((*CookieValue)(nil), "user_server.CookieValue")
	proto.RegisterType((*Author)(nil), "user_server.Author")
	proto.RegisterType((*Nothing)(nil), "user_server.Nothing")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0x68, 0x9a, 0xb6, 0x37, 0xb4, 0x0c, 0x33, 0x8d, 0xa8, 0x4c, 0x50, 0x45, 0x42, 0x4c,
	0x80, 0x8a, 0x54, 0xde, 0x10, 0x2f, 0x13, 0x45, 0xdd, 0xa4, 0x32, 0x55, 0xa9, 0xca, 0x2b, 0xf2,
	0xda, 0x4b, 0x16, 0xd1, 0xc6, 0x91, 0xed, 0x14, 0xf5, 0x6f, 0x78, 0xe3, 0x23, 0xe0, 0x2b, 0xf8,
	0x22, 0x64, 0x3b, 0x6d, 0x93, 0x75, 0x55, 0x5e, 0xe0, 0x61, 0x6f, 0xbe, 0xd7, 0xbe, 0x37, 0xe7,
	0x9e, 0x73, 0xec, 0x00, 0xa4, 0x02, 0x79, 0x37, 0xe1, 0x4c, 0x32, 0xe2, 0xaa, 0xf5, 0x17, 0x81,
	0x7c, 0x89, 0xdc, 0xff, 0x63, 0x81, 0x3d, 0x11, 0xc8, 0xc9, 0x11, 0x54, 0x87, 0x2c, 0x8c, 0x62,
	0xcf, 0xea, 0x58, 0xa7, 0x8d, 0xc0, 0x04, 0x84, 0x80, 0x7d, 0x49, 0x17, 0xe8, 0xdd, 0xd3, 0x49,
	0xbd, 0x26, 0x1e, 0xd4, 0xc6, 0x29, 0x8f, 0x55, 0xba, 0xa2, 0xd3, 0xeb, 0x50, 0xf5, 0xf8, 0xb8,
	0xa0, 0xd1, 0xdc, 0xb3, 0x4d, 0x0f, 0x1d, 0x90, 0x36, 0xd4, 0x47, 0x54, 0x88, 0xef, 0x8c, 0xcf,
	0xbc, 0xaa, 0xde, 0xd8, 0xc4, 0xaa, 0x62, 0x3c, 0x65, 0x1c, 0x3d, 0xa7, 0x63, 0x9d, 0x56, 0x03,
	0x13, 0x90, 0x13, 0x68, 0x9c, 0x2d, 0xa9, 0xa4, 0x7c, 0x12, 0x0c, 0xbd, 0x9a, 0x2e, 0xd9, 0x26,
	0x48, 0x07, 0xdc, 0x3e, 0x8a, 0x29, 0x8f, 0x12, 0x19, 0xb1, 0xd8, 0xab, 0xeb, 0xfd, 0x7c, 0xca,
	0xff, 0x65, 0x41, 0x43, 0xe3, 0xef, 0x53, 0x49, 0xff, 0xe3, 0x64, 0x1b, 0xf4, 0xd5, 0xbd, 0xe8,
	0x9d, 0x12, 0xf4, 0xb5, 0x5d, 0xf4, 0x08, 0x4d, 0x0d, 0x31, 0x40, 0x91, 0xb0, 0x58, 0x20, 0x39,
	0x06, 0x47, 0x48, 0x2a, 0x53, 0xa1, 0x27, 0x68, 0x06, 0x59, 0x44, 0x5e, 0x82, 0x3d, 0xa3, 0x92,
	0xea, 0x11, 0xdc, 0xde, 0x71, 0x37, 0xa7, 0x6b, 0x77, 0x33, 0x7e, 0xa0, 0xcf, 0x90, 0x43, 0xa8,
	0x7c, 0x12, 0x61, 0x36, 0x96, 0x5a, 0xfa, 0x23, 0x70, 0x27, 0xc9, 0x8c, 0x4a, 0xbc, 0x88, 0x93,
	0x54, 0x92, 0xe7, 0x60, 0xab, 0x7a, 0xfd, 0x09, 0xb7, 0xf7, 0xb0, 0xd0, 0x4c, 0x19, 0x24, 0xd0,
	0xdb, 0x4a, 0x4c, 0x81, 0x22, 0x62, 0xf1, 0x45, 0x3f, 0xa3, 0x6e, 0x13, 0xfb, 0xe7, 0xd0, 0x1a,
	0xb2, 0x90, 0xa5, 0xb2, 0x14, 0xf9, 0x53, 0x80, 0x01, 0x63, 0xb3, 0xab, 0x15, 0x2a, 0x50, 0xa6,
	0x4f, 0x2e, 0xe3, 0xff, 0xb4, 0xc0, 0x1d, 0xa0, 0x54, 0xdf, 0xfd, 0x97, 0x12, 0x1a, 0xb1, 0xec,
	0xbd, 0x62, 0x55, 0x4b, 0xc4, 0x72, 0x76, 0xc5, 0x8a, 0xe0, 0x41, 0x06, 0xb4, 0x74, 0xe8, 0xd7,
	0x05, 0xb9, 0xbc, 0x02, 0xc3, 0xb9, 0x61, 0xf7, 0x0a, 0xf6, 0xdb, 0x02, 0x18, 0x47, 0x61, 0x3c,
	0x49, 0xee, 0xa4, 0xad, 0x43, 0x68, 0x29, 0xf4, 0x69, 0x52, 0x4a, 0xd4, 0xab, 0x02, 0x51, 0x8f,
	0x0b, 0x44, 0x6d, 0x09, 0xd8, 0xcb, 0xd3, 0x0b, 0x68, 0x8c, 0x51, 0x18, 0x4f, 0x16, 0xfc, 0x6a,
	0xdd, 0xf0, 0xeb, 0x1b, 0x70, 0x3f, 0x30, 0xf6, 0x2d, 0xc2, 0xcf, 0x74, 0x9e, 0xa2, 0x1a, 0x61,
	0xba, 0x0d, 0xb3, 0xd3, 0xf9, 0x94, 0xdf, 0x01, 0xe7, 0x2c, 0x95, 0xd7, 0x8c, 0x2b, 0xe8, 0x54,
	0xaf, 0xb2, 0x63, 0x59, 0xe4, 0x3f, 0x83, 0xda, 0x25, 0x93, 0xd7, 0x51, 0x1c, 0x2a, 0x16, 0x67,
	0xe9, 0x62, 0xb1, 0xd2, 0x27, 0xea, 0x81, 0x09, 0x7a, 0x3f, 0x2a, 0x70, 0x5f, 0x2b, 0x8d, 0xf3,
	0x68, 0x89, 0x7c, 0x45, 0x06, 0xd0, 0x34, 0xd7, 0x70, 0xc4, 0xd9, 0xd7, 0x68, 0x8e, 0xa4, 0x68,
	0x8c, 0xdc, 0x15, 0x6d, 0xb7, 0x77, 0x6f, 0xf8, 0x9a, 0x4b, 0xff, 0x80, 0x0c, 0xe0, 0x70, 0x80,
	0xd2, 0xe0, 0x5b, 0xf7, 0x7a, 0x54, 0xa8, 0x30, 0x7b, 0xed, 0x93, 0xdb, 0x9c, 0x97, 0x6b, 0x74,
	0x0e, 0xad, 0x2c, 0xb9, 0x6e, 0x53, 0x7c, 0x5a, 0x36, 0xe4, 0x96, 0x76, 0x7a, 0x9f, 0x3d, 0xc3,
	0xfa, 0x07, 0xb3, 0xfb, 0xa4, 0x94, 0x0c, 0xf4, 0x0e, 0x1c, 0x63, 0x98, 0xdb, 0x4a, 0x9f, 0xec,
	0xb8, 0x62, 0x6b, 0x2c, 0x53, 0x6b, 0x9e, 0xa2, 0x1b, 0x74, 0xe6, 0xf4, 0x6e, 0x1f, 0x15, 0x76,
	0x32, 0xd9, 0xfc, 0x83, 0x2b, 0x47, 0xff, 0x26, 0xdf, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x82,
	0xc8, 0x43, 0xb5, 0x34, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserDeliveryClient is the client API for UserDelivery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserDeliveryClient interface {
	UpdateProfile(ctx context.Context, in *UpdateInput, opts ...grpc.CallOption) (*LoginResponse, error)
	GetAuthorProfile(ctx context.Context, in *Author, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetUserProfile(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*GetUserResponse, error)
	LoginUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*LoginResponse, error)
	Signup(ctx context.Context, in *User, opts ...grpc.CallOption) (*SignupResponse, error)
	Logout(ctx context.Context, in *CookieValue, opts ...grpc.CallOption) (*Nothing, error)
}

type userDeliveryClient struct {
	cc *grpc.ClientConn
}

func NewUserDeliveryClient(cc *grpc.ClientConn) UserDeliveryClient {
	return &userDeliveryClient{cc}
}

func (c *userDeliveryClient) UpdateProfile(ctx context.Context, in *UpdateInput, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user_server.UserDelivery/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDeliveryClient) GetAuthorProfile(ctx context.Context, in *Author, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/user_server.UserDelivery/GetAuthorProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDeliveryClient) GetUserProfile(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/user_server.UserDelivery/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDeliveryClient) LoginUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user_server.UserDelivery/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDeliveryClient) Signup(ctx context.Context, in *User, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/user_server.UserDelivery/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDeliveryClient) Logout(ctx context.Context, in *CookieValue, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/user_server.UserDelivery/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDeliveryServer is the server API for UserDelivery service.
type UserDeliveryServer interface {
	UpdateProfile(context.Context, *UpdateInput) (*LoginResponse, error)
	GetAuthorProfile(context.Context, *Author) (*GetUserResponse, error)
	GetUserProfile(context.Context, *SessionID) (*GetUserResponse, error)
	LoginUser(context.Context, *User) (*LoginResponse, error)
	Signup(context.Context, *User) (*SignupResponse, error)
	Logout(context.Context, *CookieValue) (*Nothing, error)
}

// UnimplementedUserDeliveryServer can be embedded to have forward compatible implementations.
type UnimplementedUserDeliveryServer struct {
}

func (*UnimplementedUserDeliveryServer) UpdateProfile(ctx context.Context, req *UpdateInput) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (*UnimplementedUserDeliveryServer) GetAuthorProfile(ctx context.Context, req *Author) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorProfile not implemented")
}
func (*UnimplementedUserDeliveryServer) GetUserProfile(ctx context.Context, req *SessionID) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (*UnimplementedUserDeliveryServer) LoginUser(ctx context.Context, req *User) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (*UnimplementedUserDeliveryServer) Signup(ctx context.Context, req *User) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (*UnimplementedUserDeliveryServer) Logout(ctx context.Context, req *CookieValue) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

func RegisterUserDeliveryServer(s *grpc.Server, srv UserDeliveryServer) {
	s.RegisterService(&_UserDelivery_serviceDesc, srv)
}

func _UserDelivery_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDeliveryServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_server.UserDelivery/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDeliveryServer).UpdateProfile(ctx, req.(*UpdateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDelivery_GetAuthorProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDeliveryServer).GetAuthorProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_server.UserDelivery/GetAuthorProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDeliveryServer).GetAuthorProfile(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDelivery_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDeliveryServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_server.UserDelivery/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDeliveryServer).GetUserProfile(ctx, req.(*SessionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDelivery_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDeliveryServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_server.UserDelivery/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDeliveryServer).LoginUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDelivery_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDeliveryServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_server.UserDelivery/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDeliveryServer).Signup(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDelivery_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CookieValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDeliveryServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_server.UserDelivery/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDeliveryServer).Logout(ctx, req.(*CookieValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserDelivery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user_server.UserDelivery",
	HandlerType: (*UserDeliveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateProfile",
			Handler:    _UserDelivery_UpdateProfile_Handler,
		},
		{
			MethodName: "GetAuthorProfile",
			Handler:    _UserDelivery_GetAuthorProfile_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _UserDelivery_GetUserProfile_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _UserDelivery_LoginUser_Handler,
		},
		{
			MethodName: "Signup",
			Handler:    _UserDelivery_Signup_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserDelivery_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
