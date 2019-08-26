package timeutil

import (
	"fmt"
	"time"
)

type (
	// Recoder record time of function execution.
	// Note that Recoder is not thread safe.
	Recoder struct {
		method  string
		records map[string]time.Duration
		begin   time.Time
		curr    time.Time
	}

	output struct {
		Method         string                   `json:"method"`
		TotalTimeSpent time.Duration            `json:"total_time_spent"`
		Records        map[string]time.Duration `json:"records"`
	}
)

// NewRecorder new Recoder instance where the clock is set to time.Now
// Note that Recoder is not thread safe.
func NewRecorder(method string) *Recoder {
	return &Recoder{
		method:  method,
		begin:   time.Now(),
		curr:    time.Now(),
		records: make(map[string]time.Duration),
	}
}

// Record the execution time of function f since the last call
func (tr *Recoder) Record(f string) {
	tr.records[f] = time.Since(tr.curr)
	tr.curr = time.Now()
}

// Reset reset the underlying clock to time.Now()
func (tr *Recoder) Reset() {
	tr.curr = time.Now()
}

// String return detail of recoders in JSON format
func (tr *Recoder) String() string {
	output := output{
		Method:         tr.method,
		Records:        tr.records,
		TotalTimeSpent: time.Since(tr.begin),
	}
	return fmt.Sprintf("time consumption, method: %s, total_time_spent: %v, detail: %+v", output.Method, output.TotalTimeSpent, output.Records)
}
