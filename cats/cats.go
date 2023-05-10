package cats

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Download() (string, error) {
	out, err := os.Create(fmt.Sprintf("%v.png", RandStringBytes(16)))
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get("https://cataas.com/cat")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return out.Name(), nil
}
