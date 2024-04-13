package v1

import (
	"github.com/gzwillyy/components/pkg/util/idutil"
	"gorm.io/gorm"
)

const TableNameGalloHTTPWeb = "galloHTTPWebs"

// HTTPWeb HTTP Web
type HTTPWeb struct {
	ID                 int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`            // ID
	IsOn               bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                          // 是否启用
	TemplateID         int32  `gorm:"column:templateId;comment:模版ID" json:"templateId"`                        // 模版ID
	AdminID            int32  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                             // 管理员ID
	UserID             int32  `gorm:"column:userId;comment:用户ID" json:"userId"`                                // 用户ID
	State              bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                          // 状态
	CreatedAt          int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`                          // 创建时间
	Root               string `gorm:"column:root;comment:根目录" json:"root"`                                     // 根目录
	Charset            string `gorm:"column:charset;comment:字符集" json:"charset"`                               // 字符集
	Shutdown           string `gorm:"column:shutdown;comment:临时关闭页面配置" json:"shutdown"`                        // 临时关闭页面配置
	Pages              string `gorm:"column:pages;comment:特殊页面" json:"pages"`                                  // 特殊页面
	RedirectToHTTPS    string `gorm:"column:redirectToHttps;comment:跳转到HTTPS设置" json:"redirectToHttps"`        // 跳转到HTTPS设置
	Indexes            string `gorm:"column:indexes;comment:首页文件列表" json:"indexes"`                            // 首页文件列表
	MaxRequestBodySize string `gorm:"column:maxRequestBodySize;comment:最大允许的请求内容尺寸" json:"maxRequestBodySize"` // 最大允许的请求内容尺寸
	RequestHeader      string `gorm:"column:requestHeader;comment:请求Header配置" json:"requestHeader"`            // 请求Header配置
	ResponseHeader     string `gorm:"column:responseHeader;comment:响应Header配置" json:"responseHeader"`          // 响应Header配置
	AccessLog          string `gorm:"column:accessLog;comment:访问日志配置" json:"accessLog"`                        // 访问日志配置
	Stat               string `gorm:"column:stat;comment:统计配置" json:"stat"`                                    // 统计配置
	Gzip               string `gorm:"column:gzip;comment:Gzip配置（弃用）" json:"gzip"`                              // Gzip配置（弃用）
	Compression        string `gorm:"column:compression;comment:压缩配置" json:"compression"`                      // 压缩配置
	Cache              string `gorm:"column:cache;comment:缓存配置" json:"cache"`                                  // 缓存配置
	Firewall           string `gorm:"column:firewall;comment:防火墙设置" json:"firewall"`                           // 防火墙设置
	Locations          string `gorm:"column:locations;comment:路由规则配置" json:"locations"`                        // 路由规则配置
	Websocket          string `gorm:"column:websocket;comment:Websocket设置" json:"websocket"`                   // Websocket设置
	RewriteRules       string `gorm:"column:rewriteRules;comment:重写规则配置" json:"rewriteRules"`                  // 重写规则配置
	HostRedirects      string `gorm:"column:hostRedirects;comment:域名跳转" json:"hostRedirects"`                  // 域名跳转
	Fastcgi            string `gorm:"column:fastcgi;comment:Fastcgi配置" json:"fastcgi"`                         // Fastcgi配置
	Auth               string `gorm:"column:auth;comment:认证策略配置" json:"auth"`                                  // 认证策略配置
	Webp               string `gorm:"column:webp;comment:WebP配置" json:"webp"`                                  // WebP配置
	RemoteAddr         string `gorm:"column:remoteAddr;comment:客户端IP配置" json:"remoteAddr"`                     // 客户端IP配置
	MergeSlashes       bool   `gorm:"column:mergeSlashes;comment:是否合并路径中的斜杠" json:"mergeSlashes"`              // 是否合并路径中的斜杠
	RequestLimit       string `gorm:"column:requestLimit;comment:请求限制" json:"requestLimit"`                    // 请求限制
	RequestScripts     string `gorm:"column:requestScripts;comment:请求脚本" json:"requestScripts"`                // 请求脚本
	Uam                string `gorm:"column:uam;comment:UAM设置" json:"uam"`                                     // UAM设置
	Cc                 string `gorm:"column:cc;comment:CC设置" json:"cc"`                                        // CC设置
	Referers           string `gorm:"column:referers;comment:防盗链设置" json:"referers"`                           // 防盗链设置
	UserAgent          string `gorm:"column:userAgent;comment:UserAgent设置" json:"userAgent"`                   // UserAgent设置
	Optimization       string `gorm:"column:optimization;comment:页面优化配置" json:"optimization"`                  // 页面优化配置
	Hls                string `gorm:"column:hls;comment:HLS设置" json:"hls"`                                     // HLS设置
}

// TableName HTTPWeb's table name
func (*HTTPWeb) TableName() string {
	return TableNameGalloHTTPWeb
}

// AfterCreate run after create database record.
func (u *NodeGroup) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "group-")).Error
}
