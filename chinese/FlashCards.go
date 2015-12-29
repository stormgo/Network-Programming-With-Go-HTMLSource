package flashcard

import (
	"dictionary"
	"container/vector"
)

type FlashCard struct {
	English string
	Card    DictionaryEntry
}

type FlashCards struct {
	vector.Vector
}
