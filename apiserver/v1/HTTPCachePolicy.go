package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPCachePolicy = "galloHTTPCachePolicies"

// HTTPCachePolicy HTTP缓存策略
type HTTPCachePolicy struct {
	metav1.ObjectMeta    `json:"metadata,omitempty"`
	AdminID              uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                                // 管理员ID
	UserID               uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`                                   // 用户ID
	TemplateID           uint32 `gorm:"column:templateId;comment:模版ID" json:"templateId"`                           // 模版ID
	IsOn                 bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                             // 是否启用
	Capacity             string `gorm:"column:capacity;comment:容量数据" json:"capacity"`                               // 容量数据
	MaxKeys              uint32 `gorm:"column:maxKeys;comment:最多Key值" json:"maxKeys"`                               // 最多Key值
	MaxSize              string `gorm:"column:maxSize;comment:最大缓存内容尺寸" json:"maxSize"`                             // 最大缓存内容尺寸
	Type                 string `gorm:"column:type;comment:存储类型" json:"type"`                                       // 存储类型
	Options              string `gorm:"column:options;comment:存储选项" json:"options"`                                 // 存储选项
	State                bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                             // 状态
	Description          string `gorm:"column:description;comment:描述" json:"description"`                           // 描述
	Refs                 string `gorm:"column:refs;comment:默认的缓存设置" json:"refs"`                                    // 默认的缓存设置
	SyncCompressionCache bool   `gorm:"column:syncCompressionCache;comment:是否同步写入压缩缓存" json:"syncCompressionCache"` // 是否同步写入压缩缓存
	FetchTimeout         string `gorm:"column:fetchTimeout;comment:预热超时时间" json:"fetchTimeout"`                     // 预热超时时间
}

// TableName HTTPCachePolicy's table name
func (*HTTPCachePolicy) TableName() string {
	return TableNameHTTPCachePolicy
}

// AfterCreate run after create database record.
func (u *HTTPCachePolicy) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "policy-")).Error
}

type HTTPCachePolicyList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPCachePolicy `json:"items"`
}

var HTTPCachePolicyTableZeroFields = []string{"name", "capacity", "maxSize", "type", "options", "state", "description", "refs", "fetchTimeout"}
