package customers

import (
	"BackendCRM/dto"
	"BackendCRM/utility"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
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

func (h RequestHandler) Create(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
		return
	}
	token = strings.TrimPrefix(token, "Bearer ")
	resJWT, err := utility.VerfiyJWT(token, "koentji")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
		return
	}
	roleId, err := strconv.Atoi(resJWT)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid token"})
		return
	}
	if roleId > 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak diijinkan melakukan aksi ini"})
		return
	}

	var req dto.RequestCustomer

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Read(c *gin.Context) {
	token := c.GetHeader("Authorization")
	rId := 0
	if token != "" {
		token = strings.TrimPrefix(token, "Bearer ")
		resJWT, err := utility.VerfiyJWT(token, "koentji")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}
		roleId, err := strconv.Atoi(resJWT)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid token"})
			return
		}
		rId = roleId
	}

	res, err := h.ctrl.Read()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if res.Data == nil && rId == 2 { //ketika data kosong dan role is admin read akan menambah dari API
		url := "https://reqres.in/api/users?page=2"

		// GET request ke API
		res1, err1 := http.Get(url)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err1.Error()})
			return
		}
		defer res1.Body.Close()

		var apiRequest dto.APIRequest
		if err := json.NewDecoder(res1.Body).Decode(&apiRequest); err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}
		for _, data := range apiRequest.Data {
			cust := dto.RequestCustomer{
				FirstName: data.FirstName,
				LastName:  data.LastName,
				Email:     data.Email,
				Avatar:    data.Avatar,
			}
			h.ctrl.Create(&cust)
		}
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) ReadBy(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
		return
	}
	token = strings.TrimPrefix(token, "Bearer ")
	resJWT, err := utility.VerfiyJWT(token, "koentji")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
		return
	}
	roleId, err := strconv.Atoi(resJWT)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid token"})
		return
	}
	if roleId > 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak diijinkan melakukan aksi ini"})
		return
	}

	column := c.Param("column")
	value := c.Query("value")

	var customer *dto.CustomerItemResponse

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
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
		return
	}
	token = strings.TrimPrefix(token, "Bearer ")
	resJWT, err := utility.VerfiyJWT(token, "koentji")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
		return
	}
	roleId, err := strconv.Atoi(resJWT)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid token"})
		return
	}
	if roleId > 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak diijinkan melakukan aksi ini"})
		return
	}
	userID := c.Param("id")
	res, err := h.ctrl.Delete(userID)
	if err != nil {
		return
	}
	// Tampilkan respons berhasil
	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Update(c *gin.Context) {
	var user dto.RequestCustomer
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
