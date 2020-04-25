package main

import "encoding/xml"

type Annotation struct {
	XMLName xml.Name `xml:"annotation"`
	Objects []Object `xml:"object"`
}
type Object struct {
	XMLName  xml.Name `xml:"object"`
	BndBox   BndBox   `xml:"bndbox"`
	ID       string
	Midpoint MidPoint
}

type BndBox struct {
	Xmin int `xml:"xmin"`
	Ymin int `xml:"ymin"`
	Xmax int `xml:"xmax"`
	Ymax int `xml:"ymax"`
}

type MidPoint struct {
	x float64
	y float64
}

type Img struct {
	ID      int
	Objects []Object
}
