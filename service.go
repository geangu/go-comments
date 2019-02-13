package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// CommentService class
type CommentService struct{}

// Create comment
func (c *CommentService) Create(comment Comment) (Comment, error) {
	err := collection.Insert(comment)
	return comment, err
}

// Delete method
func (c *CommentService) Delete(Purchase string) error {
	var comment = Comment{}
	err := collection.Find(bson.M{"purchase": Purchase, "deleted": false}).One(&comment)
	comment.Deleted = true

	err = collection.Update(bson.M{"purchase": Purchase}, comment)
	return err
}

// FindByPurchase get the comment for the purchase
func (c *CommentService) FindByPurchase(Purchase string) (Comment, error) {
	var comment = Comment{}
	err := collection.Find(bson.M{"purchase": Purchase, "deleted": false}).One(&comment)
	return comment, err
}

// FindByEstablishmentInRange get the comments for the establishment in a range of dates
func (c *CommentService) FindByEstablishmentInRange(Establishment string, Start time.Time, End time.Time) ([]Comment, error) {
	var comments = []Comment{}
	err := collection.Find(
		bson.M{
			"establishment": Establishment,
			"created": bson.M{
				"$gt": Start,
				"$lt": End,
			},
			"deleted": false,
		},
	).Sort("-created").All(&comments)
	return comments, err
}
