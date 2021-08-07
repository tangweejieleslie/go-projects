package main
//
//type BaseData struct {
//	Command string `json:"command"`
//}
//
//type ComputedRune struct {
//	Rune         Rune
//	BruiserTotal int
//	TurnOne      int
//	AtkNukeTotal int
//	DefNukeTotal int
//	DebuffTotal  int
//	SpeedTotal int
//}
//
//
//
//type Nuker struct {
//	TotalSpeed int
//	TotalAtkP  int
//	TotalAtkF  int
//	TotalCR    int
//	TotalCD    int
//}
//
//type RunesData struct {
//	Unit_list []Unit `json:"unit_list"`
//	Rune_list []Rune `json:"runes"`
//}
//
//type Unit struct {
//	Runes []Rune `json:"runes"`
//}
//
//type Rune struct {
//	Rune_id int32 `json:"rune_id"`
//	Slot_no int32 `json:"slot_no"`
//
//	Occupied_type int32 `json:"occupied_type"`
//	Occupited_id  int32 `json:"occupited_id"`
//
//	Rank   int32 `json:"rank"`
//	Class  int32 `json:"class"`
//	Set_id int32 `json:"set_id"`
//
//	Upgrade_curr int32 `json:"upgrade_curr"`
//
//	Pri_eff    []int32 `json:"pri_eff"`
//	Prefix_eff []int32 `json:"prefix_eff"`
//	Sec_eff    [][]int `json:"sec_eff"`
//}
//
//type SecEff struct {
//	Stat        int32
//	Value       int32
//	Grind       int32
//	Grind_value int32
//}
//
////func GetType() {
////
////	StatType := map[int]string{
////		1:  "HP flat",
////		2:  "HP%",
////		3:  "ATK flat",
////		4:  "ATK%",
////		5:  "DEF flat",
////		6:  "DEF%",
////		8:  "SPD",
////		9:  "CRate",
////		10: "CDmg",
////		11: "RES",
////		12: "ACC",
////	}
////
////	RuneSet := map[int]string{
////		1:  "Energy",
////		2:  "Guard",
////		3:  "Swift",
////		4:  "Blade",
////		5:  "Rage",
////		6:  "Focus",
////		7:  "Endure",
////		8:  "Fatal",
////		10: "Despair",
////		11: "Vampire",
////		13: "Violent",
////		14: "Nemesis",
////		15: "Will",
////		16: "Shield",
////		17: "Revenge",
////		18: "Destroy",
////		19: "Fight",
////		20: "Determination",
////		21: "Enhance",
////		22: "Accuracy",
////		23: "Tolerance",
////		99: "Immemorial",
////	}
////	RuneClass := map[int]string{
////		0: "Common",
////		1: "Magic",
////		2: "Rare",
////		3: "Hero",
////		4: "Legendary",
////	}
////	RuneQuality := map[int]string{
////
////		1: "Common",
////		2: "Magic",
////		3: "Rare",
////		4: "Hero",
////		5: "Legend",
////		// ancient rune qualities
////		11: "Common",
////		12: "Magic",
////		13: "Rare",
////		14: "Hero",
////		15: "Legend",
////	}
////}
