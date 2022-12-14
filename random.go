package mockmongo

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// DBNameLen is the length of a database name generated by RandomDatabase().
// It's OK to change this, but not concurrently with calls to RandomDatabase.
const DBNameLen = 15

// DBNameChars is the set of characters used by RandomDatabase().
// It's OK to change this, but not concurrently with calls to RandomDatabase.
const DBNameChars = "abcdefghijklmnopqrstuvwxyz"

// RandomDatabase returns a random valid mongo database name. You can use to
// to pick a new database name for each test to isolate tests from each other
// without having to tear down the whole server.
//
// This function will panic if it cannot generate a random number.
func RandomDatabase() string {
	dbChars := make([]byte, DBNameLen)
	for i := 0; i < DBNameLen; i++ {
		bigN, err := rand.Int(rand.Reader, big.NewInt(int64(len(DBNameChars))))
		if err != nil {
			panic(fmt.Errorf("error getting a random int: %s", err))
		}

		dbChars[i] = DBNameChars[int(bigN.Int64())]
	}

	return string(dbChars)
}
