package v1

const TableNameGalloACMEAuthentication = "galloACMEAuthentications"

// GalloACMEAuthentication ACME认证
type GalloACMEAuthentication struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	TaskID    int64  `gorm:"column:taskId;comment:任务ID" json:"taskId"`                     // 任务ID
	Domain    string `gorm:"column:domain;comment:域名" json:"domain"`                       // 域名
	Token     string `gorm:"column:token;comment:令牌" json:"token"`                         // 令牌
	Key       string `gorm:"column:key;comment:密钥" json:"key"`                             // 密钥
	CreatedAt int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`               // 创建时间
}

// TableName GalloACMEAuthentication's table name
func (*GalloACMEAuthentication) TableName() string {
	return TableNameGalloACMEAuthentication
}