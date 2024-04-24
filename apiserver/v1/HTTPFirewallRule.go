package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPFirewallRule = "galloHTTPFirewallRules"

// HTTPFirewallRule 防火墙规则
type HTTPFirewallRule struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                               // 是否启用
	Description       string `gorm:"column:description;comment:说明" json:"description"`                             // 说明
	Param             string `gorm:"column:param;comment:参数" json:"param"`                                         // 参数
	ParamFilters      string `gorm:"column:paramFilters;comment:处理器" json:"paramFilters"`                          // 处理器
	Operator          string `gorm:"column:operator;comment:操作符" json:"operator"`                                  // 操作符
	Value             string `gorm:"column:value;comment:对比值" json:"value"`                                        // 对比值
	IsCaseInsensitive uint32 `gorm:"column:isCaseInsensitive;default:1;comment:是否大小写不敏感" json:"isCaseInsensitive"` // 是否大小写不敏感
	CheckpointOptions string `gorm:"column:checkpointOptions;comment:检查点参数" json:"checkpointOptions"`              // 检查点参数
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                               // 状态
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                                  // 管理员ID
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`                                     // 用户ID
}

// TableName HTTPFirewallRule's table name
func (*HTTPFirewallRule) TableName() string {
	return TableNameHTTPFirewallRule
}

// AfterCreate run after create database record.
func (u *HTTPFirewallRule) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "rule-")).Error
}

type HTTPFirewallRuleList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPFirewallRule `json:"items"`
}

var HTTPFirewallRuleTableZeroFields = []string{"name", "isOn", "description", "param", "paramFilters", "operator", "value", "checkpointOptions", "state"}
