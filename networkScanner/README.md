# Port Scanner

Bu Go (Golang) programı, belirtilen hedef üzerinde belirli bir port aralığına yönelik bağlantı denemeleri yaparak açık ve kapalı portları tespit eden bir port tarayıcısıdır. Program, ağ güvenliği ve ağ yapılandırmalarını denetlemek için kullanışlıdır.

## Kurulum

Bu programı çalıştırmadan önce aşağıdaki adımları takip ederek gerekli bağımlılıkları yüklemeniz gerekmektedir:

1. Go'yu bilgisayarınıza [resmi web sitesinden](https://golang.org/dl/) indirip kurun.

2. Programı çalıştırmadan önce aşağıdaki bağımlılıkları yükleyin:

   ```bash
   go get github.com/google/gopacket
   ```

3. Programın kullanılacağı hedefi, başlangıç ve bitiş portlarını, bağlantı süresini ve maksimum eşzamanlı bağlantıları belirleyin.

## Kullanım

Program başlatıldığında, aşağıdaki bilgileri kullanıcıdan alır:

- Hedef makinenin IP adresi veya alan adı (örneğin, "example.com" veya "192.168.1.1")
- Tara-nacak port aralığı (başlangıç ve bitiş portları)
- Bağlantı denemesinin süresi (saniye cinsinden)
- Maksimum eşzamanlı bağlantı sayısı

Ardından, program belirtilen port aralığında her port için bağlantı denemeleri yapar ve açık portları tespit eder. Sonuçları ekrana ve bir dizi açık port olarak görüntüler.

Aşağıdaki komutla programı başlatabilirsiniz:

```bash
go run main.go
```

Program, hedef makinede açık olan portları bulmaya başlar.

## Örnek Kullanım

Aşağıda programın örnek bir kullanımı bulunmaktadır:

```bash
Enter target host: example.com
Enter start port: 80
Enter end port: 100
Enter connection timeout (in seconds): 2
Enter maximum concurrent connections: 10
Scanning example.com...
Open ports:
Port 80 is open
```

Bu örnekte, "example.com" hedefinde 80 ile 100 arasındaki portlar taranır ve açık olanlar bulunur.

## Notlar

- Program, belirtilen port aralığında TCP bağlantı denemeleri yapar.
- Port tarayıcısı, yalnızca açık olan portları listeleyecektir.
- Bağlantı denemeleri sırasında hedef sistem üzerinde izinsiz taramalar yapmamaya dikkat edin.

## Lisans

Bu program açık kaynaklıdır ve MIT Lisansı altında dağıtılmaktadır. Daha fazla bilgi için [LICENSE](LICENSE) dosyasına başvurun.
