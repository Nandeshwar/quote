<html>
    <head>
    <title>Admin Event Detail</title>
    </head>
    <body>
        <form action="/admin-event-detail" method="post">
            <table>
                <tr>
                    <td>Title</td>
                    <td><input type="text" name="title" required="required"  size="100"></td>
                </tr>
                <tr>
                    <td>Info</td>
                    <td><textarea name="info" required="required" rows="20" cols="100"></textarea></td>
                </tr>
                <tr>
                    <td>Event Date</td>
                    <td><input type="text" required="required" name="eventDate" size="50"></td>
                </tr>

                 <tr>
                    <td>Event Type</td>
                    <td><input type="text" name="eventType" size="50"></td>
                 </tr>
                <tr>
                    <td>Links Pipeline(|) Separated</td>
                    <td><textarea  name="link" rows="4" cols="50"></textarea></td>
                </tr>
                <tr>
                    <td>Created Date</td>
                    <td><input type="text" name="createdAt" size="50"></td>
                </tr>
                <tr>
                    <td></td>
                    <td> <input type="submit" value="Create Info"></td>
                </tr>
            </table>
        </form>

        <h1>{{.Status}}</h1>

        <a href="/admin">Admin</a>

    </body>
</html
