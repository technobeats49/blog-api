package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/blog-api/db"
	"example.com/blog-api/models"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/exp/maps"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/* stores mongodb session used by controllers methods*/
type BlogStoreController struct {
	session *mgo.Session
}

func NewBlogStoreController(s *mgo.Session) *BlogStoreController {
	return &BlogStoreController{s}
}

func (b BlogStoreController) GetArticle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// supporting obj for responsebody
	var result = make(map[string]any)

	// Handling hex value but not a valid hex representation of an ObjectId
	if !bson.IsObjectIdHex(id) {
		result["data"] = nil
		result["status"] = http.StatusNotFound
		result["message"] = "Not a valid hex representation of an ObjectId "
		w.WriteHeader(http.StatusNotFound)
		rjson, err := json.Marshal(result)

		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s\n", rjson)
		maps.Clear(result)
		return
	}

	// fetch valid ObjectId
	oid := bson.ObjectIdHex(id)

	// article object to map article data from db
	a := models.Article{}

	if err := b.session.DB(db.DB_NAME).C(db.COLL_NAME).FindId(oid).One(&a); err != nil {
		result["data"] = nil
		result["status"] = http.StatusNotFound
		result["message"] = "Article not found"
		w.WriteHeader(http.StatusNotFound)
		rjson, err := json.Marshal(result)

		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s\n", rjson)
		maps.Clear(result)
		return
	}

	// response data slice of article
	resdata := []models.Article{}

	result["data"] = append(resdata, a)
	result["status"] = http.StatusOK
	result["message"] = "success"

	rjson, err := json.Marshal(result)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", rjson)
	maps.Clear(result)
	return
}

func (b BlogStoreController) CreateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	a := models.Article{}

	var result = make(map[string]any)

	json.NewDecoder(r.Body).Decode(&a)

	a.Id = bson.NewObjectId()

	// check if any trouble inserting data
	if err := b.session.DB(db.DB_NAME).C(db.COLL_NAME).Insert(a); err != nil {
		result["data"] = nil
		result["status"] = http.StatusInternalServerError
		result["message"] = "Error inserting record in DB, check DB healthiness and errors"
		w.WriteHeader(http.StatusInternalServerError)
		rjson, err := json.Marshal(result)

		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(rjson)
	}

	// use resdata a map of string to any for building response body
	resdata := make(map[string]any)
	resdata["id"] = a.Id

	// build response body
	result["data"] = resdata
	result["status"] = http.StatusCreated
	result["message"] = "success"

	rjson, err := json.Marshal(result)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(rjson)
	maps.Clear(result)

}

func (b BlogStoreController) GetAllArticles(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var result = make(map[string]any)
	type Articles []models.Article

	allarticles := Articles{}
	if err := b.session.DB(db.DB_NAME).C(db.COLL_NAME).Find(bson.M{}).All(&allarticles); err != nil {
		result["data"] = nil
		result["status"] = http.StatusNotFound
		result["message"] = "Error fetching records from collection"
		w.WriteHeader(http.StatusNotFound)
		rjson, err := json.Marshal(result)

		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s\n", rjson)
		maps.Clear(result)
		return
	}

	// building response body
	result["data"] = allarticles
	result["status"] = http.StatusOK
	result["message"] = "Success"

	allarticlesj, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", allarticlesj)
	maps.Clear(result)

}
