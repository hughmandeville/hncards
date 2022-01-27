// Return time since in human readable format (e.g. "2 hrs", "on Sat").
const timeSince = ts => {
    var seconds = Math.floor((new Date() - ts * 1000) / 1000);
  
    var interval = seconds / 31536000;
  
    if (interval > 1) {
      if (Math.floor(interval) === 1) {
        return Math.floor(interval) + " yr";
      }
      return Math.floor(interval) + " yrs";
    }
    interval = seconds / 2592000;
    if (interval > 1) {
      if (Math.floor(interval) === 1) {
        return Math.floor(interval) + " mon";
      }
      return Math.floor(interval) + " mons";
    }
    interval = seconds / 86400;
    if (interval > 1) {
      // If 2 to 5 days ago, show day of week (e.g. "on Sat").
      if (Math.floor(interval) >= 2 && Math.floor(interval) <= 5) {
        const d = new Date(ts * 1000);
        return "on " + d.toLocaleDateString("en-US", { weekday: "short" });
      }
      return Math.floor(interval) + " d";
    }
    interval = seconds / 3600;
    if (interval > 1) {
      if (Math.floor(interval) === 1) {
        return "1 hr";
      }
      return Math.floor(interval) + " hrs";
    }
    interval = seconds / 60;
    if (interval > 1) {
      return Math.floor(interval) + " mins";
    }
    return Math.floor(seconds) + " secs";
  }

  export { timeSince };