// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RecommendationService.proto

package RecommendationService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TileClickedRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	TileId               string   `protobuf:"bytes,2,opt,name=tileId,proto3" json:"tileId,omitempty"`
	TileScore            float64  `protobuf:"fixed64,3,opt,name=tileScore,proto3" json:"tileScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *TileClickedRequest) Reset()         { *m = TileClickedRequest{} }
func (m *TileClickedRequest) String() string { return proto.CompactTextString(m) }
func (*TileClickedRequest) ProtoMessage()    {}
func (*TileClickedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcf9f213a14e9af, []int{0}
}

func (m *TileClickedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TileClickedRequest.Unmarshal(m, b)
}
func (m *TileClickedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TileClickedRequest.Marshal(b, m, deterministic)
}
func (m *TileClickedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TileClickedRequest.Merge(m, src)
}
func (m *TileClickedRequest) XXX_Size() int {
	return xxx_messageInfo_TileClickedRequest.Size(m)
}
func (m *TileClickedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TileClickedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TileClickedRequest proto.InternalMessageInfo

func (m *TileClickedRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *TileClickedRequest) GetTileId() string {
	if m != nil {
		return m.TileId
	}
	return ""
}

func (m *TileClickedRequest) GetTileScore() float64 {
	if m != nil {
		return m.TileScore
	}
	return 0
}

type InitRecommendationRequest struct {
	Genres               []string `protobuf:"bytes,1,rep,name=genres,proto3" json:"genres,omitempty"`
	Categories           []string `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
	Languages            []string `protobuf:"bytes,3,rep,name=languages,proto3" json:"languages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *InitRecommendationRequest) Reset()         { *m = InitRecommendationRequest{} }
func (m *InitRecommendationRequest) String() string { return proto.CompactTextString(m) }
func (*InitRecommendationRequest) ProtoMessage()    {}
func (*InitRecommendationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcf9f213a14e9af, []int{1}
}

func (m *InitRecommendationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRecommendationRequest.Unmarshal(m, b)
}
func (m *InitRecommendationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRecommendationRequest.Marshal(b, m, deterministic)
}
func (m *InitRecommendationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRecommendationRequest.Merge(m, src)
}
func (m *InitRecommendationRequest) XXX_Size() int {
	return xxx_messageInfo_InitRecommendationRequest.Size(m)
}
func (m *InitRecommendationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRecommendationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitRecommendationRequest proto.InternalMessageInfo

func (m *InitRecommendationRequest) GetGenres() []string {
	if m != nil {
		return m.Genres
	}
	return nil
}

func (m *InitRecommendationRequest) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *InitRecommendationRequest) GetLanguages() []string {
	if m != nil {
		return m.Languages
	}
	return nil
}

type InitRecommendationResponse struct {
	IsDone               bool     `protobuf:"varint,1,opt,name=isDone,proto3" json:"isDone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *InitRecommendationResponse) Reset()         { *m = InitRecommendationResponse{} }
func (m *InitRecommendationResponse) String() string { return proto.CompactTextString(m) }
func (*InitRecommendationResponse) ProtoMessage()    {}
func (*InitRecommendationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcf9f213a14e9af, []int{2}
}

func (m *InitRecommendationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRecommendationResponse.Unmarshal(m, b)
}
func (m *InitRecommendationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRecommendationResponse.Marshal(b, m, deterministic)
}
func (m *InitRecommendationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRecommendationResponse.Merge(m, src)
}
func (m *InitRecommendationResponse) XXX_Size() int {
	return xxx_messageInfo_InitRecommendationResponse.Size(m)
}
func (m *InitRecommendationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRecommendationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitRecommendationResponse proto.InternalMessageInfo

func (m *InitRecommendationResponse) GetIsDone() bool {
	if m != nil {
		return m.IsDone
	}
	return false
}

type GetRecommendationRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *GetRecommendationRequest) Reset()         { *m = GetRecommendationRequest{} }
func (m *GetRecommendationRequest) String() string { return proto.CompactTextString(m) }
func (*GetRecommendationRequest) ProtoMessage()    {}
func (*GetRecommendationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcf9f213a14e9af, []int{3}
}

