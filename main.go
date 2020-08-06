package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "time"
    "math/rand"
    "os"
)

type topAFIQuote struct {
    Rank int
    Quote string
    Movie string
    Year int
}

type quoteHandler []topAFIQuote

var quotes quoteHandler = make([]topAFIQuote, 100)


func getRandomRank(max int) int {
    return rand.Intn(max)
}


func checkError(e error) {
    if e != nil {
        panic(e)
    }
}


func cacheQuotes(filepath string) {
    data, err := ioutil.ReadFile(filepath)
    checkError(err)
    json.Unmarshal(data, &quotes)
}


func (qh *quoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    rank := getRandomRank(len(*qh))
    data := (*qh)[rank]
    fmt.Fprintf(w, "#%d: %s - %s (%d)", data.Rank, data.Quote, data.Movie, data.Year)
}


func main() {
    rand.Seed(time.Now().UnixNano())

    quotespath := os.Getenv("GOAFI")
    cacheQuotes(quotespath)
    http.Handle("/randomquote", &quotes)
    http.ListenAndServe(":8000", nil)
}
