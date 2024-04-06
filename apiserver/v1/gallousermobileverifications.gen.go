// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloUserMobileVerification = "galloUserMobileVerifications"

// GalloUserMobileVerification 邮箱激活邮件队列
type GalloUserMobileVerification struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Mobile     string `gorm:"column:mobile;comment:手机号码" json:"mobile"`                     // 手机号码
	UserID     int64  `gorm:"column:userId;comment:用户ID" json:"userId"`                     // 用户ID
	Code       string `gorm:"column:code;comment:激活码" json:"code"`                          // 激活码
	CreatedAt  int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`               // 创建时间
	IsSent     bool   `gorm:"column:isSent;comment:是否已发送" json:"isSent"`                    // 是否已发送
	IsVerified bool   `gorm:"column:isVerified;comment:是否已激活" json:"isVerified"`            // 是否已激活
	Day        string `gorm:"column:day;comment:YYYYMMDD" json:"day"`                       // YYYYMMDD
}

// TableName GalloUserMobileVerification's table name
func (*GalloUserMobileVerification) TableName() string {
	return TableNameGalloUserMobileVerification
}