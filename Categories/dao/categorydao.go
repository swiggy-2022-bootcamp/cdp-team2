package dao

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/db"
)

type CategoryDao struct {
	db *dynamodb.DynamoDB
}

func GetCategoryDao() *CategoryDao {
	return &CategoryDao{
		db.GetInstance(),
	}
}

func (cd *CategoryDao) GetByID(id int) (*models.Category, error) {

	cat := models.Category{
		CategoryId: id,
	}

	res, err := cd.db.GetItem(&dynamodb.GetItemInput{
		TableName: cat.GetTableName(),
		Key:       cat.GetKeyFilter(),
	})

	if err != nil {
		log.Printf("Error Fetching Category: %s", err)
		return nil, err
	}

	if err = dynamodbattribute.UnmarshalMap(res.Item, &cat); err != nil {
		log.Printf("Error unMarshalling Category: %s", err)
		return nil, err
	}

	return &cat, nil
}

func (cd *CategoryDao) GetAll() ([]models.Category, error) {

	var results []models.Category

	res, err := cd.db.Scan(&dynamodb.ScanInput{
		TableName: models.Category{}.GetTableName(),
	})

	if err != nil {
		log.Printf("Error Fetching Categories %s", err.Error())
		return nil, err
	}

	if res.Items == nil {
		return results, fmt.Errorf("no categories found")
	}

	err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &results)
	if err != nil {
		log.Printf("Failed to unmarshal categories, %v", err)
		return nil, err
	}

	return results, nil
}

func (cd *CategoryDao) Create(categories []models.Category) ([]models.Category, []error) {

	var results []models.Category
	var errorlist []error

	for _, c := range categories {
		log.Printf("new catergory %+v", c)

		av, err := dynamodbattribute.MarshalMap(c)
		if err != nil {
			log.Printf("Error marshalling new caterogry item: %s", err)
			errorlist = append(errorlist, err)
			continue
		}

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: c.GetTableName(),
		}

		resp, err := cd.db.PutItem(input)
		if err != nil {
			log.Printf("Error calling PutItem: %s", err)
			errorlist = append(errorlist, err)
			continue
		}

		log.Printf("PutItem resp %v", resp)

		savedCategory, err := cd.GetByID(c.CategoryId)

		if err != nil {
			log.Printf("Error Fetching Category: %s", err)
			errorlist = append(errorlist, err)
			continue
		}

		results = append(results, (*savedCategory))
	}

	return results, errorlist
}

func (cd *CategoryDao) UpdateByID(id int, cat models.Category) (*models.Category, error) {

	toUpd, err := getCategoryUpdExp(sanitizeUpdInput(cat))

	if err != nil {
		log.Printf("error while creating expression %s", err)
		return nil, err
	}

	resp, err := cd.db.UpdateItem(&dynamodb.UpdateItemInput{
		TableName:                 cat.GetTableName(),
		Key:                       cat.GetKeyFilter(),
		ExpressionAttributeNames:  toUpd.Names(),
		ExpressionAttributeValues: toUpd.Values(),
		UpdateExpression:          toUpd.Update(),
	})

	if err != nil {
		log.Printf("error while updating cateogories %s", err)
		return nil, err
	}

	log.Printf("category update resp %v", resp)

	return cd.GetByID(id)
}

func (cd *CategoryDao) DeleteByID(ids ...int) []error {
	var errorlist []error
	for _, id := range ids {

		toDel := models.Category{CategoryId: id}

		resp, err := cd.db.DeleteItem(&dynamodb.DeleteItemInput{
			TableName: toDel.GetTableName(),
			Key:       toDel.GetKeyFilter(),
		})

		log.Printf("delete category resp %+v", resp)

		if err != nil {
			log.Printf("delete category err %s", err.Error())
			errorlist = append(errorlist, err)
		}

	}

	return errorlist
}

//Converts Model.Category to CategoryUpdateInput
func sanitizeUpdInput(c models.Category) models.CategoryUpdateInput {
	return models.CategoryUpdateInput{
		c.CategoryDesc,
	}
}

//getCategoryUpdExp creates a Expression from CategoryUpdateInput
//this iterates through all fields
func getCategoryUpdExp(cat models.CategoryUpdateInput) (expression.Expression, error) {
	// iterate through all fields and set all fields which are not null
	v := reflect.ValueOf(cat)
	t := reflect.TypeOf(&cat).Elem()
	var updateExp expression.UpdateBuilder

	for i := 0; i < v.NumField(); i++ {
		vf := v.Field(i)
		if !vf.IsZero() {
			nameTag := t.Field(i).Tag.Get("json")
			nameTag = strings.Split(nameTag, ",")[0]
			updateExp = updateExp.Set(expression.Name(nameTag), expression.Value(vf.String()))
			fmt.Printf("Field: %s\tValue: %v\n", nameTag, vf.String())
		}
	}

	log.Printf("category update expression %+v", updateExp)
	return expression.NewBuilder().WithUpdate(updateExp).Build()
}