func (m *GetRecommendationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecommendationRequest.Unmarshal(m, b)
}
func (m *GetRecommendationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecommendationRequest.Marshal(b, m, deterministic)
}
func (m *GetRecommendationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecommendationRequest.Merge(m, src)
}
func (m *GetRecommendationRequest) XXX_Size() int {
	return xxx_messageInfo_GetRecommendationRequest.Size(m)
}
func (m *GetRecommendationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecommendationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecommendationRequest proto.InternalMessageInfo

func (m *GetRecommendationRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type TileClickedResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *TileClickedResponse) Reset()         { *m = TileClickedResponse{} }
func (m *TileClickedResponse) String() string { return proto.CompactTextString(m) }
func (*TileClickedResponse) ProtoMessage()    {}
func (*TileClickedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcf9f213a14e9af, []int{4}
}

func (m *TileClickedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TileClickedResponse.Unmarshal(m, b)
}
func (m *TileClickedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TileClickedResponse.Marshal(b, m, deterministic)
}
func (m *TileClickedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TileClickedResponse.Merge(m, src)
}
func (m *TileClickedResponse) XXX_Size() int {
	return xxx_messageInfo_TileClickedResponse.Size(m)
}
func (m *TileClickedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TileClickedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TileClickedResponse proto.InternalMessageInfo

func (m *TileClickedResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type MovieTile struct {
	RefId                string    `protobuf:"bytes,1,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	Posters              *POSTERS  `protobuf:"bytes,2,opt,name=posters,proto3" json:"posters,omitempty"`
	Content              *CONTENT  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Metadata             *METADATA `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" bson:"-"`
}

func (m *MovieTile) Reset()         { *m = MovieTile{} }
func (m *MovieTile) String() string { return proto.CompactTextString(m) }
func (*MovieTile) ProtoMessage()    {}
func (*MovieTile) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcf9f213a14e9af, []int{5}
}

func (m *MovieTile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieTile.Unmarshal(m, b)
}
func (m *MovieTile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieTile.Marshal(b, m, deterministic)
}
func (m *MovieTile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieTile.Merge(m, src)
}
func (m *MovieTile) XXX_Size() int {
	return xxx_messageInfo_MovieTile.Size(m)
}
func (m *MovieTile) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieTile.DiscardUnknown(m)
}

var xxx_messageInfo_MovieTile proto.InternalMessageInfo

func (m *MovieTile) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *MovieTile) GetPosters() *POSTERS {
	if m != nil {
		return m.Posters
	}
	return nil
}

func (m *MovieTile) GetContent() *CONTENT {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *MovieTile) GetMetadata() *METADATA {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type METADATA struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *METADATA) Reset()         { *m = METADATA{} }
func (m *METADATA) String() string { return proto.CompactTextString(m) }
func (*METADATA) ProtoMessage()    {}
func (*METADATA) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcf9f213a14e9af, []int{6}
}

func (m *METADATA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_METADATA.Unmarshal(m, b)
}
func (m *METADATA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_METADATA.Marshal(b, m, deterministic)
}
func (m *METADATA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_METADATA.Merge(m, src)
}
func (m *METADATA) XXX_Size() int {
	return xxx_messageInfo_METADATA.Size(m)
}
func (m *METADATA) XXX_DiscardUnknown() {
	xxx_messageInfo_METADATA.DiscardUnknown(m)
}

var xxx_messageInfo_METADATA proto.InternalMessageInfo

func (m *METADATA) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type POSTERS struct {
	Landscape            []string `protobuf:"bytes,1,rep,name=landscape,proto3" json:"landscape,omitempty"`
	Portrait             []string `protobuf:"bytes,2,rep,name=portrait,proto3" json:"portrait,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *POSTERS) Reset()         { *m = POSTERS{} }
func (m *POSTERS) String() string { return proto.CompactTextString(m) }
func (*POSTERS) ProtoMessage()    {}
func (*POSTERS) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcf9f213a14e9af, []int{7}
}

func (m *POSTERS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_POSTERS.Unmarshal(m, b)
}
func (m *POSTERS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_POSTERS.Marshal(b, m, deterministic)
}
func (m *POSTERS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_POSTERS.Merge(m, src)
}
func (m *POSTERS) XXX_Size() int {
	return xxx_messageInfo_POSTERS.Size(m)
}
func (m *POSTERS) XXX_DiscardUnknown() {
	xxx_messageInfo_POSTERS.DiscardUnknown(m)
}

var xxx_messageInfo_POSTERS proto.InternalMessageInfo

func (m *POSTERS) GetLandscape() []string {
	if m != nil {
		return m.Landscape
	}
	return nil
}

func (m *POSTERS) GetPortrait() []string {
	if m != nil {
		return m.Portrait
	}
	return nil
}

type CONTENT struct {
	DetailPage           bool     `protobuf:"varint,2,opt,name=detailPage,proto3" json:"detailPage,omitempty"`
	Package              string   `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *CONTENT) Reset()         { *m = CONTENT{} }
