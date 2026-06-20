package user

type handler struct {
	service service
}

func NewHandler(service service) *handler {
	return &handler{service}
}

func (h *handler) RegisterUser() {}

func (h *handler) GetAllUsers() []User {
	return []User{}
}
