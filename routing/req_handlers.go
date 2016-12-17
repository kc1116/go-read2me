package routing

import "net/http"

//MaxFileSize is max size for file uploads
const MaxFileSize = 2 * 1024 * 1024

// addImage is the request handler for the /add-image
// responds with a json object
// {status:"500", message: error message}
// {status:"200", message: current path to file uploaded}
func read(w http.ResponseWriter, r *http.Request) {
	// the FormFile function takes in the POST input id file check size

	/*text := r.FormValue("text")

	// stream straight to client(browser)
	w.Header().Set("Content-type", "audio/mpeg")

	if _, err := b.WriteTo(w); err != nil { // <----- here!
		fmt.Fprintf(w, "%s", err)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Res{"200", newImgLoc}); err != nil {
		log.Println("JSON encoder error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}*/

}
