package dao

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/db"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/internal/literals"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

const (
	customEpoch = 1300000000
	tabelName   = "team-2-orders"
)

type OrderDao struct {
	db dynamodbiface.DynamoDBAPI
}

func GetOrderDao() IDao {
	return &OrderDao{
		db.GetInstance(),
	}
}

func ordersTableName() *string {
	return aws.String(tabelName)
}

func getKeyFilter(id int) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"orderId": {
			N: aws.String(fmt.Sprint(id)),
		},
	}
}

func getRandomKey() int {
	return int(time.Now().Unix() - customEpoch)
}

func (cd *OrderDao) GetByID(id int) (*models.Order, error) {

	fmt.Println("getById", id, getKeyFilter(id))

	res, err := cd.db.GetItem(&dynamodb.GetItemInput{
		TableName: ordersTableName(),
		Key:       getKeyFilter(id),
	})

	log.Printf("res-id ", id, res)

	if err != nil {
		log.Printf("Error Fetching Category: %s", err)
		return nil, err
	}

	cat := models.Order{}

	if res.Item == nil {
		return nil, errors.New(literals.CategoryNotFound)
	}

	if err = dynamodbattribute.UnmarshalMap(res.Item, &cat); err != nil {
		log.Printf("Error unMarshalling Category: %s", err)
		return nil, err
	}

	return &cat, nil
}

//func (cd *CategoryDao) GetAll() ([]models.Category, error) {
//
//	var results []models.Category
//
//	res, err := cd.db.Scan(&dynamodb.ScanInput{
//		TableName: categoriesTableName(),
//	})
//
//	if err != nil {
//		log.Printf("Error Fetching Categories %s", err.Error())
//		return nil, err
//	}
//
//	if res.Items == nil {
//		return results, errors.New(literals.NoCategoriesFound)
//	}
//
//	err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &results)
//	if err != nil {
//		log.Printf("Failed to unmarshal categories, %v", err)
//		return nil, err
//	}
//
//	return results, nil
//}
//
func (cd *OrderDao) Create(order models.Order) (*models.Order, error) {

	newId := getRandomKey()

	log.Printf("new order %+v", order)

	var exp expression.UpdateBuilder
	exp = exp.Set(expression.Name("productDesc"), expression.Value(order.ProductDesc))
	exp = exp.Set(expression.Name("status"), expression.Value(order.Status))
	exp = exp.Set(expression.Name("customerId"), expression.Value(order.CustomerId))
	exp = exp.Set(expression.Name("totalPrice"), expression.Value(order.TotalPrice))

	express, _ := expression.NewBuilder().WithUpdate(exp).Build()

	createItemIn := dynamodb.UpdateItemInput{
		TableName:                 ordersTableName(),
		Key:                       getKeyFilter(newId),
		ExpressionAttributeNames:  express.Names(),
		ExpressionAttributeValues: express.Values(),
		UpdateExpression:          express.Update(),
		ReturnValues:              aws.String("ALL_NEW"),
	}

	log.Printf("create item input : %+v", createItemIn)

	resp, err := cd.db.UpdateItem(&createItemIn)

	if err != nil {
		log.Printf("error while creating cateogory %s", err.Error())
		return nil, err
	}

	log.Printf("category create resp %v", resp)

	savedCat := models.Order{}
	if err = dynamodbattribute.UnmarshalMap(resp.Attributes, &savedCat); err != nil {
		log.Printf("error while creating cateogory %s", err.Error())
		return nil, err
	}

	return &savedCat, nil
}

//func (cd *CategoryDao) UpdateByID(id int, cat models.Category) (*models.Category, error) {
//
//	toUpd, err := getCategoryUpdExp(cat)
//
//	if err != nil {
//		log.Printf("error while creating expression %s", err)
//		return nil, err
//	}
//
//	updItemIn := dynamodb.UpdateItemInput{
//		TableName:                 categoriesTableName(),
//		Key:                       getKeyFilter(id),
//		ExpressionAttributeNames:  toUpd.Names(),
//		ExpressionAttributeValues: toUpd.Values(),
//		UpdateExpression:          toUpd.Update(),
//		ReturnValues:              aws.String("ALL_NEW"),
//		ConditionExpression:       toUpd.Condition(),
//	}
//
//	log.Printf("update item input : %+v", updItemIn)
//
//	resp, err := cd.db.UpdateItem(&updItemIn)
//
//	if err != nil {
//		log.Printf("error while updating cateogories %s", err.Error())
//		return nil, err
//	}
//
//	log.Printf("category update resp %v", resp)
//
//	updatedAttributes := models.Category{}
//	if err = dynamodbattribute.UnmarshalMap(resp.Attributes, &updatedAttributes); err != nil {
//		log.Printf("error while updating cateogories %s", err.Error())
//		return nil, err
//	}
//
//	return &updatedAttributes, nil
//}
//
//func (cd *CategoryDao) DeleteByID(id int) error {
//
//	resp, err := cd.db.DeleteItem(&dynamodb.DeleteItemInput{
//		TableName: categoriesTableName(),
//		Key:       getKeyFilter(id),
//	})
//
//	log.Printf("delete category resp %+v", resp)
//
//	if err != nil {
//		log.Printf("delete category err %s", err.Error())
//		return err
//	}
//	return nil
//}
//
////getCategoryUpdExp creates a Expression from CategoryUpdateInput
////this iterates through all fields
//func getCategoryUpdExp(cat models.Category) (expression.Expression, error) {
//	var updateExp expression.UpdateBuilder
//
//	desc := cat.CategoryDesc
//	// descPrefix := "category_description."
//	//category description
//	if desc.Name != "" {
//		updateExp = updateExp.Set(expression.Name("category_description.name"), expression.Value(desc.Name))
//	}
//
//	if desc.Description != "" {
//		updateExp = updateExp.Set(expression.Name("category_description.description"), expression.Value(desc.Description))
//	}
//
//	if desc.MetaDescription != "" {
//		updateExp = updateExp.Set(expression.Name("category_description.meta_description"), expression.Value(desc.MetaDescription))
//	}
//
//	if desc.MetaKeyword != "" {
//		updateExp = updateExp.Set(expression.Name("category_description.meta_keyword"), expression.Value(desc.MetaKeyword))
//	}
//
//	if desc.MetaTitle != "" {
//		updateExp = updateExp.Set(expression.Name("category_description.meta_title"), expression.Value(desc.MetaTitle))
//	}
//
//	log.Printf("category update expression %+v", updateExp)
//	exp, err := expression.
//		NewBuilder().
//		WithCondition(expression.AttributeExists(expression.Name("category_id"))).
//		WithUpdate(updateExp).
//		Build()
//
//	// log.Printf("%+v", exp.Names())
//	// log.Println(exp.Values())
//	// log.Println(*exp.Update())
//	return exp, err
//}
