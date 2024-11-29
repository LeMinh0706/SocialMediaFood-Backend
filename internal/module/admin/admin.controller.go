package admin

type AdminController struct {
	service IAdminService
}

func NewAdminController(service IAdminService) AdminController {
	return AdminController{
		service: service,
	}
}
