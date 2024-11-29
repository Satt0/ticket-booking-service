package middleware

import (
	res_format "http-server/internal/shared/res-format"

	"github.com/gin-gonic/gin"
)

func (m *MiddleWare) CreateErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Execute the request handler
		defer func() {
			err := recover()
			if err != nil {
				m.logger.Out.Error(err)
				res_format.FormatResponse500(c)
				c.Abort()
			}
		}()
		c.Next()
	}
}
