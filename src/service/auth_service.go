package service

type AuthService interface {
	Auth()
	AuthCheck()
	Token()
	Certs()
	UserInfo()
}

func NewAuthService() AuthService {
	return &authService{}
}

type authService struct{}

func (s *authService) Auth() {
	panic("NotImplementedError")
}

func (s *authService) AuthCheck() {
	panic("NotImplementedError")
}

func (s *authService) Token() {
	panic("NotImplementedError")
}

func (s *authService) Certs() {
	panic("NotImplementedError")
}

func (s *authService) UserInfo() {
	panic("NotImplementedError")
}
