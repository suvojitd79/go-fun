package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func generateRandom(l int) string {

	data := strings.Split("儿,勒,屁,艾,西,艾,西,吉,艾,伊,娜,伊", ",")

	ans := make([]string, l)

	for i:=0;i<l;i+=1{
		ans[i] = data[rand.Intn(len(data))]
	}
	return strings.Join(ans, "-")
}

func sum(x ...int) (sum int) {
	for _,v := range x{
		sum += v
	}
	return
}

type Map struct {
	mu sync.Mutex
	data map[string]string
}

func (Mymap *Map) Add(x, y string) {
	Mymap.mu.Lock()
	Mymap.data[x] = y
	Mymap.mu.Unlock()
}

type Book struct{
	Name string `json:"name" xml:"name"`
	Price float64	`json:"price" xml:"price"`
	Quantity uint8	`json:"quantity" xml:"quantity"`
}


func main(){

	rand.Seed(time.Now().UTC().Unix())

	books := []Book{}

	file, _ := os.Open("google")

	fileReader := bufio.NewReader(file)

	for {

		data, e := fileReader.ReadString('\n')

		if e == io.EOF {
			Add(&books, data)
			break
		}
		Add(&books, data)
	}

	//j, _ := json.Marshal(books)


	//x := flag.String("name", "Sam", "User name")
	//y := flag.Int64("age", 22, "User Age")
	//flag.Parse()
	//
	//fmt.Println(*x, *y, flag.Args())

	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	panic(11)
}

func Add(books *[]Book, csvData string){
	csvDataSet := strings.Split(csvData, ";")
	price_val,_ := strconv.ParseFloat(csvDataSet[1], 64)
	quantity,_ := strconv.ParseInt(csvDataSet[2],10, 64)
	*books = append(*books, Book{csvDataSet[0], price_val,uint8(quantity)})
}