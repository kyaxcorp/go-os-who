package who

import (
	"context"
	"os/exec"
	"strings"
)

func linux() ([]User, error) {
	// use awk...

	var users []User

	// -u --users  list users logged in
	// -H --heading  print line of column headings
	//out, err := exec.CommandContext(context.Background(), "who", "-u", "-H").Output()

	var args = []string{
		"-u",
	}

	out, err := exec.CommandContext(context.Background(), "who", args...).Output()
	//b := bytes.NewBuffer(out)

	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	//log.Println("who is logged in")
	//log.Println(string(out))

	outStr := string(out)
	if outStr == "" {
		return users, nil
	}

	rows := strings.Split(outStr, "\n")
	if len(rows) == 0 {
		// no users have been found
		return users, nil
	}

	// TODO: on different os's time is formatted differently!  2022-07-24 18:14 or Jul 22 15:33
	for _, row := range rows {
		// split into columns...
		rowParts := strings.Split(row, " ")
		username := rowParts[0]
		if username == "" {
			continue
		}
		user := User{
			Name: username,
		}
		users = append(users, user)
	}

	return users, nil

	// TODO: try using awk?!
}
