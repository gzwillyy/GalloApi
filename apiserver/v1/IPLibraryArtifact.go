package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameIPLibraryArtifact = "galloIPLibraryArtifacts"

// IPLibraryArtifact IP库制品
type IPLibraryArtifact struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	FileID            uint32 `gorm:"column:fileId;comment:文件ID" json:"fileId"`                  // 文件ID
	LibraryFileID     uint32 `gorm:"column:libraryFileId;comment:IP库文件ID" json:"libraryFileId"` // IP库文件ID
	Meta              string `gorm:"column:meta;comment:元数据" json:"meta"`                       // 元数据
	IsPublic          uint32 `gorm:"column:isPublic;comment:是否为公用" json:"isPublic"`             // 是否为公用
	Code              string `gorm:"column:code;comment:代号" json:"code"`                        // 代号
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`            // 状态
}

// TableName IPLibraryArtifact's table name
func (*IPLibraryArtifact) TableName() string {
	return TableNameIPLibraryArtifact
}

// AfterCreate run after create database record.
func (u *IPLibraryArtifact) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "library-")).Error
}

type IPLibraryArtifactList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*IPLibraryArtifact `json:"items"`
}

var IPLibraryArtifactTableZeroFields = []string{"name", "code", "state"}

