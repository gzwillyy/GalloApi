package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPPage = "galloHTTPPages"

// HTTPPage 特殊页面
type HTTPPage struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                     // 管理员ID
	UserID            uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`                        // 用户ID
	IsOn              bool   `gorm:"column:isOn;comment:是否启用" json:"isOn"`                            // 是否启用
	StatusList        string `gorm:"column:statusList;comment:状态列表" json:"statusList"`                // 状态列表
	URL               string `gorm:"column:url;comment:页面URL" json:"url"`                             // 页面URL
	NewStatus         uint32 `gorm:"column:newStatus;comment:新状态码" json:"newStatus"`                  // 新状态码
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                  // 状态
	Body              string `gorm:"column:body;comment:页面内容" json:"body"`                            // 页面内容
	BodyType          string `gorm:"column:bodyType;default:url;comment:内容类型" json:"bodyType"`        // 内容类型
	ExceptURLPatterns string `gorm:"column:exceptURLPatterns;comment:例外URL" json:"exceptURLPatterns"` // 例外URL
	OnlyURLPatterns   string `gorm:"column:onlyURLPatterns;comment:限制URL" json:"onlyURLPatterns"`     // 限制URL
}

// TableName HTTPPage's table name
func (*HTTPPage) TableName() string {
	return TableNameHTTPPage
}

func (u *HTTPPage) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "page-")).Error
}

type HTTPPageList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPPage `json:"items"`
}
