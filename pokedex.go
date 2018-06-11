package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
	"strings"
	"strconv"
)

type Type struct {
	// Name of the type
	Name string `json:"name"`
	// The effective types, damage multiplize 2x
	EffectiveAgainst []string `json:"effectiveAgainst"`
	// The weak types that against, damage multiplize 0.5x
	WeakAgainst []string `json:"weakAgainst"`
}

type Pokemon struct {
	Number         string   `json:"Number"`
	Name           string   `json:"Name"`
	Classification string   `json:"Classification"`
	TypeI          []string `json:"Type I"`
	TypeII         []string `json:"Type II,omitempty"`
	Weaknesses     []string `json:"Weaknesses"`
	FastAttackS    []string `json:"Fast Attack(s)"`
	Weight         string   `json:"Weight"`
	Height         string   `json:"Height"`
	Candy          struct {
		Name     string `json:"Name"`
		FamilyID int    `json:"FamilyID"`
	} `json:"Candy"`
	NextEvolutionRequirements struct {
		Amount int    `json:"Amount"`
		Family int    `json:"Family"`
		Name   string `json:"Name"`
	} `json:"Next Evolution Requirements,omitempty"`
	NextEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Next evolution(s),omitempty"`
	PreviousEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Previous evolution(s),omitempty"`
	SpecialAttacks      []string `json:"Special Attack(s)"`
	BaseAttack          int      `json:"BaseAttack"`
	BaseDefense         int      `json:"BaseDefense"`
	BaseStamina         int      `json:"BaseStamina"`
	CaptureRate         float64  `json:"CaptureRate"`
	FleeRate            float64  `json:"FleeRate"`
	BuddyDistanceNeeded int      `json:"BuddyDistanceNeeded"`
}

// Move is an attack information. The
type Move struct {
	// The ID of the move
	ID int `json:"id"`
	// Name of the attack
	Name string `json:"name"`
	// Type of attack
	Type string `json:"type"`
	// The damage that enemy will take
	Damage int `json:"damage"`
	// Energy requirement of the attack
	Energy int `json:"energy"`
	// Dps is Damage Per Second
	Dps float64 `json:"dps"`
	// The duration
	Duration int `json:"duration"`
}

// BaseData is a struct for reading data.json
type BaseData struct {
	Types    []Type    `json:"types"`
	Pokemons []Pokemon `json:"pokemons"`
	Moves    []Move    `json:"moves"`
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/list url:", r.URL)

	//var queryParams = r.URL.Query()
	var typeQuery = strings.ToLower(r.URL.Query().Get("type"))

	if(len(typeQuery) > 0){

		if(isTypeValid(typeQuery)){
			fmt.Println(typeQuery+" type pokemons :" )
			printPokemonInfo(typeQuery)
		}else{
			fmt.Println("There is no type called : " + typeQuery)
		}

	}

	if(r.URL.Path == "/list/types"){
		fmt.Println("Pokemon Types:")
		for i := 0; i < len(baseData.Pokemons); i++ {
			fmt.Println(baseData.Types[i].Name)
		}
	}

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/get url:", r.URL)
	fmt.Fprint(w, "The Get Handler\n")
}

func otherwise(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}


var baseData BaseData
//the function to read json data.
func jsonReader() BaseData {

	raw, err := ioutil.ReadFile("data.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c BaseData
	json.Unmarshal(raw, &c)
	return c
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func printPokemonInfo(queryParam string) {

	for i := 0; i < len(baseData.Pokemons); i++ {
		if(strings.ToLower(baseData.Pokemons[i].TypeI[0]) == queryParam){

			fmt.Println(baseData.Pokemons[i].Name+":" )
			fmt.Println("\t Type I: " + baseData.Pokemons[i].TypeI[0])
			if(baseData.Pokemons[i].TypeII != nil){
				fmt.Println("\t Type II: " + baseData.Pokemons[i].TypeII[0])
			}
			fmt.Println("\t Height: " + baseData.Pokemons[i].Height)
			fmt.Println("\t Weight: " + baseData.Pokemons[i].Weight)
			fmt.Println("\t Base Attack: " + strconv.Itoa(baseData.Pokemons[i].BaseAttack))
			fmt.Println("\t Base Defense: " + strconv.Itoa(baseData.Pokemons[i].BaseDefense))
			fmt.Println("\t Base Stamina: " + strconv.Itoa(baseData.Pokemons[i].BaseStamina))
			fmt.Println("\t Buddy Distance Needed: " + strconv.Itoa(baseData.Pokemons[i].BuddyDistanceNeeded))
			fmt.Println("\t Candy Name: " + baseData.Pokemons[i].Candy.Name)
			fmt.Println("\t CaptureRate: " + FloatToString(baseData.Pokemons[i].CaptureRate))
			fmt.Println("\t Classification: " + baseData.Pokemons[i].Classification)
			fmt.Println("\t FastAttackS: ")
			for j := 0; j < len(baseData.Pokemons[i].FastAttackS);j++{
				fmt.Println("\t\t " + baseData.Pokemons[i].FastAttackS[j])
			}
			fmt.Println("\t FleeRate: " + FloatToString(baseData.Pokemons[i].FleeRate))
			fmt.Println("\t Weaknesses: ")
			for k := 0; k < len(baseData.Pokemons[i].Weaknesses);k++{
				fmt.Println("\t\t " + baseData.Pokemons[i].Weaknesses[k])
			}
			fmt.Println("\t Next Evolution Requirements: " )
			fmt.Println("\t\t Name: " + baseData.Pokemons[i].NextEvolutionRequirements.Name)
			fmt.Println("\t\t Amount: " + strconv.Itoa(baseData.Pokemons[i].NextEvolutionRequirements.Amount))

			if(baseData.Pokemons[i].NextEvolutions != nil){
				fmt.Println("\t Next Evolutions: ")
				for l := 0; l < len(baseData.Pokemons[i].NextEvolutions);l++{
					fmt.Println("\t\t Name: " + string(baseData.Pokemons[i].NextEvolutions[l].Name))
				}

			}
			if(baseData.Pokemons[i].PreviousEvolutions != nil){
				fmt.Println("\t Previous Evolutions: ")
				for l := 0; l < len(baseData.Pokemons[i].PreviousEvolutions);l++{
					fmt.Println("\t\t Name: " + string(baseData.Pokemons[i].PreviousEvolutions[l].Name))
				}

			}


		}
	}
}
func isTypeValid(queryParam string) bool{
	for i := 0; i<len(baseData.Types);i++{
		if(queryParam == strings.ToLower(baseData.Types[i].Name)){
			return true
		}
	}
	return false
}


func main() {
	//TODO: read data.json to a BaseData
	baseData  = jsonReader()

	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/get", getHandler)


	//List handlers
	http.HandleFunc("/list/types", listHandler)


	//TODO: add more
	http.HandleFunc("/", otherwise)
	log.Println("starting server on :8080")
	http.ListenAndServe(":8080", nil)
}