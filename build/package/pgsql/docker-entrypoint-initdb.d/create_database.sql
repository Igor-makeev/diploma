SELECT 'CREATE DATABASE "secretKeeper"'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'secretKeeper')\gexec
