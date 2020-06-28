package journal

// http://hosting.zaonce.net/community/journal/v25/Journal-Manual-v25.pdf

// 2.1
//  C:\Users\User\Saved Games\Frontier Developments\Elite Dangerous\

// 2.2
type FileheaderEvent struct {
	Base
	Part        int64  `json:"part"`
	Language    string `json:"language"`
	Gameversion string `json:"gameversion"`
	Build       string `json:"build"`
}

type Name_Count_Common struct {
	Name_Localized
	Count int64 `json:"Count"`
}

// 3.1
// cargo.json written
type CargoEvent struct {
	Base
	Vessel    string `json:"Vessel"`
	Inventory []struct {
		Name_Count_Common
		Stolen int `json:"Stolen"`
	} `json:"Inventory,omitempty"`
}

// 3.2
type ClearSavedGameEvent struct {
	Base
	Name string `json:"Name"`
	FID  string `json:"FID"`
}

// 3.3
type CommanderEvent struct {
	Base
	FID  string `json:"FID"`
	Name string `json:"Name"`
}

// 3.4
type LoadoutEvent struct {
	Base
	Ship         string  `json:"Ship"`
	ShipId       int64   `json:"ShipID"`
	ShipName     string  `json:"ShipName"`
	ShipIdent    string  `json:"ShipIdent"`
	HullValue    int64   `json:"HullValue,omitempty"`
	ModulesValue int64   `json:"ModulesValue,omitempty"`
	HullHealth   float64 `json:"HullHealth"`
	UnladenMass  float64 `json:"UnladenMass,omitempty"`
	FuelCapacity struct {
		Main    float64 `json:"Main"`
		Reserve float64 `json:"Reserve"`
	} `json:"FuelCapacity"`
	CargoCapacity int             `json:"CargoCapacity"`
	MaxJumpRange  float64         `json:"MaxJumpRange"`
	Rebuy         int64           `json:"Rebuy"`
	Hot           bool            `json:"Hot,omitempty"`
	Modules       []LoadoutModule `json:"Modules,omitempty"`
}

type LoadoutModule struct {
	Slot            string          `json:"Slot"`
	Item            string          `json:"Item"`
	On              bool            `json:"On"`
	Priority        int             `json:"Priority"`
	Health          float64         `json:"Health"`
	Value           int64           `json:"Value"`
	AmmoInClip      int64           `json:"AmmoInClip,omitempty"` // places in cabin
	AmmoInHopper    int64           `json:"AmmoInHopper,omitempty"`
	EngineeringInfo EngineeringInfo `json:"Engineering,omitempty"`
}

type EngineeringInfo struct {
	EngineerId                  int64                        `json:"EngineerID"`
	Engineer                    string                       `json:"Engineer"`
	BlueprintId                 int                          `json:"BlueprintID"`
	BlueprintName               string                       `json:"BlueprintName"`
	Level                       int                          `json:"Level"`
	Quality                     float64                      `json:"Quality"`
	ExperimentalEffect          string                       `json:"ExperimentalEffect,omitempty"`
	ExperimentalEffectLocalised string                       `json:"ExperimentalEffect_Localised,omitempty"`
	Modifiers                   []EngineeringLoadOutModifier `json:"Modifiers,omitempty"`
}

type EngineeringLoadOutModifier struct {
	Label         string  `json:"Label,omitempty"`
	Value         float64 `json:"Value,omitempty"`
	ValueStr      string  `json:"ValueStr,omitempty"`
	OriginalValue float64 `json:"OriginalValue,omitempty"`
	LessIsGood    bool    `json:"LessIsGood,omitempty"`
}

// 3.5
type MaterialsEvent struct {
	Base
	Raw          []Name_Count_Common `json:"Raw,omitempty"`
	Manufactured []Name_Count_Common `json:"Manufactured,omitempty"`
	Encoded      []Name_Count_Common `json:"Encoded,omitempty"`
}

// 3.6
type MissionsEvent struct {
	Base
	Active   []MissionInfo `json:"Active,omitempty"`
	Failed   []MissionInfo `json:"Failed,omitempty"`
	Complete []MissionInfo `json:"Complete,omitempty"`
}

type MissionInfo struct {
	MissionId        int64  `json:"MissionID"`
	Name             string `json:"Name"`
	PassengerMission bool   `json:"PassengerMission"`
	Expires          int64  `json:"Expires"`
}

// 3.7
type NewCommanderEvent struct {
	Base
	Name    string `json:"Name"`
	Fid     string `json:"FID"`
	Package string `json:"Package"`
}

// 3.8
type LoadGameEvent struct {
	Base
	Fid           string  `json:"FID"`
	Commander     string  `json:"Commander"`
	Horizons      bool    `json:"Horizons"`
	Ship          string  `json:"Ship"`
	ShipLocalised string  `json:"Ship_Localised,omitempty"`
	ShipId        int64   `json:"ShipID"`
	ShipName      string  `json:"ShipName,omitempty"`
	ShipIdent     string  `json:"ShipIdent,omitempty"`
	StartLanded   bool    `json:"StartLanded,omitempty"`
	StartDead     bool    `json:"StartDead,omitempty"`
	GameMode      string  `json:"GameMode"` // Open | Solo | Group
	Group         string  `json:"Group,omitempty"`
	Credits       int64   `json:"Credits"`
	Loan          int64   `json:"Loan"`
	FuelLevel     float64 `json:"FuelLevel"`
	FuelCapacity  float64 `json:"FuelCapacity"`
}

// 3.9
type PassengersEvent struct {
	Base
	Manifest []PassengersRecord `json:"Manifest,omitempty"`
}

type PassengersRecord struct {
	MissionId int64  `json:"MissionID"`
	Type      string `json:"Type"`
	Vip       bool   `json:"VIP,omitempty"`
	Wanted    bool   `json:"Wanted,omitempty"`
	Count     int64  `json:"Count"`
}

// 3.10
type PowerplayEvent struct {
	Base
	Power       string `json:"Power"`
	Rank        int    `json:"Rank"`
	Merits      int    `json:"Merits"`
	Votes       int    `json:"Votes"`
	TimePledged int64  `json:"TimePledged"`
}

// 3.11
type ProgressEvent struct {
	Base
	Combat     int `json:"Combat"`     // 0..100
	Trade      int `json:"Trade"`      // 0..100
	Explore    int `json:"Explore"`    // 0..100
	Empire     int `json:"Empire"`     // 0..100
	Federation int `json:"Federation"` // 0..100
	CQC        int `json:"CQC"`        // 0..100
}

// 3.12
type RankEvent struct {
	Base
	Combat      int `json:"Combat"`     // 0..8
	Trade       int `json:"Trade"`      // 0..8
	Exploration int `json:"Explore"`    // 0..8
	Empire      int `json:"Empire"`     // 0..14
	Federation  int `json:"Federation"` // 0..14
	CQC         int `json:"CQC"`        // 0..8
}

// 3.13
type ReputationEvent struct {
	Base
	Empire      float64 `json:"Empire,omitempty"`      // -100 .. 100
	Federation  float64 `json:"Federation,omitempty"`  // -100 .. 100
	Independent float64 `json:"Independent,omitempty"` // -100 .. 100
	Alliance    float64 `json:"Alliance,omitempty"`    // -100 .. 100
}

// 3.14
type StatisticsEvent struct {
	Base
	BankAccount        StatisticsBankAccount        `json:"Bank_Account"`
	Combat             StatisticsCombat             `json:"Combat"`
	Crime              StatisticsCrime              `json:"Crime"`
	Smuggling          StatisticsSmuggling          `json:"Smuggling"`
	Trading            StatisticsTrading            `json:"Trading"`
	Mining             StatisticsMining             `json:"Mining"`
	Exploration        StatisticsExploration        `json:"Exploration"`
	Passengers         StatisticsPassengers         `json:"Passengers"`
	SearchAndRescue    StatisticsSearchAndRescue    `json:"Search_And_Rescue"`
	ThargoidEncounters StatisticsThargoidEncounters `json:"TG_ENCOUNTERS"`
	Crafting           StatisticsCrafting           `json:"Crafting"`
	Crew               StatisticsCrew               `json:"Crew"`
	MultiCrew          StatisticsMultiCrew          `json:"Multicrew"`
	MaterialTrader     StatisticsMaterialTrader     `json:"Material_Trader_Stats"`
	CQC                StatisticsCQC                `json:"CQC"`
}

type StatisticsBankAccount struct {
	CurrentWealth          int64 `json:"Current_Wealth,omitempty"`
	SpentOnShips           int64 `json:"Spent_On_Ships,omitempty"`
	SpentOnOutfitting      int64 `json:"Spent_On_Outfitting,omitempty"`
	SpentOnRepairs         int64 `json:"Spent_On_Repairs,omitempty"`
	SpentOnFuel            int64 `json:"Spent_On_Fuel,omitempty"`
	SpentOnAmmoConsumables int64 `json:"Spent_On_Ammo_Consumables,omitempty"`
	InsuranceClaims        int64 `json:"Insurance_Claims,omitempty"`
	SpentOnInsurance       int64 `json:"Spent_On_Insurance,omitempty"`
	OwnedShipCount         int64 `json:"Owned_Ship_Count,omitempty"`
}

