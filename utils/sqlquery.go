/*Package utils is the package contains supporting functions and structures.
 *This file declare functions to help import from go struct to sqlite3 rows
 */
package utils

import (
	"database/sql"
	"fmt"
	"strings"

	// For sqlite3 sql plugin init
	_ "github.com/mattn/go-sqlite3"
)

const maxSQLVals = 10

//InsertClanChallenges insert from go struct 'ClanChallenge' to sqlite3 'clan_challenge' table
func InsertClanChallenges(SQLFile string, GID int, challenges []ClanChallenge) {
	db, err := sql.Open("sqlite3", SQLFile)
	if err != nil {
		fmt.Printf("sqlite file %s open failed\n", SQLFile)
		panic(err)
	}
	defer db.Close()
	ClanInsertQuery := "INSERT OR IGNORE INTO clan_challenge (gid, qqid, challenge_pcrdate, challenge_pcrtime, boss_cycle, boss_num, boss_health_ramain, challenge_damage, is_continue, message, behalf, bid) VALUES "
	InsertStr := ClanInsertQuery
	value := "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	InsertFormation := make([]string, 0)
	vals := make([]interface{}, 0)
	counter := 0
	for _, challenge := range challenges {
		if counter++; counter >= maxSQLVals {
			// sqlite3 plugin cannot support too many vals insertion at once
			counter = 0
			InsertStr += strings.Join(InsertFormation, ",")
			stm, err := db.Prepare(InsertStr)
			if err != nil {
				fmt.Println("FATAL: sqlite insertion statement incorrect")
				panic(err)
			}
			_, err = stm.Exec(vals...)
			if err != nil {
				fmt.Println("FATAL: sqlite insert values failed")
				panic(err)
			}
			stm.Close()
			InsertStr = ClanInsertQuery
			InsertFormation = make([]string, 0)
			vals = make([]interface{}, 0)
		}
		InsertFormation = append(InsertFormation, value)
		vals = append(vals,
			GID,
			challenge.QQID,
			challenge.ChallengePcrdate, challenge.ChallengePcrtime, challenge.Cycle,
			challenge.BossNum,
			challenge.HealthRemain,
			challenge.Damage,
			challenge.IsContinue,
			challenge.Message,
			challenge.Behalf,
			challenge.BattleID,
		)
	}
	InsertStr += strings.Join(InsertFormation, ",")
	if len(vals) > 0 {
		stm, err := db.Prepare(InsertStr)
		if err != nil {
			fmt.Println("FATAL: sqlite insertion statement incorrect")
			panic(err)
		}
		_, err = stm.Exec(vals...)
		stm.Close()
		if err != nil {
			fmt.Println("FATAL: sqlite insert values failed")
			panic(err)
		}
	}
}
