package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloHTTPFirewallRuleSet = "galloHTTPFirewallRuleSets"

// HTTPFirewallRuleSet 防火墙规则集
type HTTPFirewallRuleSet struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`              // 是否启用
	Code              string `gorm:"column:code;comment:代号" json:"code"`                          // 代号
	Description       string `gorm:"column:description;comment:描述" json:"description"`            // 描述
	Rules             string `gorm:"column:rules;comment:规则列表" json:"rules"`                      // 规则列表
	Connector         string `gorm:"column:connector;comment:规则之间的关系" json:"connector"`           // 规则之间的关系
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`              // 状态
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                 // 管理员ID
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`                    // 用户ID
	Action            string `gorm:"column:action;comment:执行的动作（过期）" json:"action"`               // 执行的动作（过期）
	ActionOptions     string `gorm:"column:actionOptions;comment:动作的选项（过期）" json:"actionOptions"` // 动作的选项（过期）
	Actions           string `gorm:"column:actions;comment:一组动作" json:"actions"`                  // 一组动作
	IgnoreLocal       bool   `gorm:"column:ignoreLocal;comment:忽略局域网请求" json:"ignoreLocal"`       // 忽略局域网请求
}

// TableName HTTPFirewallRuleSet's table name
func (*HTTPFirewallRuleSet) TableName() string {
	return TableNameGalloHTTPFirewallRuleSet
}

// AfterCreate run after create database record.
func (u *HTTPFirewallRuleSet) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "set-")).Error
}
