import pkg from 'pg';

const { Pool } = pkg;

const db = new Pool({
    user: process.env.DB_USER,
    host: process.env.DB_HOST,
    database: 'ugra',
    password: process.env.DB_PASS,
    port: 5432,
});

export { db }