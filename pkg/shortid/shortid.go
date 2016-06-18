package shortid

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/renstrom/shortuuid"
	"github.com/satori/go.uuid"
)

// Base58Encoder represents a Base58 encoder used in bitcoin network
type Base58Encoder struct{}

// Encode encodes UUID as Base58 string
func (enc Base58Encoder) Encode(u uuid.UUID) string {
	return base58.Encode(u.Bytes())
}

// Decode decodes Base58 string into UUID
func (enc Base58Encoder) Decode(s string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode(s))
}

// Generate creates new Base58 ID from entropy
func Generate() string {
	return shortuuid.NewWithEncoder(Base58Encoder{})
}
