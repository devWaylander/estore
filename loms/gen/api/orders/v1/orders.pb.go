// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: orders.proto

package orders

import (
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

type Statuses int32

const (
	Statuses_STATUS_UNSPECIFIED      Statuses = 0
	Statuses_STATUS_NEW              Statuses = 1
	Statuses_STATUS_AWAITING_PAYMENT Statuses = 2
	Statuses_STATUS_FAILED           Statuses = 3
	Statuses_STATUS_PAYED            Statuses = 4
	Statuses_STATUS_CANCELED         Statuses = 5
)

// Enum value maps for Statuses.
var (
	Statuses_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_NEW",
		2: "STATUS_AWAITING_PAYMENT",
		3: "STATUS_FAILED",
		4: "STATUS_PAYED",
		5: "STATUS_CANCELED",
	}
	Statuses_value = map[string]int32{
		"STATUS_UNSPECIFIED":      0,
		"STATUS_NEW":              1,
		"STATUS_AWAITING_PAYMENT": 2,
		"STATUS_FAILED":           3,
		"STATUS_PAYED":            4,
		"STATUS_CANCELED":         5,
	}
)

func (x Statuses) Enum() *Statuses {
	p := new(Statuses)
	*p = x
	return p
}

func (x Statuses) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Statuses) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_proto_enumTypes[0].Descriptor()
}

func (Statuses) Type() protoreflect.EnumType {
	return &file_orders_proto_enumTypes[0]
}

func (x Statuses) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Statuses.Descriptor instead.
func (Statuses) EnumDescriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{0}
}

// Common
type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *Item) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

// OrderCreate
type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items  []*Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[1]
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
	return file_orders_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type OrderCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *OrderCreateReq) Reset() {
	*x = OrderCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateReq) ProtoMessage() {}

func (x *OrderCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateReq.ProtoReflect.Descriptor instead.
func (*OrderCreateReq) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{2}
}

func (x *OrderCreateReq) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OrderCreateResp) Reset() {
	*x = OrderCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateResp) ProtoMessage() {}

func (x *OrderCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateResp.ProtoReflect.Descriptor instead.
func (*OrderCreateResp) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{3}
}

func (x *OrderCreateResp) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

// OrderGetInfo
type OrderGetInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OrderGetInfoReq) Reset() {
	*x = OrderGetInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderGetInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderGetInfoReq) ProtoMessage() {}

func (x *OrderGetInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderGetInfoReq.ProtoReflect.Descriptor instead.
func (*OrderGetInfoReq) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{4}
}

func (x *OrderGetInfoReq) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type OrderGetInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Statuses `protobuf:"varint,1,opt,name=status,proto3,enum=route256.ozon.ru.project.loms.api.orders.v1.Statuses" json:"status,omitempty"`
	UserId int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items  []*Item  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *OrderGetInfoResp) Reset() {
	*x = OrderGetInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderGetInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderGetInfoResp) ProtoMessage() {}

func (x *OrderGetInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderGetInfoResp.ProtoReflect.Descriptor instead.
func (*OrderGetInfoResp) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{5}
}

func (x *OrderGetInfoResp) GetStatus() Statuses {
	if x != nil {
		return x.Status
	}
	return Statuses_STATUS_UNSPECIFIED
}

func (x *OrderGetInfoResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderGetInfoResp) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

// OrderGetStockInfo
type OrderGetStockInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *OrderGetStockInfoReq) Reset() {
	*x = OrderGetStockInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderGetStockInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderGetStockInfoReq) ProtoMessage() {}

func (x *OrderGetStockInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderGetStockInfoReq.ProtoReflect.Descriptor instead.
func (*OrderGetStockInfoReq) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{6}
}

func (x *OrderGetStockInfoReq) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

type OrderGetStockInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *OrderGetStockInfoResp) Reset() {
	*x = OrderGetStockInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderGetStockInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderGetStockInfoResp) ProtoMessage() {}

func (x *OrderGetStockInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderGetStockInfoResp.ProtoReflect.Descriptor instead.
func (*OrderGetStockInfoResp) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{7}
}

func (x *OrderGetStockInfoResp) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// OrderCancel
type OrderCancelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OrderCancelReq) Reset() {
	*x = OrderCancelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCancelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCancelReq) ProtoMessage() {}

