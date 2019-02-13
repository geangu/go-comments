package service

import (
	"time"

	"../domain"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CommentService class
type CommentService struct{}

var collection *mgo.Collection

func init() {
	collection = initDatabase()
}

// initDatabase start the database and create the collection
func initDatabase() *mgo.Collection {
	session, _ := mgo.Dial("mongo:27017")
	collection := session.DB("pedidosya").C("comments")
	index := mgo.Index{
		Key:    []string{"purchase"},
		Unique: true,
	}
	_ = collection.EnsureIndex(index)
	return collection
}

// Create comment
func (c *CommentService) Create(comment domain.Comment) error {
	err := collection.Insert(comment)
	return err
}

// Delete method
func (c *CommentService) Delete(Purchase string) error {
	var comment = domain.Comment{}
	err := collection.Find(bson.M{"purchase": Purchase, "deleted": false}).One(&comment)
	comment.Deleted = true

	err = collection.Update(bson.M{"purchase": Purchase}, comment)
	return err
}

// FindByPurchase get the comment for the purchase
func (c *CommentService) FindByPurchase(Purchase string) (domain.Comment, error) {
	var comment = domain.Comment{}
	err := collection.Find(bson.M{"purchase": Purchase, "deleted": false}).One(&comment)
	return comment, err
}

// FindByEstablishmentInRange get the comments for the establishment in a range of dates
func (c *CommentService) FindByEstablishmentInRange(Establishment string, Start time.Time, End time.Time) ([]domain.Comment, error) {
	var comments = []domain.Comment{}
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
