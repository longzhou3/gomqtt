package global

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

//
// DecodeMsg implements msgp.Decodable
func (z *ProtoBufMsg) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbai uint32
	zbai, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbai > 0 {
		zbai--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ci":
			z.Cid, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "q":
			var zcmr uint32
			zcmr, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Qos) >= int(zcmr) {
				z.Qos = (z.Qos)[:zcmr]
			} else {
				z.Qos = make([]int32, zcmr)
			}
			for zxvk := range z.Qos {
				z.Qos[zxvk], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "mi":
			var zajw uint32
			zajw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.MsgIDs) >= int(zajw) {
				z.MsgIDs = (z.MsgIDs)[:zajw]
			} else {
				z.MsgIDs = make([][]byte, zajw)
			}
			for zbzg := range z.MsgIDs {
				z.MsgIDs[zbzg], err = dc.ReadBytes(z.MsgIDs[zbzg])
				if err != nil {
					return
				}
			}
		case "m":
			z.Msg, err = dc.ReadBytes(z.Msg)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ProtoBufMsg) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "ci"
	err = en.Append(0x84, 0xa2, 0x63, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Cid)
	if err != nil {
		return
	}
	// write "q"
	err = en.Append(0xa1, 0x71)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Qos)))
	if err != nil {
		return
	}
	for zxvk := range z.Qos {
		err = en.WriteInt32(z.Qos[zxvk])
		if err != nil {
			return
		}
	}
	// write "mi"
	err = en.Append(0xa2, 0x6d, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.MsgIDs)))
	if err != nil {
		return
	}
	for zbzg := range z.MsgIDs {
		err = en.WriteBytes(z.MsgIDs[zbzg])
		if err != nil {
			return
		}
	}
	// write "m"
	err = en.Append(0xa1, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Msg)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ProtoBufMsg) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "ci"
	o = append(o, 0x84, 0xa2, 0x63, 0x69)
	o = msgp.AppendInt64(o, z.Cid)
	// string "q"
	o = append(o, 0xa1, 0x71)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Qos)))
	for zxvk := range z.Qos {
		o = msgp.AppendInt32(o, z.Qos[zxvk])
	}
	// string "mi"
	o = append(o, 0xa2, 0x6d, 0x69)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MsgIDs)))
	for zbzg := range z.MsgIDs {
		o = msgp.AppendBytes(o, z.MsgIDs[zbzg])
	}
	// string "m"
	o = append(o, 0xa1, 0x6d)
	o = msgp.AppendBytes(o, z.Msg)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ProtoBufMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zwht uint32
	zwht, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zwht > 0 {
		zwht--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ci":
			z.Cid, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "q":
			var zhct uint32
			zhct, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Qos) >= int(zhct) {
				z.Qos = (z.Qos)[:zhct]
			} else {
				z.Qos = make([]int32, zhct)
			}
			for zxvk := range z.Qos {
				z.Qos[zxvk], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "mi":
			var zcua uint32
			zcua, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.MsgIDs) >= int(zcua) {
				z.MsgIDs = (z.MsgIDs)[:zcua]
			} else {
				z.MsgIDs = make([][]byte, zcua)
			}
			for zbzg := range z.MsgIDs {
				z.MsgIDs[zbzg], bts, err = msgp.ReadBytesBytes(bts, z.MsgIDs[zbzg])
				if err != nil {
					return
				}
			}
		case "m":
			z.Msg, bts, err = msgp.ReadBytesBytes(bts, z.Msg)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ProtoBufMsg) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 2 + msgp.ArrayHeaderSize + (len(z.Qos) * (msgp.Int32Size)) + 3 + msgp.ArrayHeaderSize
	for zbzg := range z.MsgIDs {
		s += msgp.BytesPrefixSize + len(z.MsgIDs[zbzg])
	}
	s += 2 + msgp.BytesPrefixSize + len(z.Msg)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TextMsg) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxhx uint32
	zxhx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "q":
			z.Qos, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "mi":
			z.MsgID, err = dc.ReadBytes(z.MsgID)
			if err != nil {
				return
			}
		case "m":
			z.Msg, err = dc.ReadBytes(z.Msg)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TextMsg) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "q"
	err = en.Append(0x83, 0xa1, 0x71)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Qos)
	if err != nil {
		return
	}
	// write "mi"
	err = en.Append(0xa2, 0x6d, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.MsgID)
	if err != nil {
		return
	}
	// write "m"
	err = en.Append(0xa1, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Msg)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TextMsg) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "q"
	o = append(o, 0x83, 0xa1, 0x71)
	o = msgp.AppendInt32(o, z.Qos)
	// string "mi"
	o = append(o, 0xa2, 0x6d, 0x69)
	o = msgp.AppendBytes(o, z.MsgID)
	// string "m"
	o = append(o, 0xa1, 0x6d)
	o = msgp.AppendBytes(o, z.Msg)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TextMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zlqf uint32
	zlqf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zlqf > 0 {
		zlqf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "q":
			z.Qos, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "mi":
			z.MsgID, bts, err = msgp.ReadBytesBytes(bts, z.MsgID)
			if err != nil {
				return
			}
		case "m":
			z.Msg, bts, err = msgp.ReadBytesBytes(bts, z.Msg)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TextMsg) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int32Size + 3 + msgp.BytesPrefixSize + len(z.MsgID) + 2 + msgp.BytesPrefixSize + len(z.Msg)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TextMsgs) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zpks uint32
	zpks, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zpks > 0 {
		zpks--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ci":
			z.Cid, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "ms":
			var zjfb uint32
			zjfb, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Msgs) >= int(zjfb) {
				z.Msgs = (z.Msgs)[:zjfb]
			} else {
				z.Msgs = make([]*TextMsg, zjfb)
			}
			for zdaf := range z.Msgs {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Msgs[zdaf] = nil
				} else {
					if z.Msgs[zdaf] == nil {
						z.Msgs[zdaf] = new(TextMsg)
					}
					var zcxo uint32
					zcxo, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zcxo > 0 {
						zcxo--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "q":
							z.Msgs[zdaf].Qos, err = dc.ReadInt32()
							if err != nil {
								return
							}
						case "mi":
							z.Msgs[zdaf].MsgID, err = dc.ReadBytes(z.Msgs[zdaf].MsgID)
							if err != nil {
								return
							}
						case "m":
							z.Msgs[zdaf].Msg, err = dc.ReadBytes(z.Msgs[zdaf].Msg)
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TextMsgs) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "ci"
	err = en.Append(0x82, 0xa2, 0x63, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Cid)
	if err != nil {
		return
	}
	// write "ms"
	err = en.Append(0xa2, 0x6d, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Msgs)))
	if err != nil {
		return
	}
	for zdaf := range z.Msgs {
		if z.Msgs[zdaf] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 3
			// write "q"
			err = en.Append(0x83, 0xa1, 0x71)
			if err != nil {
				return err
			}
			err = en.WriteInt32(z.Msgs[zdaf].Qos)
			if err != nil {
				return
			}
			// write "mi"
			err = en.Append(0xa2, 0x6d, 0x69)
			if err != nil {
				return err
			}
			err = en.WriteBytes(z.Msgs[zdaf].MsgID)
			if err != nil {
				return
			}
			// write "m"
			err = en.Append(0xa1, 0x6d)
			if err != nil {
				return err
			}
			err = en.WriteBytes(z.Msgs[zdaf].Msg)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TextMsgs) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "ci"
	o = append(o, 0x82, 0xa2, 0x63, 0x69)
	o = msgp.AppendInt64(o, z.Cid)
	// string "ms"
	o = append(o, 0xa2, 0x6d, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Msgs)))
	for zdaf := range z.Msgs {
		if z.Msgs[zdaf] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 3
			// string "q"
			o = append(o, 0x83, 0xa1, 0x71)
			o = msgp.AppendInt32(o, z.Msgs[zdaf].Qos)
			// string "mi"
			o = append(o, 0xa2, 0x6d, 0x69)
			o = msgp.AppendBytes(o, z.Msgs[zdaf].MsgID)
			// string "m"
			o = append(o, 0xa1, 0x6d)
			o = msgp.AppendBytes(o, z.Msgs[zdaf].Msg)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TextMsgs) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zeff uint32
	zeff, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zeff > 0 {
		zeff--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ci":
			z.Cid, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "ms":
			var zrsw uint32
			zrsw, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Msgs) >= int(zrsw) {
				z.Msgs = (z.Msgs)[:zrsw]
			} else {
				z.Msgs = make([]*TextMsg, zrsw)
			}
			for zdaf := range z.Msgs {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Msgs[zdaf] = nil
				} else {
					if z.Msgs[zdaf] == nil {
						z.Msgs[zdaf] = new(TextMsg)
					}
					var zxpk uint32
					zxpk, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zxpk > 0 {
						zxpk--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "q":
							z.Msgs[zdaf].Qos, bts, err = msgp.ReadInt32Bytes(bts)
							if err != nil {
								return
							}
						case "mi":
							z.Msgs[zdaf].MsgID, bts, err = msgp.ReadBytesBytes(bts, z.Msgs[zdaf].MsgID)
							if err != nil {
								return
							}
						case "m":
							z.Msgs[zdaf].Msg, bts, err = msgp.ReadBytesBytes(bts, z.Msgs[zdaf].Msg)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TextMsgs) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 3 + msgp.ArrayHeaderSize
	for zdaf := range z.Msgs {
		if z.Msgs[zdaf] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 2 + msgp.Int32Size + 3 + msgp.BytesPrefixSize + len(z.Msgs[zdaf].MsgID) + 2 + msgp.BytesPrefixSize + len(z.Msgs[zdaf].Msg)
		}
	}
	return
}
