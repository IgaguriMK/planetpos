package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/IgaguriMK/planetpos/model"
)

func main() {
	var systemName string
	flag.StringVar(&systemName, "name", "", "System name")
	var refBody string
	flag.StringVar(&refBody, "ref", "", "Reference body name")

	flag.Parse()

	if systemName == "" {
		log.Fatal("No system name")
	}
	if refBody == "" {
		log.Fatal("No reference body name")
	}

	resp, err := http.Get("https://www.edsm.net/api-system-v1/bodies?systemName=" + systemName)
	if err != nil {
		log.Fatal("API access error: ", err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	var bodies model.Bodies
	err = dec.Decode(&bodies)
	if err != nil {
		log.Fatal("JSON decode error: ", err)
	}

	bodies.RefBody = refBody

	var offset int64 = -1
	for _, b := range bodies.Bodies {
		if b.Name == refBody {
			offset = b.Offset
			break
		}
	}
	if offset == -1 {
		log.Fatalf("Can't find %q in system.", refBody)
	}

	filtered := make([]model.Body, 0)
	for _, b := range bodies.Bodies {
		if b.Offset == offset {
			filtered = append(filtered, b)
		}
	}
	bodies.Bodies = filtered

	f, err := os.Create("bodies.json")
	if err != nil {
		log.Fatal("Can't open output file: ", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	err = enc.Encode(bodies)
	if err != nil {
		log.Fatal("JSON encode error: ", err)
	}
}
