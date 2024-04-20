package v1

import metav1 "github.com/gzwillyy/components/pkg/meta/v1"

const TableNameGalloRegionCity = "galloRegionCities"

// RegionCity 区域-城市
type RegionCity struct {
	ID          uint32 `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	ValueID     uint32 `gorm:"column:valueId;comment:实际ID" json:"valueId"`                   // 实际ID
	ProvinceID  uint32 `gorm:"column:provinceId;comment:省份ID" json:"provinceId"`             // 省份ID
	Name        string `gorm:"column:name;comment:名称" json:"name"`                           // 名称
	Codes       string `gorm:"column:codes;comment:代号" json:"codes"`                         // 代号
	CustomName  string `gorm:"column:customName;comment:自定义名称" json:"customName"`            // 自定义名称
	CustomCodes string `gorm:"column:customCodes;comment:自定义代号" json:"customCodes"`          // 自定义代号
	State       bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
	DataID      string `gorm:"column:dataId;comment:原始数据ID" json:"dataId"`                   // 原始数据ID
}

// TableName RegionCity's table name
func (*RegionCity) TableName() string {
	return TableNameGalloRegionCity
}

type RegionCityList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*RegionCity `json:"items"`
}
