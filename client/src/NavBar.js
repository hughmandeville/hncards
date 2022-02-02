import './NavBar.css';

// NavBar - Render nav bar.
const NavBar = props => (
  <div id="nav">
    <img id="nav-logo" src="/hncards.png" alt="Y Combinator" />
    <h1 id="nav-title">
      <a href="https://news.ycombinator.com/news">Hacker News Cards</a>
    </h1>
    <div id="nav-subtitle">an unofficial ui for hacker news</div>
  </div>
);

export default NavBar;
