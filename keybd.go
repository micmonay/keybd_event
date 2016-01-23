package keybd_event
type KeyBounding struct {
	hasCTRL bool
	hasALT bool
	hasSHIFT bool
	hasRCTRL bool
	hasRSHIFT bool
	keys[] int
}
/**
	Use for create struct KeyBounding
 */
func NewKeyBounding() (KeyBounding,error){
	keyBounding := KeyBounding{}
	keyBounding.hasALT=false
	keyBounding.hasCTRL=false
	keyBounding.hasSHIFT=false
	keyBounding.hasRCTRL=false
	keyBounding.hasRSHIFT=false
	keyBounding.keys = []int{}
	err := initKeyBD()
	if err!=nil {
		return keyBounding,err
	}
	return keyBounding,nil
}
func (k *KeyBounding) SetKeys(keys... int){
	k.keys = keys
}
/**
	If key ALT pressed
 */
func (k *KeyBounding) HasALT(b bool) {
	k.hasALT=b
}
/**
	If key CTRL pressed
 */
func (k *KeyBounding) HasCTRL(b bool) {
	k.hasCTRL=b
}
/**
	If key SHIFT pressed
 */
func (k *KeyBounding) HasSHIFT(b bool) {
	k.hasSHIFT=b
}
