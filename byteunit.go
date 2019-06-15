package setdb

type ByteUnit struct {
	b byte
	next map[byte]*ByteUnit
	exists bool
}

func InitByteUnit(b byte) *ByteUnit {
	bu := ByteUnit{}
	bu.b = b
	bu.next = make(map[byte]*ByteUnit)

	return &bu
}

func InitAndSetByteUnit(b byte) *ByteUnit {
	byteUnit := InitByteUnit(b)
	byteUnit.Set()

	return byteUnit
}

func (bu ByteUnit) Find(b byte) *ByteUnit {
	if next, ok := bu.next[b]; ok {
		return next
	} else {
		return nil
	}
}

func (bu ByteUnit) FindByte(iBu ByteUnit) *ByteUnit {
	return bu.Find(iBu.b)
}

func (bu *ByteUnit) Set() {
	bu.exists = true
}

func (bu *ByteUnit) Unset() {
	bu.exists = false
}