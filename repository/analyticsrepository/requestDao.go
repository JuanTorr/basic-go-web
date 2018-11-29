package analyticsrepository

import (
	"fmt"
	"net/http"
)

//RequestDao defines the interface
type RequestDao struct {
}

//SaveRequest save the given request in the database
func (dao RequestDao) SaveRequest(r http.Request) error {

	/*Proto:  c.Request.Proto,c.Request.Method*/
	fmt.Printf("Saving request: %+v\n", r.URL.String())
	return nil
}
