package param

type Account struct {
	NewName        string `from:"new_username"`
	Birthday       string `from:"birthday"`
	Habitat        string `from:"habitat"`
	Hometown       string `from:"hometown"`
	HometownPublic string `from:"hometown_public"` //"1"为publilc "2"为privete
	BirthdayPublic string `form:"birthday_public"`
}
