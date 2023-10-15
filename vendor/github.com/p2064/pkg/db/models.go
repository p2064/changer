package db

type Event struct {
	ID         int64 `gorm:"primaryKey"`
	Place      string
	EventTime  string
	MaxPlayers int64
}

type UserEvent struct {
	UserID  int64
	EventID int64
}
