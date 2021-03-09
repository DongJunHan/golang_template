package main
//html/template과 text/template은 다르다. 문법에서 차이가 나는데
//html같은 경우는 특수문자가 인코딩되는거같음.
// text같은 경우는 문자그대로 나타내줌.
import(
	"os"
//	"text/template"
	"html/template"
	"fmt"
)

type User struct {
	Name string
	Email string
	Age int
}
func(u User) IsOld() bool{
	return u.Age > 30
}
func main(){
	user := User{Name : "dongjun", Email : "dongjun@naver.com", Age : 23}
	user2 := User{Name : "dong", Email : "don@naver.com", Age : 36}
	//ParseFiles함수를 사용하면 여러 탬플릿을 불러올 수 있다.	
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl","templates/tmpl2.tmpl")
	
	//Parse함수를 사용하려면 Parse함수안에 포맷을 넣어준다.
	//tmpl, err := template.New("Tmpl1").ParseFiles("Name:{{.Name}}\nEmail:{{.Email}}\nAge:{{.Age}}\n")
	if err != nil {
		panic(err)
	}
	//ParseFiles로 탬플릿을 이용하려면 ExecuteTemplate함수를 사용해야함.
	//Parse함수로 탬플릿을 이용하려면 Execute함수를 사용해야함.
	tmpl.ExecuteTemplate(os.Stdout,"tmpl2.tmpl",user)
	tmpl.ExecuteTemplate(os.Stdout,"tmpl2.tmpl",user2)
	fmt.Println("==end==")
}
