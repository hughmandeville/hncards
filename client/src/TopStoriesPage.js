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
    return <div id="ts-page">TBD</div>;
  }
}

export { TopStoriesPage };
