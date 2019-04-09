package services

type ITestService interface {
	MyTest(to string) string
}

type TestService struct {
}

func (s *TestService) MyTest(to string) string {
	return "hi " + to
}

type Test2Service struct {
}

func (s *Test2Service) MyTest(to string) string {
	return "hiiii " + to
}
