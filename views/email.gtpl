<html>
    <head>
    <title></title>
    </head>
    <body>

    <h3>Events {{.Day1}}</h3>
    <table border=1>
        {{range .TodayEvents}}
            <tr bgcolor="#ddd">
                <td>Title</td>
                <td>{{.Title}}</td>

            </tr>
            <tr>
                <td>Description</td>
                <td>{{.Info}}</td>
            </tr>
        {{end}}
    </table>

    <h3>Events {{.Day2}}</h3>
        <table border=1>
            {{range .TomorrowEvents}}
                <tr bgcolor="#ddd">
                    <td>Title</td>
                    <td>{{.Title}}</td>

                </tr>
                <tr>
                    <td>Description</td>
                    <td>{{.Info}}</td>
                </tr>
            {{end}}
        </table>

    <h3>Events {{.Day3}}</h3>
    <table border=1>
        {{range .DayAfterTomorrowEvents}}
            <tr bgcolor="#ddd">
                <td>Title</td>
                <td>{{.Title}}</td>

            </tr>
            <tr>
                <td>Description</td>
                <td>{{.Info}}</td>
            </tr>
        {{end}}
    </table>
    <h5> जय श्री कृपालु जी महाराज </h5>
    <h5> जय श्री कृष्णा जी </h5>
    <h5> जय श्री राधा जी </h5>
    <h5> जय श्री राम जी </h5>
    <h5> जय श्री सीता जी </h5>
    <h5> जय श्री लक्ष्मण जी </h5>
    <h5> जय श्री भरत जी </h5>
    <h5> जय श्री शत्रुघ्न जी </h5>
    <h5> जय श्री हनुमान जी </h5>
    <h5> जय श्री शिव शंकर महादेव जी </h5>
    <h5> जय श्री पार्वती जी </h5>
    <h5> जय श्री विष्णु जी </h5>
    <h5> जय श्री लक्ष्मी जी </h5>
    <h5> जय श्री साईं बाबा जी </h5>
    <h5> जय श्री दुर्गा माता जी </h5>
    <h5> जय श्री नारद जी </h5>

    </body>
</html>
