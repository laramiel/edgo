package journal

type Base struct {
	RAW       string `json:"-"`
	Timestamp string `json:"timestamp"`
	Event     string `json:"event"`
}

type Name_Localized struct {
	Name          string `json:"Name"`
	NameLocalised string `json:"Name_Localised,omitempty"`
}

type Category_Localized struct {
	Category          string `json:"Category"`
	CategoryLocalised string `json:"Category_Localised,omitempty"`
}

type ShipType_Localized struct {
	ShipType          string `json:"ShipType"`
	ShipTypeLocalised string `json:"ShipType_Localised,omitempty"`
}
