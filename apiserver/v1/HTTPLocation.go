package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPLocation = "galloHTTPLocations"

// HTTPLocation 路由规则配置
type HTTPLocation struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	TemplateID        uint32 `gorm:"column:templateId;comment:模版ID" json:"templateId"`     // 模版ID
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`          // 管理员ID
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`             // 用户ID
	ParentID          uint32 `gorm:"column:parentId;comment:父级ID" json:"parentId"`         // 父级ID
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`       // 状态
	Pattern           string `gorm:"column:pattern;comment:匹配规则" json:"pattern"`           // 匹配规则
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`       // 是否启用
	Description       string `gorm:"column:description;comment:描述" json:"description"`     // 描述
	WebID             uint32 `gorm:"column:webId;comment:Web配置ID" json:"webId"`            // Web配置ID
	ReverseProxy      string `gorm:"column:reverseProxy;comment:反向代理" json:"reverseProxy"` // 反向代理
	URLPrefix         string `gorm:"column:urlPrefix;comment:URL前缀" json:"urlPrefix"`      // URL前缀
	IsBreak           uint32 `gorm:"column:isBreak;comment:是否终止匹配" json:"isBreak"`         // 是否终止匹配
	Conds             string `gorm:"column:conds;comment:匹配条件" json:"conds"`               // 匹配条件
	Domains           string `gorm:"column:domains;comment:专属域名" json:"domains"`           // 专属域名
}

// TableName HTTPLocation's table name
func (*HTTPLocation) TableName() string {
	return TableNameHTTPLocation
}

func (u *HTTPLocation) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "location-")).Error
}

type HTTPLocationList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPLocation `json:"items"`
}

var HTTPLocationTableZeroFields = []string{"name", "pattern", "isOn", "description", "urlPrefix", "conds", "domains"}
