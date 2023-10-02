# Password Generator

Bu basit Go (Golang) programı, kullanıcının istediği sayıda şifre kombinasyonunu oluşturmak için kullanılabilir. Kullanıcı, kombinasyon sayısını, karakter türünü ve kullanılacak iş parçacığı (thread) sayısını belirleyebilir. Program, belirtilen kombinasyon sayısına göre şifreleri üretecek ve bunları bir metin dosyasına kaydedecektir.

## Kurulum

Bu programı çalıştırmak için öncelikle Go dilini bilgisayarınıza kurmalısınız. Go'yu [resmi web sitesinden](https://golang.org/dl/) indirebilir ve kurabilirsiniz.

Daha sonra, bu depoyu yerel bir dizine klonlayın:

```bash
git clone <repo-link>
```

Klonladığınız dizine gidin ve programı çalıştırmak için aşağıdaki komutu kullanın:

```bash
go run main.go
```

## Kullanım

Program başlatıldığında, kullanıcıdan aşağıdaki bilgileri girmesi istenir:

1. Oluşturulacak şifre kombinasyonu sayısı
2. Kullanılacak karakter türü (1-6 arası bir sayı seçimi):
   - 1: Sayılar ve harfler
   - 2: Sadece sayılar
   - 3: Sadece harfler
   - 4: Özel karakterler
   - 5: Özel karakterler ve sayılar
   - 6: Özel karakterler, sayılar ve harfler
3. Kullanılacak iş parçacığı (thread) sayısı

Program, belirtilen parametrelere göre şifre kombinasyonlarını üretecek ve `password_list.txt` adında bir metin dosyasına kaydedecektir.

## Örnek Kullanım

Aşağıda programın örnek bir kullanımı bulunmaktadır:

```bash
How many password combinations do you want to create? Exp(3): 1000
Enter combination type (1-6): 6
How many threads to use? 4
```

Bu örnek, 1000 adet şifre kombinasyonunu özel karakterler, sayılar ve harflerle oluşturmak için 4 iş parçacığı kullanılacağını belirtmektedir.

## Notlar

- Daha fazla şifre kombinasyonu oluşturmak için `possibleCombination`, karakter türünü değiştirmek için `combinationType`, ve iş parçacığı sayısını artırmak için `numThreads` değerlerini ayarlayabilirsiniz.

## Lisans

Bu program açık kaynaklıdır ve MIT Lisansı altında dağıtılmaktadır. Daha fazla bilgi için [LICENSE](LICENSE) dosyasına başvurun.
