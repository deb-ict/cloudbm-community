package mongodb

import (
	"context"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/module/user"
	"github.com/deb-ict/cloudbm-community/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	db         database
	collection *mongo.Collection
}

type userDocument struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	UserName string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
	Email    string             `bson:"email"`
}

func (d *userDocument) toViewModel() user.User {
	return user.User{
		Id:       d.Id.Hex(),
		UserName: d.UserName,
		Password: d.Password,
		Email:    d.Email,
	}
}

func (d *userDocument) fromViewModel(model user.User) error {
	var err error
	d.Id, err = primitive.ObjectIDFromHex(model.Id)
	d.UserName = strings.ToLower(model.UserName)
	d.Password = model.Password
	d.Email = strings.ToLower(model.Email)
	return err
}

func (repo *userRepository) GetUsers(ctx context.Context, pageIndex int, pageSize int) (*user.UserPage, error) {
	pageIndex, pageSize = repo.db.getNormalizedPaging(pageIndex, pageSize)

	collection := repo.getCollection()
	filter := bson.M{}
	findOptions := options.Find()

	// Get the total number of items
	totalItems, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	findOptions.SetSkip((int64(pageIndex) - 1) * int64(pageSize))
	findOptions.SetLimit(int64(pageSize))
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var data []user.User
	for cursor.Next(ctx) {
		var document userDocument
		cursor.Decode(&document)
		data = append(data, document.toViewModel())
	}

	return &user.UserPage{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Count:     int(totalItems),
		Data:      data,
	}, nil
}

func (repo *userRepository) GetUserById(ctx context.Context, id string) (*user.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := repo.getCollection()
	filter := bson.M{"_id": objectId}
	result := collection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var document userDocument
	err = result.Decode(&document)
	if err != nil {
		return nil, err
	}

	data := document.toViewModel()
	return &data, nil
}

func (repo *userRepository) GetUserByUserName(ctx context.Context, userName string) (*user.User, error) {
	userName = strings.ToLower(userName)

	collection := repo.getCollection()
	filter := bson.M{"username": userName}
	result := collection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var document userDocument
	err := result.Decode(&document)
	if err != nil {
		return nil, err
	}

	data := document.toViewModel()
	return &data, nil
}

func (repo *userRepository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	email = strings.ToLower(email)

	collection := repo.getCollection()
	filter := bson.M{"email": email}
	result := collection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var document userDocument
	err := result.Decode(&document)
	if err != nil {
		return nil, err
	}

	data := document.toViewModel()
	return &data, nil
}

func (repo *userRepository) CreateUser(ctx context.Context, model user.User) (string, error) {
	collection := repo.getCollection()

	document := &userDocument{}
	document.fromViewModel(model)

	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}

	newid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", storage.ErrDbNotCreated
	}
	return newid.Hex(), nil
}

func (repo *userRepository) UpdateUser(ctx context.Context, id string, model user.User) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	document := &userDocument{}
	document.fromViewModel(model)
	document.Id = objectId

	collection := repo.getCollection()
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": document}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return storage.ErrDbNotUpdated
	}
	return nil
}

func (repo *userRepository) DeleteUser(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := repo.getCollection()
	filter := bson.M{"_id": objectId}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return storage.ErrDbNotDeleted
	}

	return nil
}

func (repo *userRepository) getCollection() *mongo.Collection {
	if repo.collection == nil {
		repo.collection = repo.db.database.Collection("user")
	}
	return repo.collection
}
