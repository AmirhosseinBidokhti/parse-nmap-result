package utilities

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadXML(filePath string) []byte {
	xmlFile, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer xmlFile.Close()

	XMLdata, _ := ioutil.ReadAll(xmlFile)
	return XMLdata
}
