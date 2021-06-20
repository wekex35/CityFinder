package main

import (
	"flag"
	"fmt"
	"net/http"
)

type City struct {
	CityName CityInfo
}

var addr = flag.String("addr", port, "http service address")

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {

	// getCountryList()
	// getCityListByCountry("Chile")
	// searchCity("mum")
	// getContryByCity("Mumbai")
	// getAdjacentCityList("Mumbai")
	// getDataArray()
	// runtime.LockOSThread()

	// flag.Parse()
	// hub := NewHub()
	// go hub.run()
	// // fileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/", http.FileServer(assetFS()))
	http.HandleFunc("/country-list", countryList)
	http.HandleFunc("/countrywise-city", countryCity)
	http.HandleFunc("/search-city", searchCityStr)
	http.HandleFunc("/adjacent-city", adjacentCity)

	// http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))

	// http.HandleFunc("/getIp", getSystemIp)

	open("http://localhost" + port)
	fmt.Println("Server Running at Port :7007")
	http.ListenAndServe(*addr, nil)
	// go func() {
	// 	trayhost.SetUrl("http://localhost" + port)
	// }()

	// // Enter the host system's event loop
	// trayhost.EnterLoop("Barcode Anywhere", iconData)
	// fmt.Println("Exiting")
}
