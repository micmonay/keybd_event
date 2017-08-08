package keybd_event

/*
 #cgo CFLAGS: -x objective-c
 #cgo LDFLAGS: -framework Cocoa
 #import <Foundation/Foundation.h>
 CGEventRef Create(int k){
	CGEventRef event = CGEventCreateKeyboardEvent (NULL, (CGKeyCode)k, true);
	return event;
 }
 void KeyTap(CGEventRef event){
	CGEventPost(kCGAnnotatedSessionEventTap, event);
	CFRelease(event);
 }
 void AddActionKey(CGEventFlags type,CGEventRef event){
 	CGEventSetFlags(event, type);
 }
*/
import "C"
import (
	"errors"
	"time"
)

const (
	_AShift          = C.kCGEventFlagMaskAlphaShift
	_VK_SHIFT        = C.kCGEventFlagMaskShift
	_VK_CTRL         = C.kCGEventFlagMaskControl
	_VK_ALT          = C.kCGEventFlagMaskAlternate
	_VK_CMD          = C.kCGEventFlagMaskCommand
	_Help            = C.kCGEventFlagMaskHelp
	_VK_FN           = C.kCGEventFlagMaskSecondaryFn
	_NumPad          = C.kCGEventFlagMaskNumericPad
	_Coalesced       = C.kCGEventFlagMaskNonCoalesced
	_VK_Control      = 0x3B
	_VK_RightShift   = 0x3C
	_VK_RightControl = 0x3E
	_VK_Command      = 0x37
	_VK_Shift        = 0x38
)

func initKeyBD() error { return nil }

// Launch key bounding
func (k *KeyBonding) Launching() error {

	for _, key := range k.keys {
		k.tapKey(key)
	}
	return nil
}

// This function presses keys in the arguments one by one.
func (k *KeyBonding) LaunchString(str string) {
	// initialization
	k.HasSHIFT(false)

	// for every character of the input string
	for _, char := range str {

		// check whether a character needs the shift key
		if k.isNeedShift(char) {
			k.HasSHIFT(true)
		} else {
			k.HasSHIFT(false)
		}

		char = k.convertToNoShift(char)
		key, err := k.characterToKey(char)
		if err != nil {
			panic(err)
		}

		// Press the key
		k.SetKeys(key)
		err = k.Launching()
		if err != nil {
			panic(err)
		}
	}

	// reinitialization
	k.HasSHIFT(false)
}

func shift(event C.CGEventRef) {
	C.AddActionKey(_VK_SHIFT, event)
}
func ctrl(event C.CGEventRef) {
	C.AddActionKey(_VK_CTRL, event)
}
func alt(event C.CGEventRef) {
	C.AddActionKey(_VK_ALT, event)
}
func cmd(event C.CGEventRef) {
	C.AddActionKey(_VK_CMD, event)
}
func (k KeyBonding) tapKey(key int) {
	event := C.Create(C.int(key))
	if k.hasALT {
		alt(event)
	}
	if k.hasCTRL {
		ctrl(event)
	}
	if k.hasSHIFT {
		shift(event)
	}
	if k.hasRCTRL { //not support on mac
		ctrl(event)
	}
	if k.hasRSHIFT { //not support on mac
		shift(event)
	}
	C.KeyTap(event)
	time.Sleep(100 * time.Millisecond) //ignore if speed is most in my test system
}

