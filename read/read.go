package read

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/kc1116/go-read2me/parse"
)

//GetVoices . . . returns set of voices available in AWS Polly
func GetVoices(lang string) {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := polly.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	params := &polly.DescribeVoicesInput{
		LanguageCode: aws.String(lang),
	}
	resp, err := svc.DescribeVoices(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

//Read . . . Takes data input converts it SynthesizedSpeechOutput
func Read(parsed parse.Response, voiceID string) {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := polly.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	params := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		Text:         aws.String(parsed.Data),
		VoiceId:      aws.String(voiceID),
		TextType:     aws.String("text"),
	}
	resp, err := svc.SynthesizeSpeech(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
	b, err := ioutil.ReadAll(resp.AudioStream)
	if err == nil {
		err := ioutil.WriteFile("./testfiles/test.mp3", b, 0644)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
