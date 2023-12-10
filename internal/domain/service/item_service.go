package service

import (
	"esb-invoice/internal/app/handler/filters"
	"esb-invoice/internal/app/handler/response"
	"esb-invoice/internal/domain/model"
	"esb-invoice/internal/domain/repository"
	"fmt"
)

type ItemService struct {
	itemRepo repository.ItemRepo
	typeRepo repository.TypeRepo
}

func NewItemService(itemRepo repository.ItemRepo, typeRepo repository.TypeRepo) *ItemService {
	return &ItemService{
		itemRepo: itemRepo,
		typeRepo: typeRepo,
	}
}

func (s *ItemService) Create(item *model.Item) (int, error) {
	checkTypeItem, err := s.typeRepo.FindById(int(item.TypeID))

	if err != nil {
		if err.Error() == fmt.Sprintf("record not found") {
			return 0, fmt.Errorf("type item not found")
		}
		return 0, err
	}

	if checkTypeItem == nil {
		return 0, err
	}

	checkItemName, err := s.itemRepo.FindByItemName(item.Name)
	if err != nil {
		if err.Error() != fmt.Sprintf("record not found") {
			return 0, err
		}
	}

	if len(checkItemName.Name) > 0 {
		return 0, fmt.Errorf("item name already exists")
	}

	itemId, err := s.itemRepo.Create(item)

	if err != nil {
		return 0, err
	}

	return itemId, nil
}

func (s *ItemService) FindById(id int) (*response.ItemResponse, error) {
	item, err := s.itemRepo.FindById(id)
	var responseItem response.ItemResponse

	if err != nil {
		return nil, err
	}

	responseItem = response.ItemResponse{
		ID:     item.ID,
		Name:   item.Name,
		Price:  item.Price,
		TypeID: item.TypeID,
		Type: response.TypeOnItemResponse{
			ID:   item.Type.ID,
			Name: item.Type.Name,
		},
	}

	return &responseItem, nil
}

func (s *ItemService) SoftDelete(id int) error {
	err := s.itemRepo.SoftDelete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *ItemService) UpdateById(id int, update *model.Item) (*model.Item, error) {
	checkItem, err := s.itemRepo.FindByNameAndNotId(update.Name, id)
	if err != nil {
		if err.Error() != fmt.Sprintf("record not found") {
			return nil, err
		}
	}

	if len(checkItem.Name) > 0 {
		return nil, fmt.Errorf("item name already exists")
	}

	item, err := s.itemRepo.UpdateById(id, update)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ItemService) FindAll(filter *filters.ItemFilter, page int, pageSize int) ([]response.ItemResponse, int64, error) {
	var items []model.Item
	var totalRecords int64
	var responseItems []response.ItemResponse
	items, totalRecords, err := s.itemRepo.FindAll(*filter, page, pageSize)

	if err != nil {
		return nil, 0, err
	}

	for _, item := range items {
		responseItem := response.ItemResponse{
			ID:     item.ID,
			Name:   item.Name,
			Price:  item.Price,
			TypeID: item.TypeID,
			Type: response.TypeOnItemResponse{
				ID:   item.Type.ID,
				Name: item.Type.Name,
			},
		}

		responseItems = append(responseItems, responseItem)
	}

	return responseItems, totalRecords, nil
}

func (s *ItemService) CountAll() (int64, error) {
	totalRecords, err := s.itemRepo.CountAll()

	if err != nil {
		return 0, err
	}

	return totalRecords, nil
}
