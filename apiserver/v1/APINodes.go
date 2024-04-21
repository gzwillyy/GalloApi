package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloAPINode = "galloAPINodes"

// APINode API节点
type APINode struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`         // 是否启用
	ClusterID         uint32 `gorm:"column:clusterId;comment:专用集群ID" json:"clusterId"`       // 专用集群ID
	UniqueID          string `gorm:"column:uniqueId;comment:唯一ID" json:"uniqueId"`           // 唯一ID
	Secret            string `gorm:"column:secret;comment:密钥" json:"secret"`                 // 密钥
	Description       string `gorm:"column:description;comment:描述" json:"description"`       // 描述
	HTTP              string `gorm:"column:http;comment:监听的HTTP配置" json:"http"`              // 监听的HTTP配置
	HTTPS             string `gorm:"column:https;comment:监听的HTTPS配置" json:"https"`           // 监听的HTTPS配置
	RestIsOn          uint8  `gorm:"column:restIsOn;comment:是否开放REST" json:"restIsOn"`       // 是否开放REST
	RestHTTP          string `gorm:"column:restHTTP;comment:REST HTTP配置" json:"restHTTP"`    // REST HTTP配置
	RestHTTPS         string `gorm:"column:restHTTPS;comment:REST HTTPS配置" json:"restHTTPS"` // REST HTTPS配置
	AccessAddrs       string `gorm:"column:accessAddrs;comment:外部访问地址" json:"accessAddrs"`   // 外部访问地址
	Order             uint8  `gorm:"column:order;comment:排序" json:"order"`                   // 排序
	State             int    `gorm:"column:state;default:1;comment:状态" json:"state"`         // 状态
	AdminID           uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`            // 管理员ID
	Weight            uint32 `gorm:"column:weight;comment:权重" json:"weight"`                 // 权重
	Status            string `gorm:"column:status;comment:运行状态" json:"status"`               // 运行状态
	IsPrimary         uint8  `gorm:"column:isPrimary;comment:是否为主API节点" json:"isPrimary"`    // 是否为主API节点
}

// TableName APINode's table name
func (*APINode) TableName() string {
	return TableNameGalloAPINode
}

// AfterCreate run after create database record.
func (u *APINode) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "node-")).Error
}

type APINodeList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*APINode `json:"items"`
}
