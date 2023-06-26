package customers

type services struct{}

type Services interface {
	List() ([]customer, error)
}

func NewServices() Services {
	return &services{}
}

func (s *services) List() ([]customer, error) {
	return []customer{{
		FirstName: "Pepe",
		Address:   "services",
	}}, nil
}
