package response

type Login struct {
	Data struct {
		Url               string      `json:"url"`
		Username          string      `json:"username"`
		Validity          int         `json:"validity"`
		Callback          interface{} `json:"callback"`
		AccessToken       string      `json:"accessToken"`
		OriginUrlResponse struct {
			OriginUrl  string `json:"originUrl"`
			Method     string `json:"method"`
			Parameters interface{}
		} `json:"originUrlResponse"`
	} `json:"data"`
}

type IndexData struct {
	Data []CellData `json:"data"`
}

type CellData struct {
	ID                      string `json:"id"`
	PID                     string `json:"pid"`
	Text                    string `json:"text"`
	Path                    string `json:"path"`
	Description             string `json:"description"`
	DeviceType              int    `json:"deviceType"`
	EntryType               int    `json:"entryType"`
	MobileNodeIcon          string `json:"mobileNodeIcon"`
	SortIndex               int    `json:"sortIndex"`
	IsParent                bool   `json:"isParent"`
	Open                    bool   `json:"open"`
	PrivilegeDetailBeanList string `json:"privilegeDetailBeanList"`
	Cover                   string `json:"cover"`
	Icon                    string `json:"icon"`
	FullParentName          string `json:"fullParentName"`
	ShowType                int    `json:"showType"`
	Parameters              []struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"parameters"`
}

type SessionDate struct {
	FitConfig struct {
		HorFit    int  `json:"horFit"`
		VerFit    int  `json:"verFit"`
		IsUseHTML bool `json:"isUseHtml"`
		Hash5     bool `json:"hash5"`
	} `json:"fitConfig"`
	ToolbarConfig struct {
		Zoom    bool `json:"zoom"`
		Refresh bool `json:"refresh"`
	} `json:"toolbarConfig"`
	IsUseHTML                 bool `json:"isUseHTML"`
	Hash5                     bool `json:"hash5"`
	IsMobileCanvasSize        bool `json:"isMobileCanvasSize"`
	AppearRefresh             bool `json:"appearRefresh"`
	AllowFullScreen           bool `json:"allowFullScreen"`
	AllowDoubleClickOrZoom    bool `json:"allowDoubleClickOrZoom"`
	FunctionalWhenUnactivated bool `json:"functionalWhenUnactivated"`
	ReportAttr                struct {
		Listeners []struct {
			EventName string `json:"eventName"`
			Once      bool   `json:"once"`
			Action    string `json:"action"`
		} `json:"listeners"`
		UnloadCheck     bool `json:"unloadCheck"`
		HasSubmitButton bool `json:"hasSubmitButton"`
	} `json:"reportAttr"`
	SessionID    string `json:"sessionid"`
	WebTitle     string `json:"webTitle"`
	SchShowType  int    `json:"schShowType"`
	Delay        bool   `json:"delay"`
	IsShowWindow bool   `json:"isShowWindow"`
}

type Result struct {
	FrSubmitinfo struct {
		Success bool `json:"success"`
	} `json:"fr_submitinfo"`
}
