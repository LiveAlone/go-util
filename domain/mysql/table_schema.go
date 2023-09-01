package mysql

import "gorm.io/gorm"

type TableInfo struct {
	Schema *TableSchemaInfo
	Column []*TableInfoColumn
	Index  []*TableStatistics
}

type TableInfoColumn struct {
	ColumnName    string `gorm:"column:COLUMN_NAME" json:"column_name"`
	DataType      string `gorm:"column:DATA_TYPE" json:"data_type"`
	ColumnKey     string `gorm:"column:COLUMN_KEY" json:"column_key"`
	IsNullable    string `gorm:"column:IS_NULLABLE" json:"is_nullable"`
	ColumnType    string `gorm:"column:COLUMN_TYPE" json:"column_type"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT" json:"column_comment"`
}

type TableSchemaInfo struct {
	TableSchema  string `gorm:"column:TABLE_SCHEMA"`
	TableName    string `gorm:"column:TABLE_NAME"`
	TableComment string `gorm:"column:TABLE_COMMENT"`
}

// TableStatistics 数据表索引结构
type TableStatistics struct {
	TableSchema  string `gorm:"column:TABLE_SCHEMA"`
	TableName    string `gorm:"column:TABLE_NAME"`
	NoUnique     bool   `gorm:"column:NON_UNIQUE"`
	IndexName    string `gorm:"column:INDEX_NAME"`
	SeqInIndex   int    `gorm:"column:SEQ_IN_INDEX"`
	ColumnName   string `gorm:"column:COLUMN_NAME"`
	IndexComment string `gorm:"column:INDEX_COMMENT"`
}

type TableSchemaAnalyser struct {
	DbUrl  string
	Client *DBClient
}

func NewTableSchemaAnalyser(url string) (*TableSchemaAnalyser, error) {
	client, err := NewDBClient(url)
	if err != nil {
		return nil, err
	}
	return &TableSchemaAnalyser{
		DbUrl:  url,
		Client: client,
	}, nil
}

func (entity *TableSchemaAnalyser) TableInfo(databaseName, table string) (*TableInfo, error) {
	schemaInfo, err := entity.QueryTable(databaseName, table)
	if err != nil {
		return nil, err
	}
	columns, err := entity.QueryColumns(databaseName, table)
	if err != nil {
		return nil, err
	}
	indexList, err := entity.QueryStatistics(databaseName, table)
	if err != nil {
		return nil, err
	}
	return &TableInfo{
		Schema: schemaInfo,
		Column: columns,
		Index:  indexList,
	}, nil
}

// QueryTable 查询数据表基础信息
func (entity *TableSchemaAnalyser) QueryTable(databaseName, table string) (tb *TableSchemaInfo, err error) {
	err = entity.Client.Query("TABLES", &tb, func(db *gorm.DB) *gorm.DB {
		return db.Where("TABLE_SCHEMA = ? and TABLE_NAME = ?", databaseName, table)
	})
	if err != nil {
		return nil, err
	}
	return tb, nil
}

func (entity *TableSchemaAnalyser) QueryColumns(databaseName, table string) (rs []*TableInfoColumn, err error) {
	err = entity.Client.Query("COLUMNS", &rs, func(db *gorm.DB) *gorm.DB {
		return db.Where("TABLE_SCHEMA = ? and TABLE_NAME = ? order by ORDINAL_POSITION", databaseName, table)
	})
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (entity *TableSchemaAnalyser) QueryStatistics(databaseName, table string) (rs []*TableStatistics, err error) {
	err = entity.Client.Query("STATISTICS", &rs, func(db *gorm.DB) *gorm.DB {
		return db.Where("TABLE_SCHEMA = ? and TABLE_NAME = ?", databaseName, table)
	})
	if err != nil {
		return nil, err
	}
	return rs, nil
}
