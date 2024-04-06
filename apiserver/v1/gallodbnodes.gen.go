// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloDBNode = "galloDBNodes"

// GalloDBNode 数据库节点
type GalloDBNode struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	IsOn        bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`               // 是否启用
	Role        string `gorm:"column:role;comment:数据库角色" json:"role"`                        // 数据库角色
	Name        string `gorm:"column:name;comment:名称" json:"name"`                           // 名称
	Description string `gorm:"column:description;comment:描述" json:"description"`             // 描述
	Host        string `gorm:"column:host;comment:主机" json:"host"`                           // 主机
	Port        int32  `gorm:"column:port;comment:端口" json:"port"`                           // 端口
	Database    string `gorm:"column:database;comment:数据库名称" json:"database"`                // 数据库名称
	Username    string `gorm:"column:username;comment:用户名" json:"username"`                  // 用户名
	Password    string `gorm:"column:password;comment:密码" json:"password"`                   // 密码
	Charset     string `gorm:"column:charset;comment:通讯字符集" json:"charset"`                  // 通讯字符集
	ConnTimeout int32  `gorm:"column:connTimeout;comment:连接超时时间（秒）" json:"connTimeout"`      // 连接超时时间（秒）
	State       bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
	CreatedAt   int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`               // 创建时间
	Weight      int32  `gorm:"column:weight;comment:权重" json:"weight"`                       // 权重
	Order       int32  `gorm:"column:order;comment:排序" json:"order"`                         // 排序
	AdminID     int32  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                  // 管理员ID
}

// TableName GalloDBNode's table name
func (*GalloDBNode) TableName() string {
	return TableNameGalloDBNode
}