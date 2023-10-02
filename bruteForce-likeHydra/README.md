# Brute Force Cracker

Bu Go (Golang) programı, belirli bir URL'deki bir kullanıcının şifresini kaba kuvvet saldırısı yaparak çözmek için kullanılır. Program, verilen bir kullanıcı adı ve yanlış şifre girişi hatası mesajını kullanarak şifre denemelerini gerçekleştirir ve doğru şifreyi bulduğunda sonucu kullanıcıya gösterir.

## Kurulum

Bu programı çalıştırmadan önce aşağıdaki adımları takip ederek gerekli bağımlılıkları yüklemeniz gerekmektedir:

1. Go'yu bilgisayarınıza [resmi web sitesinden](https://golang.org/dl/) indirip kurun.

## Kullanım

Program başlatıldığında, aşağıdaki bilgileri kullanıcıdan alır:

- Hedef URL (şifre denemesi yapılacak web sayfasının URL'si)
- Hedef kullanıcı adı
- Yanlış şifre girişi hatası mesajı (şifre denemesi sonuçlarına dayanarak yanlış girişlerin nasıl tespit edileceği)

Ardından, program verilen bir şifre listesi dosyasını açar, her şifre için şifre denemesi yapar ve doğru şifreyi bulduğunda sonucu kullanıcıya gösterir.

Aşağıdaki komutla programı başlatabilirsiniz:

```bash
go run main.go
```

Program, belirtilen URL üzerinde kullanıcı adı ve şifre denemelerini başlatır.

## Örnek Kullanım

Aşağıda programın örnek bir kullanımı bulunmaktadır:

```bash
    Enter Target Url: https://example.com/login
    Enter Target Username: john.doe
    Enter Wrong Password Error Message: Incorrect password
```

Bu örnekte, program "https://example.com/login" URL'si için "john.doe" kullanıcısı için şifre denemelerini başlatacak ve yanlış şifre girişi hatası mesajını "Incorrect password" olarak kabul edecektir.

## Notlar

- Program, şifre denemesi yaparken hedef web sayfasına HTTP POST isteği gönderir.
- Şifre listesi `passwords.txt` adında bir dosyadan okunur. Dosyanın her satırı bir şifreyi temsil etmelidir.
- Bu programın kötü niyetli amaçlarla kullanılmasından kaçının. Sadece yasal ve etik testler veya güvenlik denemeleri için kullanılmalıdır.

## Lisans

Bu program açık kaynaklıdır ve MIT Lisansı altında dağıtılmaktadır. Daha fazla bilgi için [LICENSE](LICENSE) dosyasına başvurun.
