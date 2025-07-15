package entity

// PostCategory 文章分类关联实体
type PostCategory struct {
	ID         uint `json:"id"`
	PostID     uint `json:"post_id"`
	CategoryID uint `json:"category_id"`
}

// NewPostCategory 创建新文章分类关联
func NewPostCategory(postID, categoryID uint) *PostCategory {
	return &PostCategory{
		PostID:     postID,
		CategoryID: categoryID,
	}
}
