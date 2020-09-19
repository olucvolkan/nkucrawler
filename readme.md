
NKU CV Crawler


Proje:
Tekirdag Namik Kemal Universitesinde ki akademisyenlerin siteleri olan cv uzantili sitelerinde ki profillerindeki bilgilerinin alinin bir database e kayit edilmesi gerekmektedir.

Nasil Kurulur?

Proje Golang ile gelistirilmistir. Oncelikli olarak bilgisayariniza golang kurmaniz gerekmektedir.
Golang kurulumu icin bu kaynak kullanilabilir.
- https://www.yakuter.com/go-programlama-dili-go-kurulumu/
Veriler Mysql database yazilmistir bu yuzden bilgisayar da mysql de kurulu olmasi gerekmektedir.

Bu kurulumlar yapildiktan sonra projenin icindeki .env dosyalarina  local mysql bilgileri girilmesi gerekmektedir.
Terminal de projenin kurulu oldugu dizine gidip ilgili go package kurulmalidir.

Sirasiyla bu komutlar konsolda calistirilmalidir.

1. go get github.com/PuerkitoBio/goquery
2. go get github.com/go-sql-driver/mysql
3. go get github.com/subosito/gotenv

Bunlar calistirildiktan sonra database olusturulup proje calisitirilabilir.

1. cd /db && go build database_create.go && go run database_create.go
2. go build main.go && go run main.go

Bu komutlar calistirildiktan sonra local mysql inizde database i ve kayitlari gorebilirsiniz. Projenin icinde ornek database vardir mysql import yaparak burdanda kayitlara bakabilirsiniz.

Ogrenci Numarasi:1150606041
Isim Soyisim: Mustafa Volkan Oluc
