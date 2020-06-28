package journal

import (
	"log"
	"testing"
)

func TestParseCargo(t *testing.T) {
	obj, err := ParseStatusContents("cargo.json", []byte(`
{ "timestamp":"2020-06-28T00:16:05Z", "event":"Cargo", "Vessel":"Ship", "Count":64,
"Inventory":[ 
{ "Name":"palladium", "Count":1, "Stolen":0 }, 
{ "Name":"indium", "Count":1, "Stolen":0 }, 
{ "Name":"drones", "Name_Localised":"Limpet", "Count":62, "Stolen":0 }
] }`))
	// validate
	if err != nil {
		t.Error(err)
	}
	v, ok := obj.(*Cargo)
	if !ok {
		t.Error("Not Status")
	}
	log.Println(v)
}

func TestParseMarket(t *testing.T) {
	obj, err := ParseStatusContents("market.json", []byte(`
{ "timestamp":"2020-06-26T06:27:36Z", "event":"Market", "MarketID":3228781056, "StationName":"Hopkins Landing", "StationType":"Coriolis", "StarSystem":"Lung", "Items":[ 
{ "id":128049153, "Name":"$palladium_name;", "Name_Localised":"Palladium", "Category":"$MARKET_category_metals;", "Category_Localised":"Metals", "BuyPrice":12448, "SellPrice":12308, "MeanPrice":13244, "StockBracket":1, "DemandBracket":0, "Stock":1, "Demand":1, "Consumer":false, "Producer":true, "Rare":false }, 
{ "id":128049154, "Name":"$gold_name;", "Name_Localised":"Gold", "Category":"$MARKET_category_metals;", "Category_Localised":"Metals", "BuyPrice":9945, "SellPrice":9831, "MeanPrice":9373, "StockBracket":1, "DemandBracket":0, "Stock":51, "Demand":1, "Consumer":false, "Producer":true, "Rare":false }, 
{ "id":128049155, "Name":"$silver_name;", "Name_Localised":"Silver", "Category":"$MARKET_category_metals;", "Category_Localised":"Metals", "BuyPrice":4785, "SellPrice":4671, "MeanPrice":4759, "StockBracket":2, "DemandBracket":0, "Stock":896, "Demand":1, "Consumer":false, "Producer":true, "Rare":false }, 
{ "id":128049156, "Name":"$bertrandite_name;", "Name_Localised":"Bertrandite", "Category":"$MARKET_category_minerals;", "Category_Localised":"Minerals", "BuyPrice":0, "SellPrice":3227, "MeanPrice":2486, "StockBracket":0, "DemandBracket":3, "Stock":0, "Demand":40008, "Consumer":true, "Producer":false, "Rare":false }
] }`))
	// validate
	if err != nil {
		t.Error(err)
	}
	v, ok := obj.(*Market)
	if !ok {
		t.Error("Not Status")
	}
	log.Println(v)
}

func TestParseModulesInfo(t *testing.T) {
	obj, err := ParseStatusContents("modulesinfo.json", []byte(`
{ "timestamp":"2020-06-26T04:17:41Z", "event":"ModuleInfo", "Modules":[ 
{ "Slot":"MainEngines", "Item":"int_engine_size5_class5", "Power":6.120000, "Priority":0 }, 
{ "Slot":"Slot03_Size3", "Item":"int_shieldgenerator_size3_class5", "Power":2.520000, "Priority":0 }, 
{ "Slot":"MediumHardpoint1", "Item":"hpt_beamlaser_gimbal_medium", "Power":1.000000, "Priority":0 }, 
{ "Slot":"MediumHardpoint2", "Item":"hpt_beamlaser_gimbal_medium", "Power":1.000000, "Priority":0 }, 
{ "Slot":"Slot05_Size3", "Item":"int_buggybay_size2_class2", "Power":0.750000, "Priority":0 }
] }`))
	// validate
	if err != nil {
		t.Error(err)
	}
	v, ok := obj.(*ModulesInfo)
	if !ok {
		t.Error("Not Status")
	}
	log.Println(v)
}

func TestParseNavRoute(t *testing.T) {
	obj, err := ParseStatusContents("navroute.json", []byte(`
{ "timestamp":"2020-06-26T06:29:00Z", "event":"NavRoute", "Route":[ 
{ "StarSystem":"Lung", "SystemAddress":4511603100011, "StarPos":[3.71875,-6.53125,-13.53125], "StarClass":"K" }, 
{ "StarSystem":"LHS 1918", "SystemAddress":9467315955105, "StarPos":[30.56250,-0.50000,-22.43750], "StarClass":"M" }, 
{ "StarSystem":"Puppis Sector HW-W b1-3", "SystemAddress":7268829570457, "StarPos":[56.93750,-1.00000,-29.84375], "StarClass":"M" }, 
{ "StarSystem":"Bakas", "SystemAddress":11667144517017, "StarPos":[84.00000,-2.87500,-32.56250], "StarClass":"K" }, 
{ "StarSystem":"CD-31 4778", "SystemAddress":908687577802, "StarPos":[105.40625,-10.71875,-45.46875], "StarClass":"G" }, 
{ "StarSystem":"Deciat", "SystemAddress":6681123623626, "StarPos":[122.62500,-0.81250,-47.28125], "StarClass":"K" }
] }`))
	// validate
	if err != nil {
		t.Error(err)
	}
	v, ok := obj.(*NavRoute)
	if !ok {
		t.Error("Not Status")
	}
	log.Println(v)
}

