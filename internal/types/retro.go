package types

import "time"

type Retro struct {
	ID           int
	Title        string
	Description  string
	StartTime    time.Time
	Owner        User
	HostingGroup Group
	Columns      []RetroColumn
	Todos        []Todo
}
