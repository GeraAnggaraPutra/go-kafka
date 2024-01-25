# Go Kafka

### Consumer Kafka:

- Mengkonsumsi pesan dari topik tertentu ("my-topic" dalam contoh)
- Menampilkan pesan sebagai string ke terminal
- Mengatur strategi rebalance dan offset awal untuk grup konsumen

### Producer Kafka:

- Mengirim pesan string ("Hello, world!" dalam contoh) ke topik tertentu ("my-topic" dalam contoh)
- Menggunakan produsen asynchronous

### Struktur File:

- File ini dibagi menjadi dua package: consumer dan producer.
- Setiap package berisi fungsi main untuk menjalankan masing-masing role.

## Consumer Role:

1. Konfigurasi
   - config.ClientID: Set nama ID untuk konsumen.
   - config.Consumer.Group.Rebalance.Strategy: Pilih strategi rebalance grup (Range dalam contoh).
   - config.Consumer.Offsets.Initial: Tentukan offset awal untuk membaca pesan (Oldest dalam contoh).
2. Membuat Grup Konsumen:
   - Gunakan sarama.NewConsumerGroup untuk membuat grup konsumen dengan konfigurasi, alamat server Kafka, dan nama grup.
3. Subscribe ke Topik:
   - Gunakan consumerGroup.Consume untuk subscribe ke topik yang diinginkan dan handler pemrosesan pesan.
4. Handler Pemrosesan Pesan:
   - Fungsi ConsumeClaim dipanggil untuk setiap pesan yang tersedia.
   - for message := range claim.Messages(): Iterasi melalui pesan yang tersedia.
   - value := string(message.Value): Ubah byte pesan menjadi string.
   - fmt.Println(value): Cetak string pesan ke terminal.
   - session.MarkMessage(message, ""): Tandai pesan sebagai telah dikonsumsi.
5. Menutup Grup Konsumen:
   - Gunakan defer consumerGroup.Close() untuk menutup grup konsumen saat program berakhir.

## Producer Role:

1. Konfigurasi:
   - config.ClientID: Set nama ID untuk produsen.
2. Membuat Pesan:
   - Gunakan sarama.ProducerMessage untuk membuat objek pesan dengan topik dan nilai (string dalam contoh).
3. Membuat Produsen Asynchronous:
   - Gunakan sarama.NewAsyncProducer untuk membuat produsen dengan konfigurasi dan alamat server Kafka.
4. Kirim Pesan:
   - Gunakan p.Input() <- message untuk mengirim pesan ke channel input produsen.
5. Menutup Produsen:
   - Gunakan defer p.Close() untuk menutup produsen saat program berakhir.

## Running Kafka (Windows):

- Run zookeeper
  ```bash
  [your kafka directory] bin\windows\zookeeper-server-start.bat config\zookeeper.properties
  ```
- Run kafka server
  ```bash
  [your kafka directory] bin\windows\kafka-server-start.bat config\server.properties
  ```
- Run producer
  ```bash
  [your kafka directory] bin\windows\kafka-console-producer.bat --broker-list localhost:9092 --topic your-topic-name
  ```
- Run consumer
  ```bash
  [your kafka directory] bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic your-topic-name --from-beginning
  ```

## Running App:

- Run consumer app
  ```bash
  go run cmd/consumer/main.go
  ```
- Run producer app
  ```bash
  go run cmd/producer/main.go
  ```

## Kafka Documentation

[Install Golang Kafka Package Sarama](https://github.com/IBM/sarama)

<br />

[Download JDK(Java Development Kit) By Eclipse Temurin](https://adoptium.net/temurin/releases/)

<br />

[Download Kafka](https://kafka.apache.org/downloads)