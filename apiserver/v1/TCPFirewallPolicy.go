package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameTCPFirewallPolicy = "galloTCPFirewallPolicies"

// TCPFirewallPolicy TCP防火墙
type TCPFirewallPolicy struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`      // 管理员ID
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`         // 用户ID
	TemplateID        uint32 `gorm:"column:templateId;comment:模版ID" json:"templateId"` // 模版ID
}

// TableName TCPFirewallPolicy's table name
func (*TCPFirewallPolicy) TableName() string {
	return TableNameTCPFirewallPolicy
}

// AfterCreate run after create database record.
func (u *TCPFirewallPolicy) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "firewall-")).Error
}

type TCPFirewallPolicyList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*TCPFirewallPolicy `json:"items"`
}
