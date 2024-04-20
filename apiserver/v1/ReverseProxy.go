package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameReverseProxy = "galloReverseProxies"

// ReverseProxy 反向代理配置
type ReverseProxy struct {
	metav1.ObjectMeta        `json:"metadata,omitempty"`
	AdminID                  uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                                          // 管理员ID
	UserID                   uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`                                             // 用户ID
	TemplateID               uint32 `gorm:"column:templateId;comment:模版ID" json:"templateId"`                                     // 模版ID
	IsOn                     bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                                       // 是否启用
	Scheduling               string `gorm:"column:scheduling;comment:调度算法 json" json:"scheduling"`                                // 调度算法 json
	PrimaryOrigins           string `gorm:"column:primaryOrigins;comment:主要源站 json" json:"primaryOrigins"`                        // 主要源站 json
	BackupOrigins            string `gorm:"column:backupOrigins;comment:备用源站 json" json:"backupOrigins"`                          // 备用源站 json
	StripPrefix              string `gorm:"column:stripPrefix;comment:去除URL前缀" json:"stripPrefix"`                                // 去除URL前缀
	RequestHostType          uint8  `gorm:"column:requestHostType;comment:请求Host类型" json:"requestHostType"`                       // 请求Host类型
	RequestHost              string `gorm:"column:requestHost;comment:请求Host" json:"requestHost"`                                 // 请求Host
	RequestHostExcludingPort uint32 `gorm:"column:requestHostExcludingPort;comment:移除请求Host中的域名" json:"requestHostExcludingPort"` // 移除请求Host中的域名
	RequestURI               string `gorm:"column:requestURI;comment:请求URI" json:"requestURI"`                                    // 请求URI
	AutoFlush                uint32 `gorm:"column:autoFlush;comment:是否自动刷新缓冲区" json:"autoFlush"`                                  // 是否自动刷新缓冲区
	AddHeaders               string `gorm:"column:addHeaders;comment:自动添加的Header列表 json" json:"addHeaders"`                       // 自动添加的Header列表 json
	State                    bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                                       // 状态
	ConnTimeout              string `gorm:"column:connTimeout;comment:连接超时时间 json" json:"connTimeout"`                            // 连接超时时间 json
	ReadTimeout              string `gorm:"column:readTimeout;comment:读取超时时间 json" json:"readTimeout"`                            // 读取超时时间 json
	IdleTimeout              string `gorm:"column:idleTimeout;comment:空闲超时时间 json" json:"idleTimeout"`                            // 空闲超时时间 json
	MaxConns                 uint32 `gorm:"column:maxConns;comment:最大并发连接数" json:"maxConns"`                                      // 最大并发连接数
	MaxIdleConns             uint32 `gorm:"column:maxIdleConns;comment:最大空闲连接数" json:"maxIdleConns"`                              // 最大空闲连接数
	ProxyProtocol            string `gorm:"column:proxyProtocol;comment:Proxy Protocol配置 json" json:"proxyProtocol"`              // Proxy Protocol配置 json
	FollowRedirects          uint32 `gorm:"column:followRedirects;comment:回源跟随" json:"followRedirects"`                           // 回源跟随
	Retry50X                 bool   `gorm:"column:retry50X;comment:启用50X重试" json:"retry50X"`                                      // 启用50X重试
	Retry40X                 bool   `gorm:"column:retry40X;comment:启用40X重试" json:"retry40X"`                                      // 启用40X重试
}

// TableName ReverseProxy's table name
func (*ReverseProxy) TableName() string {
	return TableNameReverseProxy
}

// AfterCreate run after create database record.
func (u *ReverseProxy) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "proxy-")).Error
}

// ReverseProxyList 返回列表
type ReverseProxyList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*ReverseProxy `json:"items"`
}
