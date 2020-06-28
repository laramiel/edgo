package journal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"
)

// Here we register all the events and dynamically dispatch the json
// to the correct type.
type typeRegistry map[string]reflect.Type

func (t typeRegistry) Set(i interface{}) {
	v := reflect.TypeOf(i)
	t[v.Name()] = v
}

func (t typeRegistry) Get(name string) (interface{}, error) {
	if typ, ok := t[name]; ok {
		v := reflect.New(typ).Elem()
		return v.Addr().Interface(), nil
	}
	return nil, fmt.Errorf("Missing %s", name)
}

var eventRegistry = make(typeRegistry)

func init() {
	myTypes := []interface{}{
		FileheaderEvent{},
		CargoEvent{},
		ClearSavedGameEvent{},
		CommanderEvent{},
		LoadoutEvent{},
		MaterialsEvent{},
		MissionsEvent{},
		NewCommanderEvent{},
		LoadGameEvent{},
		PassengersEvent{},
		PowerplayEvent{},
		ProgressEvent{},
		RankEvent{},
		ReputationEvent{},
		StatisticsEvent{},
		ApproachBodyEvent{},
		DockedEvent{},
		DockingCancelledEvent{},
		DockingDeniedEvent{},
		DockingGrantedEvent{},
		DockingRequestedEvent{},
		DockingTimeoutEvent{},
		FSDJumpEvent{},
		FSDTargetEvent{},
		LeaveBodyEvent{},
		LiftoffEvent{},
		LocationEvent{},
		StartJumpEvent{},
		SupercruiseEntryEvent{},
		SupercruiseExitEvent{},
		TouchdownEvent{},
		UndockedEvent{},
		BountyEvent{},
		DiedEvent{},
		EscapeInterdictionEvent{},
		FactionKillBondEvent{},
		FighterDestroyedEvent{},
		HeatDamageEvent{},
		HeatWarningEvent{},
		HullDamageEvent{},
		InterdictedEvent{},
		InterdictionEvent{},
		PVPKillEvent{},
		ShieldStateEvent{},
		ShipTargetedEvent{},
		SRVDestroyedEvent{},
		UnderAttackEvent{},
		CodexEntryEvent{},
		DiscoveryScanEvent{},
		ScanEvent{},
		FSSAllBodiesFoundEvent{},
		FSSDiscoveryScanEvent{},
		FSSSignalDiscoveredEvent{},
		MaterialCollectedEvent{},
		MaterialDiscardedEvent{},
		MaterialDiscoveredEvent{},
		MultiSellExplorationDataEvent{},
		NavBeaconScanEvent{},
		BuyExplorationDataEvent{},
		SAAScanCompleteEvent{},
		SellExplorationDataEvent{},
		ScreenshotEvent{},
		AsteroidCrackedEvent{},
		BuyTradeDataEvent{},
		CollectCargoEvent{},
		EjectCargoEvent{},
		MarketBuyEvent{},
		MarketSellEvent{},
		MiningRefinedEvent{},
		BuyAmmoEvent{},
		BuyDronesEvent{},
		CargoDepotEvent{},
		CommunityGoalEvent{},
		CommunityGoalDiscardEvent{},
		CommunityGoalJoinEvent{},
		CommunityGoalRewardEvent{},
		CrewAssignEvent{},
		CrewFireEvent{},
		CrewHireEvent{},
		EngineerContributionEvent{},
		EngineerCraftEvent{},
		EngineerProgressEvent{},
		FetchRemoteModuleEvent{},
		MarketEvent{},
		MassModuleStoreEvent{},
		MaterialTradeEvent{},
		MissionAbandonedEvent{},
		MissionAcceptedEvent{},
		MissionCompletedEvent{},
		MissionFailedEvent{},
		MissionRedirectedEvent{},
		ModuleBuyEvent{},
		ModuleRetrieveEvent{},
		ModuleSellEvent{},
		ModuleSellRemoteEvent{},
		ModuleStoreEvent{},
		ModuleSwapEvent{},
		OutfittingEvent{},
		PayBountiesEvent{},
		PayFinesEvent{},
		RedeemVoucherEvent{},
		RefuelAllEvent{},
		RefuelPartialEvent{},
		RepairEvent{},
		RepairAllEvent{},
		RestockVehicleEvent{},
		ScientificResearchEvent{},
		SearchAndRescueEvent{},
		SellDronesEvent{},
		SellShipOnRebuyEvent{},
		SetUserShipNameEvent{},
		ShipyardEvent{},
		ShipyardBuyEvent{},
		ShipyardNewEvent{},
		ShipyardSellEvent{},
		ShipyardTransferEvent{},
		ShipyardSwapEvent{},
		StoredModulesEvent{},
		StoredShipsEvent{},
		TechnologyBrokerEvent{},
		PowerplayCollectEvent{},
		PowerplayDefectEvent{},
		PowerplayDeliverEvent{},
		PowerplayFastTrackEvent{},
		PowerplayJoinEvent{},
		PowerplayLeaveEvent{},
		PowerplaySalaryEvent{},
		PowerplayVoteEvent{},
		PowerplayVoucherEvent{},
		AppliedToSquadronEvent{},
		DisbandedSquadronEvent{},
		InvitedToSquadronEvent{},
		JoinedSquadronEvent{},
		KickedFromSquadronEvent{},
		LeftSquadronEvent{},
		SharedBookmarkToSquadronEvent{},
		SquadronCreatedEvent{},
		SquadronDemotionEvent{},
		SquadronPromotionEvent{},
		SquadronStartupEvent{},
		AfmuRepairsEvent{},
		ApproachSettlementEvent{},
		ChangeCrewRoleEvent{},
		CockpitBreachedEvent{},
		CommitCrimeEvent{},
		ContinuedEvent{},
		CrewLaunchFighterEvent{},
		CrewMemberJoinsEvent{},
		CrewMemberQuitsEvent{},
		CrewMemberRoleChangeEvent{},
		CrimeVictimEvent{},
		DatalinkScanEvent{},
		DatalinkVoucherEvent{},
		DataScannedEvent{},
		DockFighterEvent{},
		DockSRVEvent{},
		EndCrewSessionEvent{},
		FighterRebuiltEvent{},
		FuelScoopEvent{},
		FriendsEvent{},
		JetConeBoostEvent{},
		JetConeDamageEvent{},
		JoinACrewEvent{},
		KickCrewMemberEvent{},
		LaunchDroneEvent{},
		LaunchFighterEvent{},
		LaunchSRVEvent{},
		MusicEvent{},
		NpcCrewPaidWageEvent{},
		NpcCrewRankEvent{},
		PromotionEvent{},
		ProspectedAsteroidEvent{},
		QuitACrewEvent{},
		RebootRepairEvent{},
		ReceiveTextEvent{},
		RepairDroneEvent{},
		ReservoirReplenishedEvent{},
		ResurrectEvent{},
		ScannedEvent{},
		SelfDestructEvent{},
		SendTextEvent{},
		ShutdownEvent{},
		SynthesisEvent{},
		SystemsShutdownEvent{},
		USSDropEvent{},
		VehicleSwitchEvent{},
		WingAddEvent{},
		WingInviteEvent{},
		WingJoinEvent{},
		WingLeaveEvent{}}
	for _, v := range myTypes {
		eventRegistry.Set(v)
	}
}

