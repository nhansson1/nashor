package storage

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"nashor/internal/types"
	"os"
)

type MatchDB struct {
	MatchId            string `db:"match_id"`
	GameDuration       int    `db:"game_duration"`
	GameEndTimestamp   int    `db:"game_end_timestamp"`
	GameStartTimestamp int    `db:"game_start_timestamp"`
	QueueId            int    `db:"queue_id"`
}

type ParticipantDB struct {
	ChampionName         string         `db:"champion_name"`
	ChampionId           int            `db:"champion_id"`
	Puuid                string         `db:"puuid"`
	Kills                int            `db:"kills"`
	Deaths               int            `db:"deaths"`
	Assists              int            `db:"assists"`
	NeutralMinionsKilled int            `db:"neutral_minions_killed"`
	TotalMinionsKilled   int            `db:"total_minions_killed"`
	GoldEarned           int            `db:"gold_earned"`
	VisionScore          int            `db:"vision_score"`
	RiotIdGameName       string         `db:"riot_id_game_name"`
	RiotIdTagline        string         `db:"riot_id_tagline"`
	SummonerName         string         `db:"summoner_name"`
	Win                  bool           `db:"win"`
	Item0                int            `db:"item_0"`
	Item1                int            `db:"item_1"`
	Item2                int            `db:"item_2"`
	Item3                int            `db:"item_3"`
	Item4                int            `db:"item_4"`
	Item5                int            `db:"item_5"`
	Item6                int            `db:"item_6"`
	Summoner1Id          int            `db:"summoner_1_id"`
	Summoner2Id          int            `db:"summoner_2_id"`
	Perks                types.PerksDto `db:"perk_data"`
	MatchId              string         `db:"match_id"`
	PerksId              int            `db:"perks_id"`
	ParticipantId        int            `db:"participant_id"`
}

type PostgresClient struct {
	db *sqlx.DB
}

func (c *PostgresClient) createPerk(p types.PerksDto) (int, error) {
	perkId := 0
	q := "INSERT INTO perks (perk_data) VALUES ($1) RETURNING perks_id"

	b, err := json.Marshal(p)

	if err != nil {
		return perkId, err
	}

	err = c.db.QueryRow(q, b).Scan(&perkId)

	if err != nil {
		return perkId, err
	}

	return perkId, nil
}

func GetMatchDto(m MatchDB, p []ParticipantDB) types.MatchDto {
	var participants []types.ParticipantDto

	for _, p := range p {
		participants = append(participants, types.ParticipantDto{
			ChampionName:         p.ChampionName,
			ChampionId:           p.ChampionId,
			Puuid:                p.Puuid,
			Kills:                p.Kills,
			Deaths:               p.Deaths,
			Assists:              p.Assists,
			Item0:                p.Item0,
			Item1:                p.Item1,
			Item2:                p.Item2,
			Item3:                p.Item3,
			Item4:                p.Item4,
			Item5:                p.Item5,
			Item6:                p.Item6,
			Perks:                p.Perks,
			NeutralMinionsKilled: p.NeutralMinionsKilled,
			TotalMinionsKilled:   p.TotalMinionsKilled,
			GoldEarned:           p.GoldEarned,
			VisionScore:          p.VisionScore,
			RiotIdGameName:       p.RiotIdGameName,
			RiotIdTagline:        p.RiotIdTagline,
			SummonerName:         p.SummonerName,
			Win:                  p.Win,
			Summoner1Id:          p.Summoner1Id,
			Summoner2Id:          p.Summoner2Id,
		})
	}

	return types.MatchDto{
		Info: types.InfoDto{
			GameStartTimestamp: m.GameStartTimestamp,
			GameEndTimestamp:   m.GameEndTimestamp,
			GameDuration:       m.GameDuration,
			Participants:       participants,
			QueueId:            m.QueueId,
		},
		Metadata: types.MetadataDto{
			MatchId: m.MatchId,
		},
	}
}

func (c *PostgresClient) createParticipant(p types.ParticipantDto, matchId string, perksId int) {
	q := "INSERT INTO participants (champion_name, champion_id, puuid, kills, deaths, assists, neutral_minions_killed, total_minions_killed, gold_earned, vision_score, riot_id_game_name, riot_id_tagline, summoner_name, win, item_0, item_1, item_2, item_3, item_4, item_5, item_6, summoner_1_id, summoner_2_id, match_id, perks_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)"

	c.db.QueryRow(q, p.ChampionName, p.ChampionId, p.Puuid, p.Kills, p.Deaths, p.Assists, p.NeutralMinionsKilled, p.TotalMinionsKilled, p.GoldEarned, p.VisionScore, p.RiotIdGameName, p.RiotIdTagline, p.SummonerName, p.Win, p.Item0, p.Item1, p.Item2, p.Item3, p.Item4, p.Item5, p.Item6, p.Summoner1Id, p.Summoner2Id, matchId, perksId)
}

func (c *PostgresClient) CreateMatch(m types.MatchDto) {
	q := "INSERT INTO matches (match_id, game_duration, game_end_timestamp, game_start_timestamp, queue_id) VALUES ($1, $2, $3, $4, $5) RETURNING match_id"

	var matchId string
	err := c.db.QueryRow(q, m.Metadata.MatchId, m.Info.GameDuration, m.Info.GameEndTimestamp, m.Info.GameStartTimestamp, m.Info.QueueId).Scan(&matchId)

	if err != nil {
		fmt.Println("failed to create match in db with match id", matchId, err)
		return
	}

	for _, p := range m.Info.Participants {
		perksId, err := c.createPerk(p.Perks)

		if err != nil {
			fmt.Println("failed to create perk err: ", err)
			return
		}

		c.createParticipant(p, matchId, perksId)
	}
}

func (c *PostgresClient) GetMatchById(id string) (types.MatchDto, error) {
	m := MatchDB{}
	p := []ParticipantDB{}

	var out types.MatchDto

	q := "SELECT * FROM matches WHERE matches.match_id = $1"

	err := c.db.Get(&m, q, id)

	if err != nil {
		return out, err
	}

	q = "SELECT * FROM participants JOIN perks ON perks.perks_id = participants.perks_id WHERE participants.match_id = $1"

	err = c.db.Select(&p, q, m.MatchId)

	if err != nil {
		return out, err
	}

	out = GetMatchDto(m, p)

	return out, nil
}

func NewPostgresClient() *PostgresClient {
	connStr := os.Getenv("DB_URI")

	db := sqlx.MustOpen("postgres", connStr)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	return &PostgresClient{
		db: db,
	}
}
