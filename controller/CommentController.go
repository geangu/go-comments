package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"../domain"
	services "../service"
	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
	mgo "gopkg.in/mgo.v2"
)

// CommentController controller for the comment resource
type CommentController struct{}

var validate *validator.Validate
var service *services.CommentService

func init() {
	validate = validator.New()
	service = new(services.CommentService)
}

// Create comment method
func (c *CommentController) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var comment domain.Comment
	err := decoder.Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()

	comment.Created = time.Now()

	err = validate.Struct(comment)
	if err != nil {
		http.Error(w, "Missing or wrong request data", http.StatusBadRequest)
		return
	}

	err = service.Create(comment)
	if err != nil && mgo.IsDup(err) {
		http.Error(w, "There is already a qualification for this purchase", http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

// Delete comment method
func (c *CommentController) Delete(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id := params["id"]

	err := service.Delete(id)
	if err != nil {
		http.Error(w, "Unable to remove the resource", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// FindByPurchase get comment by purchase
func (c *CommentController) FindByPurchase(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	purchase, _ := params["id"]

	var comment, err = service.FindByPurchase(purchase)
	if err != nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)
}

// FindByEstablishment get comments by establishment
func (c *CommentController) FindByEstablishment(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	establishment := params["id"]
	format := time.RFC3339

	startDate, err := time.Parse(format, r.FormValue("start"))
	if err != nil {
		http.Error(w, "The 'start' parameter is not in the expected date format RFC3339", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse(format, r.FormValue("end"))
	if err != nil {
		http.Error(w, "The 'end' parameter is not in the expected date format RFC3339", http.StatusBadRequest)
		return
	}

	comments, _ := service.FindByEstablishmentInRange(establishment, startDate, endDate)
	if len(comments) == 0 {
		http.Error(w, "No resources found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}
