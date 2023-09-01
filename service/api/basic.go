package api

import (
	"github.com/LiveAlone/go-util/domain/template/bo"
	"log"

	"github.com/LiveAlone/go-util/domain"
	"github.com/LiveAlone/go-util/service/api/yapi"
)

// SchemaApiGen 基于APISchema 生成code
type SchemaApiGen struct {
	yapiClient *yapi.ApiClient
}

func NewSchemaGen(yapiClient *yapi.ApiClient) *SchemaApiGen {
	return &SchemaApiGen{
		yapiClient: yapiClient,
	}
}

// GenFromYapi 通过生成code
func (s *SchemaApiGen) GenFromYapi(token string, allApi bool, apiList string) (httpProject *bo.HttpProject, err error) {
	if allApi {
		httpProject, err = s.yapiClient.QueryHttpProjectInfo(token, "")
		if err != nil {
			return nil, err
		}
	} else {
		httpProject, err = s.yapiClient.QueryHttpProjectInfo(token, apiList)
		if err != nil {
			return nil, err
		}
	}

	if httpProject == nil {
		log.Fatalf("fail get project info, token:%v", token)
	}
	return
}

func ConvertProjectApisDtoDesc(apiList []*bo.HttpApi) (rs []*bo.DtoStructDesc) {
	for _, api := range apiList {
		rs = append(rs, convertBodyDescToDtoDesc(api.Prefix, api.ReqBodyDesc)...)
		rs = append(rs, convertBodyDescToDtoDesc(api.Prefix, api.ResBodyDesc)...)
	}
	return
}

func convertBodyDescToDtoDesc(prefix string, desc *bo.BodyDesc) (rs []*bo.DtoStructDesc) {
	if desc.Type != "object" {
		log.Fatalf("dto desc convert error, body:%v", desc)
	}

	fields := make([]*bo.DtoFieldDesc, 0)
	for _, property := range desc.Properties {
		if property.Type == "object" {
			loopDesc := convertBodyDescToDtoDesc(prefix, property)
			rs = append(rs, loopDesc...)
			fields = append(fields, &bo.DtoFieldDesc{
				Name:     domain.ToCamelCaseFistLarge(property.Name),
				Type:     toStructName(prefix, property.Name),
				Example:  property.Example,
				Desc:     property.Desc,
				Required: property.Required,
				Array:    property.Array,
			})
		} else {
			fields = append(fields, &bo.DtoFieldDesc{
				Name:     domain.ToCamelCaseFistLarge(property.Name),
				Type:     toStructType(property.Type),
				Example:  property.Example,
				Desc:     property.Desc,
				Required: property.Required,
				Array:    property.Array,
			})
		}
	}

	rs = append(rs, &bo.DtoStructDesc{
		Name:         toStructName(prefix, desc.Name),
		Example:      desc.Example,
		Desc:         desc.Desc,
		DtoFieldDesc: fields,
	})
	return
}

func toStructName(prefix string, name string) string {
	return prefix + domain.ToCamelCaseFistLarge(name)
}

func toStructType(fromType string) string {
	// todo yqj
	//toType, ok := config.GlobalConf.ApiTypeMap[fromType]
	//if !ok {
	//	log.Fatalf("api from type not found :%v", fromType)
	//}
	//return toType
	return fromType
}
