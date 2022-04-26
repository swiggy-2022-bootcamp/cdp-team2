package models

type CategoryDesc struct {
	Name            string `json:"name,omitempty" extensions:"x-omitempty"`
	Description     string `json:"description,omitempty"  extensions:"x-omitempty"`
	MetaDescription string `json:"meta_description,omitempty"  extensions:"x-omitempty"`
	MetaKeyword     string `json:"meta_keyword,omitempty"  extensions:"x-omitempty"`
	MetaTitle       string `json:"meta_title,omitempty"  extensions:"x-omitempty"`
}

// @Description Category Model
type Category struct {
	CategoryId   int          `json:"category_id,omitempty"  extensions:"x-omitempty"`
	CategoryDesc CategoryDesc `json:"category_description,omitempty"  extensions:"x-omitempty"`
}

type CategoryUpdateInput struct {
	CategoryDesc CategoryDesc `json:"category_description"`
}
