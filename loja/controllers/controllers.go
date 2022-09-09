package controllers

import (
	"GitHub/loja/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		stock := r.FormValue("stock")

		priceConv, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Price error:", err)
		}

		stockConv, err := strconv.Atoi(stock)
		if err != nil {
			log.Println("Stock error:", err)
		}
		models.CreateNewProduct(name, description, priceConv, stockConv)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditProduct(idProduct)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		stock := r.FormValue("stock")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("ERROR:", err)
		}

		priceConv, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("ERROR:", err)
		}

		stockConv, err := strconv.Atoi(stock)
		if err != nil {
			log.Println("ERROR:", err)
		}

		models.UpdateProduct(idConv, name, description, priceConv, stockConv)
	}
	http.Redirect(w, r, "/", 301)
}
