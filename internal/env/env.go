package env

import (
	"log"
	"strings"
	"text/scanner"

	"github.com/nakaoni/sebas/internal/util"
)

type Env struct {
	Key   string
	Value string
}

func ReadFile(path string) ([]Env, error) {
	content, err := util.GetFileContent(path)

	if err != nil {
		log.Panicln(err)
	}

	tokens, err := GetTokens(content)

	if err != nil {
		log.Panicln(err)
	}

	envs := make([]Env, 0)
	for k, v := range tokens {
		envs = append(envs, *ConvertToEnv(k, v))
	}

	return envs, nil
}

func GetTokens(envs string) (map[string]string, error) {
	tokens := make(map[string]string)
	var key string
	var value string
	var valueBuilder strings.Builder
	var lastAddTokenToValueBuilder string
	foundEqual := false

	var s scanner.Scanner
	s.Init(strings.NewReader(envs))

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()

		switch txt {
		case "=":
			if foundEqual {
				newKey := lastAddTokenToValueBuilder
				value, _ = strings.CutSuffix(valueBuilder.String(), newKey)

				tokens[key] = value
				if newKey != "" {
					key = newKey
				}

				valueBuilder.Reset()
				value = ""
				lastAddTokenToValueBuilder = ""
			} else {
				foundEqual = true
			}
			continue
		}

		if foundEqual {
			lastAddTokenToValueBuilder = txt
			valueBuilder.WriteString(txt)
		} else {
			key = txt
		}
	}

	tokens[key] = valueBuilder.String()

	return tokens, nil
}

func ConvertToEnv(key string, value string) *Env {
	return &Env{Key: key, Value: value}
}
