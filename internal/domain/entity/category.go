package entity

// Category 分类实体
type Category struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// NewCategory 创建新分类
func NewCategory(name, description string) *Category {
	return &Category{
		Name:        name,
		Description: description,
	}
}