func (x *OrderCancelReq) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCancelReq.ProtoReflect.Descriptor instead.
func (*OrderCancelReq) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{8}
}

func (x *OrderCancelReq) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type OrderCancelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderCancelResp) Reset() {
	*x = OrderCancelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCancelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCancelResp) ProtoMessage() {}

func (x *OrderCancelResp) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCancelResp.ProtoReflect.Descriptor instead.
func (*OrderCancelResp) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{9}
}

// OrderPay
type OrderPayReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OrderPayReq) Reset() {
	*x = OrderPayReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayReq) ProtoMessage() {}

func (x *OrderPayReq) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayReq.ProtoReflect.Descriptor instead.
func (*OrderPayReq) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{10}
}

func (x *OrderPayReq) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type OrderPayResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderPayResp) Reset() {
	*x = OrderPayResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayResp) ProtoMessage() {}

func (x *OrderPayResp) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayResp.ProtoReflect.Descriptor instead.
func (*OrderPayResp) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{11}
}

var File_orders_proto protoreflect.FileDescriptor

var file_orders_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x72, 0x75,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b, 0x0a, 0x03,
	0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04,
	0x20, 0x00, 0x40, 0x00, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x7e, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x22, 0x04,
	0x20, 0x00, 0x40, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x51, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x64, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x52, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e,
	0x2e, 0x72, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x32, 0x04, 0x20,
	0x00, 0x40, 0x00, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0xc3, 0x01, 0x0a,
	0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x4d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x35, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x72, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x33, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x03, 0x73, 0x6b,
	0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x20, 0x00,
	0x40, 0x00, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0x2d, 0x0a, 0x15, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x36, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x32,
	0x04, 0x20, 0x00, 0x40, 0x00, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x11,
	0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x33, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x12, 0x24, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x32, 0x04, 0x20, 0x00, 0x40, 0x00, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x2a, 0x89, 0x01, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x50,
	0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x59, 0x45, 0x44, 0x10, 0x04, 0x12, 0x13, 0x0a,
	0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44,
	0x10, 0x05, 0x32, 0xd3, 0x05, 0x0a, 0x04, 0x4c, 0x4f, 0x4d, 0x53, 0x12, 0x8a, 0x01, 0x0a, 0x0b,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x3c, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x8d, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x3d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32,
	0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x9c, 0x01, 0x0a, 0x11, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x41,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x72,
	0x75, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x1a, 0x42, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x72, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x8a, 0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x3b, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32,
	0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x1a, 0x3c, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e,
	0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x81, 0x01, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x12, 0x38, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x72, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x39, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x32, 0x35, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6c, 0x6f, 0x6d, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_proto_rawDescOnce sync.Once
	file_orders_proto_rawDescData = file_orders_proto_rawDesc
)

func file_orders_proto_rawDescGZIP() []byte {
	file_orders_proto_rawDescOnce.Do(func() {
		file_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_proto_rawDescData)
	})
	return file_orders_proto_rawDescData
}

