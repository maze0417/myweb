package services

type ITestService interface {
	MyTest(to string) string
}

type TestService struct {
}

func (s *TestService) MyTest(to string) string {
	return "hi " + to
}
