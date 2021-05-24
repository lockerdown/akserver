/* package monitoring

A web framework includes app server, logger, panicer, util and so on.
*/
package monitoring

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	monitor := r.Group("/log/hids/monitor")
	{
		monitor.POST("/file", fileMonitor)
		monitor.POST("/process", processMonitor)
		monitor.POST("/net", netMonitor)
	}
}
