package v1

import (
	"time"

	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPCacheTask = "galloHTTPCacheTasks"

// HTTPCacheTask 缓存相关任务
type HTTPCacheTask struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	UserID            uint32    `gorm:"column:userId;comment:用户ID" json:"userId"`         // 用户ID
	Type              string    `gorm:"column:type;comment:任务类型：purge|fetch" json:"type"` // 任务类型：purge|fetch
	KeyType           string    `gorm:"column:keyType;comment:Key类型" json:"keyType"`      // Key类型
	State             bool      `gorm:"column:state;default:1;comment:状态" json:"state"`   // 状态
	DoneAt            time.Time `gorm:"column:doneAt;comment:完成时间" json:"doneAt"`         // 完成时间
	Day               string    `gorm:"column:day;comment:创建日期YYYYMMDD" json:"day"`       // 创建日期YYYYMMDD
	IsDone            uint32    `gorm:"column:isDone;comment:是否已完成" json:"isDone"`        // 是否已完成
	IsOk              uint32    `gorm:"column:isOk;default:1;comment:是否完全成功" json:"isOk"` // 是否完全成功
	IsReady           uint32    `gorm:"column:isReady;comment:是否已准备好" json:"isReady"`     // 是否已准备好
	Description       string    `gorm:"column:description;comment:描述" json:"description"` // 描述
}

// TableName HTTPCacheTask's table name
func (*HTTPCacheTask) TableName() string {
	return TableNameHTTPCacheTask
}

// AfterCreate run after create database record.
func (u *HTTPCacheTask) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "task-")).Error
}

type HTTPCacheTaskList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPCacheTask `json:"items"`
}
