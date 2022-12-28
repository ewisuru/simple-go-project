package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// findMax("10293846372920384757")
	// reverseString("10293846372920384757")
	// findUniqueCounts("1a02bcaaa93846372920384757")
	// structTest()
	// arraysandslices()

	// jsonreadwrite()
	// recursiveFibCall()
	// functionReference()
	// anonymousFuncCall()
	// interfaceDemo()
	// useLogs()
	// generics()
	for true {
		httpCalls()
		time.Sleep(1 * time.Second)
	}
}

func httpCalls() {
	fmt.Println("Starting...")
	resp, err := http.Get("http://salary-mgmt:8080/v1/users")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("No errors")
		fmt.Println(resp)
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		// if u want to read the body many time
		// u need to restore
		// reader := io.NopCloser(bytes.NewReader(bodyBytes))
		if err != nil {
			fmt.Println(err)
		}

		bodyString := string(bodyBytes)
		fmt.Println(resp.Status)
		fmt.Println(bodyString)
	}
}

/*
func generics() {
	fmt.Println(minOf(1, 3))
	fmt.Println(minOf(1.2, 0.3))
}

func minOf[T comparable](first T, second T) T {
	if first > second {
		return second
	}
	return first
}

type comparable interface {
	int | float64
}
*/

/*
func useLogs() {
	flags := log.LstdFlags | log.Lshortfile
	errorLogger := log.New(os.Stderr, "ERROR: ", flags)
	log.Println("Hello")
	errorLogger.Println("Some error")
}
*/

/*
func interfaceDemo() {
	// animals := []animal{{"dog"}, {"tiger"}}
	// birds := []bird{{"pigeon"}, {"eagle"}}
	sp := []speakable{animal{"dog"}, animal{"tiger"}, bird{"pigeon"}, bird{"eagle"}}
	for _, item := range sp {
		speak(item)
		v, ok := item.(bird)
		if ok {
			v.fly()
		}
	}
}

type bird struct {
	name string
}
type animal struct {
	name string
}

func (b bird) shout() {
	fmt.Println("Chirp... my name is " + b.name)
}
func (a animal) shout() {
	fmt.Println("Grr I am a " + a.name)
}
func (b bird) fly() {
	fmt.Println("Bye... I am flying")
}

type speakable interface {
	shout()
}

func speak(s speakable) {
	s.shout()
}
*/

/*
func anonymousFuncCall() {
	anonFunc := func(input int) {
		fmt.Println("hi", input)
	}
	fmt.Println("Start")
	anonFunc(10)
}
*/

/*
func functionReference() {
	result := funcCaller(10, operation)
	fmt.Println(result)
	fmt.Printf("%T\n", funcCaller)
}

func operation(i int) int {
	return i + i
}

func funcCaller(i int, f func(input int) int) int {
	return f(i)
}
*/

/*
func recursiveFibCall() {
	fib := recursivefunction(11)
	fmt.Println(fib)
}

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144
func recursivefunction(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return recursivefunction(n-1) + recursivefunction(n-2)
	}
}
*/

/*
func jsonreadwrite() {
	str := `{"name": "abc", "age": 22}`
	fmt.Println(str)
	var c1 Client
	err := json.Unmarshal([]byte(str), &c1)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println(c1)

	c2 := Client{Name: "abcd", Age: 21}
	val, er := json.Marshal(c2)
	if er != nil {
		fmt.Println("Error: " + er.Error())
	}
	fmt.Println(string(val))
}

type Client struct {
	Name string //`json:"full_name" `
	Age  int
}
*/

/*
func arraysandslices() {
	// slice
	var a1 []int = []int{1, 2, 3, 4}
	fmt.Println(a1)
	a1 = []int{1, 2, 3}
	fmt.Println(a1)
	a1 = append(a1, []int{9, 7, 6, 5}...)
	fmt.Println(a1)
	// fixed size array
	a2 := [...]int{5, 6, 7, 8, 9, 10}
	a3 := [...]int{1, 2, 3, 4, 5, 6}
	a2 = a3
	fmt.Println(a2)
	a2[3] = 999
	fmt.Println(a2)
	fmt.Println(a3)
}
*/

/*
func structTest() {
	// creating users setting values and printing them
	user1 := user{}
	user1.name = "amal"
	user1.age = 40
	user2 := user{"name", 12, "abc"}
	user3 := user{name: "name", age: 12, loginId: "abc"}
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)

	// struct using another struct. show pass by value behavior
	var app1 application = application{
		name: "app1",
		user: user1,
	}
	fmt.Println(app1)
	app1.user.age = 35
	fmt.Println(app1)
	fmt.Println(user1.age)

	// strut using struct pointer
	app2 := applicationwithpointer{
		name:    "app2",
		userptr: &user2,
	}
	fmt.Println(app2)
	app2.userptr.age = 100
	fmt.Println(app2)
	fmt.Println(user2)

}

type user struct {
	name    string
	age     int
	loginId string
}

func (u user) String() string {
	return "name: " + u.name + ", age: " + strconv.Itoa(u.age) + ", login id: " + u.loginId
}

type application struct {
	name string
	user user
}

func (app applicationwithpointer) String() string {
	return "name: " + app.name + ", {name: " + app.userptr.name + ", age: " +
		strconv.Itoa(app.userptr.age) + ", login id: " + app.userptr.loginId + "}"
}

type applicationwithpointer struct {
	name    string
	userptr *user
}
*/

/*
func findUniqueCounts(s string) {
	arr := []rune(s)
	m := map[string]int{}
	for i := 0; i < len(arr); i++ {
		cur := string(arr[i])
		val, ok := m[cur]
		if ok {
			m[cur] = val + 1
		} else {
			m[cur] = 1
		}
	}
	for key, value := range m {
		fmt.Println(key + " : " + strconv.Itoa(value))
	}
}
*/

/*
func reverseString(s string) {
	arr := []rune(s)
	sb := strings.Builder{}
	for i := len(arr) - 1; i >= 0; i-- {
		sb.WriteString(string(arr[i]))
	}
	fmt.Println(sb.String())
}
*/

/*
func findMax(input string) {

	arr := []rune(input)
	max := 0
	for i := 0; i < len(arr); i++ {
		val, _ := strconv.ParseInt(string(arr[i]), 0, 64)
		if max < int(val) {
			max = int(val)
		}
	}
	fmt.Print(max)
}
*/
