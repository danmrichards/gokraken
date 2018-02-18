package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

const (
	templateFile = "cmd/currency.tpl"
	outputFile   = "currency.go"

	krakenAssetsURL = "https://api.kraken.com/0/public/Assets"
)

type assetsResponse struct {
	Result map[string]interface{} `json:"result"`
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	assets, err := getAssets()
	if err != nil {
		log.Fatalf("could not get latest assets: %v", err)
	}

	err = generateAssetPackage(assets)
	if err != nil {
		log.Fatalf("could not generate asset package: %v", err)
	}
}

func getAssets() (assets assetsResponse, err error) {
	res, err := http.Get(krakenAssetsURL)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	err = json.Unmarshal(b, &assets)
	return
}

func generateAssetPackage(assets assetsResponse) error {
	assetCodes := make([]string, len(assets.Result))

	i := 0
	for assetCode := range assets.Result {
		assetCodes[i] = assetCode
		i++
	}

	tpl, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return fmt.Errorf("cannot open template file: %v", err)
	}

	t := template.Must(template.New("currency").Parse(string(tpl)))
	buf := new(bytes.Buffer)
	err = t.Execute(buf, assetCodes)
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
