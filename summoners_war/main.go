package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"


)

type RunesData struct {
	Unit_list []Unit `json:"unit_list"`
	Rune_list []Rune `json:"runes"`
}

type Unit struct {
	Runes []Rune `json:"runes"`
}

type Rune struct {
	Rune_id int64 `json:"rune_id"`
	Slot_no int32 `json:"slot_no"`

	Occupied_type int32 `json:"occupied_type"`
	Occupied_id  int64 `json:"occupied_id"`

	Rank   int32 `json:"rank"`
	Class  int32 `json:"class"`
	Set_id int32 `json:"set_id"`

	Upgrade_curr int32 `json:"upgrade_curr"`

	Pri_eff    []float64 `json:"pri_eff"`
	Prefix_eff []float64 `json:"prefix_eff"`
	Sec_eff    [][]float64 `json:"sec_eff"`
}

type ComputedRune struct {
	Rune         Rune
	BruiserTotal float64
	TurnOne      float64
	AtkNukeTotal float64
	DefNukeTotal float64
	DebuffTotal  float64
	SpeedTotal  float64
	// To help search for rune in game
	HighestStat string
	SecondHighestStat string
	Slot int32
	Runeset string
}


const (
	HP_FLAT = 1
	HP_PERCENT = 2
	ATK_FLAT = 3
	ATK_PERCENT= 4
	DEF_FLAT = 5
	DEF_PERCENT = 6
	SPEED = 8
	CRATE = 9
	CDMG = 10
	RES = 11
	ACC = 12

	// +20 FROM ORIGINAL MAPPING
	ENERGY = 21
	GUARD = 22
	SWIFT = 23
	BLADE = 24
	RAGE = 25
	FOCUS = 26
	ENDURE = 27
	FATAL = 28
	DESPAIR = 30
	VAMPIRE = 31
	VIOLENT = 33
	NEMESIS = 34
	WILL = 35
	SHIELD = 36
	REVENGE = 37
	DESTROY = 38
	FIGHT = 39
	DETERMINATION = 40
	ENHANCE = 41
	ACCURACY = 42
	TOLERANCE = 43
)

var runesetmap = map[int32]string{
	1: "Energy",
	2: "Guard",
	3: "Swift",
	4: "Blade",
	5: "Rage",
	6: "Focus",
	7: "Endure",
	8: "Fatal",
	10: "Despair",
	11: "Vampire",
	13: "Violent",
	14: "Nemesis",
	15: "Will",
	16: "Shield",
	17: "Revenge",
	18: "Destroy",
	19: "Fight",
	20: "Determination",
	21: "Enhance",
	22: "Accuracy",
	23: "Tolerance",
	99: "Immemorial",
}



func main() {
	fmt.Println("hi")

	runelist := UnmarshalJson()

	datapath := ""
	datapath, _ = os.Getwd()
	datapath += "\\summoners_war\\rune_list.json"

	file, _ := json.MarshalIndent(runelist, "", " ")
	_ = ioutil.WriteFile(datapath, file, 0644)

	datapath, _ = os.Getwd()
	datapath += "\\summoners_war\\computed_rune_list.json"

	computedRuneList := ComputeRunes(runelist)
	file, _ = json.MarshalIndent(computedRuneList, "", " ")
	_ = ioutil.WriteFile(datapath, file, 0644)

}

func UnmarshalJson() (runelist []Rune) {
	datapath, _ := os.Getwd()
	datapath += "\\summoners_war\\data.json"

	fmt.Println(datapath)
	jsonFile, err := os.Open(datapath)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened JSON")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//fmt.Print32ln(byteValue)

	//runes_data := &runes_data{}

	base := RunesData{}

	err = json.Unmarshal(byteValue, &base)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(base.Unit_list[0])

	for _, unit := range base.Unit_list {
		for _, r := range unit.Runes {
			runelist = append(runelist, r)
		}
	}

	for _, r := range base.Rune_list {
		runelist = append(runelist, r)
	}

	return runelist
}

