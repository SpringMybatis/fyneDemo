package internal


type NetCard struct {
	name string
	ipv4 string
	mac  string
}

func (nc *NetCard)GetName() string {
	return nc.name
}

func (nc *NetCard)GetIPv4Addr() string {
	return nc.ipv4
}

func (nc *NetCard)GetMacAddr() string {
	return nc.mac
}