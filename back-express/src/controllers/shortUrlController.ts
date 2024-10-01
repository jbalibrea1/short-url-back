import { Request, Response } from 'express';
import shortUrl from '../services/shortUrl';
import handleError from '../utils/handleError';

const getAllShortURLs = async (_req: Request, res: Response): Promise<void> => {
  try {
    const allUrls = await shortUrl.getAll();
    res.json(allUrls);
  } catch (error) {
    handleError(res, 'Failed to get all short URLs.', error);
  }
};

const addShortURL = async (req: Request, res: Response): Promise<void> => {
  try {
    const { url }: { url: String } = req.body;
    const newShortUrlEntry = await shortUrl.create(url);
    res.json(newShortUrlEntry);
  } catch (error) {
    handleError(res, 'Failed to add short URL.', error);
  }
};

const getShortURL = async (req: Request, res: Response): Promise<void> => {
  try {
    const url = req.params.shortUrl;
    const entry = await shortUrl.get(url);
    res.json(entry);
  } catch (error) {
    handleError(res, 'Failed to get full URL.', error);
  }
};

const getRedirect = async (req: Request, res: Response): Promise<void> => {
  try {
    const url = req.params.shortUrl;
    const entry = await shortUrl.get(url);
    if (!entry.url) {
      throw new Error('No URL found for the given short URL');
    }
    res.redirect(entry.url);
  } catch (error) {
    handleError(res, 'Failed to get full URL.', error);
  }
};

export default {
  getAllShortURLs,
  addShortURL,
  getShortURL,
  getRedirect,
};
