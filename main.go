package main

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var wg sync.WaitGroup
var choices = []string{"Disburse", "Account Inquery", "Cashflow Inquery", "Lainnya"}
var choicesmf = []string{"MUF01", "MTF01"}
var choices_service = []string{"Ya", "Tidak"}
var selectedValues = "Pilihan anda adalah : "

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))

	focusedStylekey        = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	blurredStylekey        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStylekey         = focusedStyle.Copy()
	noStylekey             = lipgloss.NewStyle()
	helpStylekey           = blurredStyle.Copy()
	cursorModeHelpStylekey = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButtonkey = focusedStyle.Copy().Render("[ Submit ]")
	blurredButtonkey = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

func initialModel() model {

	outfile, _ := os.Create("./logapp.log") // update path for your needs
	infoLog := log.New(outfile, "INFO \t", log.Ldate|log.Ltime)
	//WarrLog := log.New(os.Stderr, "Warning\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(outfile, "ERROR \t", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Println("+++ Running  ETL performance API Canalis +++")

	var allkey = make(map[string]allkeykey)

	usernamemuf := "API_MUF"
	passowrdmuf := "P@ssw0rdMuf1"
	secretkeymuf := "99c051b4bc7dd4c69c551e4531419281"
	mackeymuf := "d0eba9667f8c73de78f771195d832e86"
	urlmuf := ""

	usernamemtf := "API_MTF"
	passowrdmtf := "P@ssw0rdMtf1"
	secretkeymtf := "cb4c2059df266cc722d86aebac86cceb"
	mackeymtf := "8465637562a2ef44ced654dfa274a6f2"
	urlmtf := "xxxx"

	mtfkey := allkeykey{
		username:  &usernamemtf,
		passowrd:  &passowrdmtf,
		secretkey: &secretkeymtf,
		mackey:    &mackeymtf,
		url:       &urlmtf,
	}

	mufkey := allkeykey{
		username:  &usernamemuf,
		passowrd:  &passowrdmuf,
		secretkey: &secretkeymuf,
		mackey:    &mackeymuf,
		url:       &urlmuf,
	}

	allkey["mtf"] = mtfkey
	allkey["muf"] = mufkey

	s := spinner.NewModel()
	s.Spinner = spinner.Dot

	m := model{
		inputs:        make([]textinput.Model, 4),
		inputslainnya: make([]textinput.Model, 4),
		inputsquery:   make([]textinput.Model, 3),
		selectedbool:  true,
		inputskey:     make([]textinput.Model, 5),
		spinner:       s,

		keymf: allkey,

		infoLog:  infoLog,
		errorLog: errorLog,
	}

	// all key
	var t textinput.Model
	for i := range m.inputskey {
		t = textinput.New()
		t.CursorStyle = cursorStylekey
		t.CharLimit = 300

		switch i {
		case 0:
			//
			t.Prompt = "Username  > "
			t.Placeholder = "Username"
			t.PromptStyle = focusedStylekey
			t.TextStyle = focusedStylekey
			//t.SetValue(*m.mfusername)
			t.Focus()
		case 1:
			t.Prompt = "Password  > "
			t.Placeholder = "Password"
			t.PromptStyle = focusedStylekey
			t.TextStyle = focusedStylekey
			//t.SetValue(*m.mfpassowrd)
		case 2:
			t.Prompt = "Secretkey > "
			t.Placeholder = "Secretkey"
			t.PromptStyle = focusedStylekey
			t.TextStyle = focusedStylekey
			//t.SetValue(*m.mfsecretkey)

		case 3:
			t.Prompt = "Mackey    > "
			t.Placeholder = "Mackey"
			t.PromptStyle = focusedStylekey
			t.TextStyle = focusedStylekey
			//t.SetValue(*m.mfmackey)

		case 4:
			t.Prompt = "Url       > "
			t.Placeholder = "Url"
			t.PromptStyle = focusedStylekey
			t.TextStyle = focusedStylekey
			//t.SetValue(*m.mfurl)
		}

		m.inputskey[i] = t
	}

	//var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			//
			t.Placeholder = "Loan File Name"
			t.Prompt = "File Customer > "
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.SetValue("01 Customer.txt")
			//t.Focus()

		case 1:
			t.Prompt = "File Loan     > "
			t.Placeholder = "Loan File Name"
			t.Blur()
			t.PromptStyle = noStyle
			t.TextStyle = noStyle
			t.SetValue("01 Loan.txt")
		case 2:
			t.Prompt = "File Asset    > "
			t.Placeholder = "Asset File Name"
			t.Blur()
			t.PromptStyle = noStyle
			t.TextStyle = noStyle
			t.SetValue("01 Aset.txt")

		case 3:
			t.Prompt = "BatchId    > "
			t.Placeholder = "BatchId"
			t.Blur()
			t.PromptStyle = noStyle
			t.TextStyle = noStyle
			t.SetValue("")
		}

		m.inputs[i] = t
	}

	for i := range m.inputslainnya {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			//
			//t.Placeholder = "File Name"
			t.Prompt = "File Name      > "
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.Focus()

		case 1:
			t.Prompt = "Batch ID       > "
			//t.Placeholder = "Batch ID "
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			//t.SetValue("01 Loan.txt")
		case 2:
			t.Prompt = "Transaksi Code > "
			//t.Placeholder = "Transaksi Code"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			//t.SetValue("01 Aset.txt")

		case 3:
			t.Prompt = "File Code      > "
			//t.Placeholder = "Transaksi Code"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			//t.SetValue("01 Aset.txt")

		}

		m.inputslainnya[i] = t
	}

	for i := range m.inputsquery {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			//
			//t.Placeholder = "File Name"
			t.Prompt = "File Name      > "
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.Focus()
		case 1:
			t.Prompt = "Transaksi Code > "
			//t.Placeholder = "Transaksi Code"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			//t.SetValue("01 Aset.txt")

		case 2:
			t.Prompt = "File Code      > "
			//t.Placeholder = "Transaksi Code"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			//t.SetValue("01 Aset.txt")

		}

		m.inputsquery[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	//return m.spinner.Tick
	return textinput.Blink

}

