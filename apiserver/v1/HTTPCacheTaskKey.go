package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPCacheTaskKey = "galloHTTPCacheTaskKeys"

// HTTPCacheTaskKey 缓存任务Key
type HTTPCacheTaskKey struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	TaskID            uint32  `gorm:"column:taskId;comment:任务ID" json:"taskId"`               // 任务ID
	Key               string `gorm:"column:key;comment:Key" json:"key"`                      // Key
	KeyType           string `gorm:"column:keyType;comment:Key类型：key|prefix" json:"keyType"` // Key类型：key|prefix
	Type              string `gorm:"column:type;comment:操作类型" json:"type"`                   // 操作类型
	ClusterID         uint32  `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`         // 集群ID
	Nodes             string `gorm:"column:nodes;comment:节点" json:"nodes"`                   // 节点
	Errors            string `gorm:"column:errors;comment:错误信息" json:"errors"`               // 错误信息
	IsDone            uint32  `gorm:"column:isDone;comment:是否已完成" json:"isDone"`              // 是否已完成
}

// TableName HTTPCacheTaskKey's table name
func (*HTTPCacheTaskKey) TableName() string {
	return TableNameHTTPCacheTaskKey
}

// AfterCreate run after create database record.
func (u *HTTPCacheTaskKey) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "key-")).Error
}

type HTTPCacheTaskKeyList struct {
	metav1.ListMeta `json:",inline"`
	Items []*HTTPCacheTaskKey `json:"items"`
}
var HTTPCacheTaskKeyTableZeroFields = []string{"name", "errors"}

