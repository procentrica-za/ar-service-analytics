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
	s.router.HandleFunc("/riskcriticalitydetails", s.handleGetRiskCriticalityDetails()).Methods("POST")
	s.router.HandleFunc("/riskcriticalityfilter", s.handleGetRiskCriticalityFilter()).Methods("GET")
	s.router.HandleFunc("/portfoliofiltered", s.handleGetPortfolioFilter()).Methods("POST")
	s.router.HandleFunc("/renewalprofiledetails", s.handleGetRenewalProfileDetails()).Methods("POST")
	s.router.HandleFunc("/portfoliofilteredcost", s.handleGetPortfolioFilterCost()).Methods("POST")
	s.router.HandleFunc("/yearreplacementdetails", s.handleGetYearReplacementDetails()).Methods("POST")
}
