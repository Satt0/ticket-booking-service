package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(c *gin.Context) (uint64, error) {

	m, err := GetMapFromGinContext(c, "user")
	if err != nil {
		return 0, errors.New("cannot find key in context")
	}

	userId, err := GetDataFromMap(m, "id")
	if err != nil {
		return 0, errors.New("invalid user data in context")
	}
	return GetNumberAsUint64(userId)
}
func GetNumberAsUint64(num interface{}) (uint64, error) {
	switch v := num.(type) {
	case uint64:
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case float64:
		if v == float64(uint64(v)) { // Ensure no fractional part
			return uint64(v), nil
		}
		return 0, errors.New("not a float integer")
	case string:
		intFromString, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, errors.New("not a string integer")
		}
		return uint64(intFromString), nil
	}
	return 0, errors.New("user id in context is mal-formatted")
}
