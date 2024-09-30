import { NextFunction, Response } from 'express';
import logger from '../utils/logger';

const requestLogger = (
  request: { method: any; path: any; body: any },
  _response: any,
  next: () => void
) => {
  logger.info('Method:', request.method);
  logger.info('Path:  ', request.path);
  logger.info('Body:  ', request.body);
  logger.info('---');
  next();
};

const unknownEndpoint = (_request: any, response: any) => {
  response.status(404).send({ error: 'unknown endpoint' });
};

const errorMidHandler = (
  error: { name: string; message: any },
  _request: Request,
  response: Response,
  next: NextFunction
) => {
  console.log('errorHandler error:', error);
  if (error.name === 'CastError') {
    return response.status(400).send({
      error: 'malformatted id',
    });
  } else if (error.name === 'ValidationError') {
    return response.status(400).json({
      error: error.message,
    });
  } else if (
    error.name === 'MongoServerError' &&
    error.message.includes('E11000 duplicate key error')
  ) {
    return response
      .status(400)
      .json({ error: 'expected `username` to be unique' });
  }

  logger.error(error.message);

  next(error);
  return;
};

export default {
  requestLogger,
  unknownEndpoint,
  errorMidHandler,
};
