package index

import "strings"

type Service struct {
	words map[string][]int
}

// New - констрктор индекс.
func New() *Service {
	i := Service{}
	i.words = make(map[string][]int)
	return &i
}

func (i *Service) Add(s string, num int) {
	words := strings.Fields(strings.ToLower(s))
	for _, w := range words {
		i.words[w] = append(i.words[w], num)
	}
}

func (i Service) Search(s string) []int {
	return i.words[strings.ToLower(s)]
}
