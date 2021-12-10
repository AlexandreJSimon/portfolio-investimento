package calc

import (
	"sort"
	"strings"

	"github.com/AlexandreJSimon/portfolio-investimento/services"
	"github.com/montanaflynn/stats"
)

// Parametros para inciar o serviço
type ServiceInput struct {
	Assets []*services.Assets
}

// Service ...
type Service struct {
	in ServiceInput
}

// New Service ...
func NewService(in ServiceInput) *Service {
	return &Service{in}
}

// Run ...
func (s *Service) Run() []*services.AssetsReturnRisk {
	assets := s.in.Assets

	//Ordenando ativos
	sortAssets(assets)

	// Iniciando Ativo Para Comparar com o proximo valor
	as := assets[0].Ativo
	// Atribuindo Valor inicial
	iv := assets[0].Preco
	//Iniciando dividendo, preço anterior, retorno, risco normalizado
	var dividends, pr, rt, rn float64
	//Iniciando valor total e variação
	var values, variance []float64

	assetsReturnRisk := []*services.AssetsReturnRisk{}

	for k, asset := range assets {
		// verifica se Ativo anterio é diferente do ativo atual ou se é a ultima iteracao
		if as != asset.Ativo || (k+1) == len(assets) {
			//Calculando Desvio padrão da variação
			standardDeviation, _ := stats.StandardDeviationSample(variance)
			// Calculando media do valor total
			mean, _ := stats.Mean(values)
			// Calculando risco normalizado %
			rn = standardDeviation / mean * 100
			//Calculando o retorno total %
			rt = (((dividends + pr) - iv) / iv) * 100
			assetsReturnRisk = append(assetsReturnRisk, &services.AssetsReturnRisk{
				// Ativo
				Asset: as,
				//Calculando Risco do retorno %
				ReturnRisk: ((rn / rt) * 100),
			})
			//Atribuindo valor inicial do proximo ativo
			iv = asset.Preco
			//Resetando valores
			variance = nil
			values = nil
			dividends = 0
		}

		//Somando todos os dividendos
		dividends += asset.Dividendo

		if values != nil {
			//Calculando a variação desconsiderando o calculo para o primeiro preço pois não existe dado do mes anterior
			variance = append(variance, asset.Preco-pr)
		}

		//Criando arrey com todos os preços do ativo
		values = append(values, asset.Preco)

		//Preço e ativo que serão comparados com os valores da proxima iteração
		pr = asset.Preco
		as = asset.Ativo
	}
	return assetsReturnRisk
}

func sortAssets(assets []*services.Assets) {
	//Ordenando ativos por nome e data
	sort.Slice(assets, func(i, j int) bool {
		switch strings.Compare(assets[i].Ativo, assets[j].Ativo) {
		case -1:
			return true
		case 1:
			return false
		}
		return assets[i].Data < assets[j].Data
	})
}