func (m *CONTENT) String() string { return proto.CompactTextString(m) }
func (*CONTENT) ProtoMessage()    {}
func (*CONTENT) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcf9f213a14e9af, []int{8}
}

func (m *CONTENT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CONTENT.Unmarshal(m, b)
}
func (m *CONTENT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CONTENT.Marshal(b, m, deterministic)
}
func (m *CONTENT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CONTENT.Merge(m, src)
}
func (m *CONTENT) XXX_Size() int {
	return xxx_messageInfo_CONTENT.Size(m)
}
func (m *CONTENT) XXX_DiscardUnknown() {
	xxx_messageInfo_CONTENT.DiscardUnknown(m)
}

var xxx_messageInfo_CONTENT proto.InternalMessageInfo

func (m *CONTENT) GetDetailPage() bool {
	if m != nil {
		return m.DetailPage
	}
	return false
}

func (m *CONTENT) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func init() {
	proto.RegisterType((*TileClickedRequest)(nil), "RecommendationService.TileClickedRequest")
	proto.RegisterType((*InitRecommendationRequest)(nil), "RecommendationService.InitRecommendationRequest")
	proto.RegisterType((*InitRecommendationResponse)(nil), "RecommendationService.InitRecommendationResponse")
	proto.RegisterType((*GetRecommendationRequest)(nil), "RecommendationService.GetRecommendationRequest")
	proto.RegisterType((*TileClickedResponse)(nil), "RecommendationService.TileClickedResponse")
	proto.RegisterType((*MovieTile)(nil), "RecommendationService.MovieTile")
	proto.RegisterType((*METADATA)(nil), "RecommendationService.METADATA")
	proto.RegisterType((*POSTERS)(nil), "RecommendationService.POSTERS")
	proto.RegisterType((*CONTENT)(nil), "RecommendationService.CONTENT")
}

func init() { proto.RegisterFile("RecommendationService.proto", fileDescriptor_abcf9f213a14e9af) }

