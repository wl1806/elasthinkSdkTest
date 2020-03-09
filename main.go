package main

import (
	"fmt"
	"log"

	sdk "github.com/SurgicalSteel/elasthink/sdk"
)

func main() {
	availableDocumentType := make([]string, 0)
	availableDocumentType = append(availableDocumentType, "iklan")
	initializeSpec := sdk.InitializeSpec{
		RedisConfig: sdk.RedisConfig{
			Address:   "localhost",
			MaxActive: 3,
			MaxIdle:   1000,
			Timeout:   1000,
			Port:      "6379",
		},
		SdkConfig: sdk.SdkConfig{
			AvailableDocumentType:   availableDocumentType,
			IsUsingStopWordsRemoval: false,
		},
	}
	Es, err := sdk.Initialize(initializeSpec)
	if err != nil {
		log.Println("coba err 1 ", err.Error())
		return
	}
	Es.CreateIndex(sdk.CreateIndexSpec{DocumentID: 1, DocumentName: "pergi ke pantai sendiri", DocumentType: "iklan"})
	Es.CreateIndex(sdk.CreateIndexSpec{DocumentID: 2, DocumentName: "pergi ke hutan bersama teman", DocumentType: "iklan"})
	res1, err := Es.Search(sdk.SearchSpec{DocumentType: "iklan", SearchTerm: "hutan"})
	if err != nil {
		log.Println("coba err 2 ", err.Error())
	}

	log.Printf("%+v \n", res1)
	res2, err := Es.Search(sdk.SearchSpec{DocumentType: "iklan", SearchTerm: "pergi"})
	if err != nil {
		log.Println("coba err 3 ", err.Error())
	}

	log.Printf("%+v \n", res2)
	fmt.Println("End")

}
