package pack2

var a = 1

// Что-то на русском языке. // want "comment is not written in acceptable language"
func someFunc() int {
	return a
}

// single English comment