package service

import (
	"esb-invoice/internal/app/handler/response"
	"esb-invoice/internal/domain/model"
	"esb-invoice/internal/domain/repository"
)

type TypeService struct {
	typeRepo repository.TypeRepo
}

func NewTypeService(typeRepo repository.TypeRepo) *TypeService {
	return &TypeService{
		typeRepo: typeRepo,
	}
}

func (r *TypeService) FindAll() ([]response.TypeResponse, error) {
	var types []response.TypeResponse

	typesReturn, err := r.typeRepo.FindAll()

	for _, typeItem := range typesReturn {
		types = append(types, response.TypeResponse{
			ID:   typeItem.ID,
			Name: typeItem.Name,
		})
	}

	if err != nil {
		return nil, err
	}

	return types, nil
}

func (r *TypeService) FindById(id int) (*response.TypeResponse, error) {
	typeItem, err := r.typeRepo.FindById(id)

	if err != nil {
		return nil, err
	}
	typeItemReturn := response.TypeResponse{
		ID:   typeItem.ID,
		Name: typeItem.Name,
	}

	return &typeItemReturn, nil
}

func (r *TypeService) Create(typeItem *model.Type) (int, error) {
	typeId, err := r.typeRepo.Create(typeItem)

	if err != nil {
		return 0, err
	}

	return typeId, nil
}

func (r *TypeService) UpdateById(id int, update *model.Type) (*model.Type, error) {
	typeItem, err := r.typeRepo.UpdateById(id, update)

	if err != nil {
		return nil, err
	}

	return typeItem, nil
}

func (r *TypeService) SoftDelete(id int) error {
	err := r.typeRepo.SoftDelete(id)

	if err != nil {
		return err
	}

	return nil
}

func (r *TypeService) FindByName(name string) (*response.TypeResponse, error) {
	typeItem, err := r.typeRepo.FindByName(name)

	if err != nil {
		return nil, err
	}

	typeItemReturn := response.TypeResponse{
		ID:   typeItem.ID,
		Name: typeItem.Name,
	}

	return &typeItemReturn, nil
}

func (r *TypeService) FindByNameAndNotId(name string, id int) (*response.TypeResponse, error) {
	typeItem, err := r.typeRepo.FindByNameAndNotId(name, id)

	if err != nil {
		return nil, err
	}

	typeItemReturn := response.TypeResponse{
		ID:   typeItem.ID,
		Name: typeItem.Name,
	}

	return &typeItemReturn, nil
}
