package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloFormalClientSystem = "galloFormalClientSystems"

// FormalClientSystem 终端操作系统信息
type FormalClientSystem struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Codes             string `gorm:"column:codes;comment:代号" json:"codes"`           // 代号
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
	DataID            string `gorm:"column:dataId;comment:数据ID" json:"dataId"`       // 数据ID
}

// TableName FormalClientSystem's table name
func (*FormalClientSystem) TableName() string {
	return TableNameGalloFormalClientSystem
}

// AfterCreate run after create database record.
func (u *FormalClientSystem) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "system-")).Error
}

type FormalClientSystemList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*FormalClientSystem `json:"items"`
}

var FormalClientSystemTableZeroFields = []string{"name", "codes", "dataId"}
