// Code generated by "stringer -type PersonState ."; DO NOT EDIT

package banana

import "fmt"

const _PersonState_name = "PersonHappyPersonHungryPersonGoingToStorePersonReturningFromStore"

var _PersonState_index = [...]uint8{0, 11, 23, 41, 65}

func (i PersonState) String() string {
	if i < 0 || i >= PersonState(len(_PersonState_index)-1) {
		return fmt.Sprintf("PersonState(%d)", i)
	}
	return _PersonState_name[_PersonState_index[i]:_PersonState_index[i+1]]
}
