package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameSSLPolicy = "galloSSLPolicies"

// SSLPolicy SSL配置策略
type SSLPolicy struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                          // 管理员ID
	UserID            uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`                             // 用户ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                       // 是否启用
	Certs             string `gorm:"column:certs;comment:证书列表" json:"certs"`                               // 证书列表
	ClientCACerts     string `gorm:"column:clientCACerts;comment:客户端证书" json:"clientCACerts"`              // 客户端证书
	ClientAuthType    uint32 `gorm:"column:clientAuthType;comment:客户端认证类型" json:"clientAuthType"`          // 客户端认证类型
	MinVersion        string `gorm:"column:minVersion;comment:支持的SSL最小版本" json:"minVersion"`               // 支持的SSL最小版本
	CipherSuitesIsOn  bool   `gorm:"column:cipherSuitesIsOn;comment:是否自定义加密算法套件" json:"cipherSuitesIsOn"`  // 是否自定义加密算法套件
	CipherSuites      string `gorm:"column:cipherSuites;comment:加密算法套件" json:"cipherSuites"`               // 加密算法套件
	Hsts              string `gorm:"column:hsts;comment:HSTS设置" json:"hsts"`                               // HSTS设置
	Http2Enabled      uint32 `gorm:"column:http2Enabled;default:1;comment:是否启用HTTP/2" json:"http2Enabled"` // 是否启用HTTP/2
	Http3Enabled      uint32 `gorm:"column:http3Enabled;comment:是否启用HTTP/3" json:"http3Enabled"`           // 是否启用HTTP/3
	OcspIsOn          bool   `gorm:"column:ocspIsOn;comment:是否启用OCSP" json:"ocspIsOn"`                     // 是否启用OCSP
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                       // 状态
}

// TableName SSLPolicy's table name
func (*SSLPolicy) TableName() string {
	return TableNameSSLPolicy
}

// AfterCreate run after create database record.
func (u *SSLPolicy) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "policy-")).Error
}

type SSLPolicyList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*SSLPolicy `json:"items"`
}
