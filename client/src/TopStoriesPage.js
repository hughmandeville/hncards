import React from "react";
import "./TopStoriesPage.css";

// TopStoriesPage - Render Top Stories page.
class TopStoriesPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      showMenu: false,
    };
  }
  componentDidMount() {}

  render() {
    return (
      <div id="ts-page">
        <div id="nav">
          <img id="nav-logo" src="ycombinator.png" alt="Y Combinator" />
          <div id="nav-title">
            <a href="https://news.ycombinator.com/news">Hacker News</a>
          </div>
        </div>
      </div>
    );
  }
}

export { TopStoriesPage };
