package api

import "github.com/kc1116/go-read2me/synthesize"

//AudioResponse . . .Response type for text synthesizing routines
type AudioResponse struct {
	Audio   []byte `json:"audio"`
	Error   error  `json:"error"`
	Status  int32  `json:"status"`
	Message string `json:"message"`
}

//SynthesizeText . . .
func SynthesizeText(toSynthesize string, voiceID string) AudioResponse {
	c := make(chan synthesize.Response)
	var res AudioResponse
	go func() { c <- synthesize.PlainText(toSynthesize, voiceID) }()

	synthRes := <-c

	if synthRes.Error != nil {
		res.Error = synthRes.Error
		res.Message = "There was an error converting your plain text to audio."
		res.Status = 500

		return res
	}

	res.Audio = synthRes.Audio
	res.Status = 200
	res.Message = "The plain text was converted to audio successfully."

	return res
}
