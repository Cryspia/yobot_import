/*Package utils is the package contains supporting functions and structures.
 *This file defines the go struct corresponding to the yobot JSON file
 */
package utils

//ClanChallenge is the struct that records each boss challenge
type ClanChallenge struct {
	BattleID         int    `json:"battle_id"`
	Behalf           int    `json:"behalf"`
	BossNum          int    `json:"boss_num"`
	ChallengePcrdate int    `json:"challenge_pcrdate"`
	ChallengePcrtime int    `json:"challenge_pcrtime"`
	ChallengeTime    int    `json:"challenge_time"`
	Cycle            int    `json:"cycle"`
	Damage           int    `json:"damage"`
	HealthRemain     int    `json:"health_ramain"` // It is a typo in original JSON format
	IsContinue       int    `json:"is_continue"`
	Message          string `json:"message"`
	QQID             int    `json:"qqid"`
}

//ClanGroup is the struct that record glide information
type ClanGroup struct {
	BattleID   int          `json:"battle_id"`
	GameServer string       `json:"game_server"`
	GroupID    int          `json:"group_id"`
	GroupName  string       `json:"group_name"`
	Members    []ClanMember `json:"members"`
}

//ClanMember is the struct that record playsers information
type ClanMember struct {
	Nickname string `json:"nickname"`
	QQID     int    `json:"qqid"`
}

//ClanGeneral is the struct that group all information
type ClanGeneral struct {
	APIVersion int             `json:"api_version"`
	Challenges []ClanChallenge `json:"challenges"`
	Code       int             `json:"code"`
	GroupInfo  []ClanGroup     `json:"groupinfo"`
	Members    []ClanMember    `json:"members"`
	Message    string          `json:"message"`
}
