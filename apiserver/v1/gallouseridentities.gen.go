// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloUserIdentity = "galloUserIdentities"

// GalloUserIdentity 用户实名认证信息
type GalloUserIdentity struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`            // ID
	UserID       int64  `gorm:"column:userId;comment:用户ID" json:"userId"`                                // 用户ID
	OrgType      string `gorm:"column:orgType;comment:组织类型" json:"orgType"`                              // 组织类型
	Type         string `gorm:"column:type;comment:证件类型" json:"type"`                                    // 证件类型
	RealName     string `gorm:"column:realName;comment:真实姓名" json:"realName"`                            // 真实姓名
	Number       string `gorm:"column:number;comment:编号" json:"number"`                                  // 编号
	FileIds      string `gorm:"column:fileIds;comment:文件ID" json:"fileIds"`                              // 文件ID
	Status       string `gorm:"column:status;comment:状态：none,submitted,verified,rejected" json:"status"` // 状态：none,submitted,verified,rejected
	State        bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                          // 状态
	CreatedAt    int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`                          // 创建时间
	UpdatedAt    int64  `gorm:"column:updatedAt;comment:修改时间" json:"updatedAt"`                          // 修改时间
	SubmittedAt  int64  `gorm:"column:submittedAt;comment:提交时间" json:"submittedAt"`                      // 提交时间
	RejectedAt   int64  `gorm:"column:rejectedAt;comment:拒绝时间" json:"rejectedAt"`                        // 拒绝时间
	VerifiedAt   int64  `gorm:"column:verifiedAt;comment:认证时间" json:"verifiedAt"`                        // 认证时间
	RejectReason string `gorm:"column:rejectReason;comment:拒绝原因" json:"rejectReason"`                    // 拒绝原因
}

// TableName GalloUserIdentity's table name
func (*GalloUserIdentity) TableName() string {
	return TableNameGalloUserIdentity
}