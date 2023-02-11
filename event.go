package staffing

import "time"

type Event struct {
	ID        string
	Type      string
	Timestamp time.Time
	Body      []byte
}
