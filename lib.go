package lib

type MyInterface interface {
    World(hello string) string 
}

func MyFun(i MyInterface) string {
	return i.World("Hello")
}
