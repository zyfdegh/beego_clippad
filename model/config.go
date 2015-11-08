package model

import{
  "net"
  "net/http"
}

type Config struct{
  ip net
  port int
  runMode string
}

func (s *Config)Load(f *file)(*config Config){
  return nil
}
