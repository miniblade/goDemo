package controller

import "net/http"

type Action struct{

}
func (a Action) HelloWorldAction(rw http.ResponseWriter,rPtr *http.Request){
	rw.Write([]byte("HelloWorld"))
}
