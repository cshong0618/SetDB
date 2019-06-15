package internals

const (
	OUTER_LEN = 4
	INNER_LEN = 64
)

type InnerArray struct {
	_b []*ByteUnit
}

type ByteUnit struct {
	//next map[byte]*ByteUnit
	_innerArrays []*InnerArray
	//b byte
	exists bool

}

func InitInnerArray() *InnerArray {
	ia := InnerArray{}
	ia._b = make([]*ByteUnit, INNER_LEN)

	return &ia
}

func InitByteUnit(b byte) *ByteUnit {
	bu := ByteUnit{}
	//bu.b = b
	//bu.next = make(map[byte]*ByteUnit)

	bu._innerArrays = make([]*InnerArray, OUTER_LEN)

	return &bu
}

func InitAndSetByteUnit(b byte) *ByteUnit {
	byteUnit := InitByteUnit(b)
	byteUnit.Set()

	return byteUnit
}

func (bu ByteUnit) Find(b byte) *ByteUnit {
	idx, innerIdx := getIdxAndInnerIdx(b)

	if ia := bu._innerArrays[idx]; ia != nil {
		if ib := ia._b[innerIdx]; ib != nil {
			return ib
		}
	}

	return nil
}

func (bu *ByteUnit) Put(b byte) *ByteUnit {
	idx, innerIdx := getIdxAndInnerIdx(b)

	innerBytes := bu._innerArrays[idx]
	if innerBytes == nil {
		innerBytes = InitInnerArray()
		bu._innerArrays[idx] = innerBytes
	}

	if innerBytes._b[innerIdx] == nil {
		innerBytes._b[innerIdx] = InitByteUnit(b)
	}

	return innerBytes._b[innerIdx]
}

func getIdxAndInnerIdx(b byte) (idx byte, innerIdx byte) {
	idx = b % byte(OUTER_LEN)
	innerIdx = idx % byte(INNER_LEN)

	return
}

func (bu *ByteUnit) Set() {
	bu.exists = true
}

func (bu *ByteUnit) Unset() {
	bu.exists = false
}