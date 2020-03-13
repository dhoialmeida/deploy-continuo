package main

import (
	"fmt"
	"log"
	"net/http"
	"math"
)

func NegritoSqrt(texto string) string {
	x := 0.0
	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(float64(i));
	}
	return "<b>" + texto + "</b>"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, NegritoSqrt("Code.education rocks!"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
