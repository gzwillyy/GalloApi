package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloServer = "galloServers"

// Server 服务
type Server struct {
	metav1.ObjectMeta   `json:"metadata,omitempty"`
	IsOn                bool    `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                                     // 是否启用
	UserID              uint32  `gorm:"column:userId;comment:用户ID" json:"userId"`                                           // 用户ID
	AdminID             uint32  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                                        // 管理员ID
	Type                string  `gorm:"column:type;comment:服务类型" json:"type"`                                               // 服务类型
	Description         string  `gorm:"column:description;comment:描述" json:"description"`                                   // 描述
	PlainServerNames    string  `gorm:"column:plainServerNames;comment:扁平化域名列表" json:"plainServerNames"`                    // 扁平化域名列表
	ServerNames         string  `gorm:"column:serverNames;comment:域名列表" json:"serverNames"`                                 // 域名列表
	AuditingAt          uint32   `gorm:"column:auditingAt;comment:审核提交时间" json:"auditingAt"`                                 // 审核提交时间
	AuditingServerNames string  `gorm:"column:auditingServerNames;comment:审核中的域名" json:"auditingServerNames"`               // 审核中的域名
	IsAuditing          bool    `gorm:"column:isAuditing;comment:是否正在审核" json:"isAuditing"`                                 // 是否正在审核
	AuditingResult      string  `gorm:"column:auditingResult;comment:审核结果" json:"auditingResult"`                           // 审核结果
	HTTP                string  `gorm:"column:http;comment:HTTP配置" json:"http"`                                             // HTTP配置
	HTTPS               string  `gorm:"column:https;comment:HTTPS配置" json:"https"`                                          // HTTPS配置
	Tcp                 string  `gorm:"column:tcp;comment:TCP配置" json:"tcp"`                                                // TCP配置
	TLS                 string  `gorm:"column:tls;comment:TLS配置" json:"tls"`                                                // TLS配置
	Unix                string  `gorm:"column:unix;comment:Unix配置" json:"unix"`                                             // Unix配置
	Udp                 string  `gorm:"column:udp;comment:UDP配置" json:"udp"`                                                // UDP配置
	WebID               uint32  `gorm:"column:webId;comment:WEB配置" json:"webId"`                                            // WEB配置
	ReverseProxy        string  `gorm:"column:reverseProxy;comment:反向代理配置" json:"reverseProxy"`                             // 反向代理配置
	GroupIds            string  `gorm:"column:groupIds;comment:分组ID列表" json:"groupIds"`                                     // 分组ID列表
	Config              string  `gorm:"column:config;comment:服务配置，自动生成" json:"config"`                                      // 服务配置，自动生成
	ConfigMd5           string  `gorm:"column:configMd5;comment:Md5" json:"configMd5"`                                      // Md5
	ClusterID           uint32  `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`                                     // 集群ID
	IncludeNodes        string  `gorm:"column:includeNodes;comment:部署条件" json:"includeNodes"`                               // 部署条件
	ExcludeNodes        string  `gorm:"column:excludeNodes;comment:节点排除条件" json:"excludeNodes"`                             // 节点排除条件
	Version             uint32  `gorm:"column:version;comment:版本号" json:"version"`                                          // 版本号
	State               bool    `gorm:"column:state;default:1;comment:状态" json:"state"`                                     // 状态
	DNSName             string  `gorm:"column:dnsName;comment:DNS名称" json:"dnsName"`                                        // DNS名称
	TcpPorts            string  `gorm:"column:tcpPorts;comment:所包含TCP端口" json:"tcpPorts"`                                   // 所包含TCP端口
	UdpPorts            string  `gorm:"column:udpPorts;comment:所包含UDP端口" json:"udpPorts"`                                   // 所包含UDP端口
	SupportCNAME        bool    `gorm:"column:supportCNAME;comment:允许CNAME不在域名名单" json:"supportCNAME"`                      // 允许CNAME不在域名名单
	TrafficLimit        string  `gorm:"column:trafficLimit;comment:流量限制" json:"trafficLimit"`                               // 流量限制
	TrafficDay          string  `gorm:"column:trafficDay;comment:YYYYMMDD" json:"trafficDay"`                               // YYYYMMDD
	TrafficMonth        string  `gorm:"column:trafficMonth;comment:YYYYMM" json:"trafficMonth"`                             // YYYYMM
	TotalDailyTraffic   float64 `gorm:"column:totalDailyTraffic;default:0.000000;comment:日流量" json:"totalDailyTraffic"`     // 日流量
	TotalMonthlyTraffic float64 `gorm:"column:totalMonthlyTraffic;default:0.000000;comment:月流量" json:"totalMonthlyTraffic"` // 月流量
	TrafficLimitStatus  string  `gorm:"column:trafficLimitStatus;comment:流量限制状态" json:"trafficLimitStatus"`                 // 流量限制状态
	TotalTraffic        float64 `gorm:"column:totalTraffic;default:0.000000;comment:总流量" json:"totalTraffic"`               // 总流量
	UserPlanID          uint32  `gorm:"column:userPlanId;comment:所属套餐ID" json:"userPlanId"`                                 // 所属套餐ID
	LastUserPlanID      uint32  `gorm:"column:lastUserPlanId;comment:上一次使用的套餐" json:"lastUserPlanId"`                       // 上一次使用的套餐
	Uam                 string  `gorm:"column:uam;comment:UAM设置" json:"uam"`                                                // UAM设置
	BandwidthTime       string  `gorm:"column:bandwidthTime;comment:带宽更新时间，YYYYMMDDHHII" json:"bandwidthTime"`              // 带宽更新时间，YYYYMMDDHHII
	BandwidthBytes      uint32   `gorm:"column:bandwidthBytes;comment:最近带宽峰值" json:"bandwidthBytes"`                         // 最近带宽峰值
	CountAttackRequests uint32   `gorm:"column:countAttackRequests;comment:最近攻击请求数" json:"countAttackRequests"`              // 最近攻击请求数
	CountRequests       uint32   `gorm:"column:countRequests;comment:最近总请求数" json:"countRequests"`                           // 最近总请求数
}

// TableName Server's table name
func (*Server) TableName() string {
	return TableNameGalloServer
}

func (u *Server) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "server-")).Error
}

// ServerList 返回列表
type ServerList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*ACMEUser `json:"items"`
}

var ServerTableZeroFields = []string{"name", "isOn", "description", "plainServerNames", "serverNames", "auditingServerNames", "isAuditing", "auditingResult", "http", "https", "tcp", "tls", "unix", "udp", "reverseProxy", "groupIds", "config", "configMd5", "excludeNodes", "state", "dnsName", "tcpPorts", "udpPorts", "supportCNAME", "trafficLimit", "trafficDay", "trafficLimitStatus", "uam", "bandwidthTime"}

