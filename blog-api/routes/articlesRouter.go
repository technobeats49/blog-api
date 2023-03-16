package routes

import (
	"example.com/blog-api/controllers"
	"example.com/blog-api/db"
	"github.com/julienschmidt/httprouter"
)

type Routes struct {
	router *httprouter.Router
}

// return pointer to httprouter.Router object
func New() *Routes {
	routes := &Routes{router: httprouter.New()}
	return routes
}

// return router
func (r *Routes) GetHttpRouter() *httprouter.Router {
	return r.router
}

// Routes handler calling controller methods
func (r *Routes) ArticlesRoutes() {
	router := r.GetHttpRouter()
	bsc := controllers.NewBlogStoreController(db.DBS)
	router.GET("/articles", bsc.GetAllArticles)
	router.GET("/articles/:id", bsc.GetArticle)
	router.POST("/articles", bsc.CreateArticle)
}
