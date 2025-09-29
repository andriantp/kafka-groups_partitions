# kafka-groups_partitions

Experiments and learning notes on integrating **Apache Kafka** with **Golang**, structured step-by-step through a Medium article series.

## 📚 Chapters

1. **Deploy Apache Kafka with Docker (KRaft Mode) and AKHQ**  
   👉 [Read article](https://andriantriputra.medium.com/golang-x-kafka-1-how-to-deploy-apache-kafka-with-docker-kraft-mode-and-akhq-a103e43890c3)  
   🐳 Focus: setup Kafka with Docker (KRaft mode, without Zookeeper) + AKHQ as a monitoring UI.  

2. **Understanding Consumer Groups and Partitions**  
   👉 [Read article](https://andriantriputra.medium.com/golang-x-kafka-2-understanding-consumer-groups-and-partitions-8013bef73e2e)  
   📊 Focus: concepts of partitions, consumer groups, and how load balancing works.  

3. **Offset Management & Replay** 
   👉 [Read article](https://andriantriputra.medium.com/golang-x-kafka-3-offset-management-replay-145e5b57aad9)
   📊 Focus: cover how Kafka stores offsets, manual/auto commit, and replaying events from a specific offset.  

## 🚀 How to Run Examples

1. Clone this repository:
```bash
   $ git clone https://github.com/username/golang-x-kafka.git kafka
   $ cd kafka
```

2. Start Kafka with Docker Compose (see chapter-1/ folder):
```bash
$ docker-compose up -d
```

3. Navigate to each chapter folder to explore the code examples:
- chapter-1: Kafka & AKHQ setup
- chapter-2: producer/consumer with partitions & consumer groups
- chapter-3: offset management & replay (soon)


### 🛠 Tools
- Golang
- Apache Kafka (KRaft mode)
- AKHQ
- Docker & Docker Compose


### 📌 Notes
This repository is intended for learning purposes.
For production use, please consider additional configurations (security, monitoring, scaling).
