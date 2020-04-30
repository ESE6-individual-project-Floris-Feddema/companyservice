package repositories

import (
	. "companyservice/contexts"
	. "companyservice/models"
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CompanyRepository struct{}


func (repository CompanyRepository) Create(company *Company) (*Company, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := GetCollection(ctx);


	result, err := collection.InsertOne(ctx, company.Model())
	if err != nil {
		return nil, err
	}

	var returnValue Company
	err = collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&returnValue)

	return &returnValue, nil
}

func (repository CompanyRepository) FindAll() ([]*Company, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := GetCollection(ctx);

	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer result.Close(ctx)

	var returnValues []*Company

	for result.Next(ctx) {
		var elem Company
		err := result.Decode(&elem)
		if err != nil {
			return nil, err
		}
		returnValues = append(returnValues, &elem)
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return returnValues, nil
}

func (repository CompanyRepository) FindOne(id primitive.ObjectID) (*Company, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := GetCollection(ctx)

	var returnValue Company
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&returnValue)
	if err != nil {
		return nil, err
	}

	return &returnValue, nil
}

func (repository CompanyRepository) Delete(id primitive.ObjectID) error {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := GetCollection(ctx)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (repository CompanyRepository) Update(id primitive.ObjectID, company Company) (*Company, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := GetCollection(ctx)

	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, company)
	if err != nil {
		 return nil, err
	}

	returnValue, err := repository.FindOne(result.UpsertedID.(primitive.ObjectID))
	if err != nil {
		return nil, err
	}

	return returnValue, err
}
