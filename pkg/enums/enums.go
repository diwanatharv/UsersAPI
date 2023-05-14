package enums

type Message string

var (
	Faileddecode    Message = "failed to decode the object"
	Validationerror Message = " validation failed"
	Successful      Message = "Succesfully updated "
	Servererror     Message = "Internal Server Error"
	Paramserror     Message = "Params send in wrong way"
)
