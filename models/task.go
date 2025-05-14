package models

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Completed   bool   `json:"completed" gorm:"default:false"`
	UserID      uint   `json:"user_id"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
}