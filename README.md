## gppdf
golang wrap for wkhtmltopdf command.

## Example

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/terryh/gopdf"
)

func Pdfhandle(w http.ResponseWriter, r *http.Request) {
	result, err := gopdf.Url2pdf("http://nvd3.org/examples/stackedArea.html")
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(result)
}

func Jpeghandle(w http.ResponseWriter, r *http.Request) {
	result, err := gopdf.Url2jpeg("http://nvd3.org/examples/stackedArea.html")
	fmt.Println(err)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(result)
}

func main() {

	http.HandleFunc("/", Pdfhandle)
	http.HandleFunc("/jpeg", Jpeghandle)

	http.ListenAndServe(":8080", nil)

}
```




