package gate

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/corego/tools"
)

var topicSep = []byte{'-', '-'}

func topicTrans(t []byte) ([]byte, int, error) {
	ts := bytes.Split(t, topicSep)
	if len(ts) != 2 {
		return nil, 0, errors.New("topic for subscribe mus contains just one topic type")
	}

	tp := ts[0]

	var ty int
	// check topic type
	s := tools.Bytes2String(ts[1])
	switch s {
	case "1000", "2000", "3000", "4000", "5000":
		ty, _ = strconv.Atoi(s)
	default:
		return nil, 0, errors.New("invalid topic type")
	}

	return tp, ty, nil
}
