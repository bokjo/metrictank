package interning

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Archive) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "MetricDefinition":
			err = z.MetricDefinition.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "MetricDefinition")
				return
			}
		case "SchemaId":
			z.SchemaId, err = dc.ReadUint16()
			if err != nil {
				err = msgp.WrapError(err, "SchemaId")
				return
			}
		case "AggId":
			z.AggId, err = dc.ReadUint16()
			if err != nil {
				err = msgp.WrapError(err, "AggId")
				return
			}
		case "IrId":
			z.IrId, err = dc.ReadUint16()
			if err != nil {
				err = msgp.WrapError(err, "IrId")
				return
			}
		case "LastSave":
			z.LastSave, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "LastSave")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Archive) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "MetricDefinition"
	err = en.Append(0x85, 0xb0, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = z.MetricDefinition.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "MetricDefinition")
		return
	}
	// write "SchemaId"
	err = en.Append(0xa8, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint16(z.SchemaId)
	if err != nil {
		err = msgp.WrapError(err, "SchemaId")
		return
	}
	// write "AggId"
	err = en.Append(0xa5, 0x41, 0x67, 0x67, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint16(z.AggId)
	if err != nil {
		err = msgp.WrapError(err, "AggId")
		return
	}
	// write "IrId"
	err = en.Append(0xa4, 0x49, 0x72, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint16(z.IrId)
	if err != nil {
		err = msgp.WrapError(err, "IrId")
		return
	}
	// write "LastSave"
	err = en.Append(0xa8, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x61, 0x76, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.LastSave)
	if err != nil {
		err = msgp.WrapError(err, "LastSave")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Archive) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "MetricDefinition"
	o = append(o, 0x85, 0xb0, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e)
	o, err = z.MetricDefinition.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "MetricDefinition")
		return
	}
	// string "SchemaId"
	o = append(o, 0xa8, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	o = msgp.AppendUint16(o, z.SchemaId)
	// string "AggId"
	o = append(o, 0xa5, 0x41, 0x67, 0x67, 0x49, 0x64)
	o = msgp.AppendUint16(o, z.AggId)
	// string "IrId"
	o = append(o, 0xa4, 0x49, 0x72, 0x49, 0x64)
	o = msgp.AppendUint16(o, z.IrId)
	// string "LastSave"
	o = append(o, 0xa8, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x61, 0x76, 0x65)
	o = msgp.AppendUint32(o, z.LastSave)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Archive) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "MetricDefinition":
			bts, err = z.MetricDefinition.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "MetricDefinition")
				return
			}
		case "SchemaId":
			z.SchemaId, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SchemaId")
				return
			}
		case "AggId":
			z.AggId, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AggId")
				return
			}
		case "IrId":
			z.IrId, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IrId")
				return
			}
		case "LastSave":
			z.LastSave, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastSave")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Archive) Msgsize() (s int) {
	s = 1 + 17 + z.MetricDefinition.Msgsize() + 9 + msgp.Uint16Size + 6 + msgp.Uint16Size + 5 + msgp.Uint16Size + 9 + msgp.Uint32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ArchiveInterned) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "MetricDefinitionInterned":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "MetricDefinitionInterned")
					return
				}
				z.MetricDefinitionInterned = nil
			} else {
				if z.MetricDefinitionInterned == nil {
					z.MetricDefinitionInterned = new(MetricDefinitionInterned)
				}
				err = z.MetricDefinitionInterned.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "MetricDefinitionInterned")
					return
				}
			}
		case "SchemaId":
			z.SchemaId, err = dc.ReadUint16()
			if err != nil {
				err = msgp.WrapError(err, "SchemaId")
				return
			}
		case "AggId":
			z.AggId, err = dc.ReadUint16()
			if err != nil {
				err = msgp.WrapError(err, "AggId")
				return
			}
		case "IrId":
			z.IrId, err = dc.ReadUint16()
			if err != nil {
				err = msgp.WrapError(err, "IrId")
				return
			}
		case "LastSave":
			z.LastSave, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "LastSave")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ArchiveInterned) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "MetricDefinitionInterned"
	err = en.Append(0x85, 0xb8, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x64)
	if err != nil {
		return
	}
	if z.MetricDefinitionInterned == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.MetricDefinitionInterned.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "MetricDefinitionInterned")
			return
		}
	}
	// write "SchemaId"
	err = en.Append(0xa8, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint16(z.SchemaId)
	if err != nil {
		err = msgp.WrapError(err, "SchemaId")
		return
	}
	// write "AggId"
	err = en.Append(0xa5, 0x41, 0x67, 0x67, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint16(z.AggId)
	if err != nil {
		err = msgp.WrapError(err, "AggId")
		return
	}
	// write "IrId"
	err = en.Append(0xa4, 0x49, 0x72, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint16(z.IrId)
	if err != nil {
		err = msgp.WrapError(err, "IrId")
		return
	}
	// write "LastSave"
	err = en.Append(0xa8, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x61, 0x76, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.LastSave)
	if err != nil {
		err = msgp.WrapError(err, "LastSave")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ArchiveInterned) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "MetricDefinitionInterned"
	o = append(o, 0x85, 0xb8, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x64)
	if z.MetricDefinitionInterned == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.MetricDefinitionInterned.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "MetricDefinitionInterned")
			return
		}
	}
	// string "SchemaId"
	o = append(o, 0xa8, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	o = msgp.AppendUint16(o, z.SchemaId)
	// string "AggId"
	o = append(o, 0xa5, 0x41, 0x67, 0x67, 0x49, 0x64)
	o = msgp.AppendUint16(o, z.AggId)
	// string "IrId"
	o = append(o, 0xa4, 0x49, 0x72, 0x49, 0x64)
	o = msgp.AppendUint16(o, z.IrId)
	// string "LastSave"
	o = append(o, 0xa8, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x61, 0x76, 0x65)
	o = msgp.AppendUint32(o, z.LastSave)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ArchiveInterned) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "MetricDefinitionInterned":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.MetricDefinitionInterned = nil
			} else {
				if z.MetricDefinitionInterned == nil {
					z.MetricDefinitionInterned = new(MetricDefinitionInterned)
				}
				bts, err = z.MetricDefinitionInterned.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "MetricDefinitionInterned")
					return
				}
			}
		case "SchemaId":
			z.SchemaId, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SchemaId")
				return
			}
		case "AggId":
			z.AggId, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AggId")
				return
			}
		case "IrId":
			z.IrId, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IrId")
				return
			}
		case "LastSave":
			z.LastSave, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastSave")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ArchiveInterned) Msgsize() (s int) {
	s = 1 + 25
	if z.MetricDefinitionInterned == nil {
		s += msgp.NilSize
	} else {
		s += z.MetricDefinitionInterned.Msgsize()
	}
	s += 9 + msgp.Uint16Size + 6 + msgp.Uint16Size + 5 + msgp.Uint16Size + 9 + msgp.Uint32Size
	return
}
