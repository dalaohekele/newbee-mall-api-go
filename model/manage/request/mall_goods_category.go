package request

import (
	"main.go/model/common"
	"main.go/model/common/request"
	"main.go/model/manage"
)

type MallGoodsCategoryReq struct {
	CategoryLevel int             `json:"categoryLevel" gorm:"comment:分类等级"`
	ParentId      int             `json:"parentId" gorm:"comment:父类id"`
	CategoryName  string          `json:"categoryName" gorm:"comment:分类名称"`
	CategoryRank  string          `json:"categoryRank" gorm:"comment:排序比重"`
	IsDeleted     int             `json:"isDeleted" gorm:"comment:是否删除"`
	CreateTime    common.JSONTime `json:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"` // 创建时间
	UpdateTime    common.JSONTime `json:"updateTime" gorm:"column:update_time;comment:修改时间;type:datetime"` // 更新时间
}

type SearchCategoryParams struct {
	manage.MallGoodsCategory
	request.PageInfo
}
