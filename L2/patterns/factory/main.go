package main

import "github.com/dingowd/WB/L2/patterns/factory/culinary"

func main() {
	p := []string{culinary.BunCulinary, culinary.SaladCulinary, culinary.CutletCulinary, "чизкейк"}
	for _, v := range p {
		c := culinary.New(v)
		if c == nil {
			continue
		}
		c.BuyProducts()
		c.Prepare()
		c.Sell()
	}
}

// Абстрактная фабрика (англ. Abstract factory) — порождающий шаблон
// проектирования, предоставляет интерфейс для создания семейств взаимосвязанных
// или взаимозависимых объектов, не специфицируя их конкретных классов. Шаблон
// реализуется созданием абстрактного класса Factory, который представляет собой
// интерфейс для создания компонентов системы (например, для оконного интерфейса
// он может создавать окна и кнопки). Затем пишутся классы, реализующие этот
// интерфейс

// 	Для создания объектов различных типов одним интерфейсом

// + Создание объектов, независимо от их типов и сложности процесса создания

// - Даже для одного объекта необходимо создать соответствующую фабрику, что увеличивает код