package storage

import "context"

// Storage inteface defines methods for storing to and retrieving from storage.
type Storage interface {
	StoreCharToRune(ctx context.Context, s string, r []uint32) error
	GetCharToRune(ctx context.Context, s string) ([]uint32, error)

	StoreRuneToChar(ctx context.Context, r []byte, s string) error
	GetRuneToChar(ctx context.Context, r []byte) (string, error)

	// later also char to bytes and visa versa
}
