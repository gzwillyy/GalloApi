package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloHTTPFirewallPolicy = "galloHTTPFirewallPolicies"

// HTTPFirewallPolicy HTTP防火墙
type HTTPFirewallPolicy struct {
	metav1.ObjectMeta  `json:"metadata,omitempty"`
	TemplateID         uint64 `gorm:"column:templateId;comment:模版ID" json:"templateId"`                              // 模版ID
	AdminID            uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                                   // 管理员ID
	UserID             uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`                                      // 用户ID
	ServerID           uint64 `gorm:"column:serverId;comment:服务ID" json:"serverId"`                                  // 服务ID
	GroupID            uint64 `gorm:"column:groupId;comment:服务分组ID" json:"groupId"`                                  // 服务分组ID
	State              bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                                // 状态
	IsOn               bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                                // 是否启用
	Description        string `gorm:"column:description;comment:描述" json:"description"`                              // 描述
	Inbound            string `gorm:"column:inbound;comment:入站规则" json:"inbound"`                                    // 入站规则
	Outbound           string `gorm:"column:outbound;comment:出站规则" json:"outbound"`                                  // 出站规则
	BlockOptions       string `gorm:"column:blockOptions;comment:BLOCK动作选项" json:"blockOptions"`                     // BLOCK动作选项
	PageOptions        string `gorm:"column:pageOptions;comment:PAGE动作选项" json:"pageOptions"`                        // PAGE动作选项
	CaptchaOptions     string `gorm:"column:captchaOptions;comment:验证码动作选项" json:"captchaOptions"`                   // 验证码动作选项
	Mode               string `gorm:"column:mode;default:defend;comment:模式" json:"mode"`                             // 模式
	UseLocalFirewall   bool   `gorm:"column:useLocalFirewall;default:1;comment:是否自动使用本地防火墙" json:"useLocalFirewall"` // 是否自动使用本地防火墙
	SynFlood           string `gorm:"column:synFlood;comment:SynFlood防御设置" json:"synFlood"`                          // SynFlood防御设置
	Log                string `gorm:"column:log;comment:日志配置" json:"log"`                                            // 日志配置
	MaxRequestBodySize uint64 `gorm:"column:maxRequestBodySize;comment:可以检查的最大请求内容尺寸" json:"maxRequestBodySize"`     // 可以检查的最大请求内容尺寸
	DenyCountryHTML    string `gorm:"column:denyCountryHTML;comment:区域封禁提示" json:"denyCountryHTML"`                  // 区域封禁提示
	DenyProvinceHTML   string `gorm:"column:denyProvinceHTML;comment:省份封禁提示" json:"denyProvinceHTML"`                // 省份封禁提示
}

// TableName HTTPFirewallPolicy's table name
func (*HTTPFirewallPolicy) TableName() string {
	return TableNameGalloHTTPFirewallRuleGroup
}

// AfterCreate run after create database record.
func (u *HTTPFirewallPolicy) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "policy-")).Error
}

// HTTPFirewallPolicyList 返回列表
type HTTPFirewallPolicyList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPFirewallPolicy `json:"items"`
}
