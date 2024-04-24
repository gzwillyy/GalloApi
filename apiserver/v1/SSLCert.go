package v1

import (
	"time"

	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameSSLCert = "galloSSLCerts"

// SSLCert SSL证书
type SSLCert struct {
	metav1.ObjectMeta  `json:"metadata,omitempty"`
	AdminID            uint32    `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                          // 管理员ID
	UserID             uint32    `gorm:"column:userId;comment:用户ID" json:"userId"`                             // 用户ID
	State              bool      `gorm:"column:state;default:1;comment:状态" json:"state"`                       // 状态
	IsOn               bool      `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                       // 是否启用
	Description        string    `gorm:"column:description;comment:描述" json:"description"`                     // 描述
	CertData           []byte    `gorm:"column:certData;comment:证书内容" json:"certData"`                         // 证书内容
	KeyData            []byte    `gorm:"column:keyData;comment:密钥内容" json:"keyData"`                           // 密钥内容
	ServerName         string    `gorm:"column:serverName;comment:证书使用的主机名" json:"serverName"`                 // 证书使用的主机名
	IsCA               uint32    `gorm:"column:isCA;comment:是否为CA证书" json:"isCA"`                              // 是否为CA证书
	GroupIds           string    `gorm:"column:groupIds;comment:证书分组" json:"groupIds"`                         // 证书分组
	TimeBeginAt        time.Time `gorm:"column:timeBeginAt;comment:开始时间" json:"timeBeginAt"`                   // 开始时间
	TimeEndAt          time.Time `gorm:"column:timeEndAt;comment:结束时间" json:"timeEndAt"`                       // 结束时间
	DNSNames           string    `gorm:"column:dnsNames;comment:DNS名称列表" json:"dnsNames"`                      // DNS名称列表
	CommonNames        string    `gorm:"column:commonNames;comment:发行单位列表" json:"commonNames"`                 // 发行单位列表
	IsACME             uint32    `gorm:"column:isACME;comment:是否为ACME自动生成的" json:"isACME"`                     // 是否为ACME自动生成的
	AcmeTaskID         uint32    `gorm:"column:acmeTaskId;comment:ACME任务ID" json:"acmeTaskId"`                 // ACME任务ID
	NotifiedAt         time.Time `gorm:"column:notifiedAt;comment:最后通知时间" json:"notifiedAt"`                   // 最后通知时间
	Ocsp               []byte    `gorm:"column:ocsp;comment:OCSP缓存" json:"ocsp"`                               // OCSP缓存
	OcspIsUpdated      uint32    `gorm:"column:ocspIsUpdated;comment:OCSP是否已更新" json:"ocspIsUpdated"`          // OCSP是否已更新
	OcspUpdatedAt      time.Time `gorm:"column:ocspUpdatedAt;comment:OCSP更新时间" json:"ocspUpdatedAt"`           // OCSP更新时间
	OcspError          string    `gorm:"column:ocspError;comment:OCSP更新错误" json:"ocspError"`                   // OCSP更新错误
	OcspUpdatedVersion uint32    `gorm:"column:ocspUpdatedVersion;comment:OCSP更新版本" json:"ocspUpdatedVersion"` // OCSP更新版本
	OcspExpiresAt      time.Time `gorm:"column:ocspExpiresAt;comment:OCSP过期时间(UTC)" json:"ocspExpiresAt"`      // OCSP过期时间(UTC)
	OcspTries          uint32    `gorm:"column:ocspTries;comment:OCSP尝试次数" json:"ocspTries"`                   // OCSP尝试次数
}

// TableName SSLCert's table name
func (*SSLCert) TableName() string {
	return TableNameSSLCert
}

// AfterCreate run after create database record.
func (u *SSLCert) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "cert-")).Error
}

type SSLCertList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*SSLCert `json:"items"`
}

var SSLCertTableZeroFields = []string{"name", "isOn", "description", "serverName", "groupIds", "dnsNames", "commonNames", "ocspError"}

