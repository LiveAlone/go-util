package yapi

import (
	"github.com/LiveAlone/go-util/domain/template/bo"
)

type ApiClient struct{}

func NewApiClient() *ApiClient {
	return &ApiClient{}
}

func (y *ApiClient) QueryHttpProjectInfo(token string, apiIdList string) (*bo.HttpProject, error) {
	yapiProjectInfo, err := y.QueryYapiProjectInfo(token, apiIdList)
	if err != nil {
		return nil, err
	}
	return DetailToBasicModel(yapiProjectInfo), nil
}
