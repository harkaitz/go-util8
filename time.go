package u

import (
	"time"
	"os"
)

var FakedTime time.Time
var FakedTimeEnabled bool = false

func init() {
	var err error
	if os.Getenv("FAKE_TIME") != "" {
		FakedTime, err = time.Parse("2006-01-02", os.Getenv("FAKE_TIME"))
		if err == nil {
			FakedTimeEnabled = true
		}
	}
}

// Now returns the actual time. Can be ovewritten by setting the
// FAKE_TIME environment variable.
func Now() time.Time {
	if FakedTimeEnabled {
		return FakedTime
	}
	return time.Now()
}

