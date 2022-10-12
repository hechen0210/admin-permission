package service

import "admin-permission/repository"

type RouteService struct {
	routeRepo *repository.RouteRepository
}

func NewRouteService() *RouteService {
	return &RouteService{
		routeRepo: repository.NewRouteRepository(),
	}
}

func (rs *RouteService) GetList() []string {
	data := []string{}
	list, err := rs.routeRepo.GetAll()
	if err != nil {
		return data
	}
	for _, item := range list {
		data = append(data, item.Url)
	}
	return data
}
