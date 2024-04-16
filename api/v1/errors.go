package v1

var (
	// common errors
	ErrSuccess             = newError(0, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// more biz errors
	ErrEmailAlreadyUse      = newError(1001, "The email is already in use.")
	ErrNoAccessDeleteFloder = newError(1002, "The floder is not belong you. Please  check.")
	ErrNoAccessUpdateFloder = newError(1003, "The floder is not belong you. Please  check.")
	ErrNoAccessAddDeck      = newError(1004, "The floder is not belong you. You can't add deck to it.")
	ErrNoAccessGetDecks     = newError(1005, "This deck is not belong you")
	ErrNoAccessDeleteDeck   = newError(1006, "The deck is not belong you. Please  check.")
	ErrUsernameAlreadyUse   = newError(1007, "The username is already in use.")
	ErrPrivileges           = newError(1008, "your Privileges is not enough ,can't  finish op")
)
