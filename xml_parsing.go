package main

import (
	"encoding/xml"
	"fmt"
	"github.com/rs/xid"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

func midPoint(xmin, xmax, ymin, ymax int) (float64, float64) {
	return float64(xmin+xmax) * 0.5, float64(ymin+ymax) * 0.5
}
func xmlParse() []Img {
	var img_number = 86
	var max_img_number = 376

	xmlPath := "shop/youtube_shop0"
	var imgs []Img
	var objects = make(map[string]Object)
	for img_number <= max_img_number {
		//creating path to xml file
		xmlFilename := xmlPath
		if img_number < 100 {
			xmlFilename += "0"
		}
		xmlFilename += strconv.Itoa(img_number) + ".xml"

		//load data
		xmlFile, err := os.Open(xmlFilename)
		defer xmlFile.Close()
		byteValue, _ := ioutil.ReadAll(xmlFile)
		if err != nil {
			fmt.Println(err)
		}
		var annotation Annotation
		xml.Unmarshal(byteValue, &annotation)

		var image = Img{ID: img_number}
		for i := 0; i < len(annotation.Objects); i++ {
			annotationObject := annotation.Objects[i]
			midPointX, midPointY := midPoint(annotationObject.BndBox.Xmin, annotationObject.BndBox.Xmax, annotationObject.BndBox.Ymin, annotationObject.BndBox.Ymax)
			found := false
			var guid = ""
			var min_x = 100.

			for _, v := range objects {
				dif_x := math.Abs(v.Midpoint.x - midPointX)
				dif_y := math.Abs(v.Midpoint.y-midPointY)
				if dif_x <= min_x && dif_y < 100. {
					//use ID from the previous pictures if a slight difference between objects' midpoints
					guid = v.ID
					min_x = dif_x
					found = true
				}
			}
			if !found {
				//new object
				guid = xid.New().String()
			}
			var newObject = Object{ID: guid, Midpoint: MidPoint{midPointX, midPointY}, BndBox: annotationObject.BndBox}
			objects[guid] = newObject
			image.Objects = append(image.Objects, newObject)

		}
		imgs = append(imgs, image)
		img_number += 10
	}
	return imgs
}
