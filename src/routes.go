package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/analyseassets", s.handleanalyseassets()).Methods("GET")

}
