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
	"sort"
)

var baseData BaseData

// BaseData is a struct for reading data.json
type BaseData struct {
	Types    []Type    `json:"types"`
	Pokemons []Pokemon `json:"pokemons"`
	Moves    []Move    `json:"moves"`
}

func listHandler(w http.ResponseWriter, r *http.Request) {

	//If user wants to list types with sorting parameters.
	var sortQuery = strings.ToLower(r.URL.Query().Get("sortby"))
	sortBaseData(sortQuery)


	//If user wants to list types
	if(r.URL.Path == "/list/types"){
		fmt.Println("Types:")
		for i := 0; i < len(baseData.Types); i++ {
			fmt.Println("\t" + baseData.Types[i].Name)
		}


	//If user wants to list Pokemons
	}else if(r.URL.Path == "/list/pokemons"){


		//If user wants to list pokemons certain type.
		var typeQ = strings.ToLower(r.URL.Query().Get("type"))

		if(len(typeQ)>0){
			if(isTypeValid(typeQ)){
				fmt.Println(typeQ + " type pokemons")
				for i:=0; i<len(baseData.Pokemons);i++{
					if(strings.ToLower(baseData.Pokemons[i].TypeI[0]) == typeQ){
						printPokemonInfo(baseData.Pokemons[i])
					}
				}
			}else{
				fmt.Println("Is that a valid type ?")
			}
		}else{
			for i:=0; i<len(baseData.Pokemons);i++{
				printPokemonInfo(baseData.Pokemons[i])
			}
		}


	//If user wants to list moves
	}else if(r.URL.Path == "/list/moves"){

		fmt.Println("Moves:")
		for i := 0; i < len(baseData.Moves); i++ {
			fmt.Println("\t" + baseData.Moves[i].Name)
		}
	// List of endpoints that users can call.
	}else if(r.URL.Path == "/help"){

		fmt.Println("\t get/pokemon")
		fmt.Println("\t get/type")
		fmt.Println("\t get/move")
		fmt.Println("You can call the endpoints that above with name or type filters.")
		fmt.Println()

		fmt.Println("\t list/pokemons")
		fmt.Println("\t list/types")
		fmt.Println("\t list/moves")
		fmt.Println("You can call the endpoints that above and sort them with sortby parameter.")
		fmt.Println()

		fmt.Println("You can sort the pokemons by base attack, base stamina, base defense")
		fmt.Println("You can sort the moves by damage, energy, dps, duration")
	}

}

func getHandler(w http.ResponseWriter, r *http.Request) {

	//First, get query param if exists.
	var queryParam  = getQueryParam(r)[0]



	//If user wants to get certain pokemon
	if(r.URL.Path == "/get/pokemon"){

		switch queryParam {
		case "name":
			if(isTypeValid(queryParam)) {
				var pokemonName = strings.ToLower(r.URL.Query().Get("name"))
				for i := 0; i < len(baseData.Moves); i++ {
					if (strings.ToLower(baseData.Pokemons[i].Name) == pokemonName) {
						printMoveInfo(baseData.Moves[i])
					}
				}
			}else{
				fmt.Println("Is that a valid name ?")

			}
			break
		case "type":
			if(isTypeValid(queryParam)) {
				var pokemonType= strings.ToLower(r.URL.Query().Get("type"))
				for i := 0; i < len(baseData.Moves); i++ {
					if (strings.ToLower(baseData.Pokemons[i].TypeI[0]) == pokemonType) {
						printMoveInfo(baseData.Moves[i])
					}
				}
			}else{
				fmt.Println("Is that a valid type ?")

			}
			break
		default:
			fmt.Println("Pokemon's what ???")
		}

	}

	//If user wants to get certain move
	if(r.URL.Path == "/get/move"){


		switch queryParam {
		case "name":

			if(isTypeValid(queryParam)) {
				var moveName= strings.ToLower(r.URL.Query().Get("name"))
				for i := 0; i < len(baseData.Moves); i++ {
					if (strings.ToLower(baseData.Moves[i].Name) == moveName) {
						printMoveInfo(baseData.Moves[i])
					}
				}
			}else{
				fmt.Println("Is that a valid name ?")
			}
			break
		case "type":

			if(isTypeValid(queryParam)) {
				var moveType= strings.ToLower(r.URL.Query().Get("type"))
				for i := 0; i < len(baseData.Moves); i++ {
					if (strings.ToLower(baseData.Moves[i].Type) == moveType) {
						printMoveInfo(baseData.Moves[i])
					}
				}
			}else{
				fmt.Println("Is that a valid type ?")

			}
			break
		default:
			fmt.Println("Move's what ???")

		}
	}
	//If user wants to get certain type
	if(r.URL.Path == "/get/type"){

		var typeName = strings.ToLower(r.URL.Query().Get("name"))
		if(isTypeValid(typeName)){
			for i := 0; i < len(baseData.Types); i++ {
				if(strings.ToLower(baseData.Types[i].Name) == typeName){
					printTypeInfo(baseData.Types[i])
				}
			}
		}else{
			fmt.Println("Is that a valid type ?")
		}
	}
}


