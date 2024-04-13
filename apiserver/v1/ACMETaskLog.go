package v1

const TableNameACMETaskLog = "galloACMETaskLogs"

// ACMETaskLog ACME任务运行日志
type ACMETaskLog struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	TaskID    int64  `gorm:"column:taskId;comment:任务ID" json:"taskId"`                     // 任务ID
	IsOk      bool   `gorm:"column:isOk;default:1;comment:是否成功" json:"isOk"`               // 是否成功
	Error     string `gorm:"column:error;comment:错误信息" json:"error"`                       // 错误信息
	CreatedAt int64  `gorm:"column:createdAt;comment:运行时间" json:"createdAt"`               // 运行时间
}

// TableName ACMETaskLog's table name
func (*ACMETaskLog) TableName() string {
	return TableNameACMETaskLog
}
