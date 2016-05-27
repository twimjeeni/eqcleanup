package era

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/item"
	"github.com/xackery/eqcleanup/quest"
	"github.com/xackery/eqcleanup/spawngroup"
	"github.com/xackery/eqemuconfig"
)

var focus = "era"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	//Mobs
	ids := []int64{
		111050, //Gellrazz Scalerunner
	}
	//#checkpoint_ten

	spawngroup.RemoveSpawnGroupAndEntryById(db, ids)

	//Items
	ids = []int64{
		84186,  //collector's pickaxe
		69240,  //Collector's Lightstone
		69241,  //Collector's Mistmoore Granite
		69242,  //Collector's Skunk Scent Gland
		69243,  //Collector's Fire Goblin Skin
		69244,  //Collector's Undead Froglok Tongue
		69245,  //Collector's Scythe
		69246,  //Collector's Water Ring
		69247,  //Collector's Kerran Doll
		69248,  //Collector's Preserved Split Paw Eye
		69249,  //Collector's Shark Tooth
		69250,  //Collector's Cheat Sheet
		84005,  //Collector's Fire Hornet Wing
		84006,  //Collector's Drachnid Web Sack
		84007,  //Collector's Iksar Witch Doll
		84008,  //Collector's Brittle Iksar Skull
		84009,  //Collector's Tump Stump
		84010,  //Collector's Sarnak War Braid
		84011,  //Collector's Bloodgill Scale
		84012,  //Collector's Nohope Moss
		84013,  //Collector's Excavator Claws
		84014,  //Collector's Canine
		84015,  //Collector's Cheat Sheet
		84171,  //Collector's Snow Bunny Foot
		84172,  //Collector's Cougar Tail
		84173,  //Collector's Shardwurm Scale
		84174,  //Collector's Wyvern Claw
		84175,  //Collector's Ice Sculpture
		84176,  //Collector's Brontotherium Hoof
		84177,  //Collector's Chetari Ceremonial Staff
		84178,  //Collector's Velium Trinket
		84179,  //Collector's Tizmak Horn
		84180,  //Collector's Kodiak Fang
		84181,  //Collector's Giant Sea Shell
		84182,  //Collector's Preserved Drake Wing
		84183,  //Collector's Bulthar Tongue
		84184,  //Collector's Sea Pearl
		84185,  //Collector's Lock of Mermaid Hair
		84186,  //Collector's Ry'Gorr Mining Pick
		84187,  //Collector's Terror Tentacle
		84188,  //Collector's Kromrif Signet
		84189,  //Collector's Holgresh Elder Bead
		84262,  //Collector's Crude Stone Idol
		84263,  //Collector's Sonic Wolf Ear
		84264,  //Collector's Grimling Toe
		84265,  //Collector's Elemental Focus
		84266,  //Collector's Polished Stone
		84267,  //Collector's Sun Revenant Veil
		84268,  //Collector's Lightcrawler Shell
		84269,  //Collector's Rockhopper Hide
		84270,  //Collector's Netherbian Claw
		84271,  //Collector's Fungal Fiend Flesh
		84275,  //Collector's Tribal Belt
		84276,  //Collector's Fish Hook
		84277,  //Collector's Ball of Mud
		84278,  //Collector's Scarlet Fish
		84279,  //Collector's Cht'Thk Horn
		84280,  //Collector's Shik`Nar Wing
		84296,  //Collector's Hobgoblin Tusk
		84297,  //Collector's Spider Eye
		84298,  //Collector's Pendant of Marr
		84299,  //Collector's Diaku Blade
		84300,  //Collector's Cloak of Justice
		84301,  //Collector's Frog Tongue
		84302,  //Collector's Tormentor Hide
		84303,  //Collector's Innovative Gear
		84304,  //Collector's Note
		84305,  //Collector's Soul
		110916, //Collector's Lava Protection Mask
		110931, //Collector's Sample Cloak

	}
	totalChanged, err := item.RemoveAllInstancesOfItems(db, ids)
	if err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in all player-accessible item locations.")

	filePaths := []string{
	//"befallen/Wraps_McGee.lua",
	}

	delCount, err := quest.Remove(config, filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}
