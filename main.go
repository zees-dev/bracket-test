package main

import (
	"errors"
	"log"
	"sync"
)

func main() {
	// implement a stack
	// stack pop looks at the top of stack and if invalid match, return false
	// if stack is empty, -> success (return true)
	input := "[[(){}]]"
	log.Println(validateBrackets(input))
}

func validateBrackets(bracketStr string) bool {
	bracketStack := NewBracketStack()
	for _, c := range bracketStr {
		success, err := bracketStack.tryPushPop(c)
		if err != nil {
			log.Panicln(err)
		}

		if !success {
			return false
		}
	}

	// stack is not empty, hence string is not valid
	if len(bracketStack.data) != 0 {
		return false
	}

	return true
}

type BracketStack struct {
	data       []rune
	bracketMap map[rune]rune
	mu         sync.RWMutex
}

func NewBracketStack() *BracketStack {
	return &BracketStack{
		data: []rune{},
		bracketMap: map[rune]rune{
			'{': '}',
			'[': ']',
			'(': ')',
		},
	}
}

func (b *BracketStack) tryPushPop(char rune) (bool, error) {
	for k, v := range b.bracketMap {
		// If incoming char is open bracket, add to stack
		if char == k {
			// push item and return
			b.data = append(b.data, char)
			return true, nil
		}

		// If incoming char is closing bracket
		if char == v {
			if len(b.data) == 0 {
				return false, nil
			}

			// if key for this closing brace (opening brace) is not top of stack
			if k != b.data[len(b.data)-1] {
				return false, nil
			}

			// pop last item and return
			b.data = b.data[:len(b.data)-1]
			return true, nil
		}
	}

	return false, errors.New("Invalid bracket type")
}
