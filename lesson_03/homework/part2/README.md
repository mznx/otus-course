# Часть 2

## Шаги выполнения:

### 1. Запуск окружения для проведения тестов

- развернут кластер через выполнение команды `./k3d/create.sh`

    ```sh
    developer@ubuntu-dev:~$ k3d node ls
    NAME                     ROLE           CLUSTER   STATUS
    k3d-hl-k3d-agent-0       agent          hl-k3d    running
    k3d-hl-k3d-agent-1       agent          hl-k3d    running
    k3d-hl-k3d-server-0      server         hl-k3d    running
    k3d-hl-k3d-serverlb      loadbalancer   hl-k3d    running
    k3d-registry.localhost   registry                 running
    ```

- собраны образы и залиты в registry командой `./images/build_and_push.sh`

    ```sh
    developer@ubuntu-dev:~$ docker images | grep k3d-registry
    k3d-registry.localhost:5000/backend            local          fa3c766edfe1   2 days ago      218MB
    k3d-registry.localhost:5000/postgres-replica   local          42fe1db22fa3   2 days ago      452MB
    k3d-registry.localhost:5000/haproxy            local          94cea1f8b307   2 days ago      102MB
    k3d-registry.localhost:5000/postgres-primary   local          bbb2a10a22f0   2 days ago      452MB
    k3d-registry.localhost:5000/nginx              local          6769dc3a703c   4 months ago    48.2MB
    ```

- в кластере подняты поды для тестирования командой `./k8s/apply.sh`

    ```sh
    developer@ubuntu-dev:~$ kubectl get po -n highload-dns
    NAME                                READY   STATUS    RESTARTS      AGE
    backend-c65c887b9-6bt42             1/1     Running   0             17s
    backend-c65c887b9-kkrck             1/1     Running   0             17s
    haproxy-65685c8ff7-9ww94            1/1     Running   0             17s
    nginx-5c456d678-5fqpv               1/1     Running   1 (10s ago)   17s
    postgres-primary-0                  1/1     Running   0             17s
    postgres-replicas-8c79c9899-5ct46   1/1     Running   0             17s
    postgres-replicas-8c79c9899-p95h2   1/1     Running   0             17s
    ```

### 2. Проведение тестирования

#### ⚠️ Важно

При запуске тестов с предоставленной конфигурацией, на фазе B, после скаллирования бекенда до одного инстанса и обратно, запросы переставали балансироваться между двумя бекендами. На следующих стадиях все запросы к бекенду завершались 502й ошибкой.

После ресерча было выявлено, что nginx ресолвит адреса бекенда один раз при запуске, и после смены ip у подов бекенда, продолжает пытаться ходить на первоначальные ip.

Для исправления, в конфигурацию nginx были внесены представленные ниже изменения, что позволило пройти все фазы тестирования успешно.

```diff
events {}
http {
+  resolver kube-dns.kube-system.svc.cluster.local valid=3s;
   upstream backend_up {
-    server backend-headless.highload-dns.svc.cluster.local:8000;
+    zone upstream_dynamic 64k;
+    server backend-headless.highload-dns.svc.cluster.local:8000 resolve;
   }
   server {
     listen 80;
     location / {
       proxy_pass http://backend_up;
       add_header X-Upstream-Addr $upstream_addr always;
     }
   }
}
```

#### Проверка инсталляции через скрипт имитации нагрузки

- запущен скрипт имитации нагрузки `./scripts/load-test.sh`
- вывод работы скрипта:

    ```sh
    # скрипт запущен, запросы чередуются между 2мя инстансами PostgreSQL и 2мя инстансами бекенда
    [01] 10:17:45 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
    [02] 10:17:45 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
    [03] 10:17:46 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
    [04] 10:17:47 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
    [05] 10:17:47 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
    [06] 10:17:48 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
    [07] 10:17:48 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
    [08] 10:17:49 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
    [09] 10:17:49 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
    [10] 10:17:50 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
    ```

#### Проверка отказоустойчивости инсталляции через скрипт имитации различных сценариях сбоев

Для проверки запущен скрипт имитации сбоев `./scripts/demo_failover.sh`. Ниже находится описание каждого этапа тестирования с выводом логов.

##### 1) PHASE 0: baseline (2 backend, 2 replicas)

На этом этапе в инсталляции есть 2 реплики backend и 2 реплики postgres.

