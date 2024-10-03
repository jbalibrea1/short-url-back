import { Request, RequestHandler, Response } from 'express';
import shortUrl from '../services/shortUrl';
import handleError from '../utils/handleError';

const getAllShortURLs = (async (_req: Request, res: Response) => {
  try {
    const allUrls = await shortUrl.getAll();
    res.json(allUrls);
  } catch (error) {
    handleError(res, 'Failed to get all short URLs.', error);
  }
}) as RequestHandler;

const addShortURL = (async (
  req: Request<unknown, unknown, { url: string }>,
  res: Response
) => {
  try {
    console.log(req.body);
    const { url }: { url: string } = req.body;
    const newShortUrlEntry = await shortUrl.create(url);
    res.json(newShortUrlEntry);
  } catch (error) {
    handleError(res, 'Failed to add short URL.', error);
  }
}) as RequestHandler;

const getShortURL = (async (req: Request, res: Response) => {
  try {
    const url = req.params.shortUrl;
    const entry = await shortUrl.get(url);
    res.json(entry);
  } catch (error) {
    handleError(res, 'Failed to get full URL.', error);
  }
}) as RequestHandler;

const getRedirect = (async (req: Request, res: Response) => {
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
}) as RequestHandler;

export default {
  getAllShortURLs,
  addShortURL,
  getShortURL,
  getRedirect,
};
