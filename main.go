package main

func main() {
	xmlObjects := xmlParse()
	save_to_db(xmlObjects)
}
