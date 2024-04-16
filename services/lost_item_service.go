package services

type CreateLostItemPayload struct {
	Name            string
	Description     string
	ItemTypeID      string
	ReporterName    string
	ReporterContact string
	Returned        bool
}

type LostItemService interface {
	CreateLostItem()
	UpdateLostItem()
	DeleteLostItem()
}
