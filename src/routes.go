package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/assetflexvalcondition", s.handleGetAssetFlexValCondition()).Methods("GET")
	s.router.HandleFunc("/portfolio", s.handleGetPortfolio()).Methods("GET")
}
