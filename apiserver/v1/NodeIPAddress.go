package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeIPAddress = "galloNodeIPAddresses"

// NodeIPAddress 节点IP地址
type NodeIPAddress struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	NodeID            uint32 `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`                            // 节点ID
	ClusterIds        string `gorm:"column:clusterIds;comment:所属集群IDs" json:"clusterIds"`                 // 所属集群IDs
	Role              string `gorm:"column:role;default:node;comment:节点角色" json:"role"`                   // 节点角色
	GroupID           uint32 `gorm:"column:groupId;comment:所属分组ID" json:"groupId"`                        // 所属分组ID
	IP                string `gorm:"column:ip;comment:IP地址" json:"ip"`                                    // IP地址
	Description       string `gorm:"column:description;comment:描述" json:"description"`                    // 描述
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                      // 状态
	Order             uint32 `gorm:"column:order;comment:排序" json:"order"`                                // 排序
	CanAccess         uint32 `gorm:"column:canAccess;default:1;comment:是否可以访问" json:"canAccess"`          // 是否可以访问
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                      // 是否启用
	IsUp              uint32 `gorm:"column:isUp;default:1;comment:是否上线" json:"isUp"`                      // 是否上线
	IsHealthy         uint32 `gorm:"column:isHealthy;default:1;comment:是否健康" json:"isHealthy"`            // 是否健康
	Thresholds        string `gorm:"column:thresholds;comment:上线阈值" json:"thresholds"`                    // 上线阈值
	Connectivity      string `gorm:"column:connectivity;comment:连通性状态" json:"connectivity"`               // 连通性状态
	BackupIP          string `gorm:"column:backupIP;comment:备用IP" json:"backupIP"`                        // 备用IP
	BackupThresholdID uint32 `gorm:"column:backupThresholdId;comment:触发备用IP的阈值" json:"backupThresholdId"` // 触发备用IP的阈值
	CountUp           uint32 `gorm:"column:countUp;comment:UP状态次数" json:"countUp"`                        // UP状态次数
	CountDown         uint32 `gorm:"column:countDown;comment:DOWN状态次数" json:"countDown"`                  // DOWN状态次数
}

// TableName NodeIPAddress's table name
func (*NodeIPAddress) TableName() string {
	return TableNameNodeIPAddress
}

// AfterCreate run after create database record.
func (u *NodeIPAddress) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "address-")).Error
}

type NodeIPAddressList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NodeIPAddress `json:"items"`
}

var NodeIPAddressTableZeroFields = []string{"name", "description", "state", "isOn", "thresholds", "connectivity", "backupIP"}

