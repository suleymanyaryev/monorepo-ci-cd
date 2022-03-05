docker compose -f docker-compose.backend-test.yml build
docker compose -f docker-compose.backend-test.yml up --remove-orphans --abort-on-container-exit  --exit-code-from service
docker compose -f docker-compose.backend-test.yml down --volume 

docker compose -f docker-compose.frontend-test.yml build
docker compose -f docker-compose.frontend-test.yml up --remove-orphans --abort-on-container-exit  --exit-code-from service
docker compose -f docker-compose.frontend-test.yml down --volume

docker compose -f docker-compose.e2e-test.yml build
docker compose -f docker-compose.e2e-test.yml up --remove-orphans --abort-on-container-exit  --exit-code-from frontend
docker compose -f docker-compose.e2e-test.yml down