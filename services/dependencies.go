package services

type Dependencies struct {
	UserService     UserService
	ItemTypeService ItemTypeService
	LostItemService LostItemService
	AuthService     AuthService
	EventService    EventService
}

func NewDependencies() Dependencies {
	userService := NewUserService()
	itemTypeService := NewItemTypeService()
	lostItemService := NewLostItemService()
	authService := NewAuthService()
	eventService := NewEventService()
	return Dependencies{
		UserService:     userService,
		ItemTypeService: itemTypeService,
		LostItemService: lostItemService,
		AuthService:     authService,
		EventService:    eventService,
	}
}
