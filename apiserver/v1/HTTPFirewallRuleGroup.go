package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPFirewallRuleGroup = "galloHTTPFirewallRuleGroups"

// HTTPFirewallRuleGroup 防火墙规则分组
type HTTPFirewallRuleGroup struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`      // 是否启用
	Description       string `gorm:"column:description;comment:描述" json:"description"`    // 描述
	Code              string `gorm:"column:code;comment:代号" json:"code"`                  // 代号
	IsTemplate        uint32 `gorm:"column:isTemplate;comment:是否为预置模板" json:"isTemplate"` // 是否为预置模板
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`         // 管理员ID
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`            // 用户ID
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`      // 状态
	Sets              string `gorm:"column:sets;comment:规则集列表" json:"sets"`               // 规则集列表
}

// TableName HTTPFirewallRuleGroup's table name
func (*HTTPFirewallRuleGroup) TableName() string {
	return TableNameHTTPFirewallRuleGroup
}

// AfterCreate run after create database record.
func (u *HTTPFirewallRuleGroup) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "group-")).Error
}

type HTTPFirewallRuleGroupList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPFirewallRuleGroup `json:"items"`
}

var HTTPFirewallRuleGroupTableZeroFields = []string{"name", "isOn", "description", "code", "sets"}
