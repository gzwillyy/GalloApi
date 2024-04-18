package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloServerGroup = "galloServerGroups"

// ServerGroup 服务分组
type ServerGroup struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint64 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                    // 管理员ID
	UserID            uint64 `gorm:"column:userId;comment:用户ID" json:"userId"`                       // 用户ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                 // 是否启用
	Order             uint64 `gorm:"column:order;comment:排序" json:"order"`                           // 排序
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                 // 状态
	HTTPReverseProxy  string `gorm:"column:httpReverseProxy;comment:反向代理设置" json:"httpReverseProxy"` // 反向代理设置
	TcpReverseProxy   string `gorm:"column:tcpReverseProxy;comment:TCP反向代理" json:"tcpReverseProxy"`  // TCP反向代理
	UdpReverseProxy   string `gorm:"column:udpReverseProxy;comment:UDP反向代理" json:"udpReverseProxy"`  // UDP反向代理
	WebID             uint64 `gorm:"column:webId;comment:Web配置ID" json:"webId"`                      // Web配置ID
}

// TableName ServerGroup's table name
func (*ServerGroup) TableName() string {
	return TableNameGalloServerGroup
}

func (u *ServerGroup) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "group-")).Error
}

// ServerGroupList 返回列表
type ServerGroupList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*ServerGroup `json:"items"`
}

// CreateServerGroupRequest 创建分组
type CreateServerGroupRequest struct {
	ClusterInstanceId string `json:"clusterInstanceId,omitempty"`
	Name              string `json:"name,omitempty"`
}

// UpdateServerGroupRequest 修改分组
type UpdateServerGroupRequest struct {
	InstanceID string `json:"instanceID"`
	Name       string `json:"name"`
	RootJSON   string `json:"root"`
}

// DeleteServerGroupRequest 删除分组
type DeleteServerGroupRequest struct {
	InstanceID string `json:"instanceID"`
}
