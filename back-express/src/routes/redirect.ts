import express from 'express';
import shortUrlController from '../controllers/shortUrlController';

const router = express.Router();

router.get('/:shortUrl', shortUrlController.getRedirect);

export { router };
