package quest

import (
	"fmt"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"os"
)

func Remove(config *eqemuconfig.Config, filePaths []string) (delCount int, err error) {
	if config.QuestsDir == "" {
		err = fmt.Errorf("Quests directory is not set.")
		return
	}
	for _, fileName := range filePaths {
		curFile := config.QuestsDir + "/" + fileName
		_, err = os.Stat(curFile)

		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			fmt.Printf("Error finding %s: %s", curFile, err.Error())
			continue
		}

		err = os.Remove(curFile)
		if err != nil {
			fmt.Printf("Error deleting %s: %s", curFile, err.Error())
			continue
		}
		delCount++
	}
	return
}