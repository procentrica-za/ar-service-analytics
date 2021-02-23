package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/assetflexvalcondition", s.handleGetAssetFlexValCondition()).Methods("GET")
	s.router.HandleFunc("/portfolio", s.handleGetPortfolio()).Methods("GET")
	s.router.HandleFunc("/yearreplacement", s.handleGetYearReplacement()).Methods("GET")
	s.router.HandleFunc("/renewalprofile", s.handleGetRenewalProfile()).Methods("GET")
	s.router.HandleFunc("/riskcriticality", s.handleGetRiskCriticality()).Methods("GET")
	s.router.HandleFunc("/replacementbycondition", s.handleGetReplacementByCondition()).Methods("GET")

	s.router.HandleFunc("/riskcriticalitydd", s.handleGetRiskCriticalityDrillDown()).Methods("GET")
}
