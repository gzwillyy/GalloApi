// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloHTTPFirewallRuleGroup = "galloHTTPFirewallRuleGroups"

// GalloHTTPFirewallRuleGroup 防火墙规则分组
type GalloHTTPFirewallRuleGroup struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	IsOn        bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`               // 是否启用
	Name        string `gorm:"column:name;comment:名称" json:"name"`                           // 名称
	Description string `gorm:"column:description;comment:描述" json:"description"`             // 描述
	Code        string `gorm:"column:code;comment:代号" json:"code"`                           // 代号
	IsTemplate  bool   `gorm:"column:isTemplate;comment:是否为预置模板" json:"isTemplate"`          // 是否为预置模板
	AdminID     int32  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                  // 管理员ID
	UserID      int32  `gorm:"column:userId;comment:用户ID" json:"userId"`                     // 用户ID
	State       bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
	Sets        string `gorm:"column:sets;comment:规则集列表" json:"sets"`                        // 规则集列表
	CreatedAt   int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`               // 创建时间
}

// TableName GalloHTTPFirewallRuleGroup's table name
func (*GalloHTTPFirewallRuleGroup) TableName() string {
	return TableNameGalloHTTPFirewallRuleGroup
}