type StatisticsCombat struct {
	BountiesClaimed      int64 `json:"Bounties_Claimed,omitempty"`
	BountyHuntingProfit  int64 `json:"Bounty_Hunting_Profit,omitempty"`
	CombatBonds          int64 `json:"Combat_Bonds,omitempty"`
	CombatBondProfits    int64 `json:"Combat_Bond_Profits,omitempty"`
	Assassinations       int64 `json:"Assassinations,omitempty"`
	AssassinationProfits int64 `json:"Assassination_Profits,omitempty"`
	HighestSingleReward  int64 `json:"Highest_Single_Reward,omitempty"`
	SkimmersKilled       int64 `json:"Skimmers_Killed,omitempty"`
}

type StatisticsCrime struct {
	Notoriety        int64 `json:"Notoriety,omitempty"`
	Fines            int   `json:"Fines,omitempty"`
	TotalFines       int64 `json:"Total_Fines,omitempty"`
	BountiesReceived int   `json:"Bounties_Received,omitempty"`
	TotalBounties    int64 `json:"Total_Bounties,omitempty"`
	HighestBounty    int64 `json:"Highest_Bounty,omitempty"`
}

type StatisticsSmuggling struct {
	BlackMarketsTradedWith   int   `json:"Black_Markets_Traded_With,omitempty"`
	BlackMarketsProfits      int64 `json:"Black_Markets_Profits,omitempty"`
	ResourcesSmuggled        int   `json:"Resources_Smuggled,omitempty"`
	AverageProfit            int64 `json:"Average_Profit,omitempty"`
	HighestSingleTransaction int64 `json:"Highest_Single_Transaction,omitempty"`
}

type StatisticsCrafting struct {
	CountOfUsedEngineers  int `json:"Count_Of_Used_Engineers,omitempty"`
	RecipesGenerated      int `json:"Recipes_Generated,omitempty"`
	RecipesGeneratedRank1 int `json:"Recipes_Generated_Rank_1,omitempty"`
	RecipesGeneratedRank2 int `json:"Recipes_Generated_Rank_2,omitempty"`
	RecipesGeneratedRank3 int `json:"Recipes_Generated_Rank_3,omitempty"`
	RecipesGeneratedRank4 int `json:"Recipes_Generated_Rank_4,omitempty"`
	RecipesGeneratedRank5 int `json:"Recipes_Generated_Rank_5,omitempty"`
}

type StatisticsCrew struct {
	TotalWages int64 `json:"NpcCrew_TotalWages,omitempty"`
	Hired      int   `json:"NpcCrew_Hired,omitempty"`
	Fired      int   `json:"NpcCrew_Fired,omitempty"`
	Died       int   `json:"NpcCrew_Died,omitempty"`
}

type StatisticsExploration struct {
	SystemsVisited            int     `json:"Systems_Visited,omitempty"`
	ExplorationProfits        int64   `json:"Exploration_Profits,omitempty"`
	PlanetsScannedToLevel2    int     `json:"Planets_Scanned_To_Level_2,omitempty"`
	PlanetsScannedToLevel3    int     `json:"Planets_Scanned_To_Level_3,omitempty"`
	EfficientScans            int     `json:"Efficient_Scans,omitempty"`
	HighestPayout             int64   `json:"Highest_Payout,omitempty"`
	TotalHyperspaceDistance   int     `json:"Total_Hyperspace_Distance,omitempty"`
	TotalHyperspaceJumps      int     `json:"Total_Hyperspace_Jumps,omitempty"`
	GreatestDistanceFromStart float64 `json:"Greatest_Distance_From_Start,omitempty"`
	TimePlayed                int64   `json:"Time_Played,omitempty"`
}

type StatisticsMaterialTrader struct {
	TradesCompleted        int `json:"Trades_Completed,omitempty"`
	MaterialsTraded        int `json:"Materials_Traded,omitempty"`
	EncodedMaterialsTraded int `json:"Encoded_Materials_Traded,omitempty"`
	RawMaterialsTraded     int `json:"Raw_Materials_Traded,omitempty"`
	Grade1Traded           int `json:"Grade_1_Materials_Traded,omitempty"`
	Grade2Traded           int `json:"Grade_2_Materials_Traded,omitempty"`
	Grade3Traded           int `json:"Grade_3_Materials_Traded,omitempty"`
	Grade4Traded           int `json:"Grade_4_Materials_Traded,omitempty"`
	Grade5Traded           int `json:"Grade_5_Materials_Traded,omitempty"`
}

type StatisticsMining struct {
	MiningProfits      int64 `json:"Mining_Profits,omitempty"`
	QuantityMined      int   `json:"Quantity_Mined,omitempty"`
	MaterialsCollected int   `json:"Materials_Collected,omitempty"`
}

type StatisticsMultiCrew struct {
	TimeTotal        int   `json:"Multicrew_Time_Total,omitempty"`
	GunnerTimeTotal  int   `json:"Multicrew_Gunner_Time_Total,omitempty"`
	FighterTimeTotal int   `json:"Multicrew_Fighter_Time_Total,omitempty"`
	CreditsTotal     int64 `json:"Multicrew_Credits_Total,omitempty"`
	FinesTotal       int64 `json:"Multicrew_Fines_Total,omitempty"`
}

type StatisticsPassengers struct {
	PassengersMissionsAccepted    int   `json:"Passengers_Missions_Accepted,omitempty"`
	PassengersMissionsDisgruntled int   `json:"Passengers_Missions_Disgruntled,omitempty"`
	PassengersMissionsBulk        int   `json:"Passengers_Missions_Bulk,omitempty"`
	PassengersMissionsVip         int64 `json:"Passengers_Missions_VIP,omitempty"`
	PassengersMissionsDelivered   int64 `json:"Passengers_Missions_Delivered,omitempty"`
	PassengersMissionsEjected     int64 `json:"Passengers_Missions_Ejected,omitempty"`
}

type StatisticsSearchAndRescue struct {
	SearchRescueTraded int64 `json:"SearchRescue_Traded,omitempty"`
	SearchRescueProfit int64 `json:"SearchRescue_Profit,omitempty"`
	SearchRescueCount  int64 `json:"SearchRescue_Count,omitempty"`
}

type StatisticsThargoidEncounters struct {
	Total         int64  `json:"TG_ENCOUNTER_TOTAL,omitempty"`
	LastSystem    string `json:"TG_ENCOUNTER_TOTAL_LAST_SYSTEM,omitempty"`
	LastTimestamp string `json:"TG_ENCOUNTER_TOTAL_LAST_TIMESTAMP,omitempty"`
	LastShip      string `json:"TG_ENCOUNTER_TOTAL_LAST_SHIP,omitempty"`
	ScoutCount    int64  `json:"TG_SCOUT_COUNT,omitempty"`
}

type StatisticsTrading struct {
	MarketsTradedWith        int     `json:"Markets_Traded_With,omitempty"`
	MarketProfits            int64   `json:"Market_Profits,omitempty"`
	ResourcesTraded          int     `json:"Resources_Traded,omitempty"`
	AverageProfit            float64 `json:"Average_Profit,omitempty"`
	HighestSingleTransaction int64   `json:"Highest_Single_Transaction,omitempty"`
}

type StatisticsCQC struct {
	CreditsEarned int64 `json:"CQC_Credits_Earned,omitempty"`
	TimePlayed    int64 `json:"CQC_Time_Played,omitempty"`
	KD            int64 `json:"CQC_KD,omitempty"`
	Kills         int64 `json:"CQC_Kills,omitempty"`
	WL            int64 `json:"CQC_WL,omitempty"`
}

// 4
type Base_System struct {
	StarSystem    string `json:"StarSystem"`
	SystemAddress int64  `json:"SystemAddress"`
	Body          string `json:"Body"`
	BodyId        int64  `json:"BodyID,omitempty"`
}

type Base_Dock struct {
	StationName string `json:"StationName"`
	StationType string `json:"StationType"`
	MarketId    int64  `json:"MarketID"`
}

// 4.1
type ApproachBodyEvent struct {
	Base
	Base_System
}

// 4.2
type DockedEvent struct {
	Base
	Base_Dock
	SystemAddress  int64  `json:"SystemAddress"`
	StarSystem     string `json:"StarSystem"`
	CockpitBreach  bool   `json:"CockpitBreach,omitempty"`
	Wanted         bool   `json:"Wanted,omitempty"`
	ActiveFine     bool   `json:"ActiveFine,omitempty"`
	StationFaction struct {
		Name  string `json:"Name"`
		State string `json:"State"`
	} `json:"StationFaction,omitempty"`
	StationAllegiance       string `json:"StationAllegiance"`
	StationEconomy          string `json:"StationEconomy"`
	StationEconomyLocalised string `json:"StationEconomy_Localised,omitempty"`
	StationEconomies        []struct {
		Name_Localized
		Proportion float64 `json:"Proportion"`
	} `json:"StationEconomies,omitempty"`
	StationGovernment          string   `json:"StationGovernment"`
	StationGovernmentLocalised string   `json:"StationGovernment_Localised,omitempty"`
	StationServices            []string `json:"StationServices,omitempty"`
	DistFromStarLS             float64  `json:"DistFromStarLS"`
}

// 4.3
type DockingCancelledEvent struct {
	Base
	Base_Dock
}

// 4.4
type DockingDeniedEvent struct {
	Base
	Base_Dock
	Reason string `json:"StationName"`
}

// 4.5
type DockingGrantedEvent struct {
	Base
	Base_Dock
	LandingPad int `json:"LandingPad"`
}

// 4.6
type DockingRequestedEvent struct {
	Base
	Base_Dock
}

type DockingTimeoutEvent struct {
	Base
	Base_Dock
}

