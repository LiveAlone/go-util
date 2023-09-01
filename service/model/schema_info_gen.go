package model

// todo yqj golang 代码模型生成器
//// SchemaInformationGen 基于SchemaInformation生成代码
//type SchemaInformationGen struct {
//	tempGenerator *template.Generator
//}
//
//func NewSchemaInformationGen(generator *template.Generator) *SchemaInformationGen {
//	return &SchemaInformationGen{
//		tempGenerator: generator,
//	}
//}
//
//func (s *SchemaInformationGen) Gen(url string, db string, tableList []string) (modelCode map[string]string, dataCode map[string]string, err error) {
//	if len(tableList) == 0 {
//		return nil, nil, nil
//	}
//	analyser, err := mysql.NewTableSchemaAnalyser(url)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	modelCode = make(map[string]string)
//	dataCode = make(map[string]string)
//	for _, tableName := range tableList {
//		tableInfo, err := analyser.QueryTable(db, tableName)
//		if err != nil {
//			return nil, nil, err
//		}
//		tableColumns, err := analyser.QueryColumns(db, tableName)
//		if err != nil {
//			return nil, nil, err
//		}
//
//		// 查询数据表信息
//		data, err := s.genModel(tableInfo, tableColumns)
//		if err != nil {
//			return nil, nil, err
//		}
//		modelCode[tableName] = data
//
//		// 构建data层模版
//		indexList, err := analyser.QueryStatistics(db, tableName)
//		if err != nil {
//			return nil, nil, err
//		}
//
//		data, err = s.genData(tableInfo, indexList, tableColumns)
//		if err != nil {
//			return nil, nil, err
//		}
//		dataCode[tableName] = data
//	}
//	return
//}
//
//func (s *SchemaInformationGen) genData(table *mysql.TableInfo, indexList []*mysql.TableStatistics, columns []*mysql.TableInfoColumn) (string, error) {
//	ds, err := buildDataStruct(table, indexList, columns)
//	if err != nil {
//		return "", err
//	}
//	jc, _ := jsoniter.MarshalIndent(ds, "", "  ")
//	fmt.Println("数据结构: ", string(jc))
//	return s.tempGenerator.GenerateTemplateContent(ds, map[string]any{
//		"ToCamelCaseFistLarge": domain.ToCamelCaseFistLarge,
//		"ToCamelCaseFistLower": domain.ToCamelCaseFistLower,
//	})
//}
//
//func buildDataStruct(table *mysql.TableInfo, indexList []*mysql.TableStatistics, columns []*mysql.TableInfoColumn) (*bo.DataStruct, error) {
//	columnDataTypeMap := make(map[string]string)
//	for _, column := range columns {
//		columnDataTypeMap[column.ColumnName] = column.DataType
//	}
//
//	indexMap := make(map[string]*bo.DataIndex)
//	for _, statistics := range indexList {
//		dataIndex, ok := indexMap[statistics.IndexName]
//		if !ok {
//			dataIndex = &bo.DataIndex{
//				IndexName:    statistics.IndexName,
//				Unique:       !statistics.NoUnique,
//				Fields:       make([]*bo.DataIndexField, 0),
//				IndexComment: statistics.IndexComment,
//			}
//			indexMap[statistics.IndexName] = dataIndex
//		}
//
//		fieldType, ok := config.GlobalConf.DbTypeMap[columnDataTypeMap[statistics.ColumnName]]
//		if !ok {
//			return nil, fmt.Errorf("column data type not found %s", statistics.ColumnName)
//		}
//		dataIndex.Fields = append(dataIndex.Fields, &bo.DataIndexField{
//			Index:      statistics.SeqInIndex,
//			ColumnName: statistics.ColumnName,
//			FieldType:  fieldType,
//		})
//	}
//
//	rs := make([]*bo.DataIndex, 0)
//	for _, index := range indexMap {
//		sort.Slice(index.Fields, func(i, j int) bool {
//			return index.Fields[i].Index < index.Fields[j].Index
//		})
//		rs = append(rs, index)
//	}
//	return &bo.DataStruct{
//		BeanName:  domain.ToCamelCaseFistLarge(strings.TrimPrefix(table.TableName, "tbl")),
//		DataIndex: rs,
//	}, nil
//}
//
//func (s *SchemaInformationGen) genModel(table *mysql.TableInfo, columns []*mysql.TableInfoColumn) (string, error) {
//	// 查询数据表信息
//	ms, err := buildModelStruct(table, columns)
//	if err != nil {
//		return "", err
//	}
//	ds, _ := json.MarshalIndent(ms, "", "  ")
//	fmt.Println("模型结构: ", string(ds))
//	return s.tempGenerator.GenerateTemplateContent(ms, map[string]any{
//		"ToCamelCaseFistLarge": domain.ToCamelCaseFistLarge,
//		"ToCamelCaseFistLower": domain.ToCamelCaseFistLower,
//	})
//}
//
//func buildModelStruct(table *mysql.TableInfo, columns []*mysql.TableInfoColumn) (*bo.ModelStruct, error) {
//	// 构建数据转换列表
//	cols := make([]*bo.ModelField, len(columns))
//	for i, column := range columns {
//		fieldType, ok := config.GlobalConf.DbTypeMap[column.DataType]
//		if !ok {
//			log.Fatalf("data type not found, table:%s, type:%s", table, column.DataType)
//		}
//
//		if column.IsNullable == "YES" {
//			toFieldType, ok := config.GlobalConf.GoNullableMap[fieldType]
//			if !ok {
//				log.Fatalf("go nullable type not found, table:%s, go_type:%s nullable tyle:%v", table, column.DataType, toFieldType)
//			}
//			fieldType = toFieldType
//		}
//
//		// json 类型datatypes.JSON 转换
//		if (fieldType == "string" || fieldType == "sql.NullString") && strings.Contains(column.ColumnComment, "json") {
//			fieldType = "datatypes.JSON"
//		}
//
//		cols[i] = &bo.ModelField{
//			ColumnName: column.ColumnName,
//			FieldType:  fieldType,
//			Comment:    column.ColumnComment,
//		}
//	}
//
//	return &bo.ModelStruct{
//		TableName: table.TableName,
//		BeanName:  domain.ToCamelCaseFistLarge(strings.TrimPrefix(table.TableName, "tbl")),
//		Columns:   cols,
//		Comment:   table.TableComment,
//	}, nil
//}
