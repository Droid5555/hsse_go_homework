package tools

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"hsse_go_homework/task1/pkg/book"
	"strconv"
)

func HashGen1(b book.Book) string {
	data := b.Title + b.Author
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:]) + "HASH1"
}

func HashGen2(b book.Book) string {
	data := b.Title + b.Author
	hash := sha256.Sum256([]byte(data))
	hashInt := int(binary.BigEndian.Uint64(hash[:8]))
	return strconv.Itoa(hashInt) + "HASH2"
}
