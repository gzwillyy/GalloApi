package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameUpdatingServerList = "galloUpdatingServerLists"

// UpdatingServerList 待更新服务列表
type UpdatingServerList struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	ClusterID         uint32  `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`  // 集群ID
	UniqueID          string `gorm:"column:uniqueId;comment:唯一ID" json:"uniqueId"`    // 唯一ID
	ServerIds         string `gorm:"column:serverIds;comment:服务IDs" json:"serverIds"` // 服务IDs
	Day               string `gorm:"column:day;comment:创建日期" json:"day"`              // 创建日期
}

// TableName UpdatingServerList's table name
func (*UpdatingServerList) TableName() string {
	return TableNameUpdatingServerList
}

// AfterCreate run after create database record.
func (u *UpdatingServerList) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "list-")).Error
}

type UpdatingServerListList struct {
	metav1.ListMeta `json:",inline"`
	Items []*UpdatingServerList `json:"items"`
}