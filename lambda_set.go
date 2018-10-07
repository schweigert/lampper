package lampper

type LambdaSet struct {
	Set map[string]*LambdaFunction
}

func NewLambdaSet() *LambdaSet {
	return &LambdaSet{Set: make(map[string]*LambdaFunction)}
}

func (lambdaSet *LambdaSet) Add(name string, f func(lambda *Lambda, peer *Peer)) *LambdaSet {
	lambdaSet.Set[name] = NewLambdaFunction(name, f)
	return lambdaSet
}

func (lambdaSet *LambdaSet) Get(name string) *LambdaFunction {
	return lambdaSet.Set[name]
}
