// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloReverseProxy = "galloReverseProxies"

// GalloReverseProxy 反向代理配置
type GalloReverseProxy struct {
	ID                       int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                         // ID
	AdminID                  int32  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                                          // 管理员ID
	UserID                   int32  `gorm:"column:userId;comment:用户ID" json:"userId"`                                             // 用户ID
	TemplateID               int32  `gorm:"column:templateId;comment:模版ID" json:"templateId"`                                     // 模版ID
	IsOn                     bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                                       // 是否启用
	Scheduling               string `gorm:"column:scheduling;comment:调度算法" json:"scheduling"`                                     // 调度算法
	PrimaryOrigins           string `gorm:"column:primaryOrigins;comment:主要源站" json:"primaryOrigins"`                             // 主要源站
	BackupOrigins            string `gorm:"column:backupOrigins;comment:备用源站" json:"backupOrigins"`                               // 备用源站
	StripPrefix              string `gorm:"column:stripPrefix;comment:去除URL前缀" json:"stripPrefix"`                                // 去除URL前缀
	RequestHostType          bool   `gorm:"column:requestHostType;comment:请求Host类型" json:"requestHostType"`                       // 请求Host类型
	RequestHost              string `gorm:"column:requestHost;comment:请求Host" json:"requestHost"`                                 // 请求Host
	RequestHostExcludingPort bool   `gorm:"column:requestHostExcludingPort;comment:移除请求Host中的域名" json:"requestHostExcludingPort"` // 移除请求Host中的域名
	RequestURI               string `gorm:"column:requestURI;comment:请求URI" json:"requestURI"`                                    // 请求URI
	AutoFlush                bool   `gorm:"column:autoFlush;comment:是否自动刷新缓冲区" json:"autoFlush"`                                  // 是否自动刷新缓冲区
	AddHeaders               string `gorm:"column:addHeaders;comment:自动添加的Header列表" json:"addHeaders"`                            // 自动添加的Header列表
	State                    bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                                       // 状态
	CreatedAt                int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`                                       // 创建时间
	ConnTimeout              string `gorm:"column:connTimeout;comment:连接超时时间" json:"connTimeout"`                                 // 连接超时时间
	ReadTimeout              string `gorm:"column:readTimeout;comment:读取超时时间" json:"readTimeout"`                                 // 读取超时时间
	IdleTimeout              string `gorm:"column:idleTimeout;comment:空闲超时时间" json:"idleTimeout"`                                 // 空闲超时时间
	MaxConns                 int32  `gorm:"column:maxConns;comment:最大并发连接数" json:"maxConns"`                                      // 最大并发连接数
	MaxIdleConns             int32  `gorm:"column:maxIdleConns;comment:最大空闲连接数" json:"maxIdleConns"`                              // 最大空闲连接数
	ProxyProtocol            string `gorm:"column:proxyProtocol;comment:Proxy Protocol配置" json:"proxyProtocol"`                   // Proxy Protocol配置
	FollowRedirects          bool   `gorm:"column:followRedirects;comment:回源跟随" json:"followRedirects"`                           // 回源跟随
	Retry50X                 bool   `gorm:"column:retry50X;comment:启用50X重试" json:"retry50X"`                                      // 启用50X重试
	Retry40X                 bool   `gorm:"column:retry40X;comment:启用40X重试" json:"retry40X"`                                      // 启用40X重试
}

// TableName GalloReverseProxy's table name
func (*GalloReverseProxy) TableName() string {
	return TableNameGalloReverseProxy
}