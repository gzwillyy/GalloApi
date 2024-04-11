package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloACMEUser = "galloACMEUsers"

// GalloACMEUser ACME用户
type GalloACMEUser struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           int32  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                               // 管理员ID
	UserID            int32  `gorm:"column:userId;comment:用户ID" json:"userId"`                                  // 用户ID
	PrivateKey        string `gorm:"column:privateKey;comment:私钥" json:"privateKey"`                            // 私钥
	Email             string `gorm:"column:email;comment:E-mail" json:"email"`                                  // E-mail
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                            // 状态
	Description       string `gorm:"column:description;comment:备注介绍" json:"description"`                        // 备注介绍
	Registration      string `gorm:"column:registration;comment:注册信息" json:"registration,omitempty"`            // 注册信息
	ProviderCode      string `gorm:"column:providerCode;default:letsencrypt;comment:服务商代号" json:"providerCode"` // 服务商代号
	AccountID         int64  `gorm:"column:accountId;comment:提供商ID" json:"accountId"`                           // 提供商ID
}

// TableName GalloACMEUser's table name
func (*GalloACMEUser) TableName() string {
	return TableNameGalloACMEUser
}

// AfterCreate run after create database record.
func (u *GalloACMEUser) AfterCreate(tx *gorm.DB) error {
	u.InstanceID = idutil.GetInstanceID(u.ID, "task-")

	return tx.Save(u).Error
}

// ACMEUserList 返回列表
type ACMEUserList struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	metav1.ListMeta `json:",inline"`

	// List of secrets
	Items []*GalloACMEUser `json:"items"`
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
