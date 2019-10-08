package main

import (
  "fmt"
  "os"
  "math"
)

func greetings() {
  fmt.Println("Welcome to the medicine catalog.")
  fmt.Println("")
  fmt.Println("What will you like to do today?.")
}

func farewell() {
  fmt.Println("Thank You for using MedCat, hope to see you soon!")
}

func menu() {

  fmt.Println("")
  fmt.Println("Menu:")
  fmt.Println("1. See all available medicines")
  fmt.Println("2. See current profits")
  fmt.Println("3. See upcoming profits")
  fmt.Println("4. See low-inventory")
  fmt.Println("5. Sell")
  fmt.Println("6. Restock")
  fmt.Println("7. New Medicine")
  fmt.Println("0. Exit program.")
  fmt.Println("")
  currentOption := getUserInputInt("Specify one of the option above to continue: ", 0, 7)
  callClear()
  menuSelector(currentOption)
}

func menuSelector(value int ) {
  switch value {
    case 1:
        printAllMedicine()
        break
    case 2:
        currentProfits()
        break
    case 3:
        futureProfits()
        break
    case 4:
        lowInventory()
        break
    case 5:
        sell()
        break
    case 6:
        restock()
        break
    case 7:
        newMedicine()
        break
    case 0:
        farewell()
        os.Exit(0)
        break
    default:

  }
}

func printAllMedicine() {
  printString("All medicines in our inventory:")
  for index, medicine := range catalog {
      printString( intToString(index+1) + ". " + medicine.getName())
  }
}

func currentProfits(){
  total := 0.0
  for _, medicine := range catalog {
      total += medicine.getCurrentProfit()
  }
  printString("Current total profits are $" + floatToString(total))
}

func futureProfits(){
  total := 0.0
  for _, medicine := range catalog {
      total += medicine.getFutureProfit()
  }
  printString("Future upcoming profits will be $" + floatToString(total))
}

func lowInventory(){
  printString("Medicines on low inventory:")
  currentIndex := 0
  for _, medicine := range catalog {
      if(medicine.getInventory() < 5){
        printString( intToString(currentIndex+1) + ". " + medicine.getName())
        currentIndex+=1
      }
  }
  if(currentIndex == 0){
    printString("None")
  }
}

func sell(){
  printAllMedicine()
  printString("0. Go Back")
  index := getUserInputInt("Which one are you selling: ", 0, getTotalMedicines())
  if(index != 0){
    medicine := getMedicine(index-1)
    printString("Amount of " + medicine.getName() + " in inventory is " + intToString(int(medicine.getInventory())) + ".")
    var amountSold int
    for{
      amountSold = getUserInputInt("How much are you selling? ", 1, int(medicine.getInventory()))
      if(medicine.getInventory() >= float64(amountSold)){
        break
      }
    }
    medicine.setInventory(medicine.getInventory() - float64(amountSold))
    medicine.setTotalSold(medicine.getTotalSold() + float64(amountSold))
    printString(intToString(amountSold) + " of " + medicine.getName() + " sold succesfully.")
  }
}

func restock(){
  printAllMedicine()
  printString("0. Go Back")
  index := getUserInputInt("Which one are you restocking: ", 0, getTotalMedicines())
  if(index != 0){
    medicine := getMedicine(index-1)
    printString("Amount of " + medicine.getName() + " in inventory is " + intToString(int(medicine.getInventory())) + ".")
    amountToAdd := getUserInputInt("How much are you adding? ", 1, int(math.MaxInt16))
    medicine.setInventory(medicine.getInventory() + float64(amountToAdd))
    printString(intToString(amountToAdd) + " of " + medicine.getName() + " added to inventory succesfully.")
  }
}

func newMedicine(){
  
  printString("To succesfully add a new medicine to the catalog please provide the following information: \n")

  name:= getUserInputString("Name: ")
  form:= getUserInputString("Form or structure: ")
  cat:= getUserInputString("Category: ")
  info:= getUserInputString("Information: ")
  quantity:= getUserInputInt("Quantity: ", 1, int(math.MaxInt16))
  effects:= getUserInputString("Side Effects: ")
  dir:= getUserInputString("Directions: ")
  printString("From what age to what age can this be used (1 - 99):")
  ageFrom:= getUserInputInt("From: ", 1, int(math.MaxInt16))
  ageTo:= getUserInputInt("To: ", int(ageFrom), 99)
  dosis:= getUserInputInt("Dosis: ", 1, int(math.MaxInt16))
  measure:= getUserInputString("Measurement: ")
  paid:= getUserInputFloat("Cost: ", 0, int(math.MaxInt16))
  price:= getUserInputFloat("Price: ", 0, int(math.MaxInt16))
  inventory:= getUserInputFloat("Inventory: ", 1, int(math.MaxInt16))

  newOne := Medicine{
    name,
    form,
    cat,
    info,
    quantity,
    effects,
    dir,
    ageFrom,
    ageTo,
    dosis,
    measure,
    paid,
    price,
    0,
    inventory,
  }
  addMedicine(newOne)
}
