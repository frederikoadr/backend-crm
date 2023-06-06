package account

import (
	"BackendCRM/dto"
	"BackendCRM/utility"
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

func (h RequestHandler) Create(c *gin.Context) {
	var reqActor dto.RequestActor

	if err := c.BindJSON(&reqActor); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.ctrl.Create(&reqActor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Read(c *gin.Context) {
	res, err := h.ctrl.Read()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
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
	var user dto.RequestActor
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

func (h RequestHandler) Login(c *gin.Context) {
	uname, pas, ok := c.Request.BasicAuth()
	if !ok {
		// Basic auth credentials are not provided or invalid
		c.JSON(http.StatusUnauthorized, "Otentikasi gagal")
		return
	}
	customer, err := h.ctrl.ReadBy("username", uname)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Username tidak terdaftar")
		return
	}
	if generateHash(pas) != customer.Password {
		c.JSON(http.StatusUnauthorized, "Password tidak sesuai")
		return
	}
	jwt, err := utility.GenerateJWT(customer.RoleId, "koentji")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, jwt)
}
