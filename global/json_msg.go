package global

//easyjson:json
type Messages struct {
	Compress int    `json:"compress"`
	Data     []byte `json:"data"`
}

//easyjson:json
type JsonMsg struct {
	FAcc   []byte `json:"facc"`
	FTopic []byte `json:"ftopic"`
	Type   int    `json:"type"`
	Qos    int    `json:"qos"`
	Time   int    `json:"time"`
	Nick   []byte `json:"nick"`
	MsgID  []byte `json:"mi"`
	Msg    []byte `json:"m"`
}
