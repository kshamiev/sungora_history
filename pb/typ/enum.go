// nolint
package typ

//
type Status string

const (
	Status_DRAFT  Status = "DRAFT"
	Status_NEW    Status = "NEW"
	Status_WORK   Status = "WORK"
	Status_CANCEL Status = "CANCEL"
	Status_CLOSE  Status = "CLOSE"
)

var StatusName = map[int32]Status{
	0: Status_DRAFT,
	1: Status_NEW,
	2: Status_WORK,
	3: Status_CANCEL,
	4: Status_CLOSE,
}

var StatusValue = map[Status]int32{
	Status_DRAFT:  0,
	Status_NEW:    1,
	Status_WORK:   2,
	Status_CANCEL: 3,
	Status_CLOSE:  4,
}

func (s Status) Valid() bool {
	_, ok := StatusValue[s]
	return ok
}
