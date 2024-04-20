package v1

import metav1 "github.com/gzwillyy/components/pkg/meta/v1"

const TableNameGalloRegionTown = "galloRegionTowns"

// RegionTown 区域-省份
type RegionTown struct {
	ID          uint32 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	ValueID     uint32 `gorm:"column:valueId;comment:真实ID" json:"valueId"`                   // 真实ID
	CityID      uint32 `gorm:"column:cityId;comment:城市ID" json:"cityId"`                     // 城市ID
	Name        string `gorm:"column:name;comment:名称" json:"name"`                           // 名称
	Codes       string `gorm:"column:codes;comment:代号" json:"codes"`                         // 代号
	CustomName  string `gorm:"column:customName;comment:自定义名称" json:"customName"`            // 自定义名称
	CustomCodes string `gorm:"column:customCodes;comment:自定义代号" json:"customCodes"`          // 自定义代号
	State       bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
	DataID      string `gorm:"column:dataId;comment:原始数据ID" json:"dataId"`                   // 原始数据ID
}

// TableName RegionTown's table name
func (*RegionTown) TableName() string {
	return TableNameGalloRegionTown
}

type RegionTownList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*RegionTown `json:"items"`
}
