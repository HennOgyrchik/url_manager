package links

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"gitlab.com/robotomize/gb-golang/homework/03-01-umanager/internal/database"
)

const collection = "links"

func New(db *mongo.Database, timeout time.Duration) *Repository {
	return &Repository{db: db, timeout: timeout}
}

type Repository struct {
	db      *mongo.Database
	timeout time.Duration
}

func (r *Repository) Create(ctx context.Context, req CreateReq) (database.Link, error) {
	var l database.Link

	l.ID = req.ID
	l.Title = req.Title
	l.URL = req.URL
	l.Images = req.Images
	l.Tags = req.Tags
	l.UserID = req.UserID
	l.CreatedAt = time.Now()
	l.UpdatedAt = l.CreatedAt

	_, err := r.db.Collection(collection).InsertOne(ctx, l)
	if err != nil {
		return l, fmt.Errorf("links InsertOne: %w", err)
	}

	return l, nil
}

func (r *Repository) FindByUserAndURL(ctx context.Context, link, userID string) (database.Link, error) {
	var l database.Link

	err := r.db.Collection(collection).FindOne(ctx, bson.M{
		"url":    link,
		"userID": userID,
	}).Decode(&l)

	if err != nil {
		return l, fmt.Errorf("links Decode: %w", err)
	}

	return l, nil
}

func (r *Repository) FindByCriteria(ctx context.Context, criteria Criteria) ([]database.Link, error) {
	return nil, nil
}
