package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RequestHandler struct {
	ctrl *Controller
}

func NewRequestHandler(ctrl *Controller) *RequestHandler {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

type CreateRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"  binding:"required"`
	Email     string `json:"email"  binding:"required"`
	Avatar    string `json:"avatar"  binding:"required"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h RequestHandler) Create(c *gin.Context) {
	var req CreateRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Read(c *gin.Context) {
	res, err := h.ctrl.Read()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) ReadBy(c *gin.Context) {
	column := c.Param("column")
	value := c.Query("value")
	fmt.Println(column)
	fmt.Println(value)

	var customer *UserItemResponse
	var err error

	switch column {
	case "firstname":
		customer, err = h.ctrl.ReadBy("first_name", value)
	case "lastname":
		customer, err = h.ctrl.ReadBy("last_name", value)
	case "email":
		customer, err = h.ctrl.ReadBy("email", value)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid column"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if customer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h RequestHandler) Delete(c *gin.Context) {
	userID := c.Param("id")
	res, err := h.ctrl.Delete(userID)
	if err != nil {
		return
	}
	// Tampilkan respons berhasil
	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Update(c *gin.Context) {
	var user Customers
	userID := c.Param("id")

	// Baca data JSON dari body permintaan
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ctrl.Update(&user, userID)
	if err != nil {
		return
	}
	// Tampilkan respons berhasil
	c.JSON(http.StatusOK, res)
}
