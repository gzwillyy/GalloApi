package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeIPAddressThreshold = "galloNodeIPAddressThresholds"

// NodeIPAddressThreshold IP地址阈值
type NodeIPAddressThreshold struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AddressID         uint32 `gorm:"column:addressId;comment:IP地址ID" json:"addressId"`   // IP地址ID
	Items             string `gorm:"column:items;comment:阈值条目" json:"items"`             // 阈值条目
	Actions           string `gorm:"column:actions;comment:动作" json:"actions"`           // 动作
	NotifiedAt        uint32 `gorm:"column:notifiedAt;comment:上次通知时间" json:"notifiedAt"` // 上次通知时间
	IsMatched         uint32 `gorm:"column:isMatched;comment:上次是否匹配" json:"isMatched"`   // 上次是否匹配
	State             bool   `gorm:"column:state;comment:状态" json:"state"`               // 状态
	Order             uint32 `gorm:"column:order;comment:排序" json:"order"`               // 排序
}

// TableName NodeIPAddressThreshold's table name
func (*NodeIPAddressThreshold) TableName() string {
	return TableNameNodeIPAddressThreshold
}

// AfterCreate run after create database record.
func (u *NodeIPAddressThreshold) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "hold-")).Error
}

type NodeIPAddressThresholdList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NodeIPAddressThreshold `json:"items"`
}
