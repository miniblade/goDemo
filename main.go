package main

import (
	"log"
	"net/http"
	//"reflect"
	//"github.com/miniblade/goDemo/controller"
)

func init(){

}

func main(){
	go httpServer()
	select{

	}
}

func httpServer(){
	http.HandleFunc("/gateway_admin", adminLogHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//1.按目录路由 2. panic 封装
//1.1.获取url后匹配函数
//1.2的实现方案 反射，注册全部函数
func RouterServer(f func()){

	defer func() {
		log.Println("done")  // 即使有恐慌Println也会正常执行
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
   f()
}
//func httpServer1(rw http.ResponseWriter,rPtr *http.Request){
//	var uri string=rPtr.RequestURI
//	var action =Action{}
//	reflect.TypeOf(Action{})
//}

func adminLogHandler(rw http.ResponseWriter,rPtr *http.Request){
	rw.Write([]byte("HelloWorld"))
}

func httpsServer(){

}
func grpcServer(){

}
func longLinkServer(){

}