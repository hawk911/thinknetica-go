package crawbot

// Scanner - имитация служба поискового робота.
type Scanner struct{}

// NewScanner - констрктор имитации службы поискового робота.
func New() *Scanner {
	s := Scanner{}
	return &s
}

// Scan возвращает заранее подготовленный набор данных
func (s *Scanner) Scan() (data map[string]string, err error) {

	data = map[string]string{
		"http://habr.ru/":    " There are some items.",
		"http://ya.ru/":      " I do not know it.",
		"http://golang.org/": "I am thinking about you!",
	}

	return data, nil
}
