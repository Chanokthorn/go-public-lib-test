package quack

var s Service

type Service interface {
	Execute(sound string) error
}

func Set(service Service) {
	s = service
}

func Execute(sound string) error {
	return s.Execute(sound)
}
