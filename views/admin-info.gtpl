<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/admin-info" method="post">
            <table>
                <tr>
                    <td>Title</td>
                    <td><input type="text" name="title" required="required" pattern="[A-Za-z0-9]{1,20}"  size="100"></td>
                </tr>
                <tr>
                    <td>Info</td>
                    <td><textarea name="info" required="required" pattern="[A-Za-z0-9]{1,20}" rows="20" cols="100"></textarea></td>
                </tr>
                <tr>
                    <td>Links Comma Separated</td>
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
</html>
