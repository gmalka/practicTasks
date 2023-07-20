package application

import (
	modelsjson "echohandlertest/application/modelsJSON"
	"echohandlertest/domain/customer"
	"echohandlertest/domain/product"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
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
	r := echo.New()

	r.Use(middleware.Logger())
	personGroup := r.Group("/customers")
	personGroup.GET("/", func(c echo.Context) error {
		customers, err := h.s.AllCustomers()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		custs := make([]modelsjson.Person, len(customers))
		for i, v := range customers {
			custs[i] = modelsjson.Person{
				ID:    v.GetID(),
				Name:  v.GetName(),
				Email: v.GetEmail(),
			}
		}

		c.JSON(http.StatusOK, custs)
		return nil
	})

	personGroup.POST("/", func(c echo.Context) error {
		var person modelsjson.Person

		c.Bind(&person)
		customer, err := customer.NewCustomer(person.Name, person.Email)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		h.s.NewCustomer(customer)
		return nil
	})

	productgroup := r.Group("/products")
	productgroup.GET("/", func(c echo.Context) error {
		products, err := h.s.AllProducts()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		prods := make([]modelsjson.Product, len(products))
		for i, v := range products {
			prods[i] = modelsjson.Product{
				Price:       v.GetPrice(),
				Name:        v.GetItem().Name,
				Description: v.GetItem().Description,
			}
		}

		c.JSON(http.StatusOK, prods)
		return nil
	})

	productgroup.POST("/", func(c echo.Context) error {
		var p modelsjson.Product

		c.Bind(&p)
		pr, err := product.NewProduct(p.Name, p.Description, p.Price)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		h.s.NewProduct(pr)
		return nil
	})

	return r
}
