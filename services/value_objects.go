package services

// Struct do CSV
type Assets struct {
	Ativo     string  `csv:"ativo"`
	Data      string  `csv:"data"`
	Preco     float64 `csv:"preco"`
	Valor     string  `csv:"valor"`
	Dividendo float64 `csv:"dividendo"`
}

// Struct Risco/Retorno do Ativo
type AssetsReturnRisk struct {
	Asset      string
	ReturnRisk float64
	Weight     float64
}
