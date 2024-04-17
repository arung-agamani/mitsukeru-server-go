package services

type Dependencies struct {
	UserService     UserService
	ItemTypeService ItemTypeService
}

func NewDependencies() Dependencies {
	userService := NewUserService()
	itemTypeService := NewItemTypeService()
	return Dependencies{
		UserService:     userService,
		ItemTypeService: itemTypeService,
	}
}
