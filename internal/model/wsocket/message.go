package wsocket

import "github.com/volatiletech/null"

type Message struct {
	Author string    `json:"author"`
	Body   null.JSON `json:"body"`
}
