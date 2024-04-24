package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloRegionProvince = "galloRegionProvinces"

// RegionProvince 区域-省份
type RegionProvince struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	ValueID           uint32 `gorm:"column:valueId;comment:实际ID" json:"valueId"`          // 实际ID
	CountryID         uint32 `gorm:"column:countryId;comment:国家ID" json:"countryId"`      // 国家ID
	Codes             string `gorm:"column:codes;comment:代号" json:"codes"`                // 代号
	CustomName        string `gorm:"column:customName;comment:自定义名称" json:"customName"`   // 自定义名称
	CustomCodes       string `gorm:"column:customCodes;comment:自定义代号" json:"customCodes"` // 自定义代号
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`      // 状态
	DataID            string `gorm:"column:dataId;comment:原始数据ID" json:"dataId"`          // 原始数据ID
	RouteCode         string `gorm:"column:routeCode;comment:线路代号" json:"routeCode"`      // 线路代号
}

// TableName RegionProvince's table name
func (*RegionProvince) TableName() string {
	return TableNameGalloRegionProvince
}

// AfterCreate run after create database record.
func (u *RegionProvince) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "province-")).Error
}

type RegionProvinceList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*RegionProvince `json:"items"`
}

var RegionProvinceTableZeroFields = []string{"name", "customName", "customCodes", "state", "dataId"}

