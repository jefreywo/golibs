package db

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

/*
CREATE TABLE `j_user` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
	`name` varchar(128) NOT NULL DEFAULT '' COMMENT '名称',
	`status` varchar(64) NOT NULL DEFAULT '' COMMENT '状态',
	`balance` int(12) NOT NULL DEFAULT '0' COMMENT '余额',
	`version` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '版本号',
	`create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4
*/
