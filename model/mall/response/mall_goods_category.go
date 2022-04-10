package response

//首页分类数据VO(第三级)
type ThirdLevelCategoryRes struct {
	CategoryId    int    `json:"categoryId"`
	CategoryLevel int    `json:"categoryLevel"`
	CategoryName  string `json:"categoryName" `
}

type SecondLevelCategoryRes struct {
	CategoryId            int                     `json:"categoryId"`
	ParentId              int                     `json:"parentId"`
	CategoryLevel         int                     `json:"categoryLevel"`
	CategoryName          string                  `json:"categoryName" `
	ThirdLevelCategoryVOS []ThirdLevelCategoryRes `json:"thirdLevelCategoryVOS"`
}

type FirstLevelCategoryRes struct {
	CategoryId             int                      `json:"categoryId"`
	ParentId               int                      `json:"parentId"`
	CategoryLevel          int                      `json:"categoryLevel"`
	CategoryName           string                   `json:"categoryName" `
	SecondLevelCategoryVOS []SecondLevelCategoryRes `json:"secondLevelCategoryVOS"`
}