const (
	VK_A              = 0x00
	VK_S              = 0x01
	VK_D              = 0x02
	VK_F              = 0x03
	VK_H              = 0x04
	VK_G              = 0x05
	VK_Z              = 0x06
	VK_X              = 0x07
	VK_C              = 0x08
	VK_V              = 0x09
	VK_B              = 0x0B
	VK_Q              = 0x0C
	VK_W              = 0x0D
	VK_E              = 0x0E
	VK_R              = 0x0F
	VK_Y              = 0x10
	VK_T              = 0x11
	VK_1              = 0x12
	VK_2              = 0x13
	VK_3              = 0x14
	VK_4              = 0x15
	VK_6              = 0x16
	VK_5              = 0x17
	VK_EQUAL          = 0x18
	VK_9              = 0x19
	VK_7              = 0x1A
	VK_MINUS          = 0x1B
	VK_8              = 0x1C
	VK_0              = 0x1D
	VK_RightBracket   = 0x1E
	VK_O              = 0x1F
	VK_U              = 0x20
	VK_LeftBracket    = 0x21
	VK_I              = 0x22
	VK_P              = 0x23
	VK_L              = 0x25
	VK_J              = 0x26
	VK_Quote          = 0x27
	VK_K              = 0x28
	VK_SEMICOLON      = 0x29
	VK_BACKSLASH      = 0x2A
	VK_COMMA          = 0x2B
	VK_SLASH          = 0x2C
	VK_N              = 0x2D
	VK_M              = 0x2E
	VK_Period         = 0x2F
	VK_GRAVE          = 0x32
	VK_KeypadDecimal  = 0x41
	VK_KeypadMultiply = 0x43
	VK_KeypadPlus     = 0x45
	VK_KeypadClear    = 0x47
	VK_KeypadDivide   = 0x4B
	VK_KeypadEnter    = 0x4C
	VK_KeypadMinus    = 0x4E
	VK_KeypadEquals   = 0x51
	VK_Keypad0        = 0x52
	VK_Keypad1        = 0x53
	VK_Keypad2        = 0x54
	VK_Keypad3        = 0x55
	VK_Keypad4        = 0x56
	VK_Keypad5        = 0x57
	VK_Keypad6        = 0x58
	VK_Keypad7        = 0x59
	VK_Keypad8        = 0x5B
	VK_Keypad9        = 0x5C

	VK_ENTER         = 0x24
	VK_TAB           = 0x30
	VK_SPACE         = 0x31
	VK_DELETE        = 0x33
	VK_ESC           = 0x35
	VK_CAPSLOCK      = 0x39
	VK_Option        = 0x3A
	VK_RightOption   = 0x3D
	VK_Function      = 0x3F
	VK_F17           = 0x40
	VK_VOLUMEUP      = 0x48
	VK_VOLUMEDOWN    = 0x49
	VK_MUTE          = 0x4A
	VK_F18           = 0x4F
	VK_F19           = 0x50
	VK_F20           = 0x5A
	VK_F5            = 0x60
	VK_F6            = 0x61
	VK_F7            = 0x62
	VK_F3            = 0x63
	VK_F8            = 0x64
	VK_F9            = 0x65
	VK_F11           = 0x67
	VK_F13           = 0x69
	VK_F16           = 0x6A
	VK_F14           = 0x6B
	VK_F10           = 0x6D
	VK_F12           = 0x6F
	VK_F15           = 0x71
	VK_HELP          = 0x72
	VK_HOME          = 0x73
	VK_PAGEUP        = 0x74
	VK_ForwardDelete = 0x75
	VK_F4            = 0x76
	VK_END           = 0x77
	VK_F2            = 0x78
	VK_PAGEDOWN      = 0x79
	VK_F1            = 0x7A
	VK_LEFT          = 0x7B
	VK_RIGHT         = 0x7C
	VK_DOWN          = 0x7D
	VK_UP            = 0x7E
)

// This function converts a shift-need character to a shft-not-need character
func (k *KeyBonding) convertToNoShift(c rune) rune {
	switch c {
	case 'A':
		return 'a'
	case 'S':
		return 's'
	case 'D':
		return 'd'
	case 'F':
		return 'f'
	case 'H':
		return 'h'
	case 'G':
		return 'g'
	case 'Z':
		return 'z'
	case 'X':
		return 'x'
	case 'C':
		return 'c'
	case 'V':
		return 'v'
	case 'B':
		return 'b'
	case 'Q':
		return 'q'
	case 'W':
		return 'w'
	case 'E':
		return 'e'
	case 'R':
		return 'r'
	case 'Y':
		return 'y'
	case 'T':
		return 't'
	case '!':
		return '1'
	case '@':
		return '2'
	case '#':
		return '3'
	case '$':
		return '4'
	case '^':
		return '6'
	case '%':
		return '5'
	case '+':
		return '='
	case '(':
		return '9'
	case '&':
		return '7'
	case '_':
		return '-'
	case '*':
		return '8'
	case ')':
		return '0'
	case '}':
		return ']'
	case 'O':
		return 'o'
	case 'U':
		return 'u'
	case '{':
		return '['
	case 'I':
		return 'i'
	case 'P':
		return 'p'
	case 'L':
		return 'l'
	case 'J':
		return 'j'
	case '"':
		return '\''
	case 'K':
		return 'k'
	case ':':
		return ';'
	case '|':
		return '\\'
	case '<':
		return ','
	case '?':
		return '/'
	case 'N':
		return 'n'
	case 'M':
		return 'm'
	case '>':
		return '.'
	case '~':
		return '`'

	default:
		// if the argument does not need convertion, just return it.
		return c
	}

}

