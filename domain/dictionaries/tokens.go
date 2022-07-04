package dictionaries

import (
	"errors"
	"fmt"
)

type tokens struct {
	list []Token
	mp   map[string]Token
}

func createTokens(
	list []Token,
	mp map[string]Token,
) Tokens {
	out := tokens{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the tokens
func (obj *tokens) List() []Token {
	return obj.list
}

// Find finds a token by name
func (obj *tokens) Find(name string) (Token, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the token (name: %s) is undeclared", name)
	return nil, errors.New(str)
}
