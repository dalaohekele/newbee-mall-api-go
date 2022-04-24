package mall

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"main.go/global"
	"main.go/model/common/enum"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
)

type MallGoodsCategoryService struct {
}

func (m *MallGoodsCategoryService) GetCategoriesForIndex() (err error, firstLevelCategoryVOS []mallRes.FirstLevelCategoryRes) {

	//获取一级分类的固定数量的数据
	_, firstLevelCategories := selectByLevelAndParentIdsAndNumber([]int{0}, enum.LevelOne.Code(), 10)
	if firstLevelCategories != nil {
		var firstLevelCategoryIds []int
		for _, firstLevelCategory := range firstLevelCategories {
			firstLevelCategoryIds = append(firstLevelCategoryIds, firstLevelCategory.CategoryId)
		}
		//获取二级分类的数据
		_, secondLevelCategories := selectByLevelAndParentIdsAndNumber(firstLevelCategoryIds, enum.LevelTwo.Code(), 0)
		if secondLevelCategories != nil {
			var secondLevelCategoryIds []int
			for _, secondLevelCategory := range secondLevelCategories {
				secondLevelCategoryIds = append(secondLevelCategoryIds, secondLevelCategory.CategoryId)
			}
			//获取三级分类的数据
			_, thirdLevelCategories := selectByLevelAndParentIdsAndNumber(secondLevelCategoryIds, enum.LevelThree.Code(), 0)
			if thirdLevelCategories != nil {
				//根据 parentId 将 thirdLevelCategories 分组
				var secondLevelCategoryVOS []mallRes.SecondLevelCategoryRes
				thirdLevelCategoryMap := make(map[int][]manage.MallGoodsCategory)
				for _, thirdLevelCategory := range thirdLevelCategories {
					thirdLevelCategoryMap[thirdLevelCategory.ParentId] = []manage.MallGoodsCategory{}
				}
				for k, v := range thirdLevelCategoryMap {
					for _, third := range thirdLevelCategories {
						if k == third.ParentId {
							v = append(v, third)
						}
						thirdLevelCategoryMap[k] = v
					}
				}
				str, _ := json.Marshal(thirdLevelCategoryMap)
				fmt.Println(string(str))
				//处理二级分类
				for _, secondLevelCategory := range secondLevelCategories {
					//var list []mallRes.ThirdLevelCategoryRes
					var secondLevelCategoryVO mallRes.SecondLevelCategoryRes
					err = copier.Copy(&secondLevelCategoryVO, &secondLevelCategory)
					//如果该二级分类下有数据则放入 secondLevelCategoryVOS 对象中
					if _, ok := thirdLevelCategoryMap[secondLevelCategory.CategoryId]; ok {
						//根据二级分类的id取出thirdLevelCategoryMap分组中的三级分类list
						tempGoodsCategories := thirdLevelCategoryMap[secondLevelCategory.CategoryId]
						var thirdLevelCategoryRes []mallRes.ThirdLevelCategoryRes
						err = copier.Copy(&thirdLevelCategoryRes, &tempGoodsCategories)
						secondLevelCategoryVO.ThirdLevelCategoryVOS = thirdLevelCategoryRes
						secondLevelCategoryVOS = append(secondLevelCategoryVOS, secondLevelCategoryVO)
					}

				}
				//处理一级分类
				if secondLevelCategoryVOS != nil {
					//根据 parentId 将 thirdLevelCategories 分组
					secondLevelCategoryVOMap := make(map[int]mallRes.SecondLevelCategoryRes)
					for _, v := range secondLevelCategories {
						var secondLevelCategory mallRes.SecondLevelCategoryRes
						copier.Copy(&secondLevelCategory, &v)
						secondLevelCategoryVOMap[secondLevelCategory.ParentId] = secondLevelCategory
					}
					for _, firstLevelGoodsCategory := range firstLevelCategories {
						var newBeeMallIndexCategoryVO mallRes.FirstLevelCategoryRes
						err = copier.Copy(&newBeeMallIndexCategoryVO, &firstLevelGoodsCategory)
						//如果该一级分类下有数据则放入 newBeeMallIndexCategoryVOS 对象中
						if _, ok := secondLevelCategoryVOMap[firstLevelGoodsCategory.CategoryId]; ok {
							//根据一级分类的id取出secondLevelCategoryVOMap分组中的二级级分类list
							newBeeMallIndexCategoryVO.SecondLevelCategoryVOS = secondLevelCategoryVOS
							firstLevelCategoryVOS = append(firstLevelCategoryVOS, newBeeMallIndexCategoryVO)
						}
					}
				}
			}
		}
	}
	return
}

// 获取分类数据
func selectByLevelAndParentIdsAndNumber(ids []int, level int, limit int) (err error, categories []manage.MallGoodsCategory) {

	global.GVA_DB.Where("parent_id in ?", ids).Where("category_level =? and is_deleted = 0", level).
		Order("category_rank desc").Limit(limit).Find(&categories)
	return
}
