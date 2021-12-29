psql -v ON_ERROR_STOP=1 --username "course" --dbname "course" <<-EOSQL
    CREATE TABLE users (id BIGSERIAL, name VARCHAR);
    INSERT INTO users (name) VALUES ('first'), ('second');
EOSQL