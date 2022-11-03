package repository

import "github.com/ferriantitus/golang-mock/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
