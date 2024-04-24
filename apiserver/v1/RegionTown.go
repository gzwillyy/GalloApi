package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloRegionTown = "galloRegionTowns"

// RegionTown 区域-省份
type RegionTown struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	ValueID           uint32 `gorm:"column:valueId;comment:真实ID" json:"valueId"`          // 真实ID
	CityID            uint32 `gorm:"column:cityId;comment:城市ID" json:"cityId"`            // 城市ID
	Codes             string `gorm:"column:codes;comment:代号" json:"codes"`                // 代号
	CustomName        string `gorm:"column:customName;comment:自定义名称" json:"customName"`   // 自定义名称
	CustomCodes       string `gorm:"column:customCodes;comment:自定义代号" json:"customCodes"` // 自定义代号
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`      // 状态
	DataID            string `gorm:"column:dataId;comment:原始数据ID" json:"dataId"`          // 原始数据ID
}

// TableName RegionTown's table name
func (*RegionTown) TableName() string {
	return TableNameGalloRegionTown
}

// AfterCreate run after create database record.
func (u *RegionTown) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "town-")).Error
}

type RegionTownList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*RegionTown `json:"items"`
}

var RegionTownTableZeroFields = []string{"name", "customName", "customCodes", "state", "dataId"}

