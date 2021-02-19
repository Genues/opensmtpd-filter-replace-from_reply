# OpenSMTPD opensmtpd-filter-replace-from_reply
This is a simple OpenSMTPD filter for overwriting the email address in the MAIL FROM command, in the header in all sent messages, as well as for substituting the original MAIL FROM address in the Reply-To field. Designed to send all relayed messages from a given email address (for example no-reply@example.com).

Это простой фильтр OpenSMTPD для перезаписи адреса электронной почты в команде MAIL FROM, в заголовке во всех отправляемых сообщениях, а так же для подстановки оригинального адреса MAIL FROM в поле Reply-To. 
Разработано для отправки всех ретранслируемых сообщений с заданного  адреса электронной почты (например no-reply@example.com).

## Usage
* build the filter for your target platform  
`env GOOS=linux GOARCH=amd64 go build opensmtpd-filter-replace-from_reply.go`

or for ArchLinux AUR package - <a href="https://aur.archlinux.org/packages/opensmtpd-filter-replace-from_reply">opensmtpd-filter-replace-from_reply</a>
* make OpenSMTPD use the filter
```
filter "replace-from_reply" proc-exec "opensmtpd-filter-replace-from_reply --mailFrom=no-reply@example.ru --fromToReply=true"
listen on socket filter "replace-from_reply"
```
option `--fromToReply` enables/disables the substitution of the original MAIL FROM address to the Reply-To

