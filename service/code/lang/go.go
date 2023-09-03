package lang

import (
	"errors"
	"github.com/LiveAlone/go-util/domain/mysql"
)

type GoCodeGenerator struct {
}

func NewGoCodeGenerator() *GoCodeGenerator {
	return &GoCodeGenerator{}
}

func (g *GoCodeGenerator) GenDaoFromTableInfo(tableInfo *mysql.TableInfo, params *CodeGenParams) (rs map[string]string, err error) {
	return nil, errors.New("not support")
}

// GenerateGoModel 生成go model
func GenerateGoModel() {
	//var sqlModelConfig SqlModelConfig
	//err := configLoader.LoadConfigToEntity(fmt.Sprintf("%s/%s", targetPath, "db.yaml"), &sqlModelConfig)
	//if err != nil {
	//	log.Fatalf("db.yaml load error, err :%v", err)
	//}
	//
	//// 数据表生成
	//db := sqlModelConfig.Db
	//tbs := strings.Split(db.Tables, ",")
	//
	//tableCode, dataCode, err := gen.GenDao(db.Url, db.DataBase, tbs)
	//if err != nil {
	//	log.Fatalf("db model generate error, err :%v", err)
	//}
	//
	//for tableName, code := range tableCode {
	//	fileName := domain.ToSnakeLower(strings.TrimPrefix(tableName, "tbl"))
	//
	//	// model
	//	dir := fmt.Sprintf("%s/models", targetPath)
	//	err := util.CreateDirIfNotExists(dir)
	//	if err != nil {
	//		log.Fatalf("create dir error, err :%v", err)
	//	}
	//	err = util.WriteFile(fmt.Sprintf("%s/%s.go", dir, fileName), []byte(code))
	//	if err != nil {
	//		log.Fatalf("tb file write error, err :%v", err)
	//	}
	//	fmt.Println("数据表Model 生成完成: ", tableName)
	//}
	//
	//for tableName, code := range dataCode {
	//
	//	fileName := domain.ToSnakeLower(strings.TrimPrefix(tableName, "tbl"))
	//	dataDir := fmt.Sprintf("%s/data", targetPath)
	//	err = util.CreateDirIfNotExists(dataDir)
	//	if err != nil {
	//		log.Fatalf("create dir error, err :%v", err)
	//	}
	//
	//	err = util.WriteFile(fmt.Sprintf("%s/%s.go", dataDir, fileName), []byte(code))
	//	if err != nil {
	//		log.Fatalf("tb file write error, err :%v", err)
	//	}
	//	fmt.Println("数据表Data 生成完成: ", tableName)
	//}
}
