<html>
    <head>
        <title>Other Page</title>
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

<script>
function goBack() {
  window.history.back();
}
</script>
    
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

            <form action="/otherPage" method="post">
                <table border="1" id="customers" align="center">
                    <tr>
                        <td colspan="6" align="center">  <h1>Burasıda Böyle Bir Sayfa </h1></td>
                    </tr>

                </table>
                <br><br>
                <table border="1" id="customers" align="center"> 
                    <tr>
                        <td>  Cihaz ID : </td> <td> <input type="text" name="deviceId" value={{.DeviceID}} disabled> </td>  
                    </tr>

                    <tr align="center">
                        <td><input type="button" value="Geri Dön" onclick="history.back()"></td> 
                        <td><input type="submit" value="Gönder"></td> 
                    </tr>

                </table>
            </form>
        {{else}}
             <p>Bu sayfaya erişmek için ilk önce lütfen giriş yapınız. <a href="/login">Giriş</a></p>
            <meta http-equiv="refresh" content="3;URL=/login">
        {{end}}
    </body>
</html>
