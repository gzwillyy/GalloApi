package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameIPLibraryFile = "galloIPLibraryFiles"

// IPLibraryFile IP库上传的文件
type IPLibraryFile struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	FileID            uint32 `gorm:"column:fileId;comment:原始文件ID" json:"fileId"`                    // 原始文件ID
	Template          string `gorm:"column:template;comment:模板" json:"template"`                    // 模板
	EmptyValues       string `gorm:"column:emptyValues;comment:空值列表" json:"emptyValues"`            // 空值列表
	GeneratedFileID   uint32 `gorm:"column:generatedFileId;comment:生成的文件ID" json:"generatedFileId"` // 生成的文件ID
	GeneratedAt       uint32 `gorm:"column:generatedAt;comment:生成时间" json:"generatedAt"`            // 生成时间
	IsFinished        uint32 `gorm:"column:isFinished;comment:是否已经完成" json:"isFinished"`            // 是否已经完成
	Countries         string `gorm:"column:countries;comment:国家/地区" json:"countries"`               // 国家/地区
	Provinces         string `gorm:"column:provinces;comment:省份" json:"provinces"`                  // 省份
	Cities            string `gorm:"column:cities;comment:城市" json:"cities"`                        // 城市
	Towns             string `gorm:"column:towns;comment:区县" json:"towns"`                          // 区县
	Providers         string `gorm:"column:providers;comment:ISP服务商" json:"providers"`              // ISP服务商
	Code              string `gorm:"column:code;comment:文件代号" json:"code"`                          // 文件代号
	Password          string `gorm:"column:password;comment:密码" json:"password"`                    // 密码
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                // 状态
}

// TableName IPLibraryFile's table name
func (*IPLibraryFile) TableName() string {
	return TableNameIPLibraryFile
}

// AfterCreate run after create database record.
func (u *IPLibraryFile) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "file-")).Error
}

type IPLibraryFileList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*IPLibraryFile `json:"items"`
}

var IPLibraryFileTableZeroFields = []string{"name", "emptyValues", "countries", "provinces", "cities", "towns", "providers", "code", "password", "state"}

