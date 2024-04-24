package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPBrotliPolicy = "galloHTTPBrotliPolicies"

// HTTPBrotliPolicy Gzip配置
type HTTPBrotliPolicy struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`      // 管理员ID
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`         // 用户ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`   // 是否启用
	Level             uint8  `gorm:"column:level;comment:压缩级别" json:"level"`           // 压缩级别
	MinLength         string `gorm:"column:minLength;comment:可压缩最小值" json:"minLength"` // 可压缩最小值
	MaxLength         string `gorm:"column:maxLength;comment:可压缩最大值" json:"maxLength"` // 可压缩最大值
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`   // 状态
	Conds             string `gorm:"column:conds;comment:条件" json:"conds"`             // 条件
}

// TableName HTTPBrotliPolicy's table name
func (*HTTPBrotliPolicy) TableName() string {
	return TableNameHTTPBrotliPolicy
}

// AfterCreate run after create database record.
func (u *HTTPBrotliPolicy) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "policy-")).Error
}

type HTTPBrotliPolicyList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPBrotliPolicy `json:"items"`
}

var HTTPBrotliPolicyTableZeroFields = []string{"name", "minLength", "maxLength", "state", "conds"}
