package main

import (
	"fmt"
	"os"

	// Pacotes do Gota
	"github.com/kniren/gota/dataframe"
	// Pacotes do Golearn
	//"github.com/sjwhitworth/golearn/base"
	//"github.com/sjwhitworth/golearn/evaluation"
	//"github.com/sjwhitworth/golearn/knn"
	//
)

func verificaArquivo(e error) {
	if e != nil {
		panic(e)
	}
}

func criaDataframe(file string) (df dataframe.DataFrame) {
	arquivo, _ := os.Open(file)
	return dataframe.ReadCSV(arquivo)
}

func mostraAnaliseDescritiva(df dataframe.DataFrame) {
	fmt.Println(df.Describe())
}

func main() {

	// Criando dataframe
	df := criaDataframe("googleplaystore.csv")

	// Analise descritiva dos dados
	mostraAnaliseDescritiva(df)

	// Filtrando resultados - Categoria comunicação
	//dfCatCommunication := df.Filter(dataframe.F{"Category", series.Eq, "COMMUNICATION"})
	//_ = dfCatCommunication

	// Mostrando aplicativos menos avaliados
	//dfRatings := df.Arrange(dataframe.Sort("Rating"))

	//_ = dfRatings

	//fmt.Println(dfRatings)

}
