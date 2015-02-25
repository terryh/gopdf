package gopdf

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
)

var (
	wkhtmltopdfCmd   = "wkhtmltopdf"   // the command to call wkhtmltopdf command
	wkhtmltoimageCmd = "wkhtmltoimage" // the command to call wkhtmltoimage command

	haveWKPdf   = "Looks like you have wkhtmltopdf install at your system path."
	haveWKImage = "Looks like you have wkhtmltoimage install at your system path."
	nothaveWK   = `Looks you don't have wkhtmltopdf install at your system,
please take a look on http://wkhtmltopdf.org/downloads.html`

	wkpdfcheck   = false
	wkimagecheck = false
)

func SetWkhtmltopdf(cmd string) error {
	wkhtmltopdfCmd = cmd
	return CheckWkhtmltopdf(cmd)
}

func SetWkhtmltoimage(cmd string) error {
	wkhtmltoimageCmd = cmd
	return CheckWkhtmltoimage(cmd)
}

func CheckWkhtmltopdf(cmd string) error {

	fmt.Println("Checking wkhtmltopdf..")

	_, err := exec.Command(cmd, "-V").Output()

	if err != nil {
		fmt.Println(nothaveWK)
		return nil
	}

	fmt.Println(haveWKPdf)
	// seems haaving wkhtmltopdf
	wkpdfcheck = true
	return nil
}

func CheckWkhtmltoimage(cmd string) error {

	fmt.Println("Checking wkhtmltoimage..")

	_, err := exec.Command(cmd, "-V").Output()

	if err != nil {
		fmt.Println(nothaveWK)
		return nil
	}

	fmt.Println(haveWKImage)
	// seems haaving wkhtmltopdf
	wkimagecheck = true
	return nil
}

func init() {
	CheckWkhtmltopdf(wkhtmltopdfCmd)
	CheckWkhtmltoimage(wkhtmltoimageCmd)
}

// Url2pdf is mapping wkhtmltopdf url with args
// and have result from stdout wrap in []byte
func Url2pdf(url string, args ...string) ([]byte, error) {
	if !wkpdfcheck {
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

// Url2image is mapping wkhtmltoimage url with args
// and have result from stdout wrap in []byte
// the default format is JPEG
func Url2jpeg(url string, args ...string) ([]byte, error) {
	if !wkimagecheck {
		return nil, errors.New(nothaveWK)
	}

	cmdSlice := []string{}

	for _, line := range args {
		cmdSlice = append(cmdSlice, line)
	}

	cmdSlice = append(cmdSlice, url)

	// pipe to stdout
	cmdSlice = append(cmdSlice, "-")

	var out bytes.Buffer

	cmd := exec.Command(wkhtmltoimageCmd, cmdSlice...)
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	return out.Bytes(), err

}
