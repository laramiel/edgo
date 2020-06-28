package journal

import (
	"encoding/json"
	"log"
	"testing"
)

func Parse(t *testing.T, input []byte) (interface{}, error) {
	obj, err := ParseJournalLine(input)
	if err == nil {
		return obj, err
	}

	if x, ok := err.(*json.UnmarshalFieldError); ok {
		t.Error(x)
	}
	if x, ok := err.(*json.UnmarshalTypeError); ok {
		t.Error(x)
	}
	if x, ok := err.(*json.UnsupportedTypeError); ok {
		t.Error(x)
	}
	if x, ok := err.(*json.UnsupportedValueError); ok {
		t.Error(x)
	}
	t.Error(err)
	return obj, err
}

func TestParseStatistics(t *testing.T) {
	obj, err := Parse(t, []byte(`{ "timestamp":"2020-06-27T23:33:28Z", "event":"Statistics", "Bank_Account":{ "Current_Wealth":1104888856, "Spent_On_Ships":187581397, "Spent_On_Outfitting":263525084, "Spent_On_Repairs":68367, "Spent_On_Fuel":43951, "Spent_On_Ammo_Consumables":38152, "Insurance_Claims":4, "Spent_On_Insurance":5259365, "Owned_Ship_Count":5 }, "Combat":{ "Bounties_Claimed":4, "Bounty_Hunting_Profit":53739, "Combat_Bonds":0, "Combat_Bond_Profits":0, "Assassinations":0, "Assassination_Profits":0, "Highest_Single_Reward":35250, "Skimmers_Killed":3 }, "Crime":{ "Notoriety":0, "Fines":4, "Total_Fines":1600, "Bounties_Received":2, "Total_Bounties":400, "Highest_Bounty":200 }, "Smuggling":{ "Black_Markets_Traded_With":0, "Black_Markets_Profits":0, "Resources_Smuggled":0, "Average_Profit":0, "Highest_Single_Transaction":0 }, "Trading":{ "Markets_Traded_With":21, "Market_Profits":1081121656, "Resources_Traded":42231, "Average_Profit":10496326.757282, "Highest_Single_Transaction":35028160 }, "Mining":{ "Mining_Profits":0, "Quantity_Mined":0, "Materials_Collected":411 }, "Exploration":{ "Systems_Visited":267, "Exploration_Profits":7011050, "Planets_Scanned_To_Level_2":559, "Planets_Scanned_To_Level_3":559, "Efficient_Scans":0, "Highest_Payout":1013285, "Total_Hyperspace_Distance":5905, "Total_Hyperspace_Jumps":408, "Greatest_Distance_From_Start":406.91005364237, "Time_Played":235440 }, "Passengers":{ "Passengers_Missions_Bulk":0, "Passengers_Missions_VIP":0, "Passengers_Missions_Delivered":0, "Passengers_Missions_Ejected":0 }, "Search_And_Rescue":{ "SearchRescue_Traded":0, "SearchRescue_Profit":0, "SearchRescue_Count":0 }, "Crafting":{ "Count_Of_Used_Engineers":1, "Recipes_Generated":16, "Recipes_Generated_Rank_1":10, "Recipes_Generated_Rank_2":4, "Recipes_Generated_Rank_3":2, "Recipes_Generated_Rank_4":0, "Recipes_Generated_Rank_5":0 }, "Crew":{  }, "Multicrew":{ "Multicrew_Time_Total":0, "Multicrew_Gunner_Time_Total":0, "Multicrew_Fighter_Time_Total":0, "Multicrew_Credits_Total":0, "Multicrew_Fines_Total":0 }, "Material_Trader_Stats":{ "Trades_Completed":0, "Materials_Traded":0 }, "CQC":{ "CQC_Credits_Earned":168, "CQC_Time_Played":450, "CQC_KD":0, "CQC_Kills":0, "CQC_WL":0 } }`))
	// validate
	if err != nil {
		return
	}
	v, ok := obj.(*StatisticsEvent)
	if !ok {
		t.Error("Incorrect Event")
	}
	log.Println(v)
	if v.Event != "Statistics" {
		t.Error("Event not set")
	}
	if v.BankAccount.CurrentWealth != 1104888856 {
		t.Error("Failed to parse")
	}
}
