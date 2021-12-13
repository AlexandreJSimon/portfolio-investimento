package greedy

import (
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

	var bestReturnRisk []*services.AssetsReturnRisk
	var current []*services.AssetsReturnRisk

	for i := 0; i < len(returnRisk); i++ {
		// Atribui o peso que sera distribuido
		weight := 1.0

		if len(bestReturnRisk) == 0 {
			// Atribui o peso do primeiro ativo
			current = []*services.AssetsReturnRisk{{
				Asset:      returnRisk[i].Asset,
				ReturnRisk: returnRisk[i].ReturnRisk,
				Weight:     100,
			}}
		} else {
			// Busca o peso do ativo que já foi processado para comparar o que esta sendo processado
			t := getWeightInPortfolio(bestReturnRisk, bestReturnRisk[0].Asset)
			if t > weight {
				// Realiza a subtração do peso do melhor ativo para atribuir o peso para o ativo quer esta sendo processado
				nt := t - weight
				np := []*services.AssetsReturnRisk{}
				var pp float64

				// Itera os aviso ja existentes adicionando em uma nova lista e atualizando seus pesos
				for _, v := range bestReturnRisk {
					if v.Asset == bestReturnRisk[0].Asset {
						pp = nt
					} else {
						pp = v.Weight
					}
					np = append(np, &services.AssetsReturnRisk{
						Asset:      v.Asset,
						ReturnRisk: v.ReturnRisk,
						Weight:     pp,
					})
				}

				// Adiciona o novo ativo na lista
				current = append(np, &services.AssetsReturnRisk{
					Asset:      returnRisk[i].Asset,
					ReturnRisk: returnRisk[i].ReturnRisk,
					Weight:     weight,
				})

			}

		}

		//Verifica se existe a necessidade de continuar atribuindo pesos
		if len(bestReturnRisk) == 0 || returnRiskPortfolio(current) < returnRiskPortfolio(bestReturnRisk) {
			bestReturnRisk = current
		} else {
			break
		}
	}

	return bestReturnRisk, returnRiskPortfolio(bestReturnRisk)
}

//Calcula risco retorno tota
func returnRiskPortfolio(returnRisk []*services.AssetsReturnRisk) float64 {
	rr := 0.0
	for _, v := range returnRisk {
		rr += v.ReturnRisk * (v.Weight / 100)
	}

	return rr / float64(len(returnRisk))
}

//Busca peso já existente na carteira
func getWeightInPortfolio(portfolio []*services.AssetsReturnRisk, asset string) float64 {
	for _, v := range portfolio {
		if v.Asset == asset {
			return v.Weight
		}
	}

	panic("Ativo não se encontra na carteira")
}
