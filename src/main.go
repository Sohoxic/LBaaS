package main

import(
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface{
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *httpRequest)
}

type simpleServer struct{
	addr string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer{
	serverUrl, err := url.Parse(addr)
	handleErr(err)
	return &simpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port 			string
	roundRobinCount	int
	servers			[]Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer{
	return &LoadBalancer{
		port: 			 port,
		roundRobinCount: 0,
		servers: 		 servers,

	}
}

func handleErr(err error){
	if err!= nil{
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func (lb *LoadBalancer) getNextAvailableServer() Server{}

func (lb *LoadBalancer) serverProxy(rw http.ResponseWriter, r *httpRequest){} 

func main(){
	servers := []Server{
		newSimpleServer
	}
}