package main

import (
	"fmt"
	"io/ioutil"
	"strings"
    "os"
    "math/rand"
    "time"
)

// Creates a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
    cards := deck{}

    cardSuits := [4] string {"Spades", "Diamonds", "Hearts", "Clubs"}
    cardValues := [11] string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Queen", "King"}

    for _, suit := range cardSuits {
        for _, value := range cardValues {
            cards = append(cards, value + " of " + suit)
        }
    }

    return cards;
}

// Receiver function
// reference uses the first 1 or 2 letters of the type usually
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}

func (d deck) toString() string {
    return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(name string) error {
    return ioutil.WriteFile(name, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
    s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)

    for i := range d {
        newPosition := r.Intn(len(d) - 1)

        d[i], d[newPosition] = d[newPosition], d[i]
    }
}

func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}

func newDeckFromFile(name string) deck {
    bs, err := ioutil.ReadFile(name)

    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    s := strings.Split(string(bs), ", ")
    return deck(s)
}
