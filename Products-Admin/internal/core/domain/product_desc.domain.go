package domain

type ProductDescription struct {
	LanguageID      int    `json:"language_id"      bson:"language_id"`
	Name            string `json:"name"             bson:"name"`
	Description     string `json:"description"      bson:"description"`
	MetaTitle       string `json:"meta_title"       bson:"meta_title"`
	MetaDescription string `json:"meta_description" bson:"meta_description"`
	MetaKeyword     string `json:"meta_keyword"     bson:"meta_keyword"`
	Tag             string `json:"tag"              bson:"tag"`
}
