package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/IgaguriMK/planetpos/model"
)

func main() {
	files, err := ioutil.ReadDir("dists")
	if err != nil {
		log.Fatal("Can't open 'dists' dir: ", err)
	}

	dists := make([]model.Distance, 0)

	for _, finfo := range files {
		name := finfo.Name()
		if !strings.HasSuffix(name, ".json") {
			continue
		}

		f, err := os.Open("dists/" + name)
		if err != nil {
			log.Fatalf("Can't open %q: %e", name, err)
		}
		dec := json.NewDecoder(f)
		var rd RawDist
		err = dec.Decode(&rd)
		if err != nil {
			log.Fatal("JSON decode error: ", err)
		}
		f.Close()

		dists = rd.Distances(dists)
	}

	distances := model.Distances{
		Distances: dists,
	}

	f, err := os.Create("distances.json")
	if err != nil {
		log.Fatal("Can't create 'distances.json': ", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	err = enc.Encode(distances)
	if err != nil {
		log.Fatal("JSON encode error: ", err)
	}
}

type RawDist struct {
	From  string `json:"from"`
	Time  string `json:"time"`
	Dists []struct {
		Dist float64 `json:"dist"`
		Name string  `json:"name"`
	} `json:"dists"`
}

func (rd RawDist) Distances(arr []model.Distance) []model.Distance {
	for _, d := range rd.Dists {
		arr = append(
			arr,
			model.Distance{
				Time: model.ParseRealTime(rd.Time),
				From: rd.From,
				To:   d.Name,
				Dist: d.Dist,
			},
		)
	}

	return arr
}
