package mux

import (
	"errors"
	"regexp"
	"slices"
	"strings"
)

type Tree struct {
	regPaths []regPathNode
}

func NewTree() *Tree {
	return &Tree{
		regPaths: []regPathNode{},
	}
}

func (t *Tree) Add(path string, method string, handler EndpointHandler) error {
	nodeToInsert := NewNode(path, map[string]EndpointHandler{
		method: handler,
	})

	for index, node := range t.regPaths {
		if node.path == path {
			if node.handlers[method] != nil {
				return errors.New("method already exist")
			}
			node.handlers[method] = handler
			return nil
		}
		if nodeToInsert.level > node.level ||
			(nodeToInsert.level == node.level && len(nodeToInsert.paramNames) < len(node.paramNames)) {
			t.regPaths = slices.Insert(t.regPaths, index, nodeToInsert)
			return nil
		}
	}

	t.regPaths = append(t.regPaths, nodeToInsert)
	return nil
}

func (t *Tree) Find(path string, method string) (EndpointHandler, map[string]string, error) {
	for _, node := range t.regPaths {
		preparedPath := regexp.MustCompile("^" + node.path + "/*$")
		if preparedPath.MatchString(path) {
			handler, ok := node.handlers[method]
			if !ok {
				return nil, nil, errors.New("method not registered")
			}

			paramValues := preparedPath.FindAllStringSubmatch(path, -1)[0][1:]

			paramMap := map[string]string{}

			for index, paramName := range node.paramNames {
				paramValue := paramValues[index]
				paramMap[paramName] = paramValue
			}

			return handler, paramMap, nil
		}
	}

	return nil, nil, errors.New("path not found")
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
