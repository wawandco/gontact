package core

//Contact represents the contact done throught our API,
//this struct will be helpful for our validation and so on.
type Contact struct {
	Name    string `valid:"required"`
	Email   string `valid:"email"`
	Address string
	Subject string
	Message string `valid:"required"`
	Website string
}
