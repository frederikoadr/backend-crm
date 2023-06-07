package account

import (
	"BackendCRM/dto"
	"BackendCRM/function/auth"
	"BackendCRM/utility/token"
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
func (h RequestHandler) CreateReg(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isAdmin := auth.CheckAdmin(auth.TokenParam{Data: token})
	if !isAdmin {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak diijinkan melakukan aksi ini"})
		return
	}

	userID := c.Param("id")

	res, err := h.ctrl.CreateReg(userID)
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

func (h RequestHandler) ReadBy(c *gin.Context) {
	var err error
	token := c.GetHeader("Authorization")
	isAdmin := auth.CheckAdmin(auth.TokenParam{Data: token})
	isSuperAdmin := auth.CheckSuperAdmin(auth.TokenParam{Data: token})
	if !isAdmin && !isSuperAdmin {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak diijinkan melakukan aksi ini"})
		return
	}

	column := c.Param("column")
	value := c.Query("value")

	var customer *dto.ActorItemResponse

	switch column {
	case "id":
		customer, err = h.ctrl.ReadBy("id", value)
	case "username":
		customer, err = h.ctrl.ReadBy("username", value)
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
	isSuperAdmin := auth.CheckSuperAdmin(auth.TokenParam{Data: token})
	if !isSuperAdmin {
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
	token := c.GetHeader("Authorization")
	isSuperAdmin := auth.CheckSuperAdmin(auth.TokenParam{Data: token})
	if !isSuperAdmin {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak diijinkan melakukan aksi ini"})
		return
	}

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

func (h RequestHandler) UpdateReg(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isSuperAdmin := auth.CheckSuperAdmin(auth.TokenParam{Data: token})
	if !isSuperAdmin {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak diijinkan melakukan aksi ini"})
		return
	}

	userID := c.Param("id")
	value := c.Query("status")

	res, err := h.ctrl.UpdateReg(userID, value)
	if err != nil {
		return
	}
	// Tampilkan respons berhasil
	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Login(c *gin.Context) {
	var err error
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
	jwt, err := token.GenerateJWT(customer.RoleId, "koentji")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, jwt)
}

func (h RequestHandler) ReadRegis(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isSuperAdmin := auth.CheckSuperAdmin(auth.TokenParam{Data: token})
	if !isSuperAdmin {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak diijinkan melakukan aksi ini"})
		return
	}

	res, err := h.ctrl.ReadRegis()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
