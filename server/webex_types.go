package main

import "encoding/xml"

type CreateMeetingResponse struct {
	XMLName xml.Name `xml:"message"`
	Header Header `xml:"header"`
	Body CreateMeetingResponseBody `xml:"body"`
}

type CreateMeetingResponseBody struct {
	BodyContent CreateMeetingResponseBodyContent `xml:"bodyContent"`
}

type CreateMeetingResponseBodyContent struct {
	XMLName xml.Name `xml:"bodyContent"`
	MeetingKey string `xml:"meetingkey"`
	CalendarUrls CalendarUrls `xml:"iCalendarURL`
}

type GetHostMeetingUrlResponse struct {
	XMLName xml.Name `xml:"message"`
	Header Header `xml:"header"`
	Body GetHostMeetingUrlResponseBody `xml:"body"`
}

type GetHostMeetingUrlResponseBody struct {
	BodyContent GetHostMeetingUrlResponseBodyContent `xml:"bodyContent"`
}

type GetHostMeetingUrlResponseBodyContent struct {
	XMLName xml.Name `xml:"bodyContent"`
	HostMeetingURL string `xml:"hostMeetingURL"`
}

type GetJoinMeetingUrlResponse struct {
	XMLName xml.Name `xml:"message"`
	Header Header `xml:"header"`
	Body GetJoinMeetingUrlResponseBody `xml:"body"`
}

type GetJoinMeetingUrlResponseBody struct {
	BodyContent GetJoinMeetingUrlResponseBodyContent `xml:"bodyContent"`
}

type GetJoinMeetingUrlResponseBodyContent struct {
	XMLName xml.Name `xml:"bodyContent"`
	JoinMeetingURL string `xml:"joinMeetingURL"`
}

type Header struct {
	XMLName xml.Name `xml:"header"`
	Response Response `xml:"response"`
}

type Response struct {
	XMLName xml.Name `xml:"response"`
	Result string `xml:"result"`
}

type CalendarUrls struct {
	XMLName xml.Name `xml:"iCalendarURL"`
	Host string `xml:"host"`
	Attendee string `xml:"attendee"`
}
