package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameACMEAuthentication = "galloACMEAuthentications"

// ACMEAuthentication ACME认证
type ACMEAuthentication struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	TaskID            uint32 `gorm:"column:taskId;comment:任务ID" json:"taskId"` // 任务ID
	Domain            string `gorm:"column:domain;comment:域名" json:"domain"`   // 域名
	Token             string `gorm:"column:token;comment:令牌" json:"token"`     // 令牌
	Key               string `gorm:"column:key;comment:密钥" json:"key"`         // 密钥
}

// TableName ACMEAuthentication's table name
func (*ACMEAuthentication) TableName() string {
	return TableNameACMEAuthentication
}

// AfterCreate run after create database record.
func (u *ACMEAuthentication) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "auth-")).Error
}

// ACMEAuthenticationList 返回列表
type ACMEAuthenticationList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*ACMEAuthentication `json:"items"`
}
