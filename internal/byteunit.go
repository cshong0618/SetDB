package internal

type ByteUnit struct {
	next   map[byte]*ByteUnit
	b      byte
	exists bool
}

func InitByteUnit(b byte) *ByteUnit {
	return &ByteUnit{
		b:    b,
		next: make(map[byte]*ByteUnit),
	}
}

func InitAndSetByteUnit(b byte) *ByteUnit {
	byteUnit := InitByteUnit(b)
	byteUnit.Set()

	return byteUnit
}

func (bu *ByteUnit) Put(b byte) *ByteUnit {
	if nbu, ok := bu.next[b]; ok {
		return nbu
	} else {
		bu.next[b] = InitByteUnit(b)
		return bu.next[b]
	}
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