var fileDescriptor_abcf9f213a14e9af = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xbd, 0x52, 0x1b, 0x3d,
	0x14, 0x1d, 0xf1, 0x6b, 0x2e, 0x5f, 0xf3, 0x09, 0x48, 0x16, 0x43, 0xc0, 0xa3, 0x49, 0x41, 0x5c,
	0xb0, 0x84, 0xa4, 0xc8, 0x90, 0x8a, 0x31, 0x0e, 0xe3, 0x82, 0x9f, 0x59, 0xbb, 0xcf, 0xc8, 0xbb,
	0x97, 0x1d, 0x0d, 0x8b, 0xb4, 0x48, 0xb2, 0xbb, 0x34, 0x69, 0x53, 0xa6, 0x4e, 0x9f, 0xf7, 0xe1,
	0x15, 0xd2, 0xe7, 0x15, 0x32, 0xd2, 0xee, 0x62, 0x7b, 0xf0, 0x66, 0x92, 0x22, 0xdd, 0xde, 0x73,
	0x75, 0x74, 0x8e, 0x74, 0xae, 0x16, 0x76, 0x22, 0x8c, 0xd5, 0xdd, 0x1d, 0xca, 0x84, 0x5b, 0xa1,
	0x64, 0x1f, 0xf5, 0x58, 0xc4, 0x78, 0x98, 0x6b, 0x65, 0x15, 0xdd, 0x9a, 0xdb, 0x6c, 0xee, 0xa6,
	0x4a, 0xa5, 0x19, 0x86, 0x3c, 0x17, 0x21, 0x97, 0x52, 0x59, 0xdf, 0x35, 0x05, 0x89, 0x0d, 0x81,
	0x0e, 0x44, 0x86, 0x9d, 0x4c, 0xc4, 0xb7, 0x98, 0x44, 0x78, 0x3f, 0x42, 0x63, 0xe9, 0x33, 0x58,
	0x19, 0x19, 0xd4, 0xbd, 0x24, 0x20, 0x2d, 0x72, 0xb0, 0x16, 0x95, 0x95, 0xc3, 0xad, 0xc8, 0xb0,
	0x97, 0x04, 0x0b, 0x05, 0x5e, 0x54, 0x74, 0x17, 0xd6, 0xdc, 0x57, 0x3f, 0x56, 0x1a, 0x83, 0xc5,
	0x16, 0x39, 0x20, 0xd1, 0x04, 0x60, 0xf7, 0xb0, 0xdd, 0x93, 0xc2, 0xce, 0xda, 0x9b, 0x92, 0x4a,
	0x51, 0x6a, 0x34, 0x01, 0x69, 0x2d, 0xba, 0x2d, 0x8b, 0x8a, 0xee, 0x01, 0xc4, 0xdc, 0x62, 0xaa,
	0xb4, 0x40, 0x13, 0x2c, 0xf8, 0xde, 0x14, 0xe2, 0x24, 0x33, 0x2e, 0xd3, 0x11, 0x4f, 0xd1, 0x04,
	0x8b, 0xbe, 0x3d, 0x01, 0xd8, 0x5b, 0x68, 0xce, 0x93, 0x34, 0xb9, 0x92, 0x06, 0x9d, 0xa6, 0x30,
	0x67, 0x4a, 0xa2, 0x3f, 0x5e, 0x23, 0x2a, 0x2b, 0x76, 0x0c, 0xc1, 0x39, 0xd6, 0xfb, 0x9c, 0x77,
	0x25, 0x2c, 0x84, 0x8d, 0x99, 0x0b, 0x2c, 0x25, 0x02, 0x58, 0x35, 0xa3, 0x38, 0x46, 0x63, 0x4a,
	0x8d, 0xaa, 0x64, 0x0f, 0x04, 0xd6, 0x2e, 0xd4, 0x58, 0xa0, 0xa3, 0xd1, 0x2d, 0x58, 0xd1, 0x78,
	0xf3, 0x51, 0x54, 0xdb, 0x2e, 0x6b, 0xbc, 0xe9, 0x25, 0xf4, 0x1d, 0xac, 0xe6, 0xca, 0x58, 0xd4,
	0xc6, 0xdf, 0xf4, 0xfa, 0xf1, 0xde, 0xe1, 0xfc, 0xe8, 0xaf, 0xaf, 0xfa, 0x83, 0x6e, 0xd4, 0x8f,
	0xaa, 0xe5, 0x8e, 0x19, 0x2b, 0x69, 0x51, 0x5a, 0x1f, 0x44, 0x3d, 0xb3, 0x73, 0x75, 0x39, 0xe8,
	0x5e, 0x0e, 0xa2, 0x6a, 0x39, 0x7d, 0x0f, 0x8d, 0x3b, 0xb4, 0x3c, 0xe1, 0x96, 0x07, 0x4b, 0x9e,
	0xba, 0x5f, 0x43, 0xbd, 0xe8, 0x0e, 0x4e, 0xcf, 0x4e, 0x07, 0xa7, 0xd1, 0x23, 0x81, 0xb5, 0xa0,
	0x51, 0xa1, 0x74, 0x13, 0x96, 0xad, 0xb0, 0x19, 0x56, 0x47, 0xf2, 0x05, 0xeb, 0xc0, 0x6a, 0x69,
	0xb6, 0xcc, 0x2e, 0x31, 0x31, 0xcf, 0xb1, 0x8c, 0x7d, 0x02, 0xd0, 0x26, 0x34, 0x72, 0xa5, 0xad,
	0xe6, 0xc2, 0x96, 0xb9, 0x3f, 0xd6, 0x6e, 0x93, 0xd2, 0xb7, 0x1b, 0x90, 0x04, 0x2d, 0x17, 0xd9,
	0x35, 0x4f, 0xd1, 0xdf, 0x52, 0x23, 0x9a, 0x42, 0x5c, 0x02, 0x39, 0x8f, 0x6f, 0x5d, 0xb3, 0xf0,
	0x51, 0x95, 0xc7, 0x3f, 0x97, 0x60, 0xfe, 0x5b, 0xa1, 0x9f, 0x60, 0x7d, 0x2a, 0x4c, 0xfa, 0xaa,
	0xe6, 0xfc, 0x4f, 0x5f, 0x4c, 0xb3, 0xfd, 0x27, 0x4b, 0x8b, 0xd9, 0x60, 0xcf, 0x3f, 0x3f, 0xfc,
	0xf8, 0xba, 0xf0, 0x3f, 0xfb, 0x2f, 0xb4, 0x93, 0xee, 0x09, 0x69, 0xd3, 0x6f, 0x04, 0x76, 0xcf,
	0xd1, 0x76, 0x54, 0x96, 0xf1, 0xa1, 0xe6, 0x56, 0x8c, 0xf1, 0x83, 0xc8, 0x2c, 0x6a, 0x21, 0xd3,
	0x33, 0x6e, 0x39, 0x0d, 0x6b, 0x54, 0xea, 0xa6, 0xb6, 0xd9, 0xaa, 0x4b, 0xb0, 0x1a, 0x40, 0x76,
	0xe0, 0xcd, 0x30, 0xf6, 0x22, 0x4c, 0x7f, 0xa3, 0x7c, 0x42, 0xda, 0x47, 0x84, 0x7e, 0x21, 0xb0,
	0xe1, 0xfd, 0xf9, 0x81, 0x19, 0x72, 0x83, 0xc9, 0xbf, 0xb2, 0xb5, 0xef, 0x6d, 0x6d, 0xb3, 0xcd,
	0xc2, 0xd6, 0xac, 0x60, 0xe1, 0xe6, 0x3b, 0x81, 0x1d, 0xf7, 0xc8, 0x05, 0xcf, 0x66, 0x77, 0xeb,
	0xca, 0x54, 0x48, 0xa4, 0x47, 0x35, 0x22, 0xb5, 0xff, 0xa2, 0xe6, 0xeb, 0xbf, 0x60, 0x94, 0x59,
	0xbe, 0xf4, 0x3e, 0xf7, 0xd8, 0x76, 0x28, 0x9e, 0x2c, 0x2a, 0x7c, 0x9c, 0x90, 0xf6, 0x70, 0xc5,
	0xff, 0x6c, 0xdf, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x26, 0x8e, 0xaa, 0x0d, 0xc0, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecommendationServiceClient is the client API for RecommendationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecommendationServiceClient interface {
	TileClicked(ctx context.Context, in *TileClickedRequest, opts ...grpc.CallOption) (*TileClickedResponse, error)
	GetCollabrativeFilteringData(ctx context.Context, in *GetRecommendationRequest, opts ...grpc.CallOption) (RecommendationService_GetCollabrativeFilteringDataClient, error)
	GetContentbasedData(ctx context.Context, in *GetRecommendationRequest, opts ...grpc.CallOption) (RecommendationService_GetContentbasedDataClient, error)
	InitialRecommendationEngine(ctx context.Context, in *InitRecommendationRequest, opts ...grpc.CallOption) (*InitRecommendationResponse, error)
}

