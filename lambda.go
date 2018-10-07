package lampper

type Lambda struct {
	Name string
	Body string
}

func NewLambda(name, body string) *Lambda {
	return &Lambda{Name: name, Body: body}
}
