package main

func main () { 
	c := http.Client(Timeout: time.Duration(1) * time.Second)
	resp, err := c.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}
	def resp.Body.
}