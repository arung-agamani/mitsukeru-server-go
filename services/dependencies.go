package services

type Dependencies struct {
	UserService UserService
}

func NewDependencies() Dependencies {
	userService := NewUserService()
	return Dependencies{
		UserService: userService,
	}
}
