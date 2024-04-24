package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/gzwillyy/components/pkg/meta/v1"
	"github.com/gzwillyy/components/pkg/util/idutil"
)

const TableNameNSQuestionOption = "galloNSQuestionOptions"

// NSQuestionOption DNS请求选项
type NSQuestionOption struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Values            string `gorm:"column:values;comment:选项值" json:"values"` // 选项值
}

// TableName NSQuestionOption's table name
func (*NSQuestionOption) TableName() string {
	return TableNameNSQuestionOption
}

// AfterCreate run after create database record.
func (u *NSQuestionOption) AfterCreate(tx *gorm.DB) error {
	return tx.Model(u).UpdateColumn("instanceID", idutil.GetInstanceID(u.ID, "option-")).Error
}

type NSQuestionOptionList struct {
	metav1.ListMeta `json:",inline"`
	Items []*NSQuestionOption `json:"items"`
}
var NSQuestionOptionTableZeroFields = []string{"name", "values"}

