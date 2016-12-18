package routing

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kc1116/go-read2me/api"
)

func plainText(w http.ResponseWriter, r *http.Request) {
	toSynthesize := r.FormValue("toSynthesize")
	voiceID := r.FormValue("voiceID")

	if len(toSynthesize) < 1 || len(voiceID) < 1 {
		http.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}

	results := api.SynthesizeText(toSynthesize, voiceID)

	if results.Error == nil && results.Status == 200 {

		b := bytes.NewBuffer(results.Audio)
		if _, err := b.WriteTo(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "audio/mpeg")

		if err := json.NewEncoder(w).Encode(results); err != nil {
			log.Println("JSON encoder error: " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		if err := json.NewEncoder(w).Encode(results); err != nil {
			log.Println("JSON encoder error: " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}
