package nrmock

import (
	"net/http"

	"net/http/httptest"

	"github.com/newrelic/go-agent"
)

type Transaction struct {
	http.ResponseWriter

	name     string
	WasEnded bool

	attributes map[string]interface{}
}

func NewTransaction(name string) *Transaction {
	return &Transaction{
		name:           name,
		WasEnded:       false,
		attributes:     make(map[string]interface{}),
		ResponseWriter: httptest.NewRecorder(),
	}
}

// interface
func (t *Transaction) End() error {
	t.WasEnded = true
	return nil
}

func (t *Transaction) Ignore() error {
	return nil
}

func (t *Transaction) SetName(name string) error {
	t.name = name
	return nil
}

func (t *Transaction) NoticeError(err error) error {
	return nil
}

func (t *Transaction) StartSegmentNow() newrelic.SegmentStartTime {
	return newrelic.SegmentStartTime{}
}

func (t *Transaction) Header() http.Header {
	return t.ResponseWriter.Header()
}

func (t *Transaction) Write(body []byte) (int, error) {
	return t.ResponseWriter.Write(body)
}

func (t *Transaction) WriteHeader(code int) {
	t.ResponseWriter.WriteHeader(code)
}

func (t *Transaction) AcceptDistributedTracePayload(tp newrelic.TransportType, payload interface{}) error {
	return nil
}

func (t *Transaction) CreateDistributedTracePayload() newrelic.DistributedTracePayload {
	return nil
}

func (t *Transaction) AddAttribute(key string, val interface{}) error {
	if t.attributes == nil {
		t.attributes = make(map[string]interface{})
	}

	t.attributes[key] = val

	return nil
}

// test helpers

func (t *Transaction) GetName() string {
	return t.name
}

func (t *Transaction) GetAttribute(key string) (interface{}, bool) {
	if t.attributes == nil {
		return nil, false
	}

	if val, ok := t.attributes[key]; ok {
		return val, true
	}

	return nil, false
}
