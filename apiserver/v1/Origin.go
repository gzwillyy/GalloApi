package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameOrigin = "galloOrigins"

// Origin 源站
type Origin struct {
	metav1.ObjectMeta  `json:"metadata,omitempty"`
	AdminID            uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                            // 管理员ID
	UserID             uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`                               // 用户ID
	IsOn               bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                         // 是否启用
	Version            uint32 `gorm:"column:version;comment:版本" json:"version"`                               // 版本
	Addr               string `gorm:"column:addr;comment:地址" json:"addr"`                                     // 地址
	Oss                string `gorm:"column:oss;comment:OSS配置" json:"oss"`                                    // OSS配置
	Description        string `gorm:"column:description;comment:描述" json:"description"`                       // 描述
	Code               string `gorm:"column:code;comment:代号" json:"code"`                                     // 代号
	Weight             uint32 `gorm:"column:weight;comment:权重" json:"weight"`                                 // 权重
	ConnTimeout        string `gorm:"column:connTimeout;comment:连接超时" json:"connTimeout"`                     // 连接超时
	ReadTimeout        string `gorm:"column:readTimeout;comment:读超时" json:"readTimeout"`                      // 读超时
	IdleTimeout        string `gorm:"column:idleTimeout;comment:空闲连接超时" json:"idleTimeout"`                   // 空闲连接超时
	MaxFails           uint32 `gorm:"column:maxFails;comment:最多失败次数" json:"maxFails"`                         // 最多失败次数
	MaxConns           uint32 `gorm:"column:maxConns;comment:最大并发连接数" json:"maxConns"`                        // 最大并发连接数
	MaxIdleConns       uint32 `gorm:"column:maxIdleConns;comment:最多空闲连接数" json:"maxIdleConns"`                // 最多空闲连接数
	HTTPRequestURI     string `gorm:"column:httpRequestURI;comment:转发后的请求URI" json:"httpRequestURI"`          // 转发后的请求URI
	HTTPRequestHeader  string `gorm:"column:httpRequestHeader;comment:请求Header配置" json:"httpRequestHeader"`   // 请求Header配置
	HTTPResponseHeader string `gorm:"column:httpResponseHeader;comment:响应Header配置" json:"httpResponseHeader"` // 响应Header配置
	Host               string `gorm:"column:host;comment:自定义主机名" json:"host"`                                 // 自定义主机名
	HealthCheck        string `gorm:"column:healthCheck;comment:健康检查设置" json:"healthCheck"`                   // 健康检查设置
	Cert               string `gorm:"column:cert;comment:证书设置" json:"cert"`                                   // 证书设置
	Ftp                string `gorm:"column:ftp;comment:FTP相关设置" json:"ftp"`                                  // FTP相关设置
	Domains            string `gorm:"column:domains;comment:所属域名" json:"domains"`                             // 所属域名
	FollowPort         uint32 `gorm:"column:followPort;comment:端口跟随" json:"followPort"`                       // 端口跟随
	State              bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                         // 状态
	Http2Enabled       uint32 `gorm:"column:http2Enabled;comment:是否支持HTTP/2" json:"http2Enabled"`             // 是否支持HTTP/2
}

// TableName Origin's table name
func (*Origin) TableName() string {
	return TableNameOrigin
}

// AfterCreate run after create database record.
func (u *Origin) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "origin-")).Error
}

type OriginList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*Origin `json:"items"`
}
