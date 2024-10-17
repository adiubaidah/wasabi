package model

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type History struct {
	ID        int       `gorm:"primaryKey;colum:id;autoIncrement"`
	ProcessID uint      `gorm:"colum:process_id;not null"`
	Sender    string    `gorm:"colum:sender;type:varchar(255);not null"`
	Receiver  string    `gorm:"colum:receiver;type:varchar(255);not null"`
	Content   string    `gorm:"colum:content;type:text;not null"`
	RoleAs    string    `gorm:"colum:role_as;type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"colum:created_at;autoCreateTime"`

	Process *Process `gorm:"foreignKey:process_id;references:id"`
}

func DeleteOldHistories(db *gorm.DB, hours int) {
	// Database connection string
	// Prepare the delete statement

	result := db.Exec("DELETE FROM histories WHERE created_at < NOW() - INTERVAL ? HOUR", hours)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Default().Panic(result.Error)
	}
	log.Printf("Deleted %d records\n", result.RowsAffected)
}
