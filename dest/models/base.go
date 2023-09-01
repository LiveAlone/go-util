package models

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"time"
)

const (
	// proxy 支持的hints
	HintsReadWrite = "/*#mode=READWRITE*/"
	HintsReadOnly  = "/*#mode=READONLY*/"
)

// 删除状态字段
const (
	DeletedNo  = iota //未删除
	DeletedYes        //已删除
)

type CrudModel struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

type NormalPage struct {
	No      int    // 当前第几页
	Size    int    // 每页大小
	OrderBy string `json:"orderBy"` // 排序规则
}

type Option struct {
	IsNeedCnt  bool `json:"isNeedCnt"`
	IsNeedPage bool `json:"isNeedPage"`
}

// 传统分页示例
func NormalPaginate(page *NormalPage) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageNo := 1
		if page.No > 0 {
			pageNo = page.No
		}

		pageSize := page.Size
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (pageNo - 1) * pageSize
		orderBy := "id asc"
		if len(page.OrderBy) > 0 {
			orderBy = page.OrderBy
		}
		return db.Order(orderBy).Offset(offset).Limit(pageSize)
	}
}

// 瀑布流分页示例
type ScrollPage struct {
	Start int // 当前页开始标示
	Size  int // 每页大小
}

func ScrollingPaginate(page *ScrollPage) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		start := -1
		if page.Start > 0 {
			start = page.Start
		}

		pageSize := page.Size
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		return db.Where("id > ?", start).Order("id asc").Limit(pageSize)
	}
}

//type ScopeFunc func(db *gorm.DB) *gorm.DB

type WhereScopes struct {
	Scopes []func(db *gorm.DB) *gorm.DB
}

func (wc *WhereScopes) Where(query interface{}, args ...interface{}) *WhereScopes {
	f := func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
	wc.Scopes = append(wc.Scopes, f)
	return wc
}

func (wc *WhereScopes) Scope(f func(db *gorm.DB) *gorm.DB) *WhereScopes {
	wc.Scopes = append(wc.Scopes, f)
	return wc
}

func (wc *WhereScopes) Page(normalPage *NormalPage) *WhereScopes {
	if normalPage == nil {
		return wc
	}
	// 分页
	if normalPage.No < 1 {
		normalPage.No = 1
	}
	if normalPage.Size < 1 || normalPage.Size > 100 {
		normalPage.Size = 20
	}
	wc.Scopes = append(wc.Scopes, func(db *gorm.DB) *gorm.DB {
		return db.Offset(normalPage.Size * (normalPage.No - 1)).Limit(normalPage.Size)
	})
	return wc
}

func (wc *WhereScopes) Order(orderBy string) *WhereScopes {
	if len(orderBy) == 0 {
		return wc
	}
	// 排序
	wc.Scopes = append(wc.Scopes, func(db *gorm.DB) *gorm.DB {
		return db.Order(orderBy)
	})
	return wc
}

func (wc *WhereScopes) Group(groupBy string) *WhereScopes {
	if len(groupBy) == 0 {
		return wc
	}
	// 分组
	wc.Scopes = append(wc.Scopes, func(db *gorm.DB) *gorm.DB {
		return db.Group(groupBy)
	})
	return wc
}

func (wc *WhereScopes) Preload(query string, args ...interface{}) *WhereScopes {
	wc.Scopes = append(wc.Scopes, func(db *gorm.DB) *gorm.DB {
		db = db.Preload(query, args...)
		return db
	})
	return wc

}

func (wc *WhereScopes) Unscoped() *WhereScopes {
	wc.Scopes = append(wc.Scopes, func(db *gorm.DB) *gorm.DB {
		db = db.Unscoped()
		return db
	})
	return wc
}

// WithNotDeleted 增加"未删除"状态过滤
func WithNotDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("deleted = ? ", DeletedNo)
}

func WithDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("deleted = ? ", DeletedYes)
}

var opMap = map[string]map[reflect.Kind]bool{
	"=": {
		reflect.Int8:   true,
		reflect.Int:    true,
		reflect.Int64:  true,
		reflect.String: true,
	},
	">": {
		reflect.Int8:  true,
		reflect.Int:   true,
		reflect.Int64: true,
	},
	"<": {
		reflect.Int8:  true,
		reflect.Int:   true,
		reflect.Int64: true,
	},
	">=": {
		reflect.Int8:  true,
		reflect.Int:   true,
		reflect.Int64: true,
	},
	"<=": {
		reflect.Int8:  true,
		reflect.Int:   true,
		reflect.Int64: true,
	},
	"in": {
		reflect.Array: true,
		reflect.Slice: true,
	},
	"not in": {
		reflect.Array: true,
		reflect.Slice: true,
	},
	"between": {
		reflect.Array: true,
		reflect.Slice: true,
	},
	"like": {
		reflect.String: true,
	},
	"<>": {
		reflect.Int8:   true,
		reflect.Int:    true,
		reflect.Int64:  true,
		reflect.String: true,
	},
	"udf": { //用户自定义函数
		reflect.String: true,
	},
}

// 不定条件结构体 操作类型常量
type CondItem struct {
	Index  string
	OpType string
	Data   interface{}
}

// 参数校验
func (cond CondItem) Check() (err error) {
	opType := cond.OpType
	if _, ok := opMap[opType]; !ok {
		err = errors.New(cond.Index + " 不支持的操作类型 : " + opType)
		return
	}

	if _, ok := opMap[opType][reflect.TypeOf(cond.Data).Kind()]; !ok {
		err = errors.New(cond.Index + " 条件数据于操作类型不匹配 : " + reflect.TypeOf(cond.Data).Kind().String())
		return
	}
	return
}

// BuildWhere 条件处理
func BuildWhere(db *gorm.DB, conds []CondItem) (outDb *gorm.DB, err error) {
	outDb = db
	if len(conds) == 0 {
		return
	}

	// 检查参数
	for _, v := range conds {
		err = v.Check()
		if err != nil {
			return
		}
	}

	// 处理条件
	for _, v := range conds {
		if v.OpType == "udf" {
			outDb = outDb.Where(strings.Replace(v.Data.(string), "'", "\\'", -1))
		} else {
			outDb = outDb.Where(v.Index+" "+v.OpType+" ?", v.Data)
		}

	}
	return
}
