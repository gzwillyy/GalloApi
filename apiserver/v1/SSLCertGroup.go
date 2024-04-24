package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameSSLCertGroup = "galloSSLCertGroups"

// SSLCertGroup 证书分组
type SSLCertGroup struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"` // 管理员ID
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`    // 用户ID
	Order             uint32 `gorm:"column:order;comment:分组排序" json:"order"`      // 分组排序
	State             bool   `gorm:"column:state;comment:状态" json:"state"`        // 状态
}

// TableName SSLCertGroup's table name
func (*SSLCertGroup) TableName() string {
	return TableNameSSLCertGroup
}

// AfterCreate run after create database record.
func (u *SSLCertGroup) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "group-")).Error
}

type SSLCertGroupList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*SSLCertGroup `json:"items"`
}

var SSLCertGroupTableZeroFields = []string{"name", "state"}
