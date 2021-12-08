package command

import (
	"bufio"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(b []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(b)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(b)
	}
	return str
}

func CommonExecute(command string) string {
	result := ""
	cmd := exec.Command("cmd", "/c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return ""
	}
	cmd.Start()
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		result += ConvertByte2String(in.Bytes(), "GB18030") + "\n"
	}
	cmd.Wait()
	return result
}
