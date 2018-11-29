package analyticsservice

import (
	"net/http"
)

//RequestDaoService defines the interface
type RequestDaoService interface {
	SaveRequest(http.Request) error
}

//PersistData persist the request data
func PersistData(dao RequestDaoService, r http.Request) error {
	return dao.SaveRequest(r)
}
