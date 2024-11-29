package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetMapFromGinContext(c *gin.Context, key string) (map[string]interface{}, error) {
	thisMap, exists := c.Get(key)
	if !exists {
		return nil, errors.New("value not exist in context")
	}
	return thisMap.(map[string]interface{}), nil

}
func GetDataFromMap(thisMap map[string]interface{}, key string) (interface{}, error) {
	value, ok := thisMap[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return value, nil
}
