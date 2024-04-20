package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNSRoute = "galloNSRoutes"

// NSRoute DNS线路
type NSRoute struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	IsOn              bool   `gorm:"column:isOn;comment:是否启用" json:"isOn"`                   // 是否启用
	ClusterID         uint32 `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`         // 集群ID
	CategoryID        uint32 `gorm:"column:categoryId;comment:分类ID" json:"categoryId"`       // 分类ID
	DomainID          uint32 `gorm:"column:domainId;comment:域名ID" json:"domainId"`           // 域名ID
	AdminID           uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`            // 管理员ID
	UserID            uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`               // 用户ID
	IsPublic          uint32 `gorm:"column:isPublic;comment:是否公用（管理员创建的线路）" json:"isPublic"` // 是否公用（管理员创建的线路）
	Ranges            string `gorm:"column:ranges;comment:范围" json:"ranges"`                 // 范围
	Order             uint32 `gorm:"column:order;comment:排序" json:"order"`                   // 排序
	Version           uint32 `gorm:"column:version;comment:版本号" json:"version"`              // 版本号
	Priority          uint32 `gorm:"column:priority;comment:优先级，越高越优先" json:"priority"`      // 优先级，越高越优先
	Code              string `gorm:"column:code;comment:代号" json:"code"`                     // 代号
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`         // 状态
}

// TableName NSRoute's table name
func (*NSRoute) TableName() string {
	return TableNameNSRoute
}

// AfterCreate run after create database record.
func (u *NSRoute) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "route-")).Error
}

type NSRouteList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NSRoute `json:"items"`
}
