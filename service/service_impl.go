package service

import (
	"CRUD/data/request"
	"CRUD/data/response"
	"CRUD/helper"
	"CRUD/model"
	"CRUD/repository"
	"github.com/go-playground/validator/v10"
)

type ComputadorServiceImpl struct {
	ComputadorRepository repository.ComputadorRepository
	Validate             *validator.Validate
}

func NewComputadorServiceImpl(computadorRepository repository.ComputadorRepository, validate *validator.Validate) ComputadorService {
	return &ComputadorServiceImpl{
		ComputadorRepository: computadorRepository,
		Validate:             validate,
	}
}

func (t ComputadorServiceImpl) Create(computador request.CreateComputadorRequest) {
	err := t.Validate.Struct(computador)
	helper.ErrorPanic(err)
	computadorModel := model.Computador{
		Id:             computador.Id,
		Name:           computador.Name,
		Cor:            computador.Cor,
		DataFabricacao: int8(int(computador.DataFabricacao)),
	}
	t.ComputadorRepository.Save(computadorModel)
}

func (t ComputadorServiceImpl) Update(computador request.UpdateComputadorRequest) {
	computadorData, err := t.ComputadorRepository.FindById(computador.Id)
	helper.ErrorPanic(err)
	computadorData.Name = computador.Name
	t.ComputadorRepository.Update(computadorData)
}

func (t ComputadorServiceImpl) Delete(computadorId int) {
	t.ComputadorRepository.Delete(computadorId)
}

func (t ComputadorServiceImpl) FindById(computadorId int) response.ComputadorResponse {
	computadorData, err := t.ComputadorRepository.FindById(computadorId)
	helper.ErrorPanic(err)

	computadorResponse := response.ComputadorResponse{
		Id:             computadorData.Id,
		Name:           computadorData.Name,
		Cor:            computadorData.Cor,
		DataFabricacao: int(computadorData.DataFabricacao),
	}
	return computadorResponse
}

func (t ComputadorServiceImpl) FindAll() []response.ComputadorResponse {
	result := t.ComputadorRepository.FindAll()

	var computadores []response.ComputadorResponse
	for _, value := range result {
		computador := response.ComputadorResponse{
			Id:             value.Id,
			Name:           value.Name,
			Cor:            value.Cor,
			DataFabricacao: int(value.DataFabricacao),
		}
		computadores = append(computadores, computador)
	}
	return computadores
}
