package services

import "strconv"

func CreatEmailContentUrl(Owner_email string,team_id int,description string,email_accept string) (url string) {
	id := strconv.Itoa(team_id)
	return "https://api/accpetinvite/" + "Owner_email="+Owner_email+"&&team_id="+id+"email_accept="+email_accept
}
