package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"regexp"
)

var mailFrom string
var fromToReply bool

func init() {
	flag.StringVar(&mailFrom, "mailFrom", "", "Email for rewrite")
	flag.BoolVar(&fromToReply, "fromToReply", true, "Move mail from addres to Reply-To field")
	flag.Parse()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if (mailFrom == ""){
		log.Println("Not set email address for rewrite")
		return
	}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "config|ready") {
			fmt.Println("register|filter|smtp-in|data-line")
			fmt.Println("register|filter|smtp-in|mail-from")
			fmt.Println("register|ready")
		}else{
			dataSplit := strings.Split(line, "|")
			if len(dataSplit) >= 8 {
				switch dataSplit[4] {
					case "mail-from" :
						fmt.Printf("filter-result|%s|rewrite|%s\n", strings.Join(dataSplit[5:7], "|"), "<"+mailFrom+">")
					break;
					case "data-line":
						if strings.HasPrefix(strings.ToUpper(dataSplit[7]), "FROM: "){
							re := regexp.MustCompile(`(?i)From:[\s]*([\S]+)`);
							if (fromToReply){
								var from = re.FindStringSubmatch(dataSplit[7]);
								if (len(from) >= 2){
									dataSplit[7] = "Reply-To: "+from[1]
									fmt.Printf("filter-dataline|%s\n", strings.Join(dataSplit[5:], "|"))
								}
							}
							dataSplit[7] = "From: <"+mailFrom+">"
							fmt.Printf("filter-dataline|%s\n", strings.Join(dataSplit[5:], "|"))
						}else{
							fmt.Printf("filter-dataline|%s\n", strings.Join(dataSplit[5:], "|"))
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
