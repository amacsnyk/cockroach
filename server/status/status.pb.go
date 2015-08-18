// Code generated by protoc-gen-gogo.
// source: cockroach/server/status/status.proto
// DO NOT EDIT!

/*
	Package status is a generated protocol buffer package.

	It is generated from these files:
		cockroach/server/status/status.proto

	It has these top-level messages:
		NodeStatus
*/
package status

import proto "github.com/gogo/protobuf/proto"
import math "math"
import cockroach_proto "github.com/cockroachdb/cockroach/proto"
import cockroach_storage_engine "github.com/cockroachdb/cockroach/storage/engine"

// discarding unused import gogoproto "gogoproto"

import github_com_cockroachdb_cockroach_proto "github.com/cockroachdb/cockroach/proto"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// NodeStatus contains the stats needed to calculate the current status of a
// node.
type NodeStatus struct {
	Desc                 cockroach_proto.NodeDescriptor                   `protobuf:"bytes,1,opt,name=desc" json:"desc"`
	StoreIDs             []github_com_cockroachdb_cockroach_proto.StoreID `protobuf:"varint,2,rep,name=store_ids,casttype=github.com/cockroachdb/cockroach/proto.StoreID" json:"store_ids,omitempty"`
	RangeCount           int32                                            `protobuf:"varint,3,opt,name=range_count" json:"range_count"`
	StartedAt            int64                                            `protobuf:"varint,4,opt,name=started_at" json:"started_at"`
	UpdatedAt            int64                                            `protobuf:"varint,5,opt,name=updated_at" json:"updated_at"`
	Stats                cockroach_storage_engine.MVCCStats               `protobuf:"bytes,6,opt,name=stats" json:"stats"`
	LeaderRangeCount     int32                                            `protobuf:"varint,7,opt,name=leader_range_count" json:"leader_range_count"`
	ReplicatedRangeCount int32                                            `protobuf:"varint,8,opt,name=replicated_range_count" json:"replicated_range_count"`
	AvailableRangeCount  int32                                            `protobuf:"varint,9,opt,name=available_range_count" json:"available_range_count"`
	XXX_unrecognized     []byte                                           `json:"-"`
}

func (m *NodeStatus) Reset()         { *m = NodeStatus{} }
func (m *NodeStatus) String() string { return proto.CompactTextString(m) }
func (*NodeStatus) ProtoMessage()    {}

func (m *NodeStatus) GetDesc() cockroach_proto.NodeDescriptor {
	if m != nil {
		return m.Desc
	}
	return cockroach_proto.NodeDescriptor{}
}

func (m *NodeStatus) GetStoreIDs() []github_com_cockroachdb_cockroach_proto.StoreID {
	if m != nil {
		return m.StoreIDs
	}
	return nil
}

func (m *NodeStatus) GetRangeCount() int32 {
	if m != nil {
		return m.RangeCount
	}
	return 0
}

func (m *NodeStatus) GetStartedAt() int64 {
	if m != nil {
		return m.StartedAt
	}
	return 0
}

func (m *NodeStatus) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *NodeStatus) GetStats() cockroach_storage_engine.MVCCStats {
	if m != nil {
		return m.Stats
	}
	return cockroach_storage_engine.MVCCStats{}
}

func (m *NodeStatus) GetLeaderRangeCount() int32 {
	if m != nil {
		return m.LeaderRangeCount
	}
	return 0
}

func (m *NodeStatus) GetReplicatedRangeCount() int32 {
	if m != nil {
		return m.ReplicatedRangeCount
	}
	return 0
}

func (m *NodeStatus) GetAvailableRangeCount() int32 {
	if m != nil {
		return m.AvailableRangeCount
	}
	return 0
}

func (m *NodeStatus) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Desc.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreIDs", wireType)
			}
			var v github_com_cockroachdb_cockroach_proto.StoreID
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (github_com_cockroachdb_cockroach_proto.StoreID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StoreIDs = append(m.StoreIDs, v)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeCount", wireType)
			}
			m.RangeCount = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RangeCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartedAt", wireType)
			}
			m.StartedAt = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.StartedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.UpdatedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stats.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaderRangeCount", wireType)
			}
			m.LeaderRangeCount = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.LeaderRangeCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicatedRangeCount", wireType)
			}
			m.ReplicatedRangeCount = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ReplicatedRangeCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvailableRangeCount", wireType)
			}
			m.AvailableRangeCount = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.AvailableRangeCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipStatus(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func skipStatus(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStatus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStatus(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStatus = fmt.Errorf("proto: negative length found during unmarshaling")
)

func (m *NodeStatus) Size() (n int) {
	var l int
	_ = l
	l = m.Desc.Size()
	n += 1 + l + sovStatus(uint64(l))
	if len(m.StoreIDs) > 0 {
		for _, e := range m.StoreIDs {
			n += 1 + sovStatus(uint64(e))
		}
	}
	n += 1 + sovStatus(uint64(m.RangeCount))
	n += 1 + sovStatus(uint64(m.StartedAt))
	n += 1 + sovStatus(uint64(m.UpdatedAt))
	l = m.Stats.Size()
	n += 1 + l + sovStatus(uint64(l))
	n += 1 + sovStatus(uint64(m.LeaderRangeCount))
	n += 1 + sovStatus(uint64(m.ReplicatedRangeCount))
	n += 1 + sovStatus(uint64(m.AvailableRangeCount))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStatus(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStatus(x uint64) (n int) {
	return sovStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NodeStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *NodeStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintStatus(data, i, uint64(m.Desc.Size()))
	n1, err := m.Desc.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.StoreIDs) > 0 {
		for _, num := range m.StoreIDs {
			data[i] = 0x10
			i++
			i = encodeVarintStatus(data, i, uint64(num))
		}
	}
	data[i] = 0x18
	i++
	i = encodeVarintStatus(data, i, uint64(m.RangeCount))
	data[i] = 0x20
	i++
	i = encodeVarintStatus(data, i, uint64(m.StartedAt))
	data[i] = 0x28
	i++
	i = encodeVarintStatus(data, i, uint64(m.UpdatedAt))
	data[i] = 0x32
	i++
	i = encodeVarintStatus(data, i, uint64(m.Stats.Size()))
	n2, err := m.Stats.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x38
	i++
	i = encodeVarintStatus(data, i, uint64(m.LeaderRangeCount))
	data[i] = 0x40
	i++
	i = encodeVarintStatus(data, i, uint64(m.ReplicatedRangeCount))
	data[i] = 0x48
	i++
	i = encodeVarintStatus(data, i, uint64(m.AvailableRangeCount))
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Status(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Status(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStatus(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
