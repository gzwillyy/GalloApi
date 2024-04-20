package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameServerHTTPFirewallDailyStat = "galloServerHTTPFirewallDailyStats"

// ServerHTTPFirewallDailyStat WAF统计
type ServerHTTPFirewallDailyStat struct {
	metav1.ObjectMeta       `json:"metadata,omitempty"`
	ServerID                uint32  `gorm:"column:serverId;comment:服务ID" json:"serverId"`                                  // 服务ID
	Day                     string `gorm:"column:day;comment:YYYYMMDD" json:"day"`                                        // YYYYMMDD
	HTTPFirewallRuleGroupID uint32  `gorm:"column:httpFirewallRuleGroupId;comment:WAF分组ID" json:"httpFirewallRuleGroupId"` // WAF分组ID
	Action                  string `gorm:"column:action;comment:采取的动作" json:"action"`                                     // 采取的动作
	Count                   uint32  `gorm:"column:count;comment:数量" json:"count"`                                          // 数量
}

// TableName ServerHTTPFirewallDailyStat's table name
func (*ServerHTTPFirewallDailyStat) TableName() string {
	return TableNameServerHTTPFirewallDailyStat
}

// AfterCreate run after create database record.
func (u *ServerHTTPFirewallDailyStat) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "stat-")).Error
}

type ServerHTTPFirewallDailyStatList struct {
	metav1.ListMeta `json:",inline"`
	Items []*ServerHTTPFirewallDailyStat `json:"items"`
}