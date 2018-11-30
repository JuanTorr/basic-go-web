package analytics

import (
	anRepo "github.com/JuanTorr/basic-go-web/repository/analyticsrepository"
	anservice "github.com/JuanTorr/basic-go-web/service/analyticsservice"

	"github.com/gin-gonic/gin"
)

//AnalizeRequest analizes the request
func AnalizeRequest(c *gin.Context) {
	reqDao := anRepo.RequestDao{}
	anservice.PersistData(reqDao, *c.Request)
}
