package models

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	categoriesTableName = "categories"
)

type CategoryDesc struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	MetaDescription string `json:"meta_description,omitempty"`
	MetaKeyword     string `json:"meta_keyword,omitempty"`
	MetaTitle       string `json:"meta_title,omitempty"`
}

type Category struct {
	CategoryId   int          `json:"category_id"`
	CategoryDesc CategoryDesc `json:"category_description"`
}

type CategoryUpdateInput struct {
	CategoryDesc CategoryDesc `json:"category_description"`
}

func (c Category) GetTableName() *string {
	return aws.String(categoriesTableName)
}

func (c Category) GetKeyFilter() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"category_id": {
			S: aws.Int(c.CategoryId),
		},
	}
}
