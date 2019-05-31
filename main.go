package main

import (
	"flag"
	"fmt"
)

func main() {
	user := flag.String("user", "C-Anirudh", "The GitHub username of the person.")
	flag.Parse()
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
