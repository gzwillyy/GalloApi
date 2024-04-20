package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNSNode = "galloNSNodes"

// NSNode 域名服务器节点
type NSNode struct {
	metav1.ObjectMeta  `json:"metadata,omitempty"`
	AdminID            uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                                // 管理员ID
	ClusterID          uint32 `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`                             // 集群ID
	IsOn               bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                             // 是否启用
	Status             string `gorm:"column:status;comment:运行状态" json:"status"`                                   // 运行状态
	UniqueID           string `gorm:"column:uniqueId;comment:节点ID" json:"uniqueId"`                               // 节点ID
	Secret             string `gorm:"column:secret;comment:密钥" json:"secret"`                                     // 密钥
	IsUp               uint32 `gorm:"column:isUp;default:1;comment:是否运行" json:"isUp"`                             // 是否运行
	IsInstalled        uint32 `gorm:"column:isInstalled;comment:是否已安装" json:"isInstalled"`                        // 是否已安装
	InstallStatus      string `gorm:"column:installStatus;comment:安装状态" json:"installStatus"`                     // 安装状态
	InstallDir         string `gorm:"column:installDir;comment:安装目录" json:"installDir"`                           // 安装目录
	State              bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                             // 状态
	IsActive           uint32 `gorm:"column:isActive;default:1;comment:是否活跃" json:"isActive"`                     // 是否活跃
	StatusIsNotified   uint32 `gorm:"column:statusIsNotified;default:1;comment:活跃状态已经通知" json:"statusIsNotified"` // 活跃状态已经通知
	InactiveNotifiedAt uint32 `gorm:"column:inactiveNotifiedAt;comment:离线通知时间" json:"inactiveNotifiedAt"`         // 离线通知时间
	ConnectedAPINodes  string `gorm:"column:connectedAPINodes;comment:当前连接的API节点" json:"connectedAPINodes"`       // 当前连接的API节点
	DdosProtection     string `gorm:"column:ddosProtection;comment:DDoS防护设置" json:"ddosProtection"`               // DDoS防护设置
	APINodeAddrs       string `gorm:"column:apiNodeAddrs;comment:API节点地址" json:"apiNodeAddrs"`                    // API节点地址
}

// TableName NSNode's table name
func (*NSNode) TableName() string {
	return TableNameNSNode
}

// AfterCreate run after create database record.
func (u *NSNode) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "node-")).Error
}

type NSNodeList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NSNode `json:"items"`
}
