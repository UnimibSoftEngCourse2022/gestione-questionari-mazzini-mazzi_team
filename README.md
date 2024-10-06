# Quiz App

---
### Descrizione 

Questo progetto è un'applicazione che utilizza Docker Compose per gestire i suoi servizi. Prima di avviare il progetto, è necessario configurare le variabili d'ambiente e assicurarsi che Docker e Docker Compose siano correttamente installati. 

### Prerequisiti 

- **Docker**: Assicurati di avere Docker installato. Puoi scaricarlo e installarlo seguendo [questa guida](https://docs.docker.com/get-docker/). 
- **Docker Compose**: Assicurati di avere Docker Compose installato. Puoi verificarlo con il seguente comando: ```docker-compose --version```

### Istruzioni per L'installazione:

1. clonare il repository

```bash
git clone https://github.com/UnimibSoftEngCourse2022/gestione-questionari-mazzini
cd gestione-questionari-mazzini
```

2. Creare e configurare il file `.env`

	Assicurati di avere un file `.env` nella root del progetto. Questo file contiene tutte le variabili d'ambiente necessarie per far funzionare i container Docker. Se non hai un file `.env`, puoi crearne uno basato sul file di esempio `.env.example`:

```bash 
cp .env.example .env
```

3. Compilare e avviare i container con Docker Compose

```bash
docker compose up -d --build
```

4. Accedere all'applicazione

	Una volta che tutti i container sono in esecuzione, l'applicazione sarà accessibile agll' URL http://localhost:8080.



[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/ekNf-lOK)

