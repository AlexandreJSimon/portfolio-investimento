package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/AlexandreJSimon/portfolio-investimento/services"
	"github.com/AlexandreJSimon/portfolio-investimento/services/calc"
	"github.com/gocarina/gocsv"
)

func main() {
	//Flags contendo o nome do arquivo e o numero maximo de ativos na carteira
	fileName := flag.String("fileName", "", "File name")
	nAssets := flag.Int("nAssets", 0, "Number of Assets")
	flag.Parse()

	if *fileName == "" || *nAssets == 0 {
		panic("Arquivo não informado ou numero de ativos inválido !")
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

	for _, v := range returnRiskCalc.Run() {
		fmt.Println(v)
	}

	fmt.Println(*fileName)
	fmt.Println(*nAssets)
}
