package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloACMETask = "galloACMETasks"

// GalloACMETask ACME任务
type GalloACMETask struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           int64  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`              // 管理员ID
	UserID            int64  `gorm:"column:userId;comment:用户ID" json:"userId"`                 // 用户ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`           // 是否启用
	AcmeUserID        int64  `gorm:"column:acmeUserId;comment:ACME用户ID" json:"acmeUserId"`     // ACME用户ID
	DNSDomain         string `gorm:"column:dnsDomain;comment:DNS主域名" json:"dnsDomain"`         // DNS主域名
	DNSProviderID     int64  `gorm:"column:dnsProviderId;comment:DNS服务商" json:"dnsProviderId"` // DNS服务商
	Domains           string `gorm:"column:domains;comment:证书域名" json:"domains"`               // 证书域名
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`           // 状态
	CertID            int64  `gorm:"column:certId;comment:生成的证书ID" json:"certId"`              // 生成的证书ID
	AutoRenew         bool   `gorm:"column:autoRenew;comment:是否自动更新" json:"autoRenew"`         // 是否自动更新
	AuthType          string `gorm:"column:authType;comment:认证类型" json:"authType"`             // 认证类型
	AuthURL           string `gorm:"column:authURL;comment:认证URL" json:"authURL"`              // 认证URL
}

// TableName GalloACMETask's table name
func (*GalloACMETask) TableName() string {
	return TableNameGalloACMETask
}

// AfterCreate run after create database record.
func (u *GalloACMETask) AfterCreate(tx *gorm.DB) error {
	u.InstanceID = idutil.GetInstanceID(u.ID, "task-")

	return tx.Save(u).Error
}

// ACMETaskList 返回列表
type ACMETaskList struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	metav1.ListMeta `json:",inline"`

	// List of secrets
	Items []*GalloACMETask `json:"items"`
}

type CreateACMETaskRequest struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	UserId            int64    `json:"userId"`
	AcmeUserID        int64    `json:"acmeUserId"`
	DNSProviderID     int64    `json:"dnsProviderId"`
	DNSDomain         string   `json:"dnsDomain"`
	Domains           []string `json:"domains"`
	AutoRenew         bool     `json:"autoRenew"`
	AuthType          string   `json:"authType"`
	AuthURL           string   `json:"authURL"`
}

// DeleteACMETaskRequest 删除请求
type DeleteACMETaskRequest struct {
	InstanceID string `json:"instanceID"`
}

// RunACMETaskRequest 运行请求
type RunACMETaskRequest struct {
	InstanceID string `json:"instanceID"`
}

// UpdateACMETaskRequest 修改请求
type UpdateACMETaskRequest struct {
	InstanceID    string   `json:"instanceID"`
	AcmeUserID    int64    `json:"acmeUserId"`
	DNSProviderID int64    `json:"dnsProviderId"`
	DNSDomain     string   `json:"dnsDomain"`
	Domains       []string `json:"domains"`
	AutoRenew     bool     `json:"autoRenew"`
	AuthURL       string   `json:"authURL"`
}
