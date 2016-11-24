package global

/* MessagePack protocol */

// ProtoBufMsg publish msg to client
type ProtoBufMsg struct {
	Qos    []int32  `msg:"q"`
	MsgIDs [][]byte `msg:"mi"`
	Msg    []byte   `msg:"m"` // protbuf
}

// TextMsgs
type TextMsgs struct {
	Msgs []*TextMsg `msg:"ms"`
}

// TextMsg package
type TextMsg struct {
	Qos   int32  `msg:"q"`
	MsgID []byte `msg:"mi"`
	Msg   []byte `msg:"m"`
}

type ProtoMsg struct {
	Cid int64  `msg:"ci"`
	Msg []byte `msg:"m"`
}
