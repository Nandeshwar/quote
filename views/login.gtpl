<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/login" method="post">
            Admin Username:<input type="text" name="username">
            <br>
            Admin Password:<input type="password" name="password">
            <input type="submit" value="Login">
        </form>

        <a target="_blank"  href="http://localhost:1922/quotes-devotional" > http://localhost:1922/quotes-devotional </a>
        <br>
        <a target="_blank" href="http://localhost:1922/quotes-motivational"> http://localhost:1922/quotes-motivational </a>
        <br>
        <a target="_blank" href="http://localhost:1922/search/Kaiso|Hai"> http://localhost:1922/search/Kaiso|Hai (search criteria can be delimited by '|'
) </a>

         {{range .Events}}
            <h3>Events {{.Day}}</h3>
            <table border=1>
                {{range .EventList}}
                    <tr bgcolor="#ddd">
                        <td>ID</td>
                        <td>{{.ID}}</td>

                    </tr>
                    <tr>
                        <td>Title</td>
                        <td>{{.Title}}</td>
                    </tr>

                    <tr>
                        <td>Info</td>
                        <td>{{.Info}}</td>
                    </tr>

                    <tr>
                        <td>Date</td>
                        <td>{{.EventDate.Format "Monday Jan 02, 2006"}}</td>
                    </tr>
                     <tr>
                        <td>Links</td>
                        <td>
                            <table>
                            {{ range $link := .Links}}
                               <tr>
                                    <td><a href={{$link}} target="_blank">Click Me</a></td>
                               </tr>
                            {{end}}
                            </table>
                        </td>
                     </tr>
                {{end}}
            </table>
        {{end}}


    </body>
</html>
