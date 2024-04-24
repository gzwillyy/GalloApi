package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPAuthPolicy = "galloHTTPAuthPolicies"

// HTTPAuthPolicy HTTP认证策略
type HTTPAuthPolicy struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`    // 管理员ID
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`       // 用户ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"` // 是否启用
	Type              string `gorm:"column:type;comment:类型" json:"type"`             // 类型
	Params            string `gorm:"column:params;comment:参数" json:"params"`         // 参数
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
}

// TableName HTTPAuthPolicy's table name
func (*HTTPAuthPolicy) TableName() string {
	return TableNameHTTPAuthPolicy
}

// AfterCreate run after create database record.
func (u *HTTPAuthPolicy) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "policy-")).Error
}

type HTTPAuthPolicyList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPAuthPolicy `json:"items"`
}

var HTTPAuthPolicyTableZeroFields = []string{"name", "type", "params", "state"}
