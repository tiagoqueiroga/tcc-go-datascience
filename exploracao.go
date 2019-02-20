package main

import (
	"bufio"
	"fmt"
	"os"

	// Gota
	"github.com/kniren/gota/dataframe"
	"github.com/kniren/gota/series"
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

func aguardaUsuario() {
	println("\n\n Pressione Qualquer tecla para continuar")
	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadRune()
	_ = input
}

func main() {

	// 1. Criando dataframe e selecionando principais colunas,
	// por alguma razão o a coluna "App" está vindo como \ufeffApp
	df := criaDataframe("googleplaystore.csv")

	// 2. Corrigindo nome da Coluna(App)
	df = df.Rename("App", "\ufeffApp")

	// 3. Selecionando Colunas Importantes("App", "Category","Rating","Price")
	df = df.Select([]string{"App", "Category", "Rating", "Price"})

	// 4. Apresentando as 10 primeiras linhas do dataframe
	fmt.Println("\n\t Visualização Dataframe")
	fmt.Println(df)
	aguardaUsuario()

	// 5. Filtrando Apps da Categoria "COMMUNICATION"
	fmt.Println("\n\t Filtrando Apps da Categoria 'COMMUNICATION'")
	fmt.Println(df.Filter(dataframe.F{"Category", series.Eq, "COMMUNICATION"}))
	aguardaUsuario()

	// 6. Ordenando Apps mais bem avaliados onde Rating diferente de vazio
	fmt.Println("\n\t Ordenando Apps mais bem avaliados")
	dfTopRating := df.Filter(dataframe.F{"Rating", series.GreaterEq, "4.1"}).Select([]string{"App", "Rating"})
	fmt.Println(dfTopRating)
	aguardaUsuario()

}
