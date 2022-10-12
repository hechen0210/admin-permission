package service

import "admin-permission/repository"

type PageService struct {
	pageRepo *repository.PageRepository
}

func NewPageService() *PageService {
	return &PageService{
		pageRepo: repository.NewPageRepository(),
	}
}

func (ps *PageService) Get(page, pageSize int, codition map[string]interface{}) {

}
