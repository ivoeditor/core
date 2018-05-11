package core

type Window interface {
	Close(Context)
	Command(Context, Command)
	Key(Context, Key)
	Mouse(Context, Mouse)
	Resize(Context)
}
