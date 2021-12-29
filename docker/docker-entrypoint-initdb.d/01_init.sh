psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER course WITH PASSWORD 'course';
    CREATE DATABASE course;
    GRANT ALL PRIVILEGES ON DATABASE course TO course;
EOSQL