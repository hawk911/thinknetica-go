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
		"habr1": " There are some items.",
		"habr2": " I do not know it.",
		"habr3": "I am thinking about you!",
	}

	return data, nil
}
