package main

func main() {
	//创建一个channel, 默认双向的
	ch := make(chan int)

	//双向channel能隐式转换为单向的channel，单向无法转换成双向
	var writeCh chan<- int = ch //只能写， 不能读
	var readCh chan<- int = ch  //只能读， 不能写

	writeCh <- 666 //写
	//<-writeCh //err

	<-readCh //读
	//readCh <- 666//写， error
}
