package gopdf

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
)

var (
	wkhtmltopdfCmd = "wkhtmltopdf" // the command to call wkhtmltopdf command

	haveWK    = "Looks like you have wkhtmltopdf install at your system path."
	nothaveWK = `Looks you don't have wkhtmltopdf install at your system,
please take a look on http://wkhtmltopdf.org/downloads.html`
	wkcheck = false
)

func SetWkhtmltopdf(cmd string) error {
	wkhtmltopdfCmd = cmd
	return CheckWkhtmltopdf(cmd)
}

func CheckWkhtmltopdf(cmd string) error {

	fmt.Println("Checking wkhtmltopdf..")

	_, err := exec.Command(cmd, "-V").Output()

	if err != nil {
		fmt.Println(nothaveWK)
		return nil
	}

	fmt.Println(haveWK)
	// seems haaving wkhtmltopdf
	wkcheck = true
	return nil
}

func init() {
	CheckWkhtmltopdf(wkhtmltopdfCmd)
}

// Url2pdf is mapping wkhtmltopdf url with args
// and have result from stdout wrap in []byte
func Url2pdf(url string, args ...string) ([]byte, error) {
	if !wkcheck {
		return nil, errors.New(nothaveWK)
	}

	cmdSlice := []string{url}

	for _, line := range args {
		cmdSlice = append(cmdSlice, line)
	}

	// pipe to stdout
	cmdSlice = append(cmdSlice, "-")

	//fmt.Println(wkhtmltopdfCmd, cmdSlice)

	var out bytes.Buffer

	cmd := exec.Command(wkhtmltopdfCmd, cmdSlice...)
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	return out.Bytes(), err

}
