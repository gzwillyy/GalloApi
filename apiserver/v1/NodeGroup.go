package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNodeGroup = "galloNodeGroups"

// NodeGroup 节点分组
type NodeGroup struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	ClusterID         uint32 `gorm:"column:clusterId;comment:集群ID" json:"clusterId"` // 集群ID
	Order             uint32 `gorm:"column:order;comment:排序" json:"order"`           // 排序
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
}

// TableName NodeGroup's table name
func (*NodeGroup) TableName() string {
	return TableNameNodeGroup
}

// AfterCreate run after create database record.
func (u *NodeGroup) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "group-")).Error
}

// NodeGroupList 返回列表
type NodeGroupList struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	metav1.ListMeta `json:",inline"`

	// List of secrets
	Items []*NodeGroup `json:"items"`
}

// CreateNodeGroupRequest 创建分组
type CreateNodeGroupRequest struct {
	ClusterInstanceId string `json:"clusterInstanceId,omitempty"`
	Name              string `json:"name,omitempty"`
}

// UpdateNodeGroupRequest 修改分组
type UpdateNodeGroupRequest struct {
	InstanceID string `json:"instanceID"`
	Name       string `json:"name"`
}

// DeleteNodeGroupRequest 删除分组
type DeleteNodeGroupRequest struct {
	InstanceID string `json:"instanceID"`
}
