package versions

const (
	basepath = "/usr2/fs"
)

type Creator func() (FieldSystem, error)

var Creators = map[string]Creator{}

func Add(name string, creator Creator) {
	Creators[name] = creator
}
