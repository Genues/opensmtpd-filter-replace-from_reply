# opensmtpd-filter-replace-from_reply
This is a simple OpenSMTPD filter for overwriting the email address in the MAIL FROM command, in the header in all sent messages, as well as for substituting the original MAIL FROM address in the Reply-To field. Designed to send all relayed messages from a given email address (for example no-reply@example.com).
