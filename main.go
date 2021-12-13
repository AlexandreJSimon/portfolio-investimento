package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

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

	// Ordena dados
	sort.Slice(returnRisk, func(i, j int) bool {
		return returnRisk[i].ReturnRisk < returnRisk[j].ReturnRisk
	})

	//Printa solução gulosa
	fmt.Println("Solução com algoritmo Guloso:")
	algGreedy := greedy.NewService(greedy.ServiceInput{ReturnRisk: returnRisk})
	gr, grr := algGreedy.Run()
	for _, v := range gr {
		fmt.Println(fmt.Sprintf("ativo %s com peso %f", v.Asset, v.Weight))
	}
	fmt.Println(fmt.Sprintf("Risco retorno total %f", grr))
	fmt.Println()

	//Printa solucao aleatoria
	fmt.Println("Solução com pesos aleatorios:")
	algRand := randon.NewService((randon.ServiceInput{ReturnRisk: returnRisk}))
	al, alr := algRand.Run()
	for _, v := range al {
		fmt.Println(fmt.Sprintf("ativo %s com peso %f", v.Asset, v.Weight))
	}
	fmt.Println(fmt.Sprintf("Risco retorno total %f", alr))
	fmt.Println()
}
