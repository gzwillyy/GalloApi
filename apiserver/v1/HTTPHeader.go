package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPHeader = "galloHTTPHeaders"

// HTTPHeader HTTP Header
type HTTPHeader struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                   // 管理员ID
	UserID            uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`                      // 用户ID
	TemplateID        uint32 `gorm:"column:templateId;comment:模版ID" json:"templateId"`              // 模版ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                // 是否启用
	Value             string `gorm:"column:value;comment:值" json:"value"`                           // 值
	Order             uint32 `gorm:"column:order;comment:排序" json:"order"`                          // 排序
	Status            string `gorm:"column:status;comment:状态码设置" json:"status"`                     // 状态码设置
	DisableRedirect   uint32 `gorm:"column:disableRedirect;comment:是否不支持跳转" json:"disableRedirect"` // 是否不支持跳转
	ShouldAppend      uint32 `gorm:"column:shouldAppend;comment:是否为附加" json:"shouldAppend"`         // 是否为附加
	ShouldReplace     uint32 `gorm:"column:shouldReplace;comment:是否替换变量" json:"shouldReplace"`      // 是否替换变量
	ReplaceValues     string `gorm:"column:replaceValues;comment:替换的值" json:"replaceValues"`        // 替换的值
	Methods           string `gorm:"column:methods;comment:支持的方法" json:"methods"`                   // 支持的方法
	Domains           string `gorm:"column:domains;comment:支持的域名" json:"domains"`                   // 支持的域名
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                // 状态
}

// TableName HTTPHeader's table name
func (*HTTPHeader) TableName() string {
	return TableNameHTTPHeader
}

func (u *HTTPHeader) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "header-")).Error
}

type HTTPHeaderList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPHeader `json:"items"`
}
