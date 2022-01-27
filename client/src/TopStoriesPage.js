import React, { useEffect, useState }  from "react";
import ItemBoxes from "./ItemBoxes";
import NavBar from "./NavBar";
import "./TopStoriesPage.css";

// TopStoriesPage - Render Top Stories page.
const TopStoriesPage = () => {
  const [error, setError] = useState();
  const [items, setItems] = useState();

  useEffect(() => {
    fetch(
      "https://raw.githubusercontent.com/hughmandeville/hnui/main/client/public/hn_topstories.json"
    )
      .then((response) => response.json())
      .then((items) => {
        setItems(items);
      })
      .catch((error) => {
        setError(error);
      });
  });

  return (
    <div id="ts-page">
      <NavBar />
      <div>{error ? <div className="error">{error}</div> : <div className="boxes">
        <ItemBoxes items={items} />
      </div>}</div>
    </div>)
}

export default TopStoriesPage;
