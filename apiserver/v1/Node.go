package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNode = "galloNodes"

// Node 节点
type Node struct {
	metav1.ObjectMeta      `json:"metadata,omitempty"`
	AdminID                uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                                // 管理员ID
	UserID                 uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`                                   // 用户ID
	Level                  uint32 `gorm:"column:level;default:1;comment:级别" json:"level"`                             // 级别
	LnAddrs                string `gorm:"column:lnAddrs;comment:Ln级别访问地址" json:"lnAddrs"`                             // Ln级别访问地址
	IsOn                   bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                             // 是否启用
	IsUp                   uint32 `gorm:"column:isUp;default:1;comment:是否在线" json:"isUp"`                             // 是否在线
	CountUp                uint32 `gorm:"column:countUp;comment:连续在线次数" json:"countUp"`                               // 连续在线次数
	CountDown              uint32 `gorm:"column:countDown;comment:连续下线次数" json:"countDown"`                           // 连续下线次数
	IsActive               uint32 `gorm:"column:isActive;default:1;comment:是否活跃" json:"isActive"`                     // 是否活跃
	InactiveNotifiedAt     uint32 `gorm:"column:inactiveNotifiedAt;comment:离线通知时间" json:"inactiveNotifiedAt"`         // 离线通知时间
	UniqueID               string `gorm:"column:uniqueId;comment:节点ID" json:"uniqueId"`                               // 节点ID
	Secret                 string `gorm:"column:secret;comment:密钥" json:"secret"`                                     // 密钥
	Code                   string `gorm:"column:code;comment:代号" json:"code"`                                         // 代号
	ClusterID              uint32 `gorm:"column:clusterId;comment:主集群ID" json:"clusterId"`                            // 主集群ID
	SecondaryClusterIds    string `gorm:"column:secondaryClusterIds;comment:从集群ID" json:"secondaryClusterIds"`        // 从集群ID
	RegionID               uint32 `gorm:"column:regionId;comment:区域ID" json:"regionId"`                               // 区域ID
	GroupID                uint32 `gorm:"column:groupId;comment:分组ID" json:"groupId"`                                 // 分组ID
	Status                 string `gorm:"column:status;comment:最新的状态" json:"status"`                                  // 最新的状态
	Version                uint32 `gorm:"column:version;comment:当前版本号" json:"version"`                                // 当前版本号
	LatestVersion          uint32 `gorm:"column:latestVersion;comment:最后版本号" json:"latestVersion"`                    // 最后版本号
	InstallDir             string `gorm:"column:installDir;comment:安装目录" json:"installDir"`                           // 安装目录
	IsInstalled            uint32 `gorm:"column:isInstalled;comment:是否已安装" json:"isInstalled"`                        // 是否已安装
	InstallStatus          string `gorm:"column:installStatus;comment:安装状态" json:"installStatus"`                     // 安装状态
	State                  bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                             // 状态
	ConnectedAPINodes      string `gorm:"column:connectedAPINodes;comment:当前连接的API节点" json:"connectedAPINodes"`       // 当前连接的API节点
	MaxCPU                 uint32 `gorm:"column:maxCPU;comment:可以使用的最多CPU" json:"maxCPU"`                             // 可以使用的最多CPU
	MaxThreads             uint32 `gorm:"column:maxThreads;comment:最大线程数" json:"maxThreads"`                          // 最大线程数
	DdosProtection         string `gorm:"column:ddosProtection;comment:DDOS配置" json:"ddosProtection"`                 // DDOS配置
	DNSRoutes              string `gorm:"column:dnsRoutes;comment:DNS线路设置" json:"dnsRoutes"`                          // DNS线路设置
	MaxCacheDiskCapacity   string `gorm:"column:maxCacheDiskCapacity;comment:硬盘缓存容量" json:"maxCacheDiskCapacity"`     // 硬盘缓存容量
	MaxCacheMemoryCapacity string `gorm:"column:maxCacheMemoryCapacity;comment:内存缓存容量" json:"maxCacheMemoryCapacity"` // 内存缓存容量
	CacheDiskDir           string `gorm:"column:cacheDiskDir;comment:主缓存目录" json:"cacheDiskDir"`                      // 主缓存目录
	CacheDiskSubDirs       string `gorm:"column:cacheDiskSubDirs;comment:其他缓存目录" json:"cacheDiskSubDirs"`             // 其他缓存目录
	DNSResolver            string `gorm:"column:dnsResolver;comment:DNS解析器" json:"dnsResolver"`                       // DNS解析器
	EnableIPLists          uint32 `gorm:"column:enableIPLists;default:1;comment:启用IP名单" json:"enableIPLists"`         // 启用IP名单
	APINodeAddrs           string `gorm:"column:apiNodeAddrs;comment:API节点地址" json:"apiNodeAddrs"`                    // API节点地址
	OfflineDay             string `gorm:"column:offlineDay;comment:下线日期YYYYMMDD" json:"offlineDay"`                   // 下线日期YYYYMMDD
	OfflineIsNotified      uint32 `gorm:"column:offlineIsNotified;comment:下线是否已通知" json:"offlineIsNotified"`          // 下线是否已通知
	IsBackupForCluster     uint32 `gorm:"column:isBackupForCluster;comment:是否为集群备用节点" json:"isBackupForCluster"`      // 是否为集群备用节点
	IsBackupForGroup       uint32 `gorm:"column:isBackupForGroup;comment:是否为分组备用节点" json:"isBackupForGroup"`          // 是否为分组备用节点
	BackupIPs              string `gorm:"column:backupIPs;comment:备用IP" json:"backupIPs"`                             // 备用IP
	ActionStatus           string `gorm:"column:actionStatus;comment:当前动作配置" json:"actionStatus"`                     // 当前动作配置
}

// TableName Node's table name
func (*Node) TableName() string {
	return TableNameNode
}

// AfterCreate run after create database record.
func (u *Node) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "node-")).Error
}

// NodeRef 列表数据
type NodeRef struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	NodeCount         uint32 `json:"nodeCount"`
	OnlineNodeCount   uint32 `json:"onlineNodeCount"`
	ServerCount       uint32 `json:"serverCount"`
	DNSName           string `json:"dnsName"`
}

// NodeRefList 返回列表
type NodeRefList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*NodeRef `json:"items"`
}

// NodeList 资源列表
type NodeList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*Node `json:"items"`
}

var NodeTableZeroFields = []string{"name", "lnAddrs", "isOn", "uniqueId", "code", "installDir", "installStatus", "state", "connectedAPINodes", "ddosProtection", "dnsRoutes", "maxCacheDiskCapacity", "maxCacheMemoryCapacity", "cacheDiskDir", "cacheDiskSubDirs", "dnsResolver", "apiNodeAddrs", "offlineDay", "backupIPs"}

