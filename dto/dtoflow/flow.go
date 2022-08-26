package dtoflow

type FlowRow struct {
	Flow_id     string `json:"flow_id"`
	Flowname    string `json:"flow_name"`
	Description string `json:"desc"`
	Dtime       string `json:"lastre"`
}

/**
form
*/
type FormField struct {
	Field_id    string `json:"field_id"`
	Flow_id     string `json:"flow_id"`
	Fieldname   string `json:"field_name"`
	Fieldtype   string `json:"type"`
	Tablestruct string `json:"table"`
	Apistr      string `json:"api"`
	Onlyapi     string `json:"onlyapi"`
	Sort        int    `json:"sort"`
	Cmd_        string `json:"cmd_"`
}

//node
type Nodeinfo struct {
	Node_name string `json:"node_name"`
	Node_id   string `json:"node_id"`
	Vx        int    `json:"vx"`
	Vy        int    `json:"vy"`
	Timeout   int    `json:"timeout"`
	Flow_id   string `json:"flow_id"`
	Cmd_      string `json:"cmd_"`
}

type Btncheckinfo struct {
	Field_id string `json:"field_id"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	Btn_id   string `json:"btn_id"`
}
type Btnactinfo struct {
	Btn_id string `json:"btn_id"`
	Body   string `json:"body"`
}
type BtnPathinfo struct {
	Btn_id  string `json:"btn_id"`
	Node_id string `json:"node_id"`
	Flow_id string `json:"flow_id"`
	Cmd_    string `json:"cmd"`
}
type Btninfo struct {
	Btn_id   string `json:"btn_id"`
	Btn_name string `json:"btn_name"`
	Sort     int    `json:"sort"`
	Node_id  string `json:"node_id"`
	Cmd_     string `json:"cmd_"`
}
