package mux

import (
	"errors"
	"regexp"
	"slices"
	"strings"
)

type tree struct {
	regPaths []regPathNode
}

func New() *tree {
	return &tree{
		regPaths: []regPathNode{},
	}
}

func (t *tree) Add(node regPathNode) {
	found := false
	for i, v := range t.regPaths {
		if node.level > v.level || (node.level == v.level && len(node.paramNames) < len(v.paramNames)) {
			t.regPaths = slices.Insert(t.regPaths, i, node)
			found = true
			break
		}
	}
	if !found {
		t.regPaths = append(t.regPaths, node)
	}
}

func (t *tree) Find(url string, method string) (EndpointHandler, map[string]string, error) {
	var (
		found     *regPathNode
		paramVals []string
	)

	// Dla każdej zarejestrowanej ścieżki
	for _, regPath := range t.regPaths {

		// Utwórz z niej regex
		regex := regexp.MustCompile("^" + regPath.path + "/*$")

		// Sprawdź, czy pasuje do regexu
		if regex.MatchString(url) {
			paramVals = regex.FindAllStringSubmatch(url, -1)[0][1:]
			found = &regPath
			break
		}
	}

	if found != nil {
		if handler, ok := found.handlers[method]; ok {
			params := map[string]string{}
			for i, paramName := range found.paramNames {
				params[paramName] = paramVals[i]
			}

			return handler, params, nil
		}
	}

	return nil, nil, errors.New("nie znaleziono")
}

type regPathNode struct {
	path       string
	handlers   map[string]EndpointHandler
	paramNames []string
	level      int
}

func NewNode(path string, handlers map[string]EndpointHandler) regPathNode {
	escapedPath := regexp.QuoteMeta(path)
	regex := regexp.MustCompile(`\\{([a-z]+)\\}`)

	level := len(strings.Split(path, "/"))
	matches := regex.FindAllStringSubmatch(escapedPath, -1)
	paramNames := []string{}
	for _, paramName := range matches {
		paramNames = append(paramNames, paramName[1])
	}

	regexifiedPath := regex.ReplaceAllLiteralString(escapedPath, "([[:alnum:]]+)")

	return regPathNode{
		path:       regexifiedPath,
		handlers:   handlers,
		level:      level,
		paramNames: paramNames,
	}

}
