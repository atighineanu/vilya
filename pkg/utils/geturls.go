package utils

import (
	obs "github.com/fgerling/gobs"
	"fmt"
	"strings"
)

func (config *VilyaCfg) CheckForUpd() ([]Updates, error) {
	var a string
	var updlist []Updates
	var update Updates
	client := obs.NewClient(config.OBS.User, config.OBS.Pass)
	rrs, err := client.GetReleaseRequests(config.MaintAPI.QATeam, "new,review")
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range rrs {
		a = a + "===\n"
		a = a + fmt.Sprintf("RR: %s\n", value.Id)
		update.ReleaseRequest = value.Id
		a = a + fmt.Sprintf("INCIDENT: %s\n", strings.Split(value.Actions[0].Target.Package, ".")[1])
		update.IncidentNumber = strings.Split(value.Actions[0].Target.Package, ".")[1]
		a = a + fmt.Sprintf("PACKAGE(s): \n")
		for _, val := range value.Actions {
			if !strings.Contains(val.Target.Package, "patchinfo") {
				update.SRCRPMS = append(update.SRCRPMS, strings.Split(val.Target.Package, ".")[0])
				a = a + fmt.Sprintf("             %s\n", strings.Split(val.Target.Package, ".")[0])
			}
		}
		update.Repository = obs.GetRepo(value)
		//update.AffectsCaaSP, err = CheckWichCaaSP(update.Repository)
		a = a + fmt.Sprintf("REPO: %s\n\n", update.Repository)
		updlist = append(updlist, update)
		update.SRCRPMS = []string{""}
	}
	fmt.Println(a)
	return updlist, nil
}