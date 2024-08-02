package service

func (s *Service) GetUser() (*UserResponse, error) {
	u, err := s.repository.GetUser()
	if err != nil {
		return nil, err
	}
	res := UserResponse{
		Username: u.Username,
		Password: u.Password,
	}

	return &res, err
}
