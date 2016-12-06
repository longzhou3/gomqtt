package global

//easyjson:json
type Messages struct {
	Compress int    `json:"compress"`
	Data     []byte `json:"data"`
}

//easyjson:json
type C2SMsg struct {
	Acc   string `json:"acc"`
	Topic string `json:"topic"`
	Type  int    `json:"type"`
	Qos   int    `json:"qos"`
	MsgID string `json:"msgid"`
	Msg   []byte `json:"msg"`
}
