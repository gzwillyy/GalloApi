package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeRegion = "galloNodeRegions"

// NodeRegion 节点区域
type NodeRegion struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`      // 管理员ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`   // 是否启用
	Description       string `gorm:"column:description;comment:描述" json:"description"` // 描述
	Order             uint32 `gorm:"column:order;comment:排序" json:"order"`             // 排序
	Prices            string `gorm:"column:prices;comment:流量价格" json:"prices"`         // 流量价格
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`   // 状态
}

// TableName NodeRegion's table name
func (*NodeRegion) TableName() string {
	return TableNameNodeRegion
}

// AfterCreate run after create database record.
func (u *NodeRegion) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "region-")).Error
}

type NodeRegionList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NodeRegion `json:"items"`
}

var NodeRegionTableZeroFields = []string{"name", "description", "prices", "state"}
