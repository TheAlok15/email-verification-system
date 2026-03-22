package validators

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/TheAlok15/email-verification-system/internal/verifier/core"
)

type SMTPValidator struct{}

const (
	smtpTimeout = 5 * time.Second
	maxRetries  = 2
)

func (s SMTPValidator) Validate(ctx *core.VerificationContext) error {

	mxRecords, err := net.LookupMX(ctx.Domain)
	if err != nil || len(mxRecords) == 0 {
		ctx.Result.SMTPValid = false
		ctx.Result.SMTPMessage = "no mx record"
		return nil
	}

	mxHost := strings.TrimSuffix(mxRecords[0].Host, ".")

	var finalCode int
	var finalMsg string

	for attempt := 0; attempt < maxRetries; attempt++ {

		conn, err := net.DialTimeout("tcp", mxHost+":25", smtpTimeout)
		if err != nil {
			continue
		}

		conn.SetDeadline(time.Now().Add(smtpTimeout))
		reader := bufio.NewReader(conn)

		// 1️ Reading is done
		msg, err := readLine(reader)
		if err != nil {
			conn.Close()
			continue
		}

		//  HELO is done
		if err := writeLine(conn, "HELO localhost"); err != nil {
			conn.Close()
			continue
		}
		_, _ = readLine(reader)

		//  Mail is done
		if err := writeLine(conn, "MAIL FROM:<test@example.com>"); err != nil {
			conn.Close()
			continue
		}
		_, _ = readLine(reader)

		// rcpt done
		if err := writeLine(conn, "RCPT TO:<"+ctx.Email+">"); err != nil {
			conn.Close()
			continue
		}

		msg, err = readLine(reader)
		if err != nil {
			conn.Close()
			continue
		}

		code := parseCode(msg)

		finalCode = code
		finalMsg = strings.TrimSpace(msg)

		// Decide validity
		if code == 250 {
			ctx.Result.SMTPValid = true
		} else {
			ctx.Result.SMTPValid = false
		}

		//  Retry only for temporary errors
		if code == 421 || code == 450 {
			conn.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		// Catch-all detection
		fakeEmail := "random123456@" + ctx.Domain
		if err := writeLine(conn, "RCPT TO:<"+fakeEmail+">"); err == nil {
			fakeMsg, err := readLine(reader)
			if err == nil {
				fakeCode := parseCode(fakeMsg)
				if fakeCode == 250 {
					ctx.Result.CatchAll = true
				}
			}
		}

		//  quit
		_ = writeLine(conn, "QUIT")
		conn.Close()

		break
	}

	ctx.Result.SMTPCode = finalCode
	ctx.Result.SMTPMessage = finalMsg

	return nil
}


func readLine(r *bufio.Reader) (string, error) {
	return r.ReadString('\n')
}

func writeLine(conn net.Conn, msg string) error {
	_, err := conn.Write([]byte(msg + "\r\n"))
	return err
}

func parseCode(msg string) int {
	code := 0
	fmt.Sscanf(msg, "%d", &code)
	return code
}
