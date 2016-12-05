package global

/* MessagePack protocol */

//-------------------------PUB2C-----------------------------------------
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
	FAcc       []byte `msg:"fac"`
	FTopic     []byte `msg:"ft"`
	RetryCount int32  `msg:"rc"`
	Qos        int32  `msg:"q"`
	MsgID      []byte `msg:"mi"`
	Msg        []byte `msg:"m"`
}

type JsonMsgs struct {
	RetryCount int32  `msg:"rc"`
	Qos        int32  `msg:"q"`
	Msg        []byte `msg:"m"`
}

//-------------------------APNS-----------------------------------------
type PubApns struct {
	Acc   []byte `msg:"acc"`
	AppID []byte `msg:"ai"`

	MsgID []byte `msg:"mi"`

	Msg     []byte `msg:"m"`
	JsonMsg []byte `msg:"jm"`

	Sound []byte `msg:"s"`
	Badge int    `msg:"b"`
}

type SetToken struct {
	Acc   []byte `msg:"acc"`
	AppID []byte `msg:"ai"`

	Token []byte `msg:"t"`
	// token对应的广播Topics
	Topics [][]byte `msg:"tps"`
}

type DelToken struct {
	Acc   []byte `msg:"acc"`
	AppID []byte `msg:"ai"`

	Token []byte `msg:"t"`
}
