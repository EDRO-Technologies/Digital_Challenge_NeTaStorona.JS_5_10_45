import express from 'express';
import path from 'path';
import logger from 'morgan';
import cookieParser from 'cookie-parser';
import cors from 'cors';

import indexRoutes from './routes/index.js';
import userRoutes from './routes/users.js';
import authRoutes from './routes/auth.js'

const app = express();

const corsOptions = {
    origin: '*', // Разрешаем запросы только с этого домена
    methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
    credentials: true,
    optionsSuccessStatus: 204,
};

app.use(cors(corsOptions));
app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({extended: false}));
app.use(cookieParser());
// app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRoutes);
app.use('/users', userRoutes);
app.use('/auth', authRoutes);

export default app;
