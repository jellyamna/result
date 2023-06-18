package main

import (
	"encoding/json"
	"log"
	"os"
)

func (m model) CreateFileDisburse() FinalData {
	//endcode file disburse.
	var finaldata FinalData
	var files []Files
	var databody Databody
	var batchid string

	for i, fdisburse := range m.inputs {

		s := Files{}

		switch i {
		case 0:
			s.FileCode = "CUSTOMER"
			s.FileName = fdisburse.Value()
			s.File = *EncondeFile(fdisburse.Value(), m.errorLog)
			files = append(files, s)
		case 1:
			s.FileCode = "LOAN"
			s.FileName = fdisburse.Value()
			s.File = *EncondeFile(fdisburse.Value(), m.errorLog)
			files = append(files, s)
		case 2:
			s.FileCode = "MGMTASSET"
			s.FileName = fdisburse.Value()
			s.File = *EncondeFile(fdisburse.Value(), m.errorLog)
			files = append(files, s)
		case 3:

			batchid = fdisburse.Value()

		}

	}

	databody.TransactionCode = "DIS01"
	//databody.BatchId = "AK001DIS012019010227011"
	databody.BatchId = batchid
	databody.Files = files

	bytes, err := json.Marshal(databody)

	if err != nil {
		m.errorLog.Fatal("Can not convert body origin to json format", err.Error())
	}
	// fmt.Println(string(bytes))

	var username string
	var passowrd string
	var secretkey string
	var mackey string
	var url string

	for i, datakey := range m.inputskey {

		switch i {
		case 0:
			username = datakey.Value()
		case 1:
			passowrd = datakey.Value()
		case 2:
			secretkey = datakey.Value()
		case 3:
			mackey = datakey.Value()
		case 4:
			url = datakey.Value()

		}

	}

	//urlvariable := url
	valueEncode := username + ":" + passowrd

	authorizationVariable := "Basic " + EncodeString(valueEncode)
	//fmt.Println("Authorization =Basic " + EncodeString(valueEncode))

	authentication := *Getauthentication(&secretkey, &valueEncode)

	authenticationVariable := "Basic " + EncodeString(authentication)
	//fmt.Println("Authentication =Basic " + EncodeString(authentication))

	x := string(bytes)

	mac := GetHashMac(&mackey, &x)

	macVariable := *mac
	//fmt.Println("mac =" + *mac)

	bodyrest := *Getauthentication(&secretkey, &x)

	bodyrestVariable := EncodeString(bodyrest)
	//fmt.Println("rest body=" + EncodeString(bodyrest))

	f, err := os.Create("DisburseOri.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	finaldata.Authorization = &authorizationVariable
	finaldata.Authentication = &authenticationVariable
	finaldata.Mac = &macVariable
	finaldata.Restbody = &bodyrestVariable
	finaldata.Urlrest = &url

	f.WriteString("body original =" + x + "\n")
	f.WriteString("authorizationVariable =" + authorizationVariable + "\n")
	f.WriteString("authenticationVariable =" + authenticationVariable + "\n")
	f.WriteString("mac =" + macVariable + "\n")
	f.WriteString("rest body =" + bodyrestVariable + "\n")
	f.WriteString("base url =" + url + "\n")

	return finaldata

	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

}

func (m model) CreateFileLainnya() FinalData {
	//endcode file disburse.
	var finaldata FinalData
	var files []Files
	var databody Databody

	var filename string
	var batchid string
	var transactioncode string
	var filecode string
	var fileencode string

	for i, data := range m.inputslainnya {

		//s := Files{}

		switch i {
		case 0:
			filename = data.Value()
		case 1:
			batchid = data.Value()
		case 2:
			transactioncode = data.Value()
		case 3:
			filecode = data.Value()
		}

	}

	s := Files{}
	fileencode = *EncondeFile(filename, m.errorLog)

	s.FileName = filename
	s.FileCode = filecode
	s.File = fileencode

	files = append(files, s)

	databody.TransactionCode = transactioncode
	databody.BatchId = batchid
	databody.Files = files

	bytes, err := json.Marshal(databody)

	if err != nil {
		m.errorLog.Fatal("Can not convert body origin to json format", err.Error())
	}
	// fmt.Println(string(bytes))

	var username string
	var passowrd string
	var secretkey string
	var mackey string
	var url string

	for i, datakey := range m.inputskey {

		switch i {
		case 0:
			username = datakey.Value()
		case 1:
			passowrd = datakey.Value()
		case 2:
			secretkey = datakey.Value()
		case 3:
			mackey = datakey.Value()
		case 4:
			url = datakey.Value()

		}

	}

	//urlvariable := url
	valueEncode := username + ":" + passowrd

	authorizationVariable := "Basic " + EncodeString(valueEncode)
	//fmt.Println("Authorization =Basic " + EncodeString(valueEncode))

	authentication := *Getauthentication(&secretkey, &valueEncode)

	authenticationVariable := "Basic " + EncodeString(authentication)
	//fmt.Println("Authentication =Basic " + EncodeString(authentication))

	x := string(bytes)

	mac := GetHashMac(&mackey, &x)

	macVariable := *mac
	//fmt.Println("mac =" + *mac)

	bodyrest := *Getauthentication(&secretkey, &x)

	bodyrestVariable := EncodeString(bodyrest)
	//fmt.Println("rest body=" + EncodeString(bodyrest))

	f, err := os.Create(filecode + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	finaldata.Authorization = &authorizationVariable
	finaldata.Authentication = &authenticationVariable
	finaldata.Mac = &macVariable
	finaldata.Restbody = &bodyrestVariable
	finaldata.Urlrest = &url

	f.WriteString("body original =" + x + "\n")
	f.WriteString("authorizationVariable =" + authorizationVariable + "\n")
	f.WriteString("authenticationVariable =" + authenticationVariable + "\n")
	f.WriteString("mac =" + macVariable + "\n")
	f.WriteString("rest body =" + bodyrestVariable + "\n")
	f.WriteString("base url =" + url + "\n")

	return finaldata

	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

}

func (m model) CreateFileInquery() FinalData {
	//endcode file disburse.
	var finaldata FinalData
	var files []Files
	var databody Databodyinquery

	var filename string
	var transactioncode string
	var filecode string
	var fileencode string

	for i, data := range m.inputsquery {

		//s := Files{}

		switch i {
		case 0:
			filename = data.Value()
		case 1:
			transactioncode = data.Value()
		case 2:
			filecode = data.Value()

		}

	}

	s := Files{}
	fileencode = *EncondeFile(filename, m.errorLog)

	s.FileName = filename
	s.FileCode = transactioncode
	s.File = fileencode

	files = append(files, s)

	databody.TransactionCode = transactioncode
	databody.Files = files

	bytes, err := json.Marshal(databody)

	if err != nil {
		m.errorLog.Fatal("Can not convert body origin to json format", err.Error())
	}
	// fmt.Println(string(bytes))

	var username string
	var passowrd string
	var secretkey string
	var mackey string
	var url string

	for i, datakey := range m.inputskey {

		switch i {
		case 0:
			username = datakey.Value()
		case 1:
			passowrd = datakey.Value()
		case 2:
			secretkey = datakey.Value()
		case 3:
			mackey = datakey.Value()
		case 4:
			url = datakey.Value()

		}

	}

	//urlvariable := url
	valueEncode := username + ":" + passowrd

	authorizationVariable := "Basic " + EncodeString(valueEncode)
	//fmt.Println("Authorization =Basic " + EncodeString(valueEncode))

	authentication := *Getauthentication(&secretkey, &valueEncode)

	authenticationVariable := "Basic " + EncodeString(authentication)
	//fmt.Println("Authentication =Basic " + EncodeString(authentication))

	x := string(bytes)

	mac := GetHashMac(&mackey, &x)

	macVariable := *mac
	//fmt.Println("mac =" + *mac)

	bodyrest := *Getauthentication(&secretkey, &x)

	bodyrestVariable := EncodeString(bodyrest)
	//fmt.Println("rest body=" + EncodeString(bodyrest))

	f, err := os.Create(filecode + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	finaldata.Authorization = &authorizationVariable
	finaldata.Authentication = &authenticationVariable
	finaldata.Mac = &macVariable
	finaldata.Restbody = &bodyrestVariable
	finaldata.Urlrest = &url

	f.WriteString("body original =" + x + "\n")
	f.WriteString("authorizationVariable =" + authorizationVariable + "\n")
	f.WriteString("authenticationVariable =" + authenticationVariable + "\n")
	f.WriteString("mac =" + macVariable + "\n")
	f.WriteString("rest body =" + bodyrestVariable + "\n")
	f.WriteString("base url =" + url + "\n")

	return finaldata

	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

}
