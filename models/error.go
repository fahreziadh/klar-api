package models

type ErrorResponse struct {
	HttpStatusCode int    `bson:"httpStatusCode" json:"httpStatusCode"`
	Code           string `bson:"code" json:"code"`
	Message        string `bson:"message" json:"message"`
}

const (
	S001 = "E-mail Already Exsist"
	S002 = "Username Already Exsist"
	S003 = "Failed Register"
	S004 = "Credential Error"
	S005 = "Cant Find Username"
	S006 = "Wrong Password"
	S007 = "Failed Update"
	S008 = "Failed Create"
	S009 = "Not Found"
	S010 = "Wrong Date Format"
	S011 = "Failed when Input Schedule"
	S012 = "Failed to join Class"
)
