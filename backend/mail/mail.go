package mail

import (
	"errors"
	"fmt"

	"github.com/rg-km/final-project-engineering-11/backend/config"
	"gopkg.in/gomail.v2"
)

func SendMail(emailto, BookID, name string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.CONFIG_SENDER_NAME)
	mailer.SetHeader("To", emailto)
	mailer.SetAddressHeader("Cc", "hicoder224@gmail.com", "HiCoder")
	mailer.SetHeader("Subject", "Permintaan Mentoring "+BookID)
	mailer.SetBody("text/html", `<html>
	<body>
	   <h1>Hi `+name+` Anda Mendapatkan Request Mentoring</h1>
       <p>Jika anda ingin menerima mentoring klik gambar ceklis dan jika anda menolak silahkan anda klik gambar silang!!</p>
       <br>
      
          <a href="http://localhost:8090/api/mentor/acc/`+BookID+`"> 
		  <img alt="Qries" src="https://w7.pngwing.com/pngs/412/774/png-transparent-red-mark-cross-crossed-wrong-incorrect-sign-symbol-icon.png"
		  width="90" height="70">
	   </a>
       &nbsp &nbsp &nbsp &nbsp &nbsp &nbsp
	   <a href="http://localhost:8090/api/mentor/acc/`+BookID+`">
		  <img alt="Qries" src="https://www.pngitem.com/pimgs/m/423-4236368_icon-ceklis-png-transparent-png.png"
		  width=150" height="80">
	   </a>
	</body>
 </html>`)

	dialer := gomail.NewDialer(
		config.CONFIG_SMTP_HOST,
		config.CONFIG_SMTP_PORT,
		config.CONFIG_AUTH_EMAIL,
		config.CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return errors.New("Error Request Mentoring")
	}
	return nil

}

func CreateBookId(id int) string {
	res := id + 1
	return "HICOD00" + fmt.Sprintf("%d", res)
}
