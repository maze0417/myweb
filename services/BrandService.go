package services

type IBrandService interface {
	GetBrandIdAndStatusByCredentials(clientId string, clientSecret string) (id string, status int)
}

type BrandService struct {
}

var _service ITestService

func NewBrandService(service ITestService) *BrandService {
	_service = service
	return &BrandService{}
}

func (s *BrandService) GetBrandIdAndStatusByCredentials(clientId string, clientSecret string) (id string, status int) {
	_service.Hello("1")
	return
}
