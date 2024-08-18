package service

import "github.com/ShoreLab/shorelab-backend/lib/dto"

func (s *Service) GetProjectsService() (*dto.ProjectListResponse, error) {
	var res dto.ProjectListResponse
	proj, err := s.repository.GetProjects()
	if err != nil {
		return nil, err
	}

	res.Data = append(res.Data, proj...)

	return &res, nil
}

func (s *Service) GetProjectDetailsService(projectName string) (*dto.ProjectDetailResponse, error) {
	var res dto.ProjectDetailResponse
	proj, err := s.repository.GetProjectByName(projectName)
	if err != nil {
		return nil, err
	}
	res.Data = proj
	return &res, nil
}
