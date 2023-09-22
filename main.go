package main

import (
	"fmt"

	"github.com/5paisa/go5paisa"
	"github.com/spf13/viper"
)

// func main() {
// 	url := "https://Openapi.5paisa.com/VendorsAPI/Service1.svc/V3/Holding"

// 	payload := []byte(`{
// 		"head": {
// 			"key": "LMDJW8BWpl6NR80hmcF5WudUIcvZbvHL"
// 		},
// 		"body": {
// 			"ClientCode": "{Your Client Code}"
// 		}
// 	}`)

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}

// 	req.Header.Set("Authorization", "bearer {Your Access Token}")
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Cookie", "NSC_JOh0em50e1pajl5b5jvyafempnkehc3=ffffffffaf103e0f45525d5f4f58455e445a4a423660")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error making request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("Response Status:", resp.Status)
// 	// Read and print the response body if needed
// 	// responseBody, _ := ioutil.ReadAll(resp.Body)
// 	// fmt.Println("Response Body:", string(responseBody))
// }

func main() {
	conf := &go5paisa.Config{
		AppName:       viper.GetString("APP_NAME"),
		AppSource:     viper.GetString("APP_SOURCE"),
		UserID:        viper.GetString("USER_ID"),
		Password:      viper.GetString("PASSWORD"),
		UserKey:       viper.GetString("USER_KEY"),
		EncryptionKey: viper.GetString("ENCRYPTION_KEY"),
	}
	order := go5paisa.Order{
		ClientCode:      "51611672",
		Exchange:        go5paisa.BSE,
		ScripCode:       11111,
		ExchangeSegment: go5paisa.CASH,
		Qty:             1,
		OrderType:       go5paisa.BUY,
	}

	appConfig := go5paisa.Init(conf)
	client, err := go5paisa.Login(appConfig, "kavishshah30@gmail.com", "iQ4a1yPZzl2", "20020930")

	if err != nil {
		panic(err)
	}
	res, err := client.PlaceOrder(order)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)

}