// 4.7
type FSDJumpEvent struct {
	Base
	Base_System
	StarPos       [3]float64 `json:"StarPos"`
	JumpDist      float64    `json:"JumpDist"`
	FuelUsed      float64    `json:"FuelUsed"`
	FuelLevel     float64    `json:"FuelLevel"`
	BoostUsed     bool       `json:"BoostUsed,omitempty"`
	SystemFaction struct {
		Name         string `json:"Name"`
		FactionState string `json:"FactionState"`
	} `json:"SystemFaction,omitempty"`
	SystemAllegiance             string           `json:"SystemAllegiance"`
	SystemEconomy                string           `json:"SystemEconomy"`
	SystemEconomyLocalised       string           `json:"SystemEconomy_Localised,omitempty"`
	SystemSecondEconomy          string           `json:"SystemSecondEconomy,omitempty"`
	SystemSecondEconomyLocalised string           `json:"SystemSecondEconomy_Localised,omitempty"`
	SystemGovernment             string           `json:"SystemGovernment"`
	SystemGovernmentLocalised    string           `json:"SystemGovernment_Localised,omitempty"`
	SystemSecurity               string           `json:"SystemSecurity"`
	SystemSecurityLocalised      string           `json:"SystemSecurity_Localised,omitempty"`
	Population                   int64            `json:"Population"`
	Wanted                       bool             `json:"Wanted,omitempty"`
	Factions                     []FSDJumpFaction `json:"Factions,omitempty"`
	Conflicts                    []FSDConflict    `json:"Conflicts,omitempty"`
	Powers                       []string         `json:"Powers,omitempty"`
	PowerplayState               string           `json:"PowerplayState"`
	FactionState                 string           `json:"FactionState"`
}

type FSDConflict struct {
	WarType  string             `json:"WarType"`
	Status   string             `json:"Status"`
	Faction1 FSDConflictFaction `json:"Faction1"`
	Faction2 FSDConflictFaction `json:"Faction2"`
}

type FSDConflictFaction struct {
	Name    string `json:"Name"`
	Stake   string `json:"Stake"`
	WonDays int    `json:"WonDays"`
}

type FSDJumpFaction struct {
	Name                string                `json:"Name"`
	FactionState        string                `json:"FactionState"`
	Government          string                `json:"Government"`
	Influence           float64               `json:"Influence"`
	Happiness           string                `json:"Happiness"`
	HappinessLocalised  string                `json:"Happiness_Localised,omitempty"`
	CommanderReputation float64               `json:"MyReputation,omitempty"`
	RecoveringStates    []FSDJumpFactionState `json:"RecoveringStates,omitempty"`
	PendingStates       []FSDJumpFactionState `json:"PendingStates,omitempty"`
	ActiveStates        []struct {
		State string `json:"State"`
	} `json:"ActiveStates,omitempty"`
	SquadronFaction bool
}

type FSDJumpFactionState struct {
	State string `json:"State,omitempty"`
	Trend int64  `json:"Trend,omitempty"`
}

// 4.9
type FSDTargetEvent struct {
	Base
	Name                  string `json:"Name,omitempty"`
	SystemAddress         int64  `json:"SystemAddress,omitempty"`
	RemainingJumpsInRoute int    `json:"RemainingJumpsInRoute,omitempty"`
}

// 4.10
type LeaveBodyEvent struct {
	Base
	Base_System
}

// 4.11
type LiftoffEvent struct {
	Base
	PlayerControlled bool    `json:"PlayerControlled"`
	Latitude         float64 `json:"Latitude,omitempty"`
	int64itude       float64 `json:"int64itude,omitempty"`
}

// 4.12
type LocationEvent struct {
	Base
	Base_System
	StarPos        [3]float64 `json:"StarPos"`
	BodyType       string     `json:"BodyType"`
	DistFromStarLS int64      `json:"DistFromStarLS"`
	SystemFaction  struct {
		Name         string `json:"Name"`
		FactionState string `json:"FactionState"`
	} `json:"SystemFaction,omitempty"`
	SystemAllegiance             string           `json:"SystemAllegiance"`
	SystemEconomy                string           `json:"SystemEconomy"`
	SystemEconomyLocalised       string           `json:"SystemEconomy_Localised,omitempty"`
	SystemSecondEconomy          string           `json:"SystemSecondEconomy,omitempty"`
	SystemSecondEconomyLocalised string           `json:"SystemSecondEconomy_Localised,omitempty"`
	SystemGovernment             string           `json:"SystemGovernment"`
	SystemGovernmentLocalised    string           `json:"SystemGovernment_Localised,omitempty"`
	SystemSecurity               string           `json:"SystemSecurity"`
	SystemSecurityLocalised      string           `json:"SystemSecurity_Localised,omitempty"`
	Population                   int64            `json:"Population"`
	Wanted                       bool             `json:"Wanted,omitempty"`
	Factions                     []FSDJumpFaction `json:"Factions,omitempty"`
	Conflicts                    []FSDConflict    `json:"Conflicts,omitempty"`
	Powers                       []string         `json:"Powers,omitempty"`
	PowerplayState               string           `json:"PowerplayState"`
	FactionState                 string           `json:"FactionState"`
	Docked                       bool             `json:"Docked"`
	// if landed
	Latitude  float64 `json:"Latitude,omitempty"`
	Longitude float64 `json:"Longitude,omitempty"`
	// if docked
	StationName string `json:"StationName,omitempty"`
	StationType string `json:"StationType,omitempty"`
	MarketID    string `json:"MarketID,omitempty"`
	// if starting docked
	StationFaction struct {
		Name         string `json:"Name"`
		FactionState string `json:"FactionState"`
	} `json:"SystemFaction,omitempty"`
	StationAllegiance string `json:"StationAllegiance,omitempty"`
	StationEconomies  []struct {
		Name_Localized
		Proportion float64 `json:"Proportion"`
	} `json:"StationEconomies,omitempty"`
	StationGovernment          string   `json:"StationGovernment,omitempty"`
	StationGovernmentLocalised string   `json:"StationGovernment_Localised,omitempty"`
	StationServices            []string `json:"StationServices,omitempty"`
}

// 4.13
type StartJumpEvent struct {
	Base
	JumpType      string `json:"JumpType"` // Hyperspace | Supercruise
	StarSystem    string `json:"StarSystem,omitempty"`
	SystemAddress int64  `json:"SystemAddress"`
	StarClass     string `json:"StarClass,omitempty"`
}

// 4.14
type SupercruiseEntryEvent struct {
	Base
	StarSystem    string `json:"StarSystem"`
	SystemAddress int64  `json:"SystemAddress,omitempty"`
}

// 4.15
type SupercruiseExitEvent struct {
	Base
	StarSystem    string `json:"StarSystem"`
	Body          string `json:"Body"`
	BodyId        int64  `json:"BodyID,omitempty"`
	BodyType      string `json:"BodyType,omitempty"`
	SystemAddress int64  `json:"SystemAddress,omitempty"`
}

// 4.16
type TouchdownEvent struct {
	Base
	PlayerControlled bool    `json:"PlayerControlled"`
	Latitude         float64 `json:"Latitude,omitempty"`
	Longitude        float64 `json:"Longitude,omitempty"`
}

// 4.17
type UndockedEvent struct {
	Base
	StationName string `json:"StationName"`
	StationType string `json:"StationType,omitempty"`
	MarketId    int64  `json:"MarketID,omitempty"`
}

// 5.1
type BountyEvent struct {
	Base
	Target           string `json:"Target"`
	VictimFaction    string `json:"VictimFaction"`
	SharedWithOthers int64  `json:"SharedWithOthers,omitempty"`
	TotalReward      int64  `json:"TotalReward,omitempty"`
	Rewards          []struct {
		Faction string `json:"Faction,omitempty"`
		Reward  int64  `json:"Reward,omitempty"`
	} `json:"Rewards,omitempty"`
	// Skimmer
	Reward int64 `json:"Reward,omitempty"`
}

// 5.2
type CapShipBond struct {
	Base
	Reward          int64  `json:"Reward,omitempty"`
	AwardingFaction string `json:"AwardingFaction"`
	VictimFaction   string `json:"VictimFaction"`
}

// 5.3, 5.4
type DiedEvent struct {
	Base
	KillerName          string `json:"KillerName,omitempty"`
	KillerNameLocalised string `json:"KillerName_Localised,omitempty"`
	KillerShip          string `json:"KillerShip,omitempty"`
	KillerRank          string `json:"KillerRank,omitempty"`
	// killed by a wing
	Killers []struct {
		Name_Localized
		Ship string `json:"Ship"`
		Rank string `json:"Rank"`
	} `json:"Killers,omitempty"`
}

// 5.5
type EscapeInterdictionEvent struct {
	Base
	Interdictor          string `json:"Interdictor"`
	InterdictorLocalised string `json:"Interdictor_Localised,omitempty"`
	IsPlayer             bool   `json:"IsPlayer"`
}

// 5.6
type FactionKillBondEvent struct {
	Base
	Reward                   int64  `json:"Reward"`
	AwardingFaction          string `json:"AwardingFaction"`
	AwardingFactionLocalised string `json:"AwardingFaction_Localised,omitempty"`
	VictimFaction            string `json:"VictimFaction"`
	VictimFactionLocalised   string `json:"VictimFaction_Localised,omitempty"`
}

// 5.7
type FighterDestroyedEvent struct {
	Base
}

// 5.8
type HeatDamageEvent struct {
	Base
	ID string `json:"ID,omitempty"`
}

// 5.9
type HeatWarningEvent struct {
	Base
}

