package authenticate

type UserCase interface{}

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Redirect() {}
