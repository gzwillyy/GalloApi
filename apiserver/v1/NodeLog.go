package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeLog = "galloNodeLogs"

// NodeLog 节点日志
type NodeLog struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Role              string `gorm:"column:role;comment:节点角色" json:"role"`               // 节点角色
	Type              string `gorm:"column:type;comment:类型" json:"type"`                 // 类型
	Tag               string `gorm:"column:tag;comment:标签" json:"tag"`                   // 标签
	Description       string `gorm:"column:description;comment:描述" json:"description"`   // 描述
	Level             string `gorm:"column:level;comment:级别" json:"level"`               // 级别
	NodeID            uint32  `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`           // 节点ID
	Day               string `gorm:"column:day;comment:日期" json:"day"`                   // 日期
	ServerID          uint32  `gorm:"column:serverId;comment:服务ID" json:"serverId"`       // 服务ID
	OriginID          uint32  `gorm:"column:originId;comment:源站ID" json:"originId"`       // 源站ID
	Hash              string `gorm:"column:hash;comment:信息内容Hash" json:"hash"`           // 信息内容Hash
	Count             uint32  `gorm:"column:count;comment:重复次数" json:"count"`             // 重复次数
	IsFixed           uint32  `gorm:"column:isFixed;comment:是否已处理" json:"isFixed"`        // 是否已处理
	IsRead            uint32  `gorm:"column:isRead;default:1;comment:是否已读" json:"isRead"` // 是否已读
	Params            string `gorm:"column:params;comment:参数" json:"params"`             // 参数
}

// TableName NodeLog's table name
func (*NodeLog) TableName() string {
	return TableNameNodeLog
}

// AfterCreate run after create database record.
func (u *NodeLog) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "log-")).Error
}

type NodeLogList struct {
	metav1.ListMeta `json:",inline"`
	Items []*NodeLog `json:"items"`
}
var NodeLogTableZeroFields = []string{"name", "role", "type", "tag", "description", "level", "params"}

