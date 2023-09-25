package registry

import (
	"fmt"
	"github.com/piusalfred/registry/pkg/errors"
	"regexp"
	"time"
)

var (
	ErrInvalidMacAddress   = errors.New("invalid mac address")
	ErrGeneratingNodeToken = errors.New("error generating new node token")
)

type NodeStatus int

const (
	Revoked NodeStatus = iota + 1
	AllowedOffline
	AllowedOnline
)

func (ns NodeStatus) String() string {
	switch ns {
	case Revoked:
		return "revoked"

	case AllowedOffline:
		return "allowed-offline"

	case AllowedOnline:
		return "allowed-online"

	default:
		return "unknown status"
	}

}

type Type int

const (
	Sensor Type = iota + 1
	Actuator
	Controller
)

func (t Type) String() string {
	switch t {
	case Sensor:
		return "sensor"
	case Actuator:
		return "actuator"
	case Controller:
		return "controller"
	default:
		return "unrecognized node type"
	}
}

var macAddrRegex = regexp.MustCompile("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})|([0-9a-fA-F]{4}\\\\.[0-9a-fA-F]{4}\\\\.[0-9a-fA-F]{4})$")

type NodesTree struct {
	NodeID     string   `json:"node_id"`
	ChildNodes []string `json:"child_nodes"`
}

//Node represent the edge node deployed in the igrid network
type Node struct {
	UUID    string `json:"uuid"`
	Addr    string `json:"addr"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	Region  string `json:"region"`
	Latd    string `json:"latitude"`
	Long    string `json:"longitude"`
	Created string `json:"created"`
	Master  string `json:"master,omitempty"`
}

func CreateNode(provider UUIDProvider, addr, name, region,
	lat, long, master string, typ int) (node Node, err error) {
	UUID, err := provider.ID()

	if err != nil {
		return node, err
	}

	if !validateAddr(addr) {
		err = ErrInvalidMacAddress
		return node, err
	}

	now := time.Now().Format(time.RFC3339)

	node = Node{
		UUID:    UUID,
		Addr:    addr,
		Name:    name,
		Type:    typ,
		Region:  region,
		Latd:    lat,
		Long:    long,
		Created: now,
		Master:  master,
	}

	return node, nil
}

func (n Node) String() string {
	nodeStr := fmt.Sprintf("node = [id = %s, addr = %s, name= %s, type =%s , region = %s, lat = %s, long =%s, created = %s, master = %s]",
		n.UUID, n.Addr, n.Name, Type(n.Type), n.Region, n.Latd, n.Long, n.Created, n.Master)

	return nodeStr
}

/*func (n Node) String() string {

	type PrintNode struct {
		Addr      string `json:"addr"`
		Name      string `json:"name"`
		Type      string `json:"type"`
		Status    string `json:"status"`
		Latitude  string `json:"lat"`
		Longitude string `json:"long"`
		Token     string `json:"token"`
		CreatedAt string `json:"created_at"`
	}

	//status := fmt.Sprintf(strings.Tn.Status.String())
	//nodeTyp := fmt.Sprintf(n.Type.String())

	t := time.Unix(n.CreatedAt, 0)
	t.Format(time.RFC3339)

	pn := PrintNode{
		Addr:      n.Addr,
		Name:      n.Name,
		Type:      nodeTyp,
		Status:    status,
		Latitude:  n.Latitude,
		Longitude: n.Longitude,
		Token:     n.Token,
		CreatedAt: t.String(),
	}

	empJSON, err := json.MarshalIndent(pn, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return fmt.Sprintf("node \n %s\n", string(empJSON))

}
*/
func validateAddr(macAddr string) bool {

	//regular expression for mac address validation
	//should match 00:1B:44:11:3A:B7
	//regex := "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})|([0-9a-fA-F]{4}\\.[0-9a-fA-F]{4}\\.[0-9a-fA-F]{4})$"

	if len(macAddr) == 0 || macAddr == "" {
		return false
	}

	return macAddrRegex.MatchString(macAddr)
}
