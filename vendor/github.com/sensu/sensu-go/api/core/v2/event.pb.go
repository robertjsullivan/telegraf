// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/sensu/sensu-go/api/core/v2/event.proto

package v2

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// An Event is the encapsulating type sent across the Sensu websocket transport.
type Event struct {
	// Timestamp is the time in seconds since the Epoch.
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Entity describes the entity in which the event occurred.
	Entity *Entity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	// Check describes the result of a check; if event is associated to check
	// execution.
	Check *Check `protobuf:"bytes,3,opt,name=check,proto3" json:"check,omitempty"`
	// Metrics are zero or more Sensu metrics
	Metrics *Metrics `protobuf:"bytes,4,opt,name=metrics,proto3" json:"metrics,omitempty"`
	// Metadata contains name, namespace, labels and annotations
	ObjectMeta `protobuf:"bytes,5,opt,name=metadata,proto3,embedded=metadata" json:"metadata"`
	// ID is the unique identifier of the event.
	ID []byte `protobuf:"bytes,6,opt,name=ID,proto3" json:"id"`
	// Sequence is the event sequence number. The agent increments the sequence
	// number by one for every successive event. When the agent restarts or
	// reconnects to another backend, the sequence number is reset to 1.
	Sequence int64 `protobuf:"varint,7,opt,name=Sequence,proto3" json:"sequence"`
	// Pipelines are the pipelines that should be used to process an event.
	// APIVersion should default to "core/v2" and Type should default to
	// "Pipeline".
	Pipelines            []*ResourceReference `protobuf:"bytes,8,rep,name=pipelines,proto3" json:"pipelines"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a6c1d479d0c50cf, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Event)(nil), "sensu.core.v2.Event")
}

func init() {
	proto.RegisterFile("github.com/sensu/sensu-go/api/core/v2/event.proto", fileDescriptor_4a6c1d479d0c50cf)
}

var fileDescriptor_4a6c1d479d0c50cf = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x6e, 0xd3, 0x40,
	0x18, 0xc7, 0x73, 0x4e, 0x93, 0xa6, 0xd7, 0x76, 0x39, 0x95, 0xca, 0x54, 0xc8, 0xb6, 0x98, 0xbc,
	0x70, 0x6e, 0x1c, 0xc4, 0xc0, 0x80, 0x90, 0x69, 0x87, 0x0a, 0x45, 0x48, 0xc7, 0xc6, 0x82, 0x9c,
	0xeb, 0xd7, 0xf4, 0x80, 0xf8, 0x8c, 0x7d, 0xb6, 0xc4, 0x1b, 0xf0, 0x08, 0x8c, 0x1d, 0x3b, 0xf0,
	0x00, 0x3c, 0x42, 0xc6, 0x3e, 0x81, 0x05, 0x66, 0xcb, 0x13, 0x30, 0x22, 0x9f, 0xaf, 0x09, 0x64,
	0xf2, 0x62, 0x7d, 0xfa, 0xfe, 0xff, 0xdf, 0xe7, 0xd3, 0xff, 0x8f, 0xc7, 0x73, 0xa1, 0xae, 0x8b,
	0x19, 0xe5, 0x72, 0x11, 0xe4, 0x90, 0xe4, 0x45, 0xfb, 0x7d, 0x32, 0x97, 0x41, 0x9c, 0x8a, 0x80,
	0xcb, 0x0c, 0x82, 0x32, 0x0c, 0xa0, 0x84, 0x44, 0xd1, 0x34, 0x93, 0x4a, 0x92, 0x43, 0xed, 0xa0,
	0x8d, 0x44, 0xcb, 0xf0, 0xe4, 0xe9, 0x3f, 0x17, 0xe6, 0x72, 0x2e, 0x03, 0xed, 0x9a, 0x15, 0x57,
	0x2f, 0xcb, 0x31, 0x9d, 0xd0, 0xb1, 0x5e, 0xea, 0x9d, 0x9e, 0xda, 0x23, 0x27, 0x61, 0xc7, 0xff,
	0x26, 0x4a, 0xa8, 0x2f, 0x86, 0xe9, 0xf8, 0x56, 0x7e, 0x0d, 0xfc, 0xa3, 0x41, 0x26, 0xdd, 0x90,
	0x05, 0xa8, 0x4c, 0xf0, 0xdc, 0x40, 0xa7, 0x9d, 0xa1, 0xd8, 0x10, 0x2f, 0xba, 0x11, 0x19, 0xe4,
	0xb2, 0xc8, 0x38, 0xbc, 0xcf, 0xe0, 0x0a, 0x32, 0x48, 0x38, 0xb4, 0xfc, 0xe3, 0xef, 0x7d, 0x3c,
	0x38, 0x6f, 0x22, 0x26, 0x8f, 0xf0, 0x9e, 0x12, 0x0b, 0xc8, 0x55, 0xbc, 0x48, 0x6d, 0xe4, 0x21,
	0xbf, 0xcf, 0x36, 0x0b, 0x32, 0xc1, 0xc3, 0x36, 0x11, 0xdb, 0xf2, 0x90, 0xbf, 0x1f, 0x3e, 0xa0,
	0xff, 0x75, 0x41, 0xcf, 0xb5, 0x18, 0xed, 0x2c, 0x2b, 0x17, 0x31, 0x63, 0x25, 0xa7, 0x78, 0xa0,
	0x23, 0xb1, 0xfb, 0x9a, 0x39, 0xda, 0x62, 0x5e, 0x35, 0x9a, 0x41, 0x5a, 0x23, 0x79, 0x86, 0x77,
	0x4d, 0x22, 0xf6, 0x8e, 0x66, 0x8e, 0xb7, 0x98, 0x69, 0xab, 0x1a, 0xea, 0xde, 0x4c, 0x5e, 0xe3,
	0x51, 0x13, 0xca, 0x65, 0xac, 0x62, 0x7b, 0xa0, 0xc1, 0x87, 0x5b, 0xe0, 0x9b, 0xd9, 0x07, 0xe0,
	0x6a, 0x0a, 0x2a, 0x8e, 0x8e, 0x96, 0x95, 0xdb, 0xbb, 0xab, 0x5c, 0xb4, 0xaa, 0xdc, 0x35, 0xc6,
	0xd6, 0x13, 0x39, 0xc6, 0xd6, 0xc5, 0x99, 0x3d, 0xf4, 0x90, 0x7f, 0x10, 0x0d, 0x57, 0x95, 0x6b,
	0x89, 0x4b, 0x66, 0x5d, 0x9c, 0x11, 0x1f, 0x8f, 0xde, 0xc2, 0xe7, 0xa2, 0x49, 0xcf, 0xde, 0x6d,
	0x02, 0x8a, 0x0e, 0x9a, 0x0b, 0xb9, 0xd9, 0xb1, 0xb5, 0x4a, 0xa6, 0x78, 0x2f, 0x15, 0x29, 0x7c,
	0x12, 0x09, 0xe4, 0xf6, 0xc8, 0xeb, 0xfb, 0xfb, 0xa1, 0xb7, 0xf5, 0x1e, 0x66, 0x1a, 0x61, 0xf7,
	0x85, 0x44, 0x87, 0xab, 0xca, 0xdd, 0x60, 0x6c, 0x33, 0x3e, 0x1f, 0x7d, 0xbd, 0x71, 0x7b, 0xb7,
	0x37, 0x2e, 0x8a, 0xbc, 0x3f, 0xbf, 0x1c, 0x74, 0x5b, 0x3b, 0xe8, 0x47, 0xed, 0xa0, 0x65, 0xed,
	0xa0, 0xbb, 0xda, 0x41, 0x3f, 0x6b, 0x07, 0x7d, 0xfb, 0xed, 0xf4, 0xde, 0x59, 0x65, 0x38, 0x1b,
	0xea, 0x5e, 0x27, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x67, 0xa1, 0xe8, 0x5f, 0x03, 0x00,
	0x00,
}

func (this *Event) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Event)
	if !ok {
		that2, ok := that.(Event)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	if !this.Entity.Equal(that1.Entity) {
		return false
	}
	if !this.Check.Equal(that1.Check) {
		return false
	}
	if !this.Metrics.Equal(that1.Metrics) {
		return false
	}
	if !this.ObjectMeta.Equal(&that1.ObjectMeta) {
		return false
	}
	if !bytes.Equal(this.ID, that1.ID) {
		return false
	}
	if this.Sequence != that1.Sequence {
		return false
	}
	if len(this.Pipelines) != len(that1.Pipelines) {
		return false
	}
	for i := range this.Pipelines {
		if !this.Pipelines[i].Equal(that1.Pipelines[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type EventFace interface {
	Proto() github_com_golang_protobuf_proto.Message
	GetTimestamp() int64
	GetEntity() *Entity
	GetCheck() *Check
	GetMetrics() *Metrics
	GetObjectMeta() ObjectMeta
	GetID() []byte
	GetSequence() int64
	GetPipelines() []*ResourceReference
}

func (this *Event) Proto() github_com_golang_protobuf_proto.Message {
	return this
}

func (this *Event) TestProto() github_com_golang_protobuf_proto.Message {
	return NewEventFromFace(this)
}

func (this *Event) GetTimestamp() int64 {
	return this.Timestamp
}

func (this *Event) GetEntity() *Entity {
	return this.Entity
}

func (this *Event) GetCheck() *Check {
	return this.Check
}

func (this *Event) GetMetrics() *Metrics {
	return this.Metrics
}

func (this *Event) GetObjectMeta() ObjectMeta {
	return this.ObjectMeta
}

func (this *Event) GetID() []byte {
	return this.ID
}

func (this *Event) GetSequence() int64 {
	return this.Sequence
}

func (this *Event) GetPipelines() []*ResourceReference {
	return this.Pipelines
}

func NewEventFromFace(that EventFace) *Event {
	this := &Event{}
	this.Timestamp = that.GetTimestamp()
	this.Entity = that.GetEntity()
	this.Check = that.GetCheck()
	this.Metrics = that.GetMetrics()
	this.ObjectMeta = that.GetObjectMeta()
	this.ID = that.GetID()
	this.Sequence = that.GetSequence()
	this.Pipelines = that.GetPipelines()
	return this
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Pipelines) > 0 {
		for iNdEx := len(m.Pipelines) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pipelines[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.Sequence != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x32
	}
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Metrics != nil {
		{
			size, err := m.Metrics.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Check != nil {
		{
			size, err := m.Check.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Entity != nil {
		{
			size, err := m.Entity.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedEvent(r randyEvent, easy bool) *Event {
	this := &Event{}
	this.Timestamp = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Timestamp *= -1
	}
	if r.Intn(5) != 0 {
		this.Entity = NewPopulatedEntity(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Check = NewPopulatedCheck(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Metrics = NewPopulatedMetrics(r, easy)
	}
	v1 := NewPopulatedObjectMeta(r, easy)
	this.ObjectMeta = *v1
	v2 := r.Intn(100)
	this.ID = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.ID[i] = byte(r.Intn(256))
	}
	this.Sequence = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Sequence *= -1
	}
	if r.Intn(5) != 0 {
		v3 := r.Intn(5)
		this.Pipelines = make([]*ResourceReference, v3)
		for i := 0; i < v3; i++ {
			this.Pipelines[i] = NewPopulatedResourceReference(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEvent(r, 9)
	}
	return this
}

type randyEvent interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneEvent(r randyEvent) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringEvent(r randyEvent) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneEvent(r)
	}
	return string(tmps)
}
func randUnrecognizedEvent(r randyEvent, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldEvent(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldEvent(dAtA []byte, r randyEvent, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateEvent(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateEvent(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateEvent(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateEvent(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateEvent(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateEvent(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateEvent(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovEvent(uint64(m.Timestamp))
	}
	if m.Entity != nil {
		l = m.Entity.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.Check != nil {
		l = m.Check.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.Metrics != nil {
		l = m.Metrics.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	l = m.ObjectMeta.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovEvent(uint64(m.Sequence))
	}
	if len(m.Pipelines) > 0 {
		for _, e := range m.Pipelines {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entity == nil {
				m.Entity = &Entity{}
			}
			if err := m.Entity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Check", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Check == nil {
				m.Check = &Check{}
			}
			if err := m.Check.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metrics == nil {
				m.Metrics = &Metrics{}
			}
			if err := m.Metrics.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = append(m.ID[:0], dAtA[iNdEx:postIndex]...)
			if m.ID == nil {
				m.ID = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pipelines", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pipelines = append(m.Pipelines, &ResourceReference{})
			if err := m.Pipelines[len(m.Pipelines)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)