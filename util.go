package multipass

import "bytes"

// buildString comment
func buildString(parts ...string) string {
	var buf bytes.Buffer
	for _, val := range parts {
		buf.WriteString(val)
	}
	return buf.String()
}

func pkcs5Padding(cipherBytes []byte, blockSize int) []byte {
	padding := blockSize - (len(cipherBytes) % blockSize)
	padBytes := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherBytes, padBytes...)
}
