package response

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kshamiev/sungora/pkg/errs"
)

type Request struct {
	request *http.Request
}

// JsonBodyDecode декодирование полученного тела запроса в формате json в объект
func (r *Request) JSONBodyDecode(object interface{}) (err error) {
	var body []byte
	if body, err = ioutil.ReadAll(r.request.Body); err != nil {
		return errs.NewInternalServer(err)
	}
	if len(body) == 0 {
		return errs.NewBadRequest(nil, "пустое тело запроса")
	}
	if err = json.Unmarshal(body, object); err != nil {
		return errs.NewBadRequest(err, err.Error())
	}
	return nil
}
