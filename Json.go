package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func Save() {
	saveDelivered()
	savePending()
	saveNotShipped()
}


func Load(){	
	loadDelivered()
	loadPending()
	loadNotShipped()
}




func saveToFile(jsonData []byte, filename string) {

	jsonFile, err := os.Create("./json/" + filename)

	if err != nil {
			panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}



func LoadFromFile(filename string) []byte {
	file, _ := os.Open("./json/" + filename)
	data, _ := ioutil.ReadAll(file)

	return data	
}



func saveDelivered(){
	jsonData, err := json.Marshal(mDelivered)

	if err != nil {
		panic(err)
	}

	// write to JSON file
	jsonFile, err := os.Create("./json/Delivered.json")

	if err != nil {
			panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}

func savePending(){
	jsonData, err := json.Marshal(mPending)

	if err != nil {
		panic(err)
	}

	// write to JSON file
	jsonFile, err := os.Create("./json/Pending.json")

	if err != nil {
			panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}

func saveNotShipped(){
	jsonData, err := json.Marshal(mNotShipped)

	if err != nil {
		panic(err)
	}

	// write to JSON file
	jsonFile, err := os.Create("./json/NotShipped.json")

	if err != nil {
			panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}


func loadNotShipped(){
	
	file, err := os.Open("./json/NotShipped.json")
	notShipped, _ := ioutil.ReadAll(file)

	json.Unmarshal(notShipped, &mNotShipped)

	if err != nil {
		 log.Fatal(err)
		 return
	}

	fmt.Println("Loaded "+ strconv.Itoa(len(mNotShipped)) +" items for map Not-Shipped Orders")
}

func loadPending(){
	
	file, err := os.Open("./json/Pending.json")
	pending, _ := ioutil.ReadAll(file)

	json.Unmarshal(pending, &mPending)

	if err != nil {
		 log.Fatal(err)
		 return
	}

	fmt.Println("Loaded "+ strconv.Itoa(len(mPending)) +" items for map Pending Shipments")
}

func loadDelivered(){
	
	file, err := os.Open("./json/Delivered.json")
	delivered, _ := ioutil.ReadAll(file)

	json.Unmarshal(delivered, &mDelivered)

	if err != nil {
		 log.Fatal(err)
		 return
	}

	fmt.Println("Loaded "+ strconv.Itoa(len(mDelivered)) +" items for map Delivered Shipments")
}