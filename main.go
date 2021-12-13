package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"time"

	"github.com/AlexandreJSimon/portfolio-investimento/services"
	"github.com/AlexandreJSimon/portfolio-investimento/services/calc"
	"github.com/AlexandreJSimon/portfolio-investimento/services/greedy"
	"github.com/AlexandreJSimon/portfolio-investimento/services/randon"
	"github.com/gocarina/gocsv"
)

func main() {
	//Flags contendo o nome do arquivo e o numero maximo de ativos na carteira
	fileName := flag.String("fileName", "", "File name")
	flag.Parse()

	if *fileName == "" {
		panic("Arquivo não informado !")
	}
	start := time.Now()
	//Abrindo o arquivo
	in, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	defer in.Close()

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ','
		return r
	})

	assets := []*services.Assets{}

	err = gocsv.UnmarshalFile(in, &assets)

	if err != nil {
		panic(err)
	}

	returnRiskCalc := calc.NewService(calc.ServiceInput{Assets: assets})

	returnRisk := returnRiskCalc.Run()
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("Tempo de leitura do arquivo e calculos %s", elapsed))
	fmt.Println()

	start = time.Now()
	// Ordena dados
	sort.Slice(returnRisk, func(i, j int) bool {
		return returnRisk[i].ReturnRisk < returnRisk[j].ReturnRisk
	})
	elapsed = time.Since(start)
	fmt.Println(fmt.Sprintf("Tempo de ordenação dos dados %s", elapsed))
	fmt.Println()

	//Printa solução gulosa
	fmt.Println("Solução com algoritmo Guloso:")
	start = time.Now()
	algGreedy := greedy.NewService(greedy.ServiceInput{ReturnRisk: returnRisk})
	gr, grr := algGreedy.Run()
	for _, v := range gr {
		fmt.Println(fmt.Sprintf("ativo %s com peso %f", v.Asset, v.Weight))
	}
	fmt.Println(fmt.Sprintf("Risco retorno total %f", grr))

	elapsed = time.Since(start)
	fmt.Println(fmt.Sprintf("Risco retorno total %f", grr))
	fmt.Println(fmt.Sprintf("Tempo de execução da solução gulosa %s", elapsed))
	fmt.Println()

	//Printa solucao aleatoria
	fmt.Println("Solução com pesos aleatorios:")
	start = time.Now()
	algRand := randon.NewService((randon.ServiceInput{ReturnRisk: returnRisk}))
	al, alr := algRand.Run()
	for _, v := range al {
		fmt.Println(fmt.Sprintf("ativo %s com peso %f", v.Asset, v.Weight))
	}
	fmt.Println(fmt.Sprintf("Risco retorno total %f", alr))
	elapsed = time.Since(start)
	fmt.Println(fmt.Sprintf("Tempo de execução da solução aleatoria %s", elapsed))
	fmt.Println()
}
