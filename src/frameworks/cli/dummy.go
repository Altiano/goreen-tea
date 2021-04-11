package cli

import (
	"bufio"
	"context"
	"fmt"
	"os"

	application "gitlab.com/altiano/golang-boilerplate/src/app"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks"
)

type dummyCLI struct {
	app application.App
}

func NewDummyCLI(app application.App) frameworks.Server {
	return dummyCLI{
		app: app,
	}
}

func (g dummyCLI) Run() {
	fmt.Println("Dummy CLI interface is running and receiving input from stdin")
	showAvailableCommands()

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		input := sc.Text()

		switch input {
		case "1", "oo", "onsite", "onsite order":
			callOnsite(g)

			// <-time.After(2000 * time.Millisecond)
			// fmt.Println()
			// showAvailableCommands()
		default:
			showAvailableCommands()
		}

	}
}

func callOnsite(g dummyCLI) {
	res, err := g.app.OnsiteOrder(context.Background(), application.OnsiteOrderReq{
		Email:  "altianogerung@gmail.com",
		Orders: []string{"apel", "pisang", "pepaya", "nanas"},
		// ReferralCode:            1,
	})

	if err != nil {
		fmt.Println("ERROR", err)
	} else {
		fmt.Println("SUCCESS", res)
	}
}

func showAvailableCommands() {
	fmt.Println("Available commands:")
	fmt.Println("-", "1", "oo", "onsite", "onsite order")
}
