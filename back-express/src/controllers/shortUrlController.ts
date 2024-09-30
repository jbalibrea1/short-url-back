import { Request, Response } from 'express';
import shortUrl from '../services/shortUrl';
import handleError from '../utils/handleError';

const getAllShortUrls = async (_req: Request, res: Response): Promise<void> => {
  try {
    const allUrls = await shortUrl.getAll();
    res.json(allUrls);
  } catch (error) {
    handleError(res, 'Failed to get all short URLs.', error);
  }
};

const addShortUrl = async (req: Request, res: Response): Promise<void> => {
  try {
    const { url }: { url: String } = req.body;
    const newShortUrlEntry = await shortUrl.create(url);
    res.json(newShortUrlEntry);
  } catch (error) {
    handleError(res, 'Failed to add short URL.', error);
  }
};

const getFullUrl = async (req: Request, res: Response): Promise<void> => {
  try {
    const url = req.params.shortUrl;
    const entry = await shortUrl.get(url);
    res.send(entry);
  } catch (error) {
    handleError(res, 'Failed to get full URL.', error);
  }
};

export default {
  getAllShortUrls,
  addShortUrl,
  getFullUrl,
};
