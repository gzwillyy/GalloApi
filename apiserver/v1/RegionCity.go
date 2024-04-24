package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloRegionCity = "galloRegionCities"

// RegionCity 区域-城市
type RegionCity struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	ValueID           uint32 `gorm:"column:valueId;comment:实际ID" json:"valueId"`          // 实际ID
	ProvinceID        uint32 `gorm:"column:provinceId;comment:省份ID" json:"provinceId"`    // 省份ID
	Codes             string `gorm:"column:codes;comment:代号" json:"codes"`                // 代号
	CustomName        string `gorm:"column:customName;comment:自定义名称" json:"customName"`   // 自定义名称
	CustomCodes       string `gorm:"column:customCodes;comment:自定义代号" json:"customCodes"` // 自定义代号
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`      // 状态
	DataID            string `gorm:"column:dataId;comment:原始数据ID" json:"dataId"`          // 原始数据ID
}

// TableName RegionCity's table name
func (*RegionCity) TableName() string {
	return TableNameGalloRegionCity
}

// AfterCreate run after create database record.
func (u *RegionCity) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "city-")).Error
}

type RegionCityList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*RegionCity `json:"items"`
}

var RegionCityTableZeroFields = []string{"name", "customName", "customCodes", "state", "dataId"}

