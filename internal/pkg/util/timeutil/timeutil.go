package timeutil

import "time"

// Now return time.Now as a pointer
func Now() *time.Time {
	t := time.Now()
	return &t
}
