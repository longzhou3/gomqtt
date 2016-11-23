package global

/* MessagePack protocol */

// ProtoBufMsg publish msg to client
type ProtoBufMsg struct {
	Cid    int64    `msg:"ci"`
	Qos    []int32  `msg:"q"`
	MsgIDs [][]byte `msg:"mi"`
	Msg    []byte   `msg:"m"` // protbuf
}

// TextMsgs
type TextMsgs struct {
	Cid  int64      `msg:"ci"`
	Msgs []*TextMsg `msg:"ms"`
}

// TextMsg package
type TextMsg struct {
	Qos   int32  `msg:"q"`
	MsgID []byte `msg:"mi"`
	Msg   []byte `msg:"m"`
}
