package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/abiosoft/ishell"
)

func main() {
	cliarg := os.Args
	rand.Seed(time.Now().UnixNano())
	if len(cliarg) > 1 {
		dice := cliarg[1]
		out, err := RollThem(dice)
		if err == nil {
			fmt.Printf(out)
		} else {
			print(err)
		}
		return
	} else {
		shell := ishell.New()
		shell.Println("Welcome to DiceRoll-Go.\nYou can roll dice by typing 'roll #d#'.")

		shell.AddCmd(&ishell.Cmd{
			Name: "roll",
			Help: "roll #d# dice",
			Func: func(c *ishell.Context) {
				cliargs := c.Args
				dice := cliargs[0]
				out, err := RollThem(dice)
				if err == nil {
					c.Printf(out)
				} else {
					c.Println(err)
				}
			},
		})
		shell.Run()
	}
}

// IsValidDie confirms the number of sides on the die
func IsValidDie(dice int) bool {
	switch dice {
	case
		2, 4, 6, 8, 10, 12, 20, 100:
		return true
	}
	return false
}

// TooManyDie limits dice rolled to 50
func TooManyDie(dice int) bool {
	if dice < 50 {
		return false
	}
	return true
}

// RollThem parses string and rolls dice
func RollThem(dice string) (out string, err error) {
	ps := strings.Split(dice, "d")
	dNum, errn := strconv.Atoi(ps[0])
	if errn != nil {
		return "", errors.New("Error: Please format your die roll as [number]d[type], e.g. 2d6, 5d10, etc")
	}
	dType, errt := strconv.Atoi(ps[1])
	if errt != nil {
		return "", errors.New("Error: Please format your die roll as [number]d[type], e.g. 2d6, 5d10, etc")
	}
	if TooManyDie(dNum) {
		return "", errors.New("In what world do you need so many dice rolled?")
	}
	if IsValidDie(dType) {
		total := 0
		current := 0
		out := ""
		for i := 0; i < dNum; i++ {
			current = rand.Intn(dType) + 1
			currents := strconv.Itoa(current)
			if i == dNum-1 {
				out += currents + "\n"
			} else {
				out += currents + ", "
			}
			total += current
		}
		totals := strconv.Itoa(total)
		out += "Total: " + totals + "\n"
		return out, nil
	}
	return "", errors.New("Not a valid type of die")
}
