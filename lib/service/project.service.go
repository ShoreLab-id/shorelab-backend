package service

import "github.com/ShoreLab/shorelab-backend/lib/dto"

func (s *Service) GetProjectsService() (*dto.ProjectListResponse, error) {
	var res dto.ProjectListResponse
	_res, err := s.repository.GetProjects()
	if err != nil {
		return nil, err
	}

	for _, proj := range _res {
		p := dto.Project{
			ID:       proj.ID,
			Title:    proj.Title,
			Location: proj.Location,
			Status:   proj.Status,
			Type:     proj.Type,
		}
		res.Data = append(res.Data, &p)
	}

	return &res, nil
}
