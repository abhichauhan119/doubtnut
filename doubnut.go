package main

import(
	"doubtnut/config"
	"doubtnut/pdfGenerator"
	"time"
	"os"
)

func main() {
	config.ReadConfig()
	
	m := make(map[string]string)
	for _,questions := range config.SimilarQuestions{
		for _,input := range questions{
			m[input.Question] = input.VideoLink
		}
	}
	// Creating a channel
	received := make(chan bool)
	
	// This goroutine will wait for 5 minutes of inactivity and then call createPDF
	go checkInactivity(received)

	// Assuming this main call as a hit from USER asking for PDF
	pdfGenerator.CreatePdf(m)

	// Checking whether 5 minutes elapsed - Then calling createPDF
	for {
           select {
           case <-received:
           		m["CalledThroughGoRouting"] = "https://www.google.com?q=goroutine in go"
                pdfGenerator.CreatePdf(m)
                os.Exit(0)
           }
    }

}

func checkInactivity(received chan bool){
	time.Sleep(10*time.Second)
	received<-true
}