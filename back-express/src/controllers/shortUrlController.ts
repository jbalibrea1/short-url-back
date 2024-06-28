import { Request, Response } from 'express';
import ShortUrl from '../models/shortUrl';
import getMetadata from '../utils/getMetadata';
import generateUniqueShortUrl from '../utils/randomUrl';
import toNewShortUrlEntry from '../utils/shortEntry';
import truncateString from '../utils/truncateString';

const getAllShortUrls = async (_req: Request, res: Response): Promise<void> => {
  try {
    const allUrls = await ShortUrl.find({});
    res.json(allUrls);
  } catch (error) {
    res.status(400).send({ errorMessage: handleError(error) });
  }
};

const addShortUrl = async (req: Request, res: Response): Promise<void> => {
  try {
    const newShortUrlEntry = toNewShortUrlEntry(req.body);

    //get metadata from  url
    const metadata = await getMetadata(newShortUrlEntry.url);
    console.log('metada', metadata);
    let shortUrl = generateUniqueShortUrl();
    while (await ShortUrl.findOne({ shortUrl })) {
      shortUrl = generateUniqueShortUrl();
    }

    if (metadata.title) {
      newShortUrlEntry.title = metadata.title;
    }

    if (metadata.image) {
      newShortUrlEntry.image = metadata.image;
    }
    if (metadata.logo) {
      newShortUrlEntry.logo = metadata.logo;
    }
    if (metadata.description) {
      newShortUrlEntry.description = truncateString(metadata.description, 50);
    }

    const newEntry = new ShortUrl({
      ...newShortUrlEntry,
      shortUrl,
    });

    const savedEntry = await newEntry.save();

    res.json(savedEntry);
  } catch (error) {
    res.status(400).send({ errorMessage: handleError(error) });
  }
};

const getFullUrl = async (req: Request, res: Response): Promise<void> => {
  try {
    const shortUrl = req.params.shortUrl;
    const entry = await ShortUrl.findOne({ shortUrl });

    if (!entry) {
      res
        .status(404)
        .send({ errorMessage: `Short URL not found for ${shortUrl}` });
      return;
    }

    entry.totalClicks += 1;
    await entry.save();
    // res.redirect(301, entry.url);
    res.send({ url: entry.url });
  } catch (error) {
    res.status(400).send({ errorMessage: handleError(error) });
  }
};

//![remove] Helper function
const handleError = (error: unknown): string => {
  let errorMessage = 'Something went wrong.';
  if (error instanceof Error) {
    errorMessage += ' Error: ' + error.message;
  }
  return errorMessage;
};

export default {
  addShortUrl,
  getAllShortUrls,
  getFullUrl,
};