// 5.10
type HullDamageEvent struct {
	Base
	Health      float64 `json:"Health"`
	PlayerPilot bool    `json:"PlayerPilot"`
	Fighter     bool    `json:"Fighter,omitempty"`
}

// 5.11
type InterdictedEvent struct {
	Base
	Submitted            bool   `json:"Submitted"`
	Interdictor          string `json:"Interdictor"`
	InterdictorLocalised string `json:"Interdictor_Localised,omitempty"`
	IsPlayer             bool   `json:"IsPlayer"`
	CombatRank           int    `json:"CombatRank,omitempty"` // player
	Faction              string `json:"Faction,omitempty"`    // NPC
	Power                string `json:"Power,omitempty"`      // NPC
}

// 5.12
type InterdictionEvent struct {
	Base
	Success     bool   `json:"Success"`
	IsPlayer    bool   `json:"IsPlayer"`
	Interdicted string `json:"Interdicted"`
	CombatRank  int    `json:"CombatRank,omitempty"` // player
	Faction     string `json:"Faction,omitempty"`    // NPC
	Power       string `json:"Power,omitempty"`      // NPC
}

// 5.13
type PVPKillEvent struct {
	Base
	Victim     string `json:"Victim"`
	CombatRank int    `json:"CombatRank"` // 0..8
}

// 5.14
type ShieldStateEvent struct {
	Base
	ShieldsUp bool `json:"ShieldsUp"`
}

// 5.15
type ShipTargetedEvent struct {
	Base
	TargetLocked       bool    `json:"TargetLocked"`
	Ship               string  `json:"Ship,omitempty"`
	ShipLocalised      string  `json:"Ship_Localised,omitempty"`
	ScanStage          int64   `json:"ScanStage,omitempty"`
	PilotName          string  `json:"PilotName,omitempty"`
	PilotNameLocalised string  `json:"PilotName_Localised,omitempty"`
	PilotRank          string  `json:"PilotRank,omitempty"`
	ShieldHealth       float64 `json:"ShieldHealth,omitempty"`
	HullHealth         float64 `json:"HullHealth,omitempty"`
	Faction            string  `json:"Faction,omitempty"`
	LegalStatus        string  `json:"LegalStatus,omitempty"`
	Bounty             int64   `json:"Bounty,omitempty"`
	Subsystem          string  `json:"Subsystem,omitempty"`
	SubsystemLocalised string  `json:"Subsystem_Localised,omitempty"`
	SubsystemHealth    float64 `json:"SubsystemHealth,omitempty"`
	SquadronId         int     `json:"SquadronID,omitempty"`
}

// 5.16
type SRVDestroyedEvent struct {
	Base
}

// 5.17
type UnderAttackEvent struct {
	Base
	Target string
}

// 6.1
type CodexEntryEvent struct {
	Base
	Name_Localized
	EntryId              int64    `json:"EntryID"`
	SubCategory          string   `json:"SubCategory"`
	SubCategoryLocalised string   `json:"SubCategory_Localised,omitempty"`
	Category             string   `json:"Category"`
	CategoryLocalised    string   `json:"Category_Localised,omitempty"`
	Region               string   `json:"Region"`
	RegionLocalised      string   `json:"Region_Localised,omitempty"`
	System               string   `json:"System"`
	SystemAddress        int64    `json:"SystemAddress"`
	IsNewEntry           bool     `json:"IsNewEntry"`
	VoucherAmount        int64    `json:"VoucherAmount,omitempty"`
	NewTraitsDiscovered  bool     `json:"NewTraitsDiscovered"`
	Traits               []string `json:"NewTraitsDiscovered,omitempty"`
}

// 6.2
type DiscoveryScanEvent struct {
	Base
	SystemAddress int64 `json:"SystemAddress"`
	Bodies        int64 `json:"Bodies"`
}

// 6.3
type ScanEvent struct {
	Base
	ScanType              string  `json:"ScanType,omitempty"`
	BodyName              string  `json:"BodyName"`
	BodyId                int64   `json:"BodyID"`
	StarSystem            string  `json:"StarSystem,omitempty"`
	SystemAddress         int64   `json:"SystemAddress,omitempty"`
	DistanceFromArrivalLs float64 `json:"DistanceFromArrivalLS,omitempty"`
	SurfaceTemperature    float64 `json:"SurfaceTemperature,omitempty"`
	Rings                 []Ring  `json:"Rings,omitempty"`
	// Star
	StarType          string  `json:"StarType,omitempty"`
	Subclass          int     `json:"Subclass,omitempty"` // 0..9
	StellarMass       float64 `json:"StellarMass,omitempty"`
	Radius            float64 `json:"Radius,omitempty,omitempty"`
	AbsoluteMagnitude float64 `json:"AbsoluteMagnitude,omitempty"`
	Luminosity        float64 `json:"Luminosity,omitempty"`
	AgeMY             float64 `json:"AgeMY,omitempty"`
	WasDiscovered     bool    `json:"WasDiscovered,omitempty"`
	WasMapped         bool    `json:"WasMapped,omitempty"`
	// Planet, moon
	Parents []struct {
		BodyType string `json:"BodyType"`
		BodyId   string `json:"BodyID"`
	} `json:"Parents,omitempty"`
	TidalLock             bool             `json:"TidalLock,omitempty"`
	TerraformState        string           `json:"TerraformState,omitempty"`
	PlanetClass           string           `json:"PlanetClass,omitempty"`
	Atmosphere            string           `json:"Atmosphere,omitempty"`
	AtmosphereType        string           `json:"AtmosphereType,omitempty"`
	AtmosphereComposition []string         `json:"AtmosphereComposition,omitempty"`
	Volcanism             string           `json:"Volcanism,omitempty"`
	MassEm                float64          `json:"MassEM,omitempty"`
	SurfaceGravity        float64          `json:"SurfaceGravity,omitempty"`
	SurfacePressure       float64          `json:"SurfacePressure,omitempty"`
	Landable              bool             `json:"Landable,omitempty"`
	Materials             []PlanetMaterial `json:"Materials,omitempty"`
	Composition           Composition      `json:"Composition,omitempty"`
	ReserveLevel          string           `json:"ReserveLevel,omitempty"`
	// rotation parameters
	RotationPeriod float64 `json:"RotationPeriod,omitempty"`
	AxialTilt      float64 `json:"AxialTilt,omitempty"`
	// orbital parameters
	SemiMajorAxis      float64 `json:"SemiMajorAxis,omitempty"`
	Eccentricity       float64 `json:"Eccentricity,omitempty"`
	OrbitalInclination float64 `json:"OrbitalInclination,omitempty"`
	Periapsis          float64 `json:"Periapsis,omitempty"`
	OrbitalPeriod      float64 `json:"OrbitalPeriod,omitempty"`
}

type PlanetMaterial struct {
	Name_Localized
	Percent float64 `json:"Percent,omitempty"`
}

type Ring struct {
	Name_Localized
	RingClass string  `json:"RingClass"`
	MassMt    float64 `json:"MassMT"`
	InnerRad  float64 `json:"InnerRad"`
	OuterRad  float64 `json:"OuterRad"`
}

type Composition struct {
	Ice   float64 `json:"Ice,omitempty"`
	Rock  float64 `json:"Rock,omitempty"`
	Metal float64 `json:"Metal,omitempty"`
}

// 6.4
type FSSAllBodiesFoundEvent struct {
	Base
	SystemName    string `json:"SystemName"`
	SystemAddress int64  `json:"SystemAddress"`
	Count         int64  `json:"Count"`
}

// 6.5
type FSSDiscoveryScanEvent struct {
	Base
	Progress     float64 `json:"Progress"`
	BodyCount    int64   `json:"BodyCount"`
	NonBodyCount int64   `json:"NonBodyCount"`
}

// 6.6
type FSSSignalDiscoveredEvent struct {
	Base
	SystemAddress          int64   `json:"SystemAddress"`
	SignalName             string  `json:"SignalName"`
	SignalNameLocalised    string  `json:"SignalName_Localised,omitempty"`
	UssType                string  `json:"USSType,omitempty"`
	UssTypeLocalised       string  `json:"USSType_Localised,omitempty"`
	SpawningState          string  `json:"SpawningState,omitempty"`
	SpawningStateLocalised string  `json:"SpawningState_Localised,omitempty"`
	SpawningFaction        string  `json:"SpawningFaction,omitempty"`
	ThreatLevel            int64   `json:"ThreatLevel,omitempty"`
	TimeRemaining          float64 `json:"TimeRemaining,omitempty"`
}

// 6.7
type MaterialCollectedEvent struct {
	Base
	Name_Count_Common
	Category string `json:"Category"`
}

// 6.8
type MaterialDiscardedEvent struct {
	Base
	Name_Count_Common
	Category string `json:"Category"`
}

// 6.9
type MaterialDiscoveredEvent struct {
	Base
	Name_Localized
	Category        string `json:"Category"`
	DiscoveryNumber int64  `json:"DiscoveryNumber"`
}

// 6.10
type MultiSellExplorationDataEvent struct {
	Base
	Discovered []struct {
		SystemName string `json:"SystemName"`
		NumBodies  int64  `json:"NumBodies"`
	} `json:"Discovered,omitempty"`
	BaseValue     int64 `json:"BaseValue"`
	Bonus         int64 `json:"Bonus"`
	TotalEarnings int64 `json:"TotalEarnings"`
}

// 6.11
type NavBeaconScanEvent struct {
	Base
	SystemAddress int64 `json:"SystemAddress"`
	NumBodies     int64 `json:"NumBodies"`
}

