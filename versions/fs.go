package versions

type FieldSystem interface {
	Version() string
	Log() string
	Schedule() string
	Source() string
	SemLocked(string) (bool, error)
	Semaphores() []string
	Tracking() bool
	DataValid() bool
}
