package main

import "fmt"

//region  Структура и Builder

// Pizza Настраиваемый с помощью строителя PizzaBuilder заказ пиццы
type Pizza struct {
	Size      string
	Crust     string
	Cheese    bool
	Pepperoni bool
	Mushrooms bool
	// остальные ингредиенты...
}

// PizzaBuilder Интерфейс реализует установку всех необходимых полей структуры Pizza
type PizzaBuilder interface {
	SetSize(size string) PizzaBuilder
	SetCrust(crust string) PizzaBuilder
	AddCheese() PizzaBuilder
	AddPepperoni() PizzaBuilder
	AddMushrooms() PizzaBuilder

	// Для остальных ингредиентов...

	Build() Pizza // Обратите внимание, что все методы интерфейса, кроме этого, возвращают тип PizzaBuilder, и только это т возвращает Pizza
}

//endregion

//region Concrete Builder с реализациями интерфейса

// ConcretePizzaBuilder Структура, реализующая специфические метод для настройки требуемой пиццы. Эта структура должна реализовать все методы интерфейса PizzaBuilder
type ConcretePizzaBuilder struct {
	pizza Pizza
}

func (b *ConcretePizzaBuilder) SetSize(size string) PizzaBuilder {
	b.pizza.Size = size
	return b
}
func (b *ConcretePizzaBuilder) SetCrust(crust string) PizzaBuilder {
	b.pizza.Crust = crust
	return b
}
func (b *ConcretePizzaBuilder) AddCheese() PizzaBuilder {
	b.pizza.Cheese = true
	return b
}
func (b *ConcretePizzaBuilder) AddPepperoni() PizzaBuilder {
	b.pizza.Pepperoni = true
	return b
}
func (b *ConcretePizzaBuilder) AddMushrooms() PizzaBuilder {
	b.pizza.Mushrooms = true
	return b
}

func (b *ConcretePizzaBuilder) Build() Pizza {
	return b.pizza
}

//endregion

//region Director (предустановленные)

type PizzaDirector struct{}

func (d *PizzaDirector) CreateMargherita(builder PizzaBuilder) Pizza {
	return builder.SetSize("Medium").SetCrust("Thin").AddCheese().Build()
}

// ... можно добавить еще предустановленных рецептов пиццы

//endregion

func main() {
	// создаем Builder и Director
	builder := &ConcretePizzaBuilder{}
	director := &PizzaDirector{}

	// создаем пиццу по готовому рецепту
	margherita := director.CreateMargherita(builder)
	fmt.Printf("Margherita:%#v\n", margherita)

	// кастомная пицца
	customPizza := builder.SetSize("Large").AddMushrooms().AddCheese().AddPepperoni().Build()
	fmt.Printf("Your pizza:%#v\n", customPizza)
}
