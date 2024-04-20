package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeTask = "galloNodeTasks"

// NodeTask 节点同步任务
type NodeTask struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Role              string `gorm:"column:role;default:node;comment:节点角色" json:"role"`        // 节点角色
	NodeID            uint32  `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`                 // 节点ID
	ClusterID         uint32  `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`           // 集群ID
	ServerID          uint32  `gorm:"column:serverId;comment:服务ID" json:"serverId"`             // 服务ID
	UserID            uint32  `gorm:"column:userId;comment:用户ID" json:"userId"`                 // 用户ID
	Type              string `gorm:"column:type;comment:任务类型" json:"type"`                     // 任务类型
	UniqueID          string `gorm:"column:uniqueId;comment:唯一ID：nodeId@type" json:"uniqueId"` // 唯一ID：nodeId@type
	IsDone            uint32  `gorm:"column:isDone;comment:是否已完成" json:"isDone"`                // 是否已完成
	IsOk              uint32  `gorm:"column:isOk;comment:是否已完成" json:"isOk"`                    // 是否已完成
	Error             string `gorm:"column:error;comment:错误信息" json:"error"`                   // 错误信息
	IsNotified        uint32  `gorm:"column:isNotified;comment:是否已通知更新" json:"isNotified"`      // 是否已通知更新
	Version           uint32  `gorm:"column:version;comment:版本" json:"version"`                 // 版本
}

// TableName NodeTask's table name
func (*NodeTask) TableName() string {
	return TableNameNodeTask
}

// AfterCreate run after create database record.
func (u *NodeTask) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "task-")).Error
}

type NodeTaskList struct {
	metav1.ListMeta `json:",inline"`
	Items []*NodeTask `json:"items"`
}