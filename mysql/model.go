package mysql

import "time"

type JUser struct {
	Id        int64     `gorm:"column:id;not null"`
	Name      string    `gorm:"column:name;not null"`
	Status    string    `gorm:"column:status;not null"`
	Balance   int64     `gorm:"column:balance;not null"`
	Version   int64     `gorm:"column:version;not null"`
	CreatedAt time.Time `gorm:"column:create_at;not null"`
	UpdatedAt time.Time `gorm:"column:update_at;not null"`
}

func (JUser) TableName() string {
	return "j_user"
}
