package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloFormalClientBrowser = "galloFormalClientBrowsers"

// FormalClientBrowser 终端浏览器信息
type FormalClientBrowser struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Codes             string `gorm:"column:codes;comment:代号" json:"codes"`           // 代号
	DataID            string `gorm:"column:dataId;comment:数据ID" json:"dataId"`       // 数据ID
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
}

// TableName FormalClientBrowser's table name
func (*FormalClientBrowser) TableName() string {
	return TableNameGalloFormalClientBrowser
}

// AfterCreate run after create database record.
func (u *FormalClientBrowser) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "browser-")).Error
}

type FormalClientBrowserList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*FormalClientBrowser `json:"items"`
}

var FormalClientBrowserTableZeroFields = []string{"name", "codes", "dataId"}

