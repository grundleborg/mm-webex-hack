package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net/http"
)

func CreateMeeting() (*CreateMeetingResponse, int, error) {
	bodyContent := `
           <bodyContent xsi:type="java:com.webex.service.binding.meeting.CreateMeeting">                     
              <metaData>            
                  <confName>Sample Meeting</confName>            
              </metaData>            
              <schedule>            
                  <startDate/>            
              </schedule>            
          </bodyContent>`

	buf, sc, err := doRequest(bodyContent)
	if err != nil {
		return nil, sc, err
	}

	var message CreateMeetingResponse
	err = xml.Unmarshal(buf.Bytes(), &message)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &message, 0, nil
}

func GetMeetingHostUrl(meetingId string) (*GetHostMeetingUrlResponse, int, error) {
	bodyContent := `
			<bodyContent
            	xsi:type="java:com.webex.service.binding.meeting.GethosturlMeeting">
            	<sessionKey>` + meetingId + `</sessionKey>
        	</bodyContent>
`

	buf, sc, err := doRequest(bodyContent)
	if err != nil {
		return nil, sc, err
	}

	var message GetHostMeetingUrlResponse
	err = xml.Unmarshal(buf.Bytes(), &message)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &message, 0, nil
}

func GetMeetingJoinUrl(meetingId string) (*GetJoinMeetingUrlResponse, int, error) {
	bodyContent := `
        <bodyContent
            xsi:type="java:com.webex.service.binding.meeting.GetjoinurlMeeting">
            <sessionKey>`+meetingId+`</sessionKey>
            
        </bodyContent>
`

	//<attendeeName>Unknown</attendeeName>

	buf, sc, err := doRequest(bodyContent)
	if err != nil {
		return nil, sc, err
	}

	var message GetJoinMeetingUrlResponse
	err = xml.Unmarshal(buf.Bytes(), &message)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &message, 0, nil
}

func doRequest(bodyContent string) (*bytes.Buffer, int, error) {
	payload := `
<?xml version="1.0" encoding="UTF-8"?>
  <serv:message xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">            
      <header>            
          <securityContext>            
             <siteName>apidemoeu</siteName>           
             <webExID>george</webExID>           
             <password>Jc4Hp2Le</password>           
             <siteID>690319</siteID>           
             <partnerID>g0webx!</partnerID>           
             <returnAdditionalInfo>TRUE</returnAdditionalInfo>           
         </securityContext>           
     </header>           
     <body>           
` + bodyContent + `
     </body>           
 </serv:message>
`

	rq, err := http.NewRequest("POST", "https://apidemoeu.webex.com/WBXService/XMLService", bytes.NewReader([]byte(payload)))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	rq.Header.Set("Content-Type", "text/xml")
	rq.Close = true

	httpClient := http.Client{}
	if rp, err := httpClient.Do(rq); err != nil {
		return nil, http.StatusInternalServerError, err
	} else if rp == nil {
		return nil, http.StatusInternalServerError, errors.New("Received nil response when making request")
	} else if rp.StatusCode >= 300 {
		defer closeBody(rp)
		return nil, rp.StatusCode, errors.New("Received status code above 300")
	} else {
		defer closeBody(rp)
		buf := new(bytes.Buffer)
		buf.ReadFrom(rp.Body)
		return buf, 0, nil
	}
}
