package service

type userService struct {
}

func NewUserService() UserService {
	s := &userService{}
	return s
}
