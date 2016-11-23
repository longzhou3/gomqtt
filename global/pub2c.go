package global

// Pub2C publish msg to client
type Pub2C struct {
	Cid  int64     `msg:"ci"`
	Msgs []*PubMsg `msg:"ms"`
}

// PubMsg package
type PubMsg struct {
	Qos   int    `msg:"q"`
	MsgID []byte `msg:"mi"`
	Msg   []byte `msg:"m"`
}
