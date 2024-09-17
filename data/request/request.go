package request

type CreateComputadorRequest struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Cor            string `json:"cor"`
	DataFabricacao int    `json:"dataFabricacao"`
}

type UpdateComputadorRequest struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Cor            string `json:"cor"`
	DataFabricacao int    `json:"dataFabricacao"`
}
