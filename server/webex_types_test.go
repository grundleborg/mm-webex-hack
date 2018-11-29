package main

import (
	"encoding/xml"
	"testing"
)

func TestWebexTypes(t *testing.T) {

	sample := `
<?xml version="1.0" encoding="UTF-8"?>
<serv:message xmlns:serv="http://www.webex.com/schemas/2002/06/service" xmlns:att="http://www.webex.com/schemas/2002/06/service/attendee" xmlns:com="http://www.webex.com/schemas/2002/06/common" xmlns:meet="http://www.webex.com/schemas/2002/06/service/meeting">
   <serv:header>
      <serv:response>
         <serv:result>SUCCESS</serv:result>
         <serv:gsbStatus>BACKUP</serv:gsbStatus>
         <serv:respTime>142</serv:respTime>
         <serv:caseID>gxtx7tc002.webex.com 20181128190339</serv:caseID>
      </serv:response>
   </serv:header>
   <serv:body>
      <serv:bodyContent xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="meet:createMeetingResponse">
         <meet:meetingkey>299026987</meet:meetingkey>
         <meet:iCalendarURL>
            <serv:host>https://apidemoeu.webex.com/apidemoeu/j.php?MTID=mf2c13f7f773cd7ccdf25bdd73d3fcf9a</serv:host>
            <serv:attendee>https://apidemoeu.webex.com/apidemoeu/j.php?MTID=m3e3247a2823cd37f47f7cdf084b8faf6</serv:attendee>
         </meet:iCalendarURL>
         <meet:guestToken>336b6a197d73eab748f98c6241b53ac6</meet:guestToken>
      </serv:bodyContent>
   </serv:body>
</serv:message>
`

	var message Message
	xml.Unmarshal([]byte(sample), &message)

	if message.Body.BodyContent.CalendarUrls.Attendee != "https://apidemoeu.webex.com/apidemoeu/j.php?MTID=m3e3247a2823cd37f47f7cdf084b8faf6" {
		t.Fatalf(message.Body.BodyContent.CalendarUrls.Attendee)
	}
}
