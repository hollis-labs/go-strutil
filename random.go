package strutil

import (
	"crypto/rand"
	"math/big"
)

const randomAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// Random returns a cryptographically secure random alphanumeric string
// of the given length. It uses crypto/rand as the source of entropy and
// draws from the charset [A-Za-z0-9]. A length of 0 or less returns "".
// Panics only if the system's crypto/rand is unavailable — a condition
// that already makes the process unable to function securely.
//
//	Random(8) → e.g. "xK2mP9qR"
//	Random(0) → ""
func Random(length int) string {
	if length <= 0 {
		return ""
	}
	max := big.NewInt(int64(len(randomAlphabet)))
	out := make([]byte, length)
	for i := range out {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			// crypto/rand failure is not recoverable at the API boundary.
			panic("strutil: crypto/rand unavailable: " + err.Error())
		}
		out[i] = randomAlphabet[n.Int64()]
	}
	return string(out)
}
