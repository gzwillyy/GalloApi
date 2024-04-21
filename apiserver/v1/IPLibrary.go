package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloIPLibrary = "galloIPLibraries"

// IPLibrary IP库
type IPLibrary struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`    // 管理员ID
	FileID            uint32 `gorm:"column:fileId;comment:文件ID" json:"fileId"`       // 文件ID
	Type              string `gorm:"column:type;comment:类型" json:"type"`             // 类型
	IsPublic          bool   `gorm:"column:isPublic;comment:是否公用" json:"isPublic"`   // 是否公用
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"` // 状态
}

// TableName GalloIPLibrary's table name
func (*IPLibrary) TableName() string {
	return TableNameGalloIPLibrary
}

// AfterCreate run after create database record.
func (u *IPLibrary) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "library-")).Error
}

type IPLibraryList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*IPLibrary `json:"items"`
}
