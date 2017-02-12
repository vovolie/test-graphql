package data

import "fmt"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Material struct {
	ID           string `json:"id"`
	CategoryInfo string `json:"category"`
	Name         string `json:"name"`
	Cover        string `json:"cover"`
	URL          string `json:"url"`
}

var Music = &Category{"1", "Music"}
var Sticker = &Category{"2", "Sticker"}

var Material1 = &Material{"1", "1", "小幸运", "http://p3.music.126.net/OyIYO_0pq5OLkFXb2GL2bg==/109951162850203027.jpg?param=140y140", "1.mp3"}
var Material2 = &Material{"2", "1", "陪你度过漫长岁月", "http://p3.music.126.net/GjYzJf08Xgc44_Ayznr-cA==/18681802069319982.jpg?param=140y140", "2.mp3"}
var Material3 = &Material{"3", "1", "一生有你", "http://p3.music.126.net/aNwTq7wceDZojb6vg6ngqw==/109951162845378742.jpg?param=140y140", "3.mp3"}
var Material4 = &Material{"4", "2", "夜照亮了夜", "http://p3.music.126.net/tY4h3obkIPDZcHYSYDVsKA==/18543263604213571.jpg?param=140y140", "4.mp3"}

var categoryById = map[string]*Category{
	"1": Music,
	"2": Sticker,
}

var materialsById = map[string]*Material{
	"1": Material1,
	"2": Material2,
	"3": Material3,
	"4": Material4,
}
var materialIdsByCategory = map[string][]string{
	"1": {"1", "2", "3"},
	"2": {"4"},
}

var nextMaterialId = 4

func AddMaterial(categoryInfo, name, cover, url string) string {
	material := &Material{
		ID:           fmt.Sprintf("%v", nextMaterialId),
		CategoryInfo: categoryInfo,
		Name:         name,
		Cover:        cover,
		URL:          url,
	}
	nextMaterialId = nextMaterialId + 1

	materialsById[material.ID] = material
	materialIdsByCategory[material.CategoryInfo] = append(materialIdsByCategory[material.CategoryInfo], material.ID)

	return material.ID
}

func GetMaterial(id string) *Material {
	if material, ok := materialsById[id]; ok {
		return material
	}
	return nil
}

func GetMaterials(categoryId string) []*Material {
	materials := []*Material{}
	if categoryId == "any" {
		for id := range categoryById {
			for _, materialId := range materialIdsByCategory[id] {
				if material := GetMaterial(materialId); material != nil {
					materials = append(materials, material)
				}
			}
		}
	} else {
		for _, materialId := range materialIdsByCategory[categoryId] {
			if material := GetMaterial(materialId); material != nil {
				materials = append(materials, material)
			}
		}
	}

	return materials
}

func GetCategory(id string) *Category {
	if category, ok := categoryById[id]; ok {
		return category
	}
	return nil
}

func GetViewer(id string) *Category {
	return GetCategory(id)
}

func ChangeMaterialCategory(id string, categoryInfo string) bool {
	material := GetMaterial(id)
	if material == nil {
		return false
	}
	material.CategoryInfo = categoryInfo
	updatedMaterialIdsForCategory := []string{}
	for _, materialId := range materialIdsByCategory[material.CategoryInfo] {
		updatedMaterialIdsForCategory = append(updatedMaterialIdsForCategory, materialId)
	}
	materialIdsByCategory[material.CategoryInfo] = updatedMaterialIdsForCategory
	return true
}

func ChangeMaterialUrl(id, url string) {
	material := GetMaterial(id)
	if material == nil {
		return
	}
	material.URL = url
}

func RenameMaterial(id, name string) {
	material := GetMaterial(id)
	if material != nil {
		material.Name = name
	}
}

func RemoveMaterial(id string) {
	material := GetMaterial(id)
	updatedMaterialIdsForCategory := []string{}
	for _, materialId := range materialIdsByCategory[material.CategoryInfo] {
		if materialId != id {
			updatedMaterialIdsForCategory = append(updatedMaterialIdsForCategory, materialId)
		}
	}
	materialIdsByCategory[material.CategoryInfo] = updatedMaterialIdsForCategory
	delete(materialsById, id)
}

func MaterialsToSliceInterface(materials []*Material) []interface{} {
	materialsIFace := []interface{}{}
	for _, material := range materials {
		materialsIFace = append(materialsIFace, material)
	}
	return materialsIFace
}
