// Package crawler реализует сканер содержимого веб-сайтов.

// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.

package crawler

import (
	"testing"
)

func TestScan(t *testing.T) {
	const url = "https://habr.com"
	const depth = 2
	data, err := Scan(url, depth)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range data {
		t.Logf("%s -> %s\n", k, v)
	}
}
