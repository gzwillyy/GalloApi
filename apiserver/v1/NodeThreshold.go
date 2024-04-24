package v1

import (
	"time"

	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeThreshold = "galloNodeThresholds"

// NodeThreshold 节点监控阈值设置
type NodeThreshold struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Role              string    `gorm:"column:role;comment:节点角色" json:"role"`                           // 节点角色
	ClusterID         uint32    `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`                 // 集群ID
	NodeID            uint32    `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`                       // 节点ID
	IsOn              bool      `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                 // 是否启用
	Item              string    `gorm:"column:item;comment:监控项" json:"item"`                            // 监控项
	Param             string    `gorm:"column:param;comment:参数" json:"param"`                           // 参数
	Operator          string    `gorm:"column:operator;comment:操作符" json:"operator"`                    // 操作符
	Value             string    `gorm:"column:value;comment:对比值" json:"value"`                          // 对比值
	Message           string    `gorm:"column:message;comment:消息内容" json:"message"`                     // 消息内容
	NotifyDuration    uint32    `gorm:"column:notifyDuration;comment:通知间隔（单位分钟）" json:"notifyDuration"` // 通知间隔（单位分钟）
	NotifiedAt        time.Time `gorm:"column:notifiedAt;comment:上次通知时间" json:"notifiedAt"`             // 上次通知时间
	Duration          uint32    `gorm:"column:duration;comment:时间段" json:"duration"`                    // 时间段
	DurationUnit      string    `gorm:"column:durationUnit;comment:时间段单位" json:"durationUnit"`          // 时间段单位
	SumMethod         string    `gorm:"column:sumMethod;comment:聚合方法" json:"sumMethod"`                 // 聚合方法
	Order             uint32    `gorm:"column:order;comment:排序" json:"order"`                           // 排序
	State             bool      `gorm:"column:state;default:1;comment:状态" json:"state"`                 // 状态
}

// TableName NodeThreshold's table name
func (*NodeThreshold) TableName() string {
	return TableNameNodeThreshold
}

// AfterCreate run after create database record.
func (u *NodeThreshold) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "hold-")).Error
}

type NodeThresholdList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NodeThreshold `json:"items"`
}

var NodeThresholdTableZeroFields = []string{"name", "role", "item", "param", "operator", "value", "message", "durationUnit", "sumMethod", "state"}

