package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Quote struct {
	Text   string
	Author string
}

func main() {
	// Отправляем GET-запрос на веб-сайт "https://quotes.toscrape.com"
	resp, err := http.Get("https://quotes.toscrape.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Загружаем HTML-страницу для парсинга
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Извлекаем цитаты
	var quotes []Quote
	doc.Find(".quote").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Find(".text").Text())
		author := strings.TrimSpace(s.Find(".author").Text())

		quote := Quote{
			Text:   text,
			Author: author,
		}
		quotes = append(quotes, quote)
	})

	// Выводим результат
	for _, q := range quotes {
		fmt.Printf("Цитата: %s\nАвтор: %s\n\n", q.Text, q.Author)
	}
}
