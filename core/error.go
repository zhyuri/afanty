package core

type Retry struct {
	ErrorEquals     []string `json:"errorEquals,omitempty"`
	IntervalSeconds int32    `json:"intervalSeconds,omitempty"`
	MaxAttempts     int32    `json:"maxAttempts,omitempty"`
	// A number that is the multiplier by which the retry interval increases on each attempt (default 2.0).
	BackoffRate float32 `json:"backoffRate,omitempty"`
}

type Catcher struct {
	ErrorEquals []string `json:"errorEquals,omitempty"`
	Next        string   `json:"next,omitempty"`
	ResultPath  string   `json:"resultPath,omitempty"`
}
