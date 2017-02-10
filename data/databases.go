package data


import "fmt"


type Category struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type Material struct {
	ID string `json:"id"`
	CategoryInfo string `json:"category"`
	Name string `json:"name"`
	Cover string `json:"cover"`
	URL string `json:"url"`
}

var Music = &Category{"1", "Music"}
var Sticker = &Category{"2", "Sticker"}

var Material1 = &Material{"1", "1", "小幸运", "1.jpg", "1.mp3"}
var Material2 = &Material{"2", "1", "陪你度过漫长岁月", "2.jpg", "2.mp3"}
var Material3 = &Material{"3", "1", "一生有你", "3.jpg", "3.mp3"}
var Material4 = &Material{"4", "2", "夜照亮了夜", "4.jpg", "4.mp3"}



var categoryList = map[string]*Category{
	"1": Music,
	"2": Sticker,
}

var materialList = map[string]*Material{
	"1": Material1,
	"2": Material2,
	"3": Material3,
	"4": Material4,
}

var nextMaterial = 5
var nextCategory = 3

func CreateCategory(categoryName string)  *Category {
	nextCategory = nextCategory + 1
	newCategory := &Category{
		fmt.Sprintf("%v", nextCategory),
		categoryName,
	}
	categoryList[newCategory.ID] = newCategory

	return newCategory
}

func CreateMaterial(categoryInfo, name, cover, url string) *Material {
	nextMaterial = nextMaterial + 1
	newMaterial := &Material{
		fmt.Sprintf("%v", nextMaterial),
		categoryInfo,
		name,
		cover,
		url,
	}
	materialList[newMaterial.ID] = newMaterial
	return newMaterial
}

func GetCategoryById(id string) *Category {
	if category, ok := categoryList[id]; ok {
		return category
	}
	return nil
}

func GetMaterialById(id string) *Material {
	if material, ok := materialList[id]; ok {
		return material
	}
	return nil
}

func GetMaterialByCategory(id string) map[string]*Material {
	materials := map[string]*Material{}
	if category, ok := categoryList[id]; ok {
		for _, material := range materialList {
			if material.CategoryInfo == category.ID {
				materials[material.ID] = material
			}
		}
	}
	return materials
}

func GetCategory() map[string]*Category {
	return categoryList
}

func GetMaterial() map[string]*Material {
	return materialList
}