type recommendationServiceClient struct {
	cc *grpc.ClientConn
}

func NewRecommendationServiceClient(cc *grpc.ClientConn) RecommendationServiceClient {
	return &recommendationServiceClient{cc}
}

func (c *recommendationServiceClient) TileClicked(ctx context.Context, in *TileClickedRequest, opts ...grpc.CallOption) (*TileClickedResponse, error) {
	out := new(TileClickedResponse)
	err := c.cc.Invoke(ctx, "/RecommendationService.RecommendationService/TileClicked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceClient) GetCollabrativeFilteringData(ctx context.Context, in *GetRecommendationRequest, opts ...grpc.CallOption) (RecommendationService_GetCollabrativeFilteringDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecommendationService_serviceDesc.Streams[0], "/RecommendationService.RecommendationService/GetCollabrativeFilteringData", opts...)
	if err != nil {
		return nil, err
	}
	x := &recommendationServiceGetCollabrativeFilteringDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RecommendationService_GetCollabrativeFilteringDataClient interface {
	Recv() (*MovieTile, error)
	grpc.ClientStream
}

type recommendationServiceGetCollabrativeFilteringDataClient struct {
	grpc.ClientStream
}

func (x *recommendationServiceGetCollabrativeFilteringDataClient) Recv() (*MovieTile, error) {
	m := new(MovieTile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *recommendationServiceClient) GetContentbasedData(ctx context.Context, in *GetRecommendationRequest, opts ...grpc.CallOption) (RecommendationService_GetContentbasedDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecommendationService_serviceDesc.Streams[1], "/RecommendationService.RecommendationService/GetContentbasedData", opts...)
	if err != nil {
		return nil, err
	}
	x := &recommendationServiceGetContentbasedDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RecommendationService_GetContentbasedDataClient interface {
	Recv() (*MovieTile, error)
	grpc.ClientStream
}

type recommendationServiceGetContentbasedDataClient struct {
	grpc.ClientStream
}

func (x *recommendationServiceGetContentbasedDataClient) Recv() (*MovieTile, error) {
	m := new(MovieTile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *recommendationServiceClient) InitialRecommendationEngine(ctx context.Context, in *InitRecommendationRequest, opts ...grpc.CallOption) (*InitRecommendationResponse, error) {
	out := new(InitRecommendationResponse)
	err := c.cc.Invoke(ctx, "/RecommendationService.RecommendationService/InitialRecommendationEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendationServiceServer is the server API for RecommendationService service.
type RecommendationServiceServer interface {
	TileClicked(context.Context, *TileClickedRequest) (*TileClickedResponse, error)
	GetCollabrativeFilteringData(*GetRecommendationRequest, RecommendationService_GetCollabrativeFilteringDataServer) error
	GetContentbasedData(*GetRecommendationRequest, RecommendationService_GetContentbasedDataServer) error
	InitialRecommendationEngine(context.Context, *InitRecommendationRequest) (*InitRecommendationResponse, error)
}

// UnimplementedRecommendationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRecommendationServiceServer struct {
}

func (*UnimplementedRecommendationServiceServer) TileClicked(ctx context.Context, req *TileClickedRequest) (*TileClickedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TileClicked not implemented")
}
func (*UnimplementedRecommendationServiceServer) GetCollabrativeFilteringData(req *GetRecommendationRequest, srv RecommendationService_GetCollabrativeFilteringDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCollabrativeFilteringData not implemented")
}
func (*UnimplementedRecommendationServiceServer) GetContentbasedData(req *GetRecommendationRequest, srv RecommendationService_GetContentbasedDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetContentbasedData not implemented")
}
func (*UnimplementedRecommendationServiceServer) InitialRecommendationEngine(ctx context.Context, req *InitRecommendationRequest) (*InitRecommendationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitialRecommendationEngine not implemented")
}

func RegisterRecommendationServiceServer(s *grpc.Server, srv RecommendationServiceServer) {
	s.RegisterService(&_RecommendationService_serviceDesc, srv)
}

func _RecommendationService_TileClicked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TileClickedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceServer).TileClicked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecommendationService.RecommendationService/TileClicked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceServer).TileClicked(ctx, req.(*TileClickedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationService_GetCollabrativeFilteringData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRecommendationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecommendationServiceServer).GetCollabrativeFilteringData(m, &recommendationServiceGetCollabrativeFilteringDataServer{stream})
}

type RecommendationService_GetCollabrativeFilteringDataServer interface {
	Send(*MovieTile) error
	grpc.ServerStream
}

type recommendationServiceGetCollabrativeFilteringDataServer struct {
	grpc.ServerStream
}

func (x *recommendationServiceGetCollabrativeFilteringDataServer) Send(m *MovieTile) error {
	return x.ServerStream.SendMsg(m)
}

func _RecommendationService_GetContentbasedData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRecommendationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecommendationServiceServer).GetContentbasedData(m, &recommendationServiceGetContentbasedDataServer{stream})
}

type RecommendationService_GetContentbasedDataServer interface {
	Send(*MovieTile) error
	grpc.ServerStream
}

type recommendationServiceGetContentbasedDataServer struct {
	grpc.ServerStream
}

func (x *recommendationServiceGetContentbasedDataServer) Send(m *MovieTile) error {
	return x.ServerStream.SendMsg(m)
}

func _RecommendationService_InitialRecommendationEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRecommendationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceServer).InitialRecommendationEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RecommendationService.RecommendationService/InitialRecommendationEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceServer).InitialRecommendationEngine(ctx, req.(*InitRecommendationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecommendationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RecommendationService.RecommendationService",
	HandlerType: (*RecommendationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TileClicked",
			Handler:    _RecommendationService_TileClicked_Handler,
		},
		{
			MethodName: "InitialRecommendationEngine",
			Handler:    _RecommendationService_InitialRecommendationEngine_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCollabrativeFilteringData",
			Handler:       _RecommendationService_GetCollabrativeFilteringData_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetContentbasedData",
			Handler:       _RecommendationService_GetContentbasedData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "RecommendationService.proto",
}