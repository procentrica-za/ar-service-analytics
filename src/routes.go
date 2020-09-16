package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/sheet", s.handleaddsheetdata()).Methods("POST")

}
