package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	batchId      string   = "BCR39/40"
	timing       string   = "3.00 PM"
	meetingSlot  string   = "3.00 PM - 4.00 PM"
	trainer      string   = "Sivasakthi Sir"
	coordinators []string = []string{"Abdul Rahim O.M", "Afsal"}
	activity     string   = "Individual speaking"
	members      []string = []string{"Shruthi Kiron", "Abdul Rahim", "Afsal K T", "Amal", "Anjali", "Anusha", "Gadha", "Sreedevan",
		"Arjun", "Aswin", "Alan", "Mishab", "Ajay"}
)

func main() {
	clearScreen()
	createSessionReport()
}


func createSessionReport() {
	l1 := `*🔰 Session Report- ` + batchId + `*` + "\n"

	fmt.Println("Enter the date:    (Leave empty for today (", time.Now().Format("January 02, 2006"), ")):")
	date := getAlternative(time.Now().Format("January 02, 2006"))
	l2 := `🗓 Date : ` + date + "\n"

	fmt.Println("Enter the timing:    (Leave empty for ", timing, ")")
	timing = getAlternative(timing)

	l3 := `🕜 Timing : ` + timing + "\n"
	l4 := `👨🏽‍🏫 Trainer: ` + trainer + "\n"
	l5 := `🕵🏽‍♂️ Coordinator` + func() string {
		if len(coordinators) > 1 {
			return "s"
		}
		return ""
	}() + `: ` + fmt.Sprint(coordinators) + "\n"

	fmt.Println("Enter the activity:")
	reader := bufio.NewReader(os.Stdin)
	activity, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var l6 string
	if activity != "\n" {
		l6 = `⛳ Activity: ` + activity
	} else {
		fmt.Println("Activity cannot be empty")
	}
	l7 := "\n" + `*📃 Session Summary:*` + "\n"
	fmt.Println("Enter the session summary:\n📃 Session Summary:")
	sessionSummary, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	l8 := sessionSummary + "\n"

	l9 := `*Attendance:*` + "\n"
	for i, v := range members {
		fmt.Println(i+1, ". ", v)
	}
	fmt.Println("Enter the serial numbers of participants present separated by spaces, enter any character to stop")
	presentees, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading presentees:")
		panic(err)
	}
	presentees = strings.TrimSpace(presentees)
	presenteesArr := strings.Split(presentees, " ")
	presenteesMap := make(map[int]bool)

	if presenteesArr[0] != "" {
		for _, v := range presenteesArr {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			presenteesMap[num] = true
		}
	}

	var attendance string
	for i, v := range members {
		if presenteesMap[i+1] {
			attendance += `✅ ` + v + "\n"
		} else {
			attendance += `❌ ` + v + "\n"
		}
	}

	fmt.Println("Enter the tldv link:")
	var tldv, l10 string
	fmt.Scanf("%s", &tldv)
	if tldv != "" {
		l10 = `*📽️ TLDV link:* ` + "\n" + tldv + "\n"
	}

	fmt.Println("Enter the reporter's name:   (Enter nothing if by main coordinator)")
	reporter := getAlternative(coordinators[0])

	l11 := `✒️Report prepared by :` + "\n" + `    ` + reporter

	report := l1 + l2 + l3 +
		l4 + l5 + l6 +
		l7 + l8 + l9 +
		attendance + "\n" + l10 + l11

	fmt.Println("Session Report:")
	fmt.Println("===============================")
	fmt.Println(report)
	fmt.Println("===============================")

	copyToClipboard(&report)

}

func getAlternative(defaultString string) string {
	reader := bufio.NewReader(os.Stdin)
	new, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	if new == "\n" {
		return defaultString
	}
	return new

}

func copyToClipboard(report *string) {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(*report)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Report copied to clipboard")
}

func clearScreen() {
	cmd := exec.Command("clear")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
