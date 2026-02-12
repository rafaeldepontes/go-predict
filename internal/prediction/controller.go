package prediction

import "net/http"

type Controller interface {
	Predict(w http.ResponseWriter, r *http.Request)
}
