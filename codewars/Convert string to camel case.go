//https://www.codewars.com/kata/517abf86da9663f1d2000003/train/go
// Complete the method/function so that it converts dash/underscore delimited words into camel casing. The first word within the output should be capitalized only if the original word was capitalized (known as Upper Camel Case, also often referred to as Pascal case). The next words should be always capitalized.
// Examples

// "the-stealth-warrior" gets converted to "theStealthWarrior"

// "The_Stealth_Warrior" gets converted to "TheStealthWarrior"

// "The_Stealth-Warrior" gets converted to "TheStealthWarrior"

// Вот перевод условия задачи:

// **Задание:**
// Допишите функцию так, чтобы она преобразовывала слова, разделенные тире или нижним подчеркиванием, в верблюжий регистр (**camelCase**).

// **Правила:**

// * Первое слово в результате должно начинаться с заглавной буквы только в том случае, если оно было заглавным в оригинале (такой вариант называют **Upper Camel Case** или **Pascal case**).
// * Все последующие слова должны всегда начинаться с заглавной буквы.

// **Примеры:**

// * `"the-stealth-warrior"` преобразуется в `"theStealthWarrior"`
// * `"The_Stealth_Warrior"` преобразуется в `"TheStealthWarrior"`
// * `"The_Stealth-Warrior"` преобразуется в `"TheStealthWarrior"`

package kata

import (
	"strings"
	"unicode"
)

func ToCamelCase(s string) string {
	s = strings.ReplaceAll(s, "_", "-") //заменяет все на что-то
	worldsSlice := strings.Split(s, "-")
	var builder strings.Builder
	builder.WriteString(worldsSlice[0])
	for i := 1; i < len(worldsSlice); i++ {
		oneWorld := worldsSlice[i]
		runes := []rune(oneWorld)
		runes[0] = unicode.ToUpper(runes[0])
		builder.WriteString(string(runes))
	}

	return builder.String()
}
