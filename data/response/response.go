package response

type ComputadorResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Cor            string `json:"cor"`
	DataFabricacao int    `json:"dataFabricacao"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
