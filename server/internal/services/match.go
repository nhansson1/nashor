package services

import (
	"fmt"
	"nashor/internal/helpers"
	"sync"
)

const matchBase = "/lol/match/v5/matches"

type MatchService struct {
	riotClient *RiotClient
}

func NewMatchService(rc *RiotClient) MatchService {
	return MatchService{
		riotClient: rc,
	}
}

func (s MatchService) GetMatchIdsByPuuid(server, puuid, start, count string) ([]string, error) {
	queries := make(map[string]string)
	var ids []string

	queries["count"] = count
	queries["start"] = start

	resp, err := s.riotClient.Get(helpers.GetRegionFromServer(server), fmt.Sprintf(matchBase+"/by-puuid/%s/ids", puuid), queries)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ids, err = helpers.ParseBody[[]string](resp.Body)

	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (s MatchService) GetMatchDataById(server, id string, ch chan MatchDto, wg *sync.WaitGroup) {
	defer wg.Done()
	var d MatchDto

	resp, err := s.riotClient.Get(helpers.GetRegionFromServer(server), fmt.Sprintf(matchBase+"/%s", id), nil)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	d, _ = helpers.ParseBody[MatchDto](resp.Body)

	ch <- d
}

func (s MatchService) GetMatchDataByIds(server string, ids []string) ([]MatchDto, error) {
	ch := make(chan MatchDto)
	var wg sync.WaitGroup

	for _, id := range ids {
		wg.Add(1)
		go s.GetMatchDataById(server, id, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	matches := make([]MatchDto, 0)
	for r := range ch {
		matches = append(matches, r)
	}

	return matches, nil
}

type MatchDto struct {
	Metadata MetadataDto `json:"metadata"`
	Info     InfoDto     `json:"info"`
}

type MetadataDto struct {
	DataVersion  string   `json:"dataVersion"`
	MatchId      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type InfoDto struct {
	EndOfGameResult    string           `json:"endOfGameResult"`
	GameCreation       int              `json:"gameCreation"`
	GameDuration       int              `json:"gameDuration"`
	GameEndTimestamp   int              `json:"gameEndTimestamp"`
	GameId             int              `json:"gameId"`
	GameMode           string           `json:"gameMode"`
	GameName           string           `json:"gameName"`
	GameStartTimestamp int              `json:"gameStartTimestamp"`
	GameType           string           `json:"gameType"`
	GameVersion        string           `json:"gameVersion"`
	MapId              int              `json:"mapId"`
	Participants       []ParticipantDto `json:"participants"`
	PlatformId         string           `json:"platformId"`
	QueueId            int              `json:"queueId"`
	Teams              []TeamDto        `json:"teams"`
	TournamentCode     string           `json:"tournamentCode"`
}

type ParticipantDto struct {
	AllInPings                     int           `json:"allInPings"`
	AssistMePings                  int           `json:"assistMePings"`
	Assists                        int           `json:"assists"`
	BaronKills                     int           `json:"baronKills"`
	BountyLevel                    int           `json:"bountyLevel"`
	ChampExperience                int           `json:"champExperience"`
	ChampLevel                     int           `json:"champLevel"`
	ChampionId                     int           `json:"championId"`
	ChampionName                   string        `json:"championName"`
	CommandPings                   int           `json:"commandPings"`
	ChampionTransform              int           `json:"championTransform"`
	ConsumablesPurchased           int           `json:"consumablesPurchased"`
	Challenges                     ChallengesDto `json:"challenges"`
	DamageDealtToBuildings         int           `json:"damageDealtToBuildings"`
	DamageDealtToObjectives        int           `json:"damageDealtToObjectives"`
	DamageDealtToTurrets           int           `json:"damageDealtToTurrets"`
	DamageSelfMitigated            int           `json:"damageSelfMitigated"`
	Deaths                         int           `json:"deaths"`
	DetectorWardsPlaced            int           `json:"detectorWardsPlaced"`
	DoubleKills                    int           `json:"doubleKills"`
	DragonKills                    int           `json:"dragonKills"`
	EligibleForProgression         bool          `json:"eligibleForProgression"`
	EnemyMissingPings              int           `json:"enemyMissingPings"`
	EnemyVisionPings               int           `json:"enemyVisionPings"`
	FirstBloodAssist               bool          `json:"firstBloodAssist"`
	FirstBloodKill                 bool          `json:"firstBloodKill"`
	FirstTowerAssist               bool          `json:"firstTowerAssist"`
	FirstTowerKill                 bool          `json:"firstTowerKill"`
	GameEndedInEarlySurrender      bool          `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender           bool          `json:"gameEndedInSurrender"`
	HoldPings                      int           `json:"holdPings"`
	GetBackPings                   int           `json:"getBackPings"`
	GoldEarned                     int           `json:"goldEarned"`
	GoldSpent                      int           `json:"goldSpent"`
	IndividualPosition             string        `json:"individualPosition"`
	InhibitorKills                 int           `json:"inhibitorKills"`
	InhibitorTakedowns             int           `json:"inhibitorTakedowns"`
	InhibitorsLost                 int           `json:"inhibitorsLost"`
	Item0                          int           `json:"item0"`
	Item1                          int           `json:"item1"`
	Item2                          int           `json:"item2"`
	Item3                          int           `json:"item3"`
	Item4                          int           `json:"item4"`
	Item5                          int           `json:"item5"`
	Item6                          int           `json:"item6"`
	ItemsPurchased                 int           `json:"itemsPurchased"`
	KillingSprees                  int           `json:"killingSprees"`
	Kills                          int           `json:"kills"`
	Lane                           string        `json:"lane"`
	LargestCriticalStrike          int           `json:"largestCriticalStrike"`
	LargestKillingSpree            int           `json:"largestKillingSpree"`
	LargestMultiKill               int           `json:"largestMultiKill"`
	LongestTimeSpentLiving         int           `json:"longestTimeSpentLiving"`
	MagicDamageDealt               int           `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int           `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int           `json:"magicDamageTaken"`
	Missions                       MissionsDto   `json:"missions"`
	NeutralMinionsKilled           int           `json:"neutralMinionsKilled"`
	NeedVisionPings                int           `json:"needVisionPings"`
	NexusKills                     int           `json:"nexusKills"`
	NexusTakedowns                 int           `json:"nexusTakedowns"`
	NexusLost                      int           `json:"nexusLost"`
	ObjectivesStolen               int           `json:"objectivesStolen"`
	ObjectivesStolenAssists        int           `json:"objectivesStolenAssists"`
	OnMyWayPings                   int           `json:"onMyWayPings"`
	ParticipantId                  int           `json:"participantId"`
	PlayerScore0                   int           `json:"playerScore0"`
	PlayerScore1                   int           `json:"playerScore1"`
	PlayerScore2                   int           `json:"playerScore2"`
	PlayerScore3                   int           `json:"playerScore3"`
	PlayerScore4                   int           `json:"playerScore4"`
	PlayerScore5                   int           `json:"playerScore5"`
	PlayerScore6                   int           `json:"playerScore6"`
	PlayerScore7                   int           `json:"playerScore7"`
	PlayerScore8                   int           `json:"playerScore8"`
	PlayerScore9                   int           `json:"playerScore9"`
	PlayerScore10                  int           `json:"playerScore10"`
	PlayerScore11                  int           `json:"playerScore11"`
	PentaKills                     int           `json:"pentaKills"`
	Perks                          PerksDto      `json:"perks"`
	PhysicalDamageDealt            int           `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int           `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int           `json:"physicalDamageTaken"`
	Placement                      int           `json:"placement"`
	PlayerAugment1                 int           `json:"playerAugment1"`
	PlayerAugment2                 int           `json:"playerAugment2"`
	PlayerAugment3                 int           `json:"playerAugment3"`
	PlayerAugment4                 int           `json:"playerAugment4"`
	PlayerSubteamId                int           `json:"playerSubteamId"`
	PushPings                      int           `json:"pushPings"`
	ProfileIcon                    int           `json:"profileIcon"`
	Puuid                          string        `json:"puuid"`
	QuadraKills                    int           `json:"quadraKills"`
	RiotIdGameName                 string        `json:"riotIdGameName"`
	RiotIdTagLine                  string        `json:"riotIdTagLine"`
	Role                           string        `json:"role"`
	SightWardsBoughtInGame         int           `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int           `json:"spell1Casts"`
	Spell2Casts                    int           `json:"spell2Casts"`
	Spell3Casts                    int           `json:"spell3Casts"`
	Spell4Casts                    int           `json:"spell4Casts"`
	SubteamPlacements              int           `json:"subteamPlacements"`
	Summoner1Casts                 int           `json:"summoner1Casts"`
	Summoner1Id                    int           `json:"summoner1Id"`
	Summoner2Casts                 int           `json:"summoner2Casts"`
	Summoner2Id                    int           `json:"summoner2Id"`
	SummonerId                     string        `json:"summonerId"`
	SummonerLevel                  int           `json:"summonerLevel"`
	SummonerName                   string        `json:"summonerName"`
	TeamEarlySurrendered           bool          `json:"teamEarlySurrendered"`
	TeamId                         int           `json:"teamId"`
	TeamPosition                   string        `json:"teamPosition"`
	TimeCCingOthers                int           `json:"timeCCingOthers"`
	TimePlayed                     int           `json:"timePlayed"`
	TotalAllyJungleMinionsKilled   int           `json:"totalAllyJungleMinionsKilled"`
	TotalDamageDealt               int           `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int           `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int           `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int           `json:"totalDamageTaken"`
	TotalEnemyJungleMinionsKilled  int           `json:"totalEnemyJungleMinionsKilled"`
	TotalHeal                      int           `json:"totalHeal"`
	TotalHealsOnTeammates          int           `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int           `json:"totalMinionsKilled"`
	TotalTimeCCdealt               int           `json:"totalTimeCCdealt"`
	TotalTimeSpentDead             int           `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int           `json:"totalUnitsHealed"`
	TripleKills                    int           `json:"tripleKills"`
	TrueDamageDealt                int           `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int           `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int           `json:"trueDamageTaken"`
	TurretKills                    int           `json:"turretKills"`
	TurretTakedowns                int           `json:"turretTakedowns"`
	UnrealKills                    int           `json:"unrealKills"`
	VisionScore                    int           `json:"visionScore"`
	VisionClearedPings             int           `json:"visionClearedPings"`
	VisionWardsBoughtInGame        int           `json:"visionWardsBoughtInGame"`
	WardsKilled                    int           `json:"wardsKilled"`
	WardsPlaced                    int           `json:"wardsPlaced"`
	Win                            bool          `json:"win"`
}

type TeamDto struct {
	Bans       BanDto        `json:"bans"`
	Objectives ObjectivesDto `json:"objectives"`
	TeamId     int           `json:"teamId"`
	Win        bool          `json:"win"`
}

type BanDto struct {
	ChampionId int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

type ObjectivesDto struct {
	Baron      ObjectiveDto `json:"baron"`
	Champion   ObjectiveDto `json:"champion"`
	Dragon     ObjectiveDto `json:"dragon"`
	Horde      ObjectiveDto `json:"horde"`
	Inhibitor  ObjectiveDto `json:"inhibitor"`
	RiftHerald ObjectiveDto `json:"riftHerald"`
	Tower      ObjectiveDto `json:"tower"`
}

type ObjectiveDto struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}

type ChallengesDto struct {
	AssistStreakCount                         int     `json:"12AssistStreakCount"`
	BaronBuffGoldAdvantageOverThreshold       int     `json:"baronBuffGoldAdvantageOverThreshold"`
	ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf"`
	EarliestBaron                             int     `json:"earliestBaron"`
	EarliestDragonTakedown                    int     `json:"earliestDragonTakedown"`
	EarliestElderDragon                       int     `json:"earliestElderDragon"`
	EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
	FasterSupportQuestCompletion              int     `json:"fasterSupportQuestCompletion"`
	FastestLegendary                          int     `json:"fastestLegendary"`
	HadAfkTeammate                            int     `json:"hadAfkTeammate"`
	HighestChampionDamage                     int     `json:"highestChampionDamage"`
	HighestCrowdControlScore                  int     `json:"highestCrowdControlScore"`
	HighestWardKills                          int     `json:"highestWardKills"`
	JunglerKillsEarlyJungle                   int     `json:"junglerKillsEarlyJungle"`
	KillsOnLanersEarlyJungleAsJungler         int     `json:"killsOnLanersEarlyJungleAsJungler"`
	LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
	LegendaryCount                            int     `json:"legendaryCount"`
	MaxCsAdvantageOnLaneOpponent              float64 `json:"maxCsAdvantageOnLaneOpponent"`
	MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
	MostWardsDestroyedOneSweeper              int     `json:"mostWardsDestroyedOneSweeper"`
	MythicItemUsed                            int     `json:"mythicItemUsed"`
	PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
	SoloTurretsLategame                       int     `json:"soloTurretsLategame"`
	TakedownsFirst25Minutes                   int     `json:"takedownsFirst25Minutes"`
	TeleportTakedowns                         int     `json:"teleportTakedowns"`
	ThirdInhibitorDestroyedTime               int     `json:"thirdInhibitorDestroyedTime"`
	ThreeWardsOneSweeperCount                 int     `json:"threeWardsOneSweeperCount"`
	VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
	InfernalScalePickup                       int     `json:"infernalScalePickup"`
	FistBumpParticipation                     int     `json:"fistBumpParticipation"`
	VoidMonsterKill                           int     `json:"voidMonsterKill"`
	AbilityUses                               int     `json:"abilityUses"`
	AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
	AlliedJungleMonsterKills                  float64 `json:"alliedJungleMonsterKills"`
	BaronTakedowns                            int     `json:"baronTakedowns"`
	BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
	BountyGold                                int     `json:"bountyGold"`
	BuffsStolen                               int     `json:"buffsStolen"`
	CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
	ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
	DamagePerMinute                           float64 `json:"damagePerMinute"`
	DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
	DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
	DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
	DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
	DoubleAces                                int     `json:"doubleAces"`
	DragonTakedowns                           int     `json:"dragonTakedowns"`
	LegendaryItemUsed                         []int   `json:"legendaryItemUsed"`
	EffectiveHealAndShielding                 float64 `json:"effectiveHealAndShielding"`
	ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
	ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
	EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
	EnemyJungleMonsterKills                   float64 `json:"enemyJungleMonsterKills"`
	EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
	EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
	EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
	EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
	FirstTurretKilled                         int     `json:"firstTurretKilled"`
	FirstTurretKilledTime                     float64 `json:"firstTurretKilledTime"`
	FlawlessAces                              int     `json:"flawlessAces"`
	FullTeamTakedown                          int     `json:"fullTeamTakedown"`
	GameLength                                float64 `json:"gameLength"`
	GetTakedownsInAllLanesEarlyJungleAsLaner  int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
	GoldPerMinute                             float64 `json:"goldPerMinute"`
	HadOpenNexus                              int     `json:"hadOpenNexus"`
	ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
	InitialBuffCount                          int     `json:"initialBuffCount"`
	InitialCrabCount                          int     `json:"initialCrabCount"`
	JungleCsBefore10Minutes                   float64 `json:"jungleCsBefore10Minutes"`
	JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
	Kda                                       float64 `json:"kda"`
	KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
	KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
	KillingSprees                             int     `json:"killingSprees"`
	KillParticipation                         float64 `json:"killParticipation"`
	KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
	KillsOnOtherLanesEarlyJungleAsLaner       int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
	KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
	KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
	KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
	KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
	KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
	LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
	LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
	LostAnInhibitor                           int     `json:"lostAnInhibitor"`
	MaxKillDeficit                            int     `json:"maxKillDeficit"`
	MejaisFullStackInTime                     int     `json:"mejaisFullStackInTime"`
	MoreEnemyJungleThanOpponent               float64 `json:"moreEnemyJungleThanOpponent"`
	MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
	Multikills                                int     `json:"multikills"`
	MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
	MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
	OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
	OutnumberedKills                          int     `json:"outnumberedKills"`
	OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
	PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
	PerfectGame                               int     `json:"perfectGame"`
	PickKillWithAlly                          int     `json:"pickKillWithAlly"`
	PoroExplosions                            int     `json:"poroExplosions"`
	QuickCleanse                              int     `json:"quickCleanse"`
	QuickFirstTurret                          int     `json:"quickFirstTurret"`
	QuickSoloKills                            int     `json:"quickSoloKills"`
	RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
	SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
	ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
	ShortestTimeToAceFromFirstTakedown        float64 `json:"shortestTimeToAceFromFirstTakedown"`
	SkillshotsDodged                          int     `json:"skillshotsDodged"`
	SkillshotsHit                             int     `json:"skillshotsHit"`
	SnowballsHit                              int     `json:"snowballsHit"`
	SoloBaronKills                            int     `json:"soloBaronKills"`
	SwarmDefeatAatrox                         int     `json:"SWARM_DefeatAatrox"`
	SwarmDefeatBriar                          int     `json:"SWARM_DefeatBriar"`
	SwarmDefeatMiniBosses                     int     `json:"SWARM_DefeatMiniBosses"`
	SwarmEvolveWeapon                         int     `json:"SWARM_EvolveWeapon"`
	SwarmKillEnemy                            int     `json:"SWARM_KillEnemy"`
	SwarmPickupGold                           float64 `json:"SWARM_PickupGold"`
	SwarmReachLevel50                         int     `json:"SWARM_ReachLevel50"`
	SwarmSurvive15Min                         int     `json:"SWARM_Survive15Min"`
	SwarmWinWith5EvolvedWeapons               int     `json:"SWARM_WinWith5EvolvedWeapons"`
	SoloKills                                 int     `json:"soloKills"`
	StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
	SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
	SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
	TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
	Takedowns                                 int     `json:"takedowns"`
	TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
	TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
	TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
	TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
	TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
	TeamBaronKills                            int     `json:"teamBaronKills"`
	TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
	TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
	TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
	TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
	TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
	TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
	TurretTakedowns                           int     `json:"turretTakedowns"`
	TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
	TwoWardsOneSweeperCount                   int     `json:"twoWardsOneSweeperCount"`
	UnseenRecalls                             int     `json:"unseenRecalls"`
	VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
	WardsGuarded                              int     `json:"wardsGuarded"`
	WardTakedowns                             int     `json:"wardTakedowns"`
	WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
}

type MissionsDto struct {
	PlayerScore0  int `json:"playerScore0"`
	PlayerScore1  int `json:"playerScore1"`
	PlayerScore2  int `json:"playerScore2"`
	PlayerScore3  int `json:"playerScore3"`
	PlayerScore4  int `json:"playerScore4"`
	PlayerScore5  int `json:"playerScore5"`
	PlayerScore6  int `json:"playerScore6"`
	PlayerScore7  int `json:"playerScore7"`
	PlayerScore8  int `json:"playerScore8"`
	PlayerScore9  int `json:"playerScore9"`
	PlayerScore10 int `json:"playerScore10"`
	PlayerScore11 int `json:"playerScore11"`
}

type PerksDto struct {
	StatPerks PerkStatsDto    `json:"statPerks"`
	Styles    []PerkStylesDto `json:"styles"`
}

type PerkStatsDto struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}

type PerkStylesDto struct {
	Description string                  `json:"description"`
	Selections  []PerkStyleSelectionDto `json:"selections"`
	Style       int                     `json:"style"`
}

type PerkStyleSelectionDto struct {
	Perk int `json:"perk"`
	Var1 int `json:"var1"`
	Var2 int `json:"var2"`
	Var3 int `json:"var3"`
}
