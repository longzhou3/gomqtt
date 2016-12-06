package global

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DelToken) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "acc":
			z.Acc, err = dc.ReadBytes(z.Acc)
			if err != nil {
				return
			}
		case "ai":
			z.AppID, err = dc.ReadBytes(z.AppID)
			if err != nil {
				return
			}
		case "t":
			z.Token, err = dc.ReadBytes(z.Token)
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
func (z *DelToken) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "acc"
	err = en.Append(0x83, 0xa3, 0x61, 0x63, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Acc)
	if err != nil {
		return
	}
	// write "ai"
	err = en.Append(0xa2, 0x61, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.AppID)
	if err != nil {
		return
	}
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Token)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DelToken) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "acc"
	o = append(o, 0x83, 0xa3, 0x61, 0x63, 0x63)
	o = msgp.AppendBytes(o, z.Acc)
	// string "ai"
	o = append(o, 0xa2, 0x61, 0x69)
	o = msgp.AppendBytes(o, z.AppID)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendBytes(o, z.Token)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DelToken) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "acc":
			z.Acc, bts, err = msgp.ReadBytesBytes(bts, z.Acc)
			if err != nil {
				return
			}
		case "ai":
			z.AppID, bts, err = msgp.ReadBytesBytes(bts, z.AppID)
			if err != nil {
				return
			}
		case "t":
			z.Token, bts, err = msgp.ReadBytesBytes(bts, z.Token)
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
func (z *DelToken) Msgsize() (s int) {
	s = 1 + 4 + msgp.BytesPrefixSize + len(z.Acc) + 3 + msgp.BytesPrefixSize + len(z.AppID) + 2 + msgp.BytesPrefixSize + len(z.Token)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *JsonData) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
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
		case "Msgs":
			var zajw uint32
			zajw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Msgs) >= int(zajw) {
				z.Msgs = (z.Msgs)[:zajw]
			} else {
				z.Msgs = make([]*JsonMsg, zajw)
			}
			for zbai := range z.Msgs {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Msgs[zbai] = nil
				} else {
					if z.Msgs[zbai] == nil {
						z.Msgs[zbai] = new(JsonMsg)
					}
					err = z.Msgs[zbai].DecodeMsg(dc)
					if err != nil {
						return
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
func (z *JsonData) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Msgs"
	err = en.Append(0x81, 0xa4, 0x4d, 0x73, 0x67, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Msgs)))
	if err != nil {
		return
	}
	for zbai := range z.Msgs {
		if z.Msgs[zbai] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Msgs[zbai].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *JsonData) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Msgs"
	o = append(o, 0x81, 0xa4, 0x4d, 0x73, 0x67, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Msgs)))
	for zbai := range z.Msgs {
		if z.Msgs[zbai] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Msgs[zbai].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *JsonData) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Msgs":
			var zhct uint32
			zhct, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Msgs) >= int(zhct) {
				z.Msgs = (z.Msgs)[:zhct]
			} else {
				z.Msgs = make([]*JsonMsg, zhct)
			}
			for zbai := range z.Msgs {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Msgs[zbai] = nil
				} else {
					if z.Msgs[zbai] == nil {
						z.Msgs[zbai] = new(JsonMsg)
					}
					bts, err = z.Msgs[zbai].UnmarshalMsg(bts)
					if err != nil {
						return
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
func (z *JsonData) Msgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zbai := range z.Msgs {
		if z.Msgs[zbai] == nil {
			s += msgp.NilSize
		} else {
			s += z.Msgs[zbai].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *JsonMsg) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "FAcc":
			z.FAcc, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FTopic":
			z.FTopic, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Type":
			z.Type, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Time":
			z.Time, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Nick":
			z.Nick, err = dc.ReadString()
			if err != nil {
				return
			}
		case "MsgID":
			z.MsgID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Msg":
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
func (z *JsonMsg) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "FAcc"
	err = en.Append(0x87, 0xa4, 0x46, 0x41, 0x63, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FAcc)
	if err != nil {
		return
	}
	// write "FTopic"
	err = en.Append(0xa6, 0x46, 0x54, 0x6f, 0x70, 0x69, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FTopic)
	if err != nil {
		return
	}
	// write "Type"
	err = en.Append(0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Type)
	if err != nil {
		return
	}
	// write "Time"
	err = en.Append(0xa4, 0x54, 0x69, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Time)
	if err != nil {
		return
	}
	// write "Nick"
	err = en.Append(0xa4, 0x4e, 0x69, 0x63, 0x6b)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Nick)
	if err != nil {
		return
	}
	// write "MsgID"
	err = en.Append(0xa5, 0x4d, 0x73, 0x67, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.MsgID)
	if err != nil {
		return
	}
	// write "Msg"
	err = en.Append(0xa3, 0x4d, 0x73, 0x67)
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
func (z *JsonMsg) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "FAcc"
	o = append(o, 0x87, 0xa4, 0x46, 0x41, 0x63, 0x63)
	o = msgp.AppendString(o, z.FAcc)
	// string "FTopic"
	o = append(o, 0xa6, 0x46, 0x54, 0x6f, 0x70, 0x69, 0x63)
	o = msgp.AppendString(o, z.FTopic)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt(o, z.Type)
	// string "Time"
	o = append(o, 0xa4, 0x54, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt(o, z.Time)
	// string "Nick"
	o = append(o, 0xa4, 0x4e, 0x69, 0x63, 0x6b)
	o = msgp.AppendString(o, z.Nick)
	// string "MsgID"
	o = append(o, 0xa5, 0x4d, 0x73, 0x67, 0x49, 0x44)
	o = msgp.AppendString(o, z.MsgID)
	// string "Msg"
	o = append(o, 0xa3, 0x4d, 0x73, 0x67)
	o = msgp.AppendBytes(o, z.Msg)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *JsonMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "FAcc":
			z.FAcc, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FTopic":
			z.FTopic, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Time":
			z.Time, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Nick":
			z.Nick, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "MsgID":
			z.MsgID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Msg":
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
func (z *JsonMsg) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.FAcc) + 7 + msgp.StringPrefixSize + len(z.FTopic) + 5 + msgp.IntSize + 5 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Nick) + 6 + msgp.StringPrefixSize + len(z.MsgID) + 4 + msgp.BytesPrefixSize + len(z.Msg)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *JsonMsgs) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "rc":
			z.RetryCount, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "q":
			z.Qos, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "mis":
			var zjfb uint32
			zjfb, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.MsgID) >= int(zjfb) {
				z.MsgID = (z.MsgID)[:zjfb]
			} else {
				z.MsgID = make([][]byte, zjfb)
			}
			for zlqf := range z.MsgID {
				z.MsgID[zlqf], err = dc.ReadBytes(z.MsgID[zlqf])
				if err != nil {
					return
				}
			}
		case "d":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Data = nil
			} else {
				if z.Data == nil {
					z.Data = new(JsonData)
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
					case "Msgs":
						var zeff uint32
						zeff, err = dc.ReadArrayHeader()
						if err != nil {
							return
						}
						if cap(z.Data.Msgs) >= int(zeff) {
							z.Data.Msgs = (z.Data.Msgs)[:zeff]
						} else {
							z.Data.Msgs = make([]*JsonMsg, zeff)
						}
						for zdaf := range z.Data.Msgs {
							if dc.IsNil() {
								err = dc.ReadNil()
								if err != nil {
									return
								}
								z.Data.Msgs[zdaf] = nil
							} else {
								if z.Data.Msgs[zdaf] == nil {
									z.Data.Msgs[zdaf] = new(JsonMsg)
								}
								err = z.Data.Msgs[zdaf].DecodeMsg(dc)
								if err != nil {
									return
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
func (z *JsonMsgs) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "rc"
	err = en.Append(0x84, 0xa2, 0x72, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.RetryCount)
	if err != nil {
		return
	}
	// write "q"
	err = en.Append(0xa1, 0x71)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Qos)
	if err != nil {
		return
	}
	// write "mis"
	err = en.Append(0xa3, 0x6d, 0x69, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.MsgID)))
	if err != nil {
		return
	}
	for zlqf := range z.MsgID {
		err = en.WriteBytes(z.MsgID[zlqf])
		if err != nil {
			return
		}
	}
	// write "d"
	err = en.Append(0xa1, 0x64)
	if err != nil {
		return err
	}
	if z.Data == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 1
		// write "Msgs"
		err = en.Append(0x81, 0xa4, 0x4d, 0x73, 0x67, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Data.Msgs)))
		if err != nil {
			return
		}
		for zdaf := range z.Data.Msgs {
			if z.Data.Msgs[zdaf] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Data.Msgs[zdaf].EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *JsonMsgs) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "rc"
	o = append(o, 0x84, 0xa2, 0x72, 0x63)
	o = msgp.AppendInt32(o, z.RetryCount)
	// string "q"
	o = append(o, 0xa1, 0x71)
	o = msgp.AppendInt32(o, z.Qos)
	// string "mis"
	o = append(o, 0xa3, 0x6d, 0x69, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MsgID)))
	for zlqf := range z.MsgID {
		o = msgp.AppendBytes(o, z.MsgID[zlqf])
	}
	// string "d"
	o = append(o, 0xa1, 0x64)
	if z.Data == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 1
		// string "Msgs"
		o = append(o, 0x81, 0xa4, 0x4d, 0x73, 0x67, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Data.Msgs)))
		for zdaf := range z.Data.Msgs {
			if z.Data.Msgs[zdaf] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Data.Msgs[zdaf].MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *JsonMsgs) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrsw uint32
	zrsw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrsw > 0 {
		zrsw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "rc":
			z.RetryCount, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "q":
			z.Qos, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "mis":
			var zxpk uint32
			zxpk, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.MsgID) >= int(zxpk) {
				z.MsgID = (z.MsgID)[:zxpk]
			} else {
				z.MsgID = make([][]byte, zxpk)
			}
			for zlqf := range z.MsgID {
				z.MsgID[zlqf], bts, err = msgp.ReadBytesBytes(bts, z.MsgID[zlqf])
				if err != nil {
					return
				}
			}
		case "d":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Data = nil
			} else {
				if z.Data == nil {
					z.Data = new(JsonData)
				}
				var zdnj uint32
				zdnj, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zdnj > 0 {
					zdnj--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Msgs":
						var zobc uint32
						zobc, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.Data.Msgs) >= int(zobc) {
							z.Data.Msgs = (z.Data.Msgs)[:zobc]
						} else {
							z.Data.Msgs = make([]*JsonMsg, zobc)
						}
						for zdaf := range z.Data.Msgs {
							if msgp.IsNil(bts) {
								bts, err = msgp.ReadNilBytes(bts)
								if err != nil {
									return
								}
								z.Data.Msgs[zdaf] = nil
							} else {
								if z.Data.Msgs[zdaf] == nil {
									z.Data.Msgs[zdaf] = new(JsonMsg)
								}
								bts, err = z.Data.Msgs[zdaf].UnmarshalMsg(bts)
								if err != nil {
									return
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
func (z *JsonMsgs) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 2 + msgp.Int32Size + 4 + msgp.ArrayHeaderSize
	for zlqf := range z.MsgID {
		s += msgp.BytesPrefixSize + len(z.MsgID[zlqf])
	}
	s += 2
	if z.Data == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 5 + msgp.ArrayHeaderSize
		for zdaf := range z.Data.Msgs {
			if z.Data.Msgs[zdaf] == nil {
				s += msgp.NilSize
			} else {
				s += z.Data.Msgs[zdaf].Msgsize()
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ProtoBufMsg) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zema uint32
	zema, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zema > 0 {
		zema--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "q":
			var zpez uint32
			zpez, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Qos) >= int(zpez) {
				z.Qos = (z.Qos)[:zpez]
			} else {
				z.Qos = make([]int32, zpez)
			}
			for zsnv := range z.Qos {
				z.Qos[zsnv], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "mi":
			var zqke uint32
			zqke, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.MsgIDs) >= int(zqke) {
				z.MsgIDs = (z.MsgIDs)[:zqke]
			} else {
				z.MsgIDs = make([][]byte, zqke)
			}
			for zkgt := range z.MsgIDs {
				z.MsgIDs[zkgt], err = dc.ReadBytes(z.MsgIDs[zkgt])
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
	// map header, size 3
	// write "q"
	err = en.Append(0x83, 0xa1, 0x71)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Qos)))
	if err != nil {
		return
	}
	for zsnv := range z.Qos {
		err = en.WriteInt32(z.Qos[zsnv])
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
	for zkgt := range z.MsgIDs {
		err = en.WriteBytes(z.MsgIDs[zkgt])
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
	// map header, size 3
	// string "q"
	o = append(o, 0x83, 0xa1, 0x71)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Qos)))
	for zsnv := range z.Qos {
		o = msgp.AppendInt32(o, z.Qos[zsnv])
	}
	// string "mi"
	o = append(o, 0xa2, 0x6d, 0x69)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MsgIDs)))
	for zkgt := range z.MsgIDs {
		o = msgp.AppendBytes(o, z.MsgIDs[zkgt])
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
	var zqyh uint32
	zqyh, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zqyh > 0 {
		zqyh--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "q":
			var zyzr uint32
			zyzr, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Qos) >= int(zyzr) {
				z.Qos = (z.Qos)[:zyzr]
			} else {
				z.Qos = make([]int32, zyzr)
			}
			for zsnv := range z.Qos {
				z.Qos[zsnv], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "mi":
			var zywj uint32
			zywj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.MsgIDs) >= int(zywj) {
				z.MsgIDs = (z.MsgIDs)[:zywj]
			} else {
				z.MsgIDs = make([][]byte, zywj)
			}
			for zkgt := range z.MsgIDs {
				z.MsgIDs[zkgt], bts, err = msgp.ReadBytesBytes(bts, z.MsgIDs[zkgt])
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
	s = 1 + 2 + msgp.ArrayHeaderSize + (len(z.Qos) * (msgp.Int32Size)) + 3 + msgp.ArrayHeaderSize
	for zkgt := range z.MsgIDs {
		s += msgp.BytesPrefixSize + len(z.MsgIDs[zkgt])
	}
	s += 2 + msgp.BytesPrefixSize + len(z.Msg)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PubApns) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zjpj uint32
	zjpj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjpj > 0 {
		zjpj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "acc":
			z.Acc, err = dc.ReadBytes(z.Acc)
			if err != nil {
				return
			}
		case "ai":
			z.AppID, err = dc.ReadBytes(z.AppID)
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
		case "jm":
			z.JsonMsg, err = dc.ReadBytes(z.JsonMsg)
			if err != nil {
				return
			}
		case "s":
			z.Sound, err = dc.ReadBytes(z.Sound)
			if err != nil {
				return
			}
		case "b":
			z.Badge, err = dc.ReadInt()
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
func (z *PubApns) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "acc"
	err = en.Append(0x87, 0xa3, 0x61, 0x63, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Acc)
	if err != nil {
		return
	}
	// write "ai"
	err = en.Append(0xa2, 0x61, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.AppID)
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
	// write "jm"
	err = en.Append(0xa2, 0x6a, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.JsonMsg)
	if err != nil {
		return
	}
	// write "s"
	err = en.Append(0xa1, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Sound)
	if err != nil {
		return
	}
	// write "b"
	err = en.Append(0xa1, 0x62)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Badge)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PubApns) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "acc"
	o = append(o, 0x87, 0xa3, 0x61, 0x63, 0x63)
	o = msgp.AppendBytes(o, z.Acc)
	// string "ai"
	o = append(o, 0xa2, 0x61, 0x69)
	o = msgp.AppendBytes(o, z.AppID)
	// string "mi"
	o = append(o, 0xa2, 0x6d, 0x69)
	o = msgp.AppendBytes(o, z.MsgID)
	// string "m"
	o = append(o, 0xa1, 0x6d)
	o = msgp.AppendBytes(o, z.Msg)
	// string "jm"
	o = append(o, 0xa2, 0x6a, 0x6d)
	o = msgp.AppendBytes(o, z.JsonMsg)
	// string "s"
	o = append(o, 0xa1, 0x73)
	o = msgp.AppendBytes(o, z.Sound)
	// string "b"
	o = append(o, 0xa1, 0x62)
	o = msgp.AppendInt(o, z.Badge)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PubApns) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zzpf uint32
	zzpf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zzpf > 0 {
		zzpf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "acc":
			z.Acc, bts, err = msgp.ReadBytesBytes(bts, z.Acc)
			if err != nil {
				return
			}
		case "ai":
			z.AppID, bts, err = msgp.ReadBytesBytes(bts, z.AppID)
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
		case "jm":
			z.JsonMsg, bts, err = msgp.ReadBytesBytes(bts, z.JsonMsg)
			if err != nil {
				return
			}
		case "s":
			z.Sound, bts, err = msgp.ReadBytesBytes(bts, z.Sound)
			if err != nil {
				return
			}
		case "b":
			z.Badge, bts, err = msgp.ReadIntBytes(bts)
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
func (z *PubApns) Msgsize() (s int) {
	s = 1 + 4 + msgp.BytesPrefixSize + len(z.Acc) + 3 + msgp.BytesPrefixSize + len(z.AppID) + 3 + msgp.BytesPrefixSize + len(z.MsgID) + 2 + msgp.BytesPrefixSize + len(z.Msg) + 3 + msgp.BytesPrefixSize + len(z.JsonMsg) + 2 + msgp.BytesPrefixSize + len(z.Sound) + 2 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SetToken) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zgmo uint32
	zgmo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zgmo > 0 {
		zgmo--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "acc":
			z.Acc, err = dc.ReadBytes(z.Acc)
			if err != nil {
				return
			}
		case "ai":
			z.AppID, err = dc.ReadBytes(z.AppID)
			if err != nil {
				return
			}
		case "t":
			z.Token, err = dc.ReadBytes(z.Token)
			if err != nil {
				return
			}
		case "tps":
			var ztaf uint32
			ztaf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Topics) >= int(ztaf) {
				z.Topics = (z.Topics)[:ztaf]
			} else {
				z.Topics = make([][]byte, ztaf)
			}
			for zrfe := range z.Topics {
				z.Topics[zrfe], err = dc.ReadBytes(z.Topics[zrfe])
				if err != nil {
					return
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
func (z *SetToken) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "acc"
	err = en.Append(0x84, 0xa3, 0x61, 0x63, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Acc)
	if err != nil {
		return
	}
	// write "ai"
	err = en.Append(0xa2, 0x61, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.AppID)
	if err != nil {
		return
	}
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Token)
	if err != nil {
		return
	}
	// write "tps"
	err = en.Append(0xa3, 0x74, 0x70, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Topics)))
	if err != nil {
		return
	}
	for zrfe := range z.Topics {
		err = en.WriteBytes(z.Topics[zrfe])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SetToken) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "acc"
	o = append(o, 0x84, 0xa3, 0x61, 0x63, 0x63)
	o = msgp.AppendBytes(o, z.Acc)
	// string "ai"
	o = append(o, 0xa2, 0x61, 0x69)
	o = msgp.AppendBytes(o, z.AppID)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendBytes(o, z.Token)
	// string "tps"
	o = append(o, 0xa3, 0x74, 0x70, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Topics)))
	for zrfe := range z.Topics {
		o = msgp.AppendBytes(o, z.Topics[zrfe])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SetToken) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zeth uint32
	zeth, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zeth > 0 {
		zeth--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "acc":
			z.Acc, bts, err = msgp.ReadBytesBytes(bts, z.Acc)
			if err != nil {
				return
			}
		case "ai":
			z.AppID, bts, err = msgp.ReadBytesBytes(bts, z.AppID)
			if err != nil {
				return
			}
		case "t":
			z.Token, bts, err = msgp.ReadBytesBytes(bts, z.Token)
			if err != nil {
				return
			}
		case "tps":
			var zsbz uint32
			zsbz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Topics) >= int(zsbz) {
				z.Topics = (z.Topics)[:zsbz]
			} else {
				z.Topics = make([][]byte, zsbz)
			}
			for zrfe := range z.Topics {
				z.Topics[zrfe], bts, err = msgp.ReadBytesBytes(bts, z.Topics[zrfe])
				if err != nil {
					return
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
func (z *SetToken) Msgsize() (s int) {
	s = 1 + 4 + msgp.BytesPrefixSize + len(z.Acc) + 3 + msgp.BytesPrefixSize + len(z.AppID) + 2 + msgp.BytesPrefixSize + len(z.Token) + 4 + msgp.ArrayHeaderSize
	for zrfe := range z.Topics {
		s += msgp.BytesPrefixSize + len(z.Topics[zrfe])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TextMsg) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrjx uint32
	zrjx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrjx > 0 {
		zrjx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "fac":
			z.FAcc, err = dc.ReadBytes(z.FAcc)
			if err != nil {
				return
			}
		case "ft":
			z.FTopic, err = dc.ReadBytes(z.FTopic)
			if err != nil {
				return
			}
		case "rc":
			z.RetryCount, err = dc.ReadInt32()
			if err != nil {
				return
			}
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
	// map header, size 6
	// write "fac"
	err = en.Append(0x86, 0xa3, 0x66, 0x61, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.FAcc)
	if err != nil {
		return
	}
	// write "ft"
	err = en.Append(0xa2, 0x66, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.FTopic)
	if err != nil {
		return
	}
	// write "rc"
	err = en.Append(0xa2, 0x72, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.RetryCount)
	if err != nil {
		return
	}
	// write "q"
	err = en.Append(0xa1, 0x71)
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
	// map header, size 6
	// string "fac"
	o = append(o, 0x86, 0xa3, 0x66, 0x61, 0x63)
	o = msgp.AppendBytes(o, z.FAcc)
	// string "ft"
	o = append(o, 0xa2, 0x66, 0x74)
	o = msgp.AppendBytes(o, z.FTopic)
	// string "rc"
	o = append(o, 0xa2, 0x72, 0x63)
	o = msgp.AppendInt32(o, z.RetryCount)
	// string "q"
	o = append(o, 0xa1, 0x71)
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
	var zawn uint32
	zawn, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zawn > 0 {
		zawn--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "fac":
			z.FAcc, bts, err = msgp.ReadBytesBytes(bts, z.FAcc)
			if err != nil {
				return
			}
		case "ft":
			z.FTopic, bts, err = msgp.ReadBytesBytes(bts, z.FTopic)
			if err != nil {
				return
			}
		case "rc":
			z.RetryCount, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
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
	s = 1 + 4 + msgp.BytesPrefixSize + len(z.FAcc) + 3 + msgp.BytesPrefixSize + len(z.FTopic) + 3 + msgp.Int32Size + 2 + msgp.Int32Size + 3 + msgp.BytesPrefixSize + len(z.MsgID) + 2 + msgp.BytesPrefixSize + len(z.Msg)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TextMsgs) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrbe uint32
	zrbe, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrbe > 0 {
		zrbe--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ms":
			var zmfd uint32
			zmfd, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Msgs) >= int(zmfd) {
				z.Msgs = (z.Msgs)[:zmfd]
			} else {
				z.Msgs = make([]*TextMsg, zmfd)
			}
			for zwel := range z.Msgs {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Msgs[zwel] = nil
				} else {
					if z.Msgs[zwel] == nil {
						z.Msgs[zwel] = new(TextMsg)
					}
					err = z.Msgs[zwel].DecodeMsg(dc)
					if err != nil {
						return
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
	// map header, size 1
	// write "ms"
	err = en.Append(0x81, 0xa2, 0x6d, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Msgs)))
	if err != nil {
		return
	}
	for zwel := range z.Msgs {
		if z.Msgs[zwel] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Msgs[zwel].EncodeMsg(en)
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
	// map header, size 1
	// string "ms"
	o = append(o, 0x81, 0xa2, 0x6d, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Msgs)))
	for zwel := range z.Msgs {
		if z.Msgs[zwel] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Msgs[zwel].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TextMsgs) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zzdc uint32
	zzdc, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zzdc > 0 {
		zzdc--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ms":
			var zelx uint32
			zelx, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Msgs) >= int(zelx) {
				z.Msgs = (z.Msgs)[:zelx]
			} else {
				z.Msgs = make([]*TextMsg, zelx)
			}
			for zwel := range z.Msgs {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Msgs[zwel] = nil
				} else {
					if z.Msgs[zwel] == nil {
						z.Msgs[zwel] = new(TextMsg)
					}
					bts, err = z.Msgs[zwel].UnmarshalMsg(bts)
					if err != nil {
						return
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
	s = 1 + 3 + msgp.ArrayHeaderSize
	for zwel := range z.Msgs {
		if z.Msgs[zwel] == nil {
			s += msgp.NilSize
		} else {
			s += z.Msgs[zwel].Msgsize()
		}
	}
	return
}
