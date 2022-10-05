package main

import "fmt"

const (
	pequeño = "Pequeño"
	mediano = "Mediano"
	grande  = "Grande"
)

type tienda struct {
	Productos []producto
}
type producto struct {
	Tipo   string
	Nombre string
	Precio float64
}

func (p producto) CalcularCosto() float64 {
	var r float64
	if p.Tipo == pequeño {
		r = p.Precio
	} else if p.Tipo == mediano {
		r = p.Precio + (p.Precio * 0.03)
	} else if p.Tipo == grande {
		r = p.Precio + (p.Precio * 0.06) + 2500
	}
	return r
}

type Producto interface {
	CalcularCosto() float64
}

func (t tienda) Total() (result float64) {
	for _, value := range t.Productos {
		result += value.CalcularCosto()
	}
	return
}
func (t tienda) Agregar(p producto) {
	t.Productos = append(t.Productos, p)
	fmt.Println(t.Productos)
}

type Ecommerce interface {
	Total() float64
	Agregar(p producto)
}

func nuevoProducto(tipo string, nombre string, precio float64) Producto {
	return &producto{tipo, nombre, precio}
}
func nuevaTienda() Ecommerce {
	return &tienda{}
}

func main() {
	p1 := producto{pequeño, "Producto1", 15000.95}
	p2 := producto{grande, "Producto1", 15000.95}
	ti := nuevaTienda()
	ti.Agregar(p1)
	fmt.Println("Agrega a la tienda un nuevo producto")
	ti.Agregar(p2)
	fmt.Println("Agrega a la tienda un nuevo producto")
	fmt.Println(ti.Total())
}
