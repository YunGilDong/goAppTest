package model

type Local struct {
	LOC_ID int
}

func (l *Local) TestSetId(id int) {
	l.LOC_ID = id
}
