package yapi

import (
	"github.com/LiveAlone/go-util/domain/http"
	"log"
	"strconv"
	"strings"
)

// QueryYapiProjectInfo 查询Yapi 基础信息
func (y *ApiClient) QueryYapiProjectInfo(token string, apiIdList string) (*ProjectDetailInfo, error) {
	projectBaseInfo := getProjectInfo(token)

	var apiIds []string
	if len(apiIdList) > 0 {
		apiIds = strings.Split(strings.TrimSpace(apiIdList), ",")
	} else {
		page, size := 1, 20
		for {
			pageApiInfo := pageQueryApiInfo(token, projectBaseInfo.ID, page, size)
			if len(pageApiInfo.List) == 0 {
				break
			}
			for _, info := range pageApiInfo.List {
				apiIds = append(apiIds, strconv.FormatInt(info.Id, 10))
			}
			page += 1
		}
	}

	apiList := make([]*ApiInfo, 0, len(apiIds))
	for _, apiId := range apiIds {
		interfaceApiInfo := getInterfaceApi(token, apiId)
		apiList = append(apiList, interfaceApiInfo)
	}

	if len(apiList) == 0 {
		return nil, nil
	}

	return &ProjectDetailInfo{
		ProjectInfo: projectBaseInfo,
		ApiList:     apiList,
	}, nil
}

func getInterfaceApi(token, apiId string) *ApiInfo {
	apiInfo := new(ApiInfo)
	err := http.NewWrapHttp(http.BaseWrap).GetRequest("https://yapi.zuoyebang.cc/api/interface/get", map[string]string{
		"token": token,
		"id":    apiId,
	}, apiInfo)
	if err != nil {
		log.Fatalf("single api info err:%v", err)
	}
	return apiInfo
}

func getProjectInfo(token string) *ProjectInfo {
	projectBaseInfo := new(ProjectInfo)
	err := http.NewWrapHttp(http.BaseWrap).GetRequest("https://yapi.zuoyebang.cc/api/project/get", map[string]string{
		"token": token,
	}, projectBaseInfo)
	if err != nil {
		log.Fatalf("gain basic project info error, casue:%v", err)
	}
	return projectBaseInfo
}

func pageQueryApiInfo(token string, projectId, page, size int) *PageApiInfo {
	pageApiInfo := new(PageApiInfo)
	err := http.NewWrapHttp(http.BaseWrap).GetRequest("https://yapi.zuoyebang.cc/api/interface/list", map[string]string{
		"token":      token,
		"project_id": strconv.Itoa(projectId),
		"page":       strconv.Itoa(page),
		"size":       strconv.Itoa(size),
	}, pageApiInfo)
	if err != nil {
		log.Fatalf("page api info err:%v", err)
	}
	return pageApiInfo
}
