//This script will disable spells by removing every instance of them
package spells

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"github.com/xackery/eqcleanup/item"
	"github.com/xackery/eqcleanup/spawngroup"
)

var focus = "spells"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {

	//SpawnGroups
	ids := []int64{
		//Neriakc
		//cleric
		3162, //svlia in Neriakc
		3163, //Isshia
		3164, //trik
		3160, //lyniv
		3159, //myrish
		3165, //sol
		//wizard
		3142, //jusar
		3141, //misar
		5791, //drisi
		5792, //ash
		//SK
		3152,
		//Necro
		3154,
	}
	totalChanged, err := spawngroup.RemoveSpawnGroupAndEntryById(db, ids)
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}

	//Find all item IDS

	fmt.Println("This script takes a while...")
	rows, err := db.Query("SELECT id FROM items WHERE scrolleffect > 0")
	if err != nil {
		return
	}

	//iterate results
	for rows.Next() {
		id := int64(0)
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		ids = append(ids, id)
	}

	totalChanged, err := item.RemoveAllInstancesOfItems(db, ids)
	if err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in all player-accessible item locations.")
	return
}
