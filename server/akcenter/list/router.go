/* package list */
package list

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	monitor := r.Group("/list")
	{
		monitor.GET("/user", test)
	}
}
