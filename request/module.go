package request

type Login struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	Validity             int    `json:"validity"`
	MacAddress           string `json:"macAddress"`
	DeviceName           string `json:"deviceName"`
	Encrypted            bool   `json:"encrypted"`
	SupportCustomEncrypt bool   `json:"supportCustomEncrypt"`
}

func NewLoginReq(user string, password string) Login {
	return Login{
		Username:             user,
		Password:             password,
		Validity:             -1,
		MacAddress:           "",
		DeviceName:           "",
		Encrypted:            true,
		SupportCustomEncrypt: false,
	}
}

type Report struct {
	Op              string `json:"op"`
	Cmd             string `json:"cmd"`
	EditCol         int    `json:"editcol"`
	EditRow         int    `json:"editrow"`
	EditReportIndex int    `json:"editReportIndex"`
	SessionID       string `json:"sessionID"`
	ReportXML       string `json:"reportXML"`
	Loadidxs        string `json:"loadidxs"`
}

func DefaultReportReq(sessionID string) Report {
	return Report{
		Op:              "fr_write",
		Cmd:             "cal_write_cell",
		EditCol:         3,
		EditRow:         17,
		EditReportIndex: 0,
		SessionID:       sessionID,
		ReportXML:       "%3C%3Fxml%20version%3D%221.0%22%20encoding%3D%22UTF-8%22%20%3F%3E%3CWorkBook%3E%3CVersion%3E6.5%3C%2FVersion%3E%3CReport%20class%3D%22com.fr.report.WorkSheet%22%20name%3D%220%22%3E%3CCellElementList%3E%3CC%20c%3D%223%22%20r%3D%2217%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B37%E5%BA%A6%E4%BB%A5%E4%B8%8B%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3C%2FCellElementList%3E%3C%2FReport%3E%3C%2FWorkBook%3E",
		Loadidxs:        "%5B0%5D",
	}
}

type ReportAll struct {
	Op        string `json:"op"`
	Cmd       string `json:"cmd"`
	ReportXML string `json:"reportXML"`
	/*
		op: fr_write
		cmd: submit_w_report
		reportXML: %3C%3Fxml%20version%3D%221.0%22%20encoding%3D%22UTF-8%22%20%3F%3E%3CWorkBook%3E%3CVersion%3E6.5%3C%2FVersion%3E%3CReport%20class%3D%22com.fr.report.WorkSheet%22%20name%3D%220%22%3E%3CCellElementList%3E%3CC%20c%3D%223%22%20r%3D%2216%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B37%E5%BA%A6%E4%BB%A5%E4%B8%8B%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%223%22%20r%3D%2217%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B37%E5%BA%A6%E4%BB%A5%E4%B8%8B%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2219%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E7%BB%BF%E8%89%B2%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2220%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2221%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2222%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2224%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2226%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2228%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2231%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3C%2FCellElementList%3E%3C%2FReport%3E%3C%2FWorkBook%3E
	*/
}

func DefaultReportAll() ReportAll {
	return ReportAll{
		Op:        "fr_write",
		Cmd:       "submit_w_report",
		ReportXML: "%3C%3Fxml%20version%3D%221.0%22%20encoding%3D%22UTF-8%22%20%3F%3E%3CWorkBook%3E%3CVersion%3E6.5%3C%2FVersion%3E%3CReport%20class%3D%22com.fr.report.WorkSheet%22%20name%3D%220%22%3E%3CCellElementList%3E%3CC%20c%3D%223%22%20r%3D%2216%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B37%E5%BA%A6%E4%BB%A5%E4%B8%8B%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%223%22%20r%3D%2217%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B37%E5%BA%A6%E4%BB%A5%E4%B8%8B%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2219%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E7%BB%BF%E8%89%B2%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2220%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2221%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2222%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2224%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2226%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2228%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3CC%20c%3D%225%22%20r%3D%2231%22%3E%3CO%20t%3D%22S%22%3E%3C!%5BCDATA%5B%E5%90%A6%5D%5D%3E%3C%2FO%3E%3C%2FC%3E%3C%2FCellElementList%3E%3C%2FReport%3E%3C%2FWorkBook%3E",
	}
}

func NewReportAll(reportxml string) ReportAll {
	return ReportAll{
		Op:        "fr_write",
		Cmd:       "submit_w_report",
		ReportXML: reportxml,
	}
}

type Config struct {
	Users     []User `json:"users"`
	ReportXML string `json:"reportXML"`
}

type User struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
