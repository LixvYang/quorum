#!/bin/bash

curl -X POST -H 'Content-Type: application/json' -d '{"seed": "rum://seed?v=1\u0026e=0\u0026n=0\u0026c=G7VuCagx9uj-39lULSfCeXk6cVK8-rySGHKZKMFzHGA\u0026g=7wyAnC6rQaKH_uxN57WoVQ\u0026k=Az1h8rhiLGC4u991wubZ55EtLJz0NzEjYc4K2Wc_DnD1\u0026s=M2F0hnk_jxFQSnjuZvMgWUhFXmyIERFEeWLiCHFbWRx_qEdpRZaan2zak6QpORMSPL8MJeMQOtMyvG7oM6DEtQA\u0026t=FzZDmXoI19M\u0026a=my_test_group\u0026y=test_app\u0026u=http%3A%2F%2F127.0.0.1%3A8002%3Fjwt%3DeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhbGxvd0dyb3VwcyI6WyJlZjBjODA5Yy0yZWFiLTQxYTItODdmZS1lYzRkZTdiNWE4NTUiXSwiZXhwIjoxODMwMjc4NjM4LCJuYW1lIjoiYWxsb3ctZWYwYzgwOWMtMmVhYi00MWEyLTg3ZmUtZWM0ZGU3YjVhODU1Iiwicm9sZSI6Im5vZGUifQ.j4sN9g5reJcQ6dkQLZ-8B6AORclcEUwDyUTdwalYA_g"}' http://127.0.0.1:8002/api/v2/group/join | jq
curl -X POST -H 'Content-Type: application/json' -d '{"seed": "rum://seed?v=1\u0026e=0\u0026n=0\u0026c=G7VuCagx9uj-39lULSfCeXk6cVK8-rySGHKZKMFzHGA\u0026g=7wyAnC6rQaKH_uxN57WoVQ\u0026k=Az1h8rhiLGC4u991wubZ55EtLJz0NzEjYc4K2Wc_DnD1\u0026s=M2F0hnk_jxFQSnjuZvMgWUhFXmyIERFEeWLiCHFbWRx_qEdpRZaan2zak6QpORMSPL8MJeMQOtMyvG7oM6DEtQA\u0026t=FzZDmXoI19M\u0026a=my_test_group\u0026y=test_app\u0026u=http%3A%2F%2F127.0.0.1%3A8002%3Fjwt%3DeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhbGxvd0dyb3VwcyI6WyJlZjBjODA5Yy0yZWFiLTQxYTItODdmZS1lYzRkZTdiNWE4NTUiXSwiZXhwIjoxODMwMjc4NjM4LCJuYW1lIjoiYWxsb3ctZWYwYzgwOWMtMmVhYi00MWEyLTg3ZmUtZWM0ZGU3YjVhODU1Iiwicm9sZSI6Im5vZGUifQ.j4sN9g5reJcQ6dkQLZ-8B6AORclcEUwDyUTdwalYA_g"}' http://127.0.0.1:8003/api/v2/group/join | jq
curl -X POST -H 'Content-Type: application/json' -d '{"seed": "rum://seed?v=1\u0026e=0\u0026n=0\u0026c=G7VuCagx9uj-39lULSfCeXk6cVK8-rySGHKZKMFzHGA\u0026g=7wyAnC6rQaKH_uxN57WoVQ\u0026k=Az1h8rhiLGC4u991wubZ55EtLJz0NzEjYc4K2Wc_DnD1\u0026s=M2F0hnk_jxFQSnjuZvMgWUhFXmyIERFEeWLiCHFbWRx_qEdpRZaan2zak6QpORMSPL8MJeMQOtMyvG7oM6DEtQA\u0026t=FzZDmXoI19M\u0026a=my_test_group\u0026y=test_app\u0026u=http%3A%2F%2F127.0.0.1%3A8002%3Fjwt%3DeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhbGxvd0dyb3VwcyI6WyJlZjBjODA5Yy0yZWFiLTQxYTItODdmZS1lYzRkZTdiNWE4NTUiXSwiZXhwIjoxODMwMjc4NjM4LCJuYW1lIjoiYWxsb3ctZWYwYzgwOWMtMmVhYi00MWEyLTg3ZmUtZWM0ZGU3YjVhODU1Iiwicm9sZSI6Im5vZGUifQ.j4sN9g5reJcQ6dkQLZ-8B6AORclcEUwDyUTdwalYA_g"}' http://127.0.0.1:8004/api/v2/group/join | jq
#curl -X POST -H 'Content-Type: application/json' -d '{"seed": "rum://seed?v=1\u0026e=0\u0026n=0\u0026c=G7VuCagx9uj-39lULSfCeXk6cVK8-rySGHKZKMFzHGA\u0026g=7wyAnC6rQaKH_uxN57WoVQ\u0026k=Az1h8rhiLGC4u991wubZ55EtLJz0NzEjYc4K2Wc_DnD1\u0026s=M2F0hnk_jxFQSnjuZvMgWUhFXmyIERFEeWLiCHFbWRx_qEdpRZaan2zak6QpORMSPL8MJeMQOtMyvG7oM6DEtQA\u0026t=FzZDmXoI19M\u0026a=my_test_group\u0026y=test_app\u0026u=http%3A%2F%2F127.0.0.1%3A8002%3Fjwt%3DeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhbGxvd0dyb3VwcyI6WyJlZjBjODA5Yy0yZWFiLTQxYTItODdmZS1lYzRkZTdiNWE4NTUiXSwiZXhwIjoxODMwMjc4NjM4LCJuYW1lIjoiYWxsb3ctZWYwYzgwOWMtMmVhYi00MWEyLTg3ZmUtZWM0ZGU3YjVhODU1Iiwicm9sZSI6Im5vZGUifQ.j4sN9g5reJcQ6dkQLZ-8B6AORclcEUwDyUTdwalYA_g"}' http://127.0.0.1:8006/api/v2/group/join | jq

sleep 3
curl -X POST -H 'Content-Type: application/json' -d '{"group_id":"ef0c809c-2eab-41a2-87fe-ec4de7b5a855", "action":"add", "type":"producer", "memo":"producer p1, realiable and cheap, online 24hr"}' http://127.0.0.1:8003/api/v1/group/announce | jq
curl -X POST -H 'Content-Type: application/json' -d '{"group_id":"ef0c809c-2eab-41a2-87fe-ec4de7b5a855", "action":"add", "type":"producer", "memo":"producer p2, realiable and cheap, online 24hr"}' http://127.0.0.1:8004/api/v1/group/announce | jq

sleep 3
curl -X POST -H 'Content-Type: application/json' -d '{"producer_pubkey":["Ap4s80oy_TTcx4yemN7dlND1QTdiZlgkcsWExZDkNQhS","AzG_tYeMh3n3_p05V8fJvco4n9Fgl5Wikzp7ccTqXNHL"] ,"group_id":"ef0c809c-2eab-41a2-87fe-ec4de7b5a855", "action":"add"}' http://127.0.0.1:8002/api/v1/group/producer/false