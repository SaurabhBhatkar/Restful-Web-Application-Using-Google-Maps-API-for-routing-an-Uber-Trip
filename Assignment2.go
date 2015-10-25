package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MyJsonNameResp struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID string   `json:"place_id"`
		Types   []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}

type MyJsonNameReq struct {
	Name     string        `json:"name"`
	Address  string        `json:"address"`
	City     string        `json:"city"`
	State    string        `json:"state"`
	Zip      string        `json:"zip"`
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Location struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"location"`
}

func GetPlannerReq(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Getting Plan Details for ID: %s\n", p.ByName("name"))
	//session, err := mgo.Dial("127.0.0.1")  mongodb://<dbuser>:<dbpassword>@ds045064.mongolab.com:45064/robots1
	session, err := mgo.Dial("mongodb://tr!l)l)lanner:^^&&@ds045064.mongolab.com:45064/robots1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("robots1").C("people")

	if err != nil {
		log.Fatal(err)
	}

	result := MyJsonNameReq{}
	fmt.Println()
	id := p.ByName("name")

	oid := bson.ObjectIdHex(id)
	c.FindId(oid).One(&result)

	if err != nil {
		log.Fatal(err)
	}
	b2, err := json.Marshal(result)
	if err != nil {

	}
	fmt.Fprintf(rw, string(b2))

}

func PostPlannerReq(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	var myjson3 MyJsonNameReq
	d1 := json.NewDecoder(req.Body)
	err := d1.Decode(&myjson3)

	Query1 := "http://maps.google.com/maps/api/geocode/json?address="

	Query2 := myjson3.Address + " " + myjson3.City + " " + myjson3.State + "&sensor=false"

	Query2 = strings.Replace(Query2, " ", "+", -1)

	URL := Query1 + Query2

	fmt.Println("MyUrl=: " + URL)
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	robots2, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var myjsonnameresp1 MyJsonNameResp

	err = json.Unmarshal(robots2, &myjsonnameresp1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(myjsonnameresp1.Results[0].Geometry.Location.Lat)
	fmt.Println(myjsonnameresp1.Results[0].Geometry.Location.Lng)

	myjson3.Id = bson.NewObjectId()

	myjson3.Location.Lat = myjsonnameresp1.Results[0].Geometry.Location.Lat
	myjson3.Location.Lng = myjsonnameresp1.Results[0].Geometry.Location.Lng

	session, err := mgo.Dial("mongodb://tr!l)l)lanner:^^&&@ds045064.mongolab.com:45064/robots1")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("robots1").C("people")
	err = c.Insert(&myjson3)

	if err != nil {
		log.Fatal(err)
	}

	result := MyJsonNameReq{}
	fmt.Println()

	id := myjson3.Id.Hex()

	oid := bson.ObjectIdHex(id)
	c.FindId(oid).One(&result)

	if err != nil {
		log.Fatal(err)
	}
	b2, err := json.Marshal(result)
	if err != nil {

	}
	fmt.Fprintf(rw, string(b2))

}

func UpdatePlannerReq(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Updating Planner Record for ID: %s!\n", p.ByName("name"))

	var myjson3 MyJsonNameReq
	d1 := json.NewDecoder(req.Body)
	err := d1.Decode(&myjson3)

	Query1 := "http://maps.google.com/maps/api/geocode/json?address="

	Query2 := myjson3.Address + " " + myjson3.City + " " + myjson3.State + "&sensor=false"

	Query2 = strings.Replace(Query2, " ", "+", -1)

	URL := Query1 + Query2

	fmt.Println("MyUrl=: " + URL)
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	robots2, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var myjsonnameresp1 MyJsonNameResp

	err = json.Unmarshal(robots2, &myjsonnameresp1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(myjsonnameresp1.Results[0].Geometry.Location.Lat)
	fmt.Println(myjsonnameresp1.Results[0].Geometry.Location.Lng)

	myjson3.Id = bson.NewObjectId()

	myjson3.Location.Lat = myjsonnameresp1.Results[0].Geometry.Location.Lat
	myjson3.Location.Lng = myjsonnameresp1.Results[0].Geometry.Location.Lng

	session, err := mgo.Dial("mongodb://tr!l)l)lanner:^^&&@ds045064.mongolab.com:45064/robots1")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("robots1").C("people")

	if err != nil {
		log.Fatal(err)
	}

	result := MyJsonNameReq{}

	fmt.Println()

	id := p.ByName("name")
	oid := bson.ObjectIdHex(id)
	c.FindId(oid).One(&result)

	if myjson3.Name != "" {
		result.Name = myjson3.Name
	}

	if myjson3.Address != "" {
		result.Address = myjson3.Address
	}
	if myjson3.City != "" {
		result.City = myjson3.City
	}

	if myjson3.State != "" {
		result.State = myjson3.State
	}
	if myjson3.Zip != "" {
		result.Zip = myjson3.Zip
	}

	result.Location.Lat = myjsonnameresp1.Results[0].Geometry.Location.Lat

	result.Location.Lng = myjsonnameresp1.Results[0].Geometry.Location.Lng

	err = c.UpdateId(oid, result)

	fmt.Println(result.Location.Lat)

	fmt.Println(result.Location.Lng)

	if err != nil {
		log.Fatal(err)
	}
	b2, err := json.Marshal(result)
	if err != nil {

	}

	fmt.Fprintf(rw, string(b2))

}
func DeletePlannerReq(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	session, err := mgo.Dial("mongodb://tr!l)l)lanner:^^&&@ds045064.mongolab.com:45064/robots1")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("robots1").C("people")
	id := p.ByName("name")
	oid := bson.ObjectIdHex(id)

	c.RemoveId(oid)

	fmt.Fprintf(rw, "Deleted Record, %s!\n", p.ByName("name"))

}

func main() {
	mux := httprouter.New()
	mux.GET("/locations/:name", GetPlannerReq)
	mux.POST("/locations", PostPlannerReq)
	mux.PUT("/locations/:name", UpdatePlannerReq)
	mux.DELETE("/locations/:name", DeletePlannerReq)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