// This function checks whether the argument is need the shift key.
func (k *KeyBonding) isNeedShift(c rune) bool {
	switch c {
	case 'A':
		return true
	case 'S':
		return true
	case 'D':
		return true
	case 'F':
		return true
	case 'H':
		return true
	case 'G':
		return true
	case 'Z':
		return true
	case 'X':
		return true
	case 'C':
		return true
	case 'V':
		return true
	case 'B':
		return true
	case 'Q':
		return true
	case 'W':
		return true
	case 'E':
		return true
	case 'R':
		return true
	case 'Y':
		return true
	case 'T':
		return true
	case '!':
		return true
	case '@':
		return true
	case '#':
		return true
	case '$':
		return true
	case '^':
		return true
	case '%':
		return true
	case '+':
		return true
	case '(':
		return true
	case '&':
		return true
	case '_':
		return true
	case '*':
		return true
	case ')':
		return true
	case '}':
		return true
	case 'O':
		return true
	case 'U':
		return true
	case '{':
		return true
	case 'I':
		return true
	case 'P':
		return true
	case 'L':
		return true
	case 'J':
		return true
	case '"':
		return true
	case 'K':
		return true
	case ':':
		return true
	case '|':
		return true
	case '<':
		return true
	case '?':
		return true
	case 'N':
		return true
	case 'M':
		return true
	case '>':
		return true
	case '~':
		return true

	default:
		return false
	}

}

// This function converters a character to a key number. The argument should be a character that does not need the shift key.
func (k *KeyBonding) characterToKey(c rune) (int, error) {
	switch c {
	case 'a':
		return VK_A, nil
	case 's':
		return VK_S, nil
	case 'd':
		return VK_D, nil
	case 'f':
		return VK_F, nil
	case 'h':
		return VK_H, nil
	case 'g':
		return VK_G, nil
	case 'z':
		return VK_Z, nil
	case 'x':
		return VK_X, nil
	case 'c':
		return VK_C, nil
	case 'v':
		return VK_V, nil
	case 'b':
		return VK_B, nil
	case 'q':
		return VK_Q, nil
	case 'w':
		return VK_W, nil
	case 'e':
		return VK_E, nil
	case 'r':
		return VK_R, nil
	case 'y':
		return VK_Y, nil
	case 't':
		return VK_T, nil
	case '1':
		return VK_1, nil
	case '2':
		return VK_2, nil
	case '3':
		return VK_3, nil
	case '4':
		return VK_4, nil
	case '6':
		return VK_6, nil
	case '5':
		return VK_5, nil
	case '=':
		return VK_EQUAL, nil
	case '9':
		return VK_9, nil
	case '7':
		return VK_7, nil
	case '-':
		return VK_MINUS, nil
	case '8':
		return VK_8, nil
	case '0':
		return VK_0, nil
	case ']':
		return VK_RightBracket, nil
	case 'o':
		return VK_O, nil
	case 'u':
		return VK_U, nil
	case '[':
		return VK_LeftBracket, nil
	case 'i':
		return VK_I, nil
	case 'p':
		return VK_P, nil
	case 'l':
		return VK_L, nil
	case 'j':
		return VK_J, nil
	case '\'':
		return VK_Quote, nil
	case 'k':
		return VK_K, nil
	case ';':
		return VK_SEMICOLON, nil
	case '\\':
		return VK_BACKSLASH, nil
	case ',':
		return VK_COMMA, nil
	case '/':
		return VK_SLASH, nil
	case 'n':
		return VK_N, nil
	case 'm':
		return VK_M, nil
	case '.':
		return VK_Period, nil
	case '`':
		return VK_GRAVE, nil

	default:
		//error case
	}

	// here is error handling
	errMsg := errors.New("keybd_event: " + "'" + string(c) + "' is need shift key or a functional key")
	return 0, errMsg
}
