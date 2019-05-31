<html>
    <head>
    <title>Get INFO</title>
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
        {{if .Username}}
                <table align="right" border="1" id="customers">
                    <tr> 
                        <td>Kullanıcı : </td> <td>{{.Username}}</b></td>
                    </tr>

                    <tr align="center">
                        <td colspan="2"><a href="/logout">Oturumu Kapat</td>
                    </tr>
            </table>
            <br><br>

            <form action="/getInfo" method="post">
                <table border="1" id="customers" align="center">
                    <tr>
                        <td colspan="2" align="center">  <h1>Burası Bilgilendirme Ekranı </h1></td>
                    </tr>
                    <tr>
                        <td colspan="2" align="center">  <h5>Burası Alt başlık, lütfenli cümleler falan yazıyorum :). </h5></td>
                    </tr>
                    <tr>
                        <td>  Cihaz ID : </td> <td> <input type="text" name="deviceId"></td>
                    </tr>
                       <td colspan="2" align="center"><input type="submit" value="Bilgileri Getir"></td>
                    </tr>
                    {{if eq .SendMessage 1 }}<tr><td style="color:green" colspan="2" align="center">Mesaj Gönderilmiştir...</td></tr>{{else if eq .SendMessage 0}}<tr><td style="color:red" colspan="2" align="center">Mesaj Gönderilememiştir...{{end}}
                    {{if .DeviceIDError}}<tr><td style="color:red" colspan="2" align="center">Lütfen geçerli bir Cihaz ID Giriniz!</td></tr>{{end}}
                    {{if .DeviceIDNullError}}<tr><td style="color:red" colspan="2" align="center">Lütfen Cihaz ID Bölümünü Boş Bırakmayınız!</td></tr>{{end}}

                </table>
            </form>
        {{else}}
            <p>Bu sayfaya erişmek için ilk önce lütfen giriş yapınız. <a href="/login">Giriş</a></p>
            <meta http-equiv="refresh" content="3;URL=/login">
        {{end}}
    </body>
</html>
