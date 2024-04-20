package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeTrafficMinlyStat = "galloNodeTrafficMinlyStats"

// NodeTrafficMinlyStat 总的流量统计（按分钟）
type NodeTrafficMinlyStat struct {
	metav1.ObjectMeta   `json:"metadata,omitempty"`
	Role                string `gorm:"column:role;default:node;comment:节点角色" json:"role"`                    // 节点角色
	ClusterID           uint32  `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`                       // 集群ID
	NodeID              uint32  `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`                             // 节点ID
	Minute              string `gorm:"column:minute;comment:YYYYMMDDHHSS" json:"minute"`                     // YYYYMMDDHHSS
	Bytes               uint32  `gorm:"column:bytes;comment:流量字节" json:"bytes"`                               // 流量字节
	CachedBytes         uint32  `gorm:"column:cachedBytes;comment:缓存流量" json:"cachedBytes"`                   // 缓存流量
	CountRequests       uint32  `gorm:"column:countRequests;comment:请求数" json:"countRequests"`                // 请求数
	CountCachedRequests uint32  `gorm:"column:countCachedRequests;comment:缓存的请求数" json:"countCachedRequests"` // 缓存的请求数
	CountAttackRequests uint32  `gorm:"column:countAttackRequests;comment:攻击请求数" json:"countAttackRequests"`  // 攻击请求数
	AttackBytes         uint32  `gorm:"column:attackBytes;comment:攻击流量" json:"attackBytes"`                   // 攻击流量
}

// TableName NodeTrafficMinlyStat's table name
func (*NodeTrafficMinlyStat) TableName() string {
	return TableNameNodeTrafficMinlyStat
}

// AfterCreate run after create database record.
func (u *NodeTrafficMinlyStat) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "stat-")).Error
}

type NodeTrafficMinlyStatList struct {
	metav1.ListMeta `json:",inline"`
	Items []*NodeTrafficMinlyStat `json:"items"`
}