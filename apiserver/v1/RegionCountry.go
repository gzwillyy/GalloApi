package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloRegionCountry = "galloRegionCountries"

// RegionCountry 区域-国家/地区
type RegionCountry struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	ValueID           uint32 `gorm:"column:valueId;comment:实际ID" json:"valueId"`          // 实际ID
	ValueCode         string `gorm:"column:valueCode;comment:值代号" json:"valueCode"`       // 值代号
	Codes             string `gorm:"column:codes;comment:代号" json:"codes"`                // 代号
	CustomName        string `gorm:"column:customName;comment:自定义名称" json:"customName"`   // 自定义名称
	CustomCodes       string `gorm:"column:customCodes;comment:自定义代号" json:"customCodes"` // 自定义代号
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`      // 状态
	DataID            string `gorm:"column:dataId;comment:原始数据ID" json:"dataId"`          // 原始数据ID
	Pinyin            string `gorm:"column:pinyin;comment:拼音" json:"pinyin"`              // 拼音
	IsCommon          bool   `gorm:"column:isCommon;comment:是否常用" json:"isCommon"`        // 是否常用
	RouteCode         string `gorm:"column:routeCode;comment:线路代号" json:"routeCode"`      // 线路代号
}

// TableName RegionCountry's table name
func (*RegionCountry) TableName() string {
	return TableNameGalloRegionCountry
}

// AfterCreate run after create database record.
func (u *RegionCountry) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "country-")).Error
}

type RegionCountryList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*RegionCountry `json:"items"`
}

var RegionCountryTableZeroFields = []string{"name", "codes", "customName", "customCodes", "state", "dataId", "isCommon", "routeCode"}

