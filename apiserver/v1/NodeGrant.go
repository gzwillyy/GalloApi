package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeGrant = "galloNodeGrants"

// NodeGrant 节点授权登录信息
type NodeGrant struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`      // 管理员ID
	Method            string `gorm:"column:method;comment:登录方式" json:"method"`         // 登录方式
	Username          string `gorm:"column:username;comment:用户名" json:"username"`      // 用户名
	Password          string `gorm:"column:password;comment:密码" json:"password"`       // 密码
	Su                uint32 `gorm:"column:su;comment:是否需要su" json:"su"`               // 是否需要su
	PrivateKey        string `gorm:"column:privateKey;comment:私钥" json:"privateKey"`   // 私钥
	Passphrase        string `gorm:"column:passphrase;comment:私钥密码" json:"passphrase"` // 私钥密码
	Description       string `gorm:"column:description;comment:备注" json:"description"` // 备注
	NodeID            uint32 `gorm:"column:nodeId;comment:专有节点" json:"nodeId"`         // 专有节点
	Role              string `gorm:"column:role;default:node;comment:角色" json:"role"`  // 角色
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`   // 状态
}

// TableName NodeGrant's table name
func (*NodeGrant) TableName() string {
	return TableNameNodeGrant
}

// AfterCreate run after create database record.
func (u *NodeGrant) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "grant-")).Error
}

type NodeGrantList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NodeGrant `json:"items"`
}
