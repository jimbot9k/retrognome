package types

import "time"

type RetroIdea struct {
	ID                  int
	Description         string
	Votes               int
	DiscussionStartTime time.Time
	DiscussionEndTime   time.Time
}
