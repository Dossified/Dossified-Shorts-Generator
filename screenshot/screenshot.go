package screenshot

import (
    "io"
    "os"
	"bytes"
	"net/http"
)

func ScreenshotWebsite() {

    posturl := "http://127.0.0.1:7171/api/screenshot"
    body := []byte(`{
        "url": "http://127.0.0.1:8000/vid_gen/?item_id=132&obj_type=news",
        "oneshot": "true"
    }`)

    // Create a HTTP post request
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
    r.Header.Add("Content-Type", "application/json")
    client := &http.Client{}
    res, err := client.Do(r)
    if err != nil {
        panic(err)
    }

    defer res.Body.Close()

    file, err := os.Create("output.png")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    _, err = io.Copy(file, res.Body)
    if err != nil {
        panic(err)
    }
}
