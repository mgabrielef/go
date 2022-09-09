package models

import "GitHub/loja/db"

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Stock             int
}

func SearchProducts() []Product {
	db := db.DbConnection()

	selectProducts, err := db.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, stock int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &stock)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Stock = stock

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, stock int) {
	db := db.DbConnection()

	insertDB, err := db.Prepare("insert into products(name, description, price, stock) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertDB.Exec(name, description, price, stock)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DbConnection()

	delProduct, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	delProduct.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.DbConnection()

	dbProduct, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productUpdate := Product{}

	for dbProduct.Next() {
		var id, stock int
		var name, description string
		var price float64

		err = dbProduct.Scan(&id, &name, &description, &price, &stock)
		if err != nil {
			panic(err.Error())
		}

		productUpdate.Id = id
		productUpdate.Name = name
		productUpdate.Description = description
		productUpdate.Price = price
		productUpdate.Stock = stock
	}

	defer db.Close()
	return productUpdate
}

func UpdateProduct(id int, name, description string, price float64, stock int) {
	db := db.DbConnection()

	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, stock=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateProduct.Exec(name, description, price, stock, id)

	defer db.Close()
}
