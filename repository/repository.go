package repository

import (
	"CRUD/model"
)

type ComputadorRepository interface {
	Save(computador model.Computador)
	Update(computador model.Computador)
	Delete(computadorId int)
	FindById(computadorId int) (computador model.Computador, err error)
	FindAll() []model.Computador
}
