package main

import (
	"log"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
)

type Databody struct {
	TransactionCode string  `json:"transactionCode"`
	BatchId         string  `json:"batchId"`
	Files           []Files `json:"files"`
}

type Databodyinquery struct {
	TransactionCode string  `json:"transactionCode"`
	Files           []Files `json:"files"`
}

type Files struct {
	FileCode string `json:"fileCode"`
	FileName string `json:"fileName"`
	File     string `json:"file"`
}

type model struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	err      error
	//all key and url
	typingkey     bool //flag for allkey
	focusIndexkey int
	inputskey     []textinput.Model //untuk allkey
	cursorModekey textinput.CursorMode
	//check
	checkdata bool
	//prosess
	cursor   int
	choice   string
	selected string
	//flag for select type prosess
	selectedbool bool
	//mf
	cursormf       int
	choicemf       string
	selectedmf     string
	selectedboolmf bool //flag for select type prosess
	//file disburse only
	typing bool
	//file lainnnya
	typinglainya  bool
	typinginquery bool
	focusIndex    int
	//file disburse
	inputs []textinput.Model //untuk file disburse
	//untuk file lainnya
	inputslainnya []textinput.Model
	//untuk inquery
	inputsquery []textinput.Model
	cursorMode  textinput.CursorMode
	//spiner
	spinner spinner.Model
	//loaing generate file
	loading1 bool
	//loading result rest api
	loading2 bool
	//ada data
	hasil string

	//hit services.
	hitservis        bool //hit ke service.
	hitservis_result bool //result hit ke service.
	cursor_service   int
	choice_service   string
	selected_service string
	//mesage
	showmsg     string
	showmsgbool bool

	//allkeydata mtf_muf
	keymf map[string]allkeykey

	mfusername  *string
	mfpassowrd  *string
	mfsecretkey *string
	mfmackey    *string
	mfurl       *string

	// //disburse
	// Authorization  *string
	// Authentication *string
	// Mac            *string
	// Restbody       *string
	// Urlrest        string

	//respoonserest
	// response_service *string

	FinalData   FinalData
	Apiresponse Apiresponse
}

type allkeykey struct {
	username  *string
	passowrd  *string
	secretkey *string
	mackey    *string
	url       *string
}

type FinalData struct {
	//disburse
	Authorization  *string
	Authentication *string
	Mac            *string
	Restbody       *string
	Urlrest        *string
}

type FinalDataResponse struct {
	FinalData FinalData
}

type FinalDataResponseAPI struct {
	Apiresponse Apiresponse
}

type Apiresponse struct {
	Response *string
}
