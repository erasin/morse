// morse
//  // demo
//
// m := Morse{Dict: MOEN}
//
// en := "-- --- .-. ... .  -.-. --- -.. ."
// de := "morse code"
//
// fmt.Printf("Morse: \n%s\n", en)
// fmt.Println(m.Decode(en))
//
// fmt.Printf("word: \n%s\n", de)
// fmt.Println(m.Encode(de))
// ...
//
package morse

import (
	"bytes"
	"strings"
)

var MOEN = map[string]string{
	"A": ".-", "B": "-...", "C": "-.-.", "D": "-..", "E": ".", "F": "..-.",
	"G": "--.", "H": "....", "I": "..", "J": ".---", "K": "-.-", "L": ".-..",
	"M": "--", "N": "-.", "O": "---", "P": ".--.", "Q": "--.-", "R": ".-.",
	"S": "...", "T": "-", "U": "..-", "V": "...-", "W": ".--", "X": "-..-",
	"Y": "-.--", "Z": "--..", "0": "-----", "1": ".----", "2": "..---",
	"3": "...--", "4": "....-", "5": ".....", "6": "-....", "7": "--...",
	"8": "---..", "9": "----.",

	" ": " ",

	".": "._._._", ",": "__..__", "?": "..__..", "'": ".____.",
	"/": "_.._.", "(": "_.__.", ")": "_.__._", "&": "._...",
	":": "___...", ";": "_._._.", "=": "_..._", "+": "._._.",
	"-": "_...._", "_": "..__._", "\"": "._.._.", "$": "..._.._",
	"!": "_._.__", "@": ".__._.",
}

type Morse struct {
	Dict map[string]string
}

func NewMorse() *Morse {
	return &Morse{Dict: MOEN}
}

// 解密
func (self Morse) Decode(en string) string {

	dictRe := make(map[string]string)

	for k, v := range self.Dict {
		dictRe[v] = k
	}

	re := bytes.NewBufferString("")
	for _, c := range strings.Split(en, " ") {
		if mo, ok := dictRe[string(c)]; ok {
			re.WriteString(mo)
		} else {
			// 对编码机型解析

			re.WriteString(" ")
		}
	}

	return strings.Trim(re.String(), " ")
}

// 加密
func (self Morse) Encode(de string) string {

	de = strings.ToUpper(de)
	re := bytes.NewBufferString("")

	for _, c := range []rune(de) {
		char := string(c)
		if char == " " {
			re.WriteString(" ")
		} else {
			if mo, ok := self.Dict[char]; ok {
				re.WriteString(mo + " ")
			}
		}
	}

	return strings.Trim(re.String(), " ")
}
