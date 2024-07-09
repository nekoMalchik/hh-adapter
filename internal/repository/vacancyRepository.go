package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Vacancy представляет вакансию
type Vacancy struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Company     string             `bson:"company"`
	Location    string             `bson:"location"`
	PostedAt    time.Time          `bson:"postedAt"`
}

type VacancyRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewVacancyRepository(uri string, dbName string, collectionName string) (*VacancyRepository, error) {
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create Mongo client: %w", err)
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &VacancyRepository{client: client, collection: collection}, nil
}

func (repo *VacancyRepository) Disconnect(ctx context.Context) error {
	return repo.client.Disconnect(ctx)
}

func (repo *VacancyRepository) Create(ctx context.Context, vacancy *Vacancy) (*mongo.InsertOneResult, error) {
	result, err := repo.collection.InsertOne(ctx, vacancy)
	if err != nil {
		return nil, fmt.Errorf("failed to insert vacancy: %w", err)
	}
	return result, nil
}

func (repo *VacancyRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*Vacancy, error) {
	var vacancy Vacancy
	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&vacancy)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find vacancy by ID: %w", err)
	}
	return &vacancy, nil
}

func (repo *VacancyRepository) Update(ctx context.Context, id primitive.ObjectID, update bson.M) (*mongo.UpdateResult, error) {
	result, err := repo.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	if err != nil {
		return nil, fmt.Errorf("failed to update vacancy: %w", err)
	}
	return result, nil
}

func (repo *VacancyRepository) Delete(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := repo.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, fmt.Errorf("failed to delete vacancy: %w", err)
	}
	return result, nil
}
