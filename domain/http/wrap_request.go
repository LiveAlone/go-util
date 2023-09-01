package http

import (
	"errors"
	"github.com/LiveAlone/go-util/util"
	jsoniter "github.com/json-iterator/go"
)

type WrapDecoder interface {
	data([]byte) ([]byte, error)
}

type WrapHttp struct {
	WrapDecoder WrapDecoder
}

func NewWrapHttp(wrap func() WrapDecoder) *WrapHttp {
	return &WrapHttp{WrapDecoder: wrap()}
}

func NoneWrap() WrapDecoder {
	return &NoneWrapDecoder{}
}

func BaseWrap() WrapDecoder {
	return &BaseWrapDecoder{}
}

// GetRequest get 请求
func (w *WrapHttp) GetRequest(url string, params map[string]string, result any) error {
	data, err := util.Get(url, params)
	if err != nil {
		return err
	}
	entity, err := w.WrapDecoder.data(data)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(entity, result)
}

// NoneWrapDecoder 无封装返回结果
type NoneWrapDecoder struct{}

func (n *NoneWrapDecoder) data(source []byte) ([]byte, error) {
	return source, nil
}

// BaseWrapDecoder 基础封装
type BaseWrapDecoder struct{}

func (b *BaseWrapDecoder) data(source []byte) ([]byte, error) {
	var baseResponse BasicResponse
	err := jsoniter.Unmarshal(source, &baseResponse)
	if err != nil {
		return nil, err
	}

	if baseResponse.ErrCode != 0 {
		return nil, errors.New("httpCall" + baseResponse.ErrMsg)
	}
	return baseResponse.Data, nil
}

type BasicResponse struct {
	ErrCode int                 `json:"errcode"`
	ErrMsg  string              `json:"errmsg"`
	Data    jsoniter.RawMessage `json:"data"`
}
