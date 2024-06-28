require('dotenv').config();
import express from 'express';
import connectDB from './config/db';
import middleware from './middleware';
import shortUrlRoutes from './routes/shortUrlRoutes';
const app = express();
const cors = require('cors');

// Middlewares
app.use(cors());
app.use(express.json());
app.use(middleware.requestLogger);

// DB connection
connectDB();

// Routes
app.get('/ping', (_req, res) => {
  console.log('someone pinged here');
  res.send('pong');
});
app.use('/shorturl', shortUrlRoutes);

// Middleware for handling unknown routes and errors
app.use(middleware.unknownEndpoint);
app.use(middleware.errorHandler);

export default app;
