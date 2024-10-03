package hello

type HelloService struct {
	helloRepo *HelloRepo
}

func NewHelloService() *HelloService {
	return &HelloService{
		helloRepo: NewHelloRepo(),
	}
}

func (hs *HelloService) GetName(name string) string {
	return "Hello " + hs.helloRepo.GetName(name)
}

func (hs *HelloService) GetId(id int) int {
	return hs.helloRepo.GetId(id)
}

func (hs *HelloService) GetInfo(name string, id int) Hello {
	return Hello{Name: name, Id: id}
}
