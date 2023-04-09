package utility

import (
	"github.com/PretendoNetwork/badge-arcade-secure/database"
	nex "github.com/PretendoNetwork/nex-go"
)

func ChangeFreePlayDataMeta(dataID uint64, metaBinary []byte) {
	dateTime := nex.NewDateTime(0)
	updatedTime := dateTime.Now()

	database.UpdateFreePlayDataMetaBinary(dataID, metaBinary, updatedTime)
}