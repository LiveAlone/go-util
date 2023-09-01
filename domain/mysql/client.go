package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBClient 通用查询
type DBClient struct {
	db *gorm.DB
}

func NewDBClient(dbUrl string) (*DBClient, error) {
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DBClient{
		db: db,
	}, nil
}

func (entity *DBClient) Query(table string, result interface{}, where ...func(db *gorm.DB) *gorm.DB) error {
	err := entity.db.Table(table).Scopes(where...).Find(result).Error
	return err
}