Запросы к сервису успешно балансируются между 2мя бекендами `10.42.1.9` и `10.42.0.8`, а внутри них запросы к БД балансируются между 2мя репликами postgres `pg1` и `pg2`.

Вывод работы скрипта:

```sh
======= PHASE 0: baseline (2 backend, 2 replicas) ========

[01] 10:17:58 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[02] 10:17:58 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[03] 10:17:59 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[04] 10:17:59 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[05] 10:18:00 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[06] 10:18:00 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[07] 10:18:00 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[08] 10:18:01 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[09] 10:18:01 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[10] 10:18:01 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[11] 10:18:02 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[12] 10:18:02 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
```

##### 2) PHASE A: disable ONE PG replica (scale to 1)

На этом этапе происходит отключение одной реплики postgres.

По выводу логов мы видим, что запросы к сервису продолжают балансироваться между двумя репликами бекенда `10.42.1.9` и `10.42.0.8`, но запросы к БД все начинают уходить на единственную реплику postgres `pg1`.

Вывод работы скрипта:

```sh
======= PHASE A: disable ONE PG replica (scale to 1) ========

deployment.apps/postgres-replicas scaled
NAME                ENDPOINTS        AGE
postgres-replicas   10.42.2.3:5432   9m38s
[01] 10:18:08 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[02] 10:18:08 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[03] 10:18:08 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[04] 10:18:09 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[05] 10:18:09 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[06] 10:18:09 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[07] 10:18:10 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[08] 10:18:10 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[09] 10:18:11 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[10] 10:18:11 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[11] 10:18:11 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[12] 10:18:12 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[13] 10:18:12 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[14] 10:18:12 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[15] 10:18:13 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[16] 10:18:13 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[17] 10:18:13 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[18] 10:18:14 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[19] 10:18:14 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[20] 10:18:15 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
```

##### 3) PHASE A.restore: replicas back to 2

На этом этапе для postgres возвращается вторая реплика.

По выводу логов мы видим, что после восстановления второй реплики postgres, запросы к БД продолжают уходить только на реплику `pg1`, но через время DNS сессия обновляется и запросы от бекенда начинают балансироваться между `pg1` и `pg2`.

Вывод работы скрипта:

```sh
======= PHASE A.restore: replicas back to 2 ========

deployment.apps/postgres-replicas scaled
NAME                ENDPOINTS                       AGE
postgres-replicas   10.42.1.11:5432,10.42.2.3:5432   9m50s
[01] 10:18:20 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[02] 10:18:20 | backend=10.42.0.8:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[03] 10:18:21 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[04] 10:18:21 | backend=10.42.0.8:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[05] 10:18:21 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[06] 10:18:22 | backend=10.42.0.8:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[07] 10:18:22 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[08] 10:18:23 | backend=10.42.0.8:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[09] 10:18:23 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[10] 10:18:23 | backend=10.42.0.8:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
[11] 10:18:24 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[12] 10:18:24 | backend=10.42.0.8:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-kkrck"}
```

##### 4) PHASE B: disable ONE backend (scale to 1)

На этом этапе происходит отключение одной реплики бекенда.

По выводу логов мы видим, что запросы к сервису, после неудачного запроса к реплике `10.42.0.8`, начинают уходить только на одну реплику `10.42.1.9`, но запросы к БД продолжают успешно балансироваться между `pg1` и `pg2`.

Вывод работы скрипта:

```sh
======= PHASE B: disable ONE backend (scale to 1) ========

deployment.apps/backend scaled
NAME               ENDPOINTS        AGE
backend-headless   10.42.1.9:8000   9m59s
[01] 10:18:30 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[02] 10:19:30 | backend=10.42.0.8:8000 | db=unknown | body=<html>
<head><title>504 Gateway Time-out</title></head>
<body>
<center><h1>504 Gateway Time-out</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[03] 10:19:30 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[04] 10:19:31 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[05] 10:19:31 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[06] 10:19:31 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[07] 10:19:32 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[08] 10:19:32 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[09] 10:19:32 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[10] 10:19:33 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[11] 10:19:33 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[12] 10:19:34 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[13] 10:19:34 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[14] 10:19:34 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[15] 10:19:35 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[16] 10:19:35 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[17] 10:19:35 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[18] 10:19:36 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[19] 10:19:36 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[20] 10:19:36 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
```

##### 5) PHASE B.restore: backend back to 2

На этом этапе для бекенда возвращается вторая реплика.

