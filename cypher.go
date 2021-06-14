package main

import (
	"fmt"
	"strings"
)

func main() {

	//Vizener cipher encryption
	message := "your message goes here"
	keyword := "golang"
	keyIndex := 0
	cipherText := ""

	message = strings.ToUpper(strings.Replace(message, " ", "", -1))
	keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'A' && c <= 'Z' {
			// A=0, B=1, ... Z=25
			c -= 'A'
			k := keyword[keyIndex] - 'A'
			c = (c+k)%26 + 'A'

			keyIndex++
			keyIndex %= len(keyword)
		}
		cipherText += string(c)
	}
	fmt.Println(cipherText)

	/*
	   //Vizener cipher decryption
	   cipherText := "CSOITEUIWUIZNSROCNKFD"
	   keyword := "GOLANG"
	   message := ""
	   keyIndex := 0

	   for i := 0; i < len(cipherText); i++ {
	       // A=0, B=1, ... Z=25
	       c := cipherText[i] - 'A'
	       k := keyword[keyIndex] - 'A'

	       // зашифрованная буква - ключевая буква
	       c = (c-k+26)%26 + 'A'
	       message += string(c)

	       // увеличение keyIndex
	       keyIndex++
	       keyIndex %= len(keyword)
	   }

	   fmt.Println(message)

	   //code ROT13 (--> 13 symbols)
	     text := "Hola Estación Espacial Internacional"

	     for _, c := range text {
	       if c >= 'a' && c <= 'z' {
	           c += 13
	           if c > 'z' {
	               c -= 26
	           }
	       } else if c >= 'A' && c <= 'Z' {
	         c +=13
	         if c > 'Z' {
	           c -= 26
	         }
	       }
	       fmt.Printf("%c", c)
	   }

	   //decode Cesar code -3
	     text := "L fdph, L vdz, L frqtxhuhg."

	     for i:= 0; i < len(text); i++ {
	       c := text[i]
	       if c >= 'a' && c <= 'z' {
	         c -= 3
	         if c < 'a' {
	           c += 26
	         }
	       } else if c >= 'A' && c <= 'Z' {
	         c -= 3
	         if c < 'A' {
	           c +=26
	         }
	       }
	       fmt.Printf("%c", c)
	     }
	*/
}
