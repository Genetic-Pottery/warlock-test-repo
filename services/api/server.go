package api

type Server struct{ addr string }

const ListenAddr = "0.0.0.0:8080"

func Boot() *Server { return &Server{addr: ListenAddr} }
