package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNSKey = "galloNSKeys"

// NSKey 密钥管理
type NSKey struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	IsOn              bool   `gorm:"column:isOn;default:1;comment:状态" json:"isOn"`     // 状态
	DomainID          uint32 `gorm:"column:domainId;comment:域名ID" json:"domainId"`     // 域名ID
	ZoneID            uint32 `gorm:"column:zoneId;comment:子域ID" json:"zoneId"`         // 子域ID
	Algo              string `gorm:"column:algo;comment:算法" json:"algo"`               // 算法
	Secret            string `gorm:"column:secret;comment:密码" json:"secret"`           // 密码
	SecretType        string `gorm:"column:secretType;comment:密码类型" json:"secretType"` // 密码类型
	Version           uint32 `gorm:"column:version;comment:版本号" json:"version"`        // 版本号
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`   // 状态
}

// TableName NSKey's table name
func (*NSKey) TableName() string {
	return TableNameNSKey
}

// AfterCreate run after create database record.
func (u *NSKey) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "key-")).Error
}

type NSKeyList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NSKey `json:"items"`
}

var NSKeyTableZeroFields = []string{"name", "isOn", "secret", "secretType", "state"}

