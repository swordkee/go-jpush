package jpush

type AdminRequest struct {
	AppName        string `json:"app_name,string"`
	AndroidPackage string `json:"android_package,string"`
	GroupName      string `json:"group_name,string"`
}
