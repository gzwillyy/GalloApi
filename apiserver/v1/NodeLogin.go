package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeLogin = "galloNodeLogins"

// NodeLogin 节点登录信息
type NodeLogin struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	NodeID            uint32 `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`        // 节点ID
	Role              string `gorm:"column:role;default:node;comment:角色" json:"role"` // 角色
	Type              string `gorm:"column:type;comment:类型：ssh,agent" json:"type"`    // 类型：ssh,agent
	Params            string `gorm:"column:params;comment:配置参数" json:"params"`        // 配置参数
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`  // 状态
}

// TableName NodeLogin's table name
func (*NodeLogin) TableName() string {
	return TableNameNodeLogin
}

// AfterCreate run after create database record.
func (u *NodeLogin) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "login-")).Error
}

type NodeLoginList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NodeLogin `json:"items"`
}
