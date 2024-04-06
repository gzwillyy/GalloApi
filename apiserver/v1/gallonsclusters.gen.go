// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloNSCluster = "galloNSClusters"

// GalloNSCluster 域名服务器集群
type GalloNSCluster struct {
	ID              int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`           // ID
	IsOn            bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                         // 是否启用
	Name            string `gorm:"column:name;comment:集群名" json:"name"`                                    // 集群名
	InstallDir      string `gorm:"column:installDir;comment:安装目录" json:"installDir"`                       // 安装目录
	State           bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                         // 状态
	AccessLog       string `gorm:"column:accessLog;comment:访问日志配置" json:"accessLog"`                       // 访问日志配置
	GrantID         int32  `gorm:"column:grantId;comment:授权ID" json:"grantId"`                             // 授权ID
	Recursion       string `gorm:"column:recursion;comment:递归DNS设置" json:"recursion"`                      // 递归DNS设置
	Tcp             string `gorm:"column:tcp;comment:TCP设置" json:"tcp"`                                    // TCP设置
	TLS             string `gorm:"column:tls;comment:TLS设置" json:"tls"`                                    // TLS设置
	Udp             string `gorm:"column:udp;comment:UDP设置" json:"udp"`                                    // UDP设置
	Doh             string `gorm:"column:doh;comment:DoH设置" json:"doh"`                                    // DoH设置
	DdosProtection  string `gorm:"column:ddosProtection;comment:DDoS防护设置" json:"ddosProtection"`           // DDoS防护设置
	Hosts           string `gorm:"column:hosts;comment:DNS主机地址" json:"hosts"`                              // DNS主机地址
	Soa             string `gorm:"column:soa;comment:SOA配置" json:"soa"`                                    // SOA配置
	AutoRemoteStart bool   `gorm:"column:autoRemoteStart;default:1;comment:自动远程启动" json:"autoRemoteStart"` // 自动远程启动
	TimeZone        string `gorm:"column:timeZone;comment:时区" json:"timeZone"`                             // 时区
	Answer          string `gorm:"column:answer;comment:应答设置" json:"answer"`                               // 应答设置
	SoaSerial       int64  `gorm:"column:soaSerial;comment:SOA序列号" json:"soaSerial"`                       // SOA序列号
	Email           string `gorm:"column:email;comment:管理员邮箱" json:"email"`                                // 管理员邮箱
	DetectAgents    bool   `gorm:"column:detectAgents;default:1;comment:是否监测Agents" json:"detectAgents"`   // 是否监测Agents
	CheckingPorts   bool   `gorm:"column:checkingPorts;default:1;comment:自动检测端口" json:"checkingPorts"`     // 自动检测端口
}

// TableName GalloNSCluster's table name
func (*GalloNSCluster) TableName() string {
	return TableNameGalloNSCluster
}