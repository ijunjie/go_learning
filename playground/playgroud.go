package main

import "fmt"

type RegisterReq struct {
	Name  string
	Pass  string `p:"password1"`
	Pass2 string `p:"password2"`
}

type RegisterRes struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

type Foo struct {
	Age int
}

func a(x func(int) int) int {
	return x(5)
}

func main() {

	r := a(func(a int) int {
		return a * 2
	})
	fmt.Println(r)
	//var a *Foo
	//var b Foo
	//fmt.Println(&a)
	//fmt.Printf("%p\n", &b)
	//s := g.Server()
	//s.BindHandler("/register", func(r *ghttp.Request) {
	//	var req *RegisterReq
	//	var x **RegisterReq = &req
	//	if err := r.Parse(x); err != nil {
	//		r.Response.WriteJsonExit(RegisterRes{
	//			Code:  1,
	//			Error: err.Error(),
	//		})
	//	}
	//	// ...
	//	r.Response.WriteJsonExit(RegisterRes{
	//		Data: req,
	//	})
	//})
	//s.SetPort(8199)
	//s.Run()
}