type DataReady struct {
	Err  error
	Data Databody
}

type DataReadyApi struct {
	Err  error
	Data Databody
}

func (m model) HitServices() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(1 * time.Second)
		apiresponse := m.RunServices()

		return FinalDataResponseAPI{Apiresponse: apiresponse}
	}
}
func (m model) Prepare() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(1 * time.Second)
		var finaldata FinalData
		if m.selected == "Disburse" {
			finaldata = m.CreateFileDisburse()
		} else if m.selected == "Lainnya" {
			finaldata = m.CreateFileLainnya()
		} else {
			finaldata = m.CreateFileInquery()
		}

		return FinalDataResponse{FinalData: finaldata}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "h":

			if m.showmsgbool && msg.String() == "h" {
				m.showmsgbool = false
				m.showmsg = ""
				return m, nil

			} else if m.typingkey && msg.String() == "h" {

				m.typingkey = false
				m.selectedboolmf = true
				return m, nil

			}

		case "ctrl+c", "q", "esc":

			return m, tea.Quit
		case "tab", "shift+tab", "enter", "up", "down":
			// pilih transaction
			if m.selectedbool {
				if msg.String() == "enter" {

					m.choice = choices[m.cursor]
					m.selected = m.choice

					// if (choices[m.cursor] != "Disburse") && (choices[m.cursor] != "Lainnya") {
					// 	m.showmsg = "Fungsi " + m.selected + " belum di support, Please contact Ciso"
					// 	m.showmsgbool = true
					// 	return m, nil
					// }

					if choices[m.cursor] == "Disburse" {

						m.selectedbool = false
						m.typing = true
						return m, nil
					} else if choices[m.cursor] == "Lainnya" {

						m.selectedbool = false
						m.typinglainya = true
						return m, nil
					} else { //"Account Inquery", "Cashflow Inquery"
						// m.showmsg = "Fungsi " + m.selected + " belum di support, Please contact Ciso"
						// m.showmsgbool = true
						// return m, nil
						m.selectedbool = false
						m.typinginquery = true

						return m, nil

					}

				} else if msg.String() == "up" {
					m.cursor--
					if m.cursor < 0 {
						m.cursor = len(choices) - 1
					}

				} else if msg.String() == "down" {
					m.cursor++
					if m.cursor >= len(choices) {
						m.cursor = 0
					}

				}
			}
			//hanya untuk file disburse
			if m.typing {
				s := msg.String()

				// Did the user press enter while the submit button was focused?
				// If so, exit.
				if s == "enter" && m.focusIndex == len(m.inputs) {

					for _, data := range m.inputs {

						if strings.TrimSpace(data.Value()) == "" {

							m.showmsg = "Lengkapi file terlebih dahulu"
							m.showmsgbool = true
							return m, nil

						}

					}

					m.typing = false
					m.selectedboolmf = true

					return m, nil
					//return m, tea.Quit
				}

				// Cycle indexes
				if s == "up" || s == "shift+tab" {
					m.focusIndex--
				} else {
					m.focusIndex++
				}

				if m.focusIndex > len(m.inputs) {
					m.focusIndex = 0
				} else if m.focusIndex < 0 {
					m.focusIndex = len(m.inputs)
				}

				cmds := make([]tea.Cmd, len(m.inputs))
				for i := 0; i <= len(m.inputs)-1; i++ {
					if i == m.focusIndex {
						// Set focused state
						cmds[i] = m.inputs[i].Focus()
						m.inputs[i].PromptStyle = focusedStyle
						m.inputs[i].TextStyle = focusedStyle
						continue
					}
					// Remove focused state
					m.inputs[i].Blur()
					m.inputs[i].PromptStyle = noStyle
					m.inputs[i].TextStyle = noStyle
				}

				return m, tea.Batch(cmds...)

			}

			if m.typinglainya {
				s := msg.String()

				// Did the user press enter while the submit button was focused?
				// If so, exit.
				if s == "enter" && m.focusIndex == len(m.inputslainnya) {

					for _, data := range m.inputslainnya {

						if strings.TrimSpace(data.Value()) == "" {

							m.showmsg = "Lengkapi file terlebih dahulu"
							m.showmsgbool = true
							return m, nil

						}

					}

					m.typinglainya = false
					m.selectedboolmf = true

					return m, nil
					//return m, tea.Quit
				}

				// Cycle indexes
				if s == "up" || s == "shift+tab" {
					m.focusIndex--
				} else {
					m.focusIndex++
				}

				if m.focusIndex > len(m.inputslainnya) {
					m.focusIndex = 0
				} else if m.focusIndex < 0 {
					m.focusIndex = len(m.inputslainnya)
				}

				cmds := make([]tea.Cmd, len(m.inputslainnya))
				for i := 0; i <= len(m.inputslainnya)-1; i++ {
					if i == m.focusIndex {
						// Set focused state
						cmds[i] = m.inputslainnya[i].Focus()
						m.inputslainnya[i].PromptStyle = focusedStyle
						m.inputslainnya[i].TextStyle = focusedStyle
						continue
					}
					// Remove focused state
					m.inputslainnya[i].Blur()
					m.inputslainnya[i].PromptStyle = noStyle
					m.inputslainnya[i].TextStyle = noStyle
				}

				return m, tea.Batch(cmds...)

			}

			if m.typinginquery {

				s := msg.String()

				// Did the user press enter while the submit button was focused?
				// If so, exit.
				if s == "enter" && m.focusIndex == len(m.inputsquery) {

					for _, data := range m.inputsquery {

						if strings.TrimSpace(data.Value()) == "" {

							m.showmsg = "Lengkapi file terlebih dahulu"
							m.showmsgbool = true
							return m, nil

						}

					}

					m.typinginquery = false
					m.selectedboolmf = true

					return m, nil
					//return m, tea.Quit
				}

				// Cycle indexes
				if s == "up" || s == "shift+tab" {
					m.focusIndex--
				} else {
					m.focusIndex++
				}

				if m.focusIndex > len(m.inputsquery) {
					m.focusIndex = 0
				} else if m.focusIndex < 0 {
					m.focusIndex = len(m.inputsquery)
				}

				cmds := make([]tea.Cmd, len(m.inputsquery))
				for i := 0; i <= len(m.inputsquery)-1; i++ {
					if i == m.focusIndex {
						// Set focused state
						cmds[i] = m.inputsquery[i].Focus()
						m.inputsquery[i].PromptStyle = focusedStyle
						m.inputsquery[i].TextStyle = focusedStyle
						continue
					}
					// Remove focused state
					m.inputsquery[i].Blur()
					m.inputsquery[i].PromptStyle = noStyle
					m.inputsquery[i].TextStyle = noStyle
				}

				return m, tea.Batch(cmds...)

			}
			//pilih mf
			if m.selectedboolmf {
				if msg.String() == "enter" {

					m.choicemf = choicesmf[m.cursormf]
					m.selectedmf = m.choicemf

					//m.checkdata = true

					if m.selectedmf == "MUF01" {

						m.inputskey[0].SetValue(*m.keymf["muf"].username)
						m.inputskey[1].SetValue(*m.keymf["muf"].passowrd)
						m.inputskey[2].SetValue(*m.keymf["muf"].secretkey)
						m.inputskey[3].SetValue(*m.keymf["muf"].mackey)
						m.inputskey[4].SetValue(*m.keymf["muf"].url)
					} else if m.selectedmf == "MTF01" {
						m.inputskey[0].SetValue(*m.keymf["mtf"].username)
						m.inputskey[1].SetValue(*m.keymf["mtf"].passowrd)
						m.inputskey[2].SetValue(*m.keymf["mtf"].secretkey)
						m.inputskey[3].SetValue(*m.keymf["mtf"].mackey)
						m.inputskey[4].SetValue(*m.keymf["mtf"].url)
					}

					m.typing = false
					m.selectedboolmf = false
					m.selectedbool = false
					m.typingkey = true

					return m, nil

				} else if msg.String() == "up" {
					m.cursormf--
					if m.cursormf < 0 {
						m.cursormf = len(choicesmf) - 1
					}

				} else if msg.String() == "down" {
					m.cursormf++
					if m.cursormf >= len(choicesmf) {
						m.cursormf = 0
					}

				}
			}

			//all key
			if m.typingkey {
				s := msg.String()

				// Did the user press enter while the submit button was focused?
				// If so, exit.
				if s == "enter" && m.focusIndexkey == len(m.inputskey) {

					// m.showmsg = "Fungsi " + m.selected + " belum di support, Please contact Ciso"
					// m.showmsgbool = true
					// return m, nil

					m.typing = false
					m.selectedboolmf = false
					m.typingkey = false
					m.loading1 = true

					//m.checkdata = true

					//return m, cmd
					return m, tea.Batch(
						spinner.Tick, //memangil komponent lain.
						m.Prepare(),
					)
				}

				// Cycle indexes
				if s == "up" || s == "shift+tab" {
					m.focusIndexkey--
				} else {
					m.focusIndexkey++
				}

				if m.focusIndexkey > len(m.inputskey) {
					m.focusIndexkey = 0
				} else if m.focusIndexkey < 0 {
					m.focusIndexkey = len(m.inputskey)
				}

				cmds := make([]tea.Cmd, len(m.inputskey))
				for i := 0; i <= len(m.inputskey)-1; i++ {
					if i == m.focusIndexkey {
						// Set focused state
						cmds[i] = m.inputskey[i].Focus()
						m.inputskey[i].PromptStyle = focusedStylekey
						m.inputskey[i].TextStyle = focusedStylekey
						continue
					}
					// Remove focused state
					m.inputskey[i].Blur()
					m.inputskey[i].PromptStyle = noStylekey
					m.inputskey[i].TextStyle = noStylekey
				}

				return m, tea.Batch(cmds...)

			}

			//hit service
			if m.hitservis {

				if msg.String() == "enter" {

					m.choice_service = choices_service[m.cursor_service]
					m.selected_service = m.choice_service

					if m.selected_service == "Ya" {
						m.hitservis = false
						m.loading2 = true

						return m, tea.Batch(
							spinner.Tick,
							m.HitServices(), //hit services
						)

					} else {
						return m, tea.Quit
					}

				} else if msg.String() == "up" {
					m.cursor_service--
					if m.cursor_service < 0 {
						m.cursor_service = len(choices_service) - 1
					}

				} else if msg.String() == "down" {
					m.cursor_service++
					if m.cursor_service >= len(choices_service) {
						m.cursor_service = 0
					}

				}
			}
		}

	case FinalDataResponse:
		m.FinalData = msg.FinalData
		m.loading1 = false
		m.hitservis = true

		return m, nil

	case FinalDataResponseAPI:

		m.Apiresponse = msg.Apiresponse
		m.loading2 = false
		m.hitservis_result = true

		return m, nil

	}

	if m.loading1 {
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if m.loading2 {
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	var xcmd tea.Cmd
	if m.typing {
		xcmd = m.updateInputs(msg)

	}

	if m.typingkey {
		xcmd = m.updateInputskey(msg)
	}

	if m.typinglainya {
		xcmd = m.updateInputslainnya(msg)

	}

	if m.typinginquery {
		xcmd = m.updateInputsinquery(msg)
	}

	//return m, m.updateInputskey(msg)

	//return m, m.updateInputs(msg)

	return m, xcmd

}

func (m *model) updateInputskey(msg tea.Msg) tea.Cmd {
	var cmdskey = make([]tea.Cmd, len(m.inputskey))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputskey {
		m.inputskey[i], cmdskey[i] = m.inputskey[i].Update(msg)
	}

	return tea.Batch(cmdskey...)
}

func (m *model) updateInputsinquery(msg tea.Msg) tea.Cmd {
	var cmdskey = make([]tea.Cmd, len(m.inputsquery))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputsquery {
		m.inputsquery[i], cmdskey[i] = m.inputsquery[i].Update(msg)
	}

	return tea.Batch(cmdskey...)
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *model) updateInputslainnya(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputslainnya))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputslainnya {
		m.inputslainnya[i], cmds[i] = m.inputslainnya[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
func (m model) View() string {

	if m.showmsgbool {

		s := strings.Builder{}

		s.WriteString("!!!!++++  Msg = " + m.showmsg + "  +++++!!!!\n\n")

		s.WriteString("\n(press h to home menu)\n")

		return s.String()
	}

	if m.hitservis_result {

		s := strings.Builder{}
		s.WriteString("Respons Canalis : " + *m.Apiresponse.Response + "\n")

		m.hitservis_result = false

		return s.String()

	}

	//show data
	if m.checkdata {

		s := strings.Builder{}
		s.WriteString("Proses : " + m.selected + "\n")
		s.WriteString("Name MF :" + m.selectedmf + "\n")

		for i, data := range m.inputs {
			s.WriteString("index ke : " + strconv.Itoa(i) + " Nama File : " + data.Value() + "\n")

		}

		for i, datakey := range m.inputskey {
			s.WriteString("index ke : " + strconv.Itoa(i) + " iskey : " + datakey.Value() + "\n")

		}
		return s.String()

	}

	//tranaksai
	if m.selectedbool {
		s := strings.Builder{}
		s.WriteString("Pilih Transaction yang akan di jalankan?\n\n")

		for i := 0; i < len(choices); i++ {

			if m.cursor == i {
				s.WriteString("[x] ")

				//s.WriteString("(•) ")
			} else {
				//s.WriteString("( ) ")
				s.WriteString("[ ]  ")
			}
			s.WriteString(choices[i])
			s.WriteString("\n")
		}

		s.WriteString("\n(press q to quit)\n")

		return s.String()

	}

	//file disbuser
	if m.typing {
		var b strings.Builder

		//b.WriteString("Proses " + m.selected + "\n\n")

		for i := range m.inputs {
			b.WriteString(m.inputs[i].View())
			if i < len(m.inputs)-1 {
				b.WriteRune('\n')
			}
		}

		button := &blurredButton
		if m.focusIndex == len(m.inputs) {
			button = &focusedButton
		}
		fmt.Fprintf(&b, "\n\n%s\n\n", *button)

		// b.WriteString(helpStyle.Render("cursor mode is "))
		// b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
		// b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

		return b.String()

	}

	//file disbuser
	if m.typinglainya {
		var b strings.Builder

		//b.WriteString("Proses " + m.selected + "\n\n")

		for i := range m.inputslainnya {
			b.WriteString(m.inputslainnya[i].View())
			if i < len(m.inputslainnya)-1 {
				b.WriteRune('\n')
			}
		}

		button := &blurredButton
		if m.focusIndex == len(m.inputslainnya) {
			button = &focusedButton
		}
		fmt.Fprintf(&b, "\n\n%s\n\n", *button)

		// b.WriteString(helpStyle.Render("cursor mode is "))
		// b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
		// b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

		return b.String()

	}

	//file inquery
	if m.typinginquery {
		var b strings.Builder

		//b.WriteString("Proses " + m.selected + "\n\n")

		for i := range m.inputsquery {
			b.WriteString(m.inputsquery[i].View())
			if i < len(m.inputsquery)-1 {
				b.WriteRune('\n')
			}
		}

		button := &blurredButton
		if m.focusIndex == len(m.inputsquery) {
			button = &focusedButton
		}
		fmt.Fprintf(&b, "\n\n%s\n\n", *button)

		// b.WriteString(helpStyle.Render("cursor mode is "))
		// b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
		// b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

		return b.String()

	}
	//pilih mf
	if m.selectedboolmf {

		s := strings.Builder{}
		s.WriteString("Proses " + m.selected + "\n")
		s.WriteString("Pilih MF yang akan di jalankan?\n\n")

		for i := 0; i < len(choicesmf); i++ {

			if m.cursormf == i {
				s.WriteString("[x] ")

				//s.WriteString("(•) ")
			} else {
				//s.WriteString("( ) ")
				s.WriteString("[ ]  ")
			}
			s.WriteString(choicesmf[i])
			s.WriteString("\n")
		}

		return s.String()

	}

	//all key
	if m.typingkey {

		//d := strings.Builder{}
		var d strings.Builder

		d.WriteString("Proses " + m.selected + " " + m.selectedmf + "\n\n")

		for i := range m.inputskey {
			d.WriteString(m.inputskey[i].View())
			if i < len(m.inputskey)-1 {
				d.WriteRune('\n')
			}
		}

		buttonkey := &blurredButtonkey
		if m.focusIndexkey == len(m.inputskey) {
			buttonkey = &focusedButtonkey
		}

		fmt.Fprintf(&d, "\n\n%s\n\n", *buttonkey)

		// b.WriteString(helpStyle.Render("cursor mode is "))
		// b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
		// b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

		d.WriteString("\n(press h to back menu)\n")
		return d.String()

	}

	if m.loading1 {
		return fmt.Sprintf("%s Prepre all data .... please wait.", m.spinner.View())
	}

	if m.loading2 {
		return fmt.Sprintf("%s Running hit Services Canalis .... please wait.", m.spinner.View())
	}

	if m.hitservis {

		//return fmt.Sprintf("Hasil retrive %v : ", m.hasil)

		s := strings.Builder{}
		s.WriteString("Proses Prepare " + m.selected + " " + m.selectedmf + " selesai, Apakah lanjut hit service ?\n\n")

		for i := 0; i < len(choices_service); i++ {

			if m.cursor_service == i {
				s.WriteString("[x] ")

				//s.WriteString("(•) ")
			} else {
				//s.WriteString("( ) ")
				s.WriteString("[ ]  ")
			}
			s.WriteString(choices_service[i])
			s.WriteString("\n")
		}

		return s.String()

	}

	return ""

}

func main() {

	// StartReturningModel returns the model as a tea.Model.
	//_, err := p.StartReturningModel()

	// err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Start()
	// if err != nil {
	// 	fmt.Println("Oh no:", err)
	// 	os.Exit(1)
	// }

	err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Start()

	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	// if m, ok := m.(model); ok && m.choice != "" {
	// 	fmt.Printf("\n---\nYou chose %s!\n", m.choice)
	// }

}
