package urlcheck

type Scenario []Test

func (s Scenario) Test() (err error) {
	for _, t := range s {
		err = t.Test()
		if err != nil {
			break
		}
	}
	return
}
