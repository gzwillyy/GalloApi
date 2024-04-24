package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameServerDomainHourlyStat = "galloServerDomainHourlyStats"

// ServerDomainHourlyStat 服务域名统计
type ServerDomainHourlyStat struct {
	metav1.ObjectMeta   `json:"metadata,omitempty"`
	ClusterID           uint32  `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`                      // 集群ID
	NodeID              uint32  `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`                            // 节点ID
	ServerID            uint32  `gorm:"column:serverId;comment:服务ID" json:"serverId"`                        // 服务ID
	Domain              string `gorm:"column:domain;comment:域名" json:"domain"`                              // 域名
	Hour                string `gorm:"column:hour;comment:YYYYMMDDHH" json:"hour"`                          // YYYYMMDDHH
	Bytes               uint32  `gorm:"column:bytes;comment:流量" json:"bytes"`                                // 流量
	CachedBytes         uint32  `gorm:"column:cachedBytes;comment:缓存流量" json:"cachedBytes"`                  // 缓存流量
	CountRequests       uint32  `gorm:"column:countRequests;comment:请求数" json:"countRequests"`               // 请求数
	CountCachedRequests uint32  `gorm:"column:countCachedRequests;comment:缓存请求" json:"countCachedRequests"`  // 缓存请求
	CountAttackRequests uint32  `gorm:"column:countAttackRequests;comment:攻击请求数" json:"countAttackRequests"` // 攻击请求数
	AttackBytes         uint32  `gorm:"column:attackBytes;comment:攻击流量" json:"attackBytes"`                  // 攻击流量
}

// TableName ServerDomainHourlyStat's table name
func (*ServerDomainHourlyStat) TableName() string {
	return TableNameServerDomainHourlyStat
}

// AfterCreate run after create database record.
func (u *ServerDomainHourlyStat) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "stat-")).Error
}

type ServerDomainHourlyStatList struct {
	metav1.ListMeta `json:",inline"`
	Items []*ServerDomainHourlyStat `json:"items"`
}
var ServerDomainHourlyStatTableZeroFields = []string{"name", "hour"}

