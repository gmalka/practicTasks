package application

import (
	modelsjson "echohandlertest/application/modelsJSON"
	"echohandlertest/domain/customer"
	"echohandlertest/domain/product"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service interface {
	Order(customer uuid.UUID, products []uuid.UUID) error
	NewCustomer(customer customer.Customer) error
	NewProduct(product product.Product) error
	AllProducts() ([]product.Product, error)
	AllCustomers() ([]customer.Customer, error)
}

func NewHandler(s Service) Handler {
	return Handler{s: s}
}

type Handler struct {
	s Service
}

func (h Handler) InitRouter() http.Handler {
	r := gin.Default()
	gin.Default()

	personGroup := r.Group("/customers")
	personGroup.GET("/", func(ctx *gin.Context) {
		customers, err := h.s.AllCustomers()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}
		c := make([]modelsjson.Person, len(customers))
		for i, v := range customers {
			c[i] = modelsjson.Person{
				ID: v.GetID(),
				Name: v.GetName(),
				Email: v.GetEmail(),
			}
		}

		ctx.JSON(http.StatusOK, c)
	})

	personGroup.POST("/", func(ctx *gin.Context) {
		var person modelsjson.Person

		ctx.ShouldBindJSON(&person)
		customer, err := customer.NewCustomer(person.Name, person.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}
		h.s.NewCustomer(customer)
	})

	productgroup := r.Group("/products")
	productgroup.GET("/", func(ctx *gin.Context) {
		products, err := h.s.AllProducts()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}
		c := make([]modelsjson.Product, len(products))
		for i, v := range products {
			c[i] = modelsjson.Product{
				Price: v.GetPrice(),
				Name: v.GetItem().Name,
				Description: v.GetItem().Description,
			}
		}

		ctx.JSON(http.StatusOK, c)
	})

	productgroup.POST("/", func(ctx *gin.Context) {
		var p modelsjson.Product

		ctx.ShouldBindJSON(&p)
		pr, err := product.NewProduct(p.Name, p.Description, p.Price)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}
		h.s.NewProduct(pr)
	})


	return r
}
