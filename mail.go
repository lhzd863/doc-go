import (
        "gopkg.in/gomail.v2"
        "fmt"
)

type Mail struct {
        From   string
        Host   string
        Port   int
        Usr    string
        Passwd string
}

func NewMail(host string, port int, usr string, passwd string) (*Mail, error) {
        ma := &Mail{
                From:   usr,
                Host:   host,
                Port:   port,
                Usr:    usr,
                Passwd: passwd,
        }
        return ma, nil
}

func (m *Mail) SendMsg(tousr string, subject string, context string, attach string) {
        m1 := gomail.NewMessage()
        m1.SetHeader("From", m.From)
        m1.SetHeader("To", tousr)
        //m1.SetAddressHeader("Cc", m.From, "lhzd863")
        m1.SetHeader("Subject", subject)
        m1.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
        if len(attach) != 0 {
                m1.Attach(attach)
        }

        d := gomail.NewDialer(m.Host, m.Port, m.Usr, m.Passwd)

        // Send the email to Bob, Cora and Dan.
        if err := d.DialAndSend(m1); err != nil {
                panic(err)
        }
        fmt.Println("mail send success...")
}
