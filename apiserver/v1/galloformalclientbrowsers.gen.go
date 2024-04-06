// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloFormalClientBrowser = "galloFormalClientBrowsers"

// GalloFormalClientBrowser 终端浏览器信息
type GalloFormalClientBrowser struct {
	ID     int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Name   string `gorm:"column:name;comment:浏览器名称" json:"name"`                        // 浏览器名称
	Codes  string `gorm:"column:codes;comment:代号" json:"codes"`                         // 代号
	DataID string `gorm:"column:dataId;comment:数据ID" json:"dataId"`                     // 数据ID
	State  bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
}

// TableName GalloFormalClientBrowser's table name
func (*GalloFormalClientBrowser) TableName() string {
	return TableNameGalloFormalClientBrowser
}