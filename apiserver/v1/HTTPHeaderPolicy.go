package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPHeaderPolicy = "galloHTTPHeaderPolicies"

// HTTPHeaderPolicy Header 策略
type HTTPHeaderPolicy struct {
	metav1.ObjectMeta  `json:"metadata,omitempty"`
	IsOn               bool   `gorm:"column:isOn;not null;default:1;comment:是否启用" json:"isOn"`               // 是否启用
	State              bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                        // 状态
	AdminID            uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                           // 管理员ID
	UserID             uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`                              // 用户ID
	AddHeaders         string `gorm:"column:addHeaders;comment:添加的Header" json:"addHeaders"`                 // 添加的Header
	AddTrailers        string `gorm:"column:addTrailers;comment:添加的Trailers" json:"addTrailers"`             // 添加的Trailers
	SetHeaders         string `gorm:"column:setHeaders;comment:设置Header" json:"setHeaders"`                  // 设置Header
	ReplaceHeaders     string `gorm:"column:replaceHeaders;comment:替换Header内容" json:"replaceHeaders"`        // 替换Header内容
	Expires            string `gorm:"column:expires;comment:Expires单独设置" json:"expires"`                     // Expires单独设置
	DeleteHeaders      string `gorm:"column:deleteHeaders;comment:删除的Headers" json:"deleteHeaders"`          // 删除的Headers
	NonStandardHeaders string `gorm:"column:nonStandardHeaders;comment:非标Headers" json:"nonStandardHeaders"` // 非标Headers
	Cors               string `gorm:"column:cors;comment:CORS配置" json:"cors"`                                // CORS配置
}

// TableName HTTPHeaderPolicy's table name
func (*HTTPHeaderPolicy) TableName() string {
	return TableNameHTTPHeaderPolicy
}

func (u *HTTPHeaderPolicy) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "policy-")).Error
}

type HTTPHeaderPolicyList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPHeaderPolicy `json:"items"`
}

var HTTPHeaderPolicyTableZeroFields = []string{"name", "isOn", "state", "expires", "deleteHeaders"}

