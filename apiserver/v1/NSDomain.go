package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNSDomain = "galloNSDomains"

// NSDomain DNS域名
type NSDomain struct {
	metav1.ObjectMeta  `json:"metadata,omitempty"`
	ClusterID          uint32 `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`                       // 集群ID
	UserID             uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`                             // 用户ID
	IsOn               bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                       // 是否启用
	GroupIds           string `gorm:"column:groupIds;comment:分组ID" json:"groupIds"`                         // 分组ID
	Tsig               string `gorm:"column:tsig;comment:TSIG配置" json:"tsig"`                               // TSIG配置
	VerifyTXT          string `gorm:"column:verifyTXT;comment:验证用的TXT" json:"verifyTXT"`                    // 验证用的TXT
	VerifyExpiresAt    uint32 `gorm:"column:verifyExpiresAt;comment:验证TXT过期时间" json:"verifyExpiresAt"`      // 验证TXT过期时间
	RecordsHealthCheck string `gorm:"column:recordsHealthCheck;comment:记录健康检查设置" json:"recordsHealthCheck"` // 记录健康检查设置
	Version            uint32 `gorm:"column:version;comment:版本号" json:"version"`                            // 版本号
	Status             string `gorm:"column:status;default:none;comment:状态：none|verified" json:"status"`    // 状态：none|verified
	State              bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                       // 状态
}

// TableName NSDomain's table name
func (*NSDomain) TableName() string {
	return TableNameNSDomain
}

// AfterCreate run after create database record.
func (u *NSDomain) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "domain-")).Error
}

type NSDomainList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NSDomain `json:"items"`
}

var NSDomainTableZeroFields = []string{"name", "groupIds", "verifyTXT", "recordsHealthCheck", "status"}

