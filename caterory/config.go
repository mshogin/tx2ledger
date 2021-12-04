package category

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Categories map[string]map[string]string

func CreateParser(r io.Reader) (Categories, error) {
	dec := json.NewDecoder(r)
	c := Categories{}
	if err := dec.Decode(&c); err != nil {
		return nil, fmt.Errorf("cannot decode config: %w", err)
	}
	return c, nil
}

func (m Categories) Parse(details ...string) (string, string, error) {
	for category, pats := range m {
		for pat, name := range pats {
			for _, detail := range details {
				if strings.Contains(strings.ToLower(detail), strings.ToLower(pat)) {
					return name, category, nil
				}
			}
		}
	}
	fmt.Printf("details = %+v\n", details) // output for debug
	return "", "", fmt.Errorf("cannot find Category")
}
