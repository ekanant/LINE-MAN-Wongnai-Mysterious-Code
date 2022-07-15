package main

import (
	"encoding/base64"
	"log"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		log.Fatalln("decode message error", err)
	}
	whatIsIt = string(sd)
	whatIsIt = reverseString(whatIsIt)
	log.Printf("The secret is %s\n", whatIsIt)
}

func reverseString(txt string) string {
	reverseRune := make([]byte, len(txt))

	left := 0
	right := len(txt) - 1
	for left <= right {
		reverseRune[left] = txt[right]
		reverseRune[right] = txt[left]
		right--
		left++
	}
	return string(reverseRune)
}
