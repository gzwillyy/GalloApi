// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloPost = "galloPosts"

// GalloPost 文章管理
type GalloPost struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	CategoryID  int32  `gorm:"column:categoryId;comment:文章分类" json:"categoryId"`             // 文章分类
	Type        string `gorm:"column:type;comment:类型：normal, url" json:"type"`               // 类型：normal, url
	URL         string `gorm:"column:url;comment:URL" json:"url"`                            // URL
	Subject     string `gorm:"column:subject;comment:标题" json:"subject"`                     // 标题
	Body        string `gorm:"column:body;comment:内容" json:"body"`                           // 内容
	CreatedAt   int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`               // 创建时间
	IsPublished bool   `gorm:"column:isPublished;comment:是否已发布" json:"isPublished"`          // 是否已发布
	PublishedAt int64  `gorm:"column:publishedAt;comment:发布时间" json:"publishedAt"`           // 发布时间
	ProductCode string `gorm:"column:productCode;comment:产品代号" json:"productCode"`           // 产品代号
	State       bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
}

// TableName GalloPost's table name
func (*GalloPost) TableName() string {
	return TableNameGalloPost
}