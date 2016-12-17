package parse

import (
	"bufio"
	"os"
)

//Response . . . response struct from all Parse functions, contains parsed data in long string and Error if one here
type Response struct {
	Data  string
	Error error
}

//Parse . . . Parse method implementation for .txt file parser
func (p *TxtParser) Parse(path string) Response {
	var res Response
	file, err := os.Open(path)
	if err != nil {
		res.Error = err
		return res
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res.Data += " " + scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		res.Error = err
		return res
	}

	return res
}
