package pdfGenerator

import(
	"github.com/jung-kurt/gofpdf"
)

func CreatePdf(m map[string]string){
	pdf := gofpdf.New("P", "mm", "A4", "")
	fontSize := 16.0
	ht := pdf.PointConvert(fontSize)
	pdf.AddPage()
	pdf.SetFont("Arial", "", 16)
	pdf.Cell(40,10,"Dear, Customer")
	pdf.Ln(10)
	pdf.Cell(40,10,"Greeting from Doubtnut")
	pdf.Ln(10)
	pdf.Cell(40,10, "We are delighted to have you onboard")
	pdf.Ln(10)
	pdf.Cell(40,10, "Based on your recent searches, Here are some Links of your similar interests")
	pdf.Ln(20)

	for k,v := range m {
		htmlStr := ` See <a href="` + k + `">` + v + `</a>`
		html := pdf.HTMLBasicNew()
		html.Write(ht, htmlStr)
		pdf.Ln(20)
	}
	pdf.OutputFileAndClose("similar.pdf")
}