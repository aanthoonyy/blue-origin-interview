package main

func main() {
	todos := Todos{}
	todos.add("Learn Go")
	todos.add("Learn docker")
	todos.toggle(0)
	todos.print()
}
