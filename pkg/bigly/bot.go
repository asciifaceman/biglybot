package bigly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	regomarkov "github.com/asciifaceman/regomarkov"
)

// Some matchers
const (
	ErrNgramLength = "N-gram length does not match chain order"
)

// Bot is our core bot struct
type Bot struct {
	Filepath string
	Brain    *regomarkov.Chain
}

// Train adds new training data to the bigly brain
func (b *Bot) Train(set [][]string) error {
	if len(set) < 1 {
		return fmt.Errorf("Set is too short to train")
	}
	totalLength := len(set)
	failedCount := 0
	for _, data := range set {
		if strings.Join(data, "") == "" {
			continue
		}

		if len(data) < 2 {
			//fmt.Printf("Data line [%d]: [%s] is too short to train", i, data)
			failedCount++
			continue
		}
		b.Brain.Add(data)
	}
	log.Printf("[%d of %d] records processed successfully", totalLength-failedCount, totalLength)

	return nil
}

// Speak returns a generated sentence from the chain
func (b *Bot) Speak() string {
	var spoken string
	//resp, err := b.Brain.GenerateAll()
	resp, err := b.Brain.GenerateAll()
	if err != nil {
		spoken = err.Error()
	}
	spoken = strings.Join(resp, " ")

	return spoken
}

// Save ...
func (b *Bot) Save() error {
	jsonObj, err := json.Marshal(b.Brain)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(b.Filepath, jsonObj, 0644)
	return err
}

// Load  ...
func Load(fp string) (*Bot, error) {
	if !IsValid(fp) {
		return nil, fmt.Errorf("Filepath (%s) is invalid or exists", fp)
	}

	f, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	brain := regomarkov.NewChain(2)
	err = brain.UnmarshalJSON(f)
	b := &Bot{
		Filepath: fp,
		Brain:    brain,
	}

	return b, err
}

// New ...
func New(fp string) (*Bot, error) {
	if !IsValid(fp) {
		return nil, fmt.Errorf("Filepath (%s) is invalid or exists", fp)
	}

	b := &Bot{
		Filepath: fp,
		Brain:    regomarkov.NewChain(2),
	}
	return b, nil
}
