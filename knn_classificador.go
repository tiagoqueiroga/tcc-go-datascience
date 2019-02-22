package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	rawData, err := base.ParseCSVToInstances("bases/iris.csv", true)

	// Verifica se arquivo foi encontrado
	if err != nil {
		panic(err)
	}

	//Inicia novo classificador KNN, 2 vizinhos
	cls := knn.NewKnnClassifier("manhattan", "linear", 2)

	// Separa base de treinamento e testes
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(trainData)

	// Calcula a distancia Manhattan e retorna o label mais popular
	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}

	// Imprime a precisao
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
}
