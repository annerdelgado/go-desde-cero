// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	requests := 10
// 	ch := make(chan string, requests)
// 	for i := 1; i <= requests; i++ {
// 		go func(iter int) {
// 			fmt.Println(iter)
// 			resp, err := http.Get("http://localhost:8080/api/project")
// 			if err == nil {
// 				resp.Body.Close()
// 			}
// 			ch <- fmt.Sprint("%i", iter)
// 		}(i)
// 	}
// 	for i := 1; i <= requests; i++ {
// 		fmt.Println(<-ch)
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"
)

type LogMessage struct {
	Msg      string `json:"msg"`
	Severity string `json:"severity"`
	Time     string `json:"time"`
}

func main() {
	log := LogMessage{
		Msg:      "REQUEST CreateAccountComposite {false false [{POST /services/data/v52.0/sobjects/Account map[FirstName:0xc00029da50 LastName:0xc0004989e0 PersonBirthDate:0xc0001a41a8 PersonEmail:0xc0004989f0 PersonMobilePhone:0xc000498a00 RecordTypeID:0xc00029da40] newAccount} {GET /services/data/v52.0/query/?q=SELECT+id+,+firstName+,+lastName+,+PersonMobilePhone+,+PersonEmail+,+PersonContactId+FROM+Account+WHERE+id+=+'@{newAccount.id}' <nil> newQuery}]}",
		Severity: "info",
		Time:     "2024-04-23T19:59:27Z",
	}

	// Marshal the log message to JSON
	jsonData, err := json.Marshal(log.Msg)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Print the JSON as string
	fmt.Println(string(jsonData))
}
