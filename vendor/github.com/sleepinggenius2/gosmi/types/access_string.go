// Code generated by "enumer -type=Access -autotrimprefix -json"; DO NOT EDIT

package types

import (
	"encoding/json"
	"fmt"
)

const _Access_name = "UnknownNotImplementedNotAccessibleNotifyReadOnlyReadWriteInstallInstallNotifyReportOnlyEventOnly"

var _Access_index = [...]uint8{0, 7, 21, 34, 40, 48, 57, 64, 77, 87, 96}

func (i Access) String() string {
	if i < 0 || i >= Access(len(_Access_index)-1) {
		return fmt.Sprintf("Access(%d)", i)
	}
	return _Access_name[_Access_index[i]:_Access_index[i+1]]
}

var _AccessNameToValue_map = map[string]Access{
	_Access_name[0:7]:   0,
	_Access_name[7:21]:  1,
	_Access_name[21:34]: 2,
	_Access_name[34:40]: 3,
	_Access_name[40:48]: 4,
	_Access_name[48:57]: 5,
	_Access_name[57:64]: 6,
	_Access_name[64:77]: 7,
	_Access_name[77:87]: 8,
	_Access_name[87:96]: 9,
}

func AccessFromString(s string) (Access, error) {
	if val, ok := _AccessNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Access values", s)
}

func AccessAsList() []Access {
	list := make([]Access, len(_AccessNameToValue_map))
	idx := 0
	for _, v := range _AccessNameToValue_map {
		list[idx] = v
		idx++
	}
	return list
}

func AccessAsListString() []string {
	list := make([]string, len(_AccessNameToValue_map))
	idx := 0
	for k := range _AccessNameToValue_map {
		list[idx] = k
		idx++
	}
	return list
}

func AccessIsValid(t Access) bool {
	for _, v := range AccessAsList() {
		if t == v {
			return true
		}
	}
	return false
}

func (i Access) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *Access) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Access should be a string, got %s", data)
	}

	var err error
	*i, err = AccessFromString(s)
	return err
}
