package utils

import (
	"bytes"
	"edu_v2/graph/model"
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadQuestionfile(file graphql.Upload, name string) (model.Collection, error) {
	TCH := os.Getenv("TEXT_CONVERTER_HOST")
	TCP := os.Getenv("TEXT_CONVERTER_PORT")

	url := fmt.Sprintf("http://%s:%s/getTestTextFile", TCH, TCP)

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		return model.Collection{}, err
	}

	_, err = io.Copy(part, file.File)
	if err != nil {
		return model.Collection{}, err
	}

	err = writer.Close()
	if err != nil {
		return model.Collection{}, err
	}

	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return model.Collection{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Collection{}, err
	}
	defer resp.Body.Close()

	// Read and log the raw response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Collection{}, err
	}
	fmt.Printf("Raw response: %s\n", bodyBytes)

	var response map[string]interface{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return model.Collection{}, fmt.Errorf("failed to parse response as JSON: %v", err)
	}

	questions, ok := response["Questions"].(string)
	if !ok {
		return model.Collection{}, fmt.Errorf("unexpected response format")
	}

	collection := model.Collection{
		Title:     name,
		Questions: questions,
	}

	return collection, nil
}