// 6.12
type BuyExplorationDataEvent struct {
	Base
	System string `json:"System"`
	Cost   int64  `json:"Cost"`
}

// 6.13
type SAAScanCompleteEvent struct {
	Base
	BodyName         string `json:"BodyName"`
	BodyId           int64  `json:"BodyID"`
	ProbesUsed       int    `json:"ProbesUsed"`
	EfficiencyTarget int    `json:"EfficiencyTarget"`
	SystemAddress    int64  `json:"SystemAddress,omitempty"`
}

// 6.14
type SellExplorationDataEvent struct {
	Base
	Systems       []string `json:"Systems"`
	Discovered    []string `json:"Discovered"`
	BaseValue     int64    `json:"BaseValue"`
	Bonus         int64    `json:"Bonus"`
	TotalEarnings int64    `json:"TotalEarnings"`
}

// 6.15
type ScreenshotEvent struct {
	Base
	Filename   string  `json:"Filename"`
	Width      int64   `json:"Width"`
	Height     int64   `json:"Height"`
	System     string  `json:"System"`
	Body       string  `json:"Body"`
	Latitude   float64 `json:"Latitude"`
	int64itude float64 `json:"int64itude"`
	Heading    int64   `json:"Heading"`
	Altitude   float64 `json:"Altitude"`
}

// 7 Trade
// 7.1
type AsteroidCrackedEvent struct {
	Base
	Body string `json:"Body"`
}

// 7.2
type BuyTradeDataEvent struct {
	Base
	System string `json:"System"`
	Cost   int64  `json:"Cost"`
}

// 7.3
type CollectCargoEvent struct {
	Base
	Type          string `json:"Type"`
	TypeLocalised string `json:"Type_Localised,omitempty"`
	Stolen        bool   `json:"Stolen"`
	MissionId     int64  `json:"MissionID,omitempty"`
}

// 7.4
type EjectCargoEvent struct {
	Base
	Type            string `json:"Type"`
	TypeLocalised   string `json:"Type_Localised,omitempty"`
	Count           int64  `json:"Count"`
	Abandoned       bool   `json:"Abandoned,omitempty"`
	MissionId       int64  `json:"MissionID,omitempty"`
	PowerplayOrigin string `json:"PowerplayOrigin,omitempty"`
}

// 7.5
type MarketBuyEvent struct {
	Base
	MarketId      int64  `json:"MarketID"`
	Type          string `json:"Type"`
	TypeLocalised string `json:"Type_Localised,omitempty"`
	Count         int64  `json:"Count"`
	BuyPrice      int64  `json:"BuyPrice"`
	TotalCost     int64  `json:"TotalCost"`
}

// 7.6
type MarketSellEvent struct {
	Base
	MarketId      int64  `json:"MarketID"`
	Type          string `json:"Type"`
	TypeLocalised string `json:"Type_Localised,omitempty"`
	Count         int64  `json:"Count"`
	SellPrice     int64  `json:"SellPrice"`
	TotalSale     int64  `json:"TotalSale"`
	AvgPricePaid  int64  `json:"AvgPricePaid"`
	IllegalGoods  bool   `json:"IllegalGoods,omitempty"`
	StolenGoods   bool   `json:"StolenGoods,omitempty"`
	BlackMarket   bool   `json:"BlackMarket,omitempty"`
}

// 7.7
type MiningRefinedEvent struct {
	Base
	Type          string `json:"Type"`
	TypeLocalised string `json:"Type_Localised,omitempty"`
}

// 8 Station Services

// 8.1
type BuyAmmoEvent struct {
	Base
	Cost int64 `json:"Cost"`
}

// 8.2
type BuyDronesEvent struct {
	Base
	Type      string `json:"Type"`
	Count     int64  `json:"Count"`
	BuyPrice  int64  `json:"BuyPrice"`
	TotalCost int64  `json:"TotalCost"`
}

// 8.3
type CargoDepotEvent struct {
	Base
	MissionId           int64   `json:"MissionID"`
	UpdateType          string  `json:"UpdateType"`
	CargoType           string  `json:"CargoType"`
	CargoTypeLocalised  string  `json:"CargoType_Localised,omitempty"`
	Count               int64   `json:"Count"`
	StartMarketId       int64   `json:"StartMarketID"`
	EndMarketId         int64   `json:"EndMarketID"`
	ItemsCollected      int64   `json:"ItemsCollected"`
	ItemsDelivered      int64   `json:"ItemsDelivered"`
	TotalItemsToDeliver int64   `json:"TotalItemsToDeliver"`
	Progress            float64 `json:"Progress"`
}

// 8.4
type CommunityGoalEvent struct {
	Base
	CurrentGoals []CurrentGoal `json:"CurrentGoals,omitempty"`
}

type CurrentGoal struct {
	CGID                 int64  `json:"CGID"`
	Title                string `json:"Title"`
	SystemName           string `json:"SystemName"`
	MarketName           string `json:"MarketName"`
	Expiry               string `json:"Expiry"`
	IsComplete           bool   `json:"IsComplete"`
	CurrentTotal         int64  `json:"CurrentTotal"`
	PlayerContribution   int64  `json:"PlayerContribution"`
	NumContributors      int64  `json:"NumContributors"`
	PlayerPercentileBand int64  `json:"PlayerPercentileBand"`
	TopTier              struct {
		Name  string `json:"Name"`
		Bonus string `json:"Bonus"`
	} `json:"TopTier"`
	TopRankSize     int64  `json:"TopRankSize,omitempty"`
	PlayerInTopRank bool   `json:"PlayerInTopRank,omitempty"`
	TierReached     string `json:"TierReached,omitempty"`
	Bonus           int64  `json:"Bonus,omitempty"`
}

// 8.5
type CommunityGoalDiscardEvent struct {
	Base
	CGID   int64  `json:"CGID"`
	Name   string `json:"Name"`
	System string `json:"System"`
}

// 8.6
type CommunityGoalJoinEvent struct {
	Base
	CGID   int64  `json:"CGID"`
	Name   string `json:"Name"`
	System string `json:"System"`
}

// 8.7
type CommunityGoalRewardEvent struct {
	Base
	CGID   int64  `json:"CGID"`
	Name   string `json:"Name"`
	System string `json:"System"`
	Reward int64  `json:"Reward"`
}

// 8.8
type CrewAssignEvent struct {
	Base
	Name   string `json:"Name"`
	CrewId int64  `json:"CrewID"`
	Role   string `json:"Role"`
}

// 8.9
type CrewFireEvent struct {
	Base
	Name   string `json:"Name"`
	CrewId int64  `json:"CrewID"`
}

// 8.10
type CrewHireEvent struct {
	Base
	Name       string `json:"Name"`
	CrewId     int64  `json:"CrewID"`
	Faction    string `json:"Faction"`
	Cost       int64  `json:"Cost"`
	CombatRank int    `json:"CombatRank"`
}

// 8.11 obsolete
// 8.12
type EngineerContributionEvent struct {
	Base
	Engineer      string `json:"Engineer"`
	EngineerId    int64  `json:"EngineerID"`
	Type          string `json:"Type,omitempty"`
	Commodity     string `json:"Commodity,omitempty"`
	Material      string `json:"Material,omitempty"`
	Faction       string `json:"Faction,omitempty"`
	Quantity      int64  `json:"Quantity,omitempty"`
	TotalQuantity int64  `json:"TotalQuantity,omitempty"`
}

// 8.13
type EngineerCraftEvent struct {
	Base
	Engineer                    string              `json:"Engineer"`
	EngineerId                  int64               `json:"EngineerID"`
	BlueprintId                 int64               `json:"BlueprintID"`
	BlueprintName               string              `json:"BlueprintName"`
	Level                       int64               `json:"Level"`
	Quality                     float64             `json:"Quality"`
	ApplyExperimentalEffect     string              `json:"ApplyExperimentalEffect,omitempty"`
	Slot                        string              `json:"Slot,omitempty"`
	Module                      string              `json:"Module,omitempty"`
	ExperimentalEffect          string              `json:"ExperimentalEffect,omitempty"`
	ExperimentalEffectLocalised string              `json:"ExperimentalEffect_Localised,omitempty"`
	Ingredients                 []Name_Count_Common `json:"Ingredients"`
	Modifiers                   []Modifier          `json:"Modifiers"`
}

type Modifier struct {
	Label         string  `json:"Label"`
	Value         float64 `json:"Value,omitempty"`
	ValueStr      string  `json:"ValueStr,omitempty"`
	OriginalValue float64 `json:"OriginalValue,omitempty"`
	LessIsGood    int     `json:"LessIsGood,omitempty"`
}

// 8.14 obsolete
// 8.15
type EngineerProgressEvent struct {
	Base
	Engineers []Engineer `json:"Engineers,omitempty"`
}

type Engineer struct {
	Engineer     string `json:"Engineer"`
	EngineerId   int64  `json:"EngineerID"`
	Progress     string `json:"Progress,omitempty"`
	RankProgress int    `json:"RankProgress,omitempty"`
	Rank         int    `json:"Rank,omitempty"`
}

// 8.16
type FetchRemoteModuleEvent struct {
	Base
	StorageSlot         int64  `json:"StorageSlot"`
	StoredItem          string `json:"StoredItem"`
	StoredItemLocalised string `json:"StoredItem_Localised,omitempty"`
	ServerId            int64  `json:"ServerId"`
	TransferCost        int64  `json:"TransferCost"`
	TransferTime        int64  `json:"TransferTime"`
	Ship                string `json:"Ship"`
	ShipId              int64  `json:"ShipID"`
}