func GetEventName(contents []byte) string {
	// The quick and dirty way is to find the event field and extract
	// the value from it.
	idx := bytes.Index(contents, []byte(`"event"`))
	if idx != -1 {
		tmp := contents[idx+7 : len(contents)-1]
		f := bytes.SplitN(tmp, []byte(`"`), 3)
		return string(f[1])
	}

	// The slower way is to just unmarshal into a base object, read
	// the event, and then use that.
	var b Base
	if err := json.Unmarshal(contents, &b); err != nil {
		return ""
	}
	return b.Event
}

func ParseJournalLine(contents []byte) (interface{}, error) {
	event := GetEventName(contents)
	obj, err := eventRegistry.Get(event + "Event")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(contents, obj)
	return obj, err
}

func SetRAW(obj interface{}, contents string) {
	v := reflect.Indirect(reflect.ValueOf(obj)).FieldByName("RAW")
	if v.Kind() == reflect.String {
		v.SetString(contents)
	}
}

func GetRAW(obj interface{}) string {
	v := reflect.Indirect(reflect.ValueOf(obj)).FieldByName("RAW")
	if v.Kind() == reflect.String {
		return v.String()
	}
	return ""
}

func GetEvent(obj interface{}) string {
	idr := reflect.Indirect(reflect.ValueOf(obj))
	v := idr.FieldByName("Event")
	if v.Kind() == reflect.String {
		return v.String()
	}
	return reflect.TypeOf(idr).Name()
}

func IsStatusFile(filename string) bool {
	base := strings.ToLower(filepath.Base(filename))
	switch base {
	case "cargo.json":
		return true
	case "market.json":
		return true
	case "modulesinfo.json":
		return true
	case "navroute.json":
		return true
	case "outfitting.json":
		return true
	case "shipyard.json":
		return true
	case "status.json":
		return true
	default:
		return false
	}
}

func GetStatusInterface(filename string) interface{} {
	base := strings.ToLower(filepath.Base(filename))
	switch base {
	case "cargo.json":
		return &Cargo{}
	case "market.json":
		return &Market{}
	case "modulesinfo.json":
		return &ModulesInfo{}
	case "navroute.json":
		return &NavRoute{}
	case "outfitting.json":
		return &Outfitting{}
	case "shipyard.json":
		return &Shipyard{}
	case "status.json":
		return &Status{}
	default:
		return make(map[string]interface{})
	}
}

func ParseStatusContents(filename string, content []byte) (interface{}, error) {
	if !IsStatusFile(filename) {
		return nil, errors.New("Not a status file")
	}
	obj := GetStatusInterface(filename)
	err := json.Unmarshal(content, obj)
	return obj, err
}

func ParseStatusData(filename string) (interface{}, error) {
	if !IsStatusFile(filename) {
		return nil, errors.New("Not a status file")
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	obj := GetStatusInterface(filename)
	err = json.Unmarshal(content, obj)
	return obj, err
}
