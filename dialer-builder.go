package smtp

func NewDialer(displayName string) *hostBuilder {
	return &hostBuilder{
		displayName: displayName,
	}
}

type hostBuilder struct {
	displayName string
}

// Host sets the host of the dialer
func (b *hostBuilder) Host(host string) *portBuilder {
	return &portBuilder{
		displayName: b.displayName,
		host:        host,
	}
}

type portBuilder struct {
	displayName string
	host        string
}

// Port sets the port of the dialer
func (b *portBuilder) Port(port int) *usernameBuilder {
	return &usernameBuilder{
		displayName: b.displayName,
		host:        b.host,
		port:        port,
	}
}

type usernameBuilder struct {
	displayName string
	host        string
	port        int
}

// Username sets the username of the dialer
func (b *usernameBuilder) Username(username string) *passwordBuilder {
	return &passwordBuilder{
		displayName: b.displayName,
		host:        b.host,
		port:        b.port,
		username:    username,
	}
}

type passwordBuilder struct {
	displayName string
	host        string
	port        int
	username    string
}

// Password sets the password of the dialer
func (b *passwordBuilder) Password(password string) *Dialer {
	return &Dialer{
		DisplayName: b.displayName,
		Host:        b.host,
		Port:        b.port,
		Username:    b.username,
		Password:    password,
	}
}
