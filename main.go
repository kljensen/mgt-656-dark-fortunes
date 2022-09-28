package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

var darkFortunes = []string{
	"you will die alone and poorly dressed",
	"you will be haunted by the ghosts of everyone you have ever wronged",
	"you will be forced to eat your words and like it",
	"you will be struck by lightning",
	"you will be run over by a reindeer",
	"you will be trapped in a burning building",
}

var goodFortunes = []string{
	"Today it's up to you to create the peacefulness you long for.",
	"A friend asks only for your time not your money.",
	"If you refuse to accept anything but the best, you very often get it.",
	"A smile is your passport into the hearts of others.",
	"A good way to keep healthy is to eat more Chinese food.",
	"Your high-minded principles spell success.",
	"Hard work pays off in the future, laziness pays off now.",
	"Change can hurt, but it leads a path to something better.",
	"Enjoy the good luck a companion brings you.",
	"People are naturally attracted to you.",
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/fortune", FortuneHandler)
	http.ListenAndServe(":8080", nil)
}
func FortuneHandler(w http.ResponseWriter, r *http.Request) {
	whichList := rand.Intn(2)
	fortunes := goodFortunes
	if whichList != 0 {
		fortunes = darkFortunes
	}
	index := rand.Intn(len(fortunes))
	fortune := fortunes[index]
	htmlBody := `
        <html>
            <title>Fortune</title>
            <head>
                <style>
                    body {
                        background-color: #131efe;
                        color: #fff;
                    }
                    p {
                        color: #000;
                        background-color: blue;
                    }
                </style>
            </head>
            <body>
                <h1>Fortune</h1>
                <p>%s</p>
            </body>
        </html>
    `
	fmt.Fprintf(w, htmlBody, fortune)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
