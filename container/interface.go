package container

type Queue interface{
	Put(interface{})
	Get() interface{}
	TryGet() (interface{},bool)
}

type Stack interface{
	Push(interface{})
	Pop() interface{}
	TryPop() (interface{},bool)
}