// Package keybd_event is used for a key press simulated in Windows, Linux and Mac
package keybd_event

type KeyBonding struct {
	hasCTRL   bool
	hasALT    bool
	hasSHIFT  bool
	hasRCTRL  bool
	hasRSHIFT bool
	keys      []int
}

//Use for create struct KeyBounding
func NewKeyBonding() (KeyBonding, error) {
	keyBounding := KeyBonding{}
	keyBounding.hasALT = false
	keyBounding.hasCTRL = false
	keyBounding.hasSHIFT = false
	keyBounding.hasRCTRL = false
	keyBounding.hasRSHIFT = false
	keyBounding.keys = []int{}
	err := initKeyBD()
	if err != nil {
		return keyBounding, err
	}
	return keyBounding, nil
}
func (k *KeyBonding) SetKeys(keys ...int) {
	k.keys = keys
}

//If key ALT pressed
func (k *KeyBonding) HasALT(b bool) {
	k.hasALT = b
}

//If key CTRL pressed
func (k *KeyBonding) HasCTRL(b bool) {
	k.hasCTRL = b
}

//If key SHIFT pressed
func (k *KeyBonding) HasSHIFT(b bool) {
	k.hasSHIFT = b
}
