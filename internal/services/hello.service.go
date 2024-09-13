package services

import "github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"

type HelloService struct {
	helloRepo *repo.HelloRepo
}

func NewHelloService() *HelloService {
	return &HelloService{
		helloRepo: repo.NewHelloRepo(),
	}
}

func (hs *HelloService) GetName(name string) string {
	return "Hello " + hs.helloRepo.GetName(name)
}

func (hs *HelloService) GetId(id int) int {
	return hs.helloRepo.GetId(id)
}

func (hs *HelloService) GetInfo(name string, id int) repo.Hello {
	return repo.Hello{Name: name, Id: id}
}
