package engine

import (
	"fmt"
	"strings"
)

type Scanner interface {
	Scan() (map[string]string, error)
}

func Search(s Scanner, words string) ([]string, error) {
	var found []string
	dic, err := s.Scan()
	if err != nil {
		return found, err
	}
	for k, v := range dic {
		if strings.Contains(lower(k), lower(words)) || strings.Contains(lower(v), lower(words)) {
			found = append(found, fmt.Sprintf("%s - '%s'\n", k, v))
		}
	}
	return found, nil
}

func lower(str string) string {
	strLower := strings.ToLower(str)
	return strLower
}
