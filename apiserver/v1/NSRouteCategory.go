package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNSRouteCategory = "galloNSRouteCategories"

// NSRouteCategory 线路分类
type NSRouteCategory struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"` // 是否启用
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`    // 管理员ID
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`       // 用户ID
	Order             uint32 `gorm:"column:order;comment:排序" json:"order"`           // 排序
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
}

// TableName NSRouteCategory's table name
func (*NSRouteCategory) TableName() string {
	return TableNameNSRouteCategory
}

// AfterCreate run after create database record.
func (u *NSRouteCategory) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "category-")).Error
}

type NSRouteCategoryList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NSRouteCategory `json:"items"`
}

var NSRouteCategoryTableZeroFields = []string{"name", "isOn", "state"}
