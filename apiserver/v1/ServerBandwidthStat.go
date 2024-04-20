package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameServerBandwidthStat = "galloServerBandwidthStats"

// ServerBandwidthStat 服务峰值带宽统计
type ServerBandwidthStat struct {
	metav1.ObjectMeta   `json:"metadata,omitempty"`
	UserID              uint32  `gorm:"column:userId;comment:用户ID" json:"userId"`                             // 用户ID
	ServerID            uint32  `gorm:"column:serverId;comment:服务ID" json:"serverId"`                         // 服务ID
	RegionID            uint32  `gorm:"column:regionId;comment:区域ID" json:"regionId"`                         // 区域ID
	UserPlanID          uint32  `gorm:"column:userPlanId;comment:用户套餐ID" json:"userPlanId"`                   // 用户套餐ID
	Day                 string `gorm:"column:day;comment:日期YYYYMMDD" json:"day"`                             // 日期YYYYMMDD
	TimeAt              string `gorm:"column:timeAt;comment:时间点HHMM" json:"timeAt"`                          // 时间点HHMM
	Bytes               uint32  `gorm:"column:bytes;comment:带宽字节" json:"bytes"`                               // 带宽字节
	AvgBytes            uint32  `gorm:"column:avgBytes;comment:平均流量" json:"avgBytes"`                         // 平均流量
	CachedBytes         uint32  `gorm:"column:cachedBytes;comment:缓存的流量" json:"cachedBytes"`                  // 缓存的流量
	AttackBytes         uint32  `gorm:"column:attackBytes;comment:攻击流量" json:"attackBytes"`                   // 攻击流量
	CountRequests       uint32  `gorm:"column:countRequests;comment:请求数" json:"countRequests"`                // 请求数
	CountCachedRequests uint32  `gorm:"column:countCachedRequests;comment:缓存的请求数" json:"countCachedRequests"` // 缓存的请求数
	CountAttackRequests uint32  `gorm:"column:countAttackRequests;comment:攻击请求数" json:"countAttackRequests"`  // 攻击请求数
	TotalBytes          uint32  `gorm:"column:totalBytes;comment:总流量" json:"totalBytes"`                      // 总流量
}

// TableName ServerBandwidthStat's table name
func (*ServerBandwidthStat) TableName() string {
	return TableNameServerBandwidthStat
}

// AfterCreate run after create database record.
func (u *ServerBandwidthStat) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "stat-")).Error
}

type ServerBandwidthStatList struct {
	metav1.ListMeta `json:",inline"`
	Items []*ServerBandwidthStat `json:"items"`
}