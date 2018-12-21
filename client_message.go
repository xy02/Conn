package conn

//go:generate msgp

//Message is a data of client input or output
type Message struct {
	CID  int64  `msg:"cid"`
	Type string `msg:"type"`
	Body []byte `msg:"body"`
}

type FN struct {
	Name         string                 `msg:"fn"`
	Description  string                 `msg:"description"`
	ClientFilter map[string]interface{} `msg:"client_filter"`
}
