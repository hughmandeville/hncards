import "./NavBar.css";

// NavBar - Render nav bar.
const NavBar = props => (
  <div id="nav">
    <img id="nav-logo" src="/ycombinator.png" alt="Y Combinator" />
    <div id="nav-title">
      <a href="https://news.ycombinator.com/news">Hacker News</a>
    </div>
  </div>
);


export default NavBar;
