package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/utils"
)

type usersService struct {
	UsersRepo repos.UsersRepo
}

func NewUsersService(repo repos.UsersRepo) *usersService {
	return &usersService{
		UsersRepo: repo,
	}
}

func (ss *usersService) GetAll() ([]models.UserGetAllResponse, error) {
	var items []models.UserGetAllResponse

	res, err := ss.UsersRepo.GetAll()
	if err != nil {
		return items, err
	}

	for _, item := range res {
		items = append(items, models.UserGetAllResponse{
			ID:        item.ID,
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Username:  item.Username,
			LastLogin: item.LastLogin,
			IsAdmin:   item.IsAdmin,
			UserType:  item.UserType,
			UpdatedAt: item.UpdatedAt,
			CreatedAt: item.CreatedAt,
		})
	}

	return items, err
}

func (ss *usersService) CreateOne(
	first_name,
	last_name,
	username,
	password string,
	is_admin bool,
	user_type int,
	hash_cost int,
) (*models.User, error) {
	var item *models.User

	hashed_password, err := utils.HashPassword(password, hash_cost)
	if err != nil {
		return item, err
	}

	uid, err := uuid.NewRandom()
	if err != nil {
		return item, err
	}

	time_now := time.Now()

	item, err = ss.UsersRepo.CreateOne(
		uid,
		first_name,
		last_name,
		username,
		hashed_password,
		time_now,
		is_admin,
		user_type,
		time_now, time_now)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (ss *usersService) DeleteOneByID(uid uuid.UUID) error {

	err := ss.UsersRepo.DeleteOneByID(uid)
	if err != nil {
		return err
	}

	return nil
}
