package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	user := flag.String("u", "", "The GitHub username of the person.")
	flag.Parse()

	if len(*user) == 0 {
		printUsage()
		os.Exit(1)
	}

	userInfo := getUserInfo(*user)
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
	fmt.Println("\t -u\t The GitHub username of the person")
}
