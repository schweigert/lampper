package lampper

type LambdaFunction struct {
	Name string
	F    func(*Lambda, *Peer)
}

func NewLambdaFunction(name string, f func(lambda *Lambda, peer *Peer)) *LambdaFunction {
	return &LambdaFunction{Name: name, F: f}
}