// 8.17
// market.json contains the full market price data.
type MarketEvent struct {
	Base
	MarketId    int64  `json:"MarketID"`
	StationName string `json:"StationName"`
	StarSystem  string `json:"StarSystem"`
}

// 8.18
type MassModuleStoreEvent struct {
	Base
	MarketId int64  `json:"MarketID"`
	Ship     string `json:"Ship"`
	ShipId   int64  `json:"ShipID"`
	Items    []Item `json:"Items"`
}

type Item struct {
	Name_Localized
	Slot                  string  `json:"Slot"`
	Hot                   bool    `json:"Hot"`
	EngineerModifications string  `json:"EngineerModifications,omitempty"`
	Level                 int64   `json:"Level"`
	Quality               float64 `json:"Quality"`
}

// 8.19
type MaterialTradeEvent struct {
	Base
	MarketId   int64  `json:"MarketID"`
	TraderType string `json:"TraderType"`
	Paid       Paid   `json:"Paid"`
	Received   Paid   `json:"Received"`
}

type Paid struct {
	Material          string `json:"Material"`
	MaterialLocalised string `json:"Material_Localised,omitempty"`
	Category          string `json:"Category"`
	CategoryLocalised string `json:"Category_Localised,omitempty"`
	Quantity          int64  `json:"Quantity"`
}

// 8.20
type MissionAbandonedEvent struct {
	Base
	Name      string `json:"Name"`
	MissionId int64  `json:"MissionID"`
}

// 8.21
type MissionAcceptedEvent struct {
	Base
	Name_Localized
	Faction               string `json:"Faction"`
	MissionId             int64  `json:"MissionID"`
	Influence             string `json:"Influence"`
	Reputation            string `json:"Reputation"`
	Reward                int64  `json:"Reward"`
	Wing                  bool   `json:"Wing"`
	Commodity             string `json:"Commodity,omitempty"`
	CommodityLocalised    string `json:"Commodity_Localised,omitempty"`
	Count                 int64  `json:"Count,omitempty"`
	Donation              string `json:"Donation,omitempty"`
	Donated               int64  `json:"Donated,omitempty"`
	Target                string `json:"Target,omitempty"`
	TargetType            string `json:"TargetType,omitempty"`
	TargetTypeLocalised   string `json:"TargetType_Localised,omitempty"`
	TargetFaction         string `json:"TargetFaction,omitempty"`
	KillCount             int64  `json:"KillCount,omitempty"`
	Expiry                string `json:"Expiry,omitempty"`
	DestinationSystem     string `json:"DestinationSystem,omitempty"`
	DestinationStation    string `json:"DestinationSystem,omitempty"`
	NewDestinationSystem  string `json:"DestinationSystem,omitempty"`
	NewDestinationStation string `json:"DestinationSystem,omitempty"`
	PassengerCount        int64  `json:"PassengerCount,omitempty"`
	PassengerViPs         bool   `json:"PassengerVIPs,omitempty"`
	PassengerWanted       bool   `json:"PassengerWanted,omitempty"`
	PassengerType         string `json:"PassengerType,omitempty"`
}

// 8.22
type MissionCompletedEvent struct {
	Base
	Name_Localized
	Faction             string            `json:"Faction"`
	MissionId           int64             `json:"MissionID"`
	Commodity           string            `json:"Commodity,omitempty"`
	Count               int64             `json:"Count,omitempty"`
	Target              string            `json:"Target,omitempty"`
	TargetType          string            `json:"TargetType,omitempty"`
	TargetTypeLocalised string            `json:"TargetType_Localised,omitempty"`
	TargetFaction       string            `json:"TargetFaction,omitempty"`
	Reward              int64             `json:"Reward,omitempty"`
	Donation            string            `json:"Donation,omitempty"`
	Donated             int64             `json:"Donated,omitempty"`
	PermitsAwarded      []string          `json:"PermitsAwarded,omitempty"`
	CommodityReward     []MaterialsReward `json:"CommodityReward,omitempty"`
	MaterialsReward     []MaterialsReward `json:"MaterialsReward,omitempty"`
	FactionEffects      []FactionEffect   `json:"FactionEffects,omitempty"`
}

type FactionEffect struct {
	Faction         string      `json:"Faction"`
	Effects         []Effect    `json:"Effects,omitempty"`
	Influence       []Influence `json:"Influence,omitempty"`
	ReputationTrend string      `json:"ReputationTrend,omitempty"`
	Reputation      string      `json:"Reputation,omitempty"`
}

type Effect struct {
	Effect          string `json:"Effect"`
	EffectLocalised string `json:"Effect_Localised,omitempty"`
	Trend           string `json:"Trend"`
}

type Influence struct {
	SystemAddress      int64  `json:"SystemAddress"`
	Trend              string `json:"Trend"`
	InfluenceInfluence string `json:"Influence"`
}

type MaterialsReward struct {
	Name_Count_Common
	Category          string `json:"Category"`
	CategoryLocalised string `json:"Category_Localised,omitempty"`
}

// 8.23
type MissionFailedEvent struct {
	Base
	Name      string `json:"Name"`
	MissionId int64  `json:"MissionID"`
	Fine      int64  `json:"Fine,omitempty"`
}

// 8.24
type MissionRedirectedEvent struct {
	Base
	MissionId             int64  `json:"MissionID"`
	Name                  string `json:"Name"`
	NewDestinationStation string `json:"NewDestinationStation"`
	NewDestinationSystem  string `json:"NewDestinationSystem"`
	OldDestinationStation string `json:"OldDestinationStation"`
	OldDestinationSystem  string `json:"OldDestinationSystem"`
}

// 8.25
type ModuleBuyEvent struct {
	Base
	MarketId         int64  `json:"MarketID"`
	Slot             string `json:"Slot"`
	BuyItem          string `json:"BuyItem"`
	BuyItemLocalised string `json:"BuyItem_Localised,omitempty"`
	BuyPrice         int64  `json:"BuyPrice"`
	Ship             string `json:"Ship"`
	ShipId           int64  `json:"ShipID"`
	// if stored
	StoredItem string `json:"StoredItem,omitempty"`
	// if sold
	SellItem          string `json:"SellItem,omitempty"`
	SellItemLocalised string `json:"SellItem_Localised,omitempty"`
	SellPrice         int64  `json:"SellPrice,omitempty"`
}

// 8.26
type ModuleRetrieveEvent struct {
	Base
	MarketId               int64   `json:"MarketID"`
	Slot                   string  `json:"Slot"`
	RetrievedItem          string  `json:"RetrievedItem"`
	RetrievedItemLocalised string  `json:"RetrievedItem_Localised,omitempty"`
	Ship                   string  `json:"Ship"`
	ShipId                 int64   `json:"ShipID"`
	Hot                    bool    `json:"Hot"`
	EngineerModifications  string  `json:"EngineerModifications,omitempty"`
	Level                  int64   `json:"Level"`
	Quality                float64 `json:"Quality"`
	SwapOutItem            string  `json:"SwapOutItem,omitempty"`
	SwapOutItemLocalised   string  `json:"SwapOutItem_Localised,omitempty"`
	Cost                   int64   `json:"Cost,omitempty"`
}

// 8.27
type ModuleSellEvent struct {
	Base
	MarketId          int64  `json:"MarketID"`
	Slot              string `json:"Slot"`
	SellItem          string `json:"SellItem"`
	SellItemLocalised string `json:"SellItem_Localised,omitempty"`
	SellPrice         int64  `json:"SellPrice"`
	Ship              string `json:"Ship"`
	ShipId            int64  `json:"ShipID"`
}

// 8.28
type ModuleSellRemoteEvent struct {
	Base
	StorageSlot       int64  `json:"StorageSlot"`
	SellItem          string `json:"SellItem"`
	SellItemLocalised string `json:"SellItem_Localised,omitempty"`
	ServerId          int64  `json:"ServerId"`
	SellPrice         int64  `json:"SellPrice"`
	Ship              string `json:"Ship"`
	ShipId            int64  `json:"ShipID"`
}

// 8.29
type ModuleStoreEvent struct {
	Base
	MarketId              int64   `json:"MarketID"`
	Slot                  string  `json:"Slot"`
	StoredItem            string  `json:"StoredItem"`
	StoredItemLocalised   string  `json:"StoredItem_Localised,omitempty"`
	Ship                  string  `json:"Ship"`
	ShipId                int64   `json:"ShipID"`
	Hot                   bool    `json:"Hot"`
	EngineerModifications string  `json:"EngineerModifications"`
	Level                 int64   `json:"Level"`
	Quality               float64 `json:"Quality"`
	ReplacementItem       string  `json:"ReplacementItem,omitempty"`
	Cost                  int64   `json:"Cost,omitempty"`
}

// 8.30
type ModuleSwapEvent struct {
	Base
	MarketId          int64  `json:"MarketID"`
	FromSlot          string `json:"FromSlot"`
	ToSlot            string `json:"ToSlot"`
	FromItem          string `json:"FromItem"`
	FromItemLocalised string `json:"FromItem_Localised,omitempty"`
	ToItem            string `json:"ToItem"`
	ToItemLocalised   string `json:"ToItem_Localised,omitempty"`
	Ship              string `json:"Ship"`
	ShipId            int64  `json:"ShipID"`
}

// 8.31
// Full price list written to outfitting.json
type OutfittingEvent struct {
	Base
	MarketId    int64  `json:"MarketID"`
	StationName string `json:"StationName"`
	StarSystem  string `json:"StarSystem"`
}

