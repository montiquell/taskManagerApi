package repository

type TaskModel struct {
	ID         int64  `gorm:"primaryKey;type:uuid"`
	Title      string `gorm:"not null"`
	Compeleted bool   `gorm:"not null"`
}
