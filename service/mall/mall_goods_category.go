package mall

import (
	"github.com/jinzhu/copier"
	"main.go/global"
	"main.go/model/common/enum"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
)

type MallGoodsCategoryService struct {
}

func (m *MallGoodsCategoryService) GetCategoriesForIndex() (err error, newBeeMallIndexCategoryVOS []mallRes.FirstLevelCategoryRes) {
	//获取一级分类的固定数量的数据
	_, firstLevelCategories := selectByLevelAndParentIdsAndNumber([]int{0}, enum.LevelOne.Code(), 10)
	if firstLevelCategories != nil {
		var firstLevelIds []int
		for _, firstLevelCategory := range firstLevelCategories {
			firstLevelIds = append(firstLevelIds, firstLevelCategory.CategoryId)
		}
		//获取二级分类的数据
		_, secondLevelCategories := selectByLevelAndParentIdsAndNumber(firstLevelIds, enum.LevelTwo.Code(), 0)
		if secondLevelCategories != nil {
			var secondLevelIds []int
			for _, secondLevelCategory := range secondLevelCategories {
				secondLevelIds = append(secondLevelIds, secondLevelCategory.CategoryId)
			}
			//获取三级分类的数据
			_, thirdLevelCategories := selectByLevelAndParentIdsAndNumber(secondLevelIds, enum.LevelThree.Code(), 0)
			if thirdLevelCategories != nil {
				thirdLevelCategoryMap := make(map[int]manage.MallGoodsCategory)
				//根据 parentId 将 thirdLevelCategories 分组
				for _, thirdLevelCategory := range thirdLevelCategories {
					thirdLevelCategoryMap[thirdLevelCategory.ParentId] = thirdLevelCategory
				}
				var secondLevelCategoryVOS []mallRes.SecondLevelCategoryRes
				//处理二级分类
				for _, secondLevelGoodsCategory := range secondLevelCategories {
					var secondLevelCategoryVO mallRes.SecondLevelCategoryRes
					err = copier.Copy(&secondLevelCategoryVO, &secondLevelGoodsCategory)
					//如果该二级分类下有数据则放入 secondLevelCategoryVOS 对象中
					if _, ok := thirdLevelCategoryMap[secondLevelGoodsCategory.CategoryId]; ok {
						//根据二级分类的id取出thirdLevelCategoryMap分组中的三级分类list
						tempGoodsCategories := thirdLevelCategoryMap[secondLevelGoodsCategory.CategoryId]
						var thirdLevelCategoryRes mallRes.ThirdLevelCategoryRes
						err = copier.Copy(&thirdLevelCategoryRes, &tempGoodsCategories)
						secondLevelCategoryVO.ThirdLevelCategoryVOS = append(secondLevelCategoryVO.ThirdLevelCategoryVOS, thirdLevelCategoryRes)
						secondLevelCategoryVOS = append(secondLevelCategoryVOS, secondLevelCategoryVO)
					}
				}
				//处理一级分类
				if secondLevelCategoryVOS != nil {
					secondLevelCategoryMap := make(map[int]manage.MallGoodsCategory)
					for _, secondLevelCategory := range secondLevelCategories {
						secondLevelCategoryMap[secondLevelCategory.ParentId] = secondLevelCategory
					}
					for _, firstLevelGoodsCategory := range firstLevelCategories {
						var firstLevelCategoryVO mallRes.FirstLevelCategoryRes
						err = copier.Copy(&firstLevelCategoryVO, &firstLevelGoodsCategory)
						//如果该二级分类下有数据则放入 secondLevelCategoryVOS 对象中
						if _, ok := secondLevelCategoryMap[firstLevelGoodsCategory.CategoryId]; ok {
							//根据二级分类的id取出thirdLevelCategoryMap分组中的三级分类list
							tempGoodsCategories := secondLevelCategoryMap[firstLevelGoodsCategory.CategoryId]
							var secondLevelCategoryRes mallRes.SecondLevelCategoryRes
							err = copier.Copy(&secondLevelCategoryRes, &tempGoodsCategories)
							firstLevelCategoryVO.SecondLevelCategoryVOS = append(firstLevelCategoryVO.SecondLevelCategoryVOS, secondLevelCategoryRes)
							newBeeMallIndexCategoryVOS = append(newBeeMallIndexCategoryVOS, firstLevelCategoryVO)
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
