package core

type (
	Retry struct {
		ErrorEquals     []string
		IntervalSeconds int32
		MaxAttempts     int32
		// A number that is the multiplier by which the retry interval increases on each attempt (default 2.0).
		BackoffRate float32
	}

	Catcher struct {
		ErrorEquals []string
		Next        string
		ResultPath  string
	}
)
