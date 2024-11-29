package res_format

import (
	"http-server/internal/shared/pagination"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ReponseData struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
	Status    int         `json:"status"`
	Metadata  interface{} `json:"metadata"`
}

func FormatResponse200(c *gin.Context, data interface{}, pagination *pagination.PaginationResponseDto) {
	c.JSON(http.StatusOK, &ReponseData{
		Message:   "Ok",
		Data:      data,
		Timestamp: time.Now(),
		Status:    http.StatusOK,
		Metadata:  pagination})
}

func FormatResponse400(c *gin.Context, validatorError interface{}) {
	c.JSON(http.StatusBadRequest, &ReponseData{
		Message:   "Bad Request",
		Data:      nil,
		Timestamp: time.Now(),
		Status:    http.StatusBadRequest,
		Metadata:  validatorError})
}
func FormatResponse401(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, &ReponseData{
		Message:   "Unauthorized",
		Data:      nil,
		Timestamp: time.Now(),
		Status:    http.StatusUnauthorized,
		Metadata:  nil})
}
func FormatResponse404(c *gin.Context, errCode string) {
	c.JSON(http.StatusNotFound, &ReponseData{
		Message:   errCode,
		Data:      nil,
		Timestamp: time.Now(),
		Status:    http.StatusNotFound,
		Metadata:  nil})
}
func FormatResponse500(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, &ReponseData{
		Message:   "Internal Server Error",
		Data:      nil,
		Timestamp: time.Now(),
		Status:    http.StatusInternalServerError,
		Metadata:  nil})
}
