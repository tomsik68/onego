package resources

import (
	"fmt"
	"github.com/beevik/etree"
	"strconv"
)

// DataStore struct
type DataStore struct {
	XMLData *etree.Element
}

// CreateDataStore constructs DataStore
func CreateDataStore(id int) *DataStore {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0"`)

	el := doc.CreateElement("DATASTORE").CreateElement("ID")
	el.SetText(fmt.Sprintf("%d", id))

	return &DataStore{doc.Root()}
}

// Attribute method
func (d DataStore) Attribute(path string) string {
	elements := d.XMLData.FindElements(path)
	if elements == nil {
		return ""
	}
	return elements[0].Text()
}

// Name method
func (d DataStore) Name() string {
	return d.Attribute("NAME")
}

// Clusters method
func (d DataStore) Clusters() []Cluster {
	elements := d.XMLData.FindElements("CLUSTERS")
	clusters := make([]Cluster, len(elements))
	for i, e := range elements {
		clusters[i] = Cluster{XMLData: e}
	}
	return clusters
}

// ID method
func (d DataStore) ID() int {
	i, err := strconv.Atoi(d.Attribute("ID"))
	if err != nil {
		return -1
	}
	return i
}
