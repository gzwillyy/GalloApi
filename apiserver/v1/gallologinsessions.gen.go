// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloLoginSession = "galloLoginSessions"

// GalloLoginSession 登录Session
type GalloLoginSession struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	AdminID   int64  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                  // 管理员ID
	UserID    int64  `gorm:"column:userId;comment:用户ID" json:"userId"`                     // 用户ID
	Sid       string `gorm:"column:sid;comment:令牌" json:"sid"`                             // 令牌
	Values    string `gorm:"column:values;comment:数据" json:"values"`                       // 数据
	IP        string `gorm:"column:ip;comment:登录IP" json:"ip"`                             // 登录IP
	CreatedAt int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`               // 创建时间
	ExpiresAt int64  `gorm:"column:expiresAt;comment:过期时间" json:"expiresAt"`               // 过期时间
}

// TableName GalloLoginSession's table name
func (*GalloLoginSession) TableName() string {
	return TableNameGalloLoginSession
}