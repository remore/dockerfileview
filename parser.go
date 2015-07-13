package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

func isKeyword(word string) bool {
	s := []string{"ENV", "LABEL", "MAINTAINER", "ADD", "COPY", "FROM", "ONBUILD", "WORKDIR", "RUN", "CMD", "ENTRYPOINT", "EXPOSE", "VOLUME", "USER"}
	for _, a := range s {
		if a == word {
			return true
		}
	}
	return false
}

func parseDockerfile(filename string, fileurl string, body io.Reader, colorFlag bool) {
	util := &displayUtil{colorFlag}
	scanner := bufio.NewScanner(body)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		cmdline := regexp.MustCompile(`[\t\v\f\r ]+`).Split(strings.TrimSpace(scanner.Text()), 2)
		if cmdline[0] == "FROM" {
			resp, url, err := getDockerfile(cmdline[1])
			if err == nil {
				parseDockerfile(cmdline[1], url, resp, colorFlag)
			} else {
				if cmdline[1] != "scratch" {
					fmt.Println(util.headerString(cmdline[1]))
					fmt.Println(err)
				}
			}
			if fileurl != "" {
				filename += "\n# " + fileurl
			}
			fmt.Println(util.headerString(filename))
		}
		if len(cmdline) == 2 && isKeyword(cmdline[0]) {
			fmt.Println(util.coloredString(cmdline[0], "36"), cmdline[1])
		} else if len(cmdline) == 2 && cmdline[0] == "#" {
			fmt.Println(util.coloredString(scanner.Text(), "33"))
		} else {
			fmt.Println(scanner.Text())
		}
	}
}
