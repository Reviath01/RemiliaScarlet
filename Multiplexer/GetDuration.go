package multiplexer

import (
	"fmt"
	"time"
)

// GetDuration function is for stats command
func GetDuration(duration time.Duration) string {
	return fmt.Sprintf(
		"%0.2d:%0.2d:%0.2d",
		int(duration.Hours()),
		int(duration.Minutes())%60,
		int(duration.Seconds())%60,
	)
}
