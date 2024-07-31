package service

type ImageListResponse struct {
	Data []string `json:"data"`
}

func (s *Service) GetImageList() (*ImageListResponse, error) {
	obj, err := s.repository.GetImageList()

	var res ImageListResponse
	res.Data = append(res.Data, obj...)
	return &res, err
}

func (s *Service) GetImage(img string) ([]byte, string, error) {
	return s.repository.GetImage(img)
}
