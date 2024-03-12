package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

var emojis = map[string]string{
	"happy cheerful laugh": "😊",
	"sad unhappy crying":   "😢",
	"love heart emotion":   "❤️",
	"thumbs ":              "👍",
	"cat ":                 "😺",
	"dog ":                 "🐶",
	"sun":                  "☀️",
	"moon":                 "🌙",
	"star":                 "⭐",
	"rocket":               "🚀",
	"coffee":               "☕",
	"pizza":                "🍕",
	"heart eyes":           "😍",
}

// translateToEmoji translate text to matching emoji
func translateToEmoji(input string) (string, error) {

	Input := strings.TrimSpace(input)
	inputLowerCase := strings.ToLower(Input)

	for eN, emoji := range emojis {

		if eN == inputLowerCase || strings.Contains(eN, inputLowerCase) {
			return emoji, nil
		}
	}

	return "", errors.New("no match found")
}

// translator handles the translation of input text to emoji.
func translator(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		inputText := req.FormValue("inputText")

		emoji, err := translateToEmoji(inputText)
		if err != nil {
			log.Println("Error:", err)
		}

		tmpl.ExecuteTemplate(res, "index.gohtml", map[string]string{"InputText": inputText, "Emoji": emoji})
		return
	}
	// Execute render the template without displaying the emoji
	tmpl.Execute(res, nil)
}

// index handles the rendering of the index page.
func index(res http.ResponseWriter, req *http.Request) {
	err := tmpl.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Println("template didn't execute", nil)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/translate", translator)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)

}
