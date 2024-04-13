package v1

const TableNameGalloRegionProvider = "galloRegionProviders"

// RegionProvider 区域-运营商
type RegionProvider struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	ValueID     int32  `gorm:"column:valueId;comment:实际ID" json:"valueId"`                   // 实际ID
	Name        string `gorm:"column:name;comment:名称" json:"name"`                           // 名称
	Codes       string `gorm:"column:codes;comment:代号" json:"codes"`                         // 代号
	CustomName  string `gorm:"column:customName;comment:自定义名称" json:"customName"`            // 自定义名称
	CustomCodes string `gorm:"column:customCodes;comment:自定义代号" json:"customCodes"`          // 自定义代号
	State       bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
}

// TableName RegionProvider's table name
func (*RegionProvider) TableName() string {
	return TableNameGalloRegionProvider
}
