package fly

var s Service

type Service interface {
	Execute(height int) error
}

func Set(service Service) {
	s = service
}

func Execute(height int) error {
	return s.Execute(height)
}
