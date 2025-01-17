package main

import "fmt"

type Component interface {
	Search(spec Specification)
}

type File struct {
	Name  string
	Ext   string
	Value int
}

func (f *File) Search(spec Specification) {
	if spec.IsSatisfiedBy(f) {
		fmt.Printf("File: %s, Ext: %s, Value: %d\n", f.Name, f.Ext, f.Value)
	}
}

type Folder struct {
	Name       string
	components []Component
}

func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}

func (f *Folder) Search(spec Specification) {
	for _, c := range f.components {
		c.Search(spec)
	}
}

type Specification interface {
	IsSatisfiedBy(p *File) bool
}

type NameSpecification struct {
	Name string
}

func (c *NameSpecification) IsSatisfiedBy(product *File) bool {
	return product.Name == c.Name
}

type ExtSpecification struct {
	ext string
}

func (c *ExtSpecification) IsSatisfiedBy(product *File) bool {
	return product.Ext == c.ext
}

type AndSpecification struct {
	First, Second Specification
}

func (a *AndSpecification) IsSatisfiedBy(product *File) bool {
	return a.First.IsSatisfiedBy(product) && a.Second.IsSatisfiedBy(product)
}

type OrSpecification struct {
	First, Second Specification
}

func (o *OrSpecification) IsSatisfiedBy(product *File) bool {
	return o.First.IsSatisfiedBy(product) || o.Second.IsSatisfiedBy(product)
}

type NotSpecification struct {
	Spec Specification
}

func (n *NotSpecification) IsSatisfiedBy(product *File) bool {
	return !n.Spec.IsSatisfiedBy(product)
}

func FilterProducts(component Component, spec Specification) {
	component.Search(spec)
}

func main() {
	//products := []Product{
	//	{"Laptop", 5, "Electronics"},
	//	{"Smartphone", 500, "Electronics"},
	//	{"Coffee", 10, "Groceries"},
	//	{"Bread", 2, "Groceries"},
	//}
	//priceSpec := PriceSpecification{Min: 5, Max: 500}
	//categorySpec := CategorySpecification{Category: "Electronics"}
	//andSpec := AndSpecification{First: &priceSpec, Second: &categorySpec}
	//filteredProducts := FilterProducts(products, &andSpec)
	//for _, p := range filteredProducts {
	//	fmt.Printf("Product: %s, Price: %.2f, Category: %s\n", p.Name, p.Price, p.Category)
	//}
	fol := Folder{
		Name: "sahil",
	}
	file1 := File{
		Name:  "yash",
		Ext:   "txt",
		Value: 1,
	}
	file2 := File{
		Name: "yash",
		Ext:  "pdf",
	}
	fol.Add(&file1)
	fol.Add(&file2)
	nameSpec := NameSpecification{
		Name: "yash",
	}
	extSpec := ExtSpecification{
		ext: "txt",
	}
	ANDSpec := AndSpecification{
		First:  &nameSpec,
		Second: &extSpec,
	}
	FilterProducts(&fol, &ANDSpec)
}
