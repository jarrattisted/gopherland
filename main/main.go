package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/jarrattisted/gopherworld/data"
	"github.com/jarrattisted/gopherworld/models"
)

// Main
func main() {

	var globpher *models.Gopher

	// Should only init seed once
	rand.Seed(time.Now().Unix())

	// Create a new CLI shell
	shell := ishell.New()

	// Display a welcome
	shell.Println(`Welcome to Gopherworld! 👋`)

	// Register a function for "gopher" command to generate a new Gopher
	shell.AddCmd(&ishell.Cmd{
		Name: "gopher",
		Help: "Creates a new Gopher and displays initial values",
		Func: func(c *ishell.Context) {
			// disable the '>>>' for cleaner same line input.
			c.ShowPrompt(false)
			defer c.ShowPrompt(true) // yes, revert after login.
			c.Print("What would you like to call your Gopher?\n(Leave blank if want a random name generated) ")
			name := c.ReadLine()
			if name == "" || name == " " {
				name = models.GenerateName()
			}
			c.Println("Generating Gopher...")
			g := models.GenerateGopher(name)
			globpher = g
			c.Println("⚡️ Woah! A bolt hit Gopherland and a new Gopher called", name, "was created.")
			c.Println(g.String())
		},
	})

	// Register a function for "wabbit" command to generate a new Wabbit
	shell.AddCmd(&ishell.Cmd{
		Name: "wabbit",
		Help: "Creates a new Wabbit and displays initial values",
		Func: func(c *ishell.Context) {
			// disable the '>>>' for cleaner same line input.
			c.ShowPrompt(false)
			defer c.ShowPrompt(true) // yes, revert after login.
			c.Println("Generating Wabbit...")
			w := models.GenerateWabbit()
			c.Println("🐰 Your Gopher stumbled upon a Wabbit.")
			c.Println(w.String())
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "info",
		Help: "Get information about your current Gopher",
		Func: func(c *ishell.Context) {
			if globpher == nil {
				c.Println("You need to create a Gopher first!")
				return
			}
			choice := c.MultiChoice([]string{
				"🍃 Attributes",
				"🎒 Backpack inventory",
			}, "What information would you like to view?")
			if choice == 0 {
				c.Println(globpher.AttributesString())
			} else {
				c.Println(globpher.BackpackContentsString())
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "explore",
		Help: "Get your Gopher to explore the Gopherworld!",
		Func: func(c *ishell.Context) {
			if globpher == nil {
				c.Println("You need to create a Gopher first!")
				return
			}
			choice := c.MultiChoice([]string{
				"🍇 Huckleberry Hill",
				"🌁 Peak Point",
				"⛴ Lengthy Lagoon",
				"🏜 Chalk Caves",
				"🗿 Sentry Statue",
			}, "Where would you like to explore?")
			switch choice {
			case 0:
				result := globpher.ExploreLandmark(data.HuckleberryHill)
				fmt.Println(result)
			case 1:
				result := globpher.ExploreLandmark(data.PeakPoint)
				fmt.Println(result)
			case 2:
				result := globpher.ExploreLandmark(data.LengthyLagoon)
				fmt.Println(result)
			case 3:
				result := globpher.ExploreLandmark(data.HuckleberryHill)
				fmt.Println(result)
			case 4:
				result := globpher.ExploreLandmark(data.HuckleberryHill)
				fmt.Println(result)
			}
		},
	})

	// run shell
	shell.Run()

	// landmarks := []string{"Huckleberry Hill", "Peak Point", "Lengthy Lagoon",
	//                       "Chalk Caves", "Sentry Statue"}

}
