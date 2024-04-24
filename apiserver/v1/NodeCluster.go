package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloNodeCluster = "galloNodeClusters"

// NodeCluster 节点集群
type NodeCluster struct {
	metav1.ObjectMeta    `json:"metadata,omitempty"`
	AdminID              uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                                  // 管理员ID
	UserID               uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`                                     // 用户ID
	IsOn                 bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                               // 是否启用
	UseAllAPINodes       bool   `gorm:"column:useAllAPINodes;default:1;comment:是否使用所有API节点" json:"useAllAPINodes"`    // 是否使用所有API节点
	APINodes             string `gorm:"column:apiNodes;comment:使用的API节点" json:"apiNodes"`                             // 使用的API节点
	InstallDir           string `gorm:"column:installDir;comment:安装目录" json:"installDir"`                             // 安装目录
	Order                uint32 `gorm:"column:order;comment:排序" json:"order"`                                         // 排序
	GrantID              uint32 `gorm:"column:grantId;comment:默认认证方式" json:"grantId"`                                 // 默认认证方式
	SSHParams            string `gorm:"column:sshParams;comment:SSH默认参数" json:"sshParams"`                            // SSH默认参数
	State                bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                               // 状态
	AutoRegister         bool   `gorm:"column:autoRegister;default:1;comment:是否开启自动注册" json:"autoRegister"`           // 是否开启自动注册
	UniqueID             string `gorm:"column:uniqueId;comment:唯一ID" json:"uniqueId"`                                 // 唯一ID
	Secret               string `gorm:"column:secret;comment:密钥" json:"secret"`                                       // 密钥
	HealthCheck          string `gorm:"column:healthCheck;comment:健康检查" json:"healthCheck"`                           // 健康检查
	DNSName              string `gorm:"column:dnsName;comment:DNS名称" json:"dnsName"`                                  // DNS名称
	DNSDomainID          uint32 `gorm:"column:dnsDomainId;comment:域名ID" json:"dnsDomainId"`                           // 域名ID
	DNS                  string `gorm:"column:dns;comment:DNS配置" json:"dns"`                                          // DNS配置
	Toa                  string `gorm:"column:toa;comment:TOA配置" json:"toa"`                                          // TOA配置
	CachePolicyID        uint32 `gorm:"column:cachePolicyId;comment:缓存策略ID" json:"cachePolicyId"`                     // 缓存策略ID
	HTTPFirewallPolicyID uint32 `gorm:"column:httpFirewallPolicyId;comment:WAF策略ID" json:"httpFirewallPolicyId"`      // WAF策略ID
	AccessLog            string `gorm:"column:accessLog;comment:访问日志设置" json:"accessLog"`                             // 访问日志设置
	SystemServices       string `gorm:"column:systemServices;comment:系统服务设置" json:"systemServices"`                   // 系统服务设置
	TimeZone             string `gorm:"column:timeZone;comment:时区" json:"timeZone"`                                   // 时区
	NodeMaxThreads       uint32 `gorm:"column:nodeMaxThreads;comment:节点最大线程数" json:"nodeMaxThreads"`                  // 节点最大线程数
	DdosProtection       string `gorm:"column:ddosProtection;comment:DDoS防护设置" json:"ddosProtection"`                 // DDoS防护设置
	AutoOpenPorts        bool   `gorm:"column:autoOpenPorts;default:1;comment:是否自动尝试开放端口" json:"autoOpenPorts"`       // 是否自动尝试开放端口
	IsPinned             bool   `gorm:"column:isPinned;comment:是否置顶" json:"isPinned"`                                 // 是否置顶
	Webp                 string `gorm:"column:webp;comment:WebP设置" json:"webp"`                                       // WebP设置
	Uam                  string `gorm:"column:uam;comment:UAM设置" json:"uam"`                                          // UAM设置
	Clock                string `gorm:"column:clock;comment:时钟配置" json:"clock"`                                       // 时钟配置
	GlobalServerConfig   string `gorm:"column:globalServerConfig;comment:全局服务配置" json:"globalServerConfig"`           // 全局服务配置
	AutoRemoteStart      bool   `gorm:"column:autoRemoteStart;default:1;comment:自动远程启动" json:"autoRemoteStart"`       // 自动远程启动
	AutoInstallNftables  bool   `gorm:"column:autoInstallNftables;comment:自动安装nftables" json:"autoInstallNftables"`   // 自动安装nftables
	IsAD                 bool   `gorm:"column:isAD;comment:是否为高防集群" json:"isAD"`                                      // 是否为高防集群
	HTTPPages            string `gorm:"column:httpPages;comment:自定义页面设置" json:"httpPages"`                            // 自定义页面设置
	Cc                   string `gorm:"column:cc;comment:CC设置" json:"cc"`                                             // CC设置
	Http3                string `gorm:"column:http3;comment:HTTP3设置" json:"http3"`                                    // HTTP3设置
	AutoSystemTuning     bool   `gorm:"column:autoSystemTuning;default:1;comment:是否自动调整系统参数" json:"autoSystemTuning"` // 是否自动调整系统参数
	NetworkSecurity      string `gorm:"column:networkSecurity;comment:网络安全策略" json:"networkSecurity"`                 // 网络安全策略
}

// TableName NodeCluster's table name
func (*NodeCluster) TableName() string {
	return TableNameGalloNodeCluster
}

// AfterCreate run after create database record.
func (u *NodeCluster) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "cluster-")).Error
}

// NodeClusterRef 列表数据
type NodeClusterRef struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	NodeCount         uint32 `json:"nodeCount"`
	OnlineNodeCount   uint32 `json:"onlineNodeCount"`
	ServerCount       uint32 `json:"serverCount"`
	DNSName           string `json:"dnsName"`
}

// NodeClusterRefList 返回列表
type NodeClusterRefList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NodeClusterRef `json:"items"`
}

// NodeClusterList 资源列表
type NodeClusterList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NodeCluster `json:"items"`
}

var NodeClusterTableZeroFields = []string{"name", "useAllAPINodes", "apiNodes", "installDir", "sshParams", "state", "autoRegister", "uniqueId", "healthCheck", "dnsName", "toa", "systemServices", "timeZone", "ddosProtection", "autoOpenPorts", "isPinned", "webp", "uam", "clock", "globalServerConfig", "autoRemoteStart", "autoInstallNftables", "httpPages", "cc", "http3", "autoSystemTuning", "networkSecurity"}

