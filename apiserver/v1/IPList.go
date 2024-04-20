package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameIPList = "galloIPLists"

// IPList IP黑白名单
type IPList struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`   // 是否启用
	Type              string `gorm:"column:type;comment:类型" json:"type"`               // 类型
	AdminID           uint32 `gorm:"column:adminId;comment:用户ID" json:"adminId"`       // 用户ID
	UserID            uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`         // 用户ID
	ServerID          uint32 `gorm:"column:serverId;comment:服务ID" json:"serverId"`     // 服务ID
	Code              string `gorm:"column:code;comment:代号" json:"code"`               // 代号
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`   // 状态
	Timeout           string `gorm:"column:timeout;comment:默认超时时间" json:"timeout"`     // 默认超时时间
	Actions           string `gorm:"column:actions;comment:IP触发的动作" json:"actions"`    // IP触发的动作
	Description       string `gorm:"column:description;comment:描述" json:"description"` // 描述
	IsPublic          uint32 `gorm:"column:isPublic;comment:是否公用" json:"isPublic"`     // 是否公用
	IsGlobal          uint32 `gorm:"column:isGlobal;comment:是否全局" json:"isGlobal"`     // 是否全局
}

// TableName IPList's table name
func (*IPList) TableName() string {
	return TableNameIPList
}

// AfterCreate run after create database record.
func (u *IPList) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "list-")).Error
}

type IPListList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*IPList `json:"items"`
}