func TestParseOutfitting(t *testing.T) {
	obj, err := ParseStatusContents("outfitting.json", []byte(`
{ "timestamp":"2020-06-28T00:31:59Z", "event":"Outfitting", "MarketID":128676487, "StationName":"Farseer Inc", "StarSystem":"Deciat", "Horizons":true, "Items":[ 
{ "id":128064070, "Name":"int_engine_size2_class3", "BuyPrice":17358 }, 
{ "id":128064072, "Name":"int_engine_size2_class5", "BuyPrice":156219 }, 
{ "id":128064075, "Name":"int_engine_size3_class3", "BuyPrice":55025 }, 
{ "id":128064077, "Name":"int_engine_size3_class5", "BuyPrice":495215 }, 
{ "id":128064080, "Name":"int_engine_size4_class3", "BuyPrice":174426 }, 
{ "id":128064081, "Name":"int_engine_size4_class4", "BuyPrice":523276 }, 
{ "id":128064085, "Name":"int_engine_size5_class3", "BuyPrice":552929 }, 
{ "id":128064086, "Name":"int_engine_size5_class4", "BuyPrice":1658786 }, 
{ "id":128064105, "Name":"int_hyperdrive_size2_class3", "BuyPrice":17358 }, 
{ "id":128064107, "Name":"int_hyperdrive_size2_class5", "BuyPrice":156219 }, 
{ "id":128064110, "Name":"int_hyperdrive_size3_class3", "BuyPrice":55025 }, 
{ "id":128064112, "Name":"int_hyperdrive_size3_class5", "BuyPrice":495215 }, 
{ "id":128064115, "Name":"int_hyperdrive_size4_class3", "BuyPrice":174426 }, 
{ "id":128064116, "Name":"int_hyperdrive_size4_class4", "BuyPrice":523276 }, 
{ "id":128064124, "Name":"int_hyperdrive_size6_class2", "BuyPrice":584261 }, 
{ "id":128064126, "Name":"int_hyperdrive_size6_class4", "BuyPrice":5258348 }
] }`))
	// validate
	if err != nil {
		t.Error(err)
	}
	v, ok := obj.(*Outfitting)
	if !ok {
		t.Error("Not Status")
	}
	log.Println(v)
}

func TestParseShipyard(t *testing.T) {
	obj, err := ParseStatusContents("shipyard.json", []byte(`
{ "timestamp":"2020-06-26T01:25:30Z", "event":"Shipyard", "MarketID":128087544, "StationName":"Awyra Flirble", "StarSystem":"Eurybia", "Horizons":true, "AllowCobraMkIV":false, "PriceList":[ 
{ "id":128049291, "ShipType":"dolphin", "ShipPrice":1303890 }, 
{ "id":128672269, "ShipType":"independant_trader", "ShipType_Localised":"Keelback", "ShipPrice":3048001 }, 
{ "id":128672276, "ShipType":"asp_scout", "ShipType_Localised":"Asp Scout", "ShipPrice":3862126 }, 
{ "id":128049303, "ShipType":"asp", "ShipType_Localised":"Asp Explorer", "ShipPrice":6494626 }, 
{ "id":128049321, "ShipType":"federation_dropship", "ShipType_Localised":"Federal Dropship", "ShipPrice":14314205 }, 
{ "id":128049315, "ShipType":"empire_trader", "ShipType_Localised":"Imperial Clipper", "ShipPrice":22295860 }, 
{ "id":128049327, "ShipType":"orca", "ShipPrice":47326390 }, 
{ "id":128049351, "ShipType":"ferdelance", "ShipType_Localised":"Fer-de-Lance", "ShipPrice":50277864 }, 
{ "id":128915979, "ShipType":"mamba", "ShipPrice":54470364 }, 
{ "id":128049339, "ShipType":"python", "ShipPrice":55553725 }, 
{ "id":128049345, "ShipType":"belugaliner", "ShipType_Localised":"Beluga Liner", "ShipPrice":82419445 }
 ] }`))
	// validate
	if err != nil {
		t.Error(err)
	}
	v, ok := obj.(*Shipyard)
	if !ok {
		t.Error("Not Status")
	}
	log.Println(v)
}

func TestParseStatus(t *testing.T) {
	obj, err := ParseStatusContents("status.json", []byte(`
{ "timestamp":"2020-06-28T00:59:57Z", "event":"Status", "Flags":153157645, "Pips":[2,8,2], "FireGroup":1, "GuiFocus":0, "Fuel":{ "FuelMain":32.000000, "FuelReservoir":0.630000 },
"Cargo":64.000000, "LegalState":"Clean", "Latitude":23.862598, "Longitude":152.980728, "Heading":356, "Altitude":0, "BodyName":"Deciat 6 a", "PlanetRadius":787518.625000 }
`))
	// validate
	if err != nil {
		t.Error(err)
	}
	v, ok := obj.(*Status)
	if !ok {
		t.Error("Not Status")
	}
	log.Println(v)
}