func ComputeRunes(RuneList []Rune) (computed []ComputedRune) {

	for _, rune :=range RuneList{
		computed = append(computed, Compute(rune))
	}

	return computed
}

func Compute(rune Rune) (computedRune ComputedRune) {

	// Go through each sub stat. divided by max total substat. add to respective value
	for _, substat := range rune.Sec_eff {
		switch substat[0] {
		case HP_FLAT:
			computedRune.BruiserTotal = float64((substat[1] + substat[3]) / 2250) // base + grind / max base
		case ATK_FLAT:
			computedRune.AtkNukeTotal = float64((substat[1] + substat[3]) / 110) // base + grind / max base
		case DEF_FLAT:
			computedRune.BruiserTotal = float64((substat[1] + substat[3]) / 110) // base + grind / max base
			computedRune.DefNukeTotal = float64((substat[1] + substat[3]) / 110) // base + grind / max base
		case ACC:
			computedRune.DebuffTotal = float64((substat[1] + substat[3]) / 40) // base + grind / max base
		case RES:
			computedRune.BruiserTotal = float64((substat[1] + substat[3]) / 40) // base + grind / max base
		case HP_PERCENT, DEF_PERCENT:
			computedRune.BruiserTotal = float64((substat[1] + substat[3]) / 35) // base + grind / max base
		case ATK_PERCENT:
			computedRune.AtkNukeTotal = float64((substat[1] + substat[3]) / 35) // base + grind / max base
		case CDMG:
			computedRune.AtkNukeTotal = float64((substat[1] + substat[3]) / 35) // base + grind / max base
			computedRune.DefNukeTotal = float64((substat[1] + substat[3]) / 35) // base + grind / max base
		case CRATE:
			computedRune.AtkNukeTotal = float64((substat[1] + substat[3]) / 30) // base + grind / max base
			computedRune.DefNukeTotal = float64((substat[1] + substat[3]) / 30) // base + grind / max base
		case SPEED:
			computedRune.SpeedTotal = float64((substat[1] + substat[3]) / 30) // base + grind / max base
		}
	}

	switch rune.Prefix_eff[0] {
	case HP_FLAT:
		computedRune.BruiserTotal = float64((rune.Prefix_eff[1]) / 2250) // base + grind / max base
	case ATK_FLAT:
		computedRune.AtkNukeTotal = float64(((rune.Prefix_eff[1])) / 110) // base + grind / max base
	case DEF_FLAT:
		computedRune.BruiserTotal = float64(((rune.Prefix_eff[1])) / 110) // base + grind / max base
		computedRune.DefNukeTotal = float64(((rune.Prefix_eff[1])) / 110) // base + grind / max base
	case ACC:
		computedRune.DebuffTotal = float64(((rune.Prefix_eff[1])) / 40) // base + grind / max base
	case RES:
		computedRune.BruiserTotal = float64(((rune.Prefix_eff[1])) / 40) // base + grind / max base
	case HP_PERCENT, DEF_PERCENT:
		computedRune.BruiserTotal = float64(((rune.Prefix_eff[1])) / 35) // base + grind / max base
	case ATK_PERCENT:
		computedRune.AtkNukeTotal = float64(((rune.Prefix_eff[1])) / 35) // base + grind / max base
	case CDMG:
		computedRune.AtkNukeTotal = float64(((rune.Prefix_eff[1])) / 35) // base + grind / max base
		computedRune.DefNukeTotal = float64(((rune.Prefix_eff[1])) / 35) // base + grind / max base
	case CRATE:
		computedRune.AtkNukeTotal = float64(((rune.Prefix_eff[1])) / 30) // base + grind / max base
		computedRune.DefNukeTotal = float64(((rune.Prefix_eff[1])) / 30) // base + grind / max base
	case SPEED:
		computedRune.SpeedTotal = float64((rune.Prefix_eff[1]) / 30) // base + grind / max base
	}

	computedRune.Rune = rune

	computedRune.Runeset = runesetmap[rune.Set_id]
	computedRune.Slot = rune.Slot_no
	return computedRune
}
