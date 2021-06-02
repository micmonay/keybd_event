package keybd_event

// Export downKey(), upKey() and sync() for strict reproducing of user input.

func (k *KeyBonding) Down(key uint16) error {
	return downKey(int(key))
}

func (k *KeyBonding) Up(key uint16) error {
	return upKey(int(key))
}

func (k *KeyBonding) Sync() error {
	return sync()
}
