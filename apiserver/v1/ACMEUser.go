package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameACMEUser = "galloACMEUsers"

// ACMEUser ACME用户
type ACMEUser struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                               // 管理员ID
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`                                  // 用户ID
	PrivateKey        string `gorm:"column:privateKey;comment:私钥" json:"privateKey"`                            // 私钥
	Email             string `gorm:"column:email;comment:E-mail" json:"email"`                                  // E-mail
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                            // 状态
	Description       string `gorm:"column:description;comment:备注介绍" json:"description"`                        // 备注介绍
	Registration      string `gorm:"column:registration;comment:注册信息" json:"registration,omitempty"`            // 注册信息
	ProviderCode      string `gorm:"column:providerCode;default:letsencrypt;comment:服务商代号" json:"providerCode"` // 服务商代号
	AccountID         uint32 `gorm:"column:accountId;comment:提供商ID" json:"accountId"`                           // 提供商ID
}

// TableName ACMEUser's table name
func (*ACMEUser) TableName() string {
	return TableNameACMEUser
}

// AfterCreate run after create database record.
func (u *ACMEUser) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "user-")).Error
}

// ACMEUserList 返回列表
type ACMEUserList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*ACMEUser `json:"items"`
}

// DeleteACMEUserRequest 删除请求
type DeleteACMEUserRequest struct {
	InstanceID string `json:"instanceID"`
}

// UpdateACMEUserRequest 修改请求
type UpdateACMEUserRequest struct {
	InstanceID  string `json:"instanceID"`
	Description string `json:"description"`
}

var ACMEUserTableZeroFields = []string{"name", "email", "description", "registration", "providerCode"}
