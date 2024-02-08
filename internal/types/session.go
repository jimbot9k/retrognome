package types

type Session struct {
	ID     string
	UserID int
	Token  string
}

func (session *Session) IsEmptySession() bool {
	return session.ID == "" && session.UserID == 0 && session.Token == ""
}
