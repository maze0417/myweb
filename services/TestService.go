package services

type ITestService interface {
	Hello(to string) string
}

type TestService struct {
}

func (s *TestService) Hello(to string) string {
	return "hi " + to
}

type Test2Service struct {
}

func (s *Test2Service) Hello(to string) string {
	return "hiiii " + to
}
