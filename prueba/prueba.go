package prueba

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func AnExporterFuncImJustADummyLib() {
	fmt.Println("Greetings Emoji!")
	emoji.Println(":speech_balloon: May I take your order?")

	burgerMessage := emoji.Sprint("I'd like a :hamburger: and :fries:!!")
	fmt.Println(burgerMessage)
}
