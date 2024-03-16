package database

import (
	"chatbot/internal/config/global"
	"chatbot/pkg/utils"
	"fmt"
)

type ChatRecord struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Type        uint8  `gorm:"type:tinyint(2);not null;default:0" json:"type"`
	UserID      uint32 `gorm:"not null" json:"user_id"`
	BotProvider uint8  `gorm:"type:tinyint(2);not null;default:0" json:"bot_provider"`
	BotModel    string `gorm:"type:varchar(50);not null;default:''" json:"bot_model"`
	Msg         string `gorm:"type:varchar(500);not null;default:''" json:"msg"`
	CreatedAt   uint64 `gorm:"not null;default:0" json:"created_at"`
}

/**
* 如果数据量非常庞大时读写会影响性能，导致接口响应慢，可以根据业务情况来提前规划
* 1. 可以根据 user_id % num 来hash分表
* 2. 每三个月将之前的历史数据归档，并在低峰期慢慢从业务表中清除
 */

type ChatRecordService struct{}

func getTableName(userID uint32) string {
	tableIndex := userID % 10 // 假设要分10个表
	return fmt.Sprintf("chat_records_%d", tableIndex)
}

func (s *ChatRecordService) CreateChatRecord(record ChatRecord) error {
	db := GetDB("account")
	tableName := getTableName(record.UserID)
	return db.Table(tableName).Create(&record).Error
}

func (s *ChatRecordService) DeleteChatRecord(userID uint32, id uint64) error {
	db := GetDB("account")
	tableName := getTableName(userID)
	return db.Table(tableName).Delete(&ChatRecord{}, id).Error
}

func (s *ChatRecordService) UpdateChatRecord(record ChatRecord) error {
	db := GetDB("account")
	tableName := getTableName(record.UserID)
	return db.Table(tableName).Save(&record).Error
}

func (s *ChatRecordService) GetChatRecord(id uint64) (*ChatRecord, error) {
	db := GetDB("account")
	var record ChatRecord
	tableName := getTableName(record.UserID)
	err := db.Table(tableName).First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// 分页从新到旧查询某个用户相关的聊天记录，index 为上次查询的最小created_at, 这样查询可以减少检索rows提升查询性能
func (s *ChatRecordService) GetChatRecordsByPage(userID uint32, fields string, pageSize int, index uint64) ([]ChatRecord, error) {
	db := GetDB("account")
	var records []ChatRecord
	tableName := getTableName(userID)

	var err error

	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}

	if index == 0 {
		index = utils.Now()
	}

	err = db.Table(tableName).Select(fields).Where("user_id = ? and created_at < ?", userID, index).Order("created_at desc").Limit(pageSize).Find(&records).Error

	if err != nil {
		global.Logger.Error("查询聊天记录失败: " + err.Error())
		return nil, err
	}

	return records, nil
}
