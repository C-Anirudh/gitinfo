![gitinfo-cover](logo.png)

[![GitHub Watches](https://img.shields.io/github/watchers/C-Anirudh/gitinfo.svg?style=social&label=Watch&maxAge=2592000)](https://GitHub.com/C-Anirudh/gitinfo/watchers)
[![GitHub Starts](https://img.shields.io/github/stars/C-Anirudh/gitinfo.svg?style=social&label=Star&maxAge=2592000)](https://GitHub.com/C-Anirudh/gitinfo/stargazers)
[![GitHub Forks](https://img.shields.io/github/forks/C-Anirudh/gitinfo.svg?style=social&label=Fork&maxAge=2592000)](https://GitHub.com/C-Anirudh/gitinfo/network)

CLI tool to search for GitHub users and projects.

## :minidisc: Installation
```
$ go get github.com/C-Anirudh/gitinfo
```

To be able to run the command from anywhere in the computer, run the following commands
```
$ cd $GOPATH/bin/
$ sudo cp gitinfo /usr/bin/gitinfo
```

## :rocket: Usage
```
$ gitinfo --help
```
OR
```
$ gitinfo
Usage: gitinfo [options]
Options:
         -u      The GitHub username of the person
```

```
$ gitinfo -u kenneth-reitz
Username                      :  kenneth-reitz
Name                          :  Kenneth Reitz
Email                         :  
Bio                           :  @storyscript.

I wrote @requests: HTTP for Humans. The only thing I really care about is user experience. 
Number of followers           :  26567
Count of people user follows  :  199
Number of public repositories :  37
Account created on            :  2009-08-26 21:17:47 +0000 UTC
```
