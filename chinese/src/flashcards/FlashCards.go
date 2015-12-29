package flashcards

import (
	"bufio"
	"bytes"
	"dictionary"
	"encoding/json"
	"fmt"
	"os"
)

import (
	"sort"
	"strings"
)

type FlashCard struct {
	Simplified string
	English    string
	Dictionary *dictionary.Dictionary
}

type FlashCards struct {
	Name      string
	CardOrder string
	ShowHalf  string
	Cards     []*FlashCard
}

func (cards *FlashCards) addCard(card *FlashCard) {
	cardsArray := cards.Cards
	len := len(cardsArray)

	newCardsArray := make([]*FlashCard, len+1)
	copy(newCardsArray, cardsArray)
	cards.Cards = newCardsArray
	cards.Cards[len] = card
}

func mkCard(word, pinyin, simplified,
	traditional, translations string) *FlashCard {
	card := new(FlashCard)

	entryStr :=
		`{
                           "Traditional": "` + traditional + `", 
                           "Simplified": "` + simplified + `", 
                           "Pinyin": "` + pinyin + `", 
                           "Translations": ` + translations +
			`}`
	entriesStr := `"Entries": [` + entryStr + `]`

	flashcardStr := `{"English": "` + word + `", 
                          "Simplified": "` + simplified + `", 
                          "Dictionary": {` + entriesStr + `}
                         }`

	fmt.Println(flashcardStr)
	buff := bytes.NewBufferString(flashcardStr)
	decoder := json.NewDecoder(buff)
	err := decoder.Decode(card)
	if err != nil {
		fmt.Println("Error " + err.Error())
		return nil
	} else {
		fmt.Println("Ok ")
	}
	return card
}

func (cards *FlashCards) Load(path string, d *dictionary.Dictionary) {
	fields := strings.Split(path, `/`)
	cards.Name = fields[len(fields)-1]

	println("Path is " + path)
	f, err := os.Open(path)
	r := bufio.NewReader(f)
	if err != nil {
		println(err.Error())
		return
	}

	v := make([]*FlashCard, 0, 100)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if line[0] == '#' {
			continue
		}

		println("Line in dict is", line)
		words := strings.Fields(line)
		simp := words[0]
		english := strings.Join(words[1:], ` `)
		dict := d.LookupSimplified(simp)
		if len(dict.Entries) > 1 {
			dict = dict.LookupEnglish(english)
		}
		println("Num entries for ", english, len(dict.Entries))
		for _, e := range dict.Entries {
			println("   ", e.Pinyin, e.Traditional, e.Simplified)
		}
		if len(dict.Entries) > 1 {
			dict.Entries = dict.Entries[0:1]
		}
		card := FlashCard{
			Simplified: simp,
			English:    english,
			Dictionary: dict}

		v = append(v, &card)
	}
	cards.Cards = v
}

func LoadJSON(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}

func SaveJSON(fileName string, obj interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(obj)
	checkError(err)
	outFile.Close()
}

func ListFlashCardsNames() []string {
	flashcardsDir, err := os.Open("flashcardSets")
	if err != nil {
		return nil
	}
	files, err := flashcardsDir.Readdir(-1)

	fileNames := make([]string, len(files))
	for n, f := range files {
		fileNames[n] = f.Name()
	}
	sort.Strings(fileNames)
	return fileNames
}

func NewFlashCardSet(newSet string) {
	f, _ := os.OpenFile("flashcardSets/"+newSet, os.O_CREATE|os.O_EXCL, 0666)
	f.Close()
}

func GetFlashCardsByName(name string, d *dictionary.Dictionary) *FlashCards {
	fc := new(FlashCards)
	fc.Load("flashcardSets/"+name, d)
	// fill name if possible
	return fc
}

func AddFlashEntry(cardName, word, pinyin, simplified,
	traditional, translations string) {

	cards := new(FlashCards)
	LoadJSON("flashcardSets/"+cardName, &cards)
	card := mkCard(word, pinyin, simplified,
		traditional, translations)
	cards.addCard(card)
	SaveJSON("flashcardSets/"+cardName, &cards)

	f, err := os.OpenFile("flashcardSets/"+cardName,
		os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		println(err.Error())
		return
	}
	f.Write([]byte(simplified + ` ` + word + "\n"))
	f.Close()
	println("Added entry", cardName, word, simplified)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
}
