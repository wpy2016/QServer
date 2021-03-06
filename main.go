package main

import (
	"fmt"
	"socket_component/server/connection"
	"socket_component/util"
	"socket_component/server"
)

type TopProcesser struct{

}

func (this *TopProcesser)Processe(token connection.TokenHandler,length int,bytes []byte){
	stream := util.NewStreamBuffer()
	stream.Append(bytes)
	i := stream.ReadLine()
	fmt.Println(i)
	stream.Renew()
	stream.WriteLine("LO")
	stream.InsertLen()
	token.SendAsync(stream.Bytes(), func(handler connection.TokenHandler, bytes []byte, i int, e error) {
		fmt.Println("Send Successful")
	})
}



func main() {

	s := server.NewQServer(":8888")
	tp := TopProcesser{}
	s.SetProcesser(&tp)
	s.SyncListen()
	fmt.Println("test sync")

}