По выводу логов мы видим, что после восстановления второй реплики бекенда, запросы к сервису продолжают балансироваться между репликой бекенда `10.42.1.9` и новой репликой `10.42.0.9`.

Вывод работы скрипта:

```sh
======= PHASE B.restore: backend back to 2 ========

deployment.apps/backend scaled
NAME               ENDPOINTS        AGE
backend-headless   10.42.0.9:8000,10.42.1.9:8000   11m
[01] 10:19:52 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[02] 10:19:52 | backend=10.42.0.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kstdb"}
[03] 10:19:53 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[04] 10:19:53 | backend=10.42.0.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kstdb"}
[05] 10:19:53 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[06] 10:19:54 | backend=10.42.0.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kstdb"}
[07] 10:19:54 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[08] 10:19:55 | backend=10.42.0.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kstdb"}
[09] 10:19:55 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[10] 10:19:55 | backend=10.42.0.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kstdb"}
[11] 10:19:56 | backend=10.42.1.9:8000 | db=pg2 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[12] 10:19:56 | backend=10.42.0.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-kstdb"}
```

##### 6) PHASE C: 1 backend + 1 PG replica, then delete the ONLY backend pod (observe errors -> recovery)

На этом этапе остается только одна реплика бекенда и одна реплика postgres.

По выводу логов мы видим, что произошло скаллирование деплойментов бекенда и postgres, после чего все запросы уходят только на одну доступную реплику бекенда `10.42.1.9` и одну доступную реплику postgres `pg1`.

Вывод работы скрипта:

```sh
======= PHASE C: 1 backend + 1 PG replica, then delete the ONLY backend pod (observe errors -> recovery) ========

deployment.apps/backend scaled
deployment.apps/postgres-replicas scaled
# Endpoints now:
NAME               ENDPOINTS        AGE
backend-headless   10.42.1.9:8000   11m
NAME                ENDPOINTS        AGE
postgres-replicas   10.42.2.3:5432   11m
[01] 10:20:02 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[02] 10:20:02 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[03] 10:20:02 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[04] 10:20:03 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[05] 10:20:03 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[06] 10:20:03 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[07] 10:20:04 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[08] 10:20:04 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[09] 10:20:04 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
[10] 10:20:05 | backend=10.42.1.9:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-6bt42"}
```

##### 7) PHASE C.drop: deleting the only backend pod

На этом этапе удаляется одна доступна реплика бекенда, после чего ReplicaSet автоматически поднимает новую реплику.

По выводу логов мы видим затяжные ошибки 502 запросов к бекенду, пока происходит восстановление реплики бекенда и обновление DNS на nginx. На последних интерациях теста видно, что реплика успешно поднялась и DNS обновился, после чего запросы успешно обрабатываются новой репликой бекенда `10.42.0.10`.

Вывод работы скрипта:

```sh
======= PHASE C.drop: deleting the only backend pod ========

Deleting pod/backend-c65c887b9-6bt42 ...
pod "backend-c65c887b9-6bt42" deleted from highload-dns namespace
[01] 10:20:10 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[02] 10:20:11 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[03] 10:20:11 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[04] 10:20:11 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[05] 10:20:12 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[06] 10:20:12 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[07] 10:20:12 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[08] 10:20:13 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[09] 10:20:13 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[10] 10:20:13 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[11] 10:20:14 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[12] 10:20:14 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[13] 10:20:15 | backend=backend_up | db=unknown | body=<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.27.5</center>
</body>
</html>
[14] 10:20:15 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
[15] 10:20:15 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
```

##### 8) PHASE C.restore: waiting for backend rollout, then verify

На этом этапе мы видим успешную работу одной реплики бекенда и одной реплики postgres после восстановления.

Вывод работы скрипта:

```sh
======= PHASE C.restore: waiting for backend rollout, then verify ========

deployment "backend" successfully rolled out
NAME               ENDPOINTS         AGE
backend-headless   10.42.0.10:8000   11m
[01] 10:20:18 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
[02] 10:20:18 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
[03] 10:20:19 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
[04] 10:20:19 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
[05] 10:20:19 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
[06] 10:20:20 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
[07] 10:20:20 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
[08] 10:20:20 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
[09] 10:20:21 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
[10] 10:20:21 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
[11] 10:20:21 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}
[12] 10:20:22 | backend=10.42.0.10:8000 | db=pg1 | body={"users":3,"pod":"backend-c65c887b9-98gdb"}


======= DONE ========
```
