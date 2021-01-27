package intpacker

// NewUint32 packs two uint32 into one uint64
//
// Overflows of uint32 are not fully supported.
func NewUint32(x, y uint32) *Uint32 {
	packed := uint64(x)<<32 | uint64(y)
	return &Uint32{
		val: packed,
	}
}

// Uint32 type for easy unpacking and usage with sync.atomic
type Uint32 struct {
	val uint64
}

// Unpack the value into two uint32s
func (t *Uint32) Unpack() (uint32, uint32) {
	return uint32(t.val >> 32), uint32(t.val << 32 >> 32)
}

// Ptr returns a pointer for the underlying uint64 value.
func (t *Uint32) Ptr() *uint64 {
	return &t.val
}

// Uint64 returns the uint64 underlying value.
func (t *Uint32) Uint64() uint64 {
	return t.val
}
