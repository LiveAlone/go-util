package yapi

type ProjectDetailInfo struct {
	ProjectInfo *ProjectInfo
	ApiList     []*ApiInfo
}

type ProjectInfo struct {
	ID       int    `json:"_id"`
	Name     string `json:"name"`     // brick
	Basepath string `json:"basepath"` // /brick
}

type PageApiInfo struct {
	Count int        `json:"count"`
	Total int        `json:"total"`
	List  []*ApiInfo `json:"list"`
}

type ApiInfo struct {
	Id           int64           `json:"_id"`
	Method       string          `json:"method"`
	Path         string          `json:"path"`
	Title        string          `json:"title"`
	ReqQueryList []*ReqQueryItem `json:"req_query"`     // GET
	ReqBodyType  string          `json:"req_body_type"` // POST
	ReqBodyOther string          `json:"req_body_other"`
	ResBodyType  string          `json:"res_body_type"` // POST
	ResBody      string          `json:"res_body"`
}

type ReqQueryItem struct {
	Id       string `json:"_id"`
	Name     string `json:"name"`
	Example  string `json:"example"`
	Desc     string `json:"desc"`
	Required string `json:"required"`
}
