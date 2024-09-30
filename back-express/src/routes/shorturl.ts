import express from 'express';
import shortUrlController from '../controllers/shortUrlController';

const router = express.Router();

router.get('/', shortUrlController.getAllShortUrls);
router.post('/', shortUrlController.addShortUrl);
router.get('/:shortUrl', shortUrlController.getFullUrl);

export { router };
