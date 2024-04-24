package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPWeb = "galloHTTPWebs"

// HTTPWeb HTTP Web
type HTTPWeb struct {
	metav1.ObjectMeta  `json:"metadata,omitempty"`
	IsOn               bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                               // 是否启用
	TemplateID         uint32 `gorm:"column:templateId;comment:模版ID" json:"templateId"`                             // 模版ID
	AdminID            uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                                  // 管理员ID
	UserID             uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`                                     // 用户ID
	State              bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                               // 状态
	Root               string `gorm:"column:root;comment:根目录 json" json:"root"`                                     // 根目录 json
	Charset            string `gorm:"column:charset;comment:字符集 json" json:"charset"`                               // 字符集 json
	Shutdown           string `gorm:"column:shutdown;comment:临时关闭页面配置 json" json:"shutdown"`                        // 临时关闭页面配置 json
	Pages              string `gorm:"column:pages;comment:特殊页面 json" json:"pages"`                                  // 特殊页面 json
	RedirectToHTTPS    string `gorm:"column:redirectToHttps;comment:跳转到HTTPS设置 json" json:"redirectToHttps"`        // 跳转到HTTPS设置 json
	Indexes            string `gorm:"column:indexes;comment:首页文件列表 json" json:"indexes"`                            // 首页文件列表 json
	MaxRequestBodySize string `gorm:"column:maxRequestBodySize;comment:最大允许的请求内容尺寸 json" json:"maxRequestBodySize"` // 最大允许的请求内容尺寸 json
	RequestHeader      string `gorm:"column:requestHeader;comment:请求Header配置 json" json:"requestHeader"`            // 请求Header配置 json
	ResponseHeader     string `gorm:"column:responseHeader;comment:响应Header配置 json" json:"responseHeader"`          // 响应Header配置 json
	AccessLog          string `gorm:"column:accessLog;comment:访问日志配置 json" json:"accessLog"`                        // 访问日志配置 json
	Stat               string `gorm:"column:stat;comment:统计配置 json" json:"stat"`                                    // 统计配置 json
	Gzip               string `gorm:"column:gzip;comment:Gzip配置（弃用）json" json:"gzip"`                               // Gzip配置（弃用）json
	Compression        string `gorm:"column:compression;comment:压缩配置 json" json:"compression"`                      // 压缩配置 json
	Cache              string `gorm:"column:cache;comment:缓存配置 json" json:"cache"`                                  // 缓存配置 json
	Firewall           string `gorm:"column:firewall;comment:防火墙设置 json" json:"firewall"`                           // 防火墙设置 json
	Locations          string `gorm:"column:locations;comment:路由规则配置 json" json:"locations"`                        // 路由规则配置 json
	Websocket          string `gorm:"column:websocket;comment:Websocket设置 json" json:"websocket"`                   // Websocket设置 json
	RewriteRules       string `gorm:"column:rewriteRules;comment:重写规则配置 json" json:"rewriteRules"`                  // 重写规则配置 json
	HostRedirects      string `gorm:"column:hostRedirects;comment:域名跳转 json" json:"hostRedirects"`                  // 域名跳转 json
	Fastcgi            string `gorm:"column:fastcgi;comment:Fastcgi配置 json" json:"fastcgi"`                         // Fastcgi配置 json
	Auth               string `gorm:"column:auth;comment:认证策略配置 json" json:"auth"`                                  // 认证策略配置 json
	Webp               string `gorm:"column:webp;comment:WebP配置 json" json:"webp"`                                  // WebP配置 json
	RemoteAddr         string `gorm:"column:remoteAddr;comment:客户端IP配置 json" json:"remoteAddr"`                     // 客户端IP配置 json
	MergeSlashes       uint32 `gorm:"column:mergeSlashes;comment:是否合并路径中的斜杠" json:"mergeSlashes"`                   // 是否合并路径中的斜杠
	RequestLimit       string `gorm:"column:requestLimit;comment:请求限制 json" json:"requestLimit"`                    // 请求限制 json
	RequestScripts     string `gorm:"column:requestScripts;comment:请求脚本 json" json:"requestScripts"`                // 请求脚本 json
	Uam                string `gorm:"column:uam;comment:UAM设置 json" json:"uam"`                                     // UAM设置 json
	Cc                 string `gorm:"column:cc;comment:CC设置 json" json:"cc"`                                        // CC设置 json
	Referers           string `gorm:"column:referers;comment:防盗链设置 json" json:"referers"`                           // 防盗链设置 json
	UserAgent          string `gorm:"column:userAgent;comment:UserAgent设置 json" json:"userAgent"`                   // UserAgent设置 json
	Optimization       string `gorm:"column:optimization;comment:页面优化配置 json" json:"optimization"`                  // 页面优化配置 json
	Hls                string `gorm:"column:hls;comment:HLS设置 json" json:"hls"`                                     // HLS设置 json
}

// TableName HTTPWeb's table name
func (*HTTPWeb) TableName() string {
	return TableNameHTTPWeb
}

// AfterCreate run after create database record.
func (u *HTTPWeb) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "web-")).Error
}

// HTTPWebList 返回列表
type HTTPWebList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPWeb `json:"items"`
}

// CreateHTTPWebRequest 创建分组
type CreateHTTPWebRequest struct {
	AdminID uint32 `json:"adminId"` // 管理员ID
	UserID  uint32 `json:"userId"`  // 用户ID
	Root    string `json:"root"`    // 根目录
}

// UpdateHTTPWebRequest 修改分组
type UpdateHTTPWebRequest struct {
	InstanceID string `json:"instanceID"`
	Name       string `json:"name"`
}

// DeleteHTTPWebRequest 删除分组
type DeleteHTTPWebRequest struct {
	InstanceID string `json:"instanceID"`
}

var HTTPWebTableZeroFields = []string{"name", "isOn", "root", "requestLimit"}

