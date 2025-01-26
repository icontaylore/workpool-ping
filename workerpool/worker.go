package workerpool

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Client struct {
	client http.Client
}

func CheckCode(createNewMem *Pool) {
	client := &Client{}
	wg := sync.WaitGroup{}

	wg.Add(len(createNewMem.urlStruct))
	for _, v := range createNewMem.urlStruct {
		go func() {
			defer wg.Done()

			resp, err := client.client.Get(v.Url)
			defer resp.Body.Close()

			if err != nil {
				v.Err = errors.New("Ошибка not a 200")
				log.Fatal(err)
			}
			v.Err = nil

			fmt.Printf("URL:%s   CODE:%d\n", v.Url, resp.StatusCode)
		}()
	}

	wg.Wait()
}
