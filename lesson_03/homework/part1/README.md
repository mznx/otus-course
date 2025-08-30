# Часть 1

Шаги выполнения:

1. Запуск окружения для проведения тестов

    - выполнен переход в директорию `dz2-part1/dz2-part1-postgres`
    - выполнена команда `docker compose up -d --build`
    - результат - подняты контейнеры с окружением:

    ```sh
    developer@ubuntu-dev:/$ docker ps
    CONTAINER ID   IMAGE                                COMMAND                  CREATED          STATUS          PORTS                                         NAMES
    cb8d6b10104e   dz2-part1-postgres-nginx             "/docker-entrypoint.…"   31 seconds ago   Up 29 seconds   0.0.0.0:8080->80/tcp, [::]:8080->80/tcp       nginx
    83d0a62d981a   dz2-part1-postgres-backend1          "python app.py"          31 seconds ago   Up 30 seconds   5000/tcp                                      backend1
    d936086eec66   dz2-part1-postgres-backend2          "python app.py"          31 seconds ago   Up 30 seconds   5000/tcp                                      backend2
    54ab10e0f487   dz2-part1-postgres-haproxy           "docker-entrypoint.s…"   31 seconds ago   Up 30 seconds   0.0.0.0:5432->5432/tcp, [::]:5432->5432/tcp   haproxy
    4e295d4ef894   dz2-part1-postgres-postgres-slave1   "docker-entrypoint.s…"   31 seconds ago   Up 30 seconds   0.0.0.0:5433->5432/tcp, [::]:5433->5432/tcp   postgres-slave1
    61c33e8d0668   dz2-part1-postgres-postgres-slave2   "docker-entrypoint.s…"   31 seconds ago   Up 30 seconds   0.0.0.0:5434->5432/tcp, [::]:5434->5432/tcp   postgres-slave2
    ```

2. Проведение тестирования

    A) Отключение PostgreSQL слейва

    - запущен скрипт имитации нагрузки `./load-test/run.sh`
    - остановлен один инстанс PostgreSQL командой `docker kill postgres-slave1`
    - все запросы успешно переключились на второй инстанс PostgreSQL
    - восстановление инстанста PostgreSQL командой `docker compose restart postgres-slave1`
    - запросы продолжили чередование между инстансами PostgreSQL
    - вывод работы скрипта с комментариями:

    ```sh
    # скрипт запущен, запросы чередуются между 2мя инстансами PostgreSQL
    Send 20 requests to http://localhost:8080/data...
    {"backend_id":"backend1","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:43:15.595798+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:16.114180+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:43:16.631692+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:17.195178+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:43:17.718270+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:18.240226+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:43:18.757926+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:19.277714+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:43:19.794264+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:20.320623+00:00"}

    # выполнена остановка одного инстанса, все запросы пошли на 172.19.0.2

    {"backend_id":"backend1","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:30.938690+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:41.484304+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:42.099634+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:42.644410+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:43.162306+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:43.696023+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:44.209646+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:44.724890+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:45.237114+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:45.755093+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:46.281017+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:46.850541+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:47.407863+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:47.923532+00:00"}

    # перезапущен инстанс PostgreSQL, запросы продолжили чередование

    {"backend_id":"backend1","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:43:48.453857+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:48.972428+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:43:49.491117+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:50.010394+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:43:50.536508+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:43:51.068312+00:00"}
    ```

    Б) Отключение backend-инстанса

    - запущен скрипт имитации нагрузки `./load-test/run.sh`
    - остановлен один инстанс backend командой `docker kill backend1`
    - все запросы успешно переключились на второй инстанс backend
    - восстановление инстанста backend командой `docker compose restart backend1`
    - запросы продолжили чередование между инстансами backend
    - вывод работы скрипта с комментариями:

    ```sh
    # скрипт запущен, запросы чередуются между 2мя инстансами backend
    Send 20 requests to http://localhost:8080/data...
    {"backend_id":"backend1","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:54:41.946693+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:54:42.471355+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:54:43.026532+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:54:43.553540+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:54:44.068631+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:54:44.586844+00:00"}

    # выполнена остановка одного инстанса, все запросы пошли на backend2

    {"backend_id":"backend2","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:55:45.172814+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:55:45.694219+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:55:49.352838+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:55:49.870707+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:55:50.406020+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:55:50.938609+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:55:51.472657+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:55:52.003654+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:55:52.517546+00:00"}

    # перезапущен инстанс backend, запросы продолжили чередование

    {"backend_id":"backend1","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:55:53.038428+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:55:53.576418+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:55:54.126586+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:55:54.643538+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:55:55.164796+00:00"}

    {"backend_id":"backend2","postgres_ip":"172.19.0.3","timestamp":"2025-08-30 19:55:55.682458+00:00"}

    {"backend_id":"backend1","postgres_ip":"172.19.0.2","timestamp":"2025-08-30 19:55:56.201379+00:00"}
    ```
