package room

type Store interface { SetOpen(string) error; IsOpen() bool; Validate(string) bool; Close() error }
