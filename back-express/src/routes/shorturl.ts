import express from 'express';
import shortUrlController from '../controllers/shortUrlController';

const router = express.Router();

router.get('/', shortUrlController.getAllShortURLs);
router.post('/', shortUrlController.addShortURL);
router.get('/:shortUrl', shortUrlController.getShortURL);

export { router };
