package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/qasir-id/qistream/model"
)

// PubSubMessageLogRepo interface
type PubSubMessageLogRepo interface {
	// Create PubSubMessageLog
	Create(*gorm.DB, model.PubSubMessageLog) (model.PubSubMessageLog, error)
}
