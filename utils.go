package main

import (
  "fmt"
  "strconv"
  "os"
  "os/exec"
  "runtime"
  "bufio"
)

var clear map[string]func()

func init() {
    clear = make(map[string]func())
    clear["linux"] = func() {
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["darwin"] = func() {
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func callClear() {
    value, ok := clear[runtime.GOOS]
    if ok {
        value()
    } else {
        panic(runtime.GOOS + " currently is unsupported!")
    }
}

func floatToString(value float64 ) string {
  return fmt.Sprintf("%.2f", value)
}

func intToString(value int ) string {
  return strconv.Itoa(value)
}

func printString(text string) {
  fmt.Println(text)
}

func getUserInputInt(text string, min int, max int) int {

  stdin := bufio.NewReader(os.Stdin)
  var input int
  for{
    fmt.Print(text)

    _, err := fmt.Fscan(stdin, &input)
    if err != nil || input < min || input > max {
         fmt.Println("Invalid input, try again.")
    } else {
      break
    }
    stdin.ReadString('\n')
    fmt.Println("")
  }
  return input
}

func getUserInputFloat(text string, min int, max int) float64 {

  stdin := bufio.NewReader(os.Stdin)
  var input float64
  for{
    fmt.Print(text)

    _, err := fmt.Fscan(stdin, &input)
    if err != nil || input < float64(min) || input > float64(max) {
         fmt.Println("Invalid input, try again.")
    } else {
      break
    }
    stdin.ReadString('\n')
    fmt.Println("")
  }
  return input
}

func getUserInputString(text string) string {
  fmt.Print(text)

  stdin := bufio.NewReader(os.Stdin)
  var userString string
  _, err := fmt.Fscan(stdin, &userString)
  if err != nil {
       fmt.Println("Invalid input, try again")
  }
  stdin.ReadString('\n')
  return userString;
}
