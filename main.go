package main

import (
	"os/exec"
	"bytes"
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"flag"
	"os"
	"fmt"
)

type Prediction struct {
	Class       string `json:"class"`
	Probability string `json:"probability"`
}
type Data struct {
	Value string `json:"data"`
}

func main() {modelPath := flag.String("model", "", "")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Please type a model path")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data Data
		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		output := predict(data.Value, *modelPath)
		println(data.Value)
		predictions := pretty(output.String())

		json.NewEncoder(w).Encode(predictions)
	})

	http.ListenAndServe(":8080", nil)
}

func predict(data string, modelPath string) bytes.Buffer {
	var out, buffer bytes.Buffer

	cmd := exec.Command("fasttext", "predict-prob", modelPath, "-", "3")
	cmd.Stdout = &out
	buffer.Write([]byte(data))
	cmd.Stdin = &buffer

	err := cmd.Run()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return out
}

func pretty(output string) []Prediction {
	var prediction []Prediction
	words := strings.Fields(output)

	for i := 0; i < len(words)-1; i++ {
		prediction = append(prediction, Prediction{
			Class:       words[i],
			Probability: words[i+1],
		})
		i++;
	}

	return prediction
}
