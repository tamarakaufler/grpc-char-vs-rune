package storage

// Storage inteface defines methods for storing to and retrieving from storage.
type Storage interface {
	StoreCharToRune(charToRune string) error
	RetrieveCharToRuneConversion(charToRune string) ([]rune, error)

	StoreRuneToChar(runeToChar string) error
	RetrieveRuneToCharToConversion(runeToChar []rune) (string, error)

	// later also char to bytes and visa versa
}
