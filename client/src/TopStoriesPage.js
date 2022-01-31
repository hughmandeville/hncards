import React, { useEffect, useState } from 'react';
import ItemBoxes from './ItemBoxes';
import NavBar from './NavBar';
import './TopStoriesPage.css';

// TopStoriesPage - Render Top Stories page.
const TopStoriesPage = () => {
  const [error, setError] = useState();
  const [items, setItems] = useState();

  let hnURL =
    'https://raw.githubusercontent.com/hughmandeville/hnui/main/client/public/hn_topstories.json';
  if (window.location.hostname === 'localhost' || window.location.hostname === '127.0.0.1') {
    hnURL = '/hnui/hn_topstories.json';
  }

  const fetchData = () => {
    console.log("fetching data")
    fetch(hnURL)
    .then(response => response.json())
    .then(items => {
      console.log("Fetched Hacker News data: ", items);
      setItems(items);
    })
    .catch(error => {
      console.log('Error fetching Hacker News data: ', error);
      setError(error);
    });
  }

  // NOTE: set dependency array in useEffect to empty array to prevent fetchData getting called over and over.
  // It causes React to warn about missing dependency.
  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div id="ts-page">
      <NavBar />
      <div>
        {error ? (
          <div className="error">
            <img src="/hnui/meh.png" alt="meh" />
            <p>
              There was a problem getting the <a href={hnURL}>top stories data</a>.
            </p>
            <p>
              Please try again later or view the{' '}
              <a href="https://news.ycombinator.com/">Hacker News site</a>.
            </p>
            <p>&bull;</p>
            <p>
              <i>{error.toString()}</i>
            </p>
          </div>
        ) : (
          <div className="boxes">
            <ItemBoxes items={items} />
          </div>
        )}
      </div>
      <div id="footer">hnards.com is an unofficial alternative <a href="https://news.ycombinator.com/">Hacker News</a> UI.</div>
    </div>
  );
};

export default TopStoriesPage;
