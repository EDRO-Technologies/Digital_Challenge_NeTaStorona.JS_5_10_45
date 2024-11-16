import express from "express";
import bcrypt from 'bcryptjs'
import jwt from 'jsonwebtoken'

import {db} from "../db.js";

var router = express.Router();

const TABLE = 'public."user"'

/* GET users listing. */
router.post('/login', async (req, res, next) => {
    const {email, password} = req.body;

    try {
        const {rows} = await db.query(`SELECT * FROM ${TABLE} WHERE email = $1`, [email]);

        if (rows.length === 0) {
            return res.status(401).send('Invalid username or password');
        }

        const user = rows[0];

        const isPassCorrect = await bcrypt.compare(password, user.password);

        if (!isPassCorrect) {
            return res.status(401).send('No user or wrong password');
        }

        user.password = undefined;
        const token = jwt.sign({user}, process.env.JWT_SECRET, {expiresIn: '40h'});

        res.json({
            token,
            user: {
                uuid: user.uuid,
                username: user.fio,
                email: user.email,
            }
        });
    } catch (error) {
        console.error(error);
        res.status(500).json({error: true, msg: error})
    }
});

// Регистрация пользователя
router.post('/register', async (req, res) => {
    if (!req.body || Object.keys(req.body).length === 0) {
        res.status(400).json({error: true, msg: 'Пусто в JSON :('})
        return;
    }

    const {email, password, fio, birthdate} = req.body;

    ['email', 'password', 'fio', 'birthdate'].forEach((value) => {
        if (!req.body[value]) {
            return res.status(400).json({error: true, msg: `Не введено поле ${value}!`})
        }
    })


    try {
        // Хеширование пароля
        const hashedPassword = await bcrypt.hash(password, 10);

        const {rows: [{count}]} = await db.query(`SELECT COUNT(*) FROM ${TABLE} WHERE email = $1`, [email]);

        if (+count) {
            return res.status(403).json({error: true, msg: `Пользователь с email ${email} уже существует!`})
        }

        // Вставка нового пользователя в базу данных
        const result = await db.query(
            `INSERT INTO ${TABLE} (email, password, fio, birthdate) VALUES ($1, $2, $3, $4) RETURNING *`,
            [email, hashedPassword, fio, birthdate]
        );

        const {password: _, ...user} = result.rows[0];

        // Генерация JWT токена
        const token = jwt.sign({user}, process.env.JWT_SECRET, {expiresIn: '40h'});

        res.json({token, user,});
    } catch (error) {
        console.error(error);
        res.status(500).json({error: true, msg: error})
    }
});

router.get('/me', async (req, res) => {
    try {
        const test = req.header('Authorization');

        const [_, token] = test.split(' ');

        const me = jwt.decode(token);

        if (!me) {
            return res.status(403).send('Bad token');
        }

        return res.status(200).json(me)
    } catch (error) {
        console.error(error);
        res.status(500).json({error: true, msg: error})
    }
})

export default router;
