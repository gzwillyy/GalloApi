package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameGalloFile = "galloFiles"

// File 文件管理
type File struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AdminID           uint32 `gorm:"column:adminId;comment:管理员ID" json:"adminId"`         // 管理员ID
	Code              string `gorm:"column:code;comment:代号" json:"code"`                  // 代号
	UserID            uint32 `gorm:"column:userId;comment:用户ID" json:"userId"`            // 用户ID
	Description       string `gorm:"column:description;comment:文件描述" json:"description"`  // 文件描述
	Size              uint32 `gorm:"column:size;comment:文件尺寸" json:"size"`                // 文件尺寸
	MimeType          string `gorm:"column:mimeType;comment:Mime类型" json:"mimeType"`      // Mime类型
	Order             uint32 `gorm:"column:order;comment:排序" json:"order"`                // 排序
	Type              string `gorm:"column:type;comment:类型" json:"type"`                  // 类型
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`      // 状态
	IsFinished        uint32 `gorm:"column:isFinished;comment:是否已完成上传" json:"isFinished"` // 是否已完成上传
	IsPublic          uint32 `gorm:"column:isPublic;comment:是否可以公开访问" json:"isPublic"`    // 是否可以公开访问
}

// TableName  File's table name
func (*File) TableName() string {
	return TableNameGalloFile
}

// AfterCreate run after create database record.
func (u *File) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "file-")).Error
}

type FileList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*File `json:"items"`
}
