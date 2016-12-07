package global

//easyjson:json
type Messages struct {
	Compress int    `json:"compress"`
	Data     []byte `json:"data"`
}

//easyjson:json
type C2SMsg struct {
	ToAcc string `json:"toacc"`
	Type  int    `json:"type"`
	MsgID string `json:"msgid"`
	Msg   []byte `json:"msg"`
}
