package service

type UserResponse struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (s *Service) GetUser() (*UserResponse, error) {
	u, err := s.repository.GetUser()
	if err != nil {
		return nil, err
	}
	res := UserResponse{
		Name: u.Name,
		// Password: u.Password,
	}

	return &res, err
}
