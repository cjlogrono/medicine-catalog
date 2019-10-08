package main

type Medicine struct {

  //about
  name string
  form string
  cat string
  info string
  quantity int
  sideEffects string
  direction string
  agesFrom int
  agesTo int
  dosis int
  measure string
  //buisness
  paid float64
  price float64
  totalSold float64
  inventory float64
}

var catalog []Medicine

func init() {

  catalog = []Medicine{
    Medicine {
      name: "Advil",
      form: "Capsul",
      cat: "general",
      info: "temporarily relieves minor aches & pains due to: headache, muscular aches, menstrual cramps, the common cold, backache, toothache, minor pain or arthritis.",
      quantity: 40,
      sideEffects: "Symptoms may include: hives, facial swelling, asthma, shock, skin reddening, rash, blisters.",
      direction: "1 capsule every 4 to 6 hours while symptoms persist. If pain or fever does not responsd to 1 capsule, 2 capsules may be used. Do not exceed 6 capsules in 24 hours, unless directed by a doctor.",
      agesFrom: 0,
      agesTo: 12,
      dosis: 200,
      measure: "mg",
      paid: 4,
      price: 10,
      totalSold: 0,
      inventory: 20,
    },
    Medicine {
      name: "Tylenol",
      form: "Capsul",
      cat: "general",
      info: "temporarily relieves minor aches & pains due to: headache, muscular aches, menstrual cramps, the common cold, backache, toothache, minor pain or arthritis.",
      quantity: 40,
      sideEffects: "Symptoms may include: hives, facial swelling, asthma, shock, skin reddening, rash, blisters.",
      direction: "1 capsule every 4 to 6 hours while symptoms persist. If pain or fever does not responsd to 1 capsule, 2 capsules may be used. Do not exceed 6 capsules in 24 hours, unless directed by a doctor.",
      agesFrom: 0,
      agesTo: 12,
      dosis: 200,
      measure: "mg",
      paid: 7,
      price: 13,
      totalSold: 0,
      inventory: 15,
    },
  }
}

func addMedicine(newOne Medicine){
  catalog = append(catalog, newOne)
}

func getTotalMedicines() int {
  return len(catalog)
}

func getMedicine(index int ) *Medicine {
  return &catalog[index]
}

func (medicine *Medicine) getName() string {
  return medicine.name
}

func (medicine *Medicine) getForm() string {
  return medicine.form
}

func (medicine *Medicine) getCategory() string {
  return medicine.cat
}

func (medicine *Medicine) getInfo() string {
  return medicine.info
}

func (medicine *Medicine) getAmount() int {
  return medicine.quantity
}

func (medicine *Medicine) getSideEffects() string {
  return medicine.sideEffects
}

func (medicine *Medicine) getDirections() string {
  return medicine.direction
}

func (medicine *Medicine) getAgesFrom() int {
  return medicine.agesFrom
}

func (medicine *Medicine) getAgesTo() int {
  return medicine.agesTo
}

func (medicine *Medicine) getDosis() int {
  return medicine.dosis
}

func (medicine *Medicine) getMeasure() string {
  return medicine.measure
}

func (medicine *Medicine) getCost() float64 {
  return medicine.paid
}

func (medicine *Medicine) getPrice() float64 {
  return medicine.price
}

func (medicine *Medicine) getTotalSold() float64 {
  return medicine.totalSold
}

func (medicine *Medicine) setTotalSold(amount float64) {
  medicine.totalSold = amount
}

func (medicine *Medicine) getInventory() float64 {
  return medicine.inventory
}

func (medicine *Medicine) setInventory(amount float64) {
  medicine.inventory = amount
}

func (medicine *Medicine) getCurrentProfit() float64 {
  return (medicine.price -  medicine.paid)*medicine.totalSold
}

func (medicine *Medicine) getFutureProfit() float64 {
  return (medicine.price -  medicine.paid)*medicine.inventory
}
