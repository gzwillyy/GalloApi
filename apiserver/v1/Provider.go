package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameProvider = "galloProviders"

// Provider 供应商
type Provider struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Username          string `gorm:"column:username;comment:用户名" json:"username"`    // 用户名
	Password          string `gorm:"column:password;comment:密码" json:"password"`     // 密码
	Fullname          string `gorm:"column:fullname;comment:真实姓名" json:"fullname"`   // 真实姓名
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
}

// TableName Provider's table name
func (*Provider) TableName() string {
	return TableNameProvider
}

// AfterCreate run after create database record.
func (u *Provider) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "provider-")).Error
}

type ProviderList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*Provider `json:"items"`
}

var ProviderTableZeroFields = []string{"name", "username", "password", "fullname", "state"}

