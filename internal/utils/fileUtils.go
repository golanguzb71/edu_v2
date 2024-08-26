package utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strings"

	"edu_v2/graph/model"
	"github.com/ledongthuc/pdf"
	"github.com/unidoc/unioffice/document"
)

func UploadQuestionFile(file graphql.Upload, name string) (model.Collection, error) {
	fileBytes, err := io.ReadAll(file.File)
	if err != nil {
		return model.Collection{}, fmt.Errorf("failed to read file: %v", err)
	}

	var modifiedText string
	fileName := file.Filename
	if strings.HasSuffix(fileName, ".docx") {
		modifiedText, err = processDocx(fileBytes)
		if err != nil {
			return model.Collection{}, fmt.Errorf("failed to process docx file: %v", err)
		}
	} else if strings.HasSuffix(fileName, ".pdf") {
		modifiedText, err = processPDF(fileBytes)
		if err != nil {
			return model.Collection{}, fmt.Errorf("failed to process pdf file: %v", err)
		}
	} else {
		return model.Collection{}, errors.New("unsupported file type")
	}

	return model.Collection{
		Title:     name,
		Questions: modifiedText,
	}, nil
}

func processDocx(fileBytes []byte) (string, error) {
	doc, err := document.Read(bytes.NewReader(fileBytes), int64(len(fileBytes)))
	if err != nil {
		return "", fmt.Errorf("failed to read docx file: %v", err)
	}

	var modifiedText strings.Builder
	for _, para := range doc.Paragraphs() {
		for _, run := range para.Runs() {
			paragraphText := run.Text()
			modifiedText.WriteString(replaceUnderscoreSequences(paragraphText) + "\n")
		}
	}

	return modifiedText.String(), nil
}

func processPDF(fileBytes []byte) (string, error) {
	fileReader := bytes.NewReader(fileBytes)
	pdfReader, err := pdf.NewReader(fileReader, int64(len(fileBytes)))
	if err != nil {
		return "", fmt.Errorf("failed to read pdf file: %v", err)
	}

	var text strings.Builder
	for pageIndex := 1; pageIndex <= pdfReader.NumPage(); pageIndex++ {
		page := pdfReader.Page(pageIndex)
		pageText, err := page.GetPlainText(nil)
		if err != nil {
			return "", fmt.Errorf("failed to extract text from pdf page %d: %v", pageIndex, err)
		}
		text.WriteString(replaceUnderscoreSequences(pageText) + "\n")
	}

	return text.String(), nil
}

func replaceUnderscoreSequences(text string) string {
	return strings.ReplaceAll(text, "_", "[qazwsxedc]")
}
