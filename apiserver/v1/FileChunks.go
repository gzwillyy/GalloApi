package v1

import metav1 "github.com/gzwillyy/components/pkg/meta/v1"

const TableNameGalloFileChunk = "galloFileChunks"

// FileChunk 文件片段
type FileChunk struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	FileID            uint64 `gorm:"column:fileId;comment:文件ID" json:"fileId"` // 文件ID
	Data              []byte `gorm:"column:data;comment:分块内容" json:"data"`     // 分块内容
}

// TableName FileChunk oFileChunk's table name
func (*FileChunk) TableName() string {
	return TableNameGalloFileChunk
}

type FileChunkList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*FileChunk `json:"items"`
}
