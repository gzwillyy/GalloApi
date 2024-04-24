package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeValue = "galloNodeValues"

// NodeValue 节点监控数据
type NodeValue struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	ClusterID         uint32  `gorm:"column:clusterId;comment:集群ID" json:"clusterId"` // 集群ID
	NodeID            uint32  `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`       // 节点ID
	Role              string `gorm:"column:role;comment:节点角色" json:"role"`           // 节点角色
	Item              string `gorm:"column:item;comment:监控项" json:"item"`            // 监控项
	Value             string `gorm:"column:value;comment:数据" json:"value"`           // 数据
	Day               string `gorm:"column:day;comment:日期" json:"day"`               // 日期
	Hour              string `gorm:"column:hour;comment:小时" json:"hour"`             // 小时
	Minute            string `gorm:"column:minute;comment:分钟" json:"minute"`         // 分钟
}

// TableName NodeValue's table name
func (*NodeValue) TableName() string {
	return TableNameNodeValue
}

// AfterCreate run after create database record.
func (u *NodeValue) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "value-")).Error
}

type NodeValueList struct {
	metav1.ListMeta `json:",inline"`
	Items []*NodeValue `json:"items"`
}
var NodeValueTableZeroFields = []string{"name", "item", "value", "day", "hour", "minute"}

