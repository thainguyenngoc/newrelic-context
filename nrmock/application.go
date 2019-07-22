package nrmock

import (
	"net/http"
	"time"

	"github.com/newrelic/go-agent"
)

type NewrelicApp struct {
	Tnx *Transaction
}

func (a *NewrelicApp) StartTransaction(name string, w http.ResponseWriter, r *http.Request) newrelic.Transaction {
	a.Tnx = NewTransaction(name)
	if w != nil {
		a.Tnx.ResponseWriter = w
	}

	return a.Tnx
}

func (a *NewrelicApp) RecordCustomEvent(eventType string, params map[string]interface{}) error {
	return nil
}
func (a *NewrelicApp) WaitForConnection(timeout time.Duration) error {
	return nil
}
func (a *NewrelicApp) RecordCustomMetric(name string, value float64) error {
	return nil
}
func (a *NewrelicApp) Shutdown(timeout time.Duration) {}
