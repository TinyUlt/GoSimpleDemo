package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {

	fname := "vcard.json"

	//构造结构
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	//编码到字符串
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s\n", js)

	//写入文件
	if err := ioutil.WriteFile(fname, js, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	//读取文件
	in, _ := ioutil.ReadFile(fname)
	fmt.Printf("JSON format: %s\n", in)

	//解码到结构
	var vcc VCard
	json.Unmarshal(in, &vcc)

	fmt.Printf("first name :%s\n", vcc.Remark)

	// using an encoder:
	/*
		{
			file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0)
			defer file.Close()
			enc := json.NewEncoder(file)
			err = enc.Encode(vc)
			if err != nil {
				log.Println("Error in encoding json")
			}
		}
		{
			var dvc *VCard = new(VCard)
			file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0)
			defer file.Close()
			enc := json.NewDecoder(file)
			err = enc.Decode(dvc)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("first name :%s\n", dvc.FirstName)
		}
	*/
}
