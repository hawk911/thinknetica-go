// Package crawler реализует сканер содержимого веб-сайтов.

// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.

package crawler

import (
	"testing"
)

func TestScan(t *testing.T) {

	const url = "https://habr.com"
	const depth = 2
	s := Param{Url: url, Depth: depth}
	data, err := s.Scan()
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range data {
		t.Logf("%s -> %s\n", k, v)
	}
}
