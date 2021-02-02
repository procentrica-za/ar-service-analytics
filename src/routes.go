package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/assetflexvalcondition", s.handleGetAssetFlexValCondition()).Methods("GET")

}
