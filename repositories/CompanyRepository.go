package repositories

import (
	. "companyservice/models"
	. "companyservice/repositories/contexts"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CompanyRepository struct{}

func (repository CompanyRepository) Create(company *Company) (*Company, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := GetCollection(ctx)

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
	collection := GetCollection(ctx)

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

	var oldCompany Company
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&oldCompany)
	if err != nil {
		return nil, err
	}
	company.Users = oldCompany.Users

	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", company}})
	if err != nil {
		return nil, err
	}

	returnValue, err := repository.FindOne(company.ID)
	if err != nil {
		return nil, err
	}

	return returnValue, err
}

func (repository CompanyRepository) FindAllUser(id string) ([]*Company, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := GetCollection(ctx)

	filter :=  bson.D{{"$or",
		bson.A{
			bson.D{{"users.userId", id}},
			bson.D{{"owner.userId", id}},
		},
	}}

	result, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

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

func (repository CompanyRepository) AddUser(id primitive.ObjectID, user User) error {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := GetCollection(ctx)

	company, err := repository.FindOne(id)
	if err != nil {
		return err
	}

	company.Users = append(company.Users, user)
	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", company}})
	return err
}

func (repository CompanyRepository) UpdateUser(user User) error {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := GetCollection(ctx)

	_, err := collection.UpdateMany(ctx,
		bson.M{"owner.userId": user.UserId},
		bson.D{{"$set", bson.M{"owner.name": user.Name}}})
	if err != nil {
		return err
	}

	_, err = collection.UpdateMany(ctx,
		bson.M{"users.userId": user.UserId},
		bson.D{{"$set", bson.M{"users.name": user.Name}}})
	if err != nil {
		return err
	}
	return nil
}
