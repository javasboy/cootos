package main

import (
  "fmt"
	"reflect"
)

//User结构
type User struct {
	//Id,Name,Age三个字段
	Id   int
	Name string
	Age  int
}

type cis_c struct {
	Id   int
	Name string
	Age  int
}

type cis_d struct {
	cis_c
	title string
}

type cis_e struct {
	Id   int
	Name string
	Age  int
}

//Hello方法
func (u User) Hello() {
	fmt.Println("Hello world.")
}

func main() {
	m := cis_d{cis_c: cis_c{1, "OK", 12}, title: "123"}
	//取出cis_d的类型
	t := reflect.TypeOf(m)

	//取出cis_d的第一个字段
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.Field(1))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

	x := 123
	v := reflect.ValueOf(&x)
	//elem取得v,x是int型，所以我们用SetInt
	v.Elem().SetInt(999)
	fmt.Println(x)

	//声明一个类型
	u := User{1, "OK", 12}
	//调用Info方法把u传进去
	Info(u)

	Info(&u)

	z := User{1, "OK", 12}
	Set(&z)
	fmt.Println(z)

	n := cis_e{1, "OK", 12}
	n.Hello("joe")

	h := reflect.ValueOf(n)
	hn := h.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("joe")}
	hn.Call(args)

}
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}

	/*if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}*/
}

//再使用一个方法，传入一个空接口，然后输出我们传入的结构的输入信息
func Info(o interface{}) {
	//用TypeOf获得接口的类型
	t := reflect.TypeOf(o)
	//打印名称,调用t.Name方法
	fmt.Println("Type:", t.Name())
	//取出这个类型，判断它是否等于reflect.Struct
	if k := t.Kind(); k != reflect.Struct {
		//如果不等于，我们就是要输入一个错误提示
		fmt.Println("xxx")
		//return返回
		return
	}

	//用ValueOf把字段打印出来
	v := reflect.ValueOf(o)

	fmt.Println("Fields:")
	//NumField循环到它所拥有的字段数量,因为它是通过索引来获得的
	for i := 0; i < t.NumField(); i++ {
		//这个i就是索引从0开始
		f := t.Field(i)
		//取出这个字段对应的值
		val := v.Field(i).Interface()
		//打印出来，名称，类型，值
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)

	}
	//用NumMethos方法取得方法拥有的数量
	for i := 0; i < t.NumMethod(); i++ {
		//通过索引来获得它的值
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func (e cis_e) Hello(name string) {
	fmt.Println("Hello", name, ",my name is", e.Name)
}
