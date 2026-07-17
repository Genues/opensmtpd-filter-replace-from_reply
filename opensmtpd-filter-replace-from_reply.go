package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"net/mail"
)

var mailFrom string
var mailSender string
var fromToReply bool
var debug uint

func init() {
	flag.StringVar(&mailFrom, "mailFrom", "", "Email for rewrite")
	flag.StringVar(&mailSender, "mailSender", "", "Sender name for rewrite")
	flag.BoolVar(&fromToReply, "fromToReply", true, "Move mail from address to Reply-To field")
	flag.UintVar(&debug, "debug", 0, "Debug mode")
	flag.Parse()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if (mailFrom == ""){
		log.Println("Not set email address for rewrite")
		return
	}

	skip := false

	for scanner.Scan() {
		line := scanner.Text()
		if debug > 3 {
			log.Printf("[4] %s", line);
		}
		if strings.HasPrefix(line, "config|ready") {
			fmt.Println("register|filter|smtp-in|data-line")
			fmt.Println("register|filter|smtp-in|mail-from")
			fmt.Println("register|ready")
		}else{
			dataSplit := strings.Split(line, "|")
			if debug > 2 {
				log.Printf("[3] %s", line)
			}
			if len(dataSplit) >= 8 {
				if debug > 1 {
					log.Printf("[2] %s", line)
				}
				switch dataSplit[4] {
					case "mail-from" :
						out := fmt.Sprintf("filter-result|%s|rewrite|%s\n", strings.Join(dataSplit[5:7], "|"), "<"+mailFrom+">")
						if debug > 0 {
							log.Printf("[1] %s", out)
						}
						fmt.Print(out)
					break;
					case "data-line":
						if strings.HasPrefix(strings.ToUpper(dataSplit[7]), "FROM:"){
							var from = strings.TrimSpace(dataSplit[7][5:]);
							dataSplit[7] = "From: "+mailSender+" <"+mailFrom+">"
							out := fmt.Sprintf("filter-dataline|%s\n", strings.Join(dataSplit[5:], "|"))
							if debug > 0 {
								log.Printf("[1] %s", out)
							}
							fmt.Print(out);
							if fromToReply && from != "" && valid(from) {
								dataSplit[7] = "Reply-To: "+from
								out := fmt.Sprintf("filter-dataline|%s\n", strings.Join(dataSplit[5:], "|"))
								if debug > 0 {
									log.Printf("[1] %s", out)
								}
								fmt.Print(out)
							}
							skip = true
						}else if strings.HasPrefix(strings.ToUpper(dataSplit[7]), "TO:") {
							skip = false
						}else if !skip {
							out := fmt.Sprintf("filter-dataline|%s\n", strings.Join(dataSplit[5:], "|"))
							if debug > 1 {
								log.Printf("[1] %s", out)
							}
							fmt.Print(out)
						}
					break;
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func valid(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}
