package bo

// ModelStruct template/model 模型
type ModelStruct struct {
	PackageName string        // 包名
	TableName   string        // tblModelTable
	BeanName    string        // ModelTable
	Columns     []*ModelField // 字段列表
	Comment     string        // 数据表注释
}

// ModelField 实体对象类型
type ModelField struct {
	// db 属性
	ColumnName string // db 字段名称
	Nullable   bool   // 是否可空
	IsPrimary  bool   // 是否主键

	// Field 填充属性
	FieldName string // 结构体字段名称
	FieldType string // 结构体数据类型
	Comment   string // 字段评论
}

// DaoStruct template/dao 基于数据查询对象
type DaoStruct struct {
	PackageName string      // 包名
	TableName   string      // tblModelTable
	BeanName    string      // ModelTable
	Comment     string      // 数据表注释
	IndexList   []*DaoIndex // 索引位置
}

// DaoIndex dao 包含索引
type DaoIndex struct {
	IndexName    string
	Unique       bool
	Fields       []*DaoIndexField
	IndexComment string
}

// DaoIndexField dao 索引包含字段
type DaoIndexField struct {
	Index      int    // 位置
	ColumnName string // db 字段名称
	ColumnType string // db 字段类型

	FieldName   string // 字段名称
	FieldNameFL string // 字段名称首字母大写
	FieldType   string // 结构体数据类型
}
