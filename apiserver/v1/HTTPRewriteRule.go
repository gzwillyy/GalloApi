package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPRewriteRule = "galloHTTPRewriteRules"

// HTTPRewriteRule 重写规则
type HTTPRewriteRule struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                   // 管理员ID
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`                      // 用户ID
	TemplateID        uint32 `gorm:"column:templateId;comment:模版ID" json:"templateId"`              // 模版ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                // 是否启用
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                // 状态
	Pattern           string `gorm:"column:pattern;comment:匹配规则" json:"pattern"`                    // 匹配规则
	Replace           string `gorm:"column:replace;comment:跳转后的地址" json:"replace"`                  // 跳转后的地址
	Mode              string `gorm:"column:mode;comment:替换模式" json:"mode"`                          // 替换模式
	RedirectStatus    uint32 `gorm:"column:redirectStatus;comment:跳转的状态码" json:"redirectStatus"`    // 跳转的状态码
	ProxyHost         string `gorm:"column:proxyHost;comment:代理的主机名" json:"proxyHost"`              // 代理的主机名
	IsBreak           uint32 `gorm:"column:isBreak;default:1;comment:是否终止解析" json:"isBreak"`        // 是否终止解析
	WithQuery         uint32 `gorm:"column:withQuery;default:1;comment:是否保留URI参数" json:"withQuery"` // 是否保留URI参数
	Conds             string `gorm:"column:conds;comment:匹配条件" json:"conds"`                        // 匹配条件
}

// TableName HTTPRewriteRule's table name
func (*HTTPRewriteRule) TableName() string {
	return TableNameHTTPRewriteRule
}

func (u *HTTPRewriteRule) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "rule-")).Error
}

type HTTPRewriteRuleList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPRewriteRule `json:"items"`
}

var HTTPRewriteRuleTableZeroFields = []string{"name", "state", "pattern", "replace", "mode", "proxyHost", "conds"}
