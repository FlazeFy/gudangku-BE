package validations

import (
	"gudangku/modules/auth/models"
	"gudangku/packages/helpers/converter"
	"gudangku/packages/helpers/generator"
	"gudangku/packages/utils/validator"
)

func GetValidateRegister(body models.UserRegister) (bool, string) {
	var msg = ""
	var status = true

	// Rules
	minUname, maxUname := validator.GetValidationLength("username")
	minPass, maxPass := validator.GetValidationLength("password")
	minEmail, maxEmail := validator.GetValidationLength("email")
	minFName, maxFName := validator.GetValidationLength("first_name")
	_, maxLName := validator.GetValidationLength("last_name")

	// Value
	uname := converter.TotalChar(body.Username)
	pass := converter.TotalChar(body.Password)
	email := converter.TotalChar(body.Email)
	fname := converter.TotalChar(body.FirstName)
	lname := converter.TotalChar(body.LastName)

	// Validate
	if uname <= minUname || uname >= maxUname {
		status = false
		msg += generator.GenerateValidatorMsg("Username", minUname, maxUname)
	}
	if pass <= minPass || pass >= maxPass {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("Password", minPass, maxPass)
	}
	if email <= minEmail || email >= maxEmail {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("Email", minEmail, maxEmail)
	}
	if fname <= minFName || fname >= maxFName {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("First name", minFName, maxFName)
	}
	if lname >= maxLName {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("Last name", 0, maxFName)
	}

	if status {
		return status, "Validation success"
	} else {
		return status, msg
	}
}

func GetValidateLogin(username, password string) (bool, string) {
	var msg = ""
	var status = true

	// Rules
	minUname, maxUname := validator.GetValidationLength("username")
	minPass, maxPass := validator.GetValidationLength("password")

	// Value
	uname := converter.TotalChar(username)
	pass := converter.TotalChar(password)

	// Validate
	if uname <= minUname || uname >= maxUname {
		status = false
		msg += generator.GenerateValidatorMsg("Username", minUname, maxUname)
	}
	if pass <= minPass || pass >= maxPass {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("Password", minPass, maxPass)
	}

	if status {
		return status, "Validation success"
	} else {
		return status, msg
	}
}