var file_orders_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_orders_proto_goTypes = []interface{}{
	(Statuses)(0),                 // 0: route256.ozon.ru.project.loms.api.orders.v1.Statuses
	(*Item)(nil),                  // 1: route256.ozon.ru.project.loms.api.orders.v1.Item
	(*Order)(nil),                 // 2: route256.ozon.ru.project.loms.api.orders.v1.Order
	(*OrderCreateReq)(nil),        // 3: route256.ozon.ru.project.loms.api.orders.v1.OrderCreateReq
	(*OrderCreateResp)(nil),       // 4: route256.ozon.ru.project.loms.api.orders.v1.OrderCreateResp
	(*OrderGetInfoReq)(nil),       // 5: route256.ozon.ru.project.loms.api.orders.v1.OrderGetInfoReq
	(*OrderGetInfoResp)(nil),      // 6: route256.ozon.ru.project.loms.api.orders.v1.OrderGetInfoResp
	(*OrderGetStockInfoReq)(nil),  // 7: route256.ozon.ru.project.loms.api.orders.v1.OrderGetStockInfoReq
	(*OrderGetStockInfoResp)(nil), // 8: route256.ozon.ru.project.loms.api.orders.v1.OrderGetStockInfoResp
	(*OrderCancelReq)(nil),        // 9: route256.ozon.ru.project.loms.api.orders.v1.OrderCancelReq
	(*OrderCancelResp)(nil),       // 10: route256.ozon.ru.project.loms.api.orders.v1.OrderCancelResp
	(*OrderPayReq)(nil),           // 11: route256.ozon.ru.project.loms.api.orders.v1.OrderPayReq
	(*OrderPayResp)(nil),          // 12: route256.ozon.ru.project.loms.api.orders.v1.OrderPayResp
}
var file_orders_proto_depIdxs = []int32{
	1,  // 0: route256.ozon.ru.project.loms.api.orders.v1.Order.items:type_name -> route256.ozon.ru.project.loms.api.orders.v1.Item
	2,  // 1: route256.ozon.ru.project.loms.api.orders.v1.OrderCreateReq.order:type_name -> route256.ozon.ru.project.loms.api.orders.v1.Order
	0,  // 2: route256.ozon.ru.project.loms.api.orders.v1.OrderGetInfoResp.status:type_name -> route256.ozon.ru.project.loms.api.orders.v1.Statuses
	1,  // 3: route256.ozon.ru.project.loms.api.orders.v1.OrderGetInfoResp.items:type_name -> route256.ozon.ru.project.loms.api.orders.v1.Item
	3,  // 4: route256.ozon.ru.project.loms.api.orders.v1.LOMS.OrderCreate:input_type -> route256.ozon.ru.project.loms.api.orders.v1.OrderCreateReq
	5,  // 5: route256.ozon.ru.project.loms.api.orders.v1.LOMS.OrderGetInfo:input_type -> route256.ozon.ru.project.loms.api.orders.v1.OrderGetInfoReq
	7,  // 6: route256.ozon.ru.project.loms.api.orders.v1.LOMS.OrderGetStockInfo:input_type -> route256.ozon.ru.project.loms.api.orders.v1.OrderGetStockInfoReq
	9,  // 7: route256.ozon.ru.project.loms.api.orders.v1.LOMS.OrderCancel:input_type -> route256.ozon.ru.project.loms.api.orders.v1.OrderCancelReq
	11, // 8: route256.ozon.ru.project.loms.api.orders.v1.LOMS.OrderPay:input_type -> route256.ozon.ru.project.loms.api.orders.v1.OrderPayReq
	4,  // 9: route256.ozon.ru.project.loms.api.orders.v1.LOMS.OrderCreate:output_type -> route256.ozon.ru.project.loms.api.orders.v1.OrderCreateResp
	6,  // 10: route256.ozon.ru.project.loms.api.orders.v1.LOMS.OrderGetInfo:output_type -> route256.ozon.ru.project.loms.api.orders.v1.OrderGetInfoResp
	8,  // 11: route256.ozon.ru.project.loms.api.orders.v1.LOMS.OrderGetStockInfo:output_type -> route256.ozon.ru.project.loms.api.orders.v1.OrderGetStockInfoResp
	10, // 12: route256.ozon.ru.project.loms.api.orders.v1.LOMS.OrderCancel:output_type -> route256.ozon.ru.project.loms.api.orders.v1.OrderCancelResp
	12, // 13: route256.ozon.ru.project.loms.api.orders.v1.LOMS.OrderPay:output_type -> route256.ozon.ru.project.loms.api.orders.v1.OrderPayResp
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_orders_proto_init() }
func file_orders_proto_init() {
	if File_orders_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_orders_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_orders_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateReq); i {
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
		file_orders_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateResp); i {
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
		file_orders_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderGetInfoReq); i {
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
		file_orders_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderGetInfoResp); i {
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
		file_orders_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderGetStockInfoReq); i {
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
		file_orders_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderGetStockInfoResp); i {
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
		file_orders_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCancelReq); i {
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
		file_orders_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCancelResp); i {
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
		file_orders_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayReq); i {
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
		file_orders_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayResp); i {
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
			RawDescriptor: file_orders_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_proto_goTypes,
		DependencyIndexes: file_orders_proto_depIdxs,
		EnumInfos:         file_orders_proto_enumTypes,
		MessageInfos:      file_orders_proto_msgTypes,
	}.Build()
	File_orders_proto = out.File
	file_orders_proto_rawDesc = nil
	file_orders_proto_goTypes = nil
	file_orders_proto_depIdxs = nil
}
