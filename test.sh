docker compose -f docker-compose.backend-test.yml build
docker compose -f docker-compose.backend-test.yml up --abort-on-container-exit  --exit-code-from service
docker compose -f docker-compose.backend-test.yml down --volume --remove-orphans

docker compose -f docker-compose.frontend-test.yml build
docker compose -f docker-compose.frontend-test.yml up --abort-on-container-exit  --exit-code-from service
docker compose -f docker-compose.frontend-test.yml down --volume --remove-orphans
