import { Router } from 'express';
import { readdirSync } from 'fs';

const PATH_ROUTER = `${__dirname}`;
const router = Router();

/**
 * shortUrl.ts -> shortUrl
 * @returns {string|undefined} - The file name without extension
 */
const cleanFileName = (fileName: string): string | undefined => {
  const file = fileName.split('.').shift();
  return file;
};

readdirSync(PATH_ROUTER).filter((fileName) => {
  const cleanName = cleanFileName(fileName);
  if (cleanName !== 'index') {
    import(`./${cleanName}`).then((moduleRouter) => {
      router.use(`/api/${cleanName}`, moduleRouter.router);
    });
    console.info(`Router ------> /${cleanName} loaded`);
  }
});

export { router };
