# go-read2me

These set of go packages utilize the power of [Google's translate API ](https://cloud.google.com/translate/) for language translation 
combined with the power of [Amazon AWS Polly API](https://aws.amazon.com/polly/) for text to speech. Together these API's 
can be used to; enter some text --> translate that text --> 'speak' that translated text

## Overview of directory structure 

### Package api

This package serves as helper methods for the request handlers that bring parse and synthesize methods together in the right order.

### Package configs 
This package is used to set any configurations. API keys, server ports, anything that may be needed for other packages. 

### Package parse 
This package contains a parser interface in parse.go . Because parsing of differen't file formats may be implemented 
in various different ways for different reasons you should add your parse method to it's own file and have and 
implement the parse method for that particular file type. i.e for CSV's create a file called csv.go and then implement 
the parser interfaces parse method. 

### Package routing
This package contains http routes and request handlers. 

### Package synthesize
This package contains various methods used to synthesize input with AWS polly API.

### Package translate
This package will allow translation of text before synthesizing it.

