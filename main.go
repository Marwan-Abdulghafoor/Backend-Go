package main

import (
	"fmt"
	"net/http"

	"murwan.net/fiephrs-backend/dao"
	"murwan.net/fiephrs-backend/kafka"
)

func main() {

	client, _ := dao.ConnectDB()

	profileInfoChan := kafka.StartKafka()

	go func() {
		// Receive ProfileInfo values from the channel
		for info := range profileInfoChan {
			fmt.Println("Received info in main:", info)
			err := dao.InsertProfile(client, &info)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}()

	getPatientHandler := func(w http.ResponseWriter, r *http.Request) {
		// Call getPatientById with the captured client variable
		getPatientById(w, r, client)
	}

	getEditPatientHandler := func(w http.ResponseWriter, r *http.Request) {
		// Call getPatientById with the captured client variable
		println("Here1")
		editProfile(w, r, client)
	}

	// Register the handler function for the "/getProfileInfo" endpoint
	http.HandleFunc("/getProfileInfo", getPatientHandler)

	http.HandleFunc("/editProfile", getEditPatientHandler)

	if err := http.ListenAndServe(":9093", nil); err != nil {
		fmt.Println(err.Error())
	}
}
