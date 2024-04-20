package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameHTTPWebsocket = "galloHTTPWebsockets"

// HTTPWebsocket Websocket设置
type HTTPWebsocket struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                                       // 管理员ID
	UserID            uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`                                          // 用户ID
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                                    // 状态
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                                    // 是否启用
	HandshakeTimeout  string `gorm:"column:handshakeTimeout;comment:握手超时时间" json:"handshakeTimeout"`                    // 握手超时时间
	AllowAllOrigins   uint32 `gorm:"column:allowAllOrigins;default:1;comment:是否支持所有源" json:"allowAllOrigins"`           // 是否支持所有源
	AllowedOrigins    string `gorm:"column:allowedOrigins;comment:支持的源域名列表" json:"allowedOrigins"`                      // 支持的源域名列表
	RequestSameOrigin uint32 `gorm:"column:requestSameOrigin;default:1;comment:是否请求一样的Origin" json:"requestSameOrigin"` // 是否请求一样的Origin
	RequestOrigin     string `gorm:"column:requestOrigin;comment:请求Origin" json:"requestOrigin"`                        // 请求Origin
	WebID             uint32 `gorm:"column:webId;comment:Web" json:"webId"`                                             // Web
}

// TableName HTTPWebsocket's table name
func (*HTTPWebsocket) TableName() string {
	return TableNameHTTPWebsocket
}

// AfterCreate run after create database record.
func (u *HTTPWebsocket) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "socket-")).Error
}

type HTTPWebsocketList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*HTTPWebsocket `json:"items"`
}
