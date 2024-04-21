package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloRegionProvider = "galloRegionProviders"

// RegionProvider 区域-运营商
type RegionProvider struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	ValueID           uint32 `gorm:"column:valueId;comment:实际ID" json:"valueId"`          // 实际ID
	Codes             string `gorm:"column:codes;comment:代号" json:"codes"`                // 代号
	CustomName        string `gorm:"column:customName;comment:自定义名称" json:"customName"`   // 自定义名称
	CustomCodes       string `gorm:"column:customCodes;comment:自定义代号" json:"customCodes"` // 自定义代号
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`      // 状态
}

// TableName RegionProvider's table name
func (*RegionProvider) TableName() string {
	return TableNameGalloRegionProvider
}

// AfterCreate run after create database record.
func (u *RegionProvider) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "provider-")).Error
}

type RegionProviderList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*RegionProvider `json:"items"`
}
