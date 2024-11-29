package tests

import (
	"fmt"
	"http-server/internal/shared/database/entities"
	"http-server/internal/shared/utils"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetUserIdFromContext(t *testing.T) {
	// DB, err := database.NewDatabaseConnection()
	mockUser := entities.Users{}
	mockUser.ID = 123

	user := map[string]interface{}{
		"id": float64(123),
	}
	c := &gin.Context{}
	c.Set("user", user)

	userId, err := utils.GetUserIdFromContext(c)
	if err != nil {
		panic(err)
	}
	if userId != 123 {
		t.Fail()
		return
	}
}
func TestParseUserId(t *testing.T) {

	_, err := utils.GetNumberAsUint64("123.132")
	if err == nil {
		panic("should yeild invalid integer string")
	}
	fmt.Printf("err1: %v\n", err)
	_, err2 := utils.GetNumberAsUint64(float64(123.123))
	if err2 == nil {
		panic("should yield error float must not contain precision")
	}
	fmt.Printf("err2: %v\n", err2)

}
