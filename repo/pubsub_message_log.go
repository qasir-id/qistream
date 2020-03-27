package repo

import (
	"github.com/qasir-id/qistream/model"

	"github.com/jinzhu/gorm"
)

// PubSubMessageLogGorm struct
type PubSubMessageLogGorm struct {
}

func (r *PubSubMessageLogGorm) Create(db *gorm.DB, data model.PubSubMessageLog) (model.PubSubMessageLog, error) {
	db.LogMode(true)
	if err := db.Save(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// NewPubSubMessageLogGorm initiate instance
func NewPubSubMessageLogGorm() *PubSubMessageLogGorm {
	var PubSubMessageLog PubSubMessageLogGorm
	return &PubSubMessageLog
}
