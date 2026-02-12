package prediction

type Service interface {
	Predict(body string) (string, error)
}
