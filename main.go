package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ops struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Zipcode     int    `json:"zipcode"`
	Phone       string `json:"phone"`
	DateUpdated string `json:"dateUpdated"`
}

type inf struct {
	Name string
	Ops  []ops `json:"data"`
}

var opsOut inf

func MakeRequest(index string) {
	urlRequest := fmt.Sprintf("http://basicdata.ru/api/json/zipcode/%s", index)
	resp, err := http.Get(urlRequest)
	if err != nil {
		fmt.Println("Сегодня что то не так с интернетами....")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Сегодня что то не так с интернетами....")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Видимо эту бумажку писал доктор.... она не читается.")
	}

	jsonErr := json.Unmarshal(body, &opsOut)
	if jsonErr != nil {
		fmt.Println("Что то пошло не так...")
		fmt.Println(jsonErr)
	}
	if len(opsOut.Ops) == 1 {
		fmt.Println("Индекс отделения:\t", opsOut.Ops[0].Zipcode)
		fmt.Println("Наименование:\t\t", opsOut.Ops[0].Name)
		fmt.Println("Расположено:\t\t", opsOut.Ops[0].Address)
		fmt.Println("Контактные телефоны:\t", opsOut.Ops[0].Phone)
	}
	if len(opsOut.Ops) == 0 {
		fmt.Println("Проверьте правильность почтового индекса")
	}
}

func main() {
	if len(os.Args) == 2 {
		MakeRequest(os.Args[1])
	} else {
		fmt.Println("Для корректной работы программы необходимо указать индекс ОПС в качестве аргумента")
		fmt.Println("Например: ops 353400")
	}

}
