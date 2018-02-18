package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	templateFile = "cmd/assetpair.html"
	outputFile   = "assetpair.go"

	krakenAssetPairsURL = "https://api.kraken.com/0/public/AssetPairs"
)

type assetPairsResponse struct {
	Result map[string]interface{} `json:"result"`
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	assets, err := getAssets()
	if err != nil {
		log.Fatalf("could not get latest asset pairs: %v", err)
	}

	err = generateAssetPackage(assets)
	if err != nil {
		log.Fatalf("could not generate pairs package: %v", err)
	}
}

func getAssets() (pairs assetPairsResponse, err error) {
	res, err := http.Get(krakenAssetPairsURL)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	err = json.Unmarshal(b, &pairs)
	return
}

func generateAssetPackage(pairs assetPairsResponse) error {
	assetPairs := make([]string, 0)

	for assetPair := range pairs.Result {
		if !strings.Contains(assetPair, ".") {
			assetPairs = append(assetPairs, assetPair)
		}
	}

	tpl, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return fmt.Errorf("cannot open template file: %v", err)
	}

	t := template.Must(template.New("assetpair").Parse(string(tpl)))
	buf := new(bytes.Buffer)
	err = t.Execute(buf, assetPairs)
	if err != nil {
		return err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	buf = bytes.NewBuffer(formatted)
	to, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, buf)
	if err != nil {
		return err
	}

	return nil
}
