import { Response } from 'express';
import logger from '../utils/logger';

const requestLogger = (
  request: { method: unknown; path: unknown; body: unknown },
  _response: Response,
  next: () => void
) => {
  logger.info('Method:', request.method);
  logger.info('Path:  ', request.path);
  logger.info('Body:  ', request.body);
  logger.info('---');
  next();
};

const unknownEndpoint = (_: unknown, res: Response) => {
  res.status(404).send({ error: 'unknown endpoint' });
};

const errorMidHandler = (
  error: { name: string; message: string },
  _request: unknown,
  response: Response,
  next: (arg0: unknown) => void
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
