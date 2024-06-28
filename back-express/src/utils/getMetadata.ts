import metascraper from 'metascraper';
import metascraperDescription from 'metascraper-description';
import metascraperLogoFavicon from 'metascraper-logo-favicon';
import metascraperTitle from 'metascraper-title';
import metascraperUrl from 'metascraper-url';

const scraper = metascraper([
  // metascraperImage(),
  metascraperTitle(),
  metascraperDescription(),
  metascraperUrl(),
  metascraperLogoFavicon(),
]);

async function getMetadata(url: string) {
  const response = await fetch(url, {
    headers: {
      'Content-Type': 'text/html; charset=utf-8',
    },
  });
  const html = await response.text();
  const metadata = await scraper({ html, url });
  return metadata;
}

export default getMetadata;
