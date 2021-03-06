package main

import (
	"github.com/rajdeepd/einstein-go-apis/vision"
	"os"
	"encoding/json"
	"fmt"
	"github.com/rajdeepd/einstein-go-apis/common"
)

func main() {
	modelId := "FoodImageClassifier"
	sampleLocation := "http://metamind.io/images/foodimage.jpg"
	accessToken := os.Getenv("ACCESSKEY")
	response, err := vision.Classify(modelId, sampleLocation, accessToken)
	if err != nil {
		fmt.Println(err)
	}

	var dat *common.PredictResponse
	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}
	for _, value := range dat.Probabilities {
		fmt.Println(value)
	}

}