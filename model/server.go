package model

type Server struct{
  Config *controller.Config
}

func (s *Server)NewServer()(*Server){
  return new(Server)
}

func (s *Server)Serve(){

}
