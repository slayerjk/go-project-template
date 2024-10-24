package mailing

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"
	"time"
)

type MailData struct {
	Host          string   `json:"host"`
	Port          string   `json:"port"`
	AuthUser      string   `json:"auth_user"`
	AuthPass      string   `json:"auth_pass"`
	FromAddr      string   `json:"from_addr"`
	ToAddrErrors  []string `json:"to_addr_errors"`
	ToAddrReports []string `json:"to_addr_reports"`
}

// read json mailing data
func readMailingData(dataFile string) (MailData, error) {
	var result MailData

	// open file to read
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return result, fmt.Errorf("failed to read mailing data file:\n\t%v", err)
	}

	// read file content
	errU := json.Unmarshal(data, &result)
	if errU != nil {
		return result, fmt.Errorf("failed to unmarshall mailing data:\n\t%v", errU)
	}

	return result, nil
}

// send plain text mail;
// msgType may be: report/error or anything you like;
// appName - your app name;
// subject will be like "appName - msgType"
func SendPlainEmailWoAuth(dataFile, msgType, appName string, msg []byte, curDate time.Time) error {
	// read mailing data
	mailData, err := readMailingData(dataFile)
	if err != nil {
		return fmt.Errorf("failed to get mailing data file:\n\t%v", err)
	}

	// setting mail params
	fromAddr := mailData.FromAddr
	smtpHost := mailData.Host
	smtpHostAndPort := fmt.Sprintf("%s:%s", smtpHost, mailData.Port)
	// toAddrUsers := mailData.ToAddrUsers
	subject := fmt.Sprintf("%s - %s(%v)\n", appName, msgType, curDate.Format("02.01.2006 15:04"))

	// checking type of recepients to implement(errors/reports)
	var toAddr []string
	switch msgType {
	case "error":
		toAddr = mailData.ToAddrErrors
	case "report":
		toAddr = mailData.ToAddrReports
	default:
		return fmt.Errorf("wrong msgType: neither 'error' nor 'report'")
	}

	// sending email for all recepients in list
	for _, recepient := range toAddr {
		// Generate a random Message-ID
		// r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// messageID := strconv.FormatInt(r.Int63(), 10) + "@" + smtpHost

		message := "From: " + fromAddr + "\n" +
			"To: " + recepient + "\n" +
			"Subject: " + subject + ">\n\n" +
			// "MIME-version: 1.0;\n" +
			// "Content-Type: text/html; charset=\"UTF-8\";\n" +
			// "Message-ID: <" + messageID + ">\n\n" +
			string(msg)

		// making blank auth(no auth)
		// auth := smtp.PlainAuth("", "", "", smtpHost)

		conn, err := smtp.Dial(smtpHostAndPort)
		if err != nil {
			return fmt.Errorf("failed to dial to smtp server:\n\t%v", err)
		}

		// set sender
		if err := conn.Mail(fromAddr); err != nil {
			return fmt.Errorf("failed to set sender:\n\t%v", err)
		}

		//set recepient
		if err := conn.Rcpt(recepient); err != nil {
			return fmt.Errorf("failed to set recepient:\n\t%v", err)
		}

		// send the email body
		body, err := conn.Data()
		if err != nil {
			return fmt.Errorf("failed to set data:\n\t%v", err)
		}

		// write msg to body
		_, err = fmt.Fprint(body, message)
		if err != nil {
			return fmt.Errorf("failed to write msg to data:\n\t%v", err)
		}

		// close the body
		err = body.Close()
		if err != nil {
			return fmt.Errorf("failed to close data:\n\t%v", err)
		}

		// senc QUIT command and close connection
		err = conn.Quit()
		if err != nil {
			return fmt.Errorf("failed to quit the connection:\n\t%v", err)
		}
	}

	return nil
}
