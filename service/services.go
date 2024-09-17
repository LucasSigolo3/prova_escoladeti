package service

import (
	"CRUD/data/request"
	"CRUD/data/response"
)

type ComputadorService interface {
	Create(computador request.CreateComputadorRequest)
	Update(computador request.UpdateComputadorRequest)
	Delete(computadorId int)
	FindById(computadorId int) response.ComputadorResponse
	FindAll() []response.ComputadorResponse
}
