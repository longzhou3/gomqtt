package global

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Pub2C) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
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
			var zbai uint32
			zbai, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Msgs) >= int(zbai) {
				z.Msgs = (z.Msgs)[:zbai]
			} else {
				z.Msgs = make([]*PubMsg, zbai)
			}
			for zxvk := range z.Msgs {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Msgs[zxvk] = nil
				} else {
					if z.Msgs[zxvk] == nil {
						z.Msgs[zxvk] = new(PubMsg)
					}
					var zcmr uint32
					zcmr, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zcmr > 0 {
						zcmr--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "q":
							z.Msgs[zxvk].Qos, err = dc.ReadInt32()
							if err != nil {
								return
							}
						case "mi":
							z.Msgs[zxvk].MsgID, err = dc.ReadBytes(z.Msgs[zxvk].MsgID)
							if err != nil {
								return
							}
						case "m":
							z.Msgs[zxvk].Msg, err = dc.ReadBytes(z.Msgs[zxvk].Msg)
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
func (z *Pub2C) EncodeMsg(en *msgp.Writer) (err error) {
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
	for zxvk := range z.Msgs {
		if z.Msgs[zxvk] == nil {
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
			err = en.WriteInt32(z.Msgs[zxvk].Qos)
			if err != nil {
				return
			}
			// write "mi"
			err = en.Append(0xa2, 0x6d, 0x69)
			if err != nil {
				return err
			}
			err = en.WriteBytes(z.Msgs[zxvk].MsgID)
			if err != nil {
				return
			}
			// write "m"
			err = en.Append(0xa1, 0x6d)
			if err != nil {
				return err
			}
			err = en.WriteBytes(z.Msgs[zxvk].Msg)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Pub2C) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "ci"
	o = append(o, 0x82, 0xa2, 0x63, 0x69)
	o = msgp.AppendInt64(o, z.Cid)
	// string "ms"
	o = append(o, 0xa2, 0x6d, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Msgs)))
	for zxvk := range z.Msgs {
		if z.Msgs[zxvk] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 3
			// string "q"
			o = append(o, 0x83, 0xa1, 0x71)
			o = msgp.AppendInt32(o, z.Msgs[zxvk].Qos)
			// string "mi"
			o = append(o, 0xa2, 0x6d, 0x69)
			o = msgp.AppendBytes(o, z.Msgs[zxvk].MsgID)
			// string "m"
			o = append(o, 0xa1, 0x6d)
			o = msgp.AppendBytes(o, z.Msgs[zxvk].Msg)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Pub2C) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
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
			var zwht uint32
			zwht, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Msgs) >= int(zwht) {
				z.Msgs = (z.Msgs)[:zwht]
			} else {
				z.Msgs = make([]*PubMsg, zwht)
			}
			for zxvk := range z.Msgs {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Msgs[zxvk] = nil
				} else {
					if z.Msgs[zxvk] == nil {
						z.Msgs[zxvk] = new(PubMsg)
					}
					var zhct uint32
					zhct, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zhct > 0 {
						zhct--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "q":
							z.Msgs[zxvk].Qos, bts, err = msgp.ReadInt32Bytes(bts)
							if err != nil {
								return
							}
						case "mi":
							z.Msgs[zxvk].MsgID, bts, err = msgp.ReadBytesBytes(bts, z.Msgs[zxvk].MsgID)
							if err != nil {
								return
							}
						case "m":
							z.Msgs[zxvk].Msg, bts, err = msgp.ReadBytesBytes(bts, z.Msgs[zxvk].Msg)
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
func (z *Pub2C) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 3 + msgp.ArrayHeaderSize
	for zxvk := range z.Msgs {
		if z.Msgs[zxvk] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 2 + msgp.Int32Size + 3 + msgp.BytesPrefixSize + len(z.Msgs[zxvk].MsgID) + 2 + msgp.BytesPrefixSize + len(z.Msgs[zxvk].Msg)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PubMsg) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcua uint32
	zcua, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcua > 0 {
		zcua--
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
func (z *PubMsg) EncodeMsg(en *msgp.Writer) (err error) {
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
func (z *PubMsg) MarshalMsg(b []byte) (o []byte, err error) {
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
func (z *PubMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zxhx uint32
	zxhx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
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
func (z *PubMsg) Msgsize() (s int) {
	s = 1 + 2 + msgp.Int32Size + 3 + msgp.BytesPrefixSize + len(z.MsgID) + 2 + msgp.BytesPrefixSize + len(z.Msg)
	return
}
