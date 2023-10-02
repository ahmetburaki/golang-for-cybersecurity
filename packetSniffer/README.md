# Packet Sniffer

Bu Go (Golang) programı, belirli bir ağ arabirimini dinleyerek ağ trafiğini yakalar ve bu trafiğe ilişkin bilgileri çıktıya yazar. Program, ağ trafiği analizi ve izleme için kullanışlıdır. Yakalanan veriler, bir log dosyasına yazılır ve çeşitli ağ katmanlarının bilgilerini içerir.

## Kurulum

Bu programı çalıştırmadan önce aşağıdaki adımları takip ederek gerekli bağımlılıkları yüklemeniz gerekmektedir:

1. Go'yu bilgisayarınıza [resmi web sitesinden](https://golang.org/dl/) indirip kurun.

2. Programın çalışması için `github.com/google/gopacket` paketini yüklemeniz gerekmektedir. Bu paketi aşağıdaki komutla yükleyebilirsiniz:

   ```bash
   go get github.com/google/gopacket
   ```

3. Programın kullanılacağı ağ arabirimini belirleyin ve `iface` değişkenini kod içinde güncelleyin. Varsayılan olarak `"en0"` olarak ayarlanmıştır.

## Kullanım

Program başlatıldığında, belirlediğiniz ağ arabirimini dinlemeye başlar ve yakalanan paketlerin bilgilerini ekrana ve bir log dosyasına yazar. Aşağıdaki bilgileri her paket için görüntüler:

- Kaynak MAC adresi
- Hedef MAC adresi
- Kaynak IP adresi (sadece IPv4)
- Hedef IP adresi (sadece IPv4)
- Kaynak port (sadece TCP ve UDP)
- Hedef port (sadece TCP ve UDP)
- Veri yükü (ilk 64 bayt)

Aşağıdaki komutla programı başlatabilirsiniz:

```bash
go run main.go
```

Program, ağ trafiğini dinlemeye başladığında paket bilgilerini görüntülemeye başlar.

## Örnek Kullanım

Aşağıda programın örnek bir kullanımı bulunmaktadır:

```bash
Packet Sniffer started. Listening for packets...
```

Program, "Packet Sniffer started" mesajını görüntüler ve ardından ağ trafiğini dinlemeye başlar.

## Notlar

- Program, sadece IPv4 adreslerini destekler ve sadece TCP ve UDP bağlantılarını izler.
- Yakalanan paketlerin bilgileri `outputs.log` adında bir log dosyasına yazılır.
- Ağ trafiği analizi yaparken dikkatli olun, bu tür işlemlerin yasalara ve kurallara uygun olduğundan emin olun.

## Lisans

Bu program açık kaynaklıdır ve MIT Lisansı altında dağıtılmaktadır. Daha fazla bilgi için [LICENSE](LICENSE) dosyasına başvurun.
