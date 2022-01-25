import React from "react";
import ItemBoxes from "./ItemBoxes";
import NavBar from "./NavBar";
import "./TopStoriesPage.css";

// TopStoriesPage - Render Top Stories page.
class TopStoriesPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      error: "",
      items: {},
    };
  }
  componentDidMount() {
    fetch("hn_topstories.json")
      .then((response) => response.json())
      .then((items) => {
        this.setState({
          error: "",
          items,
        });
      })
      .catch((error) => {
        this.setState({
          error: error.message,
          items: {},
        });
      });
  }

  render() {
    let outputHTML = "";
    if (this.state.error && this.state.error !== "") {
      outputHTML = <div className="error">{this.state.error}</div>;
    } else {
      outputHTML = (
        <div className="boxes">
          <ItemBoxes items={this.state.items} />
        </div>
      );
    }

    return (
      <div id="ts-page">
        <NavBar />
        <div>{outputHTML}</div>
      </div>
    );
  }
}

export { TopStoriesPage };
