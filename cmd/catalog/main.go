package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/marcospmail/devfullcycle-imersao17-goapi/internal/database"
	"github.com/marcospmail/devfullcycle-imersao17-goapi/service"
	"github.com/marcospmail/devfullcycle-imersao17-goapi/webserver"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()

	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Route("/categories", func(r chi.Router) {
		r.Get("/", webCategoryHandler.GetCategories)
		r.Get("/{id}", webCategoryHandler.GetCategory)
		r.Post("/", webCategoryHandler.CreateCategory)
	})

	c.Route("/products", func(r chi.Router) {
		r.Get("/", webProductHandler.GetProducts)
		r.Get("/{id}", webProductHandler.GetProduct)
		r.Get("/{categoryID}", webProductHandler.GetProductByCategoryID)
		r.Post("/", webProductHandler.CreateProduct)
	})

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", c)
}
