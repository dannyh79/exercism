package airportrobot

type Italian struct {}
type Portuguese struct {}
type German struct {}

func SayHello(n string, p interface{}) string {
	var l string
	switch p.(type) {
	case Italian:
		l = "Italian: Ciao "+n+"!"
	case Portuguese:
		l = "Portuguese: Olá "+n+"!"
	case German:
		l = "German: Hallo "+n+"!"
	}
	return "I can speak "+l
}
