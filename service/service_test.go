package service

import (
	"testing"
	"time"

	"../domain"
)

func TestCreate(t *testing.T) {
	comment := domain.Comment{
		Text:          "service-test",
		User:          "service-test",
		Establishment: "service-test",
		Purchase:      "service-test",
		Score:         1,
		Created:       time.Now(),
		Deleted:       false,
	}

	var service = new(CommentService)

	err := service.Create(comment)
	if err != nil {
		t.Error("Error was not expected")
	}
}

func TestFindByPurchase(t *testing.T) {
	var service = new(CommentService)
	comment, err := service.FindByPurchase("service-test")
	if err != nil {
		t.Error("Error was not expected")
	}

	if comment == (domain.Comment{}) {
		t.Error("An empty object was not expected")
	}
}

func TestFindByEstablishmentInRange(t *testing.T) {
	var service = new(CommentService)

	format := time.RFC3339
	start, err := time.Parse(format, "2019-01-01T00:00:00-05:00")
	end := time.Now()

	comments, err := service.FindByEstablishmentInRange("service-test", start, end)
	if err != nil {
		t.Error("Error was not expected")
	}

	if len(comments) == 0 {
		t.Error("Not expected empty array")
	}
}

func TestDelete(t *testing.T) {
	var service = new(CommentService)
	err := service.Delete("service-test")
	if err != nil {
		t.Error("Error was not expected")
	}
}
