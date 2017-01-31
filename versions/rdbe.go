package versions

type Rdbe interface {
	RdbeNum() int
	RdbeMap(int) (map[string]interface{}, error)
	RdbeUpdatedFn(int) (func() bool, error)
}
