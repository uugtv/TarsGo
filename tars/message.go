package tars

import (
	"time"

	"github.com/uugtv/TarsGo/tars/protocol/res/requestf"
)

// Message is a struct contains servant information
type Message struct {
	Req  *requestf.RequestPacket
	Resp *requestf.ResponsePacket

	Obj *ObjectProxy
	Ser *ServantProxy
	Adp *AdapterProxy

	BeginTime int64
	EndTime   int64
	Status    int32

	hashCode int64
	isHash   bool
}

// Init define the begintime
func (m *Message) Init() {
	m.BeginTime = time.Now().UnixNano() / 1000000
}

// End define the endtime
func (m *Message) End() {
	m.EndTime = time.Now().UnixNano() / 1000000
}

// Cost calculate the cost time
func (m *Message) Cost() int64 {
	return m.EndTime - m.BeginTime
}

// SetHashCode set hash code
func (m *Message) SetHashCode(code int64) {
	m.hashCode = code
	m.isHash = true
}
