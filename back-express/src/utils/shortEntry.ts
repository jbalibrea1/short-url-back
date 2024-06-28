import { NewShortUrlEntry } from '../types/shortUrlTypes';

const parseUrl = (url: unknown): string => {
  if (!url || typeof url !== 'string') {
    throw new Error('Incorrect or missing url');
  }

  if (!url.startsWith('http://') && !url.startsWith('https://')) {
    return 'http://' + url;
  }

  return url;
};

const toNewShortUrlEntry = (object: unknown): NewShortUrlEntry => {
  if (!object || typeof object !== 'object') {
    throw new Error('Incorrect or missing data');
  }

  if ('url' in object) {
    const newEntry: NewShortUrlEntry = {
      url: parseUrl(object.url),
    };
    return newEntry;
  }

  throw new Error('Incorrect data: a field missing');
};

export default toNewShortUrlEntry;
