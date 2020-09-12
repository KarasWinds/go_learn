package monster

import (
	"encoding/json"
	"io/ioutil"
)

// Monster is a class
type Monster struct {
	Name  string
	Age   int
	Skill string
}

// Store is save the class to file
func (m *Monster) Store() error {
	data, err := json.Marshal(&m)
	if err != nil {
		return err
	}
	filePath := "./class"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

// ReStore is load the class from file
func (m *Monster) ReStore() error {
	filePath := "./class"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	return nil
}

func main() {

}
