package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		cmd = "xdg-open"
		args = []string{url}
	}
	return exec.Command(cmd, args...).Start()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("VIP Video Parser")
	fmt.Println("-----------------")
	fmt.Println("Disclaimer:")
	fmt.Println("This software is an open-source project, prohibited for any commercial use!!!")
	fmt.Println("This tool is for learning, exchange, and technical testing purposes only, aiming to explore web video parsing technology, and does not provide any content storage or dissemination services. Please do not use it for any illegal purposes. Users should bear all legal responsibilities arising from the abuse of this tool. Please respect film and television work copyrights and support genuine content. If your rights are infringed, please contact us promptly for removal.")
	fmt.Println("The developer does not bear any direct or indirect loss responsibility.")
	fmt.Println("-----------------")

	for {
		fmt.Println("\nPlease select an option:")
		fmt.Println("1. Enter Video Link and Play VIP Video")
		fmt.Println("2. Open iQIYI Official Website")
		fmt.Println("3. Open Tencent Video Official Website")
		fmt.Println("4. Open Youku Video Official Website")
		fmt.Println("5. Open Hitmux Official Website")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		switch choiceStr {
		case "1":
			fmt.Print("Please paste the video link here: ")
			videoLink, _ := reader.ReadString('\n')
			videoLink = strings.TrimSpace(videoLink)

			if videoLink == "" {
				fmt.Println("Invalid video link. Please try again.")
				continue
			}

			parseURL := "https://jx.xmflv.cc/?url=" + videoLink
			fmt.Printf("Attempting to open: %s\n", parseURL)
			err := openBrowser(parseURL)
			if err != nil {
				fmt.Printf("Error opening browser: %v. Please copy and paste the URL manually.\n", err)
			}
			fmt.Println("Operation complete. Returning to main menu.")

		case "2":
			fmt.Println("Opening iQIYI Official Website...")
			err := openBrowser("https://www.iqiyi.com")
			if err != nil {
				fmt.Printf("Error opening browser: %v. Please copy and paste https://www.iqiyi.com manually.\n", err)
			}
			fmt.Println("Operation complete. Returning to main menu.")

		case "3":
			fmt.Println("Opening Tencent Video Official Website...")
			err := openBrowser("https://v.qq.com")
			if err != nil {
				fmt.Printf("Error opening browser: %v. Please copy and paste https://v.qq.com manually.\n", err)
			}
			fmt.Println("Operation complete. Returning to main menu.")

		case "4":
			fmt.Println("Opening Youku Video Official Website...")
			err := openBrowser("https://www.youku.com/")
			if err != nil {
				fmt.Printf("Error opening browser: %v. Please copy and paste https://www.youku.com/ manually.\n", err)
			}
			fmt.Println("Operation complete. Returning to main menu.")

		case "5":
			fmt.Println("Opening Hitmux Official Website...")
			err := openBrowser("https://hitmux.top")
			if err != nil {
				fmt.Printf("Error opening browser: %v. Please copy and paste https://hitmux.top manually.\n", err)
			}
			fmt.Println("Operation complete. Returning to main menu.")

		case "6":
			fmt.Println("Exiting. Goodbye!")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 6.")
		}
	}
}
