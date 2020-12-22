package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadModel(url string, modelName string) error {
	//Download to file
	model, err := os.Create(modelName)
	if err != nil {
		return err
	}
	defer model.Close()

	//Get response from URL
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s ", response.Status)
	}

	_, err = io.Copy(model, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	modelName := "distilbert-base-uncased"
	url := "https://huggingface.co/distilbert-base-uncased"
	err := downloadModel(url, modelName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Model %s downloaded in current working dir", modelName)
}
