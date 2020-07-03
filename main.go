package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	utils "github.com/Cryspia/yobot_import/utils"
)

func main() {
	FromFlag := flag.String("f", "", "Source json file")
	ToFlag := flag.String("t", "", "Target sqlite3 db file")
	flag.Parse()

	SourceJSON := *FromFlag
	TargetDB := *ToFlag
	JSONFile, err := os.Open(SourceJSON)
	if err != nil {
		fmt.Printf("Import json file open err: %s\n", SourceJSON)
		panic(err)
	}
	fmt.Printf("Successfully open json file: %s\n", SourceJSON)
	defer JSONFile.Close()

	JSONBytes, err := ioutil.ReadAll(JSONFile)
	if err != nil {
		fmt.Printf("Read json file open err: %s\n", SourceJSON)
		panic(err)
	}
	fmt.Printf("Successfully read json file: %s\n", SourceJSON)

	var InData utils.ClanGeneral

	json.Unmarshal(JSONBytes, &InData)
	if len(InData.GroupInfo) == 0 {
		fmt.Println("Input JSON does not contain a single glide")
		os.Exit(1)
	}
	fmt.Printf("Trying to import %d records into clan_challenge...\n", len(InData.Challenges))
	utils.InsertClanChallenges(TargetDB, InData.GroupInfo[0].GroupID, InData.Challenges)
	fmt.Println("Finished!")
}
