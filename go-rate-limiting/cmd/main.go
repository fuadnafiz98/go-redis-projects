package main

import (
	"fmt"
	"net/http"
)

/*
func client() {
	ief := net.Interface{
		Name: "fuadnet",
	}
	vief := vnet.NewInterface(ief)
	vief.AddAddr(&net.TCPAddr{
		Port: 2020,
		IP:   net.IP(net.IPv4(2, 2, 2, 2)),
	})

	fmt.Println(vief.Addrs())
	fmt.Println(vief.InterfaceBase)

	addrs, _ := vief.Addrs()

	mdil := net.Dialer{LocalAddr: addrs[0]}

	mdil.Dial("tcp", "0.0.0.0:8888")
	c := http.Client{
		// Transport: ,
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Dialaing")
	resp, _ := c.Get("0.0.0.0:8888")
	fmt.Println(resp)
}
*/

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// inspect the request
		fmt.Println(r.URL, r.RemoteAddr)
		w.WriteHeader(200)
		w.Write([]byte("HELLO WORLD!"))
	})

	server := http.Server{
		Addr:    "0.0.0.0" + ":8888",
		Handler: handler,
	}

	fmt.Println("Server is running ...")
	// go func() {
	// 	client()
	// }()
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
