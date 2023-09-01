package bo

type HttpProject struct {
	ID       int
	Title    string // 项目标题
	Name     string // 项目英文名称
	BasePath string // 基础路径
	ApiList  []*HttpApi
}

type HttpApi struct {
	Schema      string
	Path        string
	Method      string
	Prefix      string // 接口标识
	Description string
	ReqBodyType string
	ResBodyType string
	ReqBodyDesc *BodyDesc
	ResBodyDesc *BodyDesc
}

// BodyDesc 结构体描述信息
type BodyDesc struct {
	Name       string // 字段名称
	Type       string // 基础类型, 结构体类型, todo swagger format
	Example    string
	Desc       string
	Required   bool
	Array      bool
	Properties []*BodyDesc
}

// DtoStructDesc 字段结构体描述信息
type DtoStructDesc struct {
	Name         string // 字段名称
	Example      string
	Desc         string
	DtoFieldDesc []*DtoFieldDesc
}

type DtoFieldDesc struct {
	Name     string // 字段名称
	Type     string // 基础类型, 结构体类型, todo swagger format
	Example  string
	Desc     string
	Required bool
	Array    bool
}
