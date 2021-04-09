package chapter7

func Main72() {

}

//习惯使用单方法接口命名
type MyWriter interface {
	Write(p []byte) (n int, err error)
}

type MyReader interface {
	Read(p []byte) (n int, err error)
}
type MyCloser interface {
	Close() error
}

//接口内嵌
//组合接口
type MyReadWriter interface {
	MyReader
	MyWriter
}
type MyReadWriteCloser interface {
	MyReader
	MyWriter
	MyCloser
}
//混合组合
type MyReadWriter1 interface {
	Read(p []byte) (n int, err error)
	MyWriter
}
