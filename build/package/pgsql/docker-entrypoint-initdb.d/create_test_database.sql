SELECT 'CREATE DATABASE "secretKeeper_test"'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'secretKeeper_test')\gexec
