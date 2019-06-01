package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	user := flag.String("u", "", "The GitHub username of the person.   (required input)")
	fileName := flag.String("f", "", "The file name with path and .json extension to which information has to be copied. Path should be relative to home directory")
	flag.Parse()

	if len(*user) == 0 {
		printUsage()
		os.Exit(1)
	}

	out, _ := exec.Command("sh", "-c", "eval echo ~$USER").Output()
	out = out[:len(out)-1]

	userInfo := getUserInfo(*user)

	if len(*fileName) != 0 {
		f, err := os.Create(string(out) + *fileName)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		data, err := json.Marshal(userInfo)
		fmt.Fprintf(f, "%s\n", data)
	}
	printUserInfo(userInfo)
}

func printUserInfo(result User) {
	fmt.Println("Username                      : ", result.Login)
	fmt.Println("Name                          : ", result.Name)
	fmt.Println("Email                         : ", result.Email)
	fmt.Println("Bio                           : ", result.Bio)
	fmt.Println("Number of followers           : ", result.Followers)
	fmt.Println("Count of people user follows  : ", result.Following)
	fmt.Println("Number of public repositories : ", result.PublicRepos)
	fmt.Println("Account created on            : ", result.CreatedAt)
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	fmt.Println("\t -u\t The GitHub username of the person\t(required input)")
	fmt.Println("\t -f\t The file name with path and .json extension to which information has to be copied")
	fmt.Println("\t   \t Path should be relative to home directory")
}
