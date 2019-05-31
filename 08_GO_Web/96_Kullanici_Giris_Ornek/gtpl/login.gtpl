<html>

    <head>
        <title>Login</title>
        <style>
            #customers {
                font-family: "Trebuchet MS", Arial, Helvetica, sans-serif;
                border-collapse: collapse;
            }

            #customers td, #customers th {
                border: 1px solid #ddd;
                padding: 8px;
            }

            #customers tr:nth-child(even){background-color: #f2f2f2;}

            #customers tr:hover {background-color: #ddd;}

            #customers th {
                padding-top: 12px;
                padding-bottom: 12px;
                text-align: left;
                background-color: #4CAF50;
                color: white;
            }
        </style>
    </head>

    <body>
    <br><br><br><br>

        <form method="post" action="/login">
            {{if .Username}}
                <p><b>{{.Username}}</b>, you're already logged in! <a href="/logout">Logout!</a></p>
                <meta http-equiv="refresh" content="2;URL=/getInfo">
            {{else}}            
                <table border="1" id="customers" align="center">
                    <tr>
                        <td colspan="4" align="center">  <h1>Giriş Ekranı</h1></td>
                    </tr>
                    <tr>
                        <td colspan="4" align="center">  <h5>Lütfen kullanıcı adı ve şifrenizi giriniz. </h5></td>
                    </tr>
                    <tr>
                        <td><label>Username:</label></td>
                        <td><input type="text" name="Username"><br></td>
                        <td><label>Password:</label></td>
                        <td><input type="password" name="Password"></td>
                    </tr>

                    <tr>
                        <td colspan="4" align="center"><input type="submit" name="Login" value="Giriş"></td>
                    </tr>
            {{end}}
            </table>
        </form>
        {{if .LoginError}}<table id="customers" align="center"><tr><td style="color:red">Kullanıcı adı yada şifre yanlış</td></tr>{{end}}</table>
    </body>
</html>