// 8.32
type PayBountiesEvent struct {
	Base
	Amount           int64   `json:"Amount"`
	AllFines         bool    `json:"AllFines"`
	Faction          string  `json:"Faction"`
	FactionLocalised string  `json:"Faction_Localised,omitempty"`
	ShipId           int64   `json:"ShipID,omitempty"`
	BrokerPercentage float64 `json:"BrokerPercentage,omitempty"`
}

// 8.33
type PayFinesEvent struct {
	Base
	Amount           int64   `json:"Amount"`
	AllFines         bool    `json:"AllFines"`
	Faction          string  `json:"Faction"`
	FactionLocalised string  `json:"Faction_Localised,omitempty"`
	ShipId           int64   `json:"ShipID"`
	BrokerPercentage float64 `json:"BrokerPercentage,omitempty"`
}

// 8.34 (obsolete)
// 8.35
type RedeemVoucherEvent struct {
	Base
	Type             string  `json:"Type"`
	Amount           int64   `json:"Amount"`
	Faction          string  `json:"Faction,omitempty"`
	BrokerPercentage float64 `json:"BrokerPercentage,omitempty"`
	Factions         []struct {
		Faction string `json:"Faction"`
		Amount  int64  `json:"Amount"`
	} `json:"Factions,omitempty"`
}

// 8.36
type RefuelAllEvent struct {
	Base
	Cost   int64   `json:"Cost"`
	Amount float64 `json:"Amount"`
}

// 8.37
type RefuelPartialEvent struct {
	Base
	Cost   int64   `json:"Cost"`
	Amount float64 `json:"Amount"`
}

// 8.38
type RepairEvent struct {
	Base
	Item string `json:"Item"`
	Cost int64  `json:"Cost"`
}

// 8.39
type RepairAllEvent struct {
	Base
	Cost int64 `json:"Cost"`
}

// 8.40
type RestockVehicleEvent struct {
	Base
	Type    string `json:"Type"`
	Loadout string `json:"Loadout"`
	Cost    int64  `json:"Cost"`
	Count   int64  `json:"Count"`
}

// 8.41
type ScientificResearchEvent struct {
	Base
	MarketId int64  `json:"MarketID"`
	Name     string `json:"Name"`
	Category string `json:"Category"`
	Count    int64  `json:"Count"`
}

// 8.42
type SearchAndRescueEvent struct {
	Base
	Name   string `json:"Name"`
	Count  int64  `json:"Count"`
	Reward int64  `json:"Reward"`
}

// 8.43
type SellDronesEvent struct {
	Base
	Type      string `json:"Type"`
	Count     int64  `json:"Count"`
	SellPrice int64  `json:"SellPrice"`
	TotalSale int64  `json:"TotalSale"`
}

// 8.44
type SellShipOnRebuyEvent struct {
	Base
	ShipType_Localized
	SellShipId int64 `json:"SellShipId"`
	ShipPrice  int64 `json:"ShipPrice"`
	System     int64 `json:"System,omitempty"`
}

// 8.45
type SetUserShipNameEvent struct {
	Base
	Ship         string `json:"Ship"`
	ShipId       int64  `json:"ShipID"`
	UserShipName string `json:"UserShipName"`
	UserShipId   string `json:"UserShipId"`
}

// 8.46
// full price listed in shipyard.json when this event triggers.
type ShipyardEvent struct {
	Base
	MarketId    int64  `json:"MarketID,omitempty"`
	StationName string `json:"StationName,omitempty"`
	StarSystem  string `json:"StarSystem,omitempty"`
}

// 8.47
type ShipyardBuyEvent struct {
	Base
	ShipType_Localized
	MarketId     int64  `json:"MarketID"`
	ShipPrice    int64  `json:"ShipPrice"`
	StoreOldShip string `json:"StoreOldShip,omitempty"`
	StoreShipId  int64  `json:"StoreShipID,omitempty"`
	SellOldShip  string `json:"SellOldShip,omitempty"`
	SellShipId   int64  `json:"SellShipID,omitempty"`
	SellPrice    int64  `json:"SellPrice"`
}

// 8.48
type ShipyardNewEvent struct {
	Base
	ShipType_Localized
	ShipId int64 `json:"NewShipID"`
}

// 8.49
type ShipyardSellEvent struct {
	Base
	ShipType_Localized
	MarketId  int64 `json:"MarketID"`
	ShipId    int64 `json:"SellShipID"`
	ShipPrice int64 `json:"ShipPrice"`
	System    int64 `json:"System,omitempty"`
}

// 8.50
type ShipyardTransferEvent struct {
	Base
	ShipType_Localized
	MarketId      int64   `json:"MarketID"`
	ShipId        int64   `json:"ShipID"`
	System        string  `json:"System"`
	ShipMarketId  int64   `json:"ShipMarketID,omitempty"`
	Distance      float64 `json:"Distance"`
	TransferPrice int64   `json:"TransferPrice"`
	TransferTime  int64   `json:"TransferTime"`
}

// 8.51
type ShipyardSwapEvent struct {
	Base
	ShipType_Localized
	MarketId     int64  `json:"MarketID"`
	ShipId       int64  `json:"ShipID"`
	StoreOldShip string `json:"StoreOldShip,omitempty"`
	StoreShipId  int64  `json:"StoreShipID,omitempty"`
	SellOldShip  string `json:"SellOldShip,omitempty"`
	SellShipId   int64  `json:"SellShipID,omitempty"`
}

// 8.52
type StoredModulesEvent struct {
	Base
	MarketId    int64              `json:"MarketID,omitempty"`
	StationName string             `json:"StationName,omitempty"`
	StarSystem  string             `json:"StarSystem,omitempty"`
	Items       []StoredModuleItem `json:"Items,omitempty"`
}

type StoredModuleItem struct {
	Name_Localized
	StorageSlot           int64   `json:"StorageSlot"`
	StarSystem            string  `json:"StarSystem"`
	MarketId              int64   `json:"MarketID"`
	TransferCost          int64   `json:"TransferCost"`
	TransferTime          int64   `json:"TransferTime"`
	BuyPrice              int64   `json:"BuyPrice,omitempty"`
	Hot                   bool    `json:"Hot"`
	EngineerModifications string  `json:"EngineerModifications,omitempty"`
	Level                 int64   `json:"Level"`
	Quality               float64 `json:"Quality"`
	InTransit             bool    `json:"InTransit"`
}

// 8.53
type StoredShipsEvent struct {
	Base
	StationName string       `json:"StationName"`
	MarketId    int64        `json:"MarketID"`
	StarSystem  string       `json:"StarSystem"`
	ShipsHere   []StoredShip `json:"ShipsHere,omitempty"`
	ShipsRemote []StoredShip `json:"ShipsRemote,omitempty"`
}

type StoredShip struct {
	ShipType_Localized
	ShipID        int64  `json:"ShipID"`
	Name          string `json:"Name,omitempty"`
	Value         int64  `json:"Value"`
	Hot           bool   `json:"Hot"`
	InTransit     bool   `json:"InTransit,omitempty"`
	StarSystem    string `json:"StarSystem,omitempty"`
	ShipMarketId  int64  `json:"ShipMarketID,omitempty"`
	TransferPrice int64  `json:"TransferPrice,omitempty"`
	TransferTime  int64  `json:"TransferTime,omitempty"`
}

// 8.54
type TechnologyBrokerEvent struct {
	Base
	BrokerType    string              `json:"BrokerType,omitempty"`
	MarketId      int64               `json:"MarketID,omitempty"`
	ItemsUnlocked []Name_Localized    `json:"ItemsUnlocked,omitempty"`
	Commodities   []Name_Count_Common `json:"Commodities,omitempty"`
	Materials     []Material          `json:"Materials,omitempty"`
}

type Material struct {
	Name_Count_Common
	Category string `json:"Category"`
}

// 9 Powerplay
type PowerplayCollectEvent struct {
	Base
	Power         string `json:"Power"`
	Type          string `json:"Type"`
	TypeLocalised string `json:"Type_Localised,omitempty"`
	Count         int64  `json:"Count"`
}

type PowerplayDefectEvent struct {
	Base
	FromPower string `json:"FromPower"`
	ToPower   string `json:"ToPower"`
}

type PowerplayDeliverEvent struct {
	Base
	Power         string `json:"Power"`
	Type          string `json:"Type"`
	TypeLocalised string `json:"Type_Localised,omitempty"`
	Count         int64  `json:"Count"`
}

type PowerplayFastTrackEvent struct {
	Base
	Power string `json:"Power"`
	Cost  int64  `json:"Cost"`
}

type PowerplayJoinEvent struct {
	Base
	Power string `json:"Power"`
}

type PowerplayLeaveEvent struct {
	Base
	Power string `json:"Power"`
}

type PowerplaySalaryEvent struct {
	Base
	Power  string `json:"Power"`
	Amount int64  `json:"Amount"`
}

type PowerplayVoteEvent struct {
	Base
	Power             string `json:"Power"`
	Votes             int64  `json:"Votes"`
	VoteToConsolidate int64  `json:"VoteToConsolidate,omitempty"`
	System            int64  `json:"System"`
}

type PowerplayVoucherEvent struct {
	Base
	Power   string   `json:"Power"`
	Systems []string `json:"Systems"`
}

// 10 squadrons
type Base_Squadron struct {
	SquadronName string `json:"SquadronName"`
}

type AppliedToSquadronEvent struct {
	Base
	Base_Squadron
}

type DisbandedSquadronEvent struct {
	Base
	Base_Squadron
}

type InvitedToSquadronEvent struct {
	Base
	Base_Squadron
}

