package randon

import (
	"math/rand"

	"github.com/AlexandreJSimon/portfolio-investimento/services"
)

type ServiceInput struct {
	ReturnRisk []*services.AssetsReturnRisk
}

// Service ...
type Service struct {
	in ServiceInput
}

// New Service ...
func NewService(in ServiceInput) *Service {
	return &Service{in}
}

func (s *Service) Run() ([]*services.AssetsReturnRisk, float64) {
	returnRisk := s.in.ReturnRisk
	maxWeight := 100
	//Gera pesos aleatorios
	for i := 0; i < len(returnRisk); i++ {
		rd := rand.Intn(maxWeight)
		if len(returnRisk) != i+1 {
			returnRisk[i].Weight = float64(rd)
		} else {
			returnRisk[i].Weight = float64(maxWeight)
		}
		maxWeight = maxWeight - rd
	}

	return returnRisk, returnRiskPortfolio(returnRisk)
}

//Calcula risco retorno total ignorando os pesos zerados
func returnRiskPortfolio(returnRisk []*services.AssetsReturnRisk) float64 {
	rr := 0.0
	c := 1
	for _, v := range returnRisk {
		if v.Weight > 0 {
			c++
			rr += v.ReturnRisk * (v.Weight / 100)
		}
	}

	return rr / float64(c)
}