func otherwise(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Oops, there is no such a thing :(")
}

func getQueryParam(r *http.Request)[]string{
	var keys []string
	for key, _ := range r.URL.Query(){
		keys = append(keys, key)
	}
	return keys
}

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


//to sort every wanted query if query includes sortby param.
func sortBaseData(sortBy string){
	sortBy = strings.ToLower(sortBy)

	if(len(sortBy) > 0){
		fmt.Println("Sorted by: " + sortBy)
		if(sortBy == "baseattack"){
			sort.Slice(baseData.Pokemons[:], func(i, j int) bool {
				return baseData.Pokemons[i].BaseAttack > baseData.Pokemons[j].BaseAttack
			})
		}else if(sortBy == "basedefense"){
			sort.Slice(baseData.Pokemons[:], func(i, j int) bool {
				return baseData.Pokemons[i].BaseDefense > baseData.Pokemons[j].BaseDefense
			})
		}else if(sortBy == "basestamina"){
			sort.Slice(baseData.Pokemons[:], func(i, j int) bool {
				return baseData.Pokemons[i].BaseStamina > baseData.Pokemons[j].BaseStamina
			})
		}else if(sortBy == "damage"){
			sort.Slice(baseData.Moves[:], func(i, j int) bool {
				return baseData.Moves[i].Damage > baseData.Moves[j].Damage
			})
		}else if(sortBy == "energy"){
			sort.Slice(baseData.Moves[:], func(i, j int) bool {
				return baseData.Moves[i].Energy > baseData.Moves[j].Energy
			})
		}else if(sortBy == "dps"){
			sort.Slice(baseData.Moves[:], func(i, j int) bool {
				return baseData.Moves[i].Dps > baseData.Moves[j].Dps
			})
		}else if(sortBy == "duration"){
			sort.Slice(baseData.Moves[:], func(i, j int) bool {
				return baseData.Moves[i].Duration > baseData.Moves[j].Duration
			})
		}

	}

}



func printPokemonInfo(pokemon Pokemon) {

	fmt.Println(pokemon.Name+":" )
	fmt.Println("\t Type I: " + pokemon.TypeI[0])
	if(pokemon.TypeII != nil){
		fmt.Println("\t Type II: " + pokemon.TypeII[0])
	}
	fmt.Println("\t Height: " + pokemon.Height)
	fmt.Println("\t Weight: " + pokemon.Weight)
	fmt.Println("\t Base Attack: " + strconv.Itoa(pokemon.BaseAttack))
	fmt.Println("\t Base Defense: " + strconv.Itoa(pokemon.BaseDefense))
	fmt.Println("\t Base Stamina: " + strconv.Itoa(pokemon.BaseStamina))
	fmt.Println("\t Buddy Distance Needed: " + strconv.Itoa(pokemon.BuddyDistanceNeeded))
	fmt.Println("\t Candy Name: " + pokemon.Candy.Name)
	fmt.Println("\t CaptureRate: " + FloatToString(pokemon.CaptureRate))
	fmt.Println("\t FastAttackS: ")
	for j := 0; j < len(pokemon.FastAttackS);j++{
		fmt.Println("\t\t " + pokemon.FastAttackS[j])
	}
	fmt.Println("\t FleeRate: " + FloatToString(pokemon.FleeRate))
	fmt.Println("\t Weaknesses: ")
	for k := 0; k < len(pokemon.Weaknesses);k++{
		fmt.Println("\t\t " + pokemon.Weaknesses[k])
	}
	fmt.Println("\t Next Evolution Requirements: " )
	fmt.Println("\t\t Name: " + pokemon.NextEvolutionRequirements.Name)
	fmt.Println("\t\t Amount: " + strconv.Itoa(pokemon.NextEvolutionRequirements.Amount))

	if(pokemon.NextEvolutions != nil){
		fmt.Println("\t Next Evolutions: ")
		for l := 0; l < len(pokemon.NextEvolutions);l++{
			fmt.Println("\t\t Name: " + string(pokemon.NextEvolutions[l].Name))
		}
	}
	if(pokemon.PreviousEvolutions != nil){
		fmt.Println("\t Previous Evolutions: ")
		for l := 0; l < len(pokemon.PreviousEvolutions);l++{
			fmt.Println("\t\t Name: " + string(pokemon.PreviousEvolutions[l].Name))
		}
	}
}

func printMoveInfo(move Move) {

	fmt.Println(move.Name+":" )

	fmt.Println("\t Type: " + move.Type)
	fmt.Println("\t Damage: " + strconv.Itoa(move.Damage))
	fmt.Println("\t Duration: " + strconv.Itoa(move.Duration))
	fmt.Println("\t Dps: " + FloatToString(move.Dps))
	fmt.Println("\t Energy: " + strconv.Itoa(move.Energy))

}

func printTypeInfo(typ Type) {

	fmt.Println(typ.Name+":" )

	fmt.Println("\t Effective Against: ")
	for l := 0; l < len(typ.EffectiveAgainst);l++{
		fmt.Println("\t\t: " + typ.EffectiveAgainst[l])
	}
	fmt.Println("\t Weak Against: ")
	for l := 0; l < len(typ.WeakAgainst);l++{
		fmt.Println("\t\t: " + typ.WeakAgainst[l])
	}
	fmt.Println("\t Example Pokemons:")
	for j := 0; j< len(baseData.Pokemons);j++ {
		if(baseData.Pokemons[j].TypeI[0] == typ.Name && j<3){
			fmt.Println("\t\t " + baseData.Pokemons[j].Name)
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

	baseData  = jsonReader()


	//Get handlers
	http.HandleFunc("/get/type", getHandler)
	http.HandleFunc("/get/pokemon", getHandler)
	http.HandleFunc("/get/move", getHandler)

	//List handlers
	http.HandleFunc("/list/types", listHandler)
	http.HandleFunc("/list/pokemons", listHandler)
	http.HandleFunc("/list/moves", listHandler)
	http.HandleFunc("/help", listHandler)




	http.HandleFunc("/", otherwise)
	log.Println("starting server on :8080")
	http.ListenAndServe(":8080", nil)
}