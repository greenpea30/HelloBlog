package zjulogin

type Auth struct {
	am     *ZJUAM
	config Config
}

func (a *Auth) ZJUAM() *ZJUAM {
	return a.am
}