type JoinedSquadronEvent struct {
	Base
	Base_Squadron
}

type KickedFromSquadronEvent struct {
	Base
	Base_Squadron
}

type LeftSquadronEvent struct {
	Base
	Base_Squadron
}

type SharedBookmarkToSquadronEvent struct {
	Base
	Base_Squadron
}

type SquadronCreatedEvent struct {
	Base
	Base_Squadron
}

type SquadronDemotionEvent struct {
	Base
	Base_Squadron
	OldRank int `json:"OldRank"`
	NewRank int `json:"NewRank"`
}

type SquadronPromotionEvent struct {
	Base
	Base_Squadron
	OldRank int `json:"OldRank"`
	NewRank int `json:"NewRank"`
}

type SquadronStartupEvent struct {
	Base
	Base_Squadron
	CurrentRank int `json:"CurrentRank"`
}

type WonATrophyForSquadron struct {
	Base
	Base_Squadron
}

// 11 Other events

// 11.1
type AfmuRepairsEvent struct {
	Base
	Module          string  `json:"Module"`
	ModuleLocalised string  `json:"Module_Localised,omitempty"`
	FullyRepaired   bool    `json:"FullyRepaired"`
	Health          float64 `json:"Health"`
}

// 11.2
type ApproachSettlementEvent struct {
	Base
	Name_Localized
	BodyId        int64   `json:"BodyID,omitempty"`
	BodyName      string  `json:"BodyName,omitempty"`
	SystemAddress int64   `json:"SystemAddress,omitempty"`
	MarketId      int64   `json:"MarketID"`
	Latitude      float64 `json:"Latitude"`
	Longitude     float64 `json:"Longitude"`
}

// 11.3
type ChangeCrewRoleEvent struct {
	Base
	Role string `json:"Role"`
}

// 11.4
type CockpitBreachedEvent struct {
	Base
}

// 11.5
type CommitCrimeEvent struct {
	Base
	CrimeType       string `json:"CrimeType"`
	Faction         string `json:"Faction"`
	Victim          string `json:"Victim,omitempty"`
	VictimLocalised string `json:"Victim_Localised,omitempty"`
	Fine            int64  `json:"Bounty,omitempty"`
	Bounty          int64  `json:"Bounty,omitempty"`
}

// 11.6
type ContinuedEvent struct {
	Base
	Part string `json:"Part"`
}

// 11.7
type CrewLaunchFighterEvent struct {
	Base
	Crew string `json:"Crew"`
	ID   string `json:"ID"`
}

// 11.8
type CrewMemberJoinsEvent struct {
	Base
	Crew string `json:"Crew"`
}

// 11.9
type CrewMemberQuitsEvent struct {
	Base
	Crew string `json:"Crew"`
}

// 11.10
type CrewMemberRoleChangeEvent struct {
	Base
	Crew string `json:"Crew"`
	Role string `json:"Role"`
}

// 11.11
type CrimeVictimEvent struct {
	Base
	Offender  string `json:"Offender"`
	CrimeType string `json:"CrimeType"`
	Bounty    int64  `json:"Bounty"`
}

// 11.12
type DatalinkScanEvent struct {
	Base
	Message          string `json:"Message"`
	MessageLocalised string `json:"Message_Localised,omitempty"`
}

// 11.13
type DatalinkVoucherEvent struct {
	Base
	Reward        int64  `json:"Reward"`
	VictimFaction string `json:"VictimFaction"`
	PayeeFaction  string `json:"PayeeFaction"`
}

// 11.14
type DataScannedEvent struct {
	Base
	Type          string `json:"Type"`
	TypeLocalised string `json:"Type_Localised,omitempty"`
}

// 11.15
type DockFighterEvent struct {
	Base
	ID string `json:"ID"`
}

// 11.16
type DockSRVEvent struct {
	Base
	ID string `json:"ID"`
}

// 11.17
type EndCrewSessionEvent struct {
	Base
	OnCrime bool `json:"OnCrime,omitempty"`
}

// 11.18
type FighterRebuiltEvent struct {
	Base
	ID      string `json:"ID"`
	Loadout string `json:"Loadout"`
}

// 11.19
type FuelScoopEvent struct {
	Base
	Scooped float64 `json:"Scooped"`
	Total   float64 `json:"Total"`
}

// 11.20
type FriendsEvent struct {
	Base
	Status string `json:"Status"`
	Name   string `json:"Name"`
}

// 11.21
type JetConeBoostEvent struct {
	Base
	BoostValue float64 `json:"BoostValue"`
}

// 11.22
type JetConeDamageEvent struct {
	Base
	Module          string `json:"Module"`
	ModuleLocalised string `json:"Module_Localised,omitempty"`
}

// 11.23
type JoinACrewEvent struct {
	Base
	Captain string `json:"Captain"`
}

// 11.24
type KickCrewMemberEvent struct {
	Base
	Crew    string `json:"Crew"`
	OnCrime bool   `json:"OnCrime"`
}

// 11.25
type LaunchDroneEvent struct {
	Base
	Type string `json:"Type"`
}

// 11.26
type LaunchFighterEvent struct {
	Base
	Loadout          string `json:"Loadout"`
	ID               string `json:"ID"`
	PlayerControlled bool   `json:"PlayerControlled"`
}

// 11.27
type LaunchSRVEvent struct {
	Base
	Loadout          string `json:"Loadout"`
	ID               string `json:"ID"`
	PlayerControlled bool   `json:"PlayerControlled"`
}

// 11.28
// ModuleInfo is equivalent to modules.json.

// 11.29
type MusicEvent struct {
	Base
	MusicTrack string `json:"MusicTrack"`
}

// 11.30
type NpcCrewPaidWageEvent struct {
	Base
	NpcCrewName string `json:"NpcCrewName"`
	NpcCrewId   int64  `json:"NpcCrewId"`
	Amount      int64  `json:"Amount"`
}

// 11.31
type NpcCrewRankEvent struct {
	Base
	NpcCrewName string `json:"NpcCrewName"`
	NpcCrewId   int64  `json:"NpcCrewId"`
	RankCombat  int    `json:"RankCombat"`
}

// 11.32
type PromotionEvent struct {
	Base
	Combat     int `json:"Combat,omitempty"`
	Trade      int `json:"Trade,omitempty"`
	Explore    int `json:"Explore,omitempty"`
	CQC        int `json:"CQC,omitempty"`
	Federation int `json:"Federation,omitempty"`
	Empire     int `json:"Empire,omitempty"`
}

// 11.33
type ProspectedAsteroidEvent struct {
	Base
	Materials []struct {
		Name_Localized
		Event      string  `json:"event"`
		Proportion float64 `json:"Proportion"`
	} `json:"Materials,omitempty"`
	MotherlodeMaterial          string `json:"MotherlodeMaterial,omitempty"`
	MotherlodeMaterialLocalised string `json:"MotherlodeMaterial_Localised,omitempty"`
	Content                     string `json:"Content"`
	ContentLocalised            string `json:"Content_Localised,omitempty"`
	Remaining                   int    `json:"Remaining"`
}

// 11.34
type QuitACrewEvent struct {
	Base
	Captain string `json:"Captain"`
}

// 11.35
type RebootRepairEvent struct {
	Base
	Modules []string `json:"Modules"`
}

// 11.36
type ReceiveTextEvent struct {
	Base
	From             string `json:"From"`
	FromLocalised    string `json:"From_Localised,omitempty"`
	Message          string `json:"Message"`
	MessageLocalised string `json:"Message_Localised,omitempty"`
	Channel          string `json:"Channel"`
}

// 11.37
type RepairDroneEvent struct {
	Base
	HullRepaired      float64 `json:"HullRepaired"`
	CockpitRepaired   float64 `json:"CockpitRepaired"`
	CorrosionRepaired float64 `json:"CorrosionRepaired"`
}

// 11.38
type ReservoirReplenishedEvent struct {
	Base
	FuelMain      float64 `json:"FuelMain"`
	FuelReservoir float64 `json:"FuelReservoir"`
}

// 11.39
type ResurrectEvent struct {
	Base
	Option   string `json:"Option"`
	Cost     int64  `json:"Cost"`
	Bankrupt bool   `json:"Bankrupt"`
}

// 11.40
type ScannedEvent struct {
	Base
	ScanType string `json:"ScanType"`
}

// 11.41
type SelfDestructEvent struct {
	Base
}

// 11.42
type SendTextEvent struct {
	Base
	To          string `json:"To"`
	ToLocalised string `json:"To_Localised,omitempty"`
	Message     string `json:"Message"`
}

// 11.43
type ShutdownEvent struct {
	Base
}

// 11.44
type SynthesisEvent struct {
	Base
	Name      string              `json:"Name"`
	Materials []Name_Count_Common `json:"Materials,"`
}

// 11.45
type SystemsShutdownEvent struct {
	Base
}

// 11.46
type USSDropEvent struct {
	Base
	UssType          string `json:"USSType"`
	UssTypeLocalised string `json:"USSType_Localised,omitempty"`
	UssThreat        int64  `json:"USSThreat"`
}

// 11.47
type VehicleSwitchEvent struct {
	Base
	To string `json:"To"`
}

// 11.48
type WingAddEvent struct {
	Base
	Name string `json:"Name"`
}

// 11.49
type WingInviteEvent struct {
	Base
	Name string `json:"Name"`
}

// 11.50
type WingJoinEvent struct {
	Base
	Others []string `json:"Others,omitempty"`
}

// 11.51
type WingLeaveEvent struct {
	Base
}
