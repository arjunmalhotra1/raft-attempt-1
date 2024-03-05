package raftconfig

type servers struct {
	NodeNum int
	Address string
	Port    int
}

func NewServers() []servers {
	listServers := []servers{
		{
			NodeNum: 0,
			Address: "localhost",
			Port:    15000,
		},
		{
			NodeNum: 1,
			Address: "localhost",
			Port:    16000,
		},
		{
			NodeNum: 2,
			Address: "localhost",
			Port:    17000,
		},
		{
			NodeNum: 3,
			Address: "localhost",
			Port:    18000,
		},
		{
			NodeNum: 4,
			Address: "localhost",
			Port:    19000,
		},
	}

	return listServers
}
