package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameACMEProviderAccount = "galloACMEProviderAccounts"

// ACMEProviderAccount ACME提供商
type ACMEProviderAccount struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	UserID            int64  `gorm:"column:userId;comment:用户ID" json:"userId"`           // 用户ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`     // 是否启用
	ProviderCode      string `gorm:"column:providerCode;comment:代号" json:"providerCode"` // 代号
	EabKid            string `gorm:"column:eabKid;comment:KID" json:"eabKid"`            // KID
	EabKey            string `gorm:"column:eabKey;comment:Key" json:"eabKey"`            // Key
	Error             string `gorm:"column:error;comment:最后一条错误信息" json:"error"`         // 最后一条错误信息
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`     // 状态
}

// TableName ACMEProviderAccount's table name
func (*ACMEProviderAccount) TableName() string {
	return TableNameACMEProviderAccount
}

// AfterCreate run after create database record.
func (u *ACMEProviderAccount) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "account-")).Error
}

// ACMEProvider ACME提供商
type ACMEProvider struct {
	Name           string `json:"name"  validate:"omitempty"`
	Code           string `json:"code"  validate:"required,min=1,max=100"`
	Description    string `json:"description"  validate:"omitempty"`
	APIURL         string `json:"apiURL"  validate:"omitempty"`
	TestAPIURL     string `json:"testAPIURL"  validate:"omitempty"`
	RequireEAB     bool   `json:"requireEAB"  validate:"omitempty"`
	EABDescription string `json:"eabDescription"  validate:"omitempty"`
}

// ACMEProviderAccountList 返回列表
type ACMEProviderAccountList struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	metav1.ListMeta `json:",inline"`

	// List of secrets
	Items []*ACMEProviderAccount `json:"items"`
}

// DeleteACMEProviderAccountRequest 删除服务商账号请求
type DeleteACMEProviderAccountRequest struct {
	InstanceID string `json:"instanceID"`
}

// UpdateACMEProviderAccountRequest 修改服务商账号请求
type UpdateACMEProviderAccountRequest struct {
	InstanceID string `json:"instanceID"`
	Name       string `json:"name"`
	EabKid     string `json:"eabKid"`
	EabKey     string `json:"eabKey"`
}
