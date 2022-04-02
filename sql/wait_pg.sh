until docker-compose exec pgsql psql  -U user -d calendar -c "select 1" > /dev/null 2>&1; do
  >&2 echo "Postgres is unavailable - waiting..."
  sleep 1
done