package repo

type HelloRepo struct{}

func NewHelloRepo() *HelloRepo {
	return &HelloRepo{}
}

type Hello struct {
	Name string `json:"name" binding:"required" example:"Hiro"`
	Id   int    `json:"id" binding:"required"`
}

func (hr *HelloRepo) GetName(name string) string {
	return name
}

func (hr *HelloRepo) GetId(id int) int {
	return id
}

func (hr *HelloRepo) GetInfo(name string, id int) Hello {
	return Hello{Name: name, Id: id}
}
