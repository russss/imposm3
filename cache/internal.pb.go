// Code generated by protoc-gen-go.
// source: cache/internal.proto
// DO NOT EDIT!

package cache

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type DeltaCoords struct {
	Ids              []int64 `protobuf:"zigzag64,1,rep,packed,name=ids" json:"ids,omitempty"`
	Lats             []int64 `protobuf:"zigzag64,2,rep,packed,name=lats" json:"lats,omitempty"`
	Lons             []int64 `protobuf:"zigzag64,3,rep,packed,name=lons" json:"lons,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeltaCoords) Reset()         { *m = DeltaCoords{} }
func (m *DeltaCoords) String() string { return proto.CompactTextString(m) }
func (*DeltaCoords) ProtoMessage()    {}

func (m *DeltaCoords) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *DeltaCoords) GetLats() []int64 {
	if m != nil {
		return m.Lats
	}
	return nil
}

func (m *DeltaCoords) GetLons() []int64 {
	if m != nil {
		return m.Lons
	}
	return nil
}

func init() {
}
