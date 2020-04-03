<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/create" method="post">

            <table>
                <tr>
                    <td>  Kullanıcı Adı : </td> <td> <input type="text" name="username"></td>
                    <td>  Parola </td> <td><input type="password" name="password"></td> 
                </tr>

                <tr>
                    <td>  Yaş : </td> <td><input type="text" name="age"></td> 
                    <td>  E-Mail : </td> <td><input type="text" name="email"></td> 
                </tr>

                <tr>
                 <td>  Şehir : </td>
                    <td><select name="city">
                        <option value="Bursa">Bursa</option>
                        <option value="Uşak">Uşak</option>
                        <option value="Kocaeli">Kocaeli</option>
                    </select></td>

                    <td>  Cinsiyet : </td>
                    <td><input type="radio" name="gender" value="1">Erkek
                    <input type="radio" name="gender" value="2">Kadın</td>
                </tr>

                <tr>
                    <td>  Hobiler : </td>
                    <td>
                        <input type="checkbox" name="hobby" value="spor">Spor
                        <input type="checkbox" name="hobby" value="muzik">Müzik
                        <input type="checkbox" name="hobby" value="dans">Dans
                    </td>
                </tr>
                <tr>
                   <td></td><td><input type="submit" value="Giriş"></td>
                </tr>
            </table>
        </form>
    </body>
</html>