package synthesize

import (
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

//Response . . . response struct from all Parse functions, Audio contains synthesized bytes and Error if one here
type Response struct {
	Audio  []byte
	Error  error
	Voices *polly.DescribeVoicesOutput
}

//GetVoices . . . returns set of voices available in AWS Polly
func GetVoices(lang string) Response {
	var res Response
	sess, err := session.NewSession()
	if err != nil {
		res.Error = err
		return res
	}

	svc := polly.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	params := &polly.DescribeVoicesInput{
		LanguageCode: aws.String(lang),
	}
	resp, err := svc.DescribeVoices(params)
	if err != nil {
		res.Error = err
		return res
	}

	res.Voices = resp
	return res
}

//PlainText . . . Takes plain text data input converts it SynthesizedSpeechOutput
func PlainText(toSynthesize string, voiceID string) Response {
	var res Response
	sess, err := session.NewSession()
	if err != nil {
		res.Error = err
		return res
	}

	svc := polly.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	params := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		Text:         aws.String(toSynthesize),
		VoiceId:      aws.String(voiceID),
		TextType:     aws.String("text"),
	}
	resp, err := svc.SynthesizeSpeech(params)

	if err != nil {
		res.Error = err
		return res
	}

	b, err := ioutil.ReadAll(resp.AudioStream)
	if err != nil {
		res.Error = err
		return res
	}
	/*if err == nil {
		err := ioutil.WriteFile("./testfiles/test.mp3", b, 0644)
		if err != nil {
			log.Println(err.Error())
		}
	}*/

	res.Audio = b
	res.Error = nil

	return res
}
