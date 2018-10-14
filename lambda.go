package lampper

import "log"

type Lambda struct {
	Name string
	Body string
}

func NewLambda(name, body string) *Lambda {
	lambda := &Lambda{Name: name, Body: body}
	lambda.Debug()
	return lambda
}

func (lambda *Lambda) Debug() {
	log.Println("[", lambda.Name, "]", lambda.Body)
}
