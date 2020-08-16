package rotator

type Server struct {
	rotator *Rotator
}

func NewServer(rotator *Rotator) *Server {
	return &Server{rotator: rotator}
}

