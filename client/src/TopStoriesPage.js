import React, { useEffect, useState, useRef } from 'react';
import ItemBoxes from './ItemBoxes';
import NavBar from './NavBar';
import './TopStoriesPage.css';

// TopStoriesPage - Render Top Stories page.
const TopStoriesPage = () => {
  const [error, setError] = useState();
  const [items, setItems] = useState();
  const [count, setCount] = React.useState(0);
  const someRef = useRef(0);

  let hnURL = 'https://storage.googleapis.com/hncards/hn_topstories.json';
  if (window.location.hostname === 'localhost' || window.location.hostname === '127.0.0.1') {
    hnURL = '/hn_topstories.json';
  }

  const fetchIntervalInSeconds = 600; // 5 minutes

  const tick = () => {
    //console.log('timer: ', count < fetchIntervalInSeconds ? count + 1 : 0);
    setCount(prevState => (prevState < fetchIntervalInSeconds ? prevState + 1 : 0));
  };

  useEffect(() => {
    const timer = setInterval(() => tick(), 1000);
    return () => clearInterval(timer);
  });

  useEffect(() => {
    if (count === 0) {
      someRef.current = someRef.current + 1;
      setTimeout(() => {
        fetch(hnURL)
          .then(response => response.json())
          .then(items => {
            console.log('Fetched Hacker News data: ', items);
            setItems(items);
          })
          .catch(error => {
            console.log('Error fetching Hacker News data: ', error);
            setError(error);
          });
      }, 1000);
    }
  }, [count, hnURL]);

  return (
    <div id="ts-page">
      <NavBar />
      <div>
        {error ? (
          <div className="error">
            <img src="/meh.png" alt="meh" />
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
      <div id="footer">
        <a href="https://hncards.com/">hncards.com</a> is an unofficial alternative{' '}
        <a href="https://news.ycombinator.com/">Hacker News</a> UI.
      </div>
    </div>
  );
};

export default TopStoriesPage;
