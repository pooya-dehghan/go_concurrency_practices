package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "https://isna.ir", nil)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), 100*time.Microsecond)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("ERR: ", err)
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
