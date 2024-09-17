package repository

import (
	"CRUD/data/request"
	"CRUD/helper"
	"CRUD/model"
	"errors"

	"gorm.io/gorm"
)

type ComputadorRepositoryImpl struct {
	Db *gorm.DB
}

func NewComputadorRepositoryImpl(Db *gorm.DB) ComputadorRepository {
	return &ComputadorRepositoryImpl{Db: Db}
}

func (t ComputadorRepositoryImpl) Save(computador model.Computador) {
	result := t.Db.Create(&computador)
	helper.ErrorPanic(result.Error)

}

func (t ComputadorRepositoryImpl) Update(computador model.Computador) {
	var updateComputador = request.UpdateComputadorRequest{
		Id:   computador.Id,
		Name: computador.Name,
	}
	result := t.Db.Model(&computador).Updates(updateComputador)
	helper.ErrorPanic(result.Error)
}

func (t ComputadorRepositoryImpl) Delete(computadorId int) {
	var computador model.Computador
	result := t.Db.Where("id = ?", computadorId).Delete(&computador)
	helper.ErrorPanic(result.Error)
}

func (t ComputadorRepositoryImpl) FindById(computadorId int) (model.Computador, error) {
	var computador model.Computador
	result := t.Db.Find(&computador, computadorId)
	if result != nil {
		return computador, nil
	} else {
		return computador, errors.New("computador não encontrado")
	}
}

func (t ComputadorRepositoryImpl) FindAll() []model.Computador {
	var computador []model.Computador
	results := t.Db.Find(&computador)
	helper.ErrorPanic(results.Error)
	return computador
}
