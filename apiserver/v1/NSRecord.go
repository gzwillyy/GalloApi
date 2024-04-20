package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNSRecord = "galloNSRecords"

// NSRecord DNS记录
type NSRecord struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	DomainID          uint32 `gorm:"column:domainId;comment:域名ID" json:"domainId"`         // 域名ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`       // 是否启用
	Description       string `gorm:"column:description;comment:备注" json:"description"`     // 备注
	Type              string `gorm:"column:type;comment:类型" json:"type"`                   // 类型
	Value             string `gorm:"column:value;comment:值" json:"value"`                  // 值
	MxPriority        uint32 `gorm:"column:mxPriority;comment:MX优先级" json:"mxPriority"`    // MX优先级
	SrvPriority       uint32 `gorm:"column:srvPriority;comment:SRV优先级" json:"srvPriority"` // SRV优先级
	SrvWeight         uint32 `gorm:"column:srvWeight;comment:SRV权重" json:"srvWeight"`      // SRV权重
	SrvPort           uint32 `gorm:"column:srvPort;comment:SRV端口" json:"srvPort"`          // SRV端口
	CaaFlag           uint32 `gorm:"column:caaFlag;comment:CAA Flag" json:"caaFlag"`       // CAA Flag
	CaaTag            string `gorm:"column:caaTag;comment:CAA TAG" json:"caaTag"`          // CAA TAG
	TTL               uint32 `gorm:"column:ttl;comment:TTL（秒）" json:"ttl"`                 // TTL（秒）
	Weight            uint32 `gorm:"column:weight;comment:权重" json:"weight"`               // 权重
	RouteIds          string `gorm:"column:routeIds;comment:线路" json:"routeIds"`           // 线路
	HealthCheck       string `gorm:"column:healthCheck;comment:健康检查配置" json:"healthCheck"` // 健康检查配置
	CountUp           uint32 `gorm:"column:countUp;comment:连续上线次数" json:"countUp"`         // 连续上线次数
	CountDown         uint32 `gorm:"column:countDown;comment:连续离线次数" json:"countDown"`     // 连续离线次数
	IsUp              uint32 `gorm:"column:isUp;default:1;comment:是否在线" json:"isUp"`       // 是否在线
	Version           uint32 `gorm:"column:version;comment:版本号" json:"version"`            // 版本号
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`       // 状态
}

// TableName NSRecord's table name
func (*NSRecord) TableName() string {
	return TableNameNSRecord
}

// AfterCreate run after create database record.
func (u *NSRecord) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "record-")).Error
}

type NSRecordList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NSRecord `json:"items"`
}
