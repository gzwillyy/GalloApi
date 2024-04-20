package v1

import (
	"time"

	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameIPItem = "galloIPItems"

// IPItem IP
type IPItem struct {
	metav1.ObjectMeta             `json:"metadata,omitempty"`
	ListID                        uint32    `gorm:"column:listId;comment:所属名单ID" json:"listId"`                                                  // 所属名单ID
	Type                          string    `gorm:"column:type;default:ipv4;comment:类型" json:"type"`                                             // 类型
	IPFrom                        string    `gorm:"column:ipFrom;comment:开始IP" json:"ipFrom"`                                                    // 开始IP
	IPTo                          string    `gorm:"column:ipTo;comment:结束IP" json:"ipTo"`                                                        // 结束IP
	IPFromLong                    uint32    `gorm:"column:ipFromLong;comment:开始IP整型" json:"ipFromLong"`                                          // 开始IP整型
	IPToLong                      uint32    `gorm:"column:ipToLong;comment:结束IP整型" json:"ipToLong"`                                              // 结束IP整型
	Version                       uint32    `gorm:"column:version;comment:版本" json:"version"`                                                    // 版本
	Reason                        string    `gorm:"column:reason;comment:加入说明" json:"reason"`                                                    // 加入说明
	EventLevel                    string    `gorm:"column:eventLevel;comment:事件级别" json:"eventLevel"`                                            // 事件级别
	State                         bool      `gorm:"column:state;default:1;comment:状态" json:"state"`                                              // 状态
	ExpiredAt                     time.Time `gorm:"column:expiredAt;comment:过期时间" json:"expiredAt"`                                              // 过期时间
	ServerID                      uint32    `gorm:"column:serverId;comment:有效范围服务ID" json:"serverId"`                                            // 有效范围服务ID
	NodeID                        uint32    `gorm:"column:nodeId;comment:有效范围节点ID" json:"nodeId"`                                                // 有效范围节点ID
	SourceNodeID                  uint32    `gorm:"column:sourceNodeId;comment:来源节点ID" json:"sourceNodeId"`                                      // 来源节点ID
	SourceServerID                uint32    `gorm:"column:sourceServerId;comment:来源服务ID" json:"sourceServerId"`                                  // 来源服务ID
	SourceHTTPFirewallPolicyID    uint32    `gorm:"column:sourceHTTPFirewallPolicyId;comment:来源策略ID" json:"sourceHTTPFirewallPolicyId"`          // 来源策略ID
	SourceHTTPFirewallRuleGroupID uint32    `gorm:"column:sourceHTTPFirewallRuleGroupId;comment:来源规则集分组ID" json:"sourceHTTPFirewallRuleGroupId"` // 来源规则集分组ID
	SourceHTTPFirewallRuleSetID   uint32    `gorm:"column:sourceHTTPFirewallRuleSetId;comment:来源规则集ID" json:"sourceHTTPFirewallRuleSetId"`       // 来源规则集ID
	SourceUserID                  uint32    `gorm:"column:sourceUserId;comment:用户ID" json:"sourceUserId"`                                        // 用户ID
	IsRead                        uint32    `gorm:"column:isRead;default:1;comment:是否已读" json:"isRead"`                                          // 是否已读
}

// TableName IPItem's table name
func (*IPItem) TableName() string {
	return TableNameIPItem
}

// AfterCreate run after create database record.
func (u *IPItem) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "item-")).Error
}

type IPItemList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*IPItem `json:"items"`
}
