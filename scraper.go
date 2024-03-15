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
	// Определяем количество страниц для парсинга
	numPages := 5

	var quotes []Quote

	// Парсим каждую страницу
	for page := 1; page <= numPages; page++ {
		// Формируем URL с номером страницы
		url := fmt.Sprintf("https://quotes.toscrape.com/page/%d", page)

		// Отправляем GET-запрос на веб-сайт
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		// Закрываем тело ответа
		defer resp.Body.Close()

		// Загружаем HTML-страницу для парсинга
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Извлекаем цитаты
		doc.Find(".quote").Each(func(i int, s *goquery.Selection) {
			text := strings.TrimSpace(s.Find(".text").Text())
			author := strings.TrimSpace(s.Find(".author").Text())

			quote := Quote{
				Text:   text,
				Author: author,
			}
			quotes = append(quotes, quote)
		})
	}

	// Выводим результат
	for _, q := range quotes {
		fmt.Printf("Цитата: %s\nАвтор: %s\n\n", q.Text, q.Author)
	